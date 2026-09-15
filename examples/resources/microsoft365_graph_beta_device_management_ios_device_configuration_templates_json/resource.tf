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
