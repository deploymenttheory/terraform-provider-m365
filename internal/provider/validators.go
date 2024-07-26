package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/* tenant_id and client_id schema validator */
type guidValidator struct {
	attributeName string
}

func (v guidValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("%s must be a valid GUID if provided", v.attributeName)
}

func (v guidValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("%s must be a valid GUID if provided", v.attributeName)
}

func (v guidValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is null, unknown, or empty string, return without validation
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if value == "" {
		return
	}

	// Proceed with GUID validation only if the string is non-empty
	match, err := regexp.MatchString(helpers.GuidRegex, value)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Invalid %s format", v.attributeName),
			fmt.Sprintf("Error matching GUID format for %s: %s", v.attributeName, err),
		)
		return
	}

	if !match {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Invalid %s format", v.attributeName),
			fmt.Sprintf("The value %q for %s is not a valid GUID. It must match the format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", value, v.attributeName),
		)
	}
}

func validateGUID(attributeName string) validator.String {
	return guidValidator{attributeName: attributeName}
}

/* use_proxy field schema validator */

type useProxyValidator struct{}

func (v useProxyValidator) Description(ctx context.Context) string {
	return "Ensures that proxy_url is set if use_proxy is true."
}

func (v useProxyValidator) MarkdownDescription(ctx context.Context) string {
	return "Ensures that proxy_url is set if use_proxy is true."
}

func validateUseProxy() validator.Bool {
	return useProxyValidator{}
}

func (v useProxyValidator) ValidateBool(ctx context.Context, request validator.BoolRequest, response *validator.BoolResponse) {
	var useProxy types.Bool
	if diags := request.Config.GetAttribute(ctx, path.Root("use_proxy"), &useProxy); diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	var proxyURL types.String
	if diags := request.Config.GetAttribute(ctx, path.Root("proxy_url"), &proxyURL); diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	if useProxy.ValueBool() && (proxyURL.IsNull() || proxyURL.IsUnknown() || proxyURL.ValueString() == "") {
		response.Diagnostics.AddError(
			"Invalid Configuration",
			"The 'proxy_url' attribute must be set when 'use_proxy' is true.",
		)
	}
}

/* redirect_url field schema validator */
type redirectURLValidator struct{}

func (v redirectURLValidator) Description(ctx context.Context) string {
	return "Ensures that redirect_url is a well-formed URL if provided."
}

func (v redirectURLValidator) MarkdownDescription(ctx context.Context) string {
	return "Ensures that redirect_url is a well-formed URL if provided."
}

func (v redirectURLValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {

	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	redirectURL := req.ConfigValue.ValueString()
	if redirectURL == "" {
		return
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Redirect URL",
			fmt.Sprintf("The value %q for redirect_url is not a valid URL: %s", redirectURL, err),
		)
		return
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Redirect URL",
			fmt.Sprintf("The value %q for redirect_url must include a scheme and host (e.g., https://example.com)", redirectURL),
		)
		return
	}

	match, _ := regexp.MatchString(helpers.UrlValidStringRegex, redirectURL)
	if !match {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Redirect URL Format",
			fmt.Sprintf("The value %q for redirect_url contains invalid characters. It must match the format: [A-Za-z0-9-._~%%/:/?=]+", redirectURL),
		)
	}
}

func validateRedirectURL() validator.String {
	return redirectURLValidator{}
}

/* proxy_url field schema validator */

type proxyURLValidator struct{}

func (v proxyURLValidator) Description(ctx context.Context) string {
	return "Ensures that proxy_url is a well-formed URL if provided."
}

func (v proxyURLValidator) MarkdownDescription(ctx context.Context) string {
	return "Ensures that proxy_url is a well-formed URL if provided."
}

func (v proxyURLValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	proxyURL := req.ConfigValue.ValueString()
	if proxyURL == "" {
		return
	}

	_, err := url.Parse(proxyURL)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Proxy URL",
			fmt.Sprintf("The value %q for proxy_url is not a valid URL: %s", proxyURL, err),
		)
		return
	}

	match, _ := regexp.MatchString(helpers.UrlValidStringRegex, proxyURL)
	if !match {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Proxy URL Format",
			fmt.Sprintf("The value %q for proxy_url contains invalid characters. It must match the format: [A-Za-z0-9-._~%%/:/?=]+", proxyURL),
		)
	}
}

func validateProxyURL() validator.String {
	return proxyURLValidator{}
}
