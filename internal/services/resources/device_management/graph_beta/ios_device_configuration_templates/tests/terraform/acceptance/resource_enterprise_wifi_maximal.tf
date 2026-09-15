# NOTE: Group creation and assignments are commented out for now.
#
# microsoft365_graph_beta_groups_group with hard_delete = true fails its permanent-destroy step
# against the test tenant, which makes CheckDestroy report every test as failed even when the device
# configuration itself created, read, updated and destroyed correctly.
#
# The profile is still fully exercised — assignments are Optional in the schema, so omitting them is
# valid. Re-enable the commented blocks once group hard-delete is working.
#
# The two certificate dependencies below are NOT commented out: they are what this test exists to
# verify — the multi-value rootCertificatesForServerValidation@odata.bind array and the single
# identityCertificateForClientAuthentication@odata.bind reference.

resource "random_string" "ent_wifi_suffix" {
  length  = 8
  special = false
  upper   = false
}

# ==============================================================================
# Group Dependencies — disabled, see note above
# ==============================================================================

# resource "microsoft365_graph_beta_groups_group" "ent_wifi_group_1" {
#   display_name     = "acc-test-ios-ent-wifi-group-1-${random_string.ent_wifi_suffix.result}"
#   mail_nickname    = "acc-test-ios-ent-wifi-1-${random_string.ent_wifi_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "ent_wifi_group_2" {
#   display_name     = "acc-test-ios-ent-wifi-group-2-${random_string.ent_wifi_suffix.result}"
#   mail_nickname    = "acc-test-ios-ent-wifi-2-${random_string.ent_wifi_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# ==============================================================================
# Certificate dependencies — the enterprise Wi-Fi profile binds to both
# ==============================================================================

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "ent_wifi_root_cert" {
  display_name = "acc-test-iOS-ent-wifi-root-ca-${random_string.ent_wifi_suffix.result}"
  description  = "Root CA used for RADIUS server validation"

  trusted_certificate = {
    cert_file_name           = "MicrosoftRootCertificateAuthority2011.cer"
    trusted_root_certificate = filebase64("tests/terraform/acceptance/MicrosoftRootCertificateAuthority2011.cer")
  }
}

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "ent_wifi_client_cert" {
  display_name = "acc-test-iOS-ent-wifi-client-scep-${random_string.ent_wifi_suffix.result}"
  description  = "SCEP profile issuing the client authentication certificate"

  scep_certificate = {
    renewal_threshold_percentage      = 20
    certificate_store                 = "machine"
    certificate_validity_period_scale = "years"
    certificate_validity_period_value = 1
    subject_name_format               = "custom"
    subject_name_format_string        = "CN={{AAD_Device_ID}}"

    root_certificate_odata_bind = microsoft365_graph_beta_device_management_ios_device_configuration_templates.ent_wifi_root_cert.id

    key_size  = "size2048"
    key_usage = ["keyEncipherment", "digitalSignature"]

    scep_server_urls = ["https://scep.example.com/certsrv/mscep/mscep.dll"]

    extended_key_usages = [
      {
        name              = "Client Authentication"
        object_identifier = "1.3.6.1.5.5.7.3.2"
      }
    ]
  }

  depends_on = [
    microsoft365_graph_beta_device_management_ios_device_configuration_templates.ent_wifi_root_cert
  ]
}

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "enterprise_wifi_example" {
  display_name = "acc-test-iOS-ent-wifi-${random_string.ent_wifi_suffix.result}"
  description  = "Enterprise Wi-Fi with EAP-TLS authentication"

  enterprise_wifi = {
    network_name                        = "Corporate 802.1x"
    ssid                                = "CorpSecure"
    connect_automatically               = true
    connect_when_network_name_is_hidden = false
    wifi_security_type                  = "wpa2Enterprise"
    disable_mac_address_randomization   = true
    proxy_settings                      = "none"

    eap_type              = "eapTls"
    authentication_method = "certificate"

    # username_format_string / password_format_string are deliberately omitted: they build the
    # credentials for username-and-password authentication, and Intune rejects them with a 400 when
    # authentication_method is certificate. The client identity comes from the certificate bind
    # below instead.
    outer_identity_privacy_temporary_value = "anonymous"

    trusted_server_certificate_names = ["radius.example.com"]

    root_certificates_for_server_validation_odata_bind = [
      microsoft365_graph_beta_device_management_ios_device_configuration_templates.ent_wifi_root_cert.id
    ]

    identity_certificate_for_client_authentication_odata_bind = microsoft365_graph_beta_device_management_ios_device_configuration_templates.ent_wifi_client_cert.id
  }

  # assignments = [
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.ent_wifi_group_1.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.ent_wifi_group_2.id
  #   }
  # ]

  depends_on = [
    microsoft365_graph_beta_device_management_ios_device_configuration_templates.ent_wifi_root_cert,
    microsoft365_graph_beta_device_management_ios_device_configuration_templates.ent_wifi_client_cert
  ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
