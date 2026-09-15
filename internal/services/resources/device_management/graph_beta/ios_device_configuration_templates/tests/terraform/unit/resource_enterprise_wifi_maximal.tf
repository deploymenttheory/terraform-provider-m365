resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "enterprise_wifi_example" {
  display_name = "unit-test-iOS-enterprise-wifi-example"
  description  = "Enterprise Wi-Fi with EAP-TLS authentication"

  # No pre_shared_key here: enterprise networks authenticate via EAP, so the attribute is
  # deliberately not exposed on this block even though the SDK type inherits it.
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

    outer_identity_privacy_temporary_value = "anonymous"
    username_format_string                 = "{{UserPrincipalName}}"
    password_format_string                 = "unit-test-password-format"

    trusted_server_certificate_names = ["radius.example.com", "radius2.example.com"]

    # Multi-value @odata.bind: serialized as a JSON array. Bare GUIDs are expanded by the provider.
    root_certificates_for_server_validation_odata_bind = [
      "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
      "aaaaaaaa-bbbb-cccc-dddd-ffffffffffff"
    ]

    identity_certificate_for_client_authentication_odata_bind = "ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb"
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
