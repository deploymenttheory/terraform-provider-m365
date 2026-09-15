package graphBetaIosDeviceConfigurationTemplates

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	graphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/convert"
)

// IosConfigurationTemplatesAssignmentType returns the object type for IosConfigurationTemplatesAssignmentModel
func IosConfigurationTemplatesAssignmentType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"type":        types.StringType,
			"group_id":    types.StringType,
			"filter_id":   types.StringType,
			"filter_type": types.StringType,
		},
	}
}

func mapAssignmentsToTerraform(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	assignments []graphmodels.DeviceConfigurationAssignmentable,
) {
	if len(assignments) == 0 {
		tflog.Debug(ctx, "No assignments to process")
		data.Assignments = types.SetNull(IosConfigurationTemplatesAssignmentType())
		return
	}

	tflog.Debug(ctx, "Starting assignment mapping process", map[string]any{
		"assignmentCount": len(assignments),
		"resourceId":      data.ID.ValueString(),
	})

	assignmentValues := []attr.Value{}

	for i, assignment := range assignments {
		tflog.Debug(ctx, "Processing assignment", map[string]any{
			"assignmentIndex": i,
			"assignmentId":    assignment.GetId(),
			"resourceId":      data.ID.ValueString(),
		})

		target := assignment.GetTarget()
		if target == nil {
			tflog.Warn(ctx, "Assignment target is nil, skipping assignment", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"resourceId":      data.ID.ValueString(),
			})
			continue
		}

		odataType := target.GetOdataType()
		if odataType == nil {
			tflog.Warn(
				ctx,
				"Assignment target OData type is nil, skipping assignment",
				map[string]any{
					"assignmentIndex": i,
					"assignmentId":    assignment.GetId(),
					"resourceId":      data.ID.ValueString(),
				},
			)
			continue
		}

		tflog.Debug(ctx, "Processing assignment target", map[string]any{
			"assignmentIndex": i,
			"assignmentId":    assignment.GetId(),
			"targetType":      *odataType,
			"resourceId":      data.ID.ValueString(),
		})

		assignmentObj := map[string]attr.Value{
			"type":        types.StringNull(),
			"group_id":    types.StringNull(),
			"filter_id":   types.StringNull(),
			"filter_type": types.StringNull(),
		}

		switch *odataType {
		case "#microsoft.graph.allDevicesAssignmentTarget":
			tflog.Debug(ctx, "Mapping allDevicesAssignmentTarget", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"resourceId":      data.ID.ValueString(),
			})
			assignmentObj["type"] = types.StringValue("allDevicesAssignmentTarget")
			assignmentObj["group_id"] = types.StringNull()

		case "#microsoft.graph.allLicensedUsersAssignmentTarget":
			tflog.Debug(ctx, "Mapping allLicensedUsersAssignmentTarget", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"resourceId":      data.ID.ValueString(),
			})
			assignmentObj["type"] = types.StringValue("allLicensedUsersAssignmentTarget")
			assignmentObj["group_id"] = types.StringNull()

		case "#microsoft.graph.groupAssignmentTarget":
			tflog.Debug(ctx, "Mapping groupAssignmentTarget", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"resourceId":      data.ID.ValueString(),
			})
			assignmentObj["type"] = types.StringValue("groupAssignmentTarget")

			if groupTarget, ok := target.(graphmodels.GroupAssignmentTargetable); ok {
				groupId := groupTarget.GetGroupId()
				if groupId != nil && *groupId != "" {
					tflog.Debug(ctx, "Setting group ID for group assignment target", map[string]any{
						"assignmentIndex": i,
						"assignmentId":    assignment.GetId(),
						"groupId":         *groupId,
						"resourceId":      data.ID.ValueString(),
					})
					assignmentObj["group_id"] = convert.GraphToFrameworkString(groupId)
				} else {
					tflog.Warn(
						ctx,
						"Group ID is nil/empty for group assignment target",
						map[string]any{
							"assignmentIndex": i,
							"assignmentId":    assignment.GetId(),
							"resourceId":      data.ID.ValueString(),
						},
					)
					assignmentObj["group_id"] = types.StringNull()
				}
			} else {
				tflog.Error(
					ctx,
					"Failed to cast target to GroupAssignmentTargetable",
					map[string]any{
						"assignmentIndex": i,
						"assignmentId":    assignment.GetId(),
						"resourceId":      data.ID.ValueString(),
					},
				)
				assignmentObj["group_id"] = types.StringNull()
			}

		case "#microsoft.graph.exclusionGroupAssignmentTarget":
			tflog.Debug(ctx, "Mapping exclusionGroupAssignmentTarget", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"resourceId":      data.ID.ValueString(),
			})
			assignmentObj["type"] = types.StringValue("exclusionGroupAssignmentTarget")

			if groupTarget, ok := target.(graphmodels.ExclusionGroupAssignmentTargetable); ok {
				groupId := groupTarget.GetGroupId()
				if groupId != nil && *groupId != "" {
					tflog.Debug(
						ctx,
						"Setting group ID for exclusion group assignment target",
						map[string]any{
							"assignmentIndex": i,
							"assignmentId":    assignment.GetId(),
							"groupId":         *groupId,
							"resourceId":      data.ID.ValueString(),
						},
					)
					assignmentObj["group_id"] = convert.GraphToFrameworkString(groupId)
				} else {
					tflog.Warn(
						ctx,
						"Group ID is nil/empty for exclusion group assignment target",
						map[string]any{
							"assignmentIndex": i,
							"assignmentId":    assignment.GetId(),
							"resourceId":      data.ID.ValueString(),
						},
					)
					assignmentObj["group_id"] = types.StringNull()
				}
			} else {
				tflog.Error(
					ctx,
					"Failed to cast target to ExclusionGroupAssignmentTargetable",
					map[string]any{
						"assignmentIndex": i,
						"assignmentId":    assignment.GetId(),
						"resourceId":      data.ID.ValueString(),
					},
				)
				assignmentObj["group_id"] = types.StringNull()
			}

		default:
			tflog.Warn(ctx, "Unknown target type encountered", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"targetType":      *odataType,
				"resourceId":      data.ID.ValueString(),
			})
			assignmentObj["group_id"] = types.StringNull()
		}

		tflog.Debug(ctx, "Processing assignment filters", map[string]any{
			"assignmentIndex": i,
			"assignmentId":    assignment.GetId(),
			"resourceId":      data.ID.ValueString(),
		})

		filterID := target.GetDeviceAndAppManagementAssignmentFilterId()
		if filterID != nil && *filterID != "" &&
			*filterID != "00000000-0000-0000-0000-000000000000" {
			tflog.Debug(ctx, "Assignment has meaningful filter ID", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"filterId":        *filterID,
				"resourceId":      data.ID.ValueString(),
			})
			assignmentObj["filter_id"] = convert.GraphToFrameworkString(filterID)
		} else {
			tflog.Debug(
				ctx,
				"Assignment has no meaningful filter ID, using schema default",
				map[string]any{
					"assignmentIndex": i,
					"assignmentId":    assignment.GetId(),
					"resourceId":      data.ID.ValueString(),
				},
			)
			assignmentObj["filter_id"] = types.StringValue("00000000-0000-0000-0000-000000000000")
		}

		filterType := target.GetDeviceAndAppManagementAssignmentFilterType()
		if filterType != nil {
			tflog.Debug(ctx, "Processing filter type", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"filterType":      *filterType,
				"resourceId":      data.ID.ValueString(),
			})

			switch *filterType {
			case graphmodels.INCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE:
				tflog.Debug(ctx, "Setting filter type to include", map[string]any{
					"assignmentIndex": i,
					"assignmentId":    assignment.GetId(),
					"resourceId":      data.ID.ValueString(),
				})
				assignmentObj["filter_type"] = types.StringValue("include")
			case graphmodels.EXCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE:
				tflog.Debug(ctx, "Setting filter type to exclude", map[string]any{
					"assignmentIndex": i,
					"assignmentId":    assignment.GetId(),
					"resourceId":      data.ID.ValueString(),
				})
				assignmentObj["filter_type"] = types.StringValue("exclude")
			case graphmodels.NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE:
				tflog.Debug(ctx, "Setting filter type to none", map[string]any{
					"assignmentIndex": i,
					"assignmentId":    assignment.GetId(),
					"resourceId":      data.ID.ValueString(),
				})
				assignmentObj["filter_type"] = types.StringValue("none")
			default:
				tflog.Debug(ctx, "Unknown filter type, using schema default", map[string]any{
					"assignmentIndex": i,
					"assignmentId":    assignment.GetId(),
					"filterType":      *filterType,
					"resourceId":      data.ID.ValueString(),
				})
				assignmentObj["filter_type"] = types.StringValue("none")
			}
		} else {
			tflog.Debug(ctx, "No filter type specified, using schema default", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"resourceId":      data.ID.ValueString(),
			})
			assignmentObj["filter_type"] = types.StringValue("none")
		}

		tflog.Debug(ctx, "Creating assignment object value", map[string]any{
			"assignmentIndex": i,
			"assignmentId":    assignment.GetId(),
			"resourceId":      data.ID.ValueString(),
		})

		objValue, diags := types.ObjectValue(
			IosConfigurationTemplatesAssignmentType().(types.ObjectType).AttrTypes,
			assignmentObj,
		)
		if !diags.HasError() {
			tflog.Debug(ctx, "Successfully created assignment object", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"resourceId":      data.ID.ValueString(),
			})
			assignmentValues = append(assignmentValues, objValue)
		} else {
			tflog.Error(ctx, "Failed to create assignment object value", map[string]any{
				"assignmentIndex": i,
				"assignmentId":    assignment.GetId(),
				"errors":          diags.Errors(),
				"resourceId":      data.ID.ValueString(),
			})
		}
	}

	tflog.Debug(ctx, "Creating assignments set", map[string]any{
		"processedAssignments": len(assignmentValues),
		"originalAssignments":  len(assignments),
		"resourceId":           data.ID.ValueString(),
	})

	if len(assignmentValues) > 0 {
		setVal, diags := types.SetValue(IosConfigurationTemplatesAssignmentType(), assignmentValues)
		if diags.HasError() {
			tflog.Error(ctx, "Failed to create assignments set", map[string]any{
				"errors":     diags.Errors(),
				"resourceId": data.ID.ValueString(),
			})
			data.Assignments = types.SetNull(IosConfigurationTemplatesAssignmentType())
		} else {
			tflog.Debug(ctx, "Successfully created assignments set", map[string]any{
				"assignmentCount": len(assignmentValues),
				"resourceId":      data.ID.ValueString(),
			})
			data.Assignments = setVal
		}
	} else {
		tflog.Debug(
			ctx,
			"No valid assignments processed, setting assignments to null",
			map[string]any{
				"resourceId": data.ID.ValueString(),
			},
		)
		data.Assignments = types.SetNull(IosConfigurationTemplatesAssignmentType())
	}

	tflog.Debug(ctx, "Finished mapping assignments to Terraform state", map[string]any{
		"finalAssignmentCount": len(assignmentValues),
		"originalAssignments":  len(assignments),
		"resourceId":           data.ID.ValueString(),
	})
}
