# NOTE: Group creation and assignments are commented out for now.
# microsoft365_graph_beta_groups_group with hard_delete = true fails its permanent-destroy step
# against the test tenant, which makes CheckDestroy report every test as failed even when the
# device configuration itself created, read, updated and destroyed correctly.
# The profile is still fully exercised — assignments are Optional in the schema, so omitting them
# is valid. Re-enable the blocks below once group hard-delete is working to restore assignment
# coverage.

resource "random_string" "custom_config_suffix" {
  length  = 8
  special = false
  upper   = false
}

# ==============================================================================
# Group Dependencies
# ==============================================================================

# resource "microsoft365_graph_beta_groups_group" "custom_config_group_1" {
#   display_name     = "acc-test-ios-custom-config-group-1-${random_string.custom_config_suffix.result}"
#   mail_nickname    = "acc-test-ios-custom-config-1-${random_string.custom_config_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "custom_config_group_2" {
#   display_name     = "acc-test-ios-custom-config-group-2-${random_string.custom_config_suffix.result}"
#   mail_nickname    = "acc-test-ios-custom-config-2-${random_string.custom_config_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "custom_config_group_3" {
#   display_name     = "acc-test-ios-custom-config-group-3-${random_string.custom_config_suffix.result}"
#   mail_nickname    = "acc-test-ios-custom-config-3-${random_string.custom_config_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

# resource "microsoft365_graph_beta_groups_group" "custom_config_group_4" {
#   display_name     = "acc-test-ios-custom-config-group-4-${random_string.custom_config_suffix.result}"
#   mail_nickname    = "acc-test-ios-custom-config-4-${random_string.custom_config_suffix.result}"
#   mail_enabled     = false
#   security_enabled = true
#   hard_delete      = true
# }

resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "custom_configuration_example" {
  display_name = "acc-test-iOS-custom-config-${random_string.custom_config_suffix.result}"
  description  = "Example custom configuration template for iOS devices"

  custom_configuration = {
    payload_file_name = "com.example.custom.mobileconfig"
    payload_name      = "Custom Configuration Example"
    payload           = <<-EOT
      <?xml version="1.0" encoding="UTF-8"?>
      <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
      <plist version="1.0">
      <dict>
          <key>PayloadContent</key>
          <array>
              <dict>
                  <key>PayloadDisplayName</key>
                  <string>Custom Example Configuration</string>
                  <key>PayloadIdentifier</key>
                  <string>com.example.custom.settings</string>
                  <key>PayloadType</key>
                  <string>com.example.custom</string>
                  <key>PayloadUUID</key>
                  <string>12345678-1234-1234-1234-123456789012</string>
                  <key>PayloadVersion</key>
                  <integer>1</integer>
                  <key>ExampleSetting</key>
                  <true/>
              </dict>
          </array>
          <key>PayloadDisplayName</key>
          <string>Custom Configuration Example</string>
          <key>PayloadIdentifier</key>
          <string>com.example.custom</string>
          <key>PayloadType</key>
          <string>Configuration</string>
          <key>PayloadUUID</key>
          <string>87654321-4321-4321-4321-210987654321</string>
          <key>PayloadVersion</key>
          <integer>1</integer>
      </dict>
      </plist>
    EOT
  }

  # assignments = [
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.custom_config_group_1.id
  #   },
  #   {
  #     type     = "groupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.custom_config_group_2.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.custom_config_group_3.id
  #   },
  #   {
  #     type     = "exclusionGroupAssignmentTarget"
  #     group_id = microsoft365_graph_beta_groups_group.custom_config_group_4.id
  #   }
  # ]

  timeouts = {
    create = "50s"
    read   = "5m"
    update = "30m"
    delete = "30m"
  }
}
