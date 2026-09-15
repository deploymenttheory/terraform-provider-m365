package graphBetaWindowsTrustedRootCertificate

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/client"
	commonschema "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/schema"
	assignmentschema "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/schema/graph_beta/device_management"
)

const (
	ResourceName  = "microsoft365_graph_beta_device_management_windows_trusted_root_certificate"
	CreateTimeout = 180
	ReadTimeout   = 180
	UpdateTimeout = 180
	DeleteTimeout = 180
)

var (
	_ resource.Resource                   = &WindowsTrustedRootCertificateResource{}
	_ resource.ResourceWithConfigure      = &WindowsTrustedRootCertificateResource{}
	_ resource.ResourceWithImportState    = &WindowsTrustedRootCertificateResource{}
	_ resource.ResourceWithValidateConfig = &WindowsTrustedRootCertificateResource{}
)

type WindowsTrustedRootCertificateResource struct {
	client           *msgraphbetasdk.GraphServiceClient
	ReadPermissions  []string
	WritePermissions []string
	ResourcePath     string
}

func NewWindowsTrustedRootCertificateResource() resource.Resource {
	return &WindowsTrustedRootCertificateResource{
		ReadPermissions:  []string{"DeviceManagementConfiguration.Read.All"},
		WritePermissions: []string{"DeviceManagementConfiguration.ReadWrite.All"},
		ResourcePath:     "/deviceManagement/deviceConfigurations",
	}
}

func (r *WindowsTrustedRootCertificateResource) Metadata(
	_ context.Context,
	_ resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = ResourceName
}

func (r *WindowsTrustedRootCertificateResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	r.client = client.SetGraphBetaClientForResource(ctx, req, resp, ResourceName)
}

func (r *WindowsTrustedRootCertificateResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *WindowsTrustedRootCertificateResource) Schema(
	ctx context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	assignments := assignmentschema.DeviceConfigurationWithAllGroupAssignmentsAndFilterSchema()
	assignments.MarkdownDescription = "Assignments for this certificate profile. Remove all entries to unassign the profile."
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Windows trusted certificate profile in Microsoft Intune through " +
			"`/deviceManagement/deviceConfigurations` and the `#microsoft.graph.windows81TrustedRootCertificate` type. " +
			"The Graph type retains its historical name for Windows 10 and later profiles. " +
			"Upload a public root or intermediate certificate; no private key or certificate connector is required. " +
			"See [trusted certificate profiles](https://learn.microsoft.com/en-us/intune/intune-service/protect/certificates-trusted-root).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The Intune device configuration ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Required: true, MarkdownDescription: "The profile display name.",
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"description": schema.StringAttribute{
				Optional: true, Computed: true, MarkdownDescription: "The profile description.",
				Validators: []validator.String{stringvalidator.LengthAtMost(1500)},
			},
			"cert_file_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The public certificate file name shown in Intune.",
				Validators:          []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"trusted_root_certificate": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Base64-encoded DER X.509 certificate. Use `filebase64(\"root.cer\")` with a DER certificate. Do not provide a private key or a PFX file.",
			},
			"destination_store": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The Windows certificate store that receives the certificate.",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"computerCertStoreRoot",
						"computerCertStoreIntermediate",
						"userCertStoreIntermediate",
					),
				},
			},
			"role_scope_tag_ids": schema.SetAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Intune scope tag IDs. Intune supplies the default scope tag when omitted.",
			},
			"assignments": assignments,
			"timeouts":    commonschema.ResourceTimeouts(ctx),
		},
	}
}

func (r *WindowsTrustedRootCertificateResource) ValidateConfig(
	ctx context.Context,
	req resource.ValidateConfigRequest,
	resp *resource.ValidateConfigResponse,
) {
	var data WindowsTrustedRootCertificateResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !data.TrustedRootCertificate.IsNull() && !data.TrustedRootCertificate.IsUnknown() {
		if _, err := decodeCertificate(data.TrustedRootCertificate.ValueString()); err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("trusted_root_certificate"),
				"Invalid trusted certificate",
				err.Error(),
			)
		}
	}
}

func decodeCertificate(value string) ([]byte, error) {
	der, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("decode certificate base64: %w", err)
	}
	if _, err := x509.ParseCertificate(der); err != nil {
		return nil, fmt.Errorf("parse DER X.509 certificate: %w", err)
	}
	return der, nil
}
