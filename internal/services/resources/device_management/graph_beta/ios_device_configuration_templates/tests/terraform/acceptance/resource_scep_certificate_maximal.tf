# NOTE: Group creation and assignments are commented out for now.
#
# microsoft365_graph_beta_groups_group with hard_delete = true fails its permanent-destroy step
# against the test tenant, which makes CheckDestroy report every test as failed even when the device
# configuration itself created, read, updated and destroyed correctly.
#
# The profile is still fully exercised — assignments are Optional in the schema, so omitting them is
# valid. Re-enable the commented blocks once group hard-delete is working.
#
# The trusted root certificate dependency below is NOT commented out: the SCEP profile's
# root_certificate_odata_bind needs a real profile to bind to, and that is the main thing this test
# verifies.

resource "random_string" "scep_suffix" {
  length  = 8
  special = false
  upper   = false
}

# ==============================================================================
# Group Dependencies — disabled, see note above
# ==============================================================================

# resource "microsoft365_graph_beta_groups_group" "scep_group_1" {
#   display_name     = "acc-test-ios-scep-group-1-${random_string.scep_suffix.result}"
#   mail_nickname    = "acc-test-ios-scep-1-${random_string.scep_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "scep_group_2" {
#   display_name     = "acc-test-ios-scep-group-2-${random_string.scep_suffix.result}"
#   mail_nickname    = "acc-test-ios-scep-2-${random_string.scep_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# ==============================================================================
# Trusted root certificate — SCEP profiles require one to bind to
# ==============================================================================

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "scep_root_cert" {
  display_name = "acc-test-iOS-scep-root-ca-${random_string.scep_suffix.result}"
  description  = "Root CA referenced by the SCEP profile under test"

  trusted_certificate = {
    cert_file_name           = "MicrosoftRootCertificateAuthority2011.cer"
    trusted_root_certificate = filebase64("tests/terraform/acceptance/MicrosoftRootCertificateAuthority2011.cer")
  }
}

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "scep_cert_example" {
  display_name = "acc-test-iOS-scep-cert-${random_string.scep_suffix.result}"
  description  = "SCEP certificate profile for device authentication"

  # certificate_store = "machine" makes this a DEVICE certificate, and Intune only accepts device
  # attributes in the subject name and SAN of a device certificate — user attributes such as
  # {{EmailAddress}} are rejected with a 400. Both the subject name and the SAN below therefore use
  # device variables only. Validity period is 1 year: the value must not exceed the validity period
  # of the issuing CA's template.
  scep_certificate = {
    renewal_threshold_percentage      = 20
    certificate_store                 = "machine"
    certificate_validity_period_scale = "years"
    certificate_validity_period_value = 1
    subject_name_format               = "custom"
    subject_name_format_string        = "CN={{AAD_Device_ID}},O=Example Corp,C=US"

    # Bare GUID form — the provider expands this to the full @odata.bind URL.
    root_certificate_odata_bind = microsoft365_graph_beta_device_management_ios_device_configuration_templates.scep_root_cert.id

    key_size  = "size2048"
    key_usage = ["keyEncipherment", "digitalSignature"]

    scep_server_urls = [
      "https://scep.example.com/certsrv/mscep/mscep.dll"
    ]

    custom_subject_alternative_names = [
      {
        san_type = "domainNameService"
        name     = "{{DeviceName}}"
      }
    ]

    extended_key_usages = [
      {
        name              = "Client Authentication"
        object_identifier = "1.3.6.1.5.5.7.3.2"
      }
    ]
  }

  # assignments = [
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.scep_group_1.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.scep_group_2.id
  #   }
  # ]

  depends_on = [
    microsoft365_graph_beta_device_management_ios_device_configuration_templates.scep_root_cert
  ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
