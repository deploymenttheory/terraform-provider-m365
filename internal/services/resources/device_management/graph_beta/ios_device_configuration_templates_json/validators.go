package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// envelopeKeys are the root-level keys the provider injects into the request body from dedicated
// top-level attributes. Supplying them inside settings_json would mean two sources of truth for the
// same property, so they are rejected.
var envelopeKeys = []string{
	"@odata.type",
	"displayName",
	"description",
	"roleScopeTagIds",
}

// settingsJSONEnvelopeValidator rejects envelope keys at the ROOT of settings_json.
type settingsJSONEnvelopeValidator struct{}

func settingsJSONEnvelope() validator.String {
	return settingsJSONEnvelopeValidator{}
}

func (v settingsJSONEnvelopeValidator) Description(_ context.Context) string {
	return fmt.Sprintf(
		"settings_json must not contain the provider-owned root keys: %s",
		strings.Join(envelopeKeys, ", "),
	)
}

func (v settingsJSONEnvelopeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateString checks only the top level of the JSON object — it deliberately does not recurse.
//
// Nested "@odata.type" discriminators are required by Graph: homeScreenPages[].icons[] entries are
// distinguished solely by #microsoft.graph.iosHomeScreenApp vs iosHomeScreenFolder, and
// singleSignOnExtension, contentFilterSettings and notificationSettings[] each carry their own. A
// recursive check would reject every valid iosDeviceFeaturesConfiguration. The equivalent boundary
// exists in stripServerOwnedKeys for the read path.
func (v settingsJSONEnvelopeValidator) ValidateString(
	ctx context.Context,
	req validator.StringRequest,
	resp *validator.StringResponse,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var body map[string]any
	if err := json.Unmarshal([]byte(req.ConfigValue.ValueString()), &body); err != nil {
		// Malformed JSON and non-object JSON are both JSONSchemaValidator's concern; this validator
		// only has an opinion about which keys a valid object may contain.
		return
	}

	found := make([]string, 0, len(envelopeKeys))
	for _, key := range envelopeKeys {
		if _, exists := body[key]; exists {
			found = append(found, key)
		}
	}

	if len(found) == 0 {
		return
	}

	// Map iteration order is random; sort so the diagnostic is stable across runs.
	sort.Strings(found)

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Provider-Owned Keys in settings_json",
		fmt.Sprintf(
			"settings_json must contain only the iOS/iPadOS settings tree. These root keys are set "+
				"by the provider from dedicated attributes and must be removed: %s.\n\n"+
				"Use the odata_type, display_name, description and role_scope_tag_ids attributes "+
				"instead.\n\nNote that nested @odata.type discriminators (for example inside "+
				"homeScreenPages[].icons[]) are required by Graph and are not affected by this rule "+
				"— only the root level is checked.",
			strings.Join(found, ", "),
		),
	)
}
