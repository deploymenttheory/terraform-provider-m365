package graphBetaIosDeviceConfigurationTemplates_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/check"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/destroy"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/testlog"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/constants"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/helpers"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	graphBetaIosDeviceConfigurationTemplates "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/device_management/graph_beta/ios_device_configuration_templates"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var (
	resourceType = graphBetaIosDeviceConfigurationTemplates.ResourceName
	testResource = graphBetaIosDeviceConfigurationTemplates.IosDeviceConfigurationTemplatesTestResource{}
)

func loadAcceptanceTestTerraform(t *testing.T, filename string) string {
	t.Helper()
	config, err := helpers.ParseHCLFile("tests/terraform/acceptance/" + filename)
	if err != nil {
		t.Skipf("skipping acceptance test: fixture file not found: %s", filename)
		return ""
	}
	return config
}

// Custom Configuration Tests
func TestAccResourceIosDeviceConfigurationTemplates_01_CustomConfiguration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating custom configuration")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_custom_configuration_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".custom_configuration_example").ExistsInGraph(testResource),
					check.That(resourceType+".custom_configuration_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".custom_configuration_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-custom-config-[a-z0-9]{8}$`)),
					check.That(resourceType+".custom_configuration_example").Key("description").HasValue("Example custom configuration template for iOS devices"),
					check.That(resourceType+".custom_configuration_example").Key("custom_configuration.payload_file_name").HasValue("com.example.custom.mobileconfig"),
					check.That(resourceType+".custom_configuration_example").Key("custom_configuration.payload_name").HasValue("Custom Configuration Example"),
					check.That(resourceType+".custom_configuration_example").Key("custom_configuration.payload").Exists(),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".custom_configuration_example").Key("assignments.#").HasValue("4"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing custom configuration")
				},
				ResourceName:      resourceType + ".custom_configuration_example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Trusted Root Certificate Tests
func TestAccResourceIosDeviceConfigurationTemplates_02_TrustedRootCertificate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating trusted root certificate configuration")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_trusted_root_certificate_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".trusted_cert_example").ExistsInGraph(testResource),
					check.That(resourceType+".trusted_cert_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".trusted_cert_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-trusted-root-cert-[a-z0-9]{8}$`)),
					check.That(resourceType+".trusted_cert_example").Key("description").HasValue("Install company root certificate for secure connections"),
					check.That(resourceType+".trusted_cert_example").Key("trusted_certificate.cert_file_name").HasValue("MicrosoftRootCertificateAuthority2011.cer"),
					check.That(resourceType+".trusted_cert_example").Key("trusted_certificate.trusted_root_certificate").Exists(),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".trusted_cert_example").Key("assignments.#").HasValue("4"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing trusted root certificate configuration")
				},
				ResourceName:      resourceType + ".trusted_cert_example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// SCEP Certificate Tests
func TestAccResourceIosDeviceConfigurationTemplates_04_ScepCertificate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating SCEP certificate profile")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_scep_certificate_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".scep_cert_example").ExistsInGraph(testResource),
					check.That(resourceType+".scep_cert_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".scep_cert_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-scep-cert-[a-z0-9]{8}$`)),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.certificate_store").HasValue("machine"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.subject_name_format").HasValue("custom"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.key_size").HasValue("size2048"),
					// Confirms Graph round-trips the comma-joined bitmask intact. subject_alternative_name_type
					// is not asserted: it is unset in the configuration, because a machine-store (device)
					// certificate accepts only device attributes in its SAN. Its bitmask round-trip is
					// covered by the unit tests.
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.key_usage.#").HasValue("2"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.custom_subject_alternative_names.#").HasValue("1"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.root_certificate_odata_bind").Exists(),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".scep_cert_example").Key("assignments.#").HasValue("2"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing SCEP certificate profile")
				},
				ResourceName:      resourceType + ".scep_cert_example",
				ImportState:       true,
				ImportStateVerify: true,
				// Graph does not echo additionalData, so the bind reference cannot be recovered on
				// import. On a real tenant the expanded rootCertificate navigation property yields
				// the full URL form rather than the bare GUID the configuration supplies.
				ImportStateVerifyIgnore: []string{"scep_certificate.root_certificate_odata_bind"},
			},
		},
	})
}

// PKCS Certificate Tests
func TestAccResourceIosDeviceConfigurationTemplates_05_PkcsCertificate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating PKCS certificate profile")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_pkcs_certificate_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".pkcs_cert_example").ExistsInGraph(testResource),
					check.That(resourceType+".pkcs_cert_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".pkcs_cert_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-pkcs-cert-[a-z0-9]{8}$`)),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certificate_store").HasValue("user"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.subject_name_format").HasValue("commonNameIncludingEmail"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certification_authority_name").HasValue("ExampleCA-CA"),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".pkcs_cert_example").Key("assignments.#").HasValue("2"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing PKCS certificate profile")
				},
				ResourceName:      resourceType + ".pkcs_cert_example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Enterprise Wi-Fi Tests
func TestAccResourceIosDeviceConfigurationTemplates_06_EnterpriseWifi(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating enterprise Wi-Fi profile")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_enterprise_wifi_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".enterprise_wifi_example").ExistsInGraph(testResource),
					check.That(resourceType+".enterprise_wifi_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".enterprise_wifi_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-ent-wifi-[a-z0-9]{8}$`)),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.ssid").HasValue("CorpSecure"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.wifi_security_type").HasValue("wpa2Enterprise"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.eap_type").HasValue("eapTls"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.authentication_method").HasValue("certificate"),
					// Confirms Graph accepts the multi-value @odata.bind array.
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.root_certificates_for_server_validation_odata_bind.#").HasValue("1"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.identity_certificate_for_client_authentication_odata_bind").Exists(),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".enterprise_wifi_example").Key("assignments.#").HasValue("2"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing enterprise Wi-Fi profile")
				},
				ResourceName:      resourceType + ".enterprise_wifi_example",
				ImportState:       true,
				ImportStateVerify: true,
				// Graph does not echo additionalData, so neither bind reference nor the write-only
				// password can be recovered on import.
				ImportStateVerifyIgnore: []string{
					"enterprise_wifi.root_certificates_for_server_validation_odata_bind",
					"enterprise_wifi.identity_certificate_for_client_authentication_odata_bind",
					"enterprise_wifi.password_format_string",
				},
			},
		},
	})
}

// EAS Email Tests
func TestAccResourceIosDeviceConfigurationTemplates_07_EasEmail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating EAS email profile")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_eas_email_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".eas_email_example").ExistsInGraph(testResource),
					check.That(resourceType+".eas_email_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".eas_email_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-eas-email-[a-z0-9]{8}$`)),
					check.That(resourceType+".eas_email_example").Key("eas_email.host_name").HasValue("outlook.office365.com"),
					check.That(resourceType+".eas_email_example").Key("eas_email.authentication_method").HasValue("usernameAndPassword"),
					// Confirms Graph round-trips the comma-joined easServices bitmask intact.
					check.That(resourceType+".eas_email_example").Key("eas_email.eas_services.#").HasValue("3"),
					check.That(resourceType+".eas_email_example").Key("eas_email.require_ssl").HasValue("true"),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".eas_email_example").Key("assignments.#").HasValue("2"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing EAS email profile")
				},
				ResourceName:      resourceType + ".eas_email_example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// VPN Tests
func TestAccResourceIosDeviceConfigurationTemplates_08_Vpn(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating VPN profile")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_vpn_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".vpn_example").ExistsInGraph(testResource),
					check.That(resourceType+".vpn_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".vpn_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-vpn-[a-z0-9]{8}$`)),
					check.That(resourceType+".vpn_example").Key("vpn.connection_name").HasValue("Corporate VPN"),
					check.That(resourceType+".vpn_example").Key("vpn.connection_type").HasValue("ciscoAnyConnectV2"),
					// Confirms Graph round-trips the nested server/proxy objects and collections.
					check.That(resourceType+".vpn_example").Key("vpn.server.address").HasValue("vpn.example.com"),
					check.That(resourceType+".vpn_example").Key("vpn.proxy_server.port").HasValue("8080"),
					check.That(resourceType+".vpn_example").Key("vpn.safari_domains.#").HasValue("1"),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".vpn_example").Key("assignments.#").HasValue("2"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing VPN profile")
				},
				ResourceName:      resourceType + ".vpn_example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Wi-Fi Tests
func TestAccResourceIosDeviceConfigurationTemplates_03_Wifi(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { mocks.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: mocks.TestAccProtoV6ProviderFactories,
		CheckDestroy: destroy.CheckDestroyedAllFunc(
			testResource,
			resourceType,
			30*time.Second,
		),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: constants.ExternalProviderRandomVersion,
			},
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Creating wifi configuration")
				},
				Config: loadAcceptanceTestTerraform(t, "resource_wifi_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					func(_ *terraform.State) error {
						testlog.WaitForConsistency("iOS/iPadOS device configuration", 30*time.Second)
						time.Sleep(30 * time.Second)
						return nil
					},
					check.That(resourceType+".wifi_example").ExistsInGraph(testResource),
					check.That(resourceType+".wifi_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".wifi_example").Key("display_name").MatchesRegex(regexp.MustCompile(`^acc-test-iOS-wifi-[a-z0-9]{8}$`)),
					check.That(resourceType+".wifi_example").Key("wifi.network_name").HasValue("Corporate Wi-Fi"),
					check.That(resourceType+".wifi_example").Key("wifi.ssid").HasValue("CorpWiFi"),
					check.That(resourceType+".wifi_example").Key("wifi.wifi_security_type").HasValue("wpa2Personal"),
					check.That(resourceType+".wifi_example").Key("wifi.proxy_settings").HasValue("manual"),
					// Assignment assertion disabled: the group dependencies it relied on are commented
					// out in the .tf while group hard-delete is failing on the test tenant. See the
					// note at the top of the corresponding tests/terraform/acceptance file.
					// check.That(resourceType+".wifi_example").Key("assignments.#").HasValue("4"),
				),
			},
			{
				PreConfig: func() {
					testlog.StepAction(resourceType, "Importing wifi configuration")
				},
				ResourceName:      resourceType + ".wifi_example",
				ImportState:       true,
				ImportStateVerify: true,
				// Graph never returns the pre-shared key, so it cannot be verified against
				// imported state.
				ImportStateVerifyIgnore: []string{"wifi.pre_shared_key"},
			},
		},
	})
}
