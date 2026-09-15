package mocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/helpers"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
)

var mockState struct {
	sync.Mutex
	deviceConfigurations map[string]map[string]any
	assignments          map[string][]any
}

// configFixture ties an @odata.type to the id it is created under and the fixture used to answer
// GET. Keyed by @odata.type so the POST and GET responders agree on which body belongs to which
// profile type.
type configFixture struct {
	odataTypeID string
	id          string
	createFile  string
	getFile     string
}

var configFixtures = []configFixture{
	{
		odataTypeID: "#microsoft.graph.iosGeneralDeviceConfiguration",
		id:          "0e2f87bd-de2c-47c2-a69b-c8cbe9823042",
		createFile:  "../tests/responses/validate_create/post_ios_general_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_general_device_configuration.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosDeviceFeaturesConfiguration",
		id:          "1f3a98ce-ef3d-58d3-b7ac-d9dcfa934153",
		createFile:  "../tests/responses/validate_create/post_ios_device_features_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_device_features_configuration.json",
	},
}

func fixtureByOdataType(odataType string) (configFixture, bool) {
	for _, f := range configFixtures {
		if f.odataTypeID == odataType {
			return f, true
		}
	}
	return configFixture{}, false
}

func fixtureByID(id string) (configFixture, bool) {
	for _, f := range configFixtures {
		if f.id == id {
			return f, true
		}
	}
	return configFixture{}, false
}

func init() {
	mockState.deviceConfigurations = make(map[string]map[string]any)
	mockState.assignments = make(map[string][]any)
	httpmock.RegisterNoResponder(
		httpmock.NewStringResponder(
			404,
			`{"error":{"code":"ResourceNotFound","message":"Resource not found"}}`,
		),
	)
	mocks.GlobalRegistry.Register(
		"ios_device_configuration_templates_json",
		&IosDeviceConfigurationTemplatesJsonMock{},
	)
}

type IosDeviceConfigurationTemplatesJsonMock struct{}

var _ mocks.MockRegistrar = (*IosDeviceConfigurationTemplatesJsonMock)(nil)

// buildCreateResponse answers POST /deviceManagement/deviceConfigurations.
//
// Returning (body, status) rather than (*http.Response, error) keeps the httpmock call in the
// registering closure, avoiding six wrapcheck exemptions for an error that carries no useful context
// in a test double.
func buildCreateResponse(req *http.Request) (map[string]any, int) {
	badRequest := func(message string) (map[string]any, int) {
		return map[string]any{
			"error": map[string]any{"code": "BadRequest", "message": message},
		}, 400
	}
	internalError := func(message string) (map[string]any, int) {
		return map[string]any{
			"error": map[string]any{"code": "InternalError", "message": message},
		}, 500
	}

	// rawJSONRequestBody emits keys in Go map order, which is randomised — so the body must be
	// decoded rather than string-matched.
	var requestBody map[string]any
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		return badRequest("Invalid request body")
	}

	odataType, ok := requestBody["@odata.type"].(string)
	if !ok {
		return badRequest("Missing @odata.type")
	}

	fixture, ok := fixtureByOdataType(odataType)
	if !ok {
		return badRequest("Unsupported configuration type")
	}

	jsonStr, err := helpers.ParseJSONFile(fixture.createFile)
	if err != nil {
		return internalError(
			fmt.Sprintf("Failed to parse JSON file '%s': %s", fixture.createFile, err.Error()),
		)
	}

	var responseObj map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &responseObj); err != nil {
		return internalError(fmt.Sprintf("Failed to parse response JSON: %s", err.Error()))
	}

	// Echo the submitted properties back, the way Graph does. The fixture supplies only the
	// server-owned envelope fields; echoing keeps the mock honest rather than letting a stale
	// fixture masquerade as a provider bug.
	for key, value := range requestBody {
		responseObj[key] = value
	}
	responseObj["id"] = fixture.id

	// Store the submitted settings so GET can replay them over the full-surface fixture.
	mockState.Lock()
	stored := make(map[string]any, len(requestBody))
	for key, value := range requestBody {
		stored[key] = value
	}
	mockState.deviceConfigurations[fixture.id] = stored
	mockState.Unlock()

	return responseObj, 201
}

// buildGetResponse answers GET /deviceManagement/deviceConfigurations('{id}').
//
// The fixture provides the full property surface Graph really returns (~190 keys for a general
// configuration); the values actually submitted are layered on top. That combination is what makes
// the projection path meaningful under test: the response is far wider than the configuration.
func buildGetResponse(configId string) (map[string]any, int) {
	mockState.Lock()
	stored, exists := mockState.deviceConfigurations[configId]
	mockState.Unlock()

	if !exists {
		return map[string]any{
			"error": map[string]any{"code": "ResourceNotFound", "message": "Resource not found"},
		}, 404
	}

	getFile := configFixtures[0].getFile
	if fixture, ok := fixtureByID(configId); ok {
		getFile = fixture.getFile
	}

	jsonStr, err := helpers.ParseJSONFile(getFile)
	if err != nil {
		return map[string]any{
			"error": map[string]any{
				"code":    "InternalError",
				"message": fmt.Sprintf("Failed to parse JSON file '%s': %s", getFile, err.Error()),
			},
		}, 500
	}

	var responseObj map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &responseObj); err != nil {
		return map[string]any{
			"error": map[string]any{
				"code":    "InternalError",
				"message": fmt.Sprintf("Failed to parse response JSON: %s", err.Error()),
			},
		}, 500
	}

	for key, value := range stored {
		responseObj[key] = value
	}
	responseObj["id"] = configId

	// Reproduce a real Graph behaviour that would otherwise be invisible under test: the discriminator
	// is echoed back only where the enclosing collection's element type is abstract and it is needed to
	// resolve which concrete type an element is. Where the element type is already concrete it is
	// redundant, and Graph drops it from the response even though it accepted — and required — it on
	// write.
	//
	// Observed for an iosDeviceFeaturesConfiguration home screen tree:
	//   icons          []iosHomeScreenItem     abstract -> echoed
	//   homeScreenPages, pages, apps           concrete -> dropped
	//
	// Echoing the submitted tree verbatim would hide this and let a projection regression through.
	dropRedundantDiscriminators(responseObj)

	return responseObj, 200
}

// loadAssignmentsFixture decodes the assignments collection fixture. It is returned as a decoded map
// so callers can hand it to httpmock.NewJsonResponse, which sets Content-Type: application/json —
// required because the typed SDK parses this response and silently yields nil without it.
func loadAssignmentsFixture() (map[string]any, int) {
	jsonStr, err := helpers.ParseJSONFile(
		"../tests/responses/validate_assign/get_ios_device_configuration_assignments.json",
	)
	if err != nil {
		return map[string]any{"value": []any{}}, 200
	}

	var body map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &body); err != nil {
		return map[string]any{"value": []any{}}, 200
	}

	return body, 200
}

// loadErrorFixture decodes an error fixture for httpmock.NewJsonResponse. Kiota parses error bodies
// through the registered ErrorMappings, so these responses need Content-Type: application/json just
// as success responses do — otherwise the diagnostic degrades to "no response body".
func loadErrorFixture(path string, status int) (map[string]any, int) {
	jsonStr, err := helpers.ParseJSONFile(path)
	if err != nil {
		return map[string]any{
			"error": map[string]any{
				"code":    "InternalError",
				"message": "Failed to parse error response",
			},
		}, 500
	}

	var body map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &body); err != nil {
		return map[string]any{
			"error": map[string]any{
				"code":    "InternalError",
				"message": "Failed to parse error response",
			},
		}, 500
	}

	return body, status
}

// concreteCollectionKeys are collections whose declared element type is concrete, so Graph treats a
// supplied @odata.type as redundant and omits it from responses. Collections with abstract element
// types (notably homeScreenPages[].icons, which is []iosHomeScreenItem) keep theirs.
var concreteCollectionKeys = map[string]bool{
	"homeScreenPages":     true,
	"pages":               true,
	"apps":                true,
	"homeScreenDockIcons": false, // []iosHomeScreenItem — abstract, discriminator retained
}

// dropRedundantDiscriminators walks a response body and removes @odata.type from the elements of
// concrete-typed collections, mirroring what Graph really returns.
func dropRedundantDiscriminators(node any) {
	switch v := node.(type) {
	case map[string]any:
		for key, child := range v {
			if concreteCollectionKeys[key] {
				if elems, ok := child.([]any); ok {
					for _, elem := range elems {
						if m, ok := elem.(map[string]any); ok {
							delete(m, "@odata.type")
						}
					}
				}
			}
			dropRedundantDiscriminators(child)
		}
	case []any:
		for _, child := range v {
			dropRedundantDiscriminators(child)
		}
	}
}

func (m *IosDeviceConfigurationTemplatesJsonMock) RegisterMocks() {
	mockState.Lock()
	mockState.deviceConfigurations = make(map[string]map[string]any)
	mockState.assignments = make(map[string][]any)
	mockState.Unlock()

	m.registerDependencyMocks()

	httpmock.RegisterResponder(
		"POST",
		"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations",
		func(req *http.Request) (*http.Response, error) {
			body, status := buildCreateResponse(req)
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(status, body)
		},
	)

	// The provider addresses items as deviceConfigurations('{id}'), so the path segment includes the
	// surrounding quotes and parentheses.
	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations\('([^']+)'\)(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			body, status := buildGetResponse(extractQuotedID(req.URL.Path))
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(status, body)
		},
	)

	httpmock.RegisterResponder(
		"PATCH",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations\('([^']+)'\)(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			configId := extractQuotedID(req.URL.Path)

			var requestBody map[string]any
			if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
				return httpmock.NewStringResponse(
					400,
					`{"error":{"code":"BadRequest","message":"Invalid request body"}}`,
				), nil
			}

			// PATCH semantics: merge over what is stored rather than replacing it, so properties
			// absent from settings_json stay as they were — which is what the resource relies on.
			mockState.Lock()
			if existing, exists := mockState.deviceConfigurations[configId]; exists {
				for key, value := range requestBody {
					existing[key] = value
				}
			}
			mockState.Unlock()

			return httpmock.NewStringResponse(204, ""), nil
		},
	)

	// Assignments are read through the typed SDK, which addresses the collection as
	// deviceConfigurations/{id}/assignments — no quoting on this path.
	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)/assignments(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			configId := segments[len(segments)-2]

			mockState.Lock()
			assigned := len(mockState.assignments[configId]) > 0
			mockState.Unlock()

			// Nothing assigned yet: return an empty collection, as Graph would.
			if !assigned {
				//nolint:wrapcheck // test double: httpmock's error needs no additional context
				return httpmock.NewJsonResponse(200, map[string]any{"value": []any{}})
			}

			// Once assigned, serve the fixture rather than echoing back what was POSTed to /assign.
			// Those stored bodies are assignment *inputs*: they carry only a target, lacking the
			// server-assigned id/source/sourceId/intent, so replaying them does not deserialize into
			// the collection the typed SDK expects and the mapper sees nothing.
			body, status := loadAssignmentsFixture()
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(status, body)
		},
	)

	httpmock.RegisterResponder(
		"POST",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)/assign(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			configId := segments[len(segments)-2]

			var requestBody map[string]any
			if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
				return httpmock.NewStringResponse(
					400,
					`{"error":{"code":"BadRequest","message":"Invalid request body"}}`,
				), nil
			}

			mockState.Lock()
			if assignments, ok := requestBody["assignments"].([]any); ok {
				mockState.assignments[configId] = assignments
			}
			mockState.Unlock()

			body, status := loadAssignmentsFixture()
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(status, body)
		},
	)

	httpmock.RegisterResponder(
		"DELETE",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			configId := strings.Trim(segments[len(segments)-1], "()'")

			mockState.Lock()
			delete(mockState.deviceConfigurations, configId)
			delete(mockState.assignments, configId)
			mockState.Unlock()

			return httpmock.NewStringResponse(204, ""), nil
		},
	)
}

// extractQuotedID pulls the id out of a deviceConfigurations('{id}') path segment.
func extractQuotedID(path string) string {
	segments := strings.Split(path, "/")
	last := segments[len(segments)-1]
	last = strings.TrimPrefix(last, "deviceConfigurations")
	return strings.Trim(last, "()'")
}

func (m *IosDeviceConfigurationTemplatesJsonMock) RegisterErrorMocks() {
	mockState.Lock()
	mockState.deviceConfigurations = make(map[string]map[string]any)
	mockState.assignments = make(map[string][]any)
	mockState.Unlock()

	m.registerDependencyMocks()

	httpmock.RegisterResponder(
		"POST",
		"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations",
		func(req *http.Request) (*http.Response, error) {
			body, status := loadErrorFixture(
				"../tests/responses/validate_create/post_ios_device_configuration_error.json",
				400,
			)
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(status, body)
		},
	)

	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations\('([^']+)'\)(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			body, status := loadErrorFixture(
				"../tests/responses/validate_delete/get_ios_device_configuration_not_found.json",
				404,
			)
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(status, body)
		},
	)

	httpmock.RegisterResponder(
		"PATCH",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations\('([^']+)'\)(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(400, map[string]any{
				"error": map[string]any{
					"code":    "BadRequest",
					"message": "Error updating iOS/iPadOS device configuration template",
				},
			})
		},
	)
}

func (m *IosDeviceConfigurationTemplatesJsonMock) CleanupMockState() {
	mockState.Lock()
	mockState.deviceConfigurations = make(map[string]map[string]any)
	mockState.assignments = make(map[string][]any)
	mockState.Unlock()
}

// registerDependencyMocks registers mocks for groups, role scope tags, assignment filters and auth.
func (m *IosDeviceConfigurationTemplatesJsonMock) registerDependencyMocks() {
	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/roleScopeTags/([^/]+)$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			tagId := segments[len(segments)-1]

			response := map[string]any{
				"@odata.type": "#microsoft.graph.roleScopeTag",
				"id":          tagId,
				"displayName": fmt.Sprintf("Role Scope Tag %s", tagId),
				"description": "Test role scope tag",
			}

			jsonBytes, _ := json.Marshal(response)
			return httpmock.NewStringResponse(200, string(jsonBytes)), nil
		},
	)

	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/groups/([^/]+)$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			groupId := segments[len(segments)-1]

			response := map[string]any{
				"@odata.type":     "#microsoft.graph.group",
				"id":              groupId,
				"displayName":     fmt.Sprintf("Test Group %s", groupId),
				"description":     "Test group for device configuration",
				"groupTypes":      []string{},
				"securityEnabled": true,
			}

			jsonBytes, _ := json.Marshal(response)
			return httpmock.NewStringResponse(200, string(jsonBytes)), nil
		},
	)

	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/assignmentFilters/([^/]+)$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			filterId := segments[len(segments)-1]

			response := map[string]any{
				"@odata.type":   "#microsoft.graph.deviceAndAppManagementAssignmentFilter",
				"id":            filterId,
				"displayName":   fmt.Sprintf("Test Assignment Filter %s", filterId),
				"description":   "Test assignment filter",
				"platform":      "iOS",
				"rule":          "(device.deviceOwnership -eq \"Corporate\")",
				"roleScopeTags": []string{"0"},
			}

			jsonBytes, _ := json.Marshal(response)
			return httpmock.NewStringResponse(200, string(jsonBytes)), nil
		},
	)

	httpmock.RegisterResponder(
		"POST",
		"https://login.microsoftonline.com/common/oauth2/v2.0/token",
		func(req *http.Request) (*http.Response, error) {
			response := map[string]any{
				"access_token": "mock_access_token_" + uuid.New().String(),
				"token_type":   "Bearer",
				"expires_in":   3600,
			}
			jsonBytes, _ := json.Marshal(response)
			return httpmock.NewStringResponse(200, string(jsonBytes)), nil
		},
	)
}
