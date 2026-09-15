package graphBetaWindowsTrustedRootCertificate

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	graphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/convert"
)

var assignmentType = types.ObjectType{AttrTypes: map[string]attr.Type{
	"type": types.StringType, "group_id": types.StringType,
	"filter_id": types.StringType, "filter_type": types.StringType,
}}

func mapResource(
	ctx context.Context,
	data *WindowsTrustedRootCertificateResourceModel,
	remote graphmodels.DeviceConfigurationable,
) diag.Diagnostics {
	var diags diag.Diagnostics
	cert, ok := remote.(*graphmodels.Windows81TrustedRootCertificate)
	if !ok {
		diags.AddError(
			"Incorrect Intune policy type",
			fmt.Sprintf(
				"Expected windows81TrustedRootCertificate, received %T. Import a Windows trusted certificate profile ID.",
				remote,
			),
		)
		return diags
	}
	if cert.GetTrustedRootCertificate() == nil {
		diags.AddError(
			"Missing certificate data",
			"Intune does not return the trusted certificate. Retry the read; import requires certificate data.",
		)
		return diags
	}
	data.ID = convert.GraphToFrameworkString(cert.GetId())
	data.DisplayName = convert.GraphToFrameworkString(cert.GetDisplayName())
	data.Description = convert.GraphToFrameworkString(cert.GetDescription())
	data.CertFileName = convert.GraphToFrameworkString(cert.GetCertFileName())
	data.DestinationStore = convert.GraphToFrameworkEnum(cert.GetDestinationStore())
	data.TrustedRootCertificate = types.StringValue(
		base64.StdEncoding.EncodeToString(cert.GetTrustedRootCertificate()),
	)
	data.RoleScopeTagIds = convert.GraphToFrameworkStringSet(ctx, cert.GetRoleScopeTagIds())
	values := make([]attr.Value, 0, len(cert.GetAssignments()))
	for _, assignment := range cert.GetAssignments() {
		value, err := mapAssignment(assignment.GetTarget())
		if err != nil {
			diags.AddError("Cannot read certificate assignment", err.Error())
			return diags
		}
		values = append(values, value)
	}
	if len(values) == 0 && data.Assignments.IsNull() {
		data.Assignments = types.SetNull(assignmentType)
	} else {
		var setDiags diag.Diagnostics
		data.Assignments, setDiags = types.SetValue(assignmentType, values)
		diags.Append(setDiags...)
	}
	return diags
}

func mapAssignment(
	target graphmodels.DeviceAndAppManagementAssignmentTargetable,
) (attr.Value, error) {
	values := map[string]attr.Value{
		"group_id":    types.StringNull(),
		"filter_id":   types.StringValue(noFilterID),
		"filter_type": types.StringValue("none"),
	}
	switch target := target.(type) {
	case *graphmodels.AllDevicesAssignmentTarget:
		values["type"] = types.StringValue("allDevicesAssignmentTarget")
	case *graphmodels.AllLicensedUsersAssignmentTarget:
		values["type"] = types.StringValue("allLicensedUsersAssignmentTarget")
	case *graphmodels.ExclusionGroupAssignmentTarget:
		values["type"] = types.StringValue("exclusionGroupAssignmentTarget")
		values["group_id"] = convert.GraphToFrameworkString(target.GetGroupId())
	case *graphmodels.GroupAssignmentTarget:
		values["type"] = types.StringValue("groupAssignmentTarget")
		values["group_id"] = convert.GraphToFrameworkString(target.GetGroupId())
	default:
		return nil, fmt.Errorf("%w: unsupported assignment target %T", errInvalidAssignment, target)
	}
	if id := target.GetDeviceAndAppManagementAssignmentFilterId(); id != nil && *id != "" {
		values["filter_id"] = types.StringValue(*id)
	}
	if filter := target.GetDeviceAndAppManagementAssignmentFilterType(); filter != nil {
		values["filter_type"] = types.StringValue(filter.String())
	}
	value, diags := types.ObjectValue(assignmentType.AttrTypes, values)
	if diags.HasError() {
		return nil, fmt.Errorf("%w: map assignment: %v", errInvalidAssignment, diags.Errors())
	}
	return value, nil
}
