resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "vpn_ikev2_rejected" {
  display_name = "unit-test-iOS-vpn-ikev2-rejected"
  description  = "Per-app VPN for corporate applications"

  # Note ikEv2 is not an accepted connection_type: it maps to iosikEv2VpnConfiguration, a separate
  # Graph type with ~23 additional properties this resource does not model.
  vpn = {
    connection_name       = "Corporate VPN"
    connection_type       = "ikEv2"
    authentication_method = "certificate"
    provider_type         = "packetTunnel"

    enable_split_tunneling = true
    enable_per_app         = true
    include_all_networks   = false
    exclude_local_networks = true

    safari_domains   = ["intranet.example.com", "portal.example.com"]
    excluded_domains = ["public.example.com"]

    disconnect_on_idle                  = true
    disconnect_on_idle_timer_in_seconds = 300
    disable_on_demand_user_override     = true
    opt_in_to_device_id_sharing         = false

    server = {
      address           = "vpn.example.com"
      description       = "Primary VPN gateway"
      is_default_server = true
    }

    proxy_server = {
      address = "proxy.example.com"
      port    = 8080
    }

    targeted_mobile_apps = [
      {
        name      = "Outlook"
        app_id    = "com.microsoft.Office.Outlook"
        publisher = "Microsoft Corporation"
      },
      {
        name      = "Teams"
        app_id    = "com.microsoft.skype.teams"
        publisher = "Microsoft Corporation"
      }
    ]

    # customData entries use "key"; customKeyValueData entries use "name". Distinct Graph properties.
    custom_data = [
      {
        key   = "vendorSetting"
        value = "enabled"
      }
    ]

    custom_key_value_data = [
      {
        name  = "tunnelMode"
        value = "split"
      }
    ]

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
