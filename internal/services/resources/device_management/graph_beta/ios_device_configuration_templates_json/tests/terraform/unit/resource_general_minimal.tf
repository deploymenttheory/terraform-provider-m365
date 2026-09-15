# A profile declaring a single property. Graph returns ~190 keys on read, so this configuration is
# the projection test: only cameraBlocked may appear in state.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "general_minimal" {
  odata_type   = "#microsoft.graph.iosGeneralDeviceConfiguration"
  display_name = "unit-test-iOS-disable-camera"

  settings_json = jsonencode({
    cameraBlocked = true
  })

  role_scope_tag_ids = ["0"]

  assignments = [
    {
      type        = "groupAssignmentTarget"
      group_id    = "00000000-0000-0000-0000-000000000002"
      filter_id   = "00000000-0000-0000-0000-000000000003"
      filter_type = "include"
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
