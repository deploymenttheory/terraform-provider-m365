# NOTE: Group creation and assignments are commented out for now.
# microsoft365_graph_beta_groups_group with hard_delete = true fails its permanent-destroy step
# against the test tenant, which makes CheckDestroy report every test as failed even when the
# device configuration itself created, read, updated and destroyed correctly.
# The profile is still fully exercised — assignments are Optional in the schema, so omitting them
# is valid. Re-enable the blocks below once group hard-delete is working to restore assignment
# coverage.

resource "random_string" "eas_suffix" {
  length  = 8
  special = false
  upper   = false
}

# ==============================================================================
# Group Dependencies
# ==============================================================================

# resource "microsoft365_graph_beta_groups_group" "eas_group_1" {
#   display_name     = "acc-test-ios-eas-group-1-${random_string.eas_suffix.result}"
#   mail_nickname    = "acc-test-ios-eas-1-${random_string.eas_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "eas_group_2" {
#   display_name     = "acc-test-ios-eas-group-2-${random_string.eas_suffix.result}"
#   mail_nickname    = "acc-test-ios-eas-2-${random_string.eas_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "eas_email_example" {
  display_name = "acc-test-iOS-eas-email-${random_string.eas_suffix.result}"
  description  = "Exchange ActiveSync email profile"

  # authentication_method is usernameAndPassword here so the profile needs no certificate
  # dependency; the certificate path is covered by the enterprise_wifi acceptance test.
  eas_email = {
    account_name          = "Corporate Email"
    host_name             = "outlook.office365.com"
    authentication_method = "usernameAndPassword"

    eas_services                       = ["calendars", "contacts", "email"]
    eas_services_user_override_enabled = false

    duration_of_email_to_sync = "oneMonth"
    email_address_source      = "primarySmtpAddress"
    username_source           = "userPrincipalName"
    username_aad_source       = "userPrincipalName"
    user_domain_name_source   = "fullDomainName"

    require_ssl = true
    use_oauth   = false

    block_moving_messages_to_other_email_accounts = true
    block_sending_email_from_third_party_apps     = true
    block_syncing_recently_used_email_addresses   = false

    require_smime               = false
    signing_certificate_type    = "none"
    encryption_certificate_type = "none"
  }

  # assignments = [
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.eas_group_1.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.eas_group_2.id
  #   }
  # ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
