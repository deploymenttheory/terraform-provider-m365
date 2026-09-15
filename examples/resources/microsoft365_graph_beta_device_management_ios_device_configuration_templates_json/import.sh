#!/bin/bash

# Importing writes the ENTIRE stripped response into settings_json, because there is no prior
# configuration to project against. For a device restrictions profile that is roughly 190 properties,
# most of them Graph's defaults (false, null or []).
#
# After importing, prune settings_json down to the properties you actually intend to manage.
# Removing a property stops Terraform tracking it — it does not reset it on the server.

# Example 1: Import a device restrictions profile (iosGeneralDeviceConfiguration)
print_status "Example 1: Importing a device restrictions profile"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates_json.device_restrictions \"00000000-0000-0000-0000-000000000000\""
echo ""

# Example 2: Import a device features profile (iosDeviceFeaturesConfiguration)
print_status "Example 2: Importing a device features profile"
echo "terraform import microsoft365_graph_beta_device_management_ios_device_configuration_templates_json.home_screen_layout \"11111111-1111-1111-1111-111111111111\""
echo ""
