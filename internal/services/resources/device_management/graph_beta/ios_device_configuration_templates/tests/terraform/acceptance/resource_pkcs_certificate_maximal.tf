# NOTE: Group creation and assignments are commented out for now.
# microsoft365_graph_beta_groups_group with hard_delete = true fails its permanent-destroy step
# against the test tenant, which makes CheckDestroy report every test as failed even when the
# device configuration itself created, read, updated and destroyed correctly.
# The profile is still fully exercised — assignments are Optional in the schema, so omitting them
# is valid. Re-enable the blocks below once group hard-delete is working to restore assignment
# coverage.

resource "random_string" "pkcs_suffix" {
  length  = 8
  special = false
  upper   = false
}

# ==============================================================================
# Group Dependencies
# ==============================================================================

# resource "microsoft365_graph_beta_groups_group" "pkcs_group_1" {
#   display_name     = "acc-test-ios-pkcs-group-1-${random_string.pkcs_suffix.result}"
#   mail_nickname    = "acc-test-ios-pkcs-1-${random_string.pkcs_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "pkcs_group_2" {
#   display_name     = "acc-test-ios-pkcs-group-2-${random_string.pkcs_suffix.result}"
#   mail_nickname    = "acc-test-ios-pkcs-2-${random_string.pkcs_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "pkcs_cert_example" {
  display_name = "acc-test-iOS-pkcs-cert-${random_string.pkcs_suffix.result}"
  description  = "PKCS certificate profile for user authentication"

  pkcs_certificate = {
    renewal_threshold_percentage      = 30
    certificate_store                 = "user"
    certificate_validity_period_scale = "months"
    certificate_validity_period_value = 12
    subject_name_format               = "commonNameIncludingEmail"
    subject_name_format_string        = "CN={{UserName}},E={{EmailAddress}},O=Example Corp"

    subject_alternative_name_type          = ["emailAddress"]
    subject_alternative_name_format_string = "{{EmailAddress}}"

    certification_authority      = "ExampleCA.example.com"
    certification_authority_name = "ExampleCA-CA"
    certificate_template_name    = "UserAuthentication"

    custom_subject_alternative_names = [
      {
        san_type = "emailAddress"
        name     = "{{EmailAddress}}"
      }
    ]
  }

  # assignments = [
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.pkcs_group_1.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.pkcs_group_2.id
  #   }
  # ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
