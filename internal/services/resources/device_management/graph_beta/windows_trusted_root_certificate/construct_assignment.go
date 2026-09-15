package graphBetaWindowsTrustedRootCertificate

import (
	"context"
	"errors"
	"fmt"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement"
	graphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/convert"
	sharedmodels "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/shared_models/graph_beta/device_management"
)

const noFilterID = "00000000-0000-0000-0000-000000000000"

var errInvalidAssignment = errors.New("invalid assignment")

func constructAssignment(
	ctx context.Context,
	data *WindowsTrustedRootCertificateResourceModel,
) (*devicemanagement.DeviceConfigurationsItemAssignPostRequestBody, error) {
	body := devicemanagement.NewDeviceConfigurationsItemAssignPostRequestBody()
	assignments := make([]graphmodels.DeviceConfigurationAssignmentable, 0)
	if data.Assignments.IsUnknown() {
		return nil, fmt.Errorf("%w: assignments must be known before apply", errInvalidAssignment)
	}
	if !data.Assignments.IsNull() {
		var values []sharedmodels.DeviceManagementDeviceConfigurationAssignmentWithGroupFilterModel
		if diags := data.Assignments.ElementsAs(ctx, &values, false); diags.HasError() {
			return nil, fmt.Errorf("%w: read assignments: %v", errInvalidAssignment, diags.Errors())
		}
		for i, value := range values {
			target, err := constructTarget(value)
			if err != nil {
				return nil, fmt.Errorf("assignment %d: %w", i, err)
			}
			assignment := graphmodels.NewDeviceConfigurationAssignment()
			assignment.SetTarget(target)
			assignments = append(assignments, assignment)
		}
	}
	body.SetAssignments(assignments)
	return body, nil
}

func constructTarget(
	value sharedmodels.DeviceManagementDeviceConfigurationAssignmentWithGroupFilterModel,
) (graphmodels.DeviceAndAppManagementAssignmentTargetable, error) {
	var target graphmodels.DeviceAndAppManagementAssignmentTargetable
	switch value.Type.ValueString() {
	case "allDevicesAssignmentTarget":
		target = graphmodels.NewAllDevicesAssignmentTarget()
	case "allLicensedUsersAssignmentTarget":
		target = graphmodels.NewAllLicensedUsersAssignmentTarget()
	case "groupAssignmentTarget":
		group := graphmodels.NewGroupAssignmentTarget()
		convert.FrameworkToGraphString(value.GroupId, group.SetGroupId)
		target = group
	case "exclusionGroupAssignmentTarget":
		group := graphmodels.NewExclusionGroupAssignmentTarget()
		convert.FrameworkToGraphString(value.GroupId, group.SetGroupId)
		target = group
	default:
		return nil, fmt.Errorf(
			"%w: unsupported assignment target %q",
			errInvalidAssignment,
			value.Type.ValueString(),
		)
	}
	groupTarget := value.Type.ValueString() == "groupAssignmentTarget" ||
		value.Type.ValueString() == "exclusionGroupAssignmentTarget"
	if groupTarget && value.GroupId.ValueString() == "" {
		return nil, fmt.Errorf("%w: group targets require group_id", errInvalidAssignment)
	}
	if !groupTarget && !value.GroupId.IsNull() {
		return nil, fmt.Errorf("%w: group_id only applies to group targets", errInvalidAssignment)
	}
	filterType := value.FilterType.ValueString()
	filterID := value.FilterId.ValueString()
	if filterType == "include" || filterType == "exclude" {
		if filterID == "" || filterID == noFilterID {
			return nil, fmt.Errorf(
				"%w: filter_id is required for %s filters",
				errInvalidAssignment,
				filterType,
			)
		}
		if value.Type.ValueString() == "exclusionGroupAssignmentTarget" {
			return nil, fmt.Errorf(
				"%w: exclusion groups do not support assignment filters",
				errInvalidAssignment,
			)
		}
		convert.FrameworkToGraphString(
			value.FilterId,
			target.SetDeviceAndAppManagementAssignmentFilterId,
		)
		if err := convert.FrameworkToGraphEnum(
			value.FilterType,
			graphmodels.ParseDeviceAndAppManagementAssignmentFilterType,
			target.SetDeviceAndAppManagementAssignmentFilterType,
		); err != nil {
			return nil, fmt.Errorf("set assignment filter: %w", err)
		}
	} else if filterID != "" && filterID != noFilterID {
		return nil, fmt.Errorf(
			"%w: filter_id requires filter_type include or exclude",
			errInvalidAssignment,
		)
	}
	return target, nil
}
