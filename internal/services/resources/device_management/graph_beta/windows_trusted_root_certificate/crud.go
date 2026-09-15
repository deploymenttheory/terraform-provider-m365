package graphBetaWindowsTrustedRootCertificate

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/constants"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/convert"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/crud"
	kiotaerrors "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/errors/kiota"
)

func (r *WindowsTrustedRootCertificateResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var data WindowsTrustedRootCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := crud.HandleTimeout(
		ctx,
		data.Timeouts.Create,
		CreateTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()
	request, err := constructResource(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Invalid certificate profile", err.Error())
		return
	}
	assignments, err := constructAssignment(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Invalid certificate assignments", err.Error())
		return
	}
	remote, err := r.client.DeviceManagement().DeviceConfigurations().Post(ctx, request, nil)
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
	data.ID = types.StringPointerValue(remote.GetId())
	if data.Description.IsUnknown() {
		data.Description = convert.GraphToFrameworkString(remote.GetDescription())
	}
	if data.RoleScopeTagIds.IsUnknown() {
		data.RoleScopeTagIds = convert.GraphToFrameworkStringSet(ctx, remote.GetRoleScopeTagIds())
	}
	// Preserve the remote ID before the second API write so an assignment failure cannot orphan the profile.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if len(assignments.GetAssignments()) > 0 {
		_, err = r.client.DeviceManagement().
			DeviceConfigurations().
			ByDeviceConfigurationId(data.ID.ValueString()).
			Assign().
			Post(ctx, assignments, nil)
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
	}
	opts := crud.DefaultReadWithRetryOptions()
	opts.Operation = constants.TfOperationCreate
	opts.ResourceTypeName = ResourceName
	if err := crud.ReadWithRetry(
		ctx,
		r.Read,
		resource.ReadRequest{State: resp.State},
		&crud.CreateResponseContainer{CreateResponse: resp},
		opts,
	); err != nil {
		resp.Diagnostics.AddError("Cannot read created certificate profile", err.Error())
	}
}

func (r *WindowsTrustedRootCertificateResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var data WindowsTrustedRootCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := crud.HandleTimeout(
		ctx,
		data.Timeouts.Read,
		ReadTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()
	tflog.Debug(
		ctx,
		"Read Windows trusted certificate profile",
		map[string]any{"id": data.ID.ValueString()},
	)
	remote, err := r.client.DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(data.ID.ValueString()).
		Get(ctx,
			&devicemanagement.DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration{
				QueryParameters: &devicemanagement.DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters{
					Expand: []string{"assignments"},
				},
			})
	if err != nil {
		operation := constants.TfOperationRead
		if retryOperation, ok := ctx.Value("retry_operation").(string); ok {
			operation = retryOperation
		}
		kiotaerrors.HandleKiotaGraphError(ctx, err, resp, operation, r.ReadPermissions)
		return
	}
	resp.Diagnostics.Append(mapResource(ctx, &data, remote)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WindowsTrustedRootCertificateResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var data WindowsTrustedRootCertificateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := crud.HandleTimeout(
		ctx,
		data.Timeouts.Update,
		UpdateTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()
	request, err := constructResource(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Invalid certificate profile", err.Error())
		return
	}
	assignments, err := constructAssignment(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Invalid certificate assignments", err.Error())
		return
	}
	_, err = r.client.DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(data.ID.ValueString()).
		Patch(ctx, request, nil)
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
	_, err = r.client.DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(data.ID.ValueString()).
		Assign().
		Post(ctx, assignments, nil)
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	opts := crud.DefaultReadWithRetryOptions()
	opts.Operation = constants.TfOperationUpdate
	opts.ResourceTypeName = ResourceName
	if err := crud.ReadWithRetry(
		ctx,
		r.Read,
		resource.ReadRequest{State: resp.State},
		&crud.UpdateResponseContainer{UpdateResponse: resp},
		opts,
	); err != nil {
		resp.Diagnostics.AddError("Cannot read updated certificate profile", err.Error())
	}
}

func (r *WindowsTrustedRootCertificateResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var data WindowsTrustedRootCertificateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	ctx, cancel := crud.HandleTimeout(
		ctx,
		data.Timeouts.Delete,
		DeleteTimeout*time.Second,
		&resp.Diagnostics,
	)
	if cancel == nil {
		return
	}
	defer cancel()
	if err := r.client.DeviceManagement().
		DeviceConfigurations().
		ByDeviceConfigurationId(data.ID.ValueString()).
		Delete(ctx, nil); err != nil {
		kiotaerrors.HandleKiotaGraphError(
			ctx,
			err,
			resp,
			constants.TfOperationDelete,
			r.WritePermissions,
		)
		return
	}
	resp.State.RemoveResource(ctx)
}
