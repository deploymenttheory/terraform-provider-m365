# iOS/iPadOS device configuration templates (JSON) — wire traces

Reference request/response bodies for the two `@odata.type` variants handled by this resource. The
GET section is the source of truth for `serverOwnedRootKeys` in `strip.go` — it was derived from an
observed response, not from the SDK surface.

## Create — iosGeneralDeviceConfiguration

```
Request URL:    https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations
Request Method: POST
```

The provider merges the operator's `settings_json` with the envelope properties it owns:

```json
{
  "@odata.type": "#microsoft.graph.iosGeneralDeviceConfiguration",
  "displayName": "iOS - Corporate Device Restrictions",
  "description": "Baseline restrictions for corporate-owned iOS devices",
  "roleScopeTagIds": ["0"],
  "cameraBlocked": true,
  "passcodeRequired": true,
  "passcodeMinimumLength": 6
}
```

`@odata.type`, `displayName`, `description` and `roleScopeTagIds` come from the typed attributes;
everything else is the operator's settings tree passed through byte-for-byte.

Note the body is emitted by `rawJSONRequestBody`, which iterates a Go map — so **top-level key order
is nondeterministic**. Mock responders must decode the body rather than compare a raw string.

## Create — iosDeviceFeaturesConfiguration

```json
{
  "@odata.type": "#microsoft.graph.iosDeviceFeaturesConfiguration",
  "displayName": "iOS - Supervised Home Screen Layout",
  "roleScopeTagIds": ["0"],
  "homeScreenPages": [
    {
      "displayName": "Work",
      "@odata.type": "#microsoft.graph.iosHomeScreenPage",
      "icons": [
        {
          "@odata.type": "#microsoft.graph.iosHomeScreenApp",
          "displayName": "Outlook",
          "bundleID": "com.microsoft.Office.Outlook",
          "isWebClip": false
        },
        {
          "@odata.type": "#microsoft.graph.iosHomeScreenFolder",
          "displayName": "Office",
          "pages": [
            {
              "@odata.type": "#microsoft.graph.iosHomeScreenFolderPage",
              "displayName": "Office Page 1",
              "apps": [
                {
                  "@odata.type": "#microsoft.graph.iosHomeScreenApp",
                  "displayName": "Word",
                  "bundleID": "com.microsoft.Office.Word",
                  "isWebClip": false
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
```

The nested `@odata.type` values are required — Graph uses them to resolve the polymorphic
`iosHomeScreenItem` members. This is why the envelope validator and the response strip are both
**root-only**: recursing would destroy them.

### The two page types are not interchangeable

Confirmed against a live tenant (a `POST` using `icons` on a folder page returns
`400 The property 'icons' does not exist on type ...iosHomeScreenFolderPage`):

| Type | Property | Element type | May contain |
|---|---|---|---|
| `iosHomeScreenPage` | `icons` | `IosHomeScreenItemable` | apps **and** folders |
| `iosHomeScreenFolderPage` | `apps` | `IosHomeScreenAppable` | apps only — no further nesting |

Verified in the SDK: `IosHomeScreenPage` has `SetIcons([]IosHomeScreenItemable)` while
`IosHomeScreenFolderPage` has `SetApps([]IosHomeScreenAppable)`.

Graph also expects an explicit `@odata.type` on the page objects themselves, not only on the icons.

### Other polymorphic members, confirmed against a live tenant

**`contentFilterSettings`** has two implementations with entirely different property sets:

| `@odata.type` | Properties |
|---|---|
| `#microsoft.graph.iosWebContentFilterAutoFilter` | `allowedUrls` (`[]string`), `blockedUrls` (`[]string`) |
| `#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess` | `websiteList`, `specificWebsitesOnly` — both `[]iosBookmark` |

An `iosBookmark` is `{ url, displayName, bookmarkFolder }`. Note it is **not** `{ name,
bookmarkFolderName }` — an earlier draft of the examples invented those names.

Validated AutoFilter body:

```json
"contentFilterSettings": {
  "@odata.type": "#microsoft.graph.iosWebContentFilterAutoFilter",
  "allowedUrls": ["https://microsoft.com", "allowed.com"],
  "blockedUrls": ["blocked.com"]
}
```

**`wallpaperImage`** is a `mimeContent`: `{ type, value }` where `value` is the base64-encoded image
and `type` is its MIME type. `filebase64()` produces the right thing. Supervised devices on iOS 8+
only, PNG or JPEG only. Pairs with `wallpaperDisplayLocation`, one of `notConfigured`, `lockScreen`,
`homeScreen`, `lockAndHomeScreens`.

```json
"wallpaperDisplayLocation": "lockAndHomeScreens",
"wallpaperImage": { "type": "image/jpeg", "value": "/9j/4Q//9k=" }
```

Note the SDK field is `SetTypeEscaped` for the `type` property, since `type` collides with a Go
keyword — irrelevant here because this resource passes raw JSON, but worth knowing if the typed
sibling ever grows a wallpaper attribute.

### iosGeneralDeviceConfiguration nested shapes

A live restrictions profile confirms several properties are objects or object arrays rather than the
scalars they resemble:

| Property | Shape |
|---|---|
| `mediaContentRating<Country>` | object `{movieRating, tvRating}` — nine countries, each with its own rating enums |
| `mediaContentRatingApps` | **scalar** enum: `allAllowed` or `allBlocked`. Easy to conflate with the above |
| `appsSingleAppModeList`, `appsVisibilityList`, `compliantAppsList` | `[]appListItem` — `{appId, name, publisher, appStoreUrl}` |
| `networkUsageRules` | `[]iosNetworkUsageRule` — `{cellularDataBlocked, cellularDataBlockWhenRoaming, managedApps}`. Real profiles send **sparse** entries, each element setting only one property |

`compliantAppListType` and `appsVisibilityListType` are `none`, `appsInListCompliant` or
`appsNotInListCompliant`.

None of the object-array entries carry an `@odata.type`: their collections have concrete element
types, consistent with the discriminator rule above.

**Trap: root-level `movieRating` and `tvRating` are not real.** The Intune UI sends them at the root
of an `iosGeneralDeviceConfiguration` POST, but they do **not** exist on the type — there is no setter
for either anywhere in `ios_general_device_configuration.go`. Graph accepts and discards them. They
belong inside the per-country `mediaContentRating*` objects. Do not copy them out of a captured UI
payload into `settings_json`: they would be projected into state, then never returned, producing
permanent drift.

## Read

```
Request URL:    https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('{id}')
Request Method: GET
```

Three properties of the response drive the whole read path:

### 1. Every property is returned, including ones never set

A profile that configures exactly one property (`cameraBlocked: true`) comes back with roughly **190
keys** — the rest as `false`, `null`, `""` or `[]`. Unset scalars are explicit `null` rather than
omitted. This is why `ProjectOntoShape` is load-bearing rather than a nicety: without it, a three-key
configuration would diff against ~187 unwanted keys on every plan.

It also means **`false` is indistinguishable from unset**. Graph returns `airDropBlocked: false`
whether the operator set it false or never mentioned it. The prior configuration's shape is the only
record of which properties are managed.

### 2. `description` is `null`, not absent

Along with ~25 other unset properties (`passcodeMinimumLength`, all nine `mediaContentRating*`,
`kioskMode*Id`, `softwareUpdatesEnforcedDelayInDays`, and all three
`deviceManagementApplicabilityRule*`). Hydration maps JSON `null` to `types.StringNull()`, not an
empty string.

### 3. There is a sibling `assignments@odata.context` annotation

When assignments are expanded, the response carries **both**:

```json
"assignments@odata.context": "https://graph.microsoft.com/beta/$metadata#deviceManagement/deviceConfigurations('...')/microsoft.graph.iosGeneralDeviceConfiguration/assignments",
"assignments": [ ... ]
```

A fixed-name strip of `@odata.context` alone leaks the annotation straight into `settings_json`. The
strip therefore runs a **pattern pass** first, deleting any root key containing `@odata.` — which also
covers the root `@odata.context`, `@odata.type`, and any future `<nav>@odata.count` or
`@odata.nextLink`.

This resource does not use `$expand=assignments` (assignments are fetched through the typed SDK
separately, so the collection never enters the body that becomes `settings_json`), but the annotation
is stripped defensively regardless.

### 4. Nine server-computed single sign-on metadata keys

`iosDeviceFeaturesConfiguration` responses carry these, none of which has a setter **anywhere** in the
SDK:

```
ssoExtensionKey
allExtensibleSSOEntities            azureAdExtensibleSSOEntities
kerberosExtensibleSSOEntities       nonKerberosExtensibleSSOEntities
redirectExtensibleSSOEntities       nonAzureAdExtensibleSSOEntities
commonCredentialKerberosExtensibleSSOEntities
commonCredentialRedirectExtensibleSSOEntities
```

Graph derives them from the configured `iosSingleSignOnExtension` and returns them as read-only
descriptions of which properties each extension variant supports. They are large — the `all` and
`nonAzureAd` lists run to ~29 entries each — so leaking them into `settings_json` would leave an
operator with a block they cannot delete without appearing to change configuration. All nine are in
`serverOwnedRootKeys`.

The genuinely settable `iosSingleSignOnExtension` and `singleSignOnSettings` are **not** stripped —
only the derived metadata about them.

### Nested shapes confirmed on a live iosDeviceFeaturesConfiguration

| Property | Shape |
|---|---|
| `airPrintDestinations[]` | `{ipAddress, resourcePath, port, forceTls}` — on the `AppleDeviceFeaturesConfigurationBase` parent, not the iOS leaf |
| `notificationSettings[]` | `{bundleID, appName, publisher, enabled, alertType, previewVisibility, showInNotificationCenter, showOnLockScreen, badgesEnabled, soundsEnabled}` |
| `singleSignOnSettings` | `{kerberosPrincipalName, kerberosRealm, displayName, allowedUrls, allowedAppsList[{appId, name}]}` |
| `iosSingleSignOnExtension` | polymorphic — observed `#microsoft.graph.iosAzureAdSingleSignOnExtension` with `{enableSharedDeviceMode, bundleIdAccessControlList, configurations[]}`, each configuration a `#microsoft.graph.keyStringValuePair` `{key, value}` |
| `wallpaperImage` | `mimeContent` `{type, value}` — note Graph echoes a `@odata.type` here even though the property is singular and concrete |

Enum values: `alertType` is one of `deviceDefault`, `banner`, `modal`, `none`; `previewVisibility` one
of `notConfigured`, `alwaysShow`, `hideWhenLocked`, `neverShow`.

Also accepted and echoed, though absent from earlier drafts of these docs: `homeScreenGridWidth`,
`homeScreenGridHeight`, `lockScreenFootnote`, `assetTagTemplate`.

### Confirmed server-owned root key set

```
@odata.context, @odata.type, id, createdDateTime, lastModifiedDateTime, version,
supportsScopeTags, roleScopeTagIds, displayName, description,
deviceManagementApplicabilityRuleOsEdition, deviceManagementApplicabilityRuleOsVersion,
deviceManagementApplicabilityRuleDeviceMode, assignments, assignments@odata.context

plus, on iosDeviceFeaturesConfiguration only:
ssoExtensionKey, allExtensibleSSOEntities, azureAdExtensibleSSOEntities,
commonCredentialKerberosExtensibleSSOEntities, commonCredentialRedirectExtensibleSSOEntities,
kerberosExtensibleSSOEntities, nonAzureAdExtensibleSSOEntities, nonKerberosExtensibleSSOEntities,
redirectExtensibleSSOEntities
```

`isAssigned` does **not** appear, even with `$expand=assignments`. It remains in
`serverOwnedRootKeys` as a defensive delete, but nothing relies on its presence and no test asserts
on it.

### 5. Not every annotation is an `@odata.` one

A live GET carries `@microsoft.graph.tips`, a Microsoft service annotation whose value is a paragraph
of advice about using `$select`:

```json
"@microsoft.graph.tips": "Use $select to choose only the properties your app needs, as this can lead
to performance improvements. For example: GET deviceManagement/deviceConfigurations('<guid>')?$select=..."
```

An `@odata.`-only pattern misses it entirely and writes that prose into `settings_json`, where the
operator can neither remove it (it would look like a configuration change) nor reproduce it. The
pattern pass therefore matches **any** root key containing `@`: in OData, `@` always introduces an
annotation rather than a property, and no real Graph property name contains one. That is a structural
rule rather than a list to maintain.

## Update

```
Request URL:    https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('{id}')
Request Method: PATCH
```

**PATCH confirmed working against a live tenant.** A profile created with `cameraBlocked` only was
updated to also set `airDropBlocked`, and Graph accepted the PATCH. The open question of whether to
fall back to `PutRequestByResourceId` is settled: PATCH is correct and no change is needed.

PATCH rather than PUT, deliberately: PUT would clear properties absent from `settings_json`, which
would contradict the projection semantics on read ("only declared properties are managed").

Note the Intune UI sends its **entire form state** on PATCH — roughly 170 properties, most of them
`false` — rather than only what changed. This provider sends only the declared properties, which is
both smaller and consistent with the projection model. Both are valid PATCH bodies.

## Endpoint path gotcha

The `custom_requests` helpers are inconsistent about the separator between base URL and endpoint:

| Helper | URL template | Endpoint form |
|---|---|---|
| `PostRequest` | `"{+baseurl}/" + Endpoint` | **no** leading slash |
| `GetRequestByResourceId`, `PatchRequestByResourceId` (via `ByIDRequestUrlTemplate`) | `"{+baseurl}" + Endpoint` | leading slash **required** |

Hence the two constants `CollectionEndpointPath` and `ItemEndpointPath`. Using one for both silently
produces `https://graph.microsoft.com/betadeviceManagement/...`.

## Assignments

Read through the typed SDK at `deviceConfigurations/{id}/assignments` — note the plain `{id}` form
here, without the `('{id}')` quoting the raw-JSON item endpoint uses. Assignment targets come back in
the shape `state_assignment.go` already handles, including the `null`
`deviceAndAppManagementAssignmentFilterId` that is normalised to the zero GUID to keep set hashing
stable.
