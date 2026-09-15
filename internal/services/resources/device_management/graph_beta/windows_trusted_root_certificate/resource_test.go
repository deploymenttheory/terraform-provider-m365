package graphBetaWindowsTrustedRootCertificate_test

import (
	_ "embed"
	"encoding/base64"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/jarcoal/httpmock"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/check"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/helpers"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	certresource "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/device_management/graph_beta/windows_trusted_root_certificate"
	certmocks "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/device_management/graph_beta/windows_trusted_root_certificate/mocks"
)

//go:embed tests/test-root.cer
var certificate []byte

var (
	resourceType    = certresource.ResourceName
	resourceAddress = resourceType + ".test"
)

func loadConfig(t *testing.T, fixture string) string {
	t.Helper()
	config, err := helpers.ParseHCLFile("tests/terraform/" + fixture)
	if err != nil {
		t.Fatal(err)
	}
	return strings.ReplaceAll(
		config,
		"CERTIFICATE_BASE64",
		base64.StdEncoding.EncodeToString(certificate),
	)
}

func setupMockEnvironment(t *testing.T) *certmocks.WindowsTrustedRootCertificateMock {
	t.Helper()
	mocks.SetupUnitTestEnvironment(t)
	httpmock.Activate()
	mocks.NewMocks().AuthMocks.RegisterMocks()
	mock := &certmocks.WindowsTrustedRootCertificateMock{}
	mock.RegisterMocks()
	t.Cleanup(func() { httpmock.DeactivateAndReset(); mock.CleanupMockState() })
	return mock
}

func importStep() resource.TestStep {
	return resource.TestStep{
		ResourceName:            resourceAddress,
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateVerifyIgnore: []string{"timeouts"},
	}
}

func TestUnitResourceWindowsTrustedRootCertificate_01_Lifecycle(t *testing.T) {
	setupMockEnvironment(t)
	minimal := loadConfig(t, "unit/resource.tf")
	assigned := strings.Replace(
		minimal,
		"unit-test-windows-trusted-certificate",
		"updated-certificate",
		1,
	)
	assigned = strings.TrimSuffix(strings.TrimSpace(assigned), "}") + `
  description = "updated description"
  assignments = [
    { type = "allDevicesAssignmentTarget" },
    { type = "allLicensedUsersAssignmentTarget" },
    { type = "groupAssignmentTarget", group_id = "11111111-1111-4111-8111-111111111111", filter_type = "include", filter_id = "33333333-3333-4333-8333-333333333333" },
    { type = "exclusionGroupAssignmentTarget", group_id = "22222222-2222-4222-8222-222222222222" }
  ]
}`
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: minimal, Check: resource.ComposeTestCheckFunc(
				check.That(resourceAddress).
					Key("trusted_root_certificate").
					HasValue(base64.StdEncoding.EncodeToString(certificate)),
				check.That(resourceAddress).
					Key("destination_store").
					HasValue("computerCertStoreRoot"),
				check.That(resourceAddress).Key("role_scope_tag_ids.#").HasValue("1"),
			)},
			importStep(),
			{Config: assigned, Check: resource.ComposeTestCheckFunc(
				check.That(resourceAddress).Key("assignments.#").HasValue("4"),
				check.That(resourceAddress).Key("display_name").HasValue("updated-certificate"),
			)},
			importStep(),
			{Config: assigned, PlanOnly: true},
			{
				Config: strings.Replace(
					minimal,
					"computerCertStoreRoot",
					"computerCertStoreIntermediate",
					1,
				),
			},
			importStep(),
			{
				Config: strings.Replace(
					minimal,
					"computerCertStoreRoot",
					"userCertStoreIntermediate",
					1,
				),
			},
			importStep(),
		},
	})
}

func TestUnitResourceWindowsTrustedRootCertificate_02_Validation(t *testing.T) {
	for _, tc := range []struct{ name, from, to, message string }{
		{"base64", "CERTIFICATE_BASE64", "!invalid!", "decode certificate base64"},
		{"certificate", "CERTIFICATE_BASE64", "bm90IGEgY2VydGlmaWNhdGU=", "parse DER X.509 certificate"},
		{"store", "computerCertStoreRoot", "invalidStore", "destination_store"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setupMockEnvironment(t)
			config, err := helpers.ParseHCLFile("tests/terraform/unit/resource.tf")
			if err != nil {
				t.Fatal(err)
			}
			config = strings.ReplaceAll(config, tc.from, tc.to)
			config = strings.ReplaceAll(
				config,
				"CERTIFICATE_BASE64",
				base64.StdEncoding.EncodeToString(certificate),
			)
			resource.UnitTest(t, resource.TestCase{
				ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{Config: config, ExpectError: regexp.MustCompile(tc.message)},
				},
			})
		})
	}
}

func TestUnitResourceWindowsTrustedRootCertificate_03_ImportExisting(t *testing.T) {
	mock := setupMockEnvironment(t)
	id := "44444444-4444-4444-8444-444444444444"
	mock.SeedProfile(
		id,
		base64.StdEncoding.EncodeToString(certificate),
		"#microsoft.graph.windows81TrustedRootCertificate",
	)
	config := strings.Replace(
		loadConfig(t, "unit/resource.tf"),
		"unit-test-windows-trusted-certificate",
		"externally-created",
		1,
	)
	config += `
import {
  to = microsoft365_graph_beta_device_management_windows_trusted_root_certificate.test
  id = "44444444-4444-4444-8444-444444444444"
}`
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, Check: check.That(resourceAddress).Key("id").HasValue(id)},
			{Config: config, PlanOnly: true},
		},
	})
}

func TestUnitResourceWindowsTrustedRootCertificate_04_WrongPolicyImport(t *testing.T) {
	mock := setupMockEnvironment(t)
	id := "55555555-5555-4555-8555-555555555555"
	mock.SeedProfile(
		id,
		base64.StdEncoding.EncodeToString(certificate),
		"#microsoft.graph.windows10CustomConfiguration",
	)
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        loadConfig(t, "unit/resource.tf"),
				ResourceName:  resourceAddress,
				ImportState:   true,
				ImportStateId: id,
				ExpectError:   regexp.MustCompile("Incorrect Intune policy type"),
			},
		},
	})
}

func TestUnitResourceWindowsTrustedRootCertificate_05_PermissionDenied(t *testing.T) {
	setupMockEnvironment(t).RegisterErrorMocks()
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      loadConfig(t, "unit/resource.tf"),
				ExpectError: regexp.MustCompile("403|Forbidden|denied"),
			},
		},
	})
}

func TestUnitResourceWindowsTrustedRootCertificate_06_DeletedRemotely(t *testing.T) {
	mock := setupMockEnvironment(t)
	config := loadConfig(t, "unit/resource.tf")
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config},
			{
				PreConfig:          mock.CleanupMockState,
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestUnitResourceWindowsTrustedRootCertificate_07_InvalidAssignments(t *testing.T) {
	for _, tc := range []struct{ name, assignment, message string }{
		{"missing_group", `{ type = "groupAssignmentTarget" }`, "group targets require group_id"},
		{"unrelated_group", `{ type = "allDevicesAssignmentTarget", group_id = "11111111-1111-4111-8111-111111111111" }`, "group_id only applies to group targets"},
		{"filter_without_type", `{ type = "allDevicesAssignmentTarget", filter_id = "11111111-1111-4111-8111-111111111111" }`, "filter_id requires filter_type"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			setupMockEnvironment(t)
			config := strings.TrimSuffix(
				strings.TrimSpace(loadConfig(t, "unit/resource.tf")),
				"}",
			) + "\nassignments = [" + tc.assignment + "]\n}"
			resource.UnitTest(t, resource.TestCase{
				ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{Config: config, ExpectError: regexp.MustCompile(tc.message)},
				},
			})
			if count := httpmock.GetCallCountInfo()["POST https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations"]; count != 0 {
				t.Fatalf("invalid assignments created a remote policy: %d requests", count)
			}
		})
	}
}

func TestUnitResourceWindowsTrustedRootCertificate_08_AssignmentFailureRetainsID(t *testing.T) {
	setupMockEnvironment(t)
	httpmock.RegisterResponder(
		"POST",
		`=~^https://graph\.microsoft\.com/beta/deviceManagement/deviceConfigurations/[^/]+/assign$`,
		httpmock.NewStringResponder(
			400,
			`{"error":{"code":"BadRequest","message":"Assignment rejected"}}`,
		),
	)
	config := strings.TrimSuffix(strings.TrimSpace(loadConfig(t, "unit/resource.tf")), "}") + `
assignments = [{type = "groupAssignmentTarget", group_id = "11111111-1111-4111-8111-111111111111"}]
}`
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: config, ExpectError: regexp.MustCompile("Bad Request|400")},
		},
	})
	deleted := false
	for call, count := range httpmock.GetCallCountInfo() {
		if strings.HasPrefix(call, "DELETE ") && count > 0 {
			deleted = true
		}
	}
	if !deleted {
		t.Fatal("failed assignment lost the policy ID; test cleanup could not delete it")
	}
}
