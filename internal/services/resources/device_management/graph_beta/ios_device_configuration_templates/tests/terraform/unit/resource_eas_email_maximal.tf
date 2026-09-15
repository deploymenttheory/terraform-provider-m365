resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "eas_email_example" {
  display_name = "unit-test-iOS-eas-email-example"
  description  = "Exchange ActiveSync email profile"

  eas_email = {
    account_name          = "Corporate Email"
    host_name             = "outlook.office365.com"
    authentication_method = "certificate"

    # Bitmask on the wire: Graph joins these with commas into a single value.
    eas_services                       = ["calendars", "contacts", "email"]
    eas_services_user_override_enabled = false

    duration_of_email_to_sync = "oneMonth"
    email_address_source      = "primarySmtpAddress"
    username_source           = "userPrincipalName"
    # Note this uses a wider enum than username_source (it also allows samAccountName).
    username_aad_source     = "userPrincipalName"
    user_domain_name_source = "fullDomainName"

    require_ssl = true
    use_oauth   = false

    block_moving_messages_to_other_email_accounts = true
    block_sending_email_from_third_party_apps     = true
    block_syncing_recently_used_email_addresses   = false

    require_smime               = false
    signing_certificate_type    = "none"
    encryption_certificate_type = "none"

    identity_certificate_odata_bind = "ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb"
  }

  role_scope_tag_ids = ["00000000-0000-0000-0000-000000000001"]

  assignments = [
    {
      type        = "groupAssignmentTarget"
      group_id    = "00000000-0000-0000-0000-000000000002"
      filter_id   = "00000000-0000-0000-0000-000000000003"
      filter_type = "include"
    },
    {
      type        = "groupAssignmentTarget"
      group_id    = "00000000-0000-0000-0000-000000000002"
      filter_id   = "00000000-0000-0000-0000-000000000003"
      filter_type = "exclude"
    },
    {
      type     = "exclusionGroupAssignmentTarget"
      group_id = "00000000-0000-0000-0000-000000000002"
    },
    {
      type     = "exclusionGroupAssignmentTarget"
      group_id = "00000000-0000-0000-0000-000000000004"
    }
  ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
