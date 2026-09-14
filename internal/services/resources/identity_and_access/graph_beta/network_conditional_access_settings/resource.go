package graphBetaNetworkConditionalAccessSettings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/client"
	commonschema "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/schema"
)

const (
	ResourceName  = "microsoft365_graph_beta_identity_and_access_network_conditional_access_settings"
	CreateTimeout = 180
	ReadTimeout   = 180
	UpdateTimeout = 180
	DeleteTimeout = 180
	singletonID   = "conditionalAccess"
)

var (
	_ resource.Resource                = &NetworkConditionalAccessSettingsResource{}
	_ resource.ResourceWithConfigure   = &NetworkConditionalAccessSettingsResource{}
	_ resource.ResourceWithImportState = &NetworkConditionalAccessSettingsResource{}
	_ resource.ResourceWithIdentity    = &NetworkConditionalAccessSettingsResource{}
)

// NetworkConditionalAccessSettingsResource manages the tenant's Global Secure Access signaling settings.
type NetworkConditionalAccessSettingsResource struct {
	client           *msgraphbetasdk.GraphServiceClient
	ReadPermissions  []string
	WritePermissions []string
	ResourcePath     string
}

// NewNetworkConditionalAccessSettingsResource returns the singleton resource.
func NewNetworkConditionalAccessSettingsResource() resource.Resource {
	return &NetworkConditionalAccessSettingsResource{
		ReadPermissions:  []string{"NetworkAccess.Read.All"},
		WritePermissions: []string{"NetworkAccess.ReadWrite.All"},
		ResourcePath:     "/networkAccess/settings/conditionalAccess",
	}
}

func (r *NetworkConditionalAccessSettingsResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = ResourceName
}

func (r *NetworkConditionalAccessSettingsResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	r.client = client.SetGraphBetaClientForResource(ctx, req, resp, ResourceName)
}

func (r *NetworkConditionalAccessSettingsResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	if req.ID == "" && req.Identity != nil {
		resp.Diagnostics.Append(req.Identity.GetAttribute(ctx, path.Root("id"), &req.ID)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	if req.ID != singletonID {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			"Expected `conditionalAccess`, the fixed ID for this tenant-wide singleton.",
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetworkConditionalAccessSettingsResource) IdentitySchema(
	ctx context.Context,
	req resource.IdentitySchemaRequest,
	resp *resource.IdentitySchemaResponse,
) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{
		"id": identityschema.StringAttribute{RequiredForImport: true},
	}}
}

func (r *NetworkConditionalAccessSettingsResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages Microsoft Entra Global Secure Access conditional access settings using the Microsoft Graph beta `/networkAccess/settings/conditionalAccess` endpoint. This corresponds to the **Adaptive access** tab under **Session management** in the Entra admin center, and is separate from conditional access policies at `/identity/conditionalAccess/policies`.\n\n" +
			"This is a **tenant-wide singleton** with no documented create or delete operation. Initial apply updates the existing settings with PATCH when `signaling_status` is configured; otherwise it adopts the current settings with GET only. Destroy removes the resource from Terraform state only: it does not delete, disable, or reset the settings. Manage only one instance per tenant across all Terraform states.\n\n" +
			"Changing signaling can affect source IP restoration, compliant network validation, Conditional Access, Continuous Access Evaluation (CAE), and Identity Protection. Review the impact before changing it. The portal may also manage a compliant network named location; this resource manages only the settings endpoint.\n\n" +
			"See the [Microsoft Graph specification](https://learn.microsoft.com/en-us/graph/api/resources/networkaccess-conditionalaccesssettings?view=graph-rest-beta). Read requires `NetworkAccess.Read.All` or `NetworkAccess.ReadWrite.All`; write requires `NetworkAccess.ReadWrite.All`. Available in the Global service cloud only.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Fixed Terraform identifier: `conditionalAccess`. The singleton is addressed by its fixed API path rather than a server-generated ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"signaling_status": schema.StringAttribute{
				MarkdownDescription: "Enables or disables Global Secure Access conditional access signaling for Microsoft Entra ID. Values: `enabled`, `disabled`. **No default is applied.** When omitted, the current remote value is read and preserved, and this property is not sent in PATCH requests. Omission after import or after an explicit value relinquishes enforcement without resetting the setting.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"timeouts": commonschema.ResourceTimeouts(ctx),
		},
	}
}
