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

// configFixture ties an @odata.type to the id it is created under and the fixtures used to
// answer create and get. Keyed by @odata.type so the POST responder and the GET responder agree
// on which body belongs to which configuration type.
type configFixture struct {
	id          string
	createFile  string
	getFile     string
	odataTypeID string
}

var configFixtures = []configFixture{
	{
		odataTypeID: "#microsoft.graph.iosCustomConfiguration",
		id:          "12345678-1234-1234-1234-123456789012",
		createFile:  "../tests/responses/validate_create/post_ios_custom_configuration_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_custom_configuration.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosTrustedRootCertificate",
		id:          "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
		createFile:  "../tests/responses/validate_create/post_ios_trusted_certificate_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_trusted_certificate.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosWiFiConfiguration",
		id:          "bbbbbbbb-cccc-dddd-eeee-ffffffffffff",
		createFile:  "../tests/responses/validate_create/post_ios_wifi_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_wifi.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosScepCertificateProfile",
		id:          "ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb",
		createFile:  "../tests/responses/validate_create/post_ios_scep_certificate_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_scep_certificate.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosPkcsCertificateProfile",
		id:          "11111111-2222-3333-4444-555555555555",
		createFile:  "../tests/responses/validate_create/post_ios_pkcs_certificate_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_pkcs_certificate.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosEnterpriseWiFiConfiguration",
		id:          "cccccccc-dddd-eeee-ffff-000000000000",
		createFile:  "../tests/responses/validate_create/post_ios_enterprise_wifi_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_enterprise_wifi.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosEasEmailProfileConfiguration",
		id:          "dddddddd-eeee-ffff-0000-111111111111",
		createFile:  "../tests/responses/validate_create/post_ios_eas_email_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_eas_email.json",
	},
	{
		odataTypeID: "#microsoft.graph.iosVpnConfiguration",
		id:          "eeeeeeee-ffff-0000-1111-222222222222",
		createFile:  "../tests/responses/validate_create/post_ios_vpn_success.json",
		getFile:     "../tests/responses/validate_get/get_ios_vpn.json",
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
		"ios_device_configuration_templates",
		&IosDeviceConfigurationTemplatesMock{},
	)
}

type IosDeviceConfigurationTemplatesMock struct{}

var _ mocks.MockRegistrar = (*IosDeviceConfigurationTemplatesMock)(nil)

// buildCreateResponse produces the body and status for POST /deviceManagement/deviceConfigurations,
// echoing the submitted properties back the way Graph does and recording the result in mock state.
//
// Returning (body, status) rather than (*http.Response, error) keeps the httpmock call in the
// registering closure: httpmock.NewJsonResponse returns an error that would otherwise need
// wrapping at each of the six exit points, which adds nothing to a test double.
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

	// Echo every submitted property back, the way Graph does for settable properties. The fixture
	// supplies only the server-owned fields (id, timestamps, version, supportsScopeTags,
	// applicability rules). Echoing rather than serving fixture values keeps the mock honest: a
	// fixture whose payload or certificate drifts from the test configuration would otherwise
	// surface as "provider produced inconsistent result" instead of as the fixture bug it is.
	for k, v := range requestBody {
		responseObj[k] = v
	}

	// Values Graph accepts on write but never returns on read. The provider recovers each of these
	// from prior state, so the mock must not echo them back or those recovery paths go untested.
	//   - preSharedKey / passwordFormatString: write-only secrets
	//   - *@odata.bind: navigation references, not echoed as additionalData
	delete(responseObj, "preSharedKey")
	delete(responseObj, "passwordFormatString")
	for key := range responseObj {
		if strings.Contains(key, "@odata.bind") {
			delete(responseObj, key)
		}
	}

	responseObj["id"] = fixture.id

	mockState.Lock()
	mockState.deviceConfigurations[fixture.id] = responseObj
	mockState.Unlock()

	return responseObj, 201
}

func (m *IosDeviceConfigurationTemplatesMock) RegisterMocks() {
	mockState.Lock()
	mockState.deviceConfigurations = make(map[string]map[string]any)
	mockState.assignments = make(map[string][]any)
	mockState.Unlock()

	m.registerDependencyMocks()

	// GET /deviceManagement/deviceConfigurations - List device configurations
	httpmock.RegisterResponder(
		"GET",
		"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations",
		func(req *http.Request) (*http.Response, error) {
			jsonStr, err := helpers.ParseJSONFile(
				"../tests/responses/validate_get/get_ios_device_configuration_list.json",
			)
			if err != nil {
				return httpmock.NewStringResponse(
					500,
					fmt.Sprintf(
						`{"error":{"code":"InternalError","message":"Failed to parse JSON: %s"}}`,
						err.Error(),
					),
				), nil
			}
			return httpmock.NewStringResponse(200, jsonStr), nil
		},
	)

	// POST /deviceManagement/deviceConfigurations - Create device configuration
	httpmock.RegisterResponder(
		"POST",
		"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations",
		func(req *http.Request) (*http.Response, error) {
			body, status := buildCreateResponse(req)
			//nolint:wrapcheck // test double: httpmock's error needs no additional context
			return httpmock.NewJsonResponse(status, body)
		},
	)

	// GET /deviceManagement/deviceConfigurations/{id} - Get specific device configuration
	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			if len(segments) < 4 {
				return httpmock.NewJsonResponse(400, map[string]any{
					"error": map[string]any{"code": "BadRequest", "message": "Invalid URL"},
				})
			}

			configId := segments[len(segments)-1]

			mockState.Lock()
			config, exists := mockState.deviceConfigurations[configId]
			mockState.Unlock()

			if !exists {
				return httpmock.NewJsonResponse(404, map[string]any{
					"error": map[string]any{
						"code":    "ResourceNotFound",
						"message": "Resource not found",
					},
				})
			}

			// Resolve the fixture from the id being requested, not from a fixed default, so each
			// configuration type reads back its own body.
			getFile := configFixtures[0].getFile
			if fixture, ok := fixtureByID(configId); ok {
				getFile = fixture.getFile
			}

			jsonStr, err := helpers.ParseJSONFile(getFile)
			if err != nil {
				return httpmock.NewJsonResponse(500, map[string]any{
					"error": map[string]any{
						"code": "InternalError",
						"message": fmt.Sprintf(
							"Failed to parse JSON file '%s': %s",
							getFile,
							err.Error(),
						),
					},
				})
			}

			var responseObj map[string]any
			if err := json.Unmarshal([]byte(jsonStr), &responseObj); err != nil {
				return httpmock.NewJsonResponse(500, map[string]any{
					"error": map[string]any{
						"code":    "InternalError",
						"message": fmt.Sprintf("Failed to parse response JSON: %s", err.Error()),
					},
				})
			}

			// Stored state wins so an update is visible on the next read.
			for k, v := range config {
				responseObj[k] = v
			}

			if req.URL.Query().Get("$expand") == "assignments" {
				assignmentStr, assignmentErr := helpers.ParseJSONFile(
					"../tests/responses/validate_assign/post_ios_device_configuration_assign_success.json",
				)
				responseObj["assignments"] = []any{}
				if assignmentErr == nil {
					var assignmentObj map[string]any
					if json.Unmarshal([]byte(assignmentStr), &assignmentObj) == nil {
						if assignmentValue, ok := assignmentObj["value"]; ok {
							responseObj["assignments"] = assignmentValue
						}
					}
				}
			}

			return httpmock.NewJsonResponse(200, responseObj)
		},
	)

	// PATCH /deviceManagement/deviceConfigurations/{id} - Update device configuration
	httpmock.RegisterResponder(
		"PATCH",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			if len(segments) < 4 {
				return httpmock.NewStringResponse(
					400,
					`{"error":{"code":"BadRequest","message":"Invalid URL"}}`,
				), nil
			}

			configId := segments[len(segments)-1]

			var requestBody map[string]any
			if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
				return httpmock.NewStringResponse(
					400,
					`{"error":{"code":"BadRequest","message":"Invalid request body"}}`,
				), nil
			}

			mockState.Lock()
			if existing, exists := mockState.deviceConfigurations[configId]; exists {
				for k, v := range requestBody {
					existing[k] = v
				}
			}
			mockState.Unlock()

			jsonStr, err := helpers.ParseJSONFile(
				"../tests/responses/validate_update/patch_ios_device_configuration_success.json",
			)
			if err != nil {
				return httpmock.NewStringResponse(
					500,
					fmt.Sprintf(
						`{"error":{"code":"InternalError","message":"Failed to parse JSON: %s"}}`,
						err.Error(),
					),
				), nil
			}

			return httpmock.NewStringResponse(200, jsonStr), nil
		},
	)

	// DELETE /deviceManagement/deviceConfigurations/{id} - Delete device configuration
	httpmock.RegisterResponder(
		"DELETE",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			if len(segments) < 4 {
				return httpmock.NewStringResponse(
					400,
					`{"error":{"code":"BadRequest","message":"Invalid URL"}}`,
				), nil
			}

			configId := segments[len(segments)-1]

			mockState.Lock()
			delete(mockState.deviceConfigurations, configId)
			delete(mockState.assignments, configId)
			mockState.Unlock()

			return httpmock.NewStringResponse(204, ""), nil
		},
	)

	// POST /deviceManagement/deviceConfigurations/{id}/assign - Assign device configuration
	httpmock.RegisterResponder(
		"POST",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)/assign$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			if len(segments) < 5 {
				return httpmock.NewJsonResponse(400, map[string]any{
					"error": map[string]any{"code": "BadRequest", "message": "Invalid URL"},
				})
			}

			configId := segments[len(segments)-2]

			var requestBody map[string]any
			if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
				return httpmock.NewJsonResponse(400, map[string]any{
					"error": map[string]any{
						"code":    "BadRequest",
						"message": "Invalid request body",
					},
				})
			}

			mockState.Lock()
			if assignments, ok := requestBody["assignments"].([]any); ok {
				mockState.assignments[configId] = assignments
			}
			mockState.Unlock()

			jsonStr, err := helpers.ParseJSONFile(
				"../tests/responses/validate_assign/post_ios_device_configuration_assign_success.json",
			)
			if err != nil {
				return httpmock.NewJsonResponse(500, map[string]any{
					"error": map[string]any{
						"code":    "InternalError",
						"message": fmt.Sprintf("Failed to parse JSON: %s", err.Error()),
					},
				})
			}

			var responseObj map[string]any
			if err := json.Unmarshal([]byte(jsonStr), &responseObj); err != nil {
				return httpmock.NewJsonResponse(500, map[string]any{
					"error": map[string]any{
						"code":    "InternalError",
						"message": fmt.Sprintf("Failed to parse response JSON: %s", err.Error()),
					},
				})
			}

			return httpmock.NewJsonResponse(200, responseObj)
		},
	)
}

func (m *IosDeviceConfigurationTemplatesMock) RegisterErrorMocks() {
	mockState.Lock()
	mockState.deviceConfigurations = make(map[string]map[string]any)
	mockState.assignments = make(map[string][]any)
	mockState.Unlock()

	m.registerDependencyMocks()

	// POST /deviceManagement/deviceConfigurations - Create device configuration (Error)
	httpmock.RegisterResponder(
		"POST",
		"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations",
		func(req *http.Request) (*http.Response, error) {
			jsonStr, err := helpers.ParseJSONFile(
				"../tests/responses/validate_create/post_ios_device_configuration_error.json",
			)
			if err != nil {
				return httpmock.NewJsonResponse(500, map[string]any{
					"error": map[string]any{
						"code":    "InternalError",
						"message": "Failed to parse error response",
					},
				})
			}

			var errorObj map[string]any
			if err := json.Unmarshal([]byte(jsonStr), &errorObj); err != nil {
				return httpmock.NewJsonResponse(500, map[string]any{
					"error": map[string]any{
						"code":    "InternalError",
						"message": "Failed to parse error JSON",
					},
				})
			}

			return httpmock.NewJsonResponse(400, errorObj)
		},
	)

	// GET /deviceManagement/deviceConfigurations/{id} - Not found error
	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)(\?.*)?$`,
		func(req *http.Request) (*http.Response, error) {
			jsonStr, err := helpers.ParseJSONFile(
				"../tests/responses/validate_delete/get_ios_device_configuration_not_found.json",
			)
			if err != nil {
				return httpmock.NewStringResponse(
					500,
					`{"error":{"code":"InternalError","message":"Failed to parse error response"}}`,
				), nil
			}
			return httpmock.NewStringResponse(404, jsonStr), nil
		},
	)

	httpmock.RegisterResponder(
		"PATCH",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)$`,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(
				400,
				`{"error":{"code":"BadRequest","message":"Error updating iOS/iPadOS device configuration template"}}`,
			), nil
		},
	)

	httpmock.RegisterResponder(
		"DELETE",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)$`,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(
				400,
				`{"error":{"code":"BadRequest","message":"Error deleting iOS/iPadOS device configuration template"}}`,
			), nil
		},
	)

	httpmock.RegisterResponder(
		"POST",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)/assign$`,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(
				400,
				`{"error":{"code":"BadRequest","message":"Error assigning iOS/iPadOS device configuration template"}}`,
			), nil
		},
	)
}

func (m *IosDeviceConfigurationTemplatesMock) CleanupMockState() {
	mockState.Lock()
	mockState.deviceConfigurations = make(map[string]map[string]any)
	mockState.assignments = make(map[string][]any)
	mockState.Unlock()
}

// registerDependencyMocks registers mocks for dependencies like groups, role scope tags, and assignment filters
func (m *IosDeviceConfigurationTemplatesMock) registerDependencyMocks() {
	// Mock role scope tags
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

	// Mock groups
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

	// Mock assignment filters
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

	// Mock device configuration assignments endpoint for GET requests
	httpmock.RegisterResponder(
		"GET",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/([^/]+)/assignments$`,
		func(req *http.Request) (*http.Response, error) {
			segments := strings.Split(req.URL.Path, "/")
			configId := segments[len(segments)-2]

			mockState.Lock()
			assignments := mockState.assignments[configId]
			mockState.Unlock()

			response := map[string]any{
				"@odata.context": fmt.Sprintf(
					"https://graph.microsoft.com/beta/$metadata#deviceManagement/deviceConfigurations('%s')/assignments",
					configId,
				),
				"value": assignments,
			}

			if assignments == nil {
				response["value"] = []any{}
			}

			jsonBytes, _ := json.Marshal(response)
			return httpmock.NewStringResponse(200, string(jsonBytes)), nil
		},
	)

	// Register authentication mocks
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
