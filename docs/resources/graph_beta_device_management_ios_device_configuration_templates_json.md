---
page_title: "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json Resource - terraform-provider-microsoft365"
subcategory: "Device Management"

description: |-
  Manages iOS/iPadOS device restriction and device features configuration templates in Microsoft Intune using a raw JSON settings body. Use this resource for iosGeneralDeviceConfiguration (device restrictions) and iosDeviceFeaturesConfiguration (home screen layout, single sign-on, web content filters), whose property surfaces are too large to expose as typed attributes. For certificate, Wi-Fi, VPN, email and custom profiles use microsoft365_graph_beta_device_management_ios_device_configuration_templates instead.
---

# microsoft365_graph_beta_device_management_ios_device_configuration_templates_json (Resource)

Manages iOS/iPadOS device restriction and device features configuration templates in Microsoft Intune using a raw JSON settings body. Use this resource for `iosGeneralDeviceConfiguration` (device restrictions) and `iosDeviceFeaturesConfiguration` (home screen layout, single sign-on, web content filters), whose property surfaces are too large to expose as typed attributes. For certificate, Wi-Fi, VPN, email and custom profiles use `microsoft365_graph_beta_device_management_ios_device_configuration_templates` instead.

## Microsoft Documentation

- [iOS General Device Configuration resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosgeneraldeviceconfiguration?view=graph-rest-beta)
- [iOS Device Features Configuration resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosdevicefeaturesconfiguration?view=graph-rest-beta)
- [iOS Home Screen App resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-ioshomescreenapp?view=graph-rest-beta)
- [iOS Home Screen Folder resource type](https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-ioshomescreenfolder?view=graph-rest-beta)

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
# Example 1: Device restrictions (iosGeneralDeviceConfiguration)
#
# Only the properties declared here are managed. Microsoft Graph returns all ~190 properties of this
# profile type on read, so the provider projects the response onto this shape — anything absent below
# is left alone on the server rather than reset.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "device_restrictions" {
  odata_type   = "#microsoft.graph.iosGeneralDeviceConfiguration"
  display_name = "iOS - Corporate Device Restrictions"
  description  = "Baseline restrictions for corporate-owned iOS devices"

  # Note there is no @odata.type, displayName, description or roleScopeTagIds here: the provider
  # injects those from the attributes above, and including them at the root is an error.
  settings_json = jsonencode({
    # Hardware
    cameraBlocked        = true
    faceTimeBlocked      = true
    screenCaptureBlocked = true

    # App Store
    appStoreBlocked             = false
    appStoreBlockInAppPurchases = true
    appStoreRequirePassword     = true

    # iCloud
    iCloudBlockBackup       = true
    iCloudBlockDocumentSync = true
    iCloudBlockPhotoLibrary = true

    # Passcode
    passcodeRequired                      = true
    passcodeRequiredType                  = "alphanumeric"
    passcodeMinimumLength                 = 6
    passcodeMinutesOfInactivityBeforeLock = 5
    passcodeExpirationDays                = 90

    # Safari
    safariBlockAutofill       = true
    safariRequireFraudWarning = true
  })

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

# Example 2: Home screen layout (iosDeviceFeaturesConfiguration)
#
# Every icons[] entry must declare its own nested @odata.type — Graph uses it to tell an app from a
# folder, and omitting it is an error. Only the root of settings_json is checked for envelope keys, so
# these nested discriminators are left untouched.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "home_screen_layout" {
  odata_type   = "#microsoft.graph.iosDeviceFeaturesConfiguration"
  display_name = "iOS - Supervised Home Screen Layout"
  description  = "Fixed home screen layout for shared iPads"

  settings_json = jsonencode({
    homeScreenDockIcons = [
      {
        "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
        displayName   = "Phone"
        bundleID      = "com.apple.mobilephone"
      },
      {
        "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
        displayName   = "Safari"
        bundleID      = "com.apple.mobilesafari"
      }
    ]

    homeScreenPages = [
      {
        "@odata.type" = "#microsoft.graph.iosHomeScreenPage"
        displayName   = "Work"

        # A top-level page uses `icons` and may contain both apps and folders.
        icons = [
          {
            "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
            displayName   = "Outlook"
            bundleID      = "com.microsoft.Office.Outlook"
          },
          {
            "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
            displayName   = "Teams"
            bundleID      = "com.microsoft.skype.teams"
          },
          {
            "@odata.type" = "#microsoft.graph.iosHomeScreenFolder"
            displayName   = "Office"
            pages = [
              {
                # A folder page is a different type from a top-level page: it uses `apps` rather than
                # `icons`, and cannot nest further folders. Supplying `icons` here fails with
                # "The property 'icons' does not exist on type ...iosHomeScreenFolderPage".
                "@odata.type" = "#microsoft.graph.iosHomeScreenFolderPage"
                displayName   = "Office Page 1"
                apps = [
                  {
                    "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
                    displayName   = "Word"
                    bundleID      = "com.microsoft.Office.Word"
                  },
                  {
                    "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
                    displayName   = "Excel"
                    bundleID      = "com.microsoft.Office.Excel"
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  })

  assignments = [
    {
      type = "allDevicesAssignmentTarget"
    }
  ]
}

# Example 3: Web content filter — automatic filtering (iosDeviceFeaturesConfiguration)
#
# contentFilterSettings is another polymorphic member requiring a nested discriminator. There are two
# implementations, with different property sets:
#
#   iosWebContentFilterAutoFilter            -> allowedUrls, blockedUrls
#   iosWebContentFilterSpecificWebsitesAccess -> websiteList, specificWebsitesOnly
#
# This one uses Apple's automatic adult-content filter, with explicit allow/block overrides.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "web_content_filter_auto" {
  odata_type   = "#microsoft.graph.iosDeviceFeaturesConfiguration"
  display_name = "iOS - Web Content Filter (automatic)"
  description  = "Automatic adult-content filtering with allow and block overrides"

  settings_json = jsonencode({
    contentFilterSettings = {
      "@odata.type" = "#microsoft.graph.iosWebContentFilterAutoFilter"
      allowedUrls = [
        "https://microsoft.com",
        "https://intranet.example.com"
      ]
      blockedUrls = [
        "https://blocked.example.com"
      ]
    }
  })

  assignments = [
    {
      type = "allLicensedUsersAssignmentTarget"
    }
  ]
}

# Example 4: Web content filter — specific websites only (iosDeviceFeaturesConfiguration)
#
# The other contentFilterSettings implementation: Safari is limited to a bookmark list. Note the
# property is `websiteList` and each entry is an iosBookmark (`url`, `displayName`, `bookmarkFolder`).
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "web_content_filter_allowlist" {
  odata_type   = "#microsoft.graph.iosDeviceFeaturesConfiguration"
  display_name = "iOS - Web Content Filter (allowlist)"
  description  = "Restrict Safari to an approved list of sites"

  settings_json = jsonencode({
    contentFilterSettings = {
      "@odata.type" = "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess"
      websiteList = [
        {
          "@odata.type"  = "#microsoft.graph.iosBookmark"
          url            = "https://intranet.example.com"
          displayName    = "Company Intranet"
          bookmarkFolder = "Work"
        }
      ]
    }
  })

  assignments = [
    {
      type = "allLicensedUsersAssignmentTarget"
    }
  ]
}

# Example 5: Lock and home screen wallpaper (iosDeviceFeaturesConfiguration)
#
# wallpaperImage is a mimeContent object: `type` is the MIME type and `value` is the base64-encoded
# image. filebase64() produces exactly what is needed. Requires a supervised device on iOS 8+, and
# only PNG or JPEG are accepted.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "wallpaper" {
  odata_type   = "#microsoft.graph.iosDeviceFeaturesConfiguration"
  display_name = "iOS - Corporate Wallpaper"
  description  = "Sets the lock and home screen wallpaper on supervised devices"

  settings_json = jsonencode({
    # One of: notConfigured, lockScreen, homeScreen, lockAndHomeScreens
    wallpaperDisplayLocation = "lockAndHomeScreens"
    wallpaperImage = {
      type  = "image/jpeg"
      value = filebase64("${path.module}/wallpaper.jpg")
    }
  })

  assignments = [
    {
      type = "allDevicesAssignmentTarget"
    }
  ]
}

# Example 6: AirPrint, notifications and single sign-on (iosDeviceFeaturesConfiguration)
#
# All three of these nested shapes were confirmed against a live tenant.
#
# Note what is NOT in this configuration: Graph also returns nine read-only keys describing the SSO
# extension surface (ssoExtensionKey, allExtensibleSSOEntities, kerberosExtensibleSSOEntities and
# friends). They have no setters in the API, so the provider strips them from responses — do not try
# to declare them.
resource "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json" "device_features_full" {
  odata_type   = "#microsoft.graph.iosDeviceFeaturesConfiguration"
  display_name = "iOS - Device Features"
  description  = "AirPrint destinations, per-app notifications and Entra ID single sign-on"

  settings_json = jsonencode({
    lockScreenFootnote = "Property of Example Corp — if found, call +44 1234 567890"
    assetTagTemplate   = "ASSET-{{SERIALNUMBER}}"

    homeScreenGridWidth  = 4
    homeScreenGridHeight = 5

    # An airPrintDestination is {ipAddress, resourcePath, port, forceTls}.
    airPrintDestinations = [
      {
        ipAddress    = "10.0.0.1"
        resourcePath = "printers/Xerox_1024"
        forceTls     = true
      }
    ]

    # Per-app notification behaviour. alertType is one of deviceDefault, banner, modal, none;
    # previewVisibility one of notConfigured, alwaysShow, hideWhenLocked, neverShow.
    notificationSettings = [
      {
        bundleID                 = "com.microsoft.Office.Outlook"
        appName                  = "Outlook"
        publisher                = "Microsoft Corporation"
        enabled                  = true
        alertType                = "banner"
        previewVisibility        = "hideWhenLocked"
        showInNotificationCenter = true
        showOnLockScreen         = true
        badgesEnabled            = true
        soundsEnabled            = true
      }
    ]

    # Kerberos single sign-on, scoped to named apps and URLs.
    singleSignOnSettings = {
      kerberosPrincipalName = "AADDeviceId"
      kerberosRealm         = "EXAMPLE.COM"
      displayName           = "Example Corp SSO"
      allowedUrls           = ["https://www.example.com"]
      allowedAppsList = [
        {
          appId = "com.microsoft.companyportal"
          name  = "Intune Company Portal"
        }
      ]
    }

    # The Entra ID SSO extension. Note the discriminator: this member is polymorphic, with Kerberos,
    # redirect and Entra ID variants each taking different properties.
    iosSingleSignOnExtension = {
      "@odata.type"             = "#microsoft.graph.iosAzureAdSingleSignOnExtension"
      enableSharedDeviceMode    = true
      bundleIdAccessControlList = ["com.microsoft.Office.Outlook"]
      configurations = [
        {
          "@odata.type" = "#microsoft.graph.keyStringValuePair"
          key           = "device_registration"
          value         = "{{DEVICEREGISTRATION}}"
        }
      ]
    }
  })

  assignments = [
    {
      type = "allDevicesAssignmentTarget"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name for the iOS/iPadOS configuration template.
- `odata_type` (String) The Graph type of the configuration profile. Possible values are: `#microsoft.graph.iosGeneralDeviceConfiguration` (device restrictions), `#microsoft.graph.iosDeviceFeaturesConfiguration` (device features), `#microsoft.graph.iosTrustedRootCertificate` (trusted root certificate), `#microsoft.graph.iosCustomConfiguration` (custom configuration), `#microsoft.graph.iosWiFiConfiguration` (Wi-Fi), `#microsoft.graph.iosScepCertificateProfile` (SCEP certificate), `#microsoft.graph.iosPkcsCertificateProfile` (PKCS certificate), `#microsoft.graph.iosEnterpriseWiFiConfiguration` (enterprise Wi-Fi), `#microsoft.graph.iosEasEmailProfileConfiguration` (EAS email), `#microsoft.graph.iosVpnConfiguration` (VPN). Changing this forces replacement — Graph cannot mutate a profile's type in place.
- `settings_json` (String) The iOS/iPadOS settings tree as a JSON string, typically written with `jsonencode()`.

This must contain **only** the settings themselves. The provider owns the envelope and injects `@odata.type`, `displayName`, `description` and `roleScopeTagIds` from the dedicated attributes; including any of those four at the root of this value is an error.

**Only the properties you declare are managed.** Microsoft Graph returns every property of these profile types on read — including ones never configured, reported as `false`, `null` or `[]` — so the provider projects the response onto the shape of your configuration. A consequence is that `false` and "not set" are indistinguishable on the wire: removing a property from this value stops Terraform tracking it, it does **not** reset it on the server. Declaring a property is how you begin managing it.

**Nested `@odata.type` discriminators are required.** Only the root level is checked for envelope keys. Graph relies on nested discriminators to resolve polymorphic members — for example each `homeScreenPages[].icons[]` entry must declare either `#microsoft.graph.iosHomeScreenApp` or `#microsoft.graph.iosHomeScreenFolder`.

Note the two page types differ: a top-level `iosHomeScreenPage` uses `icons` and may hold both apps and folders, while an `iosHomeScreenFolderPage` nested inside a folder uses `apps` and may hold only apps. Using `icons` on a folder page is rejected by Graph.

A device features profile looks like this:

```hcl
settings_json = jsonencode({
  homeScreenPages = [
    {
      "@odata.type" = "#microsoft.graph.iosHomeScreenPage"
      displayName   = "Page 1"
      icons = [
        {
          "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
          displayName   = "Safari"
          bundleID      = "com.apple.mobilesafari"
        }
      ]
    }
  ]
})
```

On import the full property surface is written to state, which will be large; prune it down to the properties you intend to manage.

### Optional

- `assignments` (Attributes Set) Assignments for the device configuration. Each assignment specifies the target group and schedule for script execution. Supports group filters. (see [below for nested schema](#nestedatt--assignments))
- `description` (String) Optional description of the resource. Maximum length is 1500 characters.
- `role_scope_tag_ids` (Set of String) Set of scope tag IDs for this device configuration template.
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

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


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `read` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).

## Important Notes

### Which resource to use

This resource covers the two iOS/iPadOS template types whose property surface is too large to expose
as typed attributes:

| `odata_type` | Intune profile |
|---|---|
| `#microsoft.graph.iosGeneralDeviceConfiguration` | Device restrictions (187 settable properties) |
| `#microsoft.graph.iosDeviceFeaturesConfiguration` | Device features — home screen layout, single sign-on, web content filters |

For custom `.mobileconfig` payloads, trusted root certificates, SCEP/PKCS certificate profiles,
Wi-Fi (personal and enterprise), Exchange ActiveSync email, and VPN profiles, use the typed sibling
resource
[`microsoft365_graph_beta_device_management_ios_device_configuration_templates`](graph_beta_device_management_ios_device_configuration_templates)
instead.

### The provider owns the envelope

`settings_json` must contain **only** the settings tree. The provider injects `@odata.type`,
`displayName`, `description` and `roleScopeTagIds` from the dedicated attributes. Supplying any of
those four at the **root** of `settings_json` is rejected at plan time, because two sources of truth
for the same property will eventually disagree.

### Only the properties you declare are managed

Microsoft Graph returns **every** property of these profile types on read, including ones that were
never configured — reported as `false`, `null` or `[]`. A profile setting a single property comes back
with roughly 190 keys. The provider therefore projects the response onto the shape of your
configuration, keeping only the properties you declared.

The consequence is worth understanding before you rely on it:

- **`false` and "not set" are indistinguishable on the wire.** Graph reports both as `false`.
- **Removing a property from `settings_json` stops Terraform tracking it. It does not reset it on the
  server.** The value stays whatever it was.
- **Declaring a property is how you begin managing it.** To return a setting to its default, set it
  explicitly rather than deleting it.

Updates use `PATCH`, which is consistent with this model: properties absent from `settings_json` are
left untouched rather than cleared.

### Nested `@odata.type` discriminators are required

Only the **root** of `settings_json` is checked for envelope keys, and only the root is stripped from
responses. Nested `@odata.type` values are load-bearing: Graph uses them to resolve polymorphic
members, and omitting one is an error.

Most commonly this applies to home screen layouts, where every icon must declare whether it is an app
or a folder:

```hcl
settings_json = jsonencode({
  homeScreenPages = [
    {
      "@odata.type" = "#microsoft.graph.iosHomeScreenPage"
      displayName   = "Page 1"
      icons = [
        {
          "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
          displayName   = "Safari"
          bundleID      = "com.apple.mobilesafari"
        },
        {
          "@odata.type" = "#microsoft.graph.iosHomeScreenFolder"
          displayName   = "Utilities"
          pages = [
            {
              # Note: a folder page uses `apps`, not `icons`.
              "@odata.type" = "#microsoft.graph.iosHomeScreenFolderPage"
              displayName   = "Utilities Page 1"
              apps = [
                {
                  "@odata.type" = "#microsoft.graph.iosHomeScreenApp"
                  displayName   = "Calculator"
                  bundleID      = "com.apple.calculator"
                }
              ]
            }
          ]
        }
      ]
    }
  ]
})
```

The two page types are **not** interchangeable, and this is the easiest mistake to make:

| Type | Property | May contain |
|---|---|---|
| `iosHomeScreenPage` (top level) | `icons` | apps **and** folders |
| `iosHomeScreenFolderPage` (inside a folder) | `apps` | apps only |

Using `icons` on a folder page is rejected with `The property 'icons' does not exist on type
...iosHomeScreenFolderPage`.

The same nested-discriminator requirement applies to `singleSignOnExtension`,
`contentFilterSettings` and `notificationSettings[]`.

### Changing `odata_type` replaces the profile

Graph cannot mutate a profile's type in place, so `odata_type` forces replacement.

### Assignment required

Policies must be assigned to device or user groups to be deployed.

## Import

Import is supported using the following syntax:

```shell
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
```

On import there is no prior configuration to project against, so the **entire** stripped response is
written to `settings_json` — expect roughly 190 properties for a device restrictions profile. Prune it
down to the properties you actually intend to manage; anything you remove simply stops being tracked.
