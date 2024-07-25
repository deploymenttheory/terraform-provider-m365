package graphBetaAssignmentFilter

import (
	"context"
	"fmt"
	"time"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/resources/common"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Create handles the Create operation.
func (r *AssignmentFilterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AssignmentFilterResourceModel

	tflog.Debug(ctx, fmt.Sprintf("Starting creation of resource: %s_%s", r.ProviderTypeName, r.TypeName))

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createTimeout, diags := plan.Timeouts.Create(ctx, 30*time.Second)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	requestBody, err := constructResource(ctx, &plan)
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

	plan.ID = types.StringValue(*assignmentFilter.GetId())

	mapRemoteStateToTerraform(ctx, &plan, assignmentFilter)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	tflog.Debug(ctx, fmt.Sprintf("Finished creation of resource: %s_%s", r.ProviderTypeName, r.TypeName))
}

func (r *AssignmentFilterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AssignmentFilterResourceModel

	tflog.Debug(ctx, "Starting Read method for assignment filter")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.IsNull() || state.ID.ValueString() == "" {
		resp.Diagnostics.AddWarning(
			"Unable to read assignment filter",
			"Assignment filter ID is empty or null. Unable to read assignment filter.",
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Reading assignment filter with ID: %s", state.ID.ValueString()))

	readTimeout, diags := state.Timeouts.Read(ctx, 30*time.Second)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()

	assignmentFilter, err := r.client.DeviceManagement().AssignmentFilters().ByDeviceAndAppManagementAssignmentFilterId(state.ID.ValueString()).Get(ctx, nil)
	if err != nil {
		if common.IsNotFoundError(err) {
			resp.Diagnostics.AddWarning(
				"Assignment filter not found",
				fmt.Sprintf("Assignment filter with ID %s was not found. Removing from state.", state.ID.ValueString()),
			)
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading assignment filter",
			fmt.Sprintf("Could not read assignment filter with ID %s: %s", state.ID.ValueString(), err.Error()),
		)
		return
	}

	mapRemoteStateToTerraform(ctx, &state, assignmentFilter)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

	tflog.Debug(ctx, "Finished Read method for assignment filter")
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
		if common.IsNotFoundError(err) && !r.isCreate {
			resp.Diagnostics.AddWarning(
				"Resource Not Found",
				fmt.Sprintf("The resource: %s_%s with ID %s was not found and will be removed from the state.", r.ProviderTypeName, r.TypeName, data.ID.ValueString()),
			)
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading assignment filter",
			fmt.Sprintf("Could not update resource: %s_%s: %s", r.ProviderTypeName, r.TypeName, err.Error()),
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
