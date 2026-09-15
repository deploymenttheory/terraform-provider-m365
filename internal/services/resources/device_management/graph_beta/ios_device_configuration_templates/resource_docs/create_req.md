# iOS/iPadOS device configuration template — wire traces

Reference request/response bodies for the three `@odata.type` variants handled by this resource.
Note that iOS has **no `deploymentChannel`** on custom configuration or trusted root certificate
profiles, which is the main divergence from the macOS equivalent.

## iosCustomConfiguration

```
Request URL:    https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations
Request Method: POST
```

Request body (payload truncated):

```json
{
  "displayName": "iOS - Custom Configuration Example",
  "description": "Deploys a custom .mobileconfig payload",
  "roleScopeTagIds": ["0"],
  "@odata.type": "#microsoft.graph.iosCustomConfiguration",
  "payloadFileName": "com.example.custom.mobileconfig",
  "payloadName": "custom-config-profile-name",
  "payload": "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4..."
}
```

Response `201`:

```json
{
  "@odata.context": "https://graph.microsoft.com/beta/$metadata#deviceManagement/deviceConfigurations/$entity",
  "@odata.type": "#microsoft.graph.iosCustomConfiguration",
  "id": "f1bde6e1-5b28-4abf-9f82-07bda322131d",
  "lastModifiedDateTime": "2026-03-12T08:19:05.4757913Z",
  "roleScopeTagIds": ["0"],
  "supportsScopeTags": true,
  "deviceManagementApplicabilityRuleOsEdition": null,
  "deviceManagementApplicabilityRuleOsVersion": null,
  "deviceManagementApplicabilityRuleDeviceMode": null,
  "createdDateTime": "2026-03-12T08:19:05.4757913Z",
  "description": "Deploys a custom .mobileconfig payload",
  "displayName": "iOS - Custom Configuration Example",
  "version": 1,
  "payloadName": "custom-config-profile-name",
  "payloadFileName": "com.example.custom.mobileconfig",
  "payload": "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4..."
}
```

No `deploymentChannel` in either direction.

## iosTrustedRootCertificate

```
Request URL:    https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations
Request Method: POST
```

```json
{
  "displayName": "iOS - Corporate Root CA",
  "roleScopeTagIds": ["0"],
  "@odata.type": "#microsoft.graph.iosTrustedRootCertificate",
  "certFileName": "CorporateRootCA.cer",
  "trustedRootCertificate": "MIIB1TCCAX8CFHl5T3FvbHJvb3RjZXJ0ZXhhbXBsZTANBgkqhkiG9w0BAQsFADAa..."
}
```

`trustedRootCertificate` is a base64 string on the wire and `[]byte` in the SDK. The provider
accepts the `filebase64()` output, decodes it before calling `SetTrustedRootCertificate`, and
re-encodes on read so state matches configuration.

## iosWiFiConfiguration

```
Request URL:    https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations
Request Method: POST
```

```json
{
  "displayName": "iOS - Corporate Wi-Fi",
  "roleScopeTagIds": ["0"],
  "@odata.type": "#microsoft.graph.iosWiFiConfiguration",
  "networkName": "Corporate Wi-Fi",
  "ssid": "CorpWiFi",
  "connectAutomatically": true,
  "connectWhenNetworkNameIsHidden": false,
  "wiFiSecurityType": "wpa2Personal",
  "preSharedKey": "<secret>",
  "disableMacAddressRandomization": false,
  "proxySettings": "manual",
  "proxyManualAddress": "proxy.example.com",
  "proxyManualPort": 8080
}
```

Response `201` — note `preSharedKey` is **absent**. Graph accepts the key on write and never
returns it. `state.go` therefore carries the configured value forward rather than mapping the
response, and the acceptance test sets `ImportStateVerifyIgnore` on it.

Note the wire name is `wiFiSecurityType` (capital `F`), not `wifiSecurityType`. The Terraform
attribute is `wifi_security_type`.

## Confirmed against a live tenant

The following POST bodies were captured from real profile creations and match this resource's
construct output field for field.

### iosScepCertificateProfile

```json
{
  "@odata.type": "#microsoft.graph.iosScepCertificateProfile",
  "displayName": "SCEP Cert",
  "roleScopeTagIds": ["0"],
  "renewalThresholdPercentage": 20,
  "certificateStore": "machine",
  "certificateValidityPeriodScale": "years",
  "certificateValidityPeriodValue": 1,
  "subjectNameFormat": "custom",
  "subjectNameFormatString": "CN={{AAD_Device_ID}}",
  "customSubjectAlternativeNames": [{"sanType": "domainNameService", "name": "device.name"}],
  "keyUsage": "digitalSignature,keyEncipherment",
  "keySize": "size2048",
  "rootCertificate@odata.bind": "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('<guid>')",
  "extendedKeyUsages": [{"name": "Client Authentication", "objectIdentifier": "1.3.6.1.5.5.7.3.2"}],
  "scepServerUrls": ["https://contoso.com/certsrv/mscep/mscep.dll"]
}
```

Two things this confirms:

- **`keyUsage` really is a comma-joined bitmask on the wire** — `"digitalSignature,keyEncipherment"`,
  not an array. The Set + bitwise-OR construct in `setKeyUsage` produces exactly this.
- **The `@odata.bind` URL form matches `normalizeODataBind` exactly**, including the quoted-GUID
  `deviceConfigurations('<guid>')` suffix.

### iosEnterpriseWiFiConfiguration

```json
{
  "@odata.type": "#microsoft.graph.iosEnterpriseWiFiConfiguration",
  "networkName": "NetworkName",
  "ssid": "NetworkSSID",
  "wiFiSecurityType": "wpa2Enterprise",
  "eapType": "eapTls",
  "authenticationMethod": "certificate",
  "trustedServerCertificateNames": ["srv.contoso.com"],
  "outerIdentityPrivacyTemporaryValue": "anonymous",
  "rootCertificatesForServerValidation@odata.bind": [
    "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('<guid>')"
  ],
  "identityCertificateForClientAuthentication@odata.bind":
    "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('<guid>')",
  "proxySettings": "manual",
  "proxyManualAddress": "10.0.0.1",
  "proxyManualPort": 8080,
  "preSharedKey": null
}
```

Confirms the asymmetry the schema models: **`rootCertificatesForServerValidation@odata.bind` is a JSON
array** while `identityCertificateForClientAuthentication@odata.bind` is a single string. Kiota's
`WriteAdditionalData` has an explicit `[]string` case, so passing a Go slice serializes correctly.

**Trap: `alreadySetPassword` is a UI-only artefact.** The Intune UI sends
`"alreadySetPassword": "********"` on enterprise Wi-Fi profiles. It does **not exist anywhere in the
SDK** — it is the UI's own placeholder for "a password is set, redacted for display". Do not copy it
out of a captured payload; it is not a settable property.

Note also `preSharedKey: null` — the UI sends the key explicitly as null on an enterprise profile,
consistent with this resource not exposing it on the `enterprise_wifi` block at all.

### iosCustomConfiguration and iosTrustedRootCertificate

```json
{"@odata.type": "#microsoft.graph.iosCustomConfiguration",
 "payloadFileName": "SimpleProfile.mobileconfig", "payload": "<base64>", "payloadName": "SimpleCustom"}

{"@odata.type": "#microsoft.graph.iosTrustedRootCertificate",
 "certFileName": "dummy-ca.crt", "trustedRootCertificate": "<base64>"}
```

Both confirm the three- and two-property surfaces, and — importantly — that **neither carries a
`deploymentChannel`**. That was the central macOS divergence this resource was built around.

### iosScepCertificateProfile on read — the @odata.bind question, settled

A plain `GET deviceManagement/deviceConfigurations('{id}')` of the SCEP profile created above returns:

```json
{
  "@odata.type": "#microsoft.graph.iosScepCertificateProfile",
  "@microsoft.graph.tips": "Use $select to choose only the properties your app needs...",
  "renewalThresholdPercentage": 20,
  "certificateStore": "machine",
  "certificateValidityPeriodScale": "years",
  "certificateValidityPeriodValue": 1,
  "subjectNameFormat": "custom",
  "subjectNameFormatString": "CN={{AAD_Device_ID}}",
  "subjectAlternativeNameType": null,
  "subjectAlternativeNameFormatString": null,
  "keyUsage": "keyEncipherment,digitalSignature",
  "keySize": "size2048",
  "scepServerUrls": ["https://contoso.com/certsrv/mscep/mscep.dll"],
  "extendedKeyUsages": [{"name": "Client Authentication", "objectIdentifier": "1.3.6.1.5.5.7.3.2"}],
  "customSubjectAlternativeNames": [{"sanType": "domainNameService", "name": "device.name"}]
}
```

**The certificate reference is absent entirely** — neither `rootCertificate@odata.bind` nor even an
unexpanded `rootCertificate` navigation property. This settles the open question: the prior-state
recovery in `mapIosScepCertificateProfile` is **load-bearing, not defensive**. Without it the
reference is lost from state on every read. The expanded-navigation-property fallback beside it only
fires if a caller adds `$expand=rootCertificate`, which this resource does not.

It also justifies `ImportStateVerifyIgnore` on `scep_certificate.root_certificate_odata_bind`: on
import there is genuinely nothing to recover the value from.

**`keyUsage` comes back in bit order, not submission order.** Sent
`"digitalSignature,keyEncipherment"`, returned `"keyEncipherment,digitalSignature"` — Graph
reconstructs the string by walking the bitmask low bit first (`keyEncipherment` = 1,
`digitalSignature` = 2). This is why `key_usage` is modelled as a **Set** rather than a List: order is
not meaningful and a List would diff on every read. `mapKeyUsageToSet` happens to decompose in the
same bit order regardless.

Two smaller confirmations:

- `subjectAlternativeNameType` returns **`null`** when never set, not the `"none"` enum member. Since
  it is a bitmask where `none` is bit 1 rather than the absence of bits, those are different values —
  `mapSubjectAlternativeNameTypeToSet` returns a null Set for a nil pointer, which is correct.
- `managedDeviceCertificateStates` is absent, confirming it was right to skip.
- `@microsoft.graph.tips` appears on **every** GET, not only the general configuration. Harmless for
  this resource (the typed SDK ignores unknown properties) but it is the annotation that would have
  leaked into the `_json` sibling's `settings_json`.

## Read with assignments

```
Request URL:    https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations/{id}?$expand=assignments
Request Method: GET
```

Two things about the response shape that drive how state mapping is written:

1. **Every property is returned, including ones never set.** A profile that sets a single property
   comes back with roughly 190 keys, the rest as `false`, `null` or `[]`. Unset scalars are explicit
   `null` rather than omitted — `description` included, which is why it maps to `types.StringNull()`
   rather than an empty string.
2. **`assignments` arrives with a sibling `assignments@odata.context` annotation** next to it, in
   addition to the root `@odata.context`. This matters for the `_json` sibling resource, which
   strips server-owned keys out of a raw JSON body; a fixed-name strip of `@odata.context` alone
   leaks the annotation. It does not affect this resource, which uses the typed SDK.

`isAssigned` does **not** appear in the response, even with `$expand=assignments`.

Assignment targets come back as:

```json
"assignments": [
  {
    "id": "0e2f87bd-de2c-47c2-a69b-c8cbe9823042_93f57a7f-e646-4a5a-b6a4-7c19e6ca6244",
    "source": "direct",
    "sourceId": "0e2f87bd-de2c-47c2-a69b-c8cbe9823042",
    "intent": "apply",
    "target": {
      "@odata.type": "#microsoft.graph.groupAssignmentTarget",
      "deviceAndAppManagementAssignmentFilterId": null,
      "deviceAndAppManagementAssignmentFilterType": "none",
      "groupId": "93f57a7f-e646-4a5a-b6a4-7c19e6ca6244"
    }
  }
]
```

A null `deviceAndAppManagementAssignmentFilterId` is normalised to the zero GUID in state so set
hashing stays stable against the schema default.
