package graphBetaIosDeviceConfigurationTemplates_test

import (
	"regexp"
	"testing"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/check"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/helpers"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/mocks"
	iosConfigMocks "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/resources/device_management/graph_beta/ios_device_configuration_templates/mocks"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/jarcoal/httpmock"
)

func setupMockEnvironment() (*mocks.Mocks, *iosConfigMocks.IosDeviceConfigurationTemplatesMock) {
	httpmock.Activate()
	mockClient := mocks.NewMocks()
	mockClient.AuthMocks.RegisterMocks()
	configMock := &iosConfigMocks.IosDeviceConfigurationTemplatesMock{}
	configMock.RegisterMocks()
	return mockClient, configMock
}

func setupErrorMockEnvironment() (*mocks.Mocks, *iosConfigMocks.IosDeviceConfigurationTemplatesMock) {
	httpmock.Activate()
	mockClient := mocks.NewMocks()
	mockClient.AuthMocks.RegisterMocks()
	configMock := &iosConfigMocks.IosDeviceConfigurationTemplatesMock{}
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

func TestUnitResourceIosDeviceConfigurationTemplates_01_CustomConfiguration(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_custom_configuration_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".custom_configuration_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".custom_configuration_example").Key("display_name").HasValue("unit-test-iOS-custom-configuration-example"),
					check.That(resourceType+".custom_configuration_example").Key("description").HasValue("Example custom configuration template for iOS devices"),
					check.That(resourceType+".custom_configuration_example").Key("custom_configuration.payload_file_name").HasValue("com.example.custom.mobileconfig"),
					check.That(resourceType+".custom_configuration_example").Key("custom_configuration.payload_name").HasValue("Custom Configuration Example"),
					check.That(resourceType+".custom_configuration_example").Key("custom_configuration.payload").Exists(),
					check.That(resourceType+".custom_configuration_example").Key("role_scope_tag_ids.#").HasValue("1"),
					check.That(resourceType+".custom_configuration_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_02_TrustedRootCertificate(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_trusted_root_certificate_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".trusted_cert_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".trusted_cert_example").Key("display_name").HasValue("unit-test-iOS-trusted-root-certificate-example"),
					check.That(resourceType+".trusted_cert_example").Key("description").HasValue("Install company root certificate for secure connections"),
					check.That(resourceType+".trusted_cert_example").Key("trusted_certificate.cert_file_name").HasValue("MicrosoftRootCertificateAuthority2011.cer"),
					check.That(resourceType+".trusted_cert_example").Key("trusted_certificate.trusted_root_certificate").Exists(),
					check.That(resourceType+".trusted_cert_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_03_Wifi(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_wifi_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".wifi_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".wifi_example").Key("display_name").HasValue("unit-test-iOS-wifi-example"),
					check.That(resourceType+".wifi_example").Key("wifi.network_name").HasValue("Corporate Wi-Fi"),
					check.That(resourceType+".wifi_example").Key("wifi.ssid").HasValue("CorpWiFi"),
					check.That(resourceType+".wifi_example").Key("wifi.connect_automatically").HasValue("true"),
					check.That(resourceType+".wifi_example").Key("wifi.connect_when_network_name_is_hidden").HasValue("false"),
					check.That(resourceType+".wifi_example").Key("wifi.wifi_security_type").HasValue("wpa2Personal"),
					check.That(resourceType+".wifi_example").Key("wifi.proxy_settings").HasValue("manual"),
					check.That(resourceType+".wifi_example").Key("wifi.proxy_manual_address").HasValue("proxy.example.com"),
					check.That(resourceType+".wifi_example").Key("wifi.proxy_manual_port").HasValue("8080"),
					// Graph does not return the pre-shared key; it must survive the read from
					// configuration rather than being cleared out of state.
					check.That(resourceType+".wifi_example").Key("wifi.pre_shared_key").HasValue("unit-test-pre-shared-key"),
					check.That(resourceType+".wifi_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_04_ScepCertificate(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_scep_certificate_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".scep_cert_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".scep_cert_example").Key("display_name").HasValue("unit-test-iOS-scep-certificate-example"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.renewal_threshold_percentage").HasValue("20"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.certificate_store").HasValue("machine"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.certificate_validity_period_scale").HasValue("years"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.certificate_validity_period_value").HasValue("2"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.subject_name_format").HasValue("custom"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.key_size").HasValue("size2048"),
					// keyUsage round-trips through a comma-joined bitmask on the wire and must
					// decompose back into both elements.
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.key_usage.#").HasValue("2"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.key_usage.*").ContainsTypeSetElement("keyEncipherment"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.key_usage.*").ContainsTypeSetElement("digitalSignature"),
					// subjectAlternativeNameType is likewise a bitmask.
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.subject_alternative_name_type.#").HasValue("2"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.subject_alternative_name_type.*").ContainsTypeSetElement("emailAddress"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.subject_alternative_name_type.*").ContainsTypeSetElement("universalResourceIdentifier"),
					// A bare GUID in configuration must be preserved verbatim in state rather than
					// rewritten to the expanded URL sent on the wire.
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.root_certificate_odata_bind").HasValue("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.custom_subject_alternative_names.#").HasValue("2"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.extended_key_usages.#").HasValue("2"),
					check.That(resourceType+".scep_cert_example").Key("scep_certificate.scep_server_urls.#").HasValue("2"),
					check.That(resourceType+".scep_cert_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_05_PkcsCertificate(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_pkcs_certificate_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".pkcs_cert_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".pkcs_cert_example").Key("display_name").HasValue("unit-test-iOS-pkcs-certificate-example"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.renewal_threshold_percentage").HasValue("30"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certificate_store").HasValue("user"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certificate_validity_period_scale").HasValue("months"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certificate_validity_period_value").HasValue("12"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.subject_name_format").HasValue("commonNameIncludingEmail"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certification_authority").HasValue("ExampleCA.example.com"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certification_authority_name").HasValue("ExampleCA-CA"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.certificate_template_name").HasValue("UserAuthentication"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.subject_alternative_name_type.#").HasValue("1"),
					check.That(resourceType+".pkcs_cert_example").Key("pkcs_certificate.custom_subject_alternative_names.#").HasValue("1"),
					check.That(resourceType+".pkcs_cert_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

// The @odata.bind reference is the most drift-prone value in this resource: Graph accepts it on
// write but never echoes it back, so it has to be recovered from prior state on every read.
// Applying the same configuration twice must produce an empty plan — a regression here shows up as
// a non-empty second plan rather than as a failed assertion.
func TestUnitResourceIosDeviceConfigurationTemplates_06_ScepCertificateNoDriftOnRefresh(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_scep_certificate_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".scep_cert_example").Key("scep_certificate.root_certificate_odata_bind").HasValue("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
				),
			},
			{
				// Re-planning the unchanged configuration must be a no-op.
				Config:             loadUnitTestTerraform("resource_scep_certificate_maximal.tf"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_07_EnterpriseWifi(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_enterprise_wifi_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".enterprise_wifi_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".enterprise_wifi_example").Key("display_name").HasValue("unit-test-iOS-enterprise-wifi-example"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.network_name").HasValue("Corporate 802.1x"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.ssid").HasValue("CorpSecure"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.wifi_security_type").HasValue("wpa2Enterprise"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.eap_type").HasValue("eapTls"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.authentication_method").HasValue("certificate"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.outer_identity_privacy_temporary_value").HasValue("anonymous"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.trusted_server_certificate_names.#").HasValue("2"),
					// Write-only secret: must survive the read from configuration.
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.password_format_string").HasValue("unit-test-password-format"),
					// Multi-value @odata.bind array: both entries preserved as configured bare GUIDs.
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.root_certificates_for_server_validation_odata_bind.#").HasValue("2"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.root_certificates_for_server_validation_odata_bind.*").ContainsTypeSetElement("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
					check.That(resourceType+".enterprise_wifi_example").Key("enterprise_wifi.identity_certificate_for_client_authentication_odata_bind").HasValue("ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb"),
					check.That(resourceType+".enterprise_wifi_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_08_EasEmail(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_eas_email_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".eas_email_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".eas_email_example").Key("display_name").HasValue("unit-test-iOS-eas-email-example"),
					check.That(resourceType+".eas_email_example").Key("eas_email.account_name").HasValue("Corporate Email"),
					check.That(resourceType+".eas_email_example").Key("eas_email.host_name").HasValue("outlook.office365.com"),
					check.That(resourceType+".eas_email_example").Key("eas_email.authentication_method").HasValue("certificate"),
					// easServices round-trips through a comma-joined bitmask and must decompose
					// back into all three elements.
					check.That(resourceType+".eas_email_example").Key("eas_email.eas_services.#").HasValue("3"),
					check.That(resourceType+".eas_email_example").Key("eas_email.eas_services.*").ContainsTypeSetElement("calendars"),
					check.That(resourceType+".eas_email_example").Key("eas_email.eas_services.*").ContainsTypeSetElement("contacts"),
					check.That(resourceType+".eas_email_example").Key("eas_email.eas_services.*").ContainsTypeSetElement("email"),
					check.That(resourceType+".eas_email_example").Key("eas_email.duration_of_email_to_sync").HasValue("oneMonth"),
					check.That(resourceType+".eas_email_example").Key("eas_email.email_address_source").HasValue("primarySmtpAddress"),
					check.That(resourceType+".eas_email_example").Key("eas_email.username_source").HasValue("userPrincipalName"),
					check.That(resourceType+".eas_email_example").Key("eas_email.username_aad_source").HasValue("userPrincipalName"),
					check.That(resourceType+".eas_email_example").Key("eas_email.require_ssl").HasValue("true"),
					check.That(resourceType+".eas_email_example").Key("eas_email.block_moving_messages_to_other_email_accounts").HasValue("true"),
					check.That(resourceType+".eas_email_example").Key("eas_email.identity_certificate_odata_bind").HasValue("ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb"),
					check.That(resourceType+".eas_email_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

// The multi-value @odata.bind array and the write-only password are both recovered from prior
// state, so an unchanged configuration must re-plan clean.
func TestUnitResourceIosDeviceConfigurationTemplates_09_EnterpriseWifiNoDriftOnRefresh(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_enterprise_wifi_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".enterprise_wifi_example").Key("enterprise_wifi.root_certificates_for_server_validation_odata_bind.#").HasValue("2"),
				),
			},
			{
				Config:             loadUnitTestTerraform("resource_enterprise_wifi_maximal.tf"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_10_Vpn(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_vpn_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".vpn_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".vpn_example").Key("display_name").HasValue("unit-test-iOS-vpn-example"),
					check.That(resourceType+".vpn_example").Key("vpn.connection_name").HasValue("Corporate VPN"),
					check.That(resourceType+".vpn_example").Key("vpn.connection_type").HasValue("ciscoAnyConnectV2"),
					check.That(resourceType+".vpn_example").Key("vpn.authentication_method").HasValue("certificate"),
					check.That(resourceType+".vpn_example").Key("vpn.provider_type").HasValue("packetTunnel"),
					check.That(resourceType+".vpn_example").Key("vpn.enable_split_tunneling").HasValue("true"),
					check.That(resourceType+".vpn_example").Key("vpn.enable_per_app").HasValue("true"),
					check.That(resourceType+".vpn_example").Key("vpn.disconnect_on_idle_timer_in_seconds").HasValue("300"),
					check.That(resourceType+".vpn_example").Key("vpn.safari_domains.#").HasValue("2"),
					// Single-nested objects round-trip.
					check.That(resourceType+".vpn_example").Key("vpn.server.address").HasValue("vpn.example.com"),
					check.That(resourceType+".vpn_example").Key("vpn.server.is_default_server").HasValue("true"),
					check.That(resourceType+".vpn_example").Key("vpn.proxy_server.address").HasValue("proxy.example.com"),
					check.That(resourceType+".vpn_example").Key("vpn.proxy_server.port").HasValue("8080"),
					// Collections.
					check.That(resourceType+".vpn_example").Key("vpn.targeted_mobile_apps.#").HasValue("2"),
					// customData uses "key", customKeyValueData uses "name" — distinct Graph
					// properties that are easy to conflate.
					check.That(resourceType+".vpn_example").Key("vpn.custom_data.#").HasValue("1"),
					check.That(resourceType+".vpn_example").Key("vpn.custom_key_value_data.#").HasValue("1"),
					check.That(resourceType+".vpn_example").Key("vpn.identity_certificate_odata_bind").HasValue("ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb"),
					check.That(resourceType+".vpn_example").Key("assignments.#").HasValue("4"),
				),
			},
		},
	})
}

// ikEv2 maps to iosikEv2VpnConfiguration, a separate Graph type with roughly 23 extra properties
// this resource does not model. The validator must reject it rather than silently creating a profile
// that would round-trip lossily.
func TestUnitResourceIosDeviceConfigurationTemplates_11_VpnRejectsIkEv2(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      loadUnitTestTerraform("resource_vpn_ikev2_rejected.tf"),
				ExpectError: regexp.MustCompile(`(?s)connection_type.*value must be one of`),
			},
		},
	})
}

// The identity certificate @odata.bind is recovered from prior state, so an unchanged configuration
// must re-plan clean.
func TestUnitResourceIosDeviceConfigurationTemplates_12_VpnNoDriftOnRefresh(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_vpn_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".vpn_example").Key("vpn.identity_certificate_odata_bind").HasValue("ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb"),
				),
			},
			{
				Config:             loadUnitTestTerraform("resource_vpn_maximal.tf"),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_13_CreateWithError(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupErrorMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      loadUnitTestTerraform("resource_custom_configuration_maximal.tf"),
				ExpectError: regexp.MustCompile("Error creating iOS/iPadOS device configuration template"),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_14_Update(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_custom_configuration_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".custom_configuration_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".custom_configuration_example").Key("display_name").HasValue("unit-test-iOS-custom-configuration-example"),
				),
			},
			{
				Config: loadUnitTestTerraform("resource_custom_configuration_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType+".custom_configuration_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
					check.That(resourceType+".custom_configuration_example").Key("display_name").HasValue("unit-test-iOS-custom-configuration-example"),
				),
			},
		},
	})
}

func TestUnitResourceIosDeviceConfigurationTemplates_15_ImportState(t *testing.T) {
	mocks.SetupUnitTestEnvironment(t)
	_, configMock := setupMockEnvironment()
	defer httpmock.DeactivateAndReset()
	defer configMock.CleanupMockState()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: mocks.TestUnitTestProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: loadUnitTestTerraform("resource_custom_configuration_maximal.tf"),
				Check: resource.ComposeTestCheckFunc(
					check.That(resourceType + ".custom_configuration_example").Key("id").MatchesRegex(regexp.MustCompile(`^[0-9a-fA-F-]+$`)),
				),
			},
			{
				ResourceName:      resourceType + ".custom_configuration_example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
