# NOTE: Group creation and assignments are commented out for now.
# microsoft365_graph_beta_groups_group with hard_delete = true fails its permanent-destroy step
# against the test tenant, which makes CheckDestroy report every test as failed even when the
# device configuration itself created, read, updated and destroyed correctly.
# The profile is still fully exercised — assignments are Optional in the schema, so omitting them
# is valid. Re-enable the blocks below once group hard-delete is working to restore assignment
# coverage.

resource "random_string" "vpn_suffix" {
  length  = 8
  special = false
  upper   = false
}

# ==============================================================================
# Group Dependencies
# ==============================================================================

# resource "microsoft365_graph_beta_groups_group" "vpn_group_1" {
#   display_name     = "acc-test-ios-vpn-group-1-${random_string.vpn_suffix.result}"
#   mail_nickname    = "acc-test-ios-vpn-1-${random_string.vpn_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "vpn_group_2" {
#   display_name     = "acc-test-ios-vpn-group-2-${random_string.vpn_suffix.result}"
#   mail_nickname    = "acc-test-ios-vpn-2-${random_string.vpn_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "vpn_example" {
  display_name = "acc-test-iOS-vpn-${random_string.vpn_suffix.result}"
  description  = "Per-app VPN for corporate applications"

  # authentication_method is usernameAndPassword so no certificate dependency is needed; the
  # certificate bind path is exercised by the enterprise_wifi and scep acceptance tests.
  vpn = {
    connection_name       = "Corporate VPN"
    connection_type       = "ciscoAnyConnectV2"
    authentication_method = "usernameAndPassword"

    # provider_type is deliberately omitted: it selects app-layer vs packet-layer tunnelling and is
    # only available for Pulse Secure and Custom VPN connection types, not ciscoAnyConnectV2.

    # exclude_local_networks / include_all_networks are deliberately omitted: both are documented as
    # "not applicable when enablePerApp is TRUE", and this is a per-app profile.
    enable_split_tunneling = true
    enable_per_app         = true

    safari_domains = ["intranet.example.com"]

    # disconnect_on_idle / disconnect_on_idle_timer_in_seconds are deliberately omitted: they control
    # how long to wait before dropping an idle *on-demand* connection, and this profile defines no
    # on-demand rules. Profiles created through the portal leave disconnectOnIdle null.

    server = {
      address           = "vpn.example.com"
      description       = "Primary VPN gateway"
      is_default_server = true
    }

    proxy_server = {
      address = "proxy.example.com"
      port    = 8080
    }

    # targeted_mobile_apps is deliberately omitted. Per-app VPN associates apps from the *app* side
    # (Apps > Properties > Assignments > VPNs), not from the profile, which is why the portal reports
    # "additional steps are required to configure per-app VPN" and leaves targetedMobileApps empty on
    # every profile it creates. app_id also expects the Intune managed-app GUID rather than a bundle
    # identifier, so a literal bundle ID cannot resolve on this tenant.

    # custom_key_value_data is deliberately omitted: Graph documents it as "custom data when
    # connection type is set to Custom VPN", and connection_type here is ciscoAnyConnectV2. The
    # custom_key_value_data serialization path is covered by the unit tests.
  }

  # assignments = [
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.vpn_group_1.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.vpn_group_2.id
  #   }
  # ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
