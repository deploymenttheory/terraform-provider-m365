package graphBetaIosDeviceConfigurationTemplatesJson_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/jarcoal/httpmock"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/check"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/helpers"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	iosJsonMocks "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/device_management/graph_beta/ios_device_configuration_templates_json/mocks"
)

func setupMockEnvironment() (*mocks.Mocks, *iosJsonMocks.IosDeviceConfigurationTemplatesJsonMock) {
	httpmock.Activate()
	mockClient := mocks.NewMocks()
	mockClient.AuthMocks.RegisterMocks()
	configMock := &iosJsonMocks.IosDeviceConfigurationTemplatesJsonMock{}
	configMock.RegisterMocks()
	return mockClient, configMock
}

func setupErrorMockEnvironment() (*mocks.Mocks, *iosJsonMocks.IosDeviceConfigurationTemplatesJsonMock) {
	httpmock.Activate()
	mockClient := mocks.NewMocks()
	mockClient.AuthMocks.RegisterMocks()
	configMock := &iosJsonMocks.IosDeviceConfigurationTemplatesJsonMock{}
	configMock.RegisterErrorMocks()
	return mockClient, configMock
}

func loadUnitTestTerraform(filename string) string {
	config, err := helpers.ParseHCLFile("tests/terraform/unit/" + filename)
	if err != nil {
		panic("failed to load unit test config " + filename + ": " + err.Error())
	}
	return config
}

// A configuration declaring one property must not absorb the ~190 keys Graph returns. This is the
// projection path end-to-end.
func TestUnitResourceIosDeviceConfigurationTemplatesJson_01_GeneralMinimal(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_general_minimal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".general_minimal").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".general_minimal").Key("odata_type").HasValue("#microsoft.graph.iosGeneralDeviceConfiguration"),
					check.That(resourceType+".general_minimal").Key("display_name").HasValue("unit-test-iOS-disable-camera"),
					// The whole point: exactly the declared property, alphabetically normalized.
					check.That(resourceType+".general_minimal").Key("settings_json").HasValue(`{"cameraBlocked":true}`),
					check.That(resourceType+".general_minimal").Key("role_scope_tag_ids.#").HasValue("1"),
					check.That(resourceType+".general_minimal").Key("assignments.#").HasValue("2"),
				),
			},
		},
	})
}

// Nested @odata.type discriminators are load-bearing and must survive the round trip untouched.
func TestUnitResourceIosDeviceConfigurationTemplatesJson_02_DeviceFeaturesConfiguration(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_device_features_home_screen.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".device_features").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".device_features").Key("odata_type").HasValue("#microsoft.graph.iosDeviceFeaturesConfiguration"),
					// Both the app and folder discriminators, and the third at two levels of nesting.
					check.That(resourceType+".device_features").Key("settings_json").MatchesRegex(
						regexp.MustCompile(`#microsoft\.graph\.iosHomeScreenApp`),
					),
					check.That(resourceType+".device_features").Key("settings_json").MatchesRegex(
						regexp.MustCompile(`#microsoft\.graph\.iosHomeScreenFolder`),
					),
					check.That(resourceType+".device_features").Key("settings_json").MatchesRegex(
						regexp.MustCompile(`com\.apple\.mobilesafari`),
					),
					// The envelope must not leak back in via the response.
					check.That(resourceType+".device_features").Key("settings_json").MatchesRegex(
						regexp.MustCompile(`^\{"homeScreenPages"`),
					),
				),
			},
			{
				// A deep tree is the most likely thing to drift; re-planning must be a no-op.
				Config:             loadUnitTestTerraform("resource_device_features_home_screen.tf"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplatesJson_03_RejectsEnvelopeKeys(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      loadUnitTestTerraform("resource_rejects_envelope_keys.tf"),
				ExpectError: regexp.MustCompile(`(?s)Provider-Owned Keys in settings_json.*displayName`),
			},
		},
	})
}

// The minimal configuration is the most drift-prone case, since the response is ~190x wider than it.
func TestUnitResourceIosDeviceConfigurationTemplatesJson_04_NoDriftOnRefresh(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_general_minimal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".general_minimal").Key("settings_json").HasValue(`{"cameraBlocked":true}`),
				),
			},
			{
				Config:             loadUnitTestTerraform("resource_general_minimal.tf"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplatesJson_05_CreateWithError(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupErrorMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      loadUnitTestTerraform("resource_general_minimal.tf"),
				ExpectError: regexp.MustCompile("Error creating iOS/iPadOS device configuration template"),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplatesJson_06_Update(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_general_minimal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".general_minimal").Key("settings_json").HasValue(`{"cameraBlocked":true}`),
				),
			},
			{
				Config: loadUnitTestTerraform("resource_general_minimal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".general_minimal").Key("settings_json").HasValue(`{"cameraBlocked":true}`),
				),
			},
		},
	})
}

// On import there is no prior configuration to project against, so the whole stripped response is
// kept — deliberately, since discarding it would hide remote state. It must still be free of every
// envelope and server-owned key.
func TestUnitResourceIosDeviceConfigurationTemplatesJson_07_ImportState(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_general_minimal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".general_minimal").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
				),
			},
			{
				ResourceName: resourceType + ".general_minimal",
				ImportState:  true,
				// settings_json legitimately differs after import: with no prior shape to project
				// against, the full property surface is kept rather than the declared subset.
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"settings_json"},
			},
		},
	})
}
