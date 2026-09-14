package graphBetaNetworkConditionalAccessSettings

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NetworkConditionalAccessSettingsResourceModel contains the supported settings and Terraform metadata.
type NetworkConditionalAccessSettingsResourceModel struct {
	ID              types.String   `tfsdk:"id"`
	SignalingStatus types.String   `tfsdk:"signaling_status"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}

// conditionalAccessSettingsResponse decodes only the verified property; extra service fields are not replayed.
type conditionalAccessSettingsResponse struct {
	SignalingStatus *string `json:"signalingStatus"`
}
