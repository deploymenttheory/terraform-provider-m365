package graphBetaWindowsTrustedRootCertificate_test

import (
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/check"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/destroy"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/constants"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	certresource "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/device_management/graph_beta/windows_trusted_root_certificate"
)

func TestAccResourceWindowsTrustedRootCertificate_01_Lifecycle(t *testing.T) {
	testResource := certresource.WindowsTrustedRootCertificateTestResource{}
	config := loadConfig(t, "acceptance/resource.tf")
	updated := strings.Replace(
		config,
		"Trusted certificate provider acceptance test",
		"Updated certificate provider acceptance test",
		1,
	)
	updated = strings.Replace(updated, "computerCertStoreRoot", "computerCertStoreIntermediate", 1)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			10*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{Config: config, Check: resource.ComposeTestCheckFunc(
				check.That(resourceAddress).ExistsInGraph(testResource),
				check.That(resourceAddress).Key("assignments.#").HasValue("1"))},
			importStep(),
			{
				Config: updated,
				Check: check.That(resourceAddress).
					Key("destination_store").
					HasValue("computerCertStoreIntermediate"),
			},
			importStep(),
			{Config: updated, PlanOnly: true},
		},
	})
}
