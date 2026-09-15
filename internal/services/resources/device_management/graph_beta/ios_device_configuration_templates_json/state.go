package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/normalize"
)

// MapRemoteResourceStateToTerraform maps a raw GET response onto the Terraform model.
//
// The typed attributes are hydrated from the envelope properties, then settings_json is derived by
// strip → project → normalize:
//
//  1. Strip the provider-owned and server-owned root keys (see strip.go).
//  2. Project what remains onto the shape of the prior settings_json.
//  3. Sort keys alphabetically so the stored string is stable.
//
// Projection is load-bearing rather than cosmetic. Graph returns EVERY property of these types,
// including ones never set — an observed response for a profile setting a single property came back
// with roughly 190 keys, the rest as false, null or []. Without projection a three-key configuration
// would diff against all of them on every plan.
//
// A consequence worth knowing: because Graph reports an unset boolean as false, there is no way to
// distinguish "declared false" from "never declared". The prior configuration's shape is the only
// record of which properties are managed, which is why removing a property from settings_json stops
// tracking it rather than resetting it on the server.
//
// On any failure the prior settings_json is left in place rather than writing a partial value, the
// same defensive posture as normalizeSettingsCatalogJSONArray in the settings catalog resource.
func MapRemoteResourceStateToTerraform(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesJsonResourceModel,
	rawResponse []byte,
) {
	if len(rawResponse) == 0 {
		tflog.Debug(ctx, "Remote response is empty, leaving state unchanged")
		return
	}

	var body map[string]any
	if err := json.Unmarshal(rawResponse, &body); err != nil {
		tflog.Error(ctx, "Failed to unmarshal device configuration response", map[string]any{
			"error": err.Error(),
		})
		return
	}

	// Hydrate the typed attributes from the envelope before it is stripped away.
	if id, ok := body["id"].(string); ok {
		data.ID = types.StringValue(id)
	}
	if odataType, ok := body["@odata.type"].(string); ok {
		data.OdataType = types.StringValue(odataType)
	}
	if displayName, ok := body["displayName"].(string); ok {
		data.DisplayName = types.StringValue(displayName)
	}

	// description is returned as an explicit null rather than being omitted when unset, so a failed
	// string assertion means null and must map to StringNull rather than an empty string.
	if description, ok := body["description"].(string); ok {
		data.Description = types.StringValue(description)
	} else {
		data.Description = types.StringNull()
	}

	data.RoleScopeTagIds = roleScopeTagIdsToSet(ctx, body["roleScopeTagIds"])

	data.SettingsJson = projectSettingsJSON(ctx, data.SettingsJson, body)

	tflog.Debug(ctx, "Finished mapping remote state to Terraform state", map[string]any{
		"resourceId": data.ID.ValueString(),
	})
}

// projectSettingsJSON derives the settings_json value from an already-decoded response body. Returns
// prior unchanged if anything goes wrong, so a transient failure cannot corrupt state.
func projectSettingsJSON(
	ctx context.Context,
	prior types.String,
	body map[string]any,
) types.String {
	stripServerOwnedKeys(body)

	// On import there is no prior configuration to project against, so the whole stripped response
	// is kept and the operator prunes it. Expect this to be large — Graph returns the full property
	// surface, not just what was set.
	if prior.IsNull() || prior.IsUnknown() || prior.ValueString() == "" {
		tflog.Debug(
			ctx,
			"No prior settings_json to project against (import); keeping the full response",
		)
		return marshalNormalized(ctx, body, prior)
	}

	var shape any
	if err := json.Unmarshal([]byte(prior.ValueString()), &shape); err != nil {
		tflog.Error(ctx, "Failed to unmarshal prior settings_json, leaving state unchanged",
			map[string]any{"error": err.Error()})
		return prior
	}

	projected := normalize.ProjectOntoShape(shape, body)

	return marshalNormalized(ctx, projected, prior)
}

// marshalNormalized serializes value with alphabetically sorted keys, falling back to prior on error.
func marshalNormalized(ctx context.Context, value any, prior types.String) types.String {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		tflog.Error(ctx, "Failed to marshal settings JSON, leaving state unchanged",
			map[string]any{"error": err.Error()})
		return prior
	}

	normalized, err := normalize.JSONAlphabetically(string(jsonBytes))
	if err != nil {
		tflog.Error(ctx, "Failed to normalize settings JSON, leaving state unchanged",
			map[string]any{"error": err.Error()})
		return prior
	}

	return types.StringValue(normalized)
}

// roleScopeTagIdsToSet converts the decoded roleScopeTagIds value into a set of strings.
func roleScopeTagIdsToSet(ctx context.Context, value any) types.Set {
	raw, ok := value.([]any)
	if !ok {
		return types.SetNull(types.StringType)
	}
	if len(raw) == 0 {
		return types.SetValueMust(types.StringType, []attr.Value{})
	}

	tagIds := make([]string, 0, len(raw))
	for _, item := range raw {
		if tagId, ok := item.(string); ok {
			tagIds = append(tagIds, tagId)
		}
	}

	setValue, diags := types.SetValueFrom(ctx, types.StringType, tagIds)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create role scope tag ids set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(types.StringType)
	}

	return setValue
}
