package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	s "github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/constructors"
)

var (
	errUnmarshalSettingsJSON = errors.New("failed to parse settings_json")
	errExtractRoleScopeTags  = errors.New("failed to extract role scope tag ids")
)

// rawJSONRequestBody adapts an arbitrary JSON object to the s.Parsable interface that
// custom_requests requires (its POST/PATCH helpers call SetContentFromParsable, which will not
// accept a bare map).
//
// Kiota's WriteAnyValue delegates to json.Marshal, so arbitrarily nested trees round-trip
// byte-exactly — which is the point: the operator's settings tree reaches Graph unaltered.
type rawJSONRequestBody struct {
	body map[string]any
}

// Serialize writes each top-level key. Note that Go map iteration order is random, so the emitted
// key order varies between runs; mock responders must decode the body rather than compare raw
// strings.
func (b *rawJSONRequestBody) Serialize(writer s.SerializationWriter) error {
	for key, value := range b.body {
		// WriteAnyValue drops nil interfaces silently, which would turn an explicit JSON null into
		// an absent property. Writing the null explicitly preserves "unset this" semantics.
		if value == nil {
			if err := writer.WriteNullValue(key); err != nil {
				return fmt.Errorf("writing null value for %q: %w", key, err)
			}
			continue
		}
		if err := writer.WriteAnyValue(key, value); err != nil {
			return fmt.Errorf("writing value for %q: %w", key, err)
		}
	}
	return nil
}

// GetFieldDeserializers returns an empty map: this type is write-only. Responses are read as raw
// JSON via customrequests.GetRequestByResourceId rather than deserialized through Kiota.
func (b *rawJSONRequestBody) GetFieldDeserializers() map[string]func(s.ParseNode) error {
	return map[string]func(s.ParseNode) error{}
}

// constructResource builds the full request body: the operator's settings tree plus the envelope
// properties the provider owns.
func constructResource(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesJsonResourceModel,
) (*rawJSONRequestBody, error) {
	tflog.Debug(ctx, fmt.Sprintf("Constructing %s resource from model", ResourceName))

	body := map[string]any{}
	if !data.SettingsJson.IsNull() && !data.SettingsJson.IsUnknown() {
		if err := json.Unmarshal([]byte(data.SettingsJson.ValueString()), &body); err != nil {
			return nil, fmt.Errorf("%w: %w", errUnmarshalSettingsJSON, err)
		}
	}

	// Defensively remove the envelope keys. settingsJSONEnvelopeValidator already rejects them in
	// configuration, but validation does not run on the import path, so a settings_json recovered
	// from a hand-edited state file could still carry them.
	for _, key := range envelopeKeys {
		delete(body, key)
	}

	body["@odata.type"] = data.OdataType.ValueString()
	body["displayName"] = data.DisplayName.ValueString()

	// Send description only when set: an empty string is a meaningful value to Graph, distinct from
	// omitting the property.
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body["description"] = data.Description.ValueString()
	}

	if !data.RoleScopeTagIds.IsNull() && !data.RoleScopeTagIds.IsUnknown() {
		var tagIds []string
		if diags := data.RoleScopeTagIds.ElementsAs(ctx, &tagIds, false); diags.HasError() {
			return nil, fmt.Errorf("%w: %v", errExtractRoleScopeTags, diags.Errors())
		}
		body["roleScopeTagIds"] = tagIds
	}

	requestBody := &rawJSONRequestBody{body: body}

	if err := constructors.DebugLogGraphObject(
		ctx,
		fmt.Sprintf("Final JSON to be sent to Graph API for resource %s", ResourceName),
		requestBody,
	); err != nil {
		tflog.Error(ctx, "Failed to debug log object", map[string]any{
			"error": err.Error(),
		})
	}

	tflog.Debug(ctx, fmt.Sprintf("Finished constructing %s resource", ResourceName))

	return requestBody, nil
}
