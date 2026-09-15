#!/bin/bash

# Example 1: Import Custom Configuration Template
print_status "Example 1: Importing a Custom Configuration Template"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.custom_configuration_example \"12345678-1234-1234-1234-123456789012\""
echo ""

# Example 2: Import Trusted Root Certificate Configuration
print_status "Example 2: Importing a Trusted Root Certificate Configuration"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.trusted_cert_example \"aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\""
echo ""

# Example 3: Import Wi-Fi Configuration
# Note: pre_shared_key is write-only in Microsoft Graph and is not returned on read. After
# importing a Wi-Fi profile, add the key to your configuration; Terraform cannot recover it.
print_status "Example 3: Importing a Wi-Fi Configuration"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.wifi_example \"bbbbbbbb-cccc-dddd-eeee-ffffffffffff\""
echo ""

# Example 4: Import SCEP Certificate Profile
# Note: root_certificate_odata_bind is an @odata.bind navigation reference that Graph does not
# return on read. After importing, expect the expanded URL form rather than the bare GUID; adjust
# your configuration to match or re-apply to normalise it.
print_status "Example 4: Importing a SCEP Certificate Profile"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.scep_cert_example \"ffffffff-eeee-dddd-cccc-bbbbbbbbbbbb\""
echo ""

# Example 5: Import PKCS Certificate Profile
print_status "Example 5: Importing a PKCS Certificate Profile"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.pkcs_cert_example \"11111111-2222-3333-4444-555555555555\""
echo ""

# Example 6: Import Enterprise Wi-Fi Configuration
# Note: neither the @odata.bind certificate references nor password_format_string are returned by
# Graph on read. After importing, re-add them to your configuration.
print_status "Example 6: Importing an Enterprise Wi-Fi Configuration"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.enterprise_wifi_example \"cccccccc-dddd-eeee-ffff-000000000000\""
echo ""

# Example 7: Import EAS Email Profile
print_status "Example 7: Importing an EAS Email Profile"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.eas_email_example \"dddddddd-eeee-ffff-0000-111111111111\""
echo ""

# Example 8: Import VPN Profile
# Note: importing an IKEv2 profile (iosikEv2VpnConfiguration) logs a warning and brings in only the
# shared VPN properties — IKEv2-specific settings are not managed by this resource.
print_status "Example 8: Importing a VPN Profile"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates.vpn_example \"eeeeeeee-ffff-0000-1111-222222222222\""
echo ""
