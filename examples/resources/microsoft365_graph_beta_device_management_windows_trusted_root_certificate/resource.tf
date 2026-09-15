resource "microsoft365_graph_beta_device_management_windows_trusted_root_certificate" "example" {
  display_name             = "Windows trusted root certificate"
  description              = "Trust the organization's root certificate on Windows devices."
  cert_file_name           = "root.cer"
  trusted_root_certificate = filebase64("root.cer")
  destination_store        = "computerCertStoreRoot"
  role_scope_tag_ids       = ["0"]

  assignments = [{
    type     = "groupAssignmentTarget"
    group_id = "00000000-0000-0000-0000-000000000001" # Replace with your test group ID.
  }]
}
