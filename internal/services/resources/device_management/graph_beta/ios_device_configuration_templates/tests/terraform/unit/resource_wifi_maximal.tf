resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "wifi_example" {
  display_name = "unit-test-iOS-wifi-example"
  description  = "Corporate Wi-Fi network for iOS devices"

  wifi = {
    network_name                        = "Corporate Wi-Fi"
    ssid                                = "CorpWiFi"
    connect_automatically               = true
    connect_when_network_name_is_hidden = false
    wifi_security_type                  = "wpa2Personal"
    pre_shared_key                      = "unit-test-pre-shared-key"
    disable_mac_address_randomization   = false
    proxy_settings                      = "manual"
    proxy_manual_address                = "proxy.example.com"
    proxy_manual_port                   = 8080
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
