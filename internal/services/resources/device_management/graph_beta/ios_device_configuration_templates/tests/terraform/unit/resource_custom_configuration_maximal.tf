resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "custom_configuration_example" {
  display_name = "unit-test-iOS-custom-configuration-example"
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
