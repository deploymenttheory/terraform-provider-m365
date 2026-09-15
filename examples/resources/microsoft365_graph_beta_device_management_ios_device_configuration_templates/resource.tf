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
