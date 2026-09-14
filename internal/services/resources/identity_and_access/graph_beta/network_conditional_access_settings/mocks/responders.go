package mocks

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/jarcoal/httpmock"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks/factories"
)

const BaseURL = "https://graph.microsoft.com/beta/networkAccess/settings/conditionalAccess"

// ConditionalAccessSettingsMock records partial writes against an existing singleton.
type ConditionalAccessSettingsMock struct {
	sync.Mutex
	Status   string
	Requests []map[string]any
	Methods  []string
}

var _ mocks.MockRegistrar = (*ConditionalAccessSettingsMock)(nil)

func init() {
	mocks.GlobalRegistry.Register(
		"network_conditional_access_settings",
		&ConditionalAccessSettingsMock{},
	)
}

func (m *ConditionalAccessSettingsMock) CleanupMockState() {
	m.Lock()
	defer m.Unlock()
	m.Status = "enabled"
	m.Requests = nil
	m.Methods = nil
}

func (m *ConditionalAccessSettingsMock) RegisterMocks() {
	m.CleanupMockState()
	for _, method := range []string{"GET", "PATCH", "POST", "DELETE"} {
		httpmock.RegisterResponder(method, BaseURL, m.respond)
	}
}

func (m *ConditionalAccessSettingsMock) RegisterErrorMocks() {
	httpmock.RegisterResponder(
		"PATCH",
		BaseURL,
		factories.ErrorResponse(403, "Forbidden", "Synthetic permission error"),
	)
}

func (m *ConditionalAccessSettingsMock) respond(req *http.Request) (*http.Response, error) {
	m.Lock()
	defer m.Unlock()
	m.Methods = append(m.Methods, req.Method)
	switch req.Method {
	case "GET":
		// Extra remote properties must survive all writes and never enter Terraform's request body.
		return httpmock.NewJsonResponse(
			200,
			map[string]any{"signalingStatus": m.Status, "dataPlaneSignalingOptions": "entraId"},
		)
	case "PATCH":
		var body map[string]any
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			return factories.ErrorResponse(400, "BadRequest", "Invalid JSON")(req)
		}
		m.Requests = append(m.Requests, body)
		if len(body) != 1 ||
			(body["signalingStatus"] != "enabled" && body["signalingStatus"] != "disabled") {
			return factories.ErrorResponse(400, "BadRequest", "Expected only signalingStatus")(req)
		}
		m.Status = body["signalingStatus"].(string)
		return httpmock.NewStringResponse(204, ""), nil
	default:
		return factories.ErrorResponse(
			405,
			"MethodNotAllowed",
			"No create or delete operation",
		)(
			req,
		)
	}
}
