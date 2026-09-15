package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// capturedResponsePath is a verbatim copy of a real
// `GET beta/deviceManagement/deviceConfigurations/{id}?$expand=assignments` response for an
// iosGeneralDeviceConfiguration. serverOwnedRootKeys and the pattern pass were both derived from it,
// so the tests below assert against the wire shape rather than against a hand-written approximation.
const capturedResponsePath = "tests/responses/validate_get/get_ios_general_device_configuration.json"

func loadCapturedResponse(t *testing.T) map[string]any {
	t.Helper()

	raw, err := os.ReadFile(capturedResponsePath)
	require.NoError(t, err, "captured response fixture must be readable")

	var body map[string]any
	require.NoError(t, json.Unmarshal(raw, &body), "captured response must be valid JSON")

	return body
}

// The regression this whole strip design exists for: the response carries a sibling
// "assignments@odata.context" annotation next to "assignments", so a fixed-name strip of
// "@odata.context" alone would leak it into settings_json.
func TestStripRemovesAllRootODataKeys(t *testing.T) {
	body := loadCapturedResponse(t)

	// Guard the fixture itself: if these ever stop being present the test below proves nothing.
	require.Contains(t, body, "@odata.context", "fixture must carry the root @odata.context")
	require.Contains(t, body, "@odata.type", "fixture must carry @odata.type")
	require.Contains(t, body, "assignments@odata.context",
		"fixture must carry the sibling annotation — this is the case the pattern pass exists for")
	require.Contains(t, body, "@microsoft.graph.tips",
		"fixture must carry the Microsoft service annotation observed on live responses")

	stripServerOwnedKeys(body)

	for key := range body {
		assert.NotContains(t, key, "@",
			"no root key may contain an annotation marker after stripping")
	}
}

// Not every Graph annotation is an @odata. one. Live GET responses carry "@microsoft.graph.tips",
// whose value is a paragraph of advisory text about using $select. An "@odata."-only pattern misses
// it entirely and writes that prose into settings_json, where the operator can neither remove nor
// reproduce it.
func TestStripRemovesNonODataAnnotations(t *testing.T) {
	body := map[string]any{
		"@odata.context":          "https://graph.microsoft.com/beta/$metadata#...",
		"@odata.type":             "#microsoft.graph.iosGeneralDeviceConfiguration",
		"@microsoft.graph.tips":   "Use $select to choose only the properties your app needs...",
		"assignments@odata.count": 2,
		"cameraBlocked":           true,
	}

	stripServerOwnedKeys(body)

	assert.NotContains(t, body, "@microsoft.graph.tips",
		"a vendor annotation is not a setting")
	assert.NotContains(t, body, "assignments@odata.count",
		"property-scoped annotations of any kind must go, not just @odata.context")
	assert.Equal(t, map[string]any{"cameraBlocked": true}, body,
		"only the real setting may remain")
}

// Guard the annotation predicate directly, since it is the structural rule the pattern pass rests on.
func TestIsODataAnnotation(t *testing.T) {
	for _, key := range []string{
		"@odata.context",
		"@odata.type",
		"@odata.nextLink",
		"@microsoft.graph.tips",
		"assignments@odata.context",
		"assignments@odata.count",
	} {
		assert.True(t, isODataAnnotation(key), "%s is an annotation", key)
	}

	// Real Graph property names never contain "@".
	for _, key := range []string{
		"cameraBlocked",
		"homeScreenPages",
		"emailInDomainSuffixes",
		"mediaContentRatingApps",
		"iosSingleSignOnExtension",
	} {
		assert.False(t, isODataAnnotation(key), "%s is a real property", key)
	}
}

func TestStripRemovesEveryServerOwnedKey(t *testing.T) {
	body := loadCapturedResponse(t)

	stripServerOwnedKeys(body)

	for _, key := range serverOwnedRootKeys {
		assert.NotContains(t, body, key,
			"%s is provider- or server-owned and must not reach settings_json", key)
	}
}

// The strip must remove only what it claims to. cameraBlocked is the one property this profile
// actually sets, so its survival is the positive control.
func TestStripPreservesSettings(t *testing.T) {
	body := loadCapturedResponse(t)

	stripServerOwnedKeys(body)

	assert.Equal(t, true, body["cameraBlocked"],
		"the one configured property must survive with its value")

	// A representative sample across the value shapes Graph returns.
	assert.Contains(t, body, "airDropBlocked", "unset booleans are returned as false and kept")
	assert.Contains(t, body, "passcodeMinimumLength", "unset scalars are returned as null and kept")
	assert.Contains(t, body, "emailInDomainSuffixes", "unset collections are returned as [] and kept")
	assert.Contains(t, body, "mediaContentRatingApps", "enum-valued properties are kept")
	assert.Contains(t, body, "networkUsageRules", "object collections are kept")

	// Keys that merely resemble stripped ones must not be caught by the named pass.
	assert.NotContains(t, serverOwnedRootKeys, "kioskModeAppType",
		"sanity: the strip list must not contain real settings")
}

// Graph returns every property, not just the configured ones. This is what makes projection
// load-bearing rather than cosmetic, so it is worth pinning down explicitly.
func TestStripLeavesFullPropertySurface(t *testing.T) {
	body := loadCapturedResponse(t)
	before := len(body)

	stripServerOwnedKeys(body)

	assert.Greater(t, len(body), 150,
		"a profile setting one property still returns its whole surface (~190 keys); "+
			"projection is what stops that diffing against a small configuration")
	assert.Less(t, len(body), before, "stripping must actually remove keys")

	// Count how many of the remaining properties are falsy — the ones indistinguishable from unset.
	falsy := 0
	for _, value := range body {
		switch v := value.(type) {
		case bool:
			if !v {
				falsy++
			}
		case nil:
			falsy++
		}
	}
	assert.Greater(t, falsy, 100,
		"most returned properties are false or null; false is indistinguishable from unset, "+
			"which is why removing a property stops tracking rather than resetting it")
}

// isAssigned did not appear in the observed response even with $expand=assignments. It stays in the
// strip list defensively, but nothing may depend on its presence.
func TestStripHandlesIsAssignedDefensively(t *testing.T) {
	body := loadCapturedResponse(t)
	require.NotContains(t, body, "isAssigned",
		"fixture reflects reality: Graph did not return isAssigned")

	assert.Contains(t, serverOwnedRootKeys, "isAssigned",
		"kept as a defensive delete in case Graph starts returning it")

	// Injecting it must still result in removal.
	body["isAssigned"] = true
	stripServerOwnedKeys(body)
	assert.NotContains(t, body, "isAssigned")
}

// A restrictions profile mixes scalars with object-valued properties and object arrays. Those are the
// shapes most likely to break projection, so the fixture carries populated examples of each and this
// test pins down that the strip leaves them intact.
func TestStripPreservesNestedObjectSettings(t *testing.T) {
	body := loadCapturedResponse(t)

	stripServerOwnedKeys(body)

	// mediaContentRating* are objects, not scalars — mediaContentRatingCanada is {movieRating,
	// tvRating}. Only mediaContentRatingApps is a plain enum, which is easy to conflate.
	canada, ok := body["mediaContentRatingCanada"].(map[string]any)
	require.True(t, ok, "mediaContentRatingCanada is a nested object and must survive")
	assert.Equal(t, "allAllowed", canada["movieRating"])
	assert.Equal(t, "allAllowed", canada["tvRating"])
	assert.Equal(t, "allAllowed", body["mediaContentRatingApps"],
		"mediaContentRatingApps is a scalar enum, unlike the per-country ratings")

	// App collections are []appListItem. Note the entries carry no @odata.type: the collection's
	// element type is concrete, so Graph omits the discriminator.
	for _, key := range []string{"appsSingleAppModeList", "compliantAppsList"} {
		list, ok := body[key].([]any)
		require.True(t, ok, "%s must survive as an array", key)
		require.NotEmpty(t, list, "%s must be populated in the fixture", key)

		entry, ok := list[0].(map[string]any)
		require.True(t, ok)
		assert.Contains(t, entry, "appId", "%s entries are appListItem objects", key)
		assert.NotContains(t, entry, "@odata.type",
			"%s has a concrete element type, so Graph omits the discriminator", key)
	}

	// networkUsageRules is []iosNetworkUsageRule, and real profiles send sparse entries — each
	// element setting only the one property it cares about.
	rules, ok := body["networkUsageRules"].([]any)
	require.True(t, ok, "networkUsageRules must survive")
	require.Len(t, rules, 2)

	first, ok := rules[0].(map[string]any)
	require.True(t, ok)
	assert.Contains(t, first, "cellularDataBlocked")

	second, ok := rules[1].(map[string]any)
	require.True(t, ok)
	assert.Contains(t, second, "cellularDataBlockWhenRoaming",
		"sparse rule entries must survive with their own distinct keys")
}

// iosDeviceFeaturesConfiguration responses carry nine server-computed single sign-on metadata keys
// that have no setters anywhere in the SDK. Graph derives them from the configured
// iosSingleSignOnExtension and returns them as read-only descriptions of which properties each
// extension variant supports. They are large — the "all" and "nonAzureAd" lists run to ~29 entries
// each on a real profile — so leaking them into settings_json would leave an operator with a large
// block they cannot remove without appearing to change configuration.
func TestStripRemovesServerComputedSSOMetadata(t *testing.T) {
	raw, err := os.ReadFile("tests/responses/validate_get/get_ios_device_features_configuration.json")
	require.NoError(t, err)

	var body map[string]any
	require.NoError(t, json.Unmarshal(raw, &body))

	ssoMetadataKeys := []string{
		"ssoExtensionKey",
		"allExtensibleSSOEntities",
		"azureAdExtensibleSSOEntities",
		"commonCredentialKerberosExtensibleSSOEntities",
		"commonCredentialRedirectExtensibleSSOEntities",
		"kerberosExtensibleSSOEntities",
		"nonAzureAdExtensibleSSOEntities",
		"nonKerberosExtensibleSSOEntities",
		"redirectExtensibleSSOEntities",
	}

	// Guard the fixture: if these stop being present the assertions below prove nothing.
	for _, key := range ssoMetadataKeys {
		require.Contains(t, body, key, "fixture must carry the observed SSO metadata key %s", key)
	}

	stripServerOwnedKeys(body)

	for _, key := range ssoMetadataKeys {
		assert.NotContains(t, body, key,
			"%s is server-computed with no SDK setter and must not reach settings_json", key)
	}

	// The genuinely settable property must survive: only the derived metadata about it is stripped.
	assert.Contains(t, body, "iosSingleSignOnExtension",
		"iosSingleSignOnExtension is settable and must be preserved")
	assert.Contains(t, body, "singleSignOnSettings",
		"singleSignOnSettings is settable and must be preserved")
}

// The pattern pass must not recurse: nested @odata.type discriminators are load-bearing. This is the
// single most destructive mistake available in this file, so it is asserted directly.
func TestStripDoesNotTouchNestedODataDiscriminators(t *testing.T) {
	body := map[string]any{
		"@odata.context": "https://graph.microsoft.com/beta/$metadata#...",
		"@odata.type":    "#microsoft.graph.iosDeviceFeaturesConfiguration",
		"id":             "0e2f87bd-de2c-47c2-a69b-c8cbe9823042",
		// Four levels of nested discriminator: page -> icons -> folder -> pages -> apps -> app.
		// Note the two page types use different property names — iosHomeScreenPage has "icons" and may
		// hold folders, while iosHomeScreenFolderPage has "apps" and may not. Graph rejects the wrong
		// one outright, so the fixture mirrors the real shape.
		"homeScreenPages": []any{
			map[string]any{
				"@odata.type": "#microsoft.graph.iosHomeScreenPage",
				"displayName": "Page 1",
				"icons": []any{
					map[string]any{
						"@odata.type": "#microsoft.graph.iosHomeScreenApp",
						"displayName": "Safari",
						"bundleID":    "com.apple.mobilesafari",
					},
					map[string]any{
						"@odata.type": "#microsoft.graph.iosHomeScreenFolder",
						"displayName": "Utilities",
						"pages": []any{
							map[string]any{
								"@odata.type": "#microsoft.graph.iosHomeScreenFolderPage",
								"displayName": "Utilities Page 1",
								"apps": []any{
									map[string]any{
										"@odata.type": "#microsoft.graph.iosHomeScreenApp",
										"displayName": "Calculator",
										"bundleID":    "com.apple.calculator",
									},
								},
							},
						},
					},
				},
			},
		},
		"singleSignOnExtension": map[string]any{
			"@odata.type": "#microsoft.graph.iosKerberosSingleSignOnExtension",
			"realm":       "EXAMPLE.COM",
		},
		// contentFilterSettings is polymorphic: the two implementations share no properties, so the
		// discriminator carries all the meaning. Its bookmark entries nest one further.
		"contentFilterSettings": map[string]any{
			"@odata.type": "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess",
			"websiteList": []any{
				map[string]any{
					"@odata.type":    "#microsoft.graph.iosBookmark",
					"url":            "https://intranet.example.com",
					"displayName":    "Company Intranet",
					"bookmarkFolder": "Work",
				},
			},
		},
		// wallpaperImage is a mimeContent {type, value}. Included because "type" is the sort of
		// innocuous-looking key a careless strip might target.
		"wallpaperImage": map[string]any{
			"type":  "image/jpeg",
			"value": "/9j/4Q//9k=",
		},
	}

	stripServerOwnedKeys(body)

	// Root annotations gone.
	assert.NotContains(t, body, "@odata.context")
	assert.NotContains(t, body, "@odata.type")
	assert.NotContains(t, body, "id")

	// Nested discriminators intact — Graph cannot resolve these polymorphic members without them.
	pages, ok := body["homeScreenPages"].([]any)
	require.True(t, ok, "homeScreenPages must survive")
	require.Len(t, pages, 1)

	page, ok := pages[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "#microsoft.graph.iosHomeScreenPage", page["@odata.type"],
		"page discriminator must survive")
	assert.Equal(t, "Page 1", page["displayName"],
		"a nested displayName is a real setting, not an envelope key")

	icons, ok := page["icons"].([]any)
	require.True(t, ok, "top-level pages use icons")
	require.Len(t, icons, 2)

	app, ok := icons[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "#microsoft.graph.iosHomeScreenApp", app["@odata.type"],
		"nested app discriminator must survive")

	folder, ok := icons[1].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "#microsoft.graph.iosHomeScreenFolder", folder["@odata.type"],
		"nested folder discriminator must survive")

	// Deepest level: a folder page uses "apps", not "icons", and its own discriminator plus the app's
	// must both survive.
	folderPages, ok := folder["pages"].([]any)
	require.True(t, ok, "folder pages must survive")
	require.Len(t, folderPages, 1)

	folderPage, ok := folderPages[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "#microsoft.graph.iosHomeScreenFolderPage", folderPage["@odata.type"],
		"folder page discriminator must survive")
	assert.NotContains(t, folderPage, "icons",
		"a folder page uses apps, not icons — Graph rejects icons on this type")

	folderApps, ok := folderPage["apps"].([]any)
	require.True(t, ok, "folder pages use apps")
	require.Len(t, folderApps, 1)

	folderApp, ok := folderApps[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "#microsoft.graph.iosHomeScreenApp", folderApp["@odata.type"],
		"four-levels-deep app discriminator must survive")

	sso, ok := body["singleSignOnExtension"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "#microsoft.graph.iosKerberosSingleSignOnExtension", sso["@odata.type"],
		"nested extension discriminator must survive")

	filter, ok := body["contentFilterSettings"].(map[string]any)
	require.True(t, ok, "contentFilterSettings must survive")
	assert.Equal(t, "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess", filter["@odata.type"],
		"content filter discriminator must survive — the two implementations share no properties")

	websites, ok := filter["websiteList"].([]any)
	require.True(t, ok)
	require.Len(t, websites, 1)

	bookmark, ok := websites[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "#microsoft.graph.iosBookmark", bookmark["@odata.type"],
		"bookmark discriminator must survive")
	assert.Equal(t, "Work", bookmark["bookmarkFolder"],
		"the field is bookmarkFolder, not bookmarkFolderName")

	wallpaper, ok := body["wallpaperImage"].(map[string]any)
	require.True(t, ok, "wallpaperImage must survive")
	assert.Equal(t, "image/jpeg", wallpaper["type"],
		"a nested 'type' key is a real mimeContent property, not something to strip")
	assert.Equal(t, "/9j/4Q//9k=", wallpaper["value"])
}

func TestStripIsIdempotent(t *testing.T) {
	body := loadCapturedResponse(t)

	stripServerOwnedKeys(body)
	afterFirst := len(body)

	stripServerOwnedKeys(body)

	assert.Equal(t, afterFirst, len(body), "a second strip must be a no-op")
}

func TestStripHandlesEmptyBody(t *testing.T) {
	body := map[string]any{}

	stripServerOwnedKeys(body)

	assert.Empty(t, body)
}

// The strip list and the validator's envelope list overlap but are not the same: the validator only
// rejects what an operator could wrongly supply, while the strip also removes server-owned fields.
func TestEnvelopeKeysAreASubsetOfServerOwnedKeys(t *testing.T) {
	for _, key := range envelopeKeys {
		assert.Contains(t, serverOwnedRootKeys, key,
			"every provider-owned envelope key must also be stripped from responses")
	}
}

// Guard against a stray leading/trailing space or casing slip in the strip list, which would
// silently stop matching.
func TestServerOwnedRootKeysAreWellFormed(t *testing.T) {
	seen := map[string]bool{}
	for _, key := range serverOwnedRootKeys {
		assert.Equal(t, strings.TrimSpace(key), key, "key must not have surrounding whitespace")
		assert.NotEmpty(t, key)
		assert.False(t, seen[key], "duplicate key in serverOwnedRootKeys: %s", key)
		seen[key] = true
	}
}
