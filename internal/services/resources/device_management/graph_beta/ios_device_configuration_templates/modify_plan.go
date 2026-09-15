package graphBetaIosDeviceConfigurationTemplates

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ModifyPlan validates that exactly one configuration type is specified
func (r *IosDeviceConfigurationTemplatesResource) ModifyPlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan IosDeviceConfigurationTemplatesResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Count how many configuration types are specified
	configCount := 0
	if !plan.GeneralDeviceConfiguration.IsNull() && !plan.GeneralDeviceConfiguration.IsUnknown() {
		configCount++
	}
	if !plan.CustomConfiguration.IsNull() && !plan.CustomConfiguration.IsUnknown() {
		configCount++
	}
	if !plan.TrustedCertificate.IsNull() && !plan.TrustedCertificate.IsUnknown() {
		configCount++
	}
	if !plan.Wifi.IsNull() && !plan.Wifi.IsUnknown() {
		configCount++
	}
	if !plan.ScepCertificate.IsNull() && !plan.ScepCertificate.IsUnknown() {
		configCount++
	}
	if !plan.PkcsCertificate.IsNull() && !plan.PkcsCertificate.IsUnknown() {
		configCount++
	}
	if !plan.EnterpriseWifi.IsNull() && !plan.EnterpriseWifi.IsUnknown() {
		configCount++
	}
	if !plan.EasEmail.IsNull() && !plan.EasEmail.IsUnknown() {
		configCount++
	}
	if !plan.Vpn.IsNull() && !plan.Vpn.IsUnknown() {
		configCount++
	}

	const blockList = "general_device_configuration, custom_configuration, trusted_certificate, wifi, " +
		"scep_certificate, pkcs_certificate, enterprise_wifi, eas_email, or vpn"

	if configCount == 0 {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"Exactly one of "+blockList+" must be specified.",
		)
	} else if configCount > 1 {
		resp.Diagnostics.AddError(
			"Multiple Configurations",
			"Only one of "+blockList+" may be specified.",
		)
	}
}
