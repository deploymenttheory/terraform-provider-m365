package mocks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
)

const endpoint = "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations"

type WindowsTrustedRootCertificateMock struct {
	sync.Mutex
	profiles map[string]map[string]any
}

var _ mocks.MockRegistrar = (*WindowsTrustedRootCertificateMock)(nil)

func init() {
	mocks.GlobalRegistry.Register(
		"windows_trusted_root_certificate",
		&WindowsTrustedRootCertificateMock{},
	)
}

func (m *WindowsTrustedRootCertificateMock) SeedProfile(id, certificate, policyType string) {
	m.Lock()
	defer m.Unlock()
	m.profiles[id] = map[string]any{
		"id":                     id,
		"@odata.type":            policyType,
		"displayName":            "externally-created",
		"description":            "",
		"certFileName":           "root.cer",
		"trustedRootCertificate": certificate,
		"destinationStore":       "computerCertStoreRoot",
		"roleScopeTagIds":        []string{"0"},
		"assignments":            []any{},
	}
}

func (m *WindowsTrustedRootCertificateMock) RegisterMocks() {
	m.CleanupMockState()
	httpmock.RegisterResponder("POST", endpoint, func(req *http.Request) (*http.Response, error) {
		var body map[string]any
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			return nil, fmt.Errorf("decode mock request: %w", err)
		}
		if body["@odata.type"] != "#microsoft.graph.windows81TrustedRootCertificate" {
			return httpmock.NewStringResponse(
				400,
				`{"error":{"code":"BadRequest","message":"Incorrect policy type"}}`,
			), nil
		}
		id := uuid.NewString()
		m.Lock()
		defer m.Unlock()
		body["id"] = id
		if _, ok := body["description"]; !ok {
			body["description"] = ""
		}
		if _, ok := body["roleScopeTagIds"]; !ok {
			body["roleScopeTagIds"] = []string{"0"}
		}
		body["assignments"] = []any{}
		m.profiles[id] = body
		return httpmock.NewJsonResponse(201, body)
	})
	pattern := `=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/[^/?]+(\?.*)?$`
	for _, method := range []string{"GET", "PATCH", "DELETE"} {
		httpmock.RegisterResponder(method, pattern, m.profile)
	}
	httpmock.RegisterResponder(
		"POST",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/[^/]+/assign$`,
		func(req *http.Request) (*http.Response, error) {
			var body struct {
				Assignments []any `json:"assignments"`
			}
			if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
				return nil, fmt.Errorf("decode mock request: %w", err)
			}
			m.Lock()
			defer m.Unlock()
			id := strings.Split(req.URL.Path, "/")[4]
			profile, ok := m.profiles[id]
			if !ok {
				return notFound(), nil
			}
			profile["assignments"] = body.Assignments
			return httpmock.NewJsonResponse(200, map[string]any{"value": body.Assignments})
		},
	)
}

func (m *WindowsTrustedRootCertificateMock) profile(req *http.Request) (*http.Response, error) {
	m.Lock()
	defer m.Unlock()
	id := strings.Split(req.URL.Path, "/")[4]
	profile, ok := m.profiles[id]
	if !ok {
		return notFound(), nil
	}
	switch req.Method {
	case "DELETE":
		delete(m.profiles, id)
		return httpmock.NewStringResponse(204, ""), nil
	case "PATCH":
		var body map[string]any
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			return nil, fmt.Errorf("decode mock request: %w", err)
		}
		for key, value := range body {
			profile[key] = value
		}
		return httpmock.NewStringResponse(204, ""), nil
	default:
		response, err := httpmock.NewJsonResponse(200, profile)
		if err != nil {
			return nil, fmt.Errorf("encode mock response: %w", err)
		}
		return response, nil
	}
}

func (m *WindowsTrustedRootCertificateMock) RegisterErrorMocks() {
	m.RegisterMocks()
	httpmock.RegisterResponder(
		"POST",
		endpoint,
		httpmock.NewStringResponder(
			403,
			`{"error":{"code":"Forbidden","message":"Policy write denied"}}`,
		),
	)
}

func (m *WindowsTrustedRootCertificateMock) CleanupMockState() {
	m.Lock()
	defer m.Unlock()
	m.profiles = make(map[string]map[string]any)
}

func notFound() *http.Response {
	return httpmock.NewStringResponse(
		404,
		`{"error":{"code":"ResourceNotFound","message":"Policy not found"}}`,
	)
}
