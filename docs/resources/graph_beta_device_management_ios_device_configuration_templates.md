---
page_title: "microsoft365_graph_beta_device_management_ios_device_configuration_templates Resource - terraform-provider-microsoft365"
subcategory: "Device Management"

description: |-
  Manages iOS/iPadOS configuration templates in Microsoft Intune. This resource creates device configurations for iOS/iPadOS devices including custom configuration profiles, trusted root certificates, certificate profiles (SCEP/PKCS), Wi-Fi profiles (personal and enterprise), Exchange ActiveSync email profiles, and VPN profiles.
---

# microsoft365_graph_beta_device_management_ios_device_configuration_templates (Resource)

Manages iOS/iPadOS configuration templates in Microsoft Intune. This resource creates device configurations for iOS/iPadOS devices including custom configuration profiles, trusted root certificates, certificate profiles (SCEP/PKCS), Wi-Fi profiles (personal and enterprise), Exchange ActiveSync email profiles, and VPN profiles.

## Microsoft Documentation

- [iOS Custom Configuration resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-ioscustomconfiguration?view=graph-rest-beta)
- [iOS Trusted Root Certificate resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iostrustedrootcertificate?view=graph-rest-beta)
- [iOS Wi-Fi Configuration resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-ioswificonfiguration?view=graph-rest-beta)
- [iOS SCEP Certificate Profile resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosscepcertificateprofile?view=graph-rest-beta)
- [iOS PKCS Certificate Profile resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iospkcscertificateprofile?view=graph-rest-beta)
- [iOS Enterprise Wi-Fi Configuration resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosenterprisewificonfiguration?view=graph-rest-beta)
- [iOS EAS Email Profile Configuration resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-ioseasemailprofileconfiguration?view=graph-rest-beta)
- [iOS VPN Configuration resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosvpnconfiguration?view=graph-rest-beta)

## Microsoft Graph API Permissions

The following client `application` permissions are needed in order to use this resource:

**Required:**
- `DeviceManagementConfiguration.Read.All`
- `DeviceManagementConfiguration.ReadWrite.All`

**Optional:**
- `None` `[N/A]`

## Version History

| Version | Status | Notes |
|---------|--------|-------|
| v1.1.0 | Experimental | Initial release |

## Example Usage

```terraform
# Example 1: iOS/iPadOS Custom Configuration Template
#
# Deploys a .mobileconfig payload built with Apple Configurator or Profile Manager. Use this for
# settings Intune does not expose natively. Unlike the macOS equivalent there is no
# deployment_channel — iOS profiles are always delivered over the device channel.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "custom_configuration_example" {
  display_name = "iOS - Custom Configuration Example"
  description  = "Deploys a custom .mobileconfig payload to iOS/iPadOS devices"

  custom_configuration = {
    payload_file_name = "com.example.custom.mobileconfig"
    payload_name      = "Custom Configuration Example"
    payload           = filebase64("${path.module}/com.example.custom.mobileconfig")
  }

  role_scope_tag_ids = ["0"]

  assignments = [
    {
      type     = "groupAssignmentTarget"
      group_id = "00000000-0000-0000-0000-000000000001"
    }
  ]

  timeouts = {
    create = "3m"
    read   = "3m"
    update = "3m"
    delete = "3m"
  }
}

# Example 2: iOS/iPadOS Trusted Root Certificate
#
# Installs a root CA onto the device so it trusts certificates issued from your own PKI. This is
# commonly a prerequisite for enterprise Wi-Fi and VPN profiles that validate a server certificate.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "trusted_cert_example" {
  display_name = "iOS - Corporate Root CA"
  description  = "Installs the corporate root certificate authority"

  trusted_certificate = {
    cert_file_name           = "CorporateRootCA.cer"
    trusted_root_certificate = filebase64("${path.module}/CorporateRootCA.cer")
  }

  role_scope_tag_ids = ["0"]

  assignments = [
    {
      type = "allDevicesAssignmentTarget"
    }
  ]
}

# Example 3: iOS/iPadOS Wi-Fi (pre-shared key)
#
# A WPA2 Personal network. pre_shared_key is write-only in Microsoft Graph: it is never returned on
# read, so Terraform preserves the configured value in state and cannot detect a key that was
# changed directly in Intune. Store it in a variable or secret manager rather than inline.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "wifi_example" {
  display_name = "iOS - Corporate Wi-Fi"
  description  = "Corporate Wi-Fi network with a manual proxy"

  wifi = {
    network_name                        = "Corporate Wi-Fi"
    ssid                                = "CorpWiFi"
    connect_automatically               = true
    connect_when_network_name_is_hidden = false
    wifi_security_type                  = "wpa2Personal"
    pre_shared_key                      = var.corporate_wifi_password

    # Set this when the network authenticates devices by MAC address, since iOS otherwise presents
    # a randomised private address per network.
    disable_mac_address_randomization = false

    proxy_settings       = "manual"
    proxy_manual_address = "proxy.example.com"
    proxy_manual_port    = 8080
  }

  role_scope_tag_ids = ["0"]

  assignments = [
    {
      type        = "groupAssignmentTarget"
      group_id    = "00000000-0000-0000-0000-000000000001"
      filter_id   = "00000000-0000-0000-0000-000000000002"
      filter_type = "include"
    }
  ]
}

# Example 4: iOS/iPadOS Wi-Fi (open network with PAC file)
#
# An open guest network that discovers its proxy from a PAC URL. No pre_shared_key applies.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "wifi_guest_example" {
  display_name = "iOS - Guest Wi-Fi"
  description  = "Open guest network using automatic proxy configuration"

  wifi = {
    network_name                      = "Guest Wi-Fi"
    ssid                              = "GuestWiFi"
    connect_automatically             = false
    wifi_security_type                = "open"
    proxy_settings                    = "automatic"
    proxy_automatic_configuration_url = "https://proxy.example.com/proxy.pac"
  }

  assignments = [
    {
      type = "allLicensedUsersAssignmentTarget"
    }
  ]
}

# Example 5: iOS/iPadOS SCEP Certificate Profile
#
# Issues device certificates via SCEP. Requires an existing trusted root certificate profile, which
# root_certificate_odata_bind references — accepting either a bare GUID (as here) or the full
# https://graph.microsoft.com/beta/... URL. Graph never returns this reference on read, so Terraform
# preserves the configured value; on import it cannot be recovered as a bare GUID.
#
# key_usage and subject_alternative_name_type are bitmasks in Graph (stored as a single
# comma-joined value), so multiple entries in these sets are combined into one wire value.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "scep_example" {
  display_name = "iOS - Device Authentication (SCEP)"
  description  = "SCEP certificate profile for 802.1x device authentication"

  scep_certificate = {
    renewal_threshold_percentage      = 20
    certificate_store                 = "machine"
    certificate_validity_period_scale = "years"
    certificate_validity_period_value = 2

    subject_name_format        = "custom"
    subject_name_format_string = "CN={{AAD_Device_ID}},O=Example Corp,C=US"

    subject_alternative_name_type          = ["emailAddress", "universalResourceIdentifier"]
    subject_alternative_name_format_string = "{{EmailAddress}}"

    # References the trusted_cert_example profile defined above.
    root_certificate_odata_bind = microsoft365_graph_beta_device_management_ios_device_configuration_templates.trusted_cert_example.id

    key_size  = "size2048"
    key_usage = ["keyEncipherment", "digitalSignature"]

    scep_server_urls = [
      "https://scep.example.com/certsrv/mscep/mscep.dll"
    ]

    extended_key_usages = [
      {
        name              = "Client Authentication"
        object_identifier = "1.3.6.1.5.5.7.3.2"
      }
    ]
  }

  assignments = [
    {
      type = "allDevicesAssignmentTarget"
    }
  ]
}

# Example 6: iOS/iPadOS PKCS Certificate Profile
#
# Issues user certificates from an on-premises CA via the Intune Certificate Connector. Note that
# iOS PKCS profiles have no key_size, key_usage, extended_key_usages or scep_server_urls — those
# exist only on scep_certificate.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "pkcs_example" {
  display_name = "iOS - User Authentication (PKCS)"
  description  = "PKCS certificate profile for user authentication"

  pkcs_certificate = {
    renewal_threshold_percentage      = 30
    certificate_store                 = "user"
    certificate_validity_period_scale = "months"
    certificate_validity_period_value = 12

    subject_name_format        = "commonNameIncludingEmail"
    subject_name_format_string = "CN={{UserName}},E={{EmailAddress}},O=Example Corp"

    subject_alternative_name_type          = ["emailAddress"]
    subject_alternative_name_format_string = "{{EmailAddress}}"

    certification_authority      = "ExampleCA.example.com"
    certification_authority_name = "ExampleCA-CA"
    certificate_template_name    = "UserAuthentication"

    custom_subject_alternative_names = [
      {
        san_type = "emailAddress"
        name     = "{{EmailAddress}}"
      }
    ]
  }

  assignments = [
    {
      type = "allLicensedUsersAssignmentTarget"
    }
  ]
}

# Example 7: iOS/iPadOS Enterprise Wi-Fi (802.1x, EAP-TLS)
#
# For 802.1x networks. Note there is no pre_shared_key here — authentication is via EAP using the
# bound client certificate. Use the `wifi` block instead for personal/WEP networks.
#
# root_certificates_for_server_validation_odata_bind is a multi-value @odata.bind: Graph accepts it
# as a JSON array, and each entry may be a bare GUID or the full URL.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "enterprise_wifi_example" {
  display_name = "iOS - Corporate 802.1x"
  description  = "Enterprise Wi-Fi using EAP-TLS certificate authentication"

  enterprise_wifi = {
    network_name                        = "Corporate 802.1x"
    ssid                                = "CorpSecure"
    connect_automatically               = true
    connect_when_network_name_is_hidden = false
    wifi_security_type                  = "wpa2Enterprise"
    proxy_settings                      = "none"

    eap_type              = "eapTls"
    authentication_method = "certificate"

    # Sent in the clear during the outer EAP exchange, hiding the real identity.
    outer_identity_privacy_temporary_value = "anonymous"
    username_format_string                 = "{{UserPrincipalName}}"

    trusted_server_certificate_names = ["radius.example.com"]

    root_certificates_for_server_validation_odata_bind = [
      microsoft365_graph_beta_device_management_ios_device_configuration_templates.trusted_cert_example.id
    ]

    identity_certificate_for_client_authentication_odata_bind = microsoft365_graph_beta_device_management_ios_device_configuration_templates.scep_example.id
  }

  assignments = [
    {
      type = "allDevicesAssignmentTarget"
    }
  ]
}

# Example 8: iOS/iPadOS Exchange ActiveSync Email Profile
#
# eas_services is a bitmask in Graph (stored as a single comma-joined value), so the set entries are
# combined into one wire value.
#
# Note username_source and username_aad_source are different Graph enums: only the latter accepts
# samAccountName.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "eas_email_example" {
  display_name = "iOS - Corporate Email"
  description  = "Exchange ActiveSync profile for Exchange Online"

  eas_email = {
    account_name          = "Corporate Email"
    host_name             = "outlook.office365.com"
    authentication_method = "certificate"

    eas_services                       = ["calendars", "contacts", "email"]
    eas_services_user_override_enabled = false

    duration_of_email_to_sync = "oneMonth"
    email_address_source      = "primarySmtpAddress"
    username_source           = "userPrincipalName"
    user_domain_name_source   = "fullDomainName"

    require_ssl = true

    block_moving_messages_to_other_email_accounts = true
    block_sending_email_from_third_party_apps     = true

    identity_certificate_odata_bind = microsoft365_graph_beta_device_management_ios_device_configuration_templates.pkcs_example.id
  }

  assignments = [
    {
      type = "allLicensedUsersAssignmentTarget"
    }
  ]
}

# Example 9: iOS/iPadOS per-app VPN
#
# Note `ikEv2` is not an accepted connection_type: IKEv2 uses a separate Graph type with ~23
# additional properties this resource does not model, so the validator rejects it rather than
# creating a profile that would round-trip lossily. On-demand rules are likewise not yet modelled.
#
# custom_data and custom_key_value_data are two distinct Graph properties: entries in the former use
# a `key` field, entries in the latter use `name`.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates" "vpn_example" {
  display_name = "iOS - Corporate VPN"
  description  = "Per-app VPN tunnelling traffic for selected corporate apps"

  vpn = {
    connection_name       = "Corporate VPN"
    connection_type       = "ciscoAnyConnectV2"
    authentication_method = "certificate"
    provider_type         = "packetTunnel"

    enable_split_tunneling = true
    enable_per_app         = true
    exclude_local_networks = true

    # Visiting these in Safari brings the per-app VPN up.
    safari_domains = ["intranet.example.com", "portal.example.com"]

    disconnect_on_idle                  = true
    disconnect_on_idle_timer_in_seconds = 300

    server = {
      address           = "vpn.example.com"
      description       = "Primary VPN gateway"
      is_default_server = true
    }

    proxy_server = {
      address = "proxy.example.com"
      port    = 8080
    }

    # Only these apps' traffic is tunnelled, since enable_per_app is true.
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

    custom_key_value_data = [
      {
        name  = "tunnelMode"
        value = "split"
      }
    ]

    identity_certificate_odata_bind = microsoft365_graph_beta_device_management_ios_device_configuration_templates.scep_example.id
  }

  assignments = [
    {
      type = "allLicensedUsersAssignmentTarget"
    }
  ]
}

variable "corporate_wifi_password" {
  description = "Pre-shared key for the corporate Wi-Fi network."
  type        = string
  sensitive   = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name for the iOS/iPadOS configuration template.

### Optional

- `assignments` (Attributes Set) Assignments for the device configuration. Each assignment specifies the target group and schedule for script execution. Supports group filters. (see [below for nested schema](#nestedatt--assignments))
- `custom_configuration` (Attributes) The custom configuration template allows IT admins to assign settings that aren't built into Intune yet. For iOS/iPadOS devices, you can import a .mobileconfig file that you created using Apple Configurator or Profile Manager. (see [below for nested schema](#nestedatt--custom_configuration))
- `description` (String) Optional description of the resource. Maximum length is 1500 characters.
- `eas_email` (Attributes) Exchange ActiveSync email profile configuration for iOS/iPadOS devices. (see [below for nested schema](#nestedatt--eas_email))
- `enterprise_wifi` (Attributes) Enterprise (802.1x) Wi-Fi configuration for iOS/iPadOS devices. Authentication is via EAP, so there is no pre-shared key — use the `wifi` block for personal/WEP networks instead. (see [below for nested schema](#nestedatt--enterprise_wifi))
- `general_device_configuration` (Attributes) General iOS/iPadOS device restriction policy (`iosGeneralDeviceConfiguration`). Set this block (with no inner attributes required) to create a general device restrictions profile. Configure the full property set via the JSON resource or Intune portal after creation. (see [below for nested schema](#nestedatt--general_device_configuration))
- `pkcs_certificate` (Attributes) PKCS certificate profile configuration for iOS/iPadOS devices. Unlike the macOS equivalent, iOS PKCS profiles have no key size, key usage or extended key usage properties. (see [below for nested schema](#nestedatt--pkcs_certificate))
- `role_scope_tag_ids` (Set of String) Set of scope tag IDs for this device configuration template.
- `scep_certificate` (Attributes) SCEP certificate profile configuration for iOS/iPadOS devices. Requires a trusted root certificate profile to already exist, referenced via root_certificate_odata_bind. (see [below for nested schema](#nestedatt--scep_certificate))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `trusted_certificate` (Attributes) Trusted root certificate configuration for iOS/iPadOS devices. (see [below for nested schema](#nestedatt--trusted_certificate))
- `vpn` (Attributes) VPN configuration for iOS/iPadOS devices. Note that IKEv2 profiles use a distinct Graph type with a substantially different property set and are not supported by this block — `ikEv2` is therefore not an accepted connection_type. On-demand rules are also not yet modelled. (see [below for nested schema](#nestedatt--vpn))
- `wifi` (Attributes) Wi-Fi configuration for iOS/iPadOS devices. For enterprise (802.1x) networks use the enterprise Wi-Fi profile instead. (see [below for nested schema](#nestedatt--wifi))

### Read-Only

- `id` (String) The unique identifier for the iOS/iPadOS configuration template.

<a id="nestedatt--assignments"></a>
### Nested Schema for `assignments`

Required:

- `type` (String) Type of assignment target. Must be one of: 'allDevicesAssignmentTarget', 'allLicensedUsersAssignmentTarget', 'groupAssignmentTarget', 'exclusionGroupAssignmentTarget'.

Optional:

- `filter_id` (String) ID of the filter to apply to the assignment. Required when filter_type is 'include' or 'exclude'. Should be omitted when filter_type is 'none'.
- `filter_type` (String) Type of filter to apply. Must be one of: 'include', 'exclude', or 'none'.
- `group_id` (String) The Entra ID group ID to include or exclude in the assignment. Required when type is 'groupAssignmentTarget' or 'exclusionGroupAssignmentTarget'.


<a id="nestedatt--custom_configuration"></a>
### Nested Schema for `custom_configuration`

Required:

- `payload` (String) The iOS/iPadOS configuration payload (.mobileconfig / .plist) file content.
- `payload_file_name` (String) The profile name displayed to users.
- `payload_name` (String) The name of the payload configuration.


<a id="nestedatt--eas_email"></a>
### Nested Schema for `eas_email`

Required:

- `authentication_method` (String) How the device authenticates to Exchange. Possible values are: usernameAndPassword, certificate, derivedCredential.
- `host_name` (String) The Exchange server hostname that the device connects to.

Optional:

- `account_name` (String) The display name of the email account as shown to users on the device.
- `block_moving_messages_to_other_email_accounts` (Boolean) Whether to prevent users moving messages out of this account into another.
- `block_sending_email_from_third_party_apps` (Boolean) Whether to prevent third-party apps sending email from this account.
- `block_syncing_recently_used_email_addresses` (Boolean) Whether to prevent syncing of recently used email addresses.
- `custom_domain_name` (String) A custom domain name value used instead of the one derived from Entra ID.
- `duration_of_email_to_sync` (String) How much email history to synchronise. Possible values are: userDefined, oneDay, threeDays, oneWeek, twoWeeks, oneMonth, unlimited.
- `eas_services` (Set of String) The Exchange data types to synchronise. Graph models this as a bitmask, so multiple values may be combined. Possible values are: none, calendars, contacts, email, notes, reminders.
- `eas_services_user_override_enabled` (Boolean) Whether users may change which Exchange data types are synchronised.
- `email_address_source` (String) Which Entra ID attribute supplies the email address. Possible values are: userPrincipalName, primarySmtpAddress.
- `encryption_certificate_type` (String) The type of certificate used for S/MIME encryption. Possible values are: none, certificate, derivedCredential.
- `identity_certificate_odata_bind` (String) Reference to a pre-existing iOS SCEP or PKCS certificate profile used to authenticate to Exchange. Accepts a bare GUID or the full URL. Required when authentication_method is certificate. Graph does not return this reference on read.
- `per_app_vpn_profile_id` (String) The identifier of a per-app VPN profile to associate with this email profile.
- `require_smime` (Boolean) Whether S/MIME is required for outgoing messages.
- `require_ssl` (Boolean) Whether SSL is required for connections to the Exchange server.
- `signing_certificate_type` (String) The type of certificate used for S/MIME signing. Possible values are: none, certificate, derivedCredential.
- `smime_enable_per_message_switch` (Boolean) Whether users may toggle S/MIME on a per-message basis.
- `smime_encrypt_by_default_enabled` (Boolean) Whether outgoing messages are S/MIME encrypted by default.
- `smime_encrypt_by_default_user_override_enabled` (Boolean) Whether users may change the encrypt-by-default setting.
- `smime_encryption_certificate_odata_bind` (String) Reference to a pre-existing iOS certificate profile used for S/MIME encryption. Accepts a bare GUID or the full URL. Graph does not return this reference on read.
- `smime_encryption_certificate_user_override_enabled` (Boolean) Whether users may select the S/MIME encryption certificate.
- `smime_signing_certificate_odata_bind` (String) Reference to a pre-existing iOS certificate profile used for S/MIME signing. Accepts a bare GUID or the full URL. Graph does not return this reference on read.
- `smime_signing_certificate_user_override_enabled` (Boolean) Whether users may select the S/MIME signing certificate.
- `smime_signing_enabled` (Boolean) Whether S/MIME signing is enabled for this account.
- `smime_signing_user_override_enabled` (Boolean) Whether users may change the S/MIME signing setting.
- `use_oauth` (Boolean) Whether the connection uses OAuth for authentication.
- `user_domain_name_source` (String) Which form of the domain name to use. Possible values are: fullDomainName, netBiosDomainName.
- `username_aad_source` (String) Which Entra ID attribute supplies the username. Note this uses a wider set of values than username_source. Possible values are: userPrincipalName, primarySmtpAddress, samAccountName.
- `username_source` (String) Which Entra ID attribute supplies the username. Possible values are: userPrincipalName, primarySmtpAddress.


<a id="nestedatt--enterprise_wifi"></a>
### Nested Schema for `enterprise_wifi`

Required:

- `network_name` (String) The name of the Wi-Fi network shown to users when browsing available networks.
- `ssid` (String) The service set identifier (SSID) of the Wi-Fi network the device connects to.

Optional:

- `authentication_method` (String) How the device authenticates to the network. Possible values are: certificate, usernameAndPassword, derivedCredential.
- `connect_automatically` (Boolean) Whether the device connects automatically to this Wi-Fi network when in range.
- `connect_when_network_name_is_hidden` (Boolean) Whether the device connects to the network even when the SSID is not broadcast.
- `disable_mac_address_randomization` (Boolean) Whether to disable the device's private (randomized) Wi-Fi MAC address for this network.
- `eap_fast_configuration` (String) Protected Access Credential (PAC) handling, applicable when eap_type is eapFast. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
- `eap_type` (String) The extensible authentication protocol (EAP) type used for 802.1x authentication. Possible values are: eapTls, leap, eapSim, eapTtls, peap, eapFast, teap.
- `identity_certificate_for_client_authentication_odata_bind` (String) Reference to a pre-existing iOS SCEP or PKCS certificate profile used for client authentication. Accepts a bare GUID or the full URL. Required when authentication_method is certificate. Graph does not return this reference on read.
- `inner_authentication_protocol_for_eap_ttls` (String) The non-EAP inner authentication protocol, applicable when eap_type is eapTtls. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
- `outer_identity_privacy_temporary_value` (String) The identity sent in the clear during the outer EAP exchange, hiding the real user identity until the tunnel is established.
- `password_format_string` (String, Sensitive) The password format used for authentication. Graph does not return this value on read, so it is preserved from configuration.
- `proxy_automatic_configuration_url` (String) The URL of the proxy auto-configuration (PAC) file. Applies when proxy_settings is automatic.
- `proxy_manual_address` (String) The IP address or hostname of the proxy server. Applies when proxy_settings is manual.
- `proxy_manual_port` (Number) The port of the proxy server. Applies when proxy_settings is manual.
- `proxy_settings` (String) How the device obtains its proxy configuration. Possible values are: none, manual, automatic.
- `root_certificates_for_server_validation_odata_bind` (Set of String) References to pre-existing iOS trusted root certificate profiles used to validate the RADIUS server. Each entry accepts a bare GUID or the full "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('...')" URL. Graph does not return these references on read, so they are preserved from configuration.
- `trusted_server_certificate_names` (Set of String) The common names of the certificates the device should trust for server validation.
- `username_format_string` (String) The username format used for authentication. Example: {{UserPrincipalName}}
- `wifi_security_type` (String) The Wi-Fi security protocol. For enterprise networks use wpaEnterprise or wpa2Enterprise. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise, wpa3Personal.


<a id="nestedatt--general_device_configuration"></a>
### Nested Schema for `general_device_configuration`


<a id="nestedatt--pkcs_certificate"></a>
### Nested Schema for `pkcs_certificate`

Required:

- `subject_name_format` (String) How Intune builds the subject name in the certificate request. Possible values are: commonName, commonNameAsEmail, custom, commonNameIncludingEmail, commonNameAsIMEI, commonNameAsSerialNumber.

Optional:

- `certificate_store` (String) The certificate store location. Possible values are: user, machine.
- `certificate_template_name` (String) The name of the certificate template on the issuing certification authority.
- `certificate_validity_period_scale` (String) The unit for the certificate validity period. Possible values are: days, months, years.
- `certificate_validity_period_value` (Number) The certificate validity period, in the unit given by certificate_validity_period_scale.
- `certification_authority` (String) The fully qualified domain name of the issuing certification authority.
- `certification_authority_name` (String) The name of the issuing certification authority.
- `custom_subject_alternative_names` (Attributes Set) Custom Subject Alternative Names for the certificate. (see [below for nested schema](#nestedatt--pkcs_certificate--custom_subject_alternative_names))
- `renewal_threshold_percentage` (Number) The percentage of the certificate lifetime remaining when renewal is attempted (1-99).
- `subject_alternative_name_format_string` (String) The custom subject alternative name format string.
- `subject_alternative_name_type` (Set of String) The subject alternative name types to include. Graph models this as a bitmask, so multiple values may be combined. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
- `subject_name_format_string` (String) The custom subject name format, used when subject_name_format is custom. Example: CN={{UserName}},E={{EmailAddress}},O=Example Corp

<a id="nestedatt--pkcs_certificate--custom_subject_alternative_names"></a>
### Nested Schema for `pkcs_certificate.custom_subject_alternative_names`

Required:

- `name` (String) The SAN value/name.
- `san_type` (String) The SAN type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.



<a id="nestedatt--scep_certificate"></a>
### Nested Schema for `scep_certificate`

Required:

- `certificate_validity_period_scale` (String) The unit for the certificate validity period. Possible values are: days, months, years.
- `certificate_validity_period_value` (Number) The certificate validity period, in the unit given by certificate_validity_period_scale.
- `extended_key_usages` (Attributes Set) Extended key usage settings for the certificate. (see [below for nested schema](#nestedatt--scep_certificate--extended_key_usages))
- `key_size` (String) The key size in bits. 2048 is the recommended minimum. Possible values are: size1024, size2048, size4096.
- `key_usage` (Set of String) Key usage options for the certificate. Graph models this as a bitmask, so both values may be combined. Possible values are: keyEncipherment, digitalSignature.
- `renewal_threshold_percentage` (Number) The percentage of the certificate lifetime remaining when renewal is attempted (1-99).
- `root_certificate_odata_bind` (String) Reference to a pre-existing iOS trusted root certificate profile. Accepts either a bare GUID (e.g. '00000000-0000-0000-0000-000000000000') or the full URL "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('00000000-0000-0000-0000-000000000000')". Graph does not return this reference on read, so it is preserved from configuration.
- `subject_name_format` (String) How Intune builds the subject name in the certificate request. Possible values are: commonName, commonNameAsEmail, custom, commonNameIncludingEmail, commonNameAsIMEI, commonNameAsSerialNumber. Use custom together with subject_name_format_string. See https://learn.microsoft.com/en-us/intune/intune-service/protect/certificates-profile-scep

Optional:

- `certificate_store` (String) The certificate store location. Possible values are: user, machine.
- `custom_subject_alternative_names` (Attributes Set) Custom Subject Alternative Names for the certificate. (see [below for nested schema](#nestedatt--scep_certificate--custom_subject_alternative_names))
- `scep_server_urls` (Set of String) SCEP server URL(s) for certificate enrollment.
- `subject_alternative_name_format_string` (String) The custom subject alternative name format string.
- `subject_alternative_name_type` (Set of String) The subject alternative name types to include. Graph models this as a bitmask, so multiple values may be combined. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
- `subject_name_format_string` (String) The custom subject name format, used when subject_name_format is custom. Example: CN={{AAD_Device_ID}},O={{Organization}}

<a id="nestedatt--scep_certificate--extended_key_usages"></a>
### Nested Schema for `scep_certificate.extended_key_usages`

Required:

- `name` (String) The extended key usage name.
- `object_identifier` (String) The extended key usage object identifier (OID).


<a id="nestedatt--scep_certificate--custom_subject_alternative_names"></a>
### Nested Schema for `scep_certificate.custom_subject_alternative_names`

Required:

- `name` (String) The SAN value/name.
- `san_type` (String) The SAN type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `read` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).


<a id="nestedatt--trusted_certificate"></a>
### Nested Schema for `trusted_certificate`

Required:

- `cert_file_name` (String) The file name of the certificate file (.cer file).
- `trusted_root_certificate` (String) The base64-encoded trusted root certificate content. This should be a filebase64() encoded string. e.g filebase64("my-root-cert.cer")


<a id="nestedatt--vpn"></a>
### Nested Schema for `vpn`

Required:

- `connection_name` (String) The connection name displayed to users on the device.
- `connection_type` (String) The VPN vendor/provider. Possible values are: ciscoAnyConnect, pulseSecure, f5EdgeClient, dellSonicWallMobileConnect, checkPointCapsuleVpn, customVpn, ciscoIPSec, citrix, ciscoAnyConnectV2, paloAltoGlobalProtect, zscalerPrivateAccess, f5Access2018, citrixSso, paloAltoGlobalProtectV2, alwaysOn, microsoftTunnel, netMotionMobility, microsoftProtect. `ikEv2` is intentionally excluded: it maps to a separate Graph type this resource does not model.

Optional:

- `associated_domains` (Set of String) Domains associated with this VPN profile.
- `authentication_method` (String) How the device authenticates to the VPN. Possible values are: certificate, usernameAndPassword, sharedSecret, derivedCredential, azureAD.
- `cloud_name` (String) The Zscaler cloud name. Applies when connection_type is zscalerPrivateAccess.
- `custom_data` (Attributes Set) Vendor-specific key/value configuration. Note this maps to Graph's `customData` property, whose entries use a `key` field — distinct from custom_key_value_data below. (see [below for nested schema](#nestedatt--vpn--custom_data))
- `custom_key_value_data` (Attributes Set) Vendor-specific name/value configuration. Note this maps to Graph's `customKeyValueData` property, whose entries use a `name` field — distinct from custom_data above. (see [below for nested schema](#nestedatt--vpn--custom_key_value_data))
- `disable_on_demand_user_override` (Boolean) Whether users are prevented from disabling on-demand VPN.
- `disconnect_on_idle` (Boolean) Whether the VPN disconnects after an idle period.
- `disconnect_on_idle_timer_in_seconds` (Number) How long the connection may be idle before disconnecting, in seconds.
- `enable_per_app` (Boolean) Whether this is a per-app VPN, tunnelling only traffic from the apps listed in targeted_mobile_apps.
- `enable_split_tunneling` (Boolean) Whether only traffic destined for the VPN's networks is tunnelled.
- `exclude_list` (Set of String) Hostnames excluded from the Zscaler tunnel. Applies when connection_type is zscalerPrivateAccess.
- `exclude_local_networks` (Boolean) Whether traffic to local networks bypasses the VPN.
- `excluded_domains` (Set of String) Domains whose traffic bypasses the VPN even while it is connected.
- `identifier` (String) The vendor-supplied bundle identifier of the VPN app. Required when connection_type is customVpn.
- `identity_certificate_odata_bind` (String) Reference to a pre-existing iOS SCEP or PKCS certificate profile used to authenticate to the VPN. Accepts a bare GUID or the full URL. Required when authentication_method is certificate. Graph does not return this reference on read.
- `include_all_networks` (Boolean) Whether all network traffic is routed through the VPN.
- `login_group_or_domain` (String) The login group or domain, used by some VPN vendors.
- `microsoft_tunnel_site_id` (String) The Microsoft Tunnel site identifier. Applies when connection_type is microsoftTunnel.
- `opt_in_to_device_id_sharing` (Boolean) Whether the device identifier is shared with the VPN provider.
- `provider_type` (String) The tunnel provider type. Possible values are: notConfigured, appProxy, packetTunnel.
- `proxy_server` (Attributes) The proxy server used while the VPN is connected. (see [below for nested schema](#nestedatt--vpn--proxy_server))
- `realm` (String) The authentication realm, used by some VPN vendors.
- `role` (String) The authentication role, used by some VPN vendors.
- `safari_domains` (Set of String) Domains that trigger the per-app VPN when visited in Safari. Applies when enable_per_app is true.
- `server` (Attributes) The VPN server the device connects to. (see [below for nested schema](#nestedatt--vpn--server))
- `strict_enforcement` (Boolean) Whether the VPN stays connected and blocks traffic when it cannot be established. Applies to Microsoft Tunnel.
- `targeted_mobile_apps` (Attributes Set) The apps whose traffic is tunnelled by this per-app VPN. Applies when enable_per_app is true. (see [below for nested schema](#nestedatt--vpn--targeted_mobile_apps))
- `user_domain` (String) The Zscaler user domain. Applies when connection_type is zscalerPrivateAccess.

<a id="nestedatt--vpn--custom_data"></a>
### Nested Schema for `vpn.custom_data`

Required:

- `key` (String) The configuration key.
- `value` (String) The configuration value.


<a id="nestedatt--vpn--custom_key_value_data"></a>
### Nested Schema for `vpn.custom_key_value_data`

Required:

- `name` (String) The configuration name.
- `value` (String) The configuration value.


<a id="nestedatt--vpn--proxy_server"></a>
### Nested Schema for `vpn.proxy_server`

Optional:

- `address` (String) The IP address or hostname of the proxy server.
- `automatic_configuration_script_url` (String) The URL of the proxy auto-configuration (PAC) script.
- `port` (Number) The port of the proxy server.


<a id="nestedatt--vpn--server"></a>
### Nested Schema for `vpn.server`

Required:

- `address` (String) The IP address or fully qualified domain name of the VPN server.

Optional:

- `description` (String) A description of the VPN server.
- `is_default_server` (Boolean) Whether this is the default server for the connection.


<a id="nestedatt--vpn--targeted_mobile_apps"></a>
### Nested Schema for `vpn.targeted_mobile_apps`

Required:

- `app_id` (String) The bundle identifier of the app. e.g. com.microsoft.Office.Outlook

Optional:

- `app_store_url` (String) The App Store URL of the app.
- `name` (String) The display name of the app.
- `publisher` (String) The publisher of the app.



<a id="nestedatt--wifi"></a>
### Nested Schema for `wifi`

Required:

- `network_name` (String) The name of the Wi-Fi network shown to users when browsing available networks on the device.
- `ssid` (String) The service set identifier (SSID) of the Wi-Fi network the device connects to.

Optional:

- `connect_automatically` (Boolean) Whether the device connects automatically to this Wi-Fi network when in range.
- `connect_when_network_name_is_hidden` (Boolean) Whether the device connects to the network even when the SSID is not broadcast.
- `disable_mac_address_randomization` (Boolean) Whether to disable the device's private (randomized) Wi-Fi MAC address for this network. Set this when the network relies on MAC-based authentication.
- `pre_shared_key` (String, Sensitive) The pre-shared key (password) for the Wi-Fi network. Applies to the personal and WEP security types. Graph does not return this value on read, so it is preserved from configuration.
- `proxy_automatic_configuration_url` (String) The URL of the proxy auto-configuration (PAC) file. Applies when proxy_settings is automatic.
- `proxy_manual_address` (String) The IP address or hostname of the proxy server. Applies when proxy_settings is manual.
- `proxy_manual_port` (Number) The port of the proxy server. Applies when proxy_settings is manual.
- `proxy_settings` (String) How the device obtains its proxy configuration for this network. Possible values are: none, manual, automatic.
- `wifi_security_type` (String) The Wi-Fi security protocol. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, wpa2Enterprise, wpa3Personal. For the personal and WEP variants, set pre_shared_key.

## Important Notes

- **One configuration per resource**: Exactly one of `custom_configuration`, `trusted_certificate`,
  `wifi`, `scep_certificate`, `pkcs_certificate`, `enterprise_wifi`, `eas_email`, or `vpn` must be
  set. Each maps to a distinct Graph `@odata.type`, and the type of an existing profile cannot be
  changed in place — switching blocks recreates the profile.
- **IKEv2 VPN profiles are not supported**: `ikEv2` is deliberately excluded from `connection_type`.
  IKEv2 uses a separate Graph type (`iosikEv2VpnConfiguration`) with roughly 23 additional
  properties — including nested security association parameters — that this resource does not model.
  Creating one here would round-trip lossily, so the validator rejects it. An existing IKEv2 profile
  imported into this resource logs a warning and exposes only the shared VPN properties.
- **On-demand VPN rules are not yet modelled**: the `onDemandRules` property is not exposed. Rules
  configured outside Terraform are left untouched.
- **Personal vs enterprise Wi-Fi**: use `wifi` for open/WEP/WPA-Personal networks (which take a
  `pre_shared_key`) and `enterprise_wifi` for 802.1x networks (which authenticate via EAP). The
  enterprise block intentionally has no `pre_shared_key`.
- **No deployment channel on iOS**: Unlike the macOS equivalent, iOS profiles have no
  `deployment_channel` property. iOS profiles are always delivered over the device channel.
- **iOS PKCS is narrower than macOS PKCS**: iOS PKCS certificate profiles have no key size, key
  usage, extended key usage or SCEP server URL properties. Those apply only to `scep_certificate`.
- **`key_usage`, `subject_alternative_name_type` and `eas_services` are bitmasks**: Graph stores each
  as a single comma-joined value rather than a list. They are modelled here as sets and combined on
  write / decomposed on read, so multiple values may be supplied.
- **Boolean attributes are computed**: Graph returns every boolean property, using `false` for ones
  that were never set. Every optional boolean in this resource is therefore also computed — omitting
  one means "let the server decide" rather than "set it to false", and its value will appear in state
  after apply.
- **`eas_email` has two distinct username source attributes**: `username_source` and
  `email_address_source` accept `userPrincipalName` / `primarySmtpAddress`, while
  `username_aad_source` additionally accepts `samAccountName`. They map to different Graph enums and
  are not interchangeable.
- **`root_certificate_odata_bind` is not returned by Graph**: The SCEP root certificate reference is
  a navigation property sent as an `@odata.bind` annotation. Graph accepts it on write but does not
  echo it back, so the provider preserves the configured value in state. Accepts either a bare GUID
  or the full `https://graph.microsoft.com/beta/...` URL. On import the bare-GUID form cannot be
  recovered; the expanded URL is used instead.
- **`pre_shared_key` is never returned by Graph**: The Wi-Fi pre-shared key is write-only. It is
  preserved in state from configuration rather than read back, so a key changed directly in Intune
  will not be detected as drift. On import the value is unknown and must be supplied in
  configuration.
- **Assignment Required**: Policies must be assigned to device or user groups to be deployed.

## Import

Import is supported using the following syntax:

```shell
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
```
