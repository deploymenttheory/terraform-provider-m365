package graphBetaNetworkConditionalAccessSettings

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"
	models "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
	builders "github.com/microsoftgraph/msgraph-beta-sdk-go/networkaccess"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/convert"
)

var errMissingSignalingStatus = errors.New(
	"signaling_status must be a known enabled or disabled value before writing",
)

func constructResource(value types.String) (models.ConditionalAccessSettingsable, error) {
	if value.IsNull() || value.IsUnknown() {
		return nil, errMissingSignalingStatus
	}
	body := models.NewConditionalAccessSettings()
	if err := convert.FrameworkToGraphEnum(
		value,
		models.ParseStatus,
		body.SetSignalingStatus,
	); err != nil {
		return nil, fmt.Errorf("construct signaling status: %w", err)
	}
	return body, nil
}

//nolint:wrapcheck // Shared Graph diagnostics need the original SDK error type.
func (r *NetworkConditionalAccessSettingsResource) patchSettings(
	ctx context.Context,
	body models.ConditionalAccessSettingsable,
) error {
	_, err := r.client.NetworkAccess().Settings().ConditionalAccess().Patch(ctx, body,
		&builders.SettingsConditionalAccessRequestBuilderPatchRequestConfiguration{
			// Avoid Kiota replaying an uncompressed body with a stale gzip header on retry.
			Options: []abstractions.RequestOption{kiotahttp.NewCompressionOptionsReference(false)},
		})
	return err
}

var errInvalidSettingsResponse = errors.New("invalid conditional access settings response")

//nolint:wrapcheck // Preserve SDK errors for shared Graph diagnostics.
func (r *NetworkConditionalAccessSettingsResource) getSettings(
	ctx context.Context,
) (*conditionalAccessSettingsResponse, error) {
	request, err := r.client.NetworkAccess().
		Settings().
		ConditionalAccess().
		ToGetRequestInformation(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Kiota's backing-store JSON parser panics on a JSON null root. Read bytes through the
	// same authenticated adapter and decode explicitly so malformed responses preserve state.
	result, err := r.client.GetAdapter().
		SendPrimitive(ctx, request, "[]byte", abstractions.ErrorMappings{"XXX": odataerrors.CreateODataErrorFromDiscriminatorValue})
	if err != nil {
		return nil, err
	}
	raw, ok := result.([]byte)
	if !ok || len(raw) == 0 {
		return nil, errInvalidSettingsResponse
	}
	var settings *conditionalAccessSettingsResponse
	if err := json.Unmarshal(raw, &settings); err != nil {
		return nil, fmt.Errorf("%w: %w", errInvalidSettingsResponse, err)
	}
	if settings == nil {
		return nil, errInvalidSettingsResponse
	}
	return settings, nil
}
