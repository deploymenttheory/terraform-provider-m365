package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	graphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/constants"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/crud"
	customrequest "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/custom_requests"
	kiotaerrors "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/errors/kiota"
	sharedmodels "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/shared_models/graph_beta"
)

// deviceConfigurationErrorMapping lets Kiota parse a Graph error body into an ODataError so the
// diagnostic carries the server's actual message. Without it a non-2xx surfaces only as "The server
// returned an unexpected status code and no error factory is registered for this code", hiding the
// reason. "XXX" is Kiota's catch-all status pattern, following the convention in
// identity_and_access/graph_beta/network_web_filtering_policy/requests.go.
var deviceConfigurationErrorMapping = abstractions.ErrorMappings{
	"XXX": odataerrors.CreateODataErrorFromDiscriminatorValue,
}

// Create handles the Create operation.
//
// The profile itself is created through custom_requests so the raw settings tree reaches Graph
// unaltered; assignments then go through the typed SDK exactly as in the sibling typed resource.
func (r *IosDeviceConfigurationTemplatesJsonResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var object IosDeviceConfigurationTemplatesJsonResourceModel

	tflog.Debug(ctx, fmt.Sprintf("Starting creation of resource: %s", ResourceName))

	resp.Diagnostics.Append(req.Plan.Get(ctx, &object)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := crud.HandleTimeout(
		ctx,
		object.Timeouts.Create,
		CreateTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()

	requestBody, err := constructResource(ctx, &object)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error constructing resource",
			fmt.Sprintf("Could not construct resource: %s: %s", ResourceName, err.Error()),
		)
		return
	}

	createdResource, err := customrequest.PostRequest(
		ctx,
		r.client.GetAdapter(),
		customrequest.PostRequestConfig{
			APIVersion:  customrequest.GraphAPIBeta,
			Endpoint:    CollectionEndpointPath,
			RequestBody: requestBody,
		},
		graphmodels.CreateDeviceConfigurationFromDiscriminatorValue,
		deviceConfigurationErrorMapping,
	)
	if err != nil {
		kiotaerrors.HandleKiotaGraphError(
			ctx,
			err,
			resp,
			constants.TfOperationCreate,
			r.WritePermissions,
		)
		return
	}

	// The response is parsed as a DeviceConfiguration purely to recover the server-assigned id; the
	// settings themselves are read back as raw JSON in Read.
	deviceConfig, ok := createdResource.(graphmodels.DeviceConfigurationable)
	if !ok || deviceConfig.GetId() == nil {
		resp.Diagnostics.AddError(
			"Error creating resource",
			fmt.Sprintf("Could not determine the id of the created resource: %s", ResourceName),
		)
		return
	}

	object.ID = types.StringValue(*deviceConfig.GetId())

	requestAssignment, err := constructAssignment(ctx, &object)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error constructing assignment for create method",
			fmt.Sprintf("Could not construct assignment: %s: %s", ResourceName, err.Error()),
		)
		return
	}

	_, err = r.client.
		DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(object.ID.ValueString()).
		Assign().
		Post(ctx, requestAssignment, nil)
	if err != nil {
		kiotaerrors.HandleKiotaGraphError(
			ctx,
			err,
			resp,
			constants.TfOperationCreate,
			r.WritePermissions,
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &object)...)
	if resp.Diagnostics.HasError() {
		return
	}

	readReq := resource.ReadRequest{State: resp.State}
	stateContainer := &crud.CreateResponseContainer{CreateResponse: resp}

	opts := crud.DefaultReadWithRetryOptions()
	opts.Operation = constants.TfOperationCreate
	opts.ResourceTypeName = ResourceName

	err = crud.ReadWithRetry(ctx, r.Read, readReq, stateContainer, opts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading resource state after create",
			fmt.Sprintf("Could not read resource state: %s: %s", ResourceName, err.Error()),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Finished Create Method: %s", ResourceName))
}

// Read handles the Read operation.
//
// The profile body is fetched as raw JSON. Assignments are fetched separately through the typed SDK
// rather than with $expand: expanding them would fold the assignments collection into the same body
// that becomes settings_json.
func (r *IosDeviceConfigurationTemplatesJsonResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var object IosDeviceConfigurationTemplatesJsonResourceModel
	var identity sharedmodels.ResourceIdentity

	tflog.Debug(ctx, fmt.Sprintf("Starting Read method for: %s", ResourceName))

	operation := constants.TfOperationRead
	if ctxOp := ctx.Value("retry_operation"); ctxOp != nil {
		if opStr, ok := ctxOp.(string); ok {
			operation = opStr
		}
	}

	resp.Diagnostics.Append(req.State.Get(ctx, &object)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Reading %s with ID: %s", ResourceName, object.ID.ValueString()))

	ctx, cancel := crud.HandleTimeout(
		ctx,
		object.Timeouts.Read,
		ReadTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()

	identity.ID = object.ID.ValueString()

	if resp.Identity != nil {
		resp.Diagnostics.Append(resp.Identity.Set(ctx, identity)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	rawResponse, err := customrequest.GetRequestByResourceId(
		ctx,
		r.client.GetAdapter(),
		customrequest.GetRequestConfig{
			APIVersion:        customrequest.GraphAPIBeta,
			Endpoint:          ItemEndpointPath,
			ResourceIDPattern: "('id')",
			ResourceID:        object.ID.ValueString(),
		},
	)
	if err != nil {
		kiotaerrors.HandleKiotaGraphError(ctx, err, resp, operation, r.ReadPermissions)
		return
	}

	MapRemoteResourceStateToTerraform(ctx, &object, rawResponse)

	assignments, err := r.client.
		DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(object.ID.ValueString()).
		Assignments().
		Get(ctx, nil)
	if err != nil {
		kiotaerrors.HandleKiotaGraphError(ctx, err, resp, operation, r.ReadPermissions)
		return
	}

	if assignments == nil || len(assignments.GetValue()) == 0 {
		object.Assignments = types.SetNull(IosConfigurationTemplatesJsonAssignmentType())
	} else {
		mapAssignmentsToTerraform(ctx, &object, assignments.GetValue())
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &object)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Finished Read Method: %s", ResourceName))
}

// Update handles the Update operation.
//
// PATCH is used rather than PUT so that properties absent from settings_json are left untouched
// server-side, which matches the projection semantics on read: only declared properties are managed.
func (r *IosDeviceConfigurationTemplatesJsonResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan IosDeviceConfigurationTemplatesJsonResourceModel
	var state IosDeviceConfigurationTemplatesJsonResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating %s with ID: %s", ResourceName, state.ID.ValueString()))

	ctx, cancel := crud.HandleTimeout(
		ctx,
		plan.Timeouts.Update,
		UpdateTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()

	requestBody, err := constructResource(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error constructing resource",
			fmt.Sprintf("Could not construct resource: %s: %s", ResourceName, err.Error()),
		)
		return
	}

	err = customrequest.PatchRequestByResourceId(
		ctx,
		r.client.GetAdapter(),
		customrequest.PatchRequestConfig{
			APIVersion:        customrequest.GraphAPIBeta,
			Endpoint:          ItemEndpointPath,
			ResourceID:        state.ID.ValueString(),
			ResourceIDPattern: "('id')",
			RequestBody:       requestBody,
		},
	)
	if err != nil {
		kiotaerrors.HandleKiotaGraphError(
			ctx,
			err,
			resp,
			constants.TfOperationUpdate,
			r.WritePermissions,
		)
		return
	}

	// Always reassign, so that removing assignments from configuration removes them remotely.
	requestAssignment, err := constructAssignment(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error constructing assignment for update method",
			fmt.Sprintf("Could not construct assignment: %s: %s", ResourceName, err.Error()),
		)
		return
	}

	_, err = r.client.
		DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(state.ID.ValueString()).
		Assign().
		Post(ctx, requestAssignment, nil)
	if err != nil {
		kiotaerrors.HandleKiotaGraphError(
			ctx,
			err,
			resp,
			constants.TfOperationUpdate,
			r.WritePermissions,
		)
		return
	}

	// Carry the id forward: it is not part of the plan and would otherwise be unknown.
	plan.ID = state.ID

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	readReq := resource.ReadRequest{State: resp.State}
	stateContainer := &crud.UpdateResponseContainer{UpdateResponse: resp}

	opts := crud.DefaultReadWithRetryOptions()
	opts.Operation = constants.TfOperationUpdate
	opts.ResourceTypeName = ResourceName

	err = crud.ReadWithRetry(ctx, r.Read, readReq, stateContainer, opts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading resource state after update",
			fmt.Sprintf("Could not read resource state: %s: %s", ResourceName, err.Error()),
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Finished Update Method: %s", ResourceName))
}

// Delete handles the Delete operation. Assignments and settings are removed with the profile.
func (r *IosDeviceConfigurationTemplatesJsonResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var object IosDeviceConfigurationTemplatesJsonResourceModel

	tflog.Debug(ctx, fmt.Sprintf("Starting Delete of resource: %s", ResourceName))

	resp.Diagnostics.Append(req.State.Get(ctx, &object)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx, cancel := crud.HandleTimeout(
		ctx,
		object.Timeouts.Delete,
		DeleteTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()

	err := r.client.
		DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(object.ID.ValueString()).
		Delete(ctx, nil)
	if err != nil {
		kiotaerrors.HandleKiotaGraphError(
			ctx,
			err,
			resp,
			constants.TfOperationDelete,
			r.WritePermissions,
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Removing %s from Terraform state", ResourceName))

	resp.State.RemoveResource(ctx)

	tflog.Debug(ctx, fmt.Sprintf("Finished Delete Method: %s", ResourceName))
}
