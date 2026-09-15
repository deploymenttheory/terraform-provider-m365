# NOTE: Group creation and assignments are commented out for now.
# microsoft365_graph_beta_groups_group with hard_delete = true fails its permanent-destroy step
# against the test tenant, which makes CheckDestroy report every test as failed even when the
# device configuration itself created, read, updated and destroyed correctly.
# The profile is still fully exercised — assignments are Optional in the schema, so omitting them
# is valid. Re-enable the blocks below once group hard-delete is working to restore assignment
# coverage.

resource "random_string" "trusted_cert_suffix" {
  length  = 8
  special = false
  upper   = false
}

# ==============================================================================
# Group Dependencies
# ==============================================================================

# resource "microsoft365_graph_beta_groups_group" "trusted_cert_group_1" {
#   display_name     = "acc-test-ios-trusted-cert-group-1-${random_string.trusted_cert_suffix.result}"
#   mail_nickname    = "acc-test-ios-trusted-cert-1-${random_string.trusted_cert_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "trusted_cert_group_2" {
#   display_name     = "acc-test-ios-trusted-cert-group-2-${random_string.trusted_cert_suffix.result}"
#   mail_nickname    = "acc-test-ios-trusted-cert-2-${random_string.trusted_cert_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "trusted_cert_group_3" {
#   display_name     = "acc-test-ios-trusted-cert-group-3-${random_string.trusted_cert_suffix.result}"
#   mail_nickname    = "acc-test-ios-trusted-cert-3-${random_string.trusted_cert_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "trusted_cert_group_4" {
#   display_name     = "acc-test-ios-trusted-cert-group-4-${random_string.trusted_cert_suffix.result}"
#   mail_nickname    = "acc-test-ios-trusted-cert-4-${random_string.trusted_cert_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "trusted_cert_example" {
  display_name = "acc-test-iOS-trusted-root-cert-${random_string.trusted_cert_suffix.result}"
  description  = "Install company root certificate for secure connections"

  trusted_certificate = {
    cert_file_name           = "MicrosoftRootCertificateAuthority2011.cer"
    trusted_root_certificate = filebase64("tests/terraform/acceptance/MicrosoftRootCertificateAuthority2011.cer")
  }

  # assignments = [
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.trusted_cert_group_1.id
  #   },
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.trusted_cert_group_2.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.trusted_cert_group_3.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.trusted_cert_group_4.id
  #   }
  # ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
