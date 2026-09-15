// REF: https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosgeneraldeviceconfiguration?view=graph-rest-beta
// REF: https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosdevicefeaturesconfiguration?view=graph-rest-beta
package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// IosDeviceConfigurationTemplatesJsonResourceModel describes the resource data model.
//
// This resource covers the two iOS/iPadOS template types whose property surface is too large to
// type usefully: iosGeneralDeviceConfiguration (187 settable properties) and
// iosDeviceFeaturesConfiguration (a deeply nested home-screen layout tree). The settings tree is
// carried as raw JSON while the provider owns the envelope.
type IosDeviceConfigurationTemplatesJsonResourceModel struct {
	ID              types.String   `tfsdk:"id"`
	OdataType       types.String   `tfsdk:"odata_type"`
	DisplayName     types.String   `tfsdk:"display_name"`
	Description     types.String   `tfsdk:"description"`
	RoleScopeTagIds types.Set      `tfsdk:"role_scope_tag_ids"`
	SettingsJson    types.String   `tfsdk:"settings_json"`
	Assignments     types.Set      `tfsdk:"assignments"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}
