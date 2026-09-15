# displayName belongs in the display_name attribute, not inside settings_json. The envelope validator
# must reject it rather than letting two sources of truth diverge.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "rejects_envelope" {
  odata_type   = "#microsoft.graph.iosGeneralDeviceConfiguration"
  display_name = "unit-test-iOS-rejects-envelope"

  settings_json = jsonencode({
    displayName   = "duplicated in the wrong place"
    cameraBlocked = true
  })
}
