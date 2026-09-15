package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/client"
	planmodifiers "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/plan_modifiers"
	commonschema "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/schema"
	commonschemagraphbeta "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/schema/graph_beta/device_management"
	customValidator "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/validate/attribute"
)

const (
	ResourceName  = "microsoft365_graph_beta_device_management_ios_device_configuration_templates_json"
	CreateTimeout = 180
	UpdateTimeout = 180
	ReadTimeout   = 180
	DeleteTimeout = 180

	// The custom_requests helpers are inconsistent about the separator between base URL and endpoint,
	// so two spellings are needed. Getting this wrong yields a URL like
	// "https://graph.microsoft.com/betadeviceManagement/..." rather than an obvious error.
	//
	//   PostRequest builds "{+baseurl}/" + Endpoint  -> endpoint must NOT have a leading slash.
	//   ByIDRequestUrlTemplate (used by GetRequestByResourceId and PatchRequestByResourceId) builds
	//   "{+baseurl}" + Endpoint                      -> endpoint MUST have a leading slash.
	CollectionEndpointPath = "deviceManagement/deviceConfigurations"
	ItemEndpointPath       = "/deviceManagement/deviceConfigurations"
)

var (
	// Basic resource interface (CRUD operations)
	_ resource.Resource = &IosDeviceConfigurationTemplatesJsonResource{}

	// Allows the resource to be configured with the provider client
	_ resource.ResourceWithConfigure = &IosDeviceConfigurationTemplatesJsonResource{}

	// Enables import functionality
	_ resource.ResourceWithImportState = &IosDeviceConfigurationTemplatesJsonResource{}

	// Enables identity schema for list resource support
	_ resource.ResourceWithIdentity = &IosDeviceConfigurationTemplatesJsonResource{}
)

func NewIosDeviceConfigurationTemplatesJsonResource() resource.Resource {
	return &IosDeviceConfigurationTemplatesJsonResource{
		ReadPermissions: []string{
			"DeviceManagementConfiguration.Read.All",
		},
		WritePermissions: []string{
			"DeviceManagementConfiguration.ReadWrite.All",
		},
		ResourcePath: "/deviceManagement/deviceConfigurations",
	}
}

type IosDeviceConfigurationTemplatesJsonResource struct {
	client           *msgraphbetasdk.GraphServiceClient
	ReadPermissions  []string
	WritePermissions []string
	ResourcePath     string
}

// Metadata returns the resource type name.
func (r *IosDeviceConfigurationTemplatesJsonResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = ResourceName
}

// Configure sets the client for the resource.
func (r *IosDeviceConfigurationTemplatesJsonResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	r.client = client.SetGraphBetaClientForResource(ctx, req, resp, ResourceName)
}

// ImportState imports the resource state.
func (r *IosDeviceConfigurationTemplatesJsonResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// IdentitySchema defines the identity schema for this resource, used by list operations to uniquely identify instances
func (r *IosDeviceConfigurationTemplatesJsonResource) IdentitySchema(
	ctx context.Context,
	req resource.IdentitySchemaRequest,
	resp *resource.IdentitySchemaResponse,
) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

// Schema defines the schema for the resource.
func (r *IosDeviceConfigurationTemplatesJsonResource) Schema(
	ctx context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages iOS/iPadOS device restriction and device features configuration templates in " +
			"Microsoft Intune using a raw JSON settings body. Use this resource for `iosGeneralDeviceConfiguration` " +
			"(device restrictions) and `iosDeviceFeaturesConfiguration` (home screen layout, single sign-on, web " +
			"content filters), whose property surfaces are too large to expose as typed attributes. For certificate, " +
			"Wi-Fi, VPN, email and custom profiles use " +
			"`microsoft365_graph_beta_device_management_ios_device_configuration_templates` instead.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the iOS/iPadOS configuration template.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"odata_type": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "The Graph type of the configuration profile. Possible values are: " +
					"`#microsoft.graph.iosGeneralDeviceConfiguration` (device restrictions), " +
					"`#microsoft.graph.iosDeviceFeaturesConfiguration` (device features), " +
					"`#microsoft.graph.iosTrustedRootCertificate` (trusted root certificate), " +
					"`#microsoft.graph.iosCustomConfiguration` (custom configuration), " +
					"`#microsoft.graph.iosWiFiConfiguration` (Wi-Fi), " +
					"`#microsoft.graph.iosScepCertificateProfile` (SCEP certificate), " +
					"`#microsoft.graph.iosPkcsCertificateProfile` (PKCS certificate), " +
					"`#microsoft.graph.iosEnterpriseWiFiConfiguration` (enterprise Wi-Fi), " +
					"`#microsoft.graph.iosEasEmailProfileConfiguration` (EAS email), " +
					"`#microsoft.graph.iosVpnConfiguration` (VPN). Changing this forces " +
					"replacement — Graph cannot mutate a profile's type in place.",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"#microsoft.graph.iosGeneralDeviceConfiguration",
						"#microsoft.graph.iosDeviceFeaturesConfiguration",
						"#microsoft.graph.iosTrustedRootCertificate",
						"#microsoft.graph.iosCustomConfiguration",
						"#microsoft.graph.iosWiFiConfiguration",
						"#microsoft.graph.iosScepCertificateProfile",
						"#microsoft.graph.iosPkcsCertificateProfile",
						"#microsoft.graph.iosEnterpriseWiFiConfiguration",
						"#microsoft.graph.iosEasEmailProfileConfiguration",
						"#microsoft.graph.iosVpnConfiguration",
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"display_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The display name for the iOS/iPadOS configuration template.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Optional description of the resource. Maximum length is 1500 characters.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1500),
				},
			},
			"role_scope_tag_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Set of scope tag IDs for this device configuration template.",
				PlanModifiers: []planmodifier.Set{
					planmodifiers.DefaultSetValue(
						[]attr.Value{types.StringValue("0")},
					),
				},
			},
			"settings_json": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "The iOS/iPadOS settings tree as a JSON string, typically written with " +
					"`jsonencode()`.\n\n" +
					"This must contain **only** the settings themselves. The provider owns the envelope and injects " +
					"`@odata.type`, `displayName`, `description` and `roleScopeTagIds` from the dedicated attributes; " +
					"including any of those four at the root of this value is an error.\n\n" +
					"**Only the properties you declare are managed.** Microsoft Graph returns every property of these " +
					"profile types on read — including ones never configured, reported as `false`, `null` or `[]` — so " +
					"the provider projects the response onto the shape of your configuration. A consequence is that " +
					"`false` and \"not set\" are indistinguishable on the wire: removing a property from this value " +
					"stops Terraform tracking it, it does **not** reset it on the server. Declaring a property is how " +
					"you begin managing it.\n\n" +
					"**Nested `@odata.type` discriminators are required.** Only the root level is checked for " +
					"envelope keys. Graph relies on nested discriminators to resolve polymorphic members — for " +
					"example each `homeScreenPages[].icons[]` entry must declare either " +
					"`#microsoft.graph.iosHomeScreenApp` or `#microsoft.graph.iosHomeScreenFolder`.\n\n" +
					"Note the two page types differ: a top-level `iosHomeScreenPage` uses `icons` and may hold " +
					"both apps and folders, while an `iosHomeScreenFolderPage` nested inside a folder uses " +
					"`apps` and may hold only apps. Using `icons` on a folder page is rejected by Graph.\n\n" +
					"A device features profile looks like this:\n\n" +
					"```hcl\n" +
					"settings_json = jsonencode({\n" +
					"  homeScreenPages = [\n" +
					"    {\n" +
					"      \"@odata.type\" = \"#microsoft.graph.iosHomeScreenPage\"\n" +
					"      displayName   = \"Page 1\"\n" +
					"      icons = [\n" +
					"        {\n" +
					"          \"@odata.type\" = \"#microsoft.graph.iosHomeScreenApp\"\n" +
					"          displayName   = \"Safari\"\n" +
					"          bundleID      = \"com.apple.mobilesafari\"\n" +
					"        }\n" +
					"      ]\n" +
					"    }\n" +
					"  ]\n" +
					"})\n" +
					"```\n\n" +
					"On import the full property surface is written to state, which will be large; prune it down to " +
					"the properties you intend to manage.",
				Validators: []validator.String{
					customValidator.JSONSchemaValidator(),
					settingsJSONEnvelope(),
				},
				PlanModifiers: []planmodifier.String{
					planmodifiers.NormalizeJSONPlanModifier{},
				},
			},
			"assignments": commonschemagraphbeta.DeviceConfigurationWithAllGroupAssignmentsAndFilterSchema(),
			"timeouts":    commonschema.ResourceTimeouts(ctx),
		},
	}
}
