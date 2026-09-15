resource "microsoft365_graph_beta_device_management_windows_trusted_root_certificate" "test" {
  display_name             = "unit-test-windows-trusted-certificate"
  cert_file_name           = "root.cer"
  trusted_root_certificate = "CERTIFICATE_BASE64"
  destination_store        = "computerCertStoreRoot"
}
