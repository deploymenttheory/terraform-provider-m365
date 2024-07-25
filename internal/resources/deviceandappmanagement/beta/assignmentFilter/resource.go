// REF: https://learn.microsoft.com/en-us/graph/api/resources/intune-policyset-deviceandappmanagementassignmentfilter?view=graph-rest-beta
package graphBetaAssignmentFilter

import (
	"context"
	"fmt"
	"time"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/client"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
)

var _ resource.Resource = &AssignmentFilterResource{}
var _ resource.ResourceWithConfigure = &AssignmentFilterResource{}
var _ resource.ResourceWithImportState = &AssignmentFilterResource{}

func NewAssignmentFilterResource() resource.Resource {
	return &AssignmentFilterResource{}
}

type AssignmentFilterResource struct {
	client           *msgraphbetasdk.GraphServiceClient
	ProviderTypeName string
	TypeName         string
}

type AssignmentFilterResourceModel struct {
	ID                             types.String   `tfsdk:"id"`
	DisplayName                    types.String   `tfsdk:"display_name"`
	Description                    types.String   `tfsdk:"description"`
	Platform                       types.String   `tfsdk:"platform"`
	Rule                           types.String   `tfsdk:"rule"`
	AssignmentFilterManagementType types.String   `tfsdk:"assignment_filter_management_type"`
	CreatedDateTime                types.String   `tfsdk:"created_date_time"`
	LastModifiedDateTime           types.String   `tfsdk:"last_modified_date_time"`
	RoleScopeTags                  types.List     `tfsdk:"role_scope_tags"`
	Payloads                       types.List     `tfsdk:"payloads"`
	Timeouts                       timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *AssignmentFilterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_graph_beta_device_and_app_management_assignment_filter"
}

// Configure sets the client for the resource.
func (r *AssignmentFilterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Configuring AssignmentFilterResource")

	if req.ProviderData == nil {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we reject this. Please report this to the provider developers.",
		)
		return
	}

	clients, ok := req.ProviderData.(*client.GraphClients)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Provider Data Type",
			fmt.Sprintf("Expected *client.GraphClients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	if clients.BetaClient == nil {
		resp.Diagnostics.AddError(
			"BetaClient is nil",
			"The BetaClient in the provider data is nil. This could indicate a configuration error in the provider. Please check your provider configuration and try again.",
		)
		return
	}

	r.client = clients.BetaClient
	tflog.Debug(ctx, "Initialized graphBetaAssignmentFilter resource with BetaClient")
}

// ImportState imports the resource state.
func (r *AssignmentFilterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Schema returns the schema for the resource.
func (r *AssignmentFilterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the assignment filter.",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The display name of the assignment filter.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The optional description of the assignment filter.",
			},
			"platform": schema.StringAttribute{
				Required:    true,
				Description: fmt.Sprintf("The Intune device management type (platform) for the assignment filter. Supported types: %v", getAllPlatformStrings()),
				Validators: []validator.String{
					platformValidator{},
				},
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Rule definition of the assignment filter.",
			},
			"assignment_filter_management_type": schema.StringAttribute{
				Optional:    true,
				Description: fmt.Sprintf("Indicates filter is applied to either 'devices' or 'apps' management type. Possible values are: %v. Default filter will be applied to 'devices'.", getAllManagementTypeStrings()),
				Validators: []validator.String{
					assignmentFilterManagementTypeValidator{},
				},
			},

			"created_date_time": schema.StringAttribute{
				Computed:    true,
				Description: "The creation time of the assignment filter.",
			},
			"last_modified_date_time": schema.StringAttribute{
				Computed:    true,
				Description: "Last modified time of the assignment filter.",
			},
			"role_scope_tags": schema.ListAttribute{
				Optional:    true,
				Description: "Indicates role scope tags assigned for the assignment filter.",
				ElementType: types.StringType,
			},
			"payloads": schema.ListNestedAttribute{
				Optional:    true,
				Description: "Indicates associated assignments for a specific filter.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"payload_id": schema.StringAttribute{
							Required:    true,
							Description: "The ID of the payload.",
						},
						"payload_type": schema.StringAttribute{
							Required:    true,
							Description: "The type of the payload.",
						},
						"group_id": schema.StringAttribute{
							Required:    true,
							Description: "The group ID associated with the payload.",
						},
						"assignment_filter_type": schema.StringAttribute{
							Required:    true,
							Description: fmt.Sprintf("The assignment filter type. Supported types: %v", getAllAssignmentFilterTypes()),
							Validators: []validator.String{
								assignmentFilterTypeValidator{},
							},
						},
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

// Create handles the Create operation.
func (r *AssignmentFilterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AssignmentFilterResourceModel

	tflog.Debug(ctx, fmt.Sprintf("Starting creation of resource: %s_%s", r.ProviderTypeName, r.TypeName))

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.client == nil {
		resp.Diagnostics.AddError(
			"Client is not initialized",
			"Cannot create assignment filter because the client is not initialized.",
		)
		return
	}

	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Second)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	requestBody, err := constructResource(ctx, &data)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error constructing assignment filter",
			fmt.Sprintf("Could not construct resource: %s_%s: %s", r.ProviderTypeName, r.TypeName, err.Error()),
		)
		return
	}

	assignmentFilter, err := r.client.DeviceManagement().AssignmentFilters().Post(ctx, requestBody, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating assignment filter",
			fmt.Sprintf("Could not create assignment filter: %s", err.Error()),
		)
		return
	}

	data.ID = types.StringValue(*assignmentFilter.GetId())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Debug(ctx, fmt.Sprintf("Finished creation of resource: %s_%s", r.ProviderTypeName, r.TypeName))
}

// Read handles the read operation and stating.
func (r *AssignmentFilterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AssignmentFilterResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	readTimeout, diags := data.Timeouts.Read(ctx, 30*time.Second)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()

	remoteResource, err := r.client.DeviceManagement().AssignmentFilters().ByDeviceAndAppManagementAssignmentFilterId(data.ID.ValueString()).Get(ctx, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading assignment filter",
			fmt.Sprintf("Could not read assignment filter: %s", err.Error()),
		)
		return
	}

	mapRemoteStateToTerraform(&data, remoteResource)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update handles the Update operation.
func (r *AssignmentFilterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AssignmentFilterResourceModel

	tflog.Debug(ctx, fmt.Sprintf("Starting Update of resource: %s_%s", r.ProviderTypeName, r.TypeName))

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateTimeout, diags := data.Timeouts.Update(ctx, 30*time.Second)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, updateTimeout)
	defer cancel()

	requestBody, err := constructResource(ctx, &data)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error constructing assignment filter",
			fmt.Sprintf("Could not construct resource: %s_%s: %s", r.ProviderTypeName, r.TypeName, err.Error()),
		)
		return
	}

	_, err = r.client.DeviceManagement().AssignmentFilters().ByDeviceAndAppManagementAssignmentFilterId(data.ID.ValueString()).Patch(ctx, requestBody, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating assignment filter",
			fmt.Sprintf("Could not update assignment filter: %s", err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Debug(ctx, fmt.Sprintf("Finished Update of resource: %s_%s", r.ProviderTypeName, r.TypeName))
}

// Delete handles the Delete operation.
func (r *AssignmentFilterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AssignmentFilterResourceModel

	tflog.Debug(ctx, fmt.Sprintf("Starting deletion of resource: %s_%s", r.ProviderTypeName, r.TypeName))

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteTimeout, diags := data.Timeouts.Delete(ctx, 30*time.Second)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, deleteTimeout)
	defer cancel()

	err := r.client.DeviceManagement().AssignmentFilters().ByDeviceAndAppManagementAssignmentFilterId(data.ID.ValueString()).Delete(ctx, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Client error when deleting %s_%s", r.ProviderTypeName, r.TypeName), err.Error())
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Completed deletion of resource: %s_%s", r.ProviderTypeName, r.TypeName))

	resp.State.RemoveResource(ctx)
}
