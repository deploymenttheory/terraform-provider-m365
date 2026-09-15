package graphBetaIosDeviceConfigurationTemplatesJson

import "strings"

// serverOwnedRootKeys are the root-level properties the provider owns or the server controls, all of
// which must be removed from a GET response before it becomes settings_json.
//
// This list was derived from an observed
// `GET beta/deviceManagement/deviceConfigurations/{id}?$expand=assignments` response for an
// iosGeneralDeviceConfiguration, not from the SDK surface. Two things that observation corrected:
//
//   - isAssigned does NOT appear, even with $expand=assignments. It is kept here as a defensive
//     delete in case Graph starts returning it, but nothing relies on its presence.
//   - assignments arrives with a sibling "assignments@odata.context" annotation alongside it. A
//     fixed-name strip of "@odata.context" alone would leak that annotation into settings_json,
//     which is why stripServerOwnedKeys also runs a pattern pass (see below).
//
// @odata.type and @odata.context are listed for documentation only — the pattern pass removes them
// first.
var serverOwnedRootKeys = []string{
	// Provider-owned: surfaced as dedicated attributes.
	"@odata.type",
	"displayName",
	"description",
	"roleScopeTagIds",
	// Server-owned.
	"@odata.context",
	"id",
	"createdDateTime",
	"lastModifiedDateTime",
	"version",
	"supportsScopeTags",
	"deviceManagementApplicabilityRuleOsEdition",
	"deviceManagementApplicabilityRuleOsVersion",
	"deviceManagementApplicabilityRuleDeviceMode",
	// Modelled as its own block.
	"assignments",
	// Defensive only: not observed in any response.
	"isAssigned",

	// Server-computed single sign-on metadata, observed on iosDeviceFeaturesConfiguration responses.
	//
	// None of these exist as setters anywhere in the SDK — Graph derives them from the configured
	// iosSingleSignOnExtension and returns them as read-only descriptions of which properties each
	// extension variant supports. They are large (the two "all"/"nonAzureAd" lists run to ~29 entries
	// each) and would otherwise be written verbatim into settings_json on import, where an operator
	// could not remove them without appearing to change configuration.
	//
	// Note the actual settable property, iosSingleSignOnExtension, is NOT stripped — only this
	// derived metadata about it.
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

// stripServerOwnedKeys removes provider-owned and server-owned properties from a decoded GET
// response, leaving only the settings tree.
//
// Two passes:
//
//  1. Pattern pass — deletes any ROOT key containing an "@" annotation marker. In OData, "@" always
//     introduces an annotation rather than a property, so this is a structural rule rather than a
//     list to maintain. It covers:
//     - "@odata.context", "@odata.type" — standard OData annotations
//     - "assignments@odata.context"     — a property-scoped annotation sibling
//     - "@microsoft.graph.tips"         — a Microsoft service annotation carrying advisory text,
//     observed on live GET responses. It matches no "@odata."
//     pattern, so an earlier "@odata."-only rule would have
//     written a paragraph of performance advice into
//     settings_json.
//     Any future "@odata.nextLink", "<nav>@odata.count" or vendor annotation is covered without a
//     code change.
//
//  2. Named pass — deletes each key in serverOwnedRootKeys.
//
// The pattern pass is deliberately restricted to the root. Nested "@odata.type" discriminators are
// load-bearing — homeScreenPages[].icons[] entries are distinguished only by
// #microsoft.graph.iosHomeScreenApp vs iosHomeScreenFolder, and singleSignOnExtension,
// contentFilterSettings and notificationSettings[] all rely on theirs. Recursing would silently
// destroy them, so this must not be generalized into a tree walk. The same reasoning drives the
// no-recursion rule in settingsJSONEnvelopeValidator.
//
// The input map is mutated in place.
func stripServerOwnedKeys(body map[string]any) {
	for key := range body {
		if isODataAnnotation(key) {
			delete(body, key)
		}
	}

	for _, key := range serverOwnedRootKeys {
		delete(body, key)
	}
}

// isODataAnnotation reports whether a root key is an OData annotation rather than a property.
//
// Annotations either begin with "@" (instance annotations such as "@odata.context" or
// "@microsoft.graph.tips") or contain "@" after a property name (property-scoped annotations such as
// "assignments@odata.context"). No real Graph property name contains "@", so the presence of one is
// sufficient.
func isODataAnnotation(key string) bool {
	return strings.Contains(key, "@")
}
