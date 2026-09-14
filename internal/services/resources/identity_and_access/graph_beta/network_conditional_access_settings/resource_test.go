package graphBetaNetworkConditionalAccessSettings_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/check"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	target "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/identity_and_access/graph_beta/network_conditional_access_settings"
	settingsmocks "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/identity_and_access/graph_beta/network_conditional_access_settings/mocks"
)

const resourceType = target.ResourceName

func setupMockEnvironment(t *testing.T) *settingsmocks.ConditionalAccessSettingsMock {
	t.Helper()
	mocks.SetupUnitTestEnvironment(t)
	httpmock.Activate()
	mocks.NewMocks().AuthMocks.RegisterMocks()
	m := &settingsmocks.ConditionalAccessSettingsMock{}
	m.RegisterMocks()
	t.Cleanup(func() { httpmock.DeactivateAndReset(); m.CleanupMockState() })
	return m
}

func TestUnitResourceNetworkConditionalAccessSettings_01_Lifecycle(t *testing.T) {
	m := setupMockEnvironment(t)
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: mocks.LoadUnitTerraformConfig("enabled.tf"),
				Check:  check.That(resourceType + ".test").Key("id").HasValue("conditionalAccess"),
			},
			{
				Config: mocks.LoadUnitTerraformConfig("disabled.tf"),
				Check: check.That(resourceType + ".test").
					Key("signaling_status").
					HasValue("disabled"),
			},
			{Config: mocks.LoadUnitTerraformConfig("omitted.tf"), PlanOnly: true},
			{
				ResourceName:      resourceType + ".test",
				ImportState:       true,
				ImportStateId:     "conditionalAccess",
				ImportStateVerify: true,
			},
			{Config: mocks.LoadUnitTerraformConfig("disabled.tf"), PlanOnly: true},
		},
	})
	m.Lock()
	defer m.Unlock()
	require.Equal(t, "disabled", m.Status, "destroy must preserve the remote settings")
	require.Len(t, m.Requests, 2)
	require.Equal(t, map[string]any{"signalingStatus": "enabled"}, m.Requests[0])
	require.Equal(t, map[string]any{"signalingStatus": "disabled"}, m.Requests[1])
	require.NotContains(t, m.Methods, "POST")
	require.NotContains(t, m.Methods, "DELETE")
}

func TestUnitResourceNetworkConditionalAccessSettings_02_NoDefault(t *testing.T) {
	for _, initial := range []string{"enabled", "disabled"} {
		t.Run(initial, func(t *testing.T) {
			m := setupMockEnvironment(t)
			m.Status = initial
			resource.UnitTest(
				t,
				resource.TestCase{
					ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
					Steps: []resource.TestStep{
						{
							Config: mocks.LoadUnitTerraformConfig("omitted.tf"),
							Check: check.That(resourceType + ".test").
								Key("signaling_status").
								HasValue(initial),
						},
						{
							ResourceName:      resourceType + ".test",
							ImportState:       true,
							ImportStateId:     "conditionalAccess",
							ImportStateVerify: true,
						},
						{Config: mocks.LoadUnitTerraformConfig("omitted.tf"), PlanOnly: true},
					},
				},
			)
			require.Empty(t, m.Requests)
			require.Equal(t, initial, m.Status)
		})
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_03_InvalidValue(t *testing.T) {
	m := setupMockEnvironment(t)
	resource.UnitTest(
		t,
		resource.TestCase{
			ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config:      mocks.LoadUnitTerraformConfig("invalid.tf"),
					ExpectError: regexp.MustCompile("value must be one of"),
				},
			},
		},
	)
	require.Empty(t, m.Requests)
}

func TestUnitResourceNetworkConditionalAccessSettings_04_Drift(t *testing.T) {
	m := setupMockEnvironment(t)
	resource.UnitTest(
		t,
		resource.TestCase{
			ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{Config: mocks.LoadUnitTerraformConfig("enabled.tf")},
				{
					PreConfig:          func() { m.Lock(); defer m.Unlock(); m.Status = "disabled" },
					Config:             mocks.LoadUnitTerraformConfig("enabled.tf"),
					PlanOnly:           true,
					ExpectNonEmptyPlan: true,
				},
				{
					Config: mocks.LoadUnitTerraformConfig("enabled.tf"),
					Check: check.That(resourceType + ".test").
						Key("signaling_status").
						HasValue("enabled"),
				},
				{
					PreConfig: func() { m.Lock(); defer m.Unlock(); m.Status = "disabled" },
					Config:    mocks.LoadUnitTerraformConfig("omitted.tf"),
					Check: check.That(resourceType + ".test").
						Key("signaling_status").
						HasValue("disabled"),
				},
			},
		},
	)
	require.Len(t, m.Requests, 2, "omission must not re-enable signaling after drift")
}
