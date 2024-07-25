package graphCloudPcProvisioningPolicy

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/client"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

var _ resource.Resource = &CloudPcProvisioningPolicyResource{}
var _ resource.ResourceWithConfigure = &CloudPcProvisioningPolicyResource{}
var _ resource.ResourceWithImportState = &CloudPcProvisioningPolicyResource{}

func NewCloudPcProvisioningPolicyResource() resource.Resource {
	return &CloudPcProvisioningPolicyResource{}
}

type CloudPcProvisioningPolicyResource struct {
	client           *msgraphsdk.GraphServiceClient
	ProviderTypeName string
	TypeName         string
}

type CloudPcProvisioningPolicyResourceModel struct {
	ID                       types.String                   `tfsdk:"id"`
	AlternateResourceUrl     types.String                   `tfsdk:"alternate_resource_url"`
	CloudPcGroupDisplayName  types.String                   `tfsdk:"cloud_pc_group_display_name"`
	CloudPcNamingTemplate    types.String                   `tfsdk:"cloud_pc_naming_template"`
	Description              types.String                   `tfsdk:"description"`
	DisplayName              types.String                   `tfsdk:"display_name"`
	DomainJoinConfigurations []DomainJoinConfigurationModel `tfsdk:"domain_join_configurations"`
	EnableSingleSignOn       types.Bool                     `tfsdk:"enable_single_sign_on"`
	GracePeriodInHours       types.Int64                    `tfsdk:"grace_period_in_hours"`
	ImageDisplayName         types.String                   `tfsdk:"image_display_name"`
	ImageId                  types.String                   `tfsdk:"image_id"`
	ImageType                types.String                   `tfsdk:"image_type"`
	LocalAdminEnabled        types.Bool                     `tfsdk:"local_admin_enabled"`
	MicrosoftManagedDesktop  *MicrosoftManagedDesktopModel  `tfsdk:"microsoft_managed_desktop"`
	ProvisioningType         types.String                   `tfsdk:"provisioning_type"`
	WindowsSetting           *WindowsSettingModel           `tfsdk:"windows_setting"`
	Timeouts                 timeouts.Value                 `tfsdk:"timeouts"`
}

type DomainJoinConfigurationModel struct {
	DomainJoinType         types.String `tfsdk:"domain_join_type"`
	OnPremisesConnectionId types.String `tfsdk:"on_premises_connection_id"`
	RegionName             types.String `tfsdk:"region_name"`
}

type MicrosoftManagedDesktopModel struct {
	ManagedType types.String `tfsdk:"managed_type"`
	Profile     types.String `tfsdk:"profile"`
}

type WindowsSettingModel struct {
	Locale types.String `tfsdk:"locale"`
}

// GetID returns the ID of a resource from the state model.
func (s *CloudPcProvisioningPolicyResourceModel) GetID() string {
	return s.ID.ValueString()
}

// GetTypeName returns the type name of the resource from the state model.
func (r *CloudPcProvisioningPolicyResource) GetTypeName() string {
	return r.TypeName
}

// Metadata returns the resource type name.
func (r *CloudPcProvisioningPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_graph_cloud_pc_provisioning_policy"
}

// Configure sets the client for the resource.
func (r *CloudPcProvisioningPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Configuring CloudPcProvisioningPolicyResource")

	if req.ProviderData == nil {
		tflog.Warn(ctx, "Provider data is nil, skipping resource configuration")
		return
	}

	clients, ok := req.ProviderData.(*client.GraphClients)
	if !ok {
		tflog.Error(ctx, "Unexpected Provider Data Type", map[string]interface{}{
			"expected": "*client.GraphClients",
			"actual":   fmt.Sprintf("%T", req.ProviderData),
		})
		resp.Diagnostics.AddError(
			"Unexpected Provider Data Type",
			fmt.Sprintf("Expected *client.GraphClients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	if clients.StableClient == nil {
		tflog.Warn(ctx, "StableClient is nil, resource may not be fully configured")
		return
	}

	r.client = clients.StableClient
	tflog.Debug(ctx, "Initialized graphCloudPcProvisioningPolicy resource with Graph Client")
}

// ImportState imports the resource state.
func (r *CloudPcProvisioningPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Schema returns the schema for the resource.
func (r *CloudPcProvisioningPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the provisioning policy.",
			},
			"alternate_resource_url": schema.StringAttribute{
				Computed:    true,
				Description: "The URL of the alternate resource that links to this provisioning policy. Read-only.",
			},
			"cloud_pc_group_display_name": schema.StringAttribute{
				Computed:    true,
				Description: "The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.",
			},
			"cloud_pc_naming_template": schema.StringAttribute{
				Required: true,
				Description: "The template used to name Cloud PCs provisioned using this policy. The template can contain custom text and replacement tokens, including %USERNAME:x% and %RAND:x%, which represent the user's name and a randomly generated number, respectively. " +
					"For example, CPC-%USERNAME:4%-%RAND:5% means that the name of the Cloud PC starts with CPC-, followed by a four-character username, a - character, and then five random characters. The total length of the text generated by the template can't exceed 15 characters. Supports $filter, $select, and $orderby.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The provisioning policy description. Supports $filter, $select, and $orderBy.",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The display name for the provisioning policy.",
			},
			"domain_join_configurations": schema.ListNestedAttribute{
				Optional:    true,
				Description: "Specifies a list ordered by priority on how Cloud PCs join Microsoft Entra ID (Azure AD).",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"domain_join_type": schema.StringAttribute{
							Required:    true,
							Description: "Specifies the method by which the provisioned Cloud PC joins Microsoft Entra ID.",
							Validators: []validator.String{
								stringvalidator.OneOf("azureADJoin", "hybridAzureADJoin", "unknownFutureValue"),
							},
						},
						"on_premises_connection_id": schema.StringAttribute{
							Optional:    true,
							Description: "The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs.",
						},
						"region_name": schema.StringAttribute{
							Optional:    true,
							Description: "The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs.",
						},
					},
				},
			},
			"enable_single_sign_on": schema.BoolAttribute{
				Optional:    true,
				Description: "True if the provisioned Cloud PC can be accessed by single sign-on. False indicates that the provisioned Cloud PC doesn't support this feature. The default value is false. Supports $filter, $select, and $orderby.",
			},
			"grace_period_in_hours": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.",
			},
			"image_display_name": schema.StringAttribute{
				Computed:    true,
				Description: "The display name of the operating system image that is used for provisioning. Supports $filter, $select, and $orderBy.",
			},
			"image_id": schema.StringAttribute{
				Required: true,
				Description: "The unique identifier that represents an operating system image used for provisioning new Cloud PCs. " +
					"The format for a gallery type image is: {publisherName_offerName_skuName}. " +
					"Supported values: " +
					"publisher: 'Microsoftwindowsdesktop', " +
					"offer: 'windows-ent-cpc', " +
					"sku: '21h1-ent-cpc-m365', '21h1-ent-cpc-os', '20h2-ent-cpc-m365', '20h2-ent-cpc-os', " +
					"'20h1-ent-cpc-m365', '20h1-ent-cpc-os', '19h2-ent-cpc-m365', '19h2-ent-cpc-os'. " +
					"Supports $filter, $select, and $orderBy.",
			},
			"image_type": schema.StringAttribute{
				Required:    true,
				Description: "The type of operating system image (custom or gallery) that is used for provisioning on Cloud PCs. Possible values are: gallery, custom. The default value is gallery. Supports $filter, $select, and $orderBy.",
				Validators: []validator.String{
					stringvalidator.OneOf("gallery", "custom"),
				},
			},
			"local_admin_enabled": schema.BoolAttribute{
				Optional:    true,
				Description: "When true, the local admin is enabled for Cloud PCs; false indicates that the local admin isn't enabled for Cloud PCs. The default value is false. Supports $filter, $select, and $orderBy.",
			},
			"microsoft_managed_desktop": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "The specific settings for Microsoft Managed Desktop that enables Microsoft Managed Desktop customers to get device managed experience for Cloud PC.",
				Attributes: map[string]schema.Attribute{
					"managed_type": schema.StringAttribute{
						Optional:    true,
						Description: "Indicates the provisioning policy associated with Microsoft Managed Desktop settings.",
						Validators: []validator.String{
							stringvalidator.OneOf("notManaged", "premiumManaged", "standardManaged", "starterManaged", "unknownFutureValue"),
						},
					},
					"profile": schema.StringAttribute{
						Optional:    true,
						Description: "The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.",
					},
				},
			},
			"provisioning_type": schema.StringAttribute{
				Required:    true,
				Description: "Specifies the type of license used when provisioning Cloud PCs using this policy. By default, the license type is dedicated if the provisioningType isn't specified when you create the cloudPcProvisioningPolicy. Possible values are: dedicated, shared, unknownFutureValue. Supports $filter, $select, and $orderBy.",
				Validators: []validator.String{
					stringvalidator.OneOf("dedicated", "shared", "unknownFutureValue"),
				},
			},
			"windows_setting": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Indicates a specific Windows setting to configure during the creation of Cloud PCs for this provisioning policy.",
				Attributes: map[string]schema.Attribute{
					"locale": schema.StringAttribute{
						Optional:    true,
						Description: "The Windows language or region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).",
					},
				},
			},
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Create: true,
				Read:   true,
				Update: true,
				Delete: true,
			}),
		},
	}
}