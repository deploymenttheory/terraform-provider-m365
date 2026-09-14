package graphBetaNetworkConditionalAccessSettings_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/check"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/helpers"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	target "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/identity_and_access/graph_beta/network_conditional_access_settings"
)

// This test adopts without changing tenant security settings. Toggle validation requires an approved manual test window.
func TestAccResourceNetworkConditionalAccessSettings_01_Adopt(t *testing.T) {
	hcl, err := helpers.ParseHCLFile("tests/terraform/acceptance/omitted.tf")
	require.NoError(t, err)
	config := acceptance.ConfiguredM365ProviderBlock(hcl)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		// The singleton must remain in Graph after destroy; it has no resource deletion lifecycle.
		Steps: []resource.TestStep{
			{Config: config, Check: resource.ComposeTestCheckFunc(
				check.That(resourceType+".test").Key("id").HasValue("conditionalAccess"),
				check.That(resourceType+".test").
					ExistsInGraph(target.NetworkConditionalAccessSettingsTestResource{}),
			)},
			{
				ResourceName:            resourceType + ".test",
				ImportState:             true,
				ImportStateId:           "conditionalAccess",
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
			{Config: config, PlanOnly: true},
		},
	})
}
