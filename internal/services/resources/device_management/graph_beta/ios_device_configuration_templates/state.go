package graphBetaIosDeviceConfigurationTemplates

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	graphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/convert"
)

// Attribute type helpers for ObjectNull calls

func GeneralDeviceConfigurationType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

func CustomConfigurationType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"payload_file_name": types.StringType,
			"payload":           types.StringType,
			"payload_name":      types.StringType,
		},
	}
}

func TrustedCertificateType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cert_file_name":           types.StringType,
			"trusted_root_certificate": types.StringType,
		},
	}
}

func WifiType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_name":                        types.StringType,
			"ssid":                                types.StringType,
			"connect_automatically":               types.BoolType,
			"connect_when_network_name_is_hidden": types.BoolType,
			"wifi_security_type":                  types.StringType,
			"pre_shared_key":                      types.StringType,
			"disable_mac_address_randomization":   types.BoolType,
			"proxy_settings":                      types.StringType,
			"proxy_manual_address":                types.StringType,
			"proxy_manual_port":                   types.Int32Type,
			"proxy_automatic_configuration_url":   types.StringType,
		},
	}
}

// customSanAttrTypes is the element type for the custom_subject_alternative_names set, shared by
// both certificate blocks.
func customSanAttrTypes() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"san_type": types.StringType,
		"name":     types.StringType,
	}}
}

// extendedKeyUsageAttrTypes is the element type for the extended_key_usages set.
func extendedKeyUsageAttrTypes() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"name":              types.StringType,
		"object_identifier": types.StringType,
	}}
}

func ScepCertificateType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"renewal_threshold_percentage":           types.Int32Type,
			"certificate_store":                      types.StringType,
			"certificate_validity_period_scale":      types.StringType,
			"certificate_validity_period_value":      types.Int32Type,
			"subject_name_format":                    types.StringType,
			"subject_name_format_string":             types.StringType,
			"subject_alternative_name_type":          types.SetType{ElemType: types.StringType},
			"subject_alternative_name_format_string": types.StringType,
			"root_certificate_odata_bind":            types.StringType,
			"key_size":                               types.StringType,
			"key_usage":                              types.SetType{ElemType: types.StringType},
			"custom_subject_alternative_names":       types.SetType{ElemType: customSanAttrTypes()},
			"extended_key_usages": types.SetType{
				ElemType: extendedKeyUsageAttrTypes(),
			},
			"scep_server_urls": types.SetType{ElemType: types.StringType},
		},
	}
}

func PkcsCertificateType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"renewal_threshold_percentage":           types.Int32Type,
			"certificate_store":                      types.StringType,
			"certificate_validity_period_scale":      types.StringType,
			"certificate_validity_period_value":      types.Int32Type,
			"subject_name_format":                    types.StringType,
			"subject_name_format_string":             types.StringType,
			"subject_alternative_name_type":          types.SetType{ElemType: types.StringType},
			"subject_alternative_name_format_string": types.StringType,
			"certification_authority":                types.StringType,
			"certification_authority_name":           types.StringType,
			"certificate_template_name":              types.StringType,
			"custom_subject_alternative_names":       types.SetType{ElemType: customSanAttrTypes()},
		},
	}
}

func EnterpriseWifiType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_name":                               types.StringType,
			"ssid":                                       types.StringType,
			"connect_automatically":                      types.BoolType,
			"connect_when_network_name_is_hidden":        types.BoolType,
			"wifi_security_type":                         types.StringType,
			"disable_mac_address_randomization":          types.BoolType,
			"proxy_settings":                             types.StringType,
			"proxy_manual_address":                       types.StringType,
			"proxy_manual_port":                          types.Int32Type,
			"proxy_automatic_configuration_url":          types.StringType,
			"eap_type":                                   types.StringType,
			"eap_fast_configuration":                     types.StringType,
			"authentication_method":                      types.StringType,
			"inner_authentication_protocol_for_eap_ttls": types.StringType,
			"outer_identity_privacy_temporary_value":     types.StringType,
			"username_format_string":                     types.StringType,
			"password_format_string":                     types.StringType,
			"trusted_server_certificate_names": types.SetType{
				ElemType: types.StringType,
			},
			"root_certificates_for_server_validation_odata_bind": types.SetType{
				ElemType: types.StringType,
			},
			"identity_certificate_for_client_authentication_odata_bind": types.StringType,
		},
	}
}

func EasEmailType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_name":                       types.StringType,
			"host_name":                          types.StringType,
			"authentication_method":              types.StringType,
			"eas_services":                       types.SetType{ElemType: types.StringType},
			"eas_services_user_override_enabled": types.BoolType,
			"duration_of_email_to_sync":          types.StringType,
			"email_address_source":               types.StringType,
			"username_source":                    types.StringType,
			"username_aad_source":                types.StringType,
			"user_domain_name_source":            types.StringType,
			"custom_domain_name":                 types.StringType,
			"require_ssl":                        types.BoolType,
			"use_oauth":                          types.BoolType,
			"per_app_vpn_profile_id":             types.StringType,
			"block_moving_messages_to_other_email_accounts": types.BoolType,
			"block_sending_email_from_third_party_apps":     types.BoolType,
			"block_syncing_recently_used_email_addresses":   types.BoolType,
			"require_smime":                                      types.BoolType,
			"smime_enable_per_message_switch":                    types.BoolType,
			"smime_signing_enabled":                              types.BoolType,
			"smime_signing_user_override_enabled":                types.BoolType,
			"smime_signing_certificate_user_override_enabled":    types.BoolType,
			"smime_encrypt_by_default_enabled":                   types.BoolType,
			"smime_encrypt_by_default_user_override_enabled":     types.BoolType,
			"smime_encryption_certificate_user_override_enabled": types.BoolType,
			"signing_certificate_type":                           types.StringType,
			"encryption_certificate_type":                        types.StringType,
			"identity_certificate_odata_bind":                    types.StringType,
			"smime_signing_certificate_odata_bind":               types.StringType,
			"smime_encryption_certificate_odata_bind":            types.StringType,
		},
	}
}

func vpnServerAttrTypes() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"address":           types.StringType,
		"description":       types.StringType,
		"is_default_server": types.BoolType,
	}}
}

func vpnProxyServerAttrTypes() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"address":                            types.StringType,
		"port":                               types.Int32Type,
		"automatic_configuration_script_url": types.StringType,
	}}
}

func appListItemAttrTypes() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"name":          types.StringType,
		"app_id":        types.StringType,
		"publisher":     types.StringType,
		"app_store_url": types.StringType,
	}}
}

// keyValueAttrTypes is the element type for custom_data, whose Graph entries use a "key" field.
func keyValueAttrTypes() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	}}
}

// keyValuePairAttrTypes is the element type for custom_key_value_data, whose Graph entries use a
// "name" field rather than "key".
func keyValuePairAttrTypes() types.ObjectType {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"name":  types.StringType,
		"value": types.StringType,
	}}
}

func VpnType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name":                     types.StringType,
			"connection_type":                     types.StringType,
			"authentication_method":               types.StringType,
			"identifier":                          types.StringType,
			"provider_type":                       types.StringType,
			"realm":                               types.StringType,
			"role":                                types.StringType,
			"login_group_or_domain":               types.StringType,
			"enable_split_tunneling":              types.BoolType,
			"enable_per_app":                      types.BoolType,
			"include_all_networks":                types.BoolType,
			"exclude_local_networks":              types.BoolType,
			"safari_domains":                      types.SetType{ElemType: types.StringType},
			"associated_domains":                  types.SetType{ElemType: types.StringType},
			"excluded_domains":                    types.SetType{ElemType: types.StringType},
			"disconnect_on_idle":                  types.BoolType,
			"disconnect_on_idle_timer_in_seconds": types.Int32Type,
			"disable_on_demand_user_override":     types.BoolType,
			"opt_in_to_device_id_sharing":         types.BoolType,
			"microsoft_tunnel_site_id":            types.StringType,
			"strict_enforcement":                  types.BoolType,
			"cloud_name":                          types.StringType,
			"user_domain":                         types.StringType,
			"exclude_list":                        types.SetType{ElemType: types.StringType},
			"server":                              vpnServerAttrTypes(),
			"proxy_server":                        vpnProxyServerAttrTypes(),
			"targeted_mobile_apps":                types.SetType{ElemType: appListItemAttrTypes()},
			"custom_data":                         types.SetType{ElemType: keyValueAttrTypes()},
			"custom_key_value_data":               types.SetType{ElemType: keyValuePairAttrTypes()},
			"identity_certificate_odata_bind":     types.StringType,
		},
	}
}

func MapRemoteResourceStateToTerraform(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	remoteResource graphmodels.DeviceConfigurationable,
) {
	if remoteResource == nil {
		tflog.Debug(ctx, "Remote resource is nil")
		return
	}

	tflog.Debug(ctx, "Starting to map remote state to Terraform state", map[string]any{
		"resourceId": convert.GraphToFrameworkString(remoteResource.GetId()).ValueString(),
	})

	// Map common properties
	data.ID = convert.GraphToFrameworkString(remoteResource.GetId())
	data.DisplayName = convert.GraphToFrameworkString(remoteResource.GetDisplayName())
	data.Description = convert.GraphToFrameworkString(remoteResource.GetDescription())
	data.RoleScopeTagIds = convert.GraphToFrameworkStringSet(
		ctx,
		remoteResource.GetRoleScopeTagIds(),
	)

	// Snapshot the incoming state before clearing. Graph does not return write-only values
	// (pre_shared_key) or @odata.bind navigation references (root_certificate_odata_bind), so the
	// mappers recover them from the prior state — they must read the snapshot, not the live model.
	// types.Object is an immutable value type, so copying the struct is a safe point-in-time copy.
	prior := *data

	// Null every block before mapping, so each map function only has to set its own. With five
	// mutually exclusive blocks, clearing the others inside each function is easy to get wrong.
	clearConfigurationBlocks(data)

	// Map specific configuration based on type
	switch config := remoteResource.(type) {
	case *graphmodels.IosGeneralDeviceConfiguration:
		mapIosGeneralDeviceConfiguration(ctx, data, config)
	case *graphmodels.IosCustomConfiguration:
		mapIosCustomConfiguration(ctx, data, config)
	case *graphmodels.IosTrustedRootCertificate:
		mapIosTrustedRootCertificate(ctx, data, config)
	case *graphmodels.IosWiFiConfiguration:
		mapIosWiFiConfiguration(ctx, data, &prior, config)
	case *graphmodels.IosScepCertificateProfile:
		mapIosScepCertificateProfile(ctx, data, &prior, config)
	case *graphmodels.IosPkcsCertificateProfile:
		mapIosPkcsCertificateProfile(ctx, data, config)
	// IosEnterpriseWiFiConfiguration embeds IosWiFiConfiguration, but Go type switches match the
	// concrete type, so this arm is reached for enterprise profiles and the plain wifi arm is not.
	// Ordering between them is therefore not load-bearing.
	case *graphmodels.IosEnterpriseWiFiConfiguration:
		mapIosEnterpriseWiFiConfiguration(ctx, data, &prior, config)
	case *graphmodels.IosEasEmailProfileConfiguration:
		mapIosEasEmailProfileConfiguration(ctx, data, &prior, config)
	case *graphmodels.IosVpnConfiguration:
		mapIosVpnConfiguration(ctx, data, &prior, config)
	// IosikEv2VpnConfiguration embeds IosVpnConfiguration, and a Go type switch matches the concrete
	// type — so without this arm an IKEv2 profile would fall through to default and vanish from
	// state. The schema excludes ikEv2 from connection_type, so this is only reachable via import or
	// an out-of-band type change: map the shared fields and warn that the rest is unmanaged.
	case *graphmodels.IosikEv2VpnConfiguration:
		tflog.Warn(
			ctx,
			"Profile is an iosikEv2VpnConfiguration. IKEv2-specific properties are not managed by "+
				"this resource and will not appear in state.",
			map[string]any{"resourceId": data.ID.ValueString()},
		)
		mapIosVpnConfiguration(ctx, data, &prior, &config.IosVpnConfiguration)
	default:
		tflog.Error(ctx, "Unknown device configuration type", map[string]any{
			"type": fmt.Sprintf("%T", config),
		})
	}

	// Map assignments
	assignments := remoteResource.GetAssignments()
	tflog.Debug(ctx, "Retrieved assignments from remote resource", map[string]any{
		"assignmentCount": len(assignments),
		"resourceId":      data.ID.ValueString(),
	})

	if len(assignments) == 0 {
		data.Assignments = types.SetNull(IosConfigurationTemplatesAssignmentType())
	} else {
		mapAssignmentsToTerraform(ctx, data, assignments)
	}

	tflog.Debug(
		ctx,
		fmt.Sprintf("Finished mapping resource %s with id %s", ResourceName, data.ID.ValueString()),
	)
}

func mapIosGeneralDeviceConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	_ *graphmodels.IosGeneralDeviceConfiguration,
) {
	tflog.Debug(ctx, "Mapping IosGeneralDeviceConfiguration")

	objectValue, diags := types.ObjectValueFrom(
		ctx,
		GeneralDeviceConfigurationType().AttrTypes,
		GeneralDeviceConfigurationResourceModel{},
	)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create general device configuration object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}
	data.GeneralDeviceConfiguration = objectValue
}

func mapIosCustomConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosCustomConfiguration,
) {
	tflog.Debug(ctx, "Mapping IosCustomConfiguration")

	customConfigModel := CustomConfigurationResourceModel{
		PayloadFileName: convert.GraphToFrameworkString(config.GetPayloadFileName()),
		Payload:         convert.GraphToFrameworkBytes(config.GetPayload()),
		PayloadName:     convert.GraphToFrameworkString(config.GetPayloadName()),
	}

	objectValue, diags := types.ObjectValueFrom(
		ctx,
		CustomConfigurationType().AttrTypes,
		customConfigModel,
	)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create custom configuration object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.CustomConfiguration = objectValue
}

func mapIosTrustedRootCertificate(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosTrustedRootCertificate,
) {
	tflog.Debug(ctx, "Mapping IosTrustedRootCertificate")

	// Convert binary certificate data back to base64 for state consistency
	var trustedRootCert types.String
	if certBytes := config.GetTrustedRootCertificate(); certBytes != nil {
		trustedRootCert = types.StringValue(base64.StdEncoding.EncodeToString(certBytes))
	} else {
		trustedRootCert = types.StringNull()
	}

	certModel := TrustedCertificateResourceModel{
		CertFileName:           convert.GraphToFrameworkString(config.GetCertFileName()),
		TrustedRootCertificate: trustedRootCert,
	}

	objectValue, diags := types.ObjectValueFrom(ctx, TrustedCertificateType().AttrTypes, certModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create trusted certificate object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.TrustedCertificate = objectValue
}

func mapIosWiFiConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	prior *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosWiFiConfiguration,
) {
	tflog.Debug(ctx, "Mapping IosWiFiConfiguration")

	// Preserve the pre-shared key from existing state. Graph does not return the secret on read,
	// so mapping the response value directly would clear it from state and produce perpetual drift.
	preSharedKey := types.StringNull()
	if !prior.Wifi.IsNull() && !prior.Wifi.IsUnknown() {
		var existingWifiData WifiResourceModel
		diags := prior.Wifi.As(ctx, &existingWifiData, basetypes.ObjectAsOptions{})
		if !diags.HasError() && !existingWifiData.PreSharedKey.IsNull() {
			preSharedKey = existingWifiData.PreSharedKey
		}
	}

	// If the response does carry a value (some tenants echo it back), prefer the remote one so a
	// key changed outside Terraform still surfaces as drift.
	if remoteKey := config.GetPreSharedKey(); remoteKey != nil && *remoteKey != "" {
		preSharedKey = types.StringValue(*remoteKey)
	}

	wifiModel := WifiResourceModel{
		NetworkName: convert.GraphToFrameworkString(config.GetNetworkName()),
		Ssid:        convert.GraphToFrameworkString(config.GetSsid()),
		ConnectAutomatically: convert.GraphToFrameworkBool(
			config.GetConnectAutomatically(),
		),
		ConnectWhenNetworkNameIsHidden: convert.GraphToFrameworkBool(
			config.GetConnectWhenNetworkNameIsHidden(),
		),
		WifiSecurityType: convert.GraphToFrameworkEnum(config.GetWiFiSecurityType()),
		PreSharedKey:     preSharedKey,
		DisableMacAddressRandomization: convert.GraphToFrameworkBool(
			config.GetDisableMacAddressRandomization(),
		),
		ProxySettings: convert.GraphToFrameworkEnum(config.GetProxySettings()),
		ProxyManualAddress: convert.GraphToFrameworkString(
			config.GetProxyManualAddress(),
		),
		ProxyManualPort: convert.GraphToFrameworkInt32(config.GetProxyManualPort()),
		ProxyAutomaticConfigurationUrl: convert.GraphToFrameworkString(
			config.GetProxyAutomaticConfigurationUrl(),
		),
	}

	objectValue, diags := types.ObjectValueFrom(ctx, WifiType().AttrTypes, wifiModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create wifi object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.Wifi = objectValue
}

// clearConfigurationBlocks nulls all mutually exclusive configuration blocks. Called before
// mapping so each map function sets only its own block.
func clearConfigurationBlocks(data *IosDeviceConfigurationTemplatesResourceModel) {
	data.GeneralDeviceConfiguration = types.ObjectNull(GeneralDeviceConfigurationType().AttrTypes)
	data.CustomConfiguration = types.ObjectNull(CustomConfigurationType().AttrTypes)
	data.TrustedCertificate = types.ObjectNull(TrustedCertificateType().AttrTypes)
	data.Wifi = types.ObjectNull(WifiType().AttrTypes)
	data.ScepCertificate = types.ObjectNull(ScepCertificateType().AttrTypes)
	data.PkcsCertificate = types.ObjectNull(PkcsCertificateType().AttrTypes)
	data.EnterpriseWifi = types.ObjectNull(EnterpriseWifiType().AttrTypes)
	data.EasEmail = types.ObjectNull(EasEmailType().AttrTypes)
	data.Vpn = types.ObjectNull(VpnType().AttrTypes)
}

func mapIosVpnConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	prior *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosVpnConfiguration,
) {
	tflog.Debug(ctx, "Mapping IosVpnConfiguration")

	// The identity certificate is a navigation property Graph does not echo back as additionalData.
	identityCertRef := types.StringNull()
	if !prior.Vpn.IsNull() && !prior.Vpn.IsUnknown() {
		var existing VpnResourceModel
		diags := prior.Vpn.As(ctx, &existing, basetypes.ObjectAsOptions{})
		if !diags.HasError() && !existing.IdentityCertificateOdataBind.IsNull() {
			identityCertRef = existing.IdentityCertificateOdataBind
		}
	}
	if identityCertRef.IsNull() {
		if cert := config.GetIdentityCertificate(); cert != nil {
			if id := cert.GetId(); id != nil {
				identityCertRef = types.StringValue(odataBindURL(*id))
			}
		}
	}

	vpnModel := VpnResourceModel{
		ConnectionName:       convert.GraphToFrameworkString(config.GetConnectionName()),
		ConnectionType:       convert.GraphToFrameworkEnum(config.GetConnectionType()),
		AuthenticationMethod: convert.GraphToFrameworkEnum(config.GetAuthenticationMethod()),
		Identifier:           convert.GraphToFrameworkString(config.GetIdentifier()),
		ProviderType:         convert.GraphToFrameworkEnum(config.GetProviderType()),
		Realm:                convert.GraphToFrameworkString(config.GetRealm()),
		Role:                 convert.GraphToFrameworkString(config.GetRole()),
		LoginGroupOrDomain:   convert.GraphToFrameworkString(config.GetLoginGroupOrDomain()),
		EnableSplitTunneling: convert.GraphToFrameworkBool(config.GetEnableSplitTunneling()),
		EnablePerApp:         convert.GraphToFrameworkBool(config.GetEnablePerApp()),
		IncludeAllNetworks:   convert.GraphToFrameworkBool(config.GetIncludeAllNetworks()),
		ExcludeLocalNetworks: convert.GraphToFrameworkBool(config.GetExcludeLocalNetworks()),
		SafariDomains:        convert.GraphToFrameworkStringSet(ctx, config.GetSafariDomains()),
		AssociatedDomains:    convert.GraphToFrameworkStringSet(ctx, config.GetAssociatedDomains()),
		ExcludedDomains:      convert.GraphToFrameworkStringSet(ctx, config.GetExcludedDomains()),
		DisconnectOnIdle:     convert.GraphToFrameworkBool(config.GetDisconnectOnIdle()),
		DisconnectOnIdleTimerInSeconds: convert.GraphToFrameworkInt32(
			config.GetDisconnectOnIdleTimerInSeconds(),
		),
		DisableOnDemandUserOverride: convert.GraphToFrameworkBool(
			config.GetDisableOnDemandUserOverride(),
		),
		OptInToDeviceIdSharing: convert.GraphToFrameworkBool(
			config.GetOptInToDeviceIdSharing(),
		),
		MicrosoftTunnelSiteId: convert.GraphToFrameworkString(
			config.GetMicrosoftTunnelSiteId(),
		),
		StrictEnforcement: convert.GraphToFrameworkBool(config.GetStrictEnforcement()),
		CloudName:         convert.GraphToFrameworkString(config.GetCloudName()),
		UserDomain:        convert.GraphToFrameworkString(config.GetUserDomain()),
		ExcludeList: convert.GraphToFrameworkStringSet(
			ctx,
			config.GetExcludeList(),
		),
		Server:             mapVpnServerToObject(ctx, config.GetServer()),
		ProxyServer:        mapVpnProxyServerToObject(ctx, config.GetProxyServer()),
		TargetedMobileApps: mapAppListItemsToSet(ctx, config.GetTargetedMobileApps()),
		CustomData:         mapCustomDataToSet(ctx, config.GetCustomData()),
		CustomKeyValueData: mapCustomKeyValueDataToSet(
			ctx,
			config.GetCustomKeyValueData(),
		),
		IdentityCertificateOdataBind: identityCertRef,
	}

	objectValue, diags := types.ObjectValueFrom(ctx, VpnType().AttrTypes, vpnModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create vpn object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.Vpn = objectValue
}

func mapVpnServerToObject(ctx context.Context, server graphmodels.VpnServerable) types.Object {
	if server == nil {
		return types.ObjectNull(vpnServerAttrTypes().AttrTypes)
	}

	serverModel := VpnServerResourceModel{
		Address:         convert.GraphToFrameworkString(server.GetAddress()),
		Description:     convert.GraphToFrameworkString(server.GetDescription()),
		IsDefaultServer: convert.GraphToFrameworkBool(server.GetIsDefaultServer()),
	}

	objectValue, diags := types.ObjectValueFrom(ctx, vpnServerAttrTypes().AttrTypes, serverModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create vpn server object", map[string]any{
			"errors": diags.Errors(),
		})
		return types.ObjectNull(vpnServerAttrTypes().AttrTypes)
	}

	return objectValue
}

// stringNullIfEmpty converts a Graph string pointer to a Terraform string, treating both nil and the
// empty string as null. Use it for Optional attributes that Graph returns as "" when unset, where
// convert.GraphToFrameworkString would surface "" against a null configuration value and trip
// Terraform's consistency check.
func stringNullIfEmpty(value *string) types.String {
	if value == nil || *value == "" {
		return types.StringNull()
	}
	return types.StringValue(*value)
}

func mapVpnProxyServerToObject(
	ctx context.Context,
	proxy graphmodels.VpnProxyServerable,
) types.Object {
	if proxy == nil {
		return types.ObjectNull(vpnProxyServerAttrTypes().AttrTypes)
	}

	// Graph returns an empty string rather than omitting automaticConfigurationScriptUrl when no PAC
	// script is configured. The attribute is Optional and not Computed, so surfacing "" where the
	// configuration left the field unset looks like provider-produced drift and fails the apply with
	// `was null, but now cty.StringVal("")`. Normalize empty back to null.
	proxyModel := VpnProxyServerResourceModel{
		Address: convert.GraphToFrameworkString(proxy.GetAddress()),
		Port:    convert.GraphToFrameworkInt32(proxy.GetPort()),
		AutomaticConfigurationScriptUrl: stringNullIfEmpty(
			proxy.GetAutomaticConfigurationScriptUrl(),
		),
	}

	objectValue, diags := types.ObjectValueFrom(
		ctx,
		vpnProxyServerAttrTypes().AttrTypes,
		proxyModel,
	)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create vpn proxy server object", map[string]any{
			"errors": diags.Errors(),
		})
		return types.ObjectNull(vpnProxyServerAttrTypes().AttrTypes)
	}

	return objectValue
}

func mapAppListItemsToSet(ctx context.Context, apps []graphmodels.AppListItemable) types.Set {
	if len(apps) == 0 {
		return types.SetNull(appListItemAttrTypes())
	}

	appModels := make([]AppListItemResourceModel, 0, len(apps))
	for _, app := range apps {
		if app == nil {
			continue
		}
		appModels = append(appModels, AppListItemResourceModel{
			Name:        convert.GraphToFrameworkString(app.GetName()),
			AppId:       convert.GraphToFrameworkString(app.GetAppId()),
			Publisher:   convert.GraphToFrameworkString(app.GetPublisher()),
			AppStoreUrl: convert.GraphToFrameworkString(app.GetAppStoreUrl()),
		})
	}

	setValue, diags := types.SetValueFrom(ctx, appListItemAttrTypes(), appModels)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create targeted mobile apps set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(appListItemAttrTypes())
	}

	return setValue
}

func mapCustomDataToSet(ctx context.Context, kvs []graphmodels.KeyValueable) types.Set {
	if len(kvs) == 0 {
		return types.SetNull(keyValueAttrTypes())
	}

	kvModels := make([]KeyValueResourceModel, 0, len(kvs))
	for _, kv := range kvs {
		if kv == nil {
			continue
		}
		kvModels = append(kvModels, KeyValueResourceModel{
			Key:   convert.GraphToFrameworkString(kv.GetKey()),
			Value: convert.GraphToFrameworkString(kv.GetValue()),
		})
	}

	setValue, diags := types.SetValueFrom(ctx, keyValueAttrTypes(), kvModels)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create custom data set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(keyValueAttrTypes())
	}

	return setValue
}

func mapCustomKeyValueDataToSet(
	ctx context.Context,
	kvs []graphmodels.KeyValuePairable,
) types.Set {
	if len(kvs) == 0 {
		return types.SetNull(keyValuePairAttrTypes())
	}

	kvModels := make([]KeyValuePairResourceModel, 0, len(kvs))
	for _, kv := range kvs {
		if kv == nil {
			continue
		}
		kvModels = append(kvModels, KeyValuePairResourceModel{
			Name:  convert.GraphToFrameworkString(kv.GetName()),
			Value: convert.GraphToFrameworkString(kv.GetValue()),
		})
	}

	setValue, diags := types.SetValueFrom(ctx, keyValuePairAttrTypes(), kvModels)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create custom key value data set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(keyValuePairAttrTypes())
	}

	return setValue
}

// odataBindURL renders a device configuration id as the full Graph URL used by @odata.bind
// references. Mirrors normalizeODataBind in construct_resource.go, which accepts either form.
func odataBindURL(id string) string {
	return fmt.Sprintf(
		"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('%s')",
		id,
	)
}

func mapIosEnterpriseWiFiConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	prior *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosEnterpriseWiFiConfiguration,
) {
	tflog.Debug(ctx, "Mapping IosEnterpriseWiFiConfiguration")

	// Recover the values Graph never returns: the password format string (write-only) and both
	// certificate @odata.bind references (navigation properties, not echoed as additionalData).
	passwordFormatString := types.StringNull()
	identityCertRef := types.StringNull()
	rootCertRefs := types.SetNull(types.StringType)

	if !prior.EnterpriseWifi.IsNull() && !prior.EnterpriseWifi.IsUnknown() {
		var existing EnterpriseWifiResourceModel
		diags := prior.EnterpriseWifi.As(ctx, &existing, basetypes.ObjectAsOptions{})
		if !diags.HasError() {
			if !existing.PasswordFormatString.IsNull() {
				passwordFormatString = existing.PasswordFormatString
			}
			if !existing.IdentityCertificateForClientAuthenticationOdataBind.IsNull() {
				identityCertRef = existing.IdentityCertificateForClientAuthenticationOdataBind
			}
			if !existing.RootCertificatesForServerValidationOdataBind.IsNull() {
				rootCertRefs = existing.RootCertificatesForServerValidationOdataBind
			}
		}
	}

	// If Graph does echo the password back, prefer the remote value so an out-of-band change shows
	// up as drift rather than being masked by the configured value.
	if remote := config.GetPasswordFormatString(); remote != nil && *remote != "" {
		passwordFormatString = types.StringValue(*remote)
	}

	// With no prior state (import), fall back to the expanded navigation properties.
	if identityCertRef.IsNull() {
		if cert := config.GetIdentityCertificateForClientAuthentication(); cert != nil {
			if id := cert.GetId(); id != nil {
				identityCertRef = types.StringValue(odataBindURL(*id))
			}
		}
	}
	if rootCertRefs.IsNull() {
		if certs := config.GetRootCertificatesForServerValidation(); len(certs) > 0 {
			urls := make([]string, 0, len(certs))
			for _, cert := range certs {
				if cert == nil {
					continue
				}
				if id := cert.GetId(); id != nil {
					urls = append(urls, odataBindURL(*id))
				}
			}
			if len(urls) > 0 {
				if setValue, diags := types.SetValueFrom(
					ctx,
					types.StringType,
					urls,
				); !diags.HasError() {
					rootCertRefs = setValue
				}
			}
		}
	}

	wifiModel := EnterpriseWifiResourceModel{
		NetworkName:          convert.GraphToFrameworkString(config.GetNetworkName()),
		Ssid:                 convert.GraphToFrameworkString(config.GetSsid()),
		ConnectAutomatically: convert.GraphToFrameworkBool(config.GetConnectAutomatically()),
		ConnectWhenNetworkNameIsHidden: convert.GraphToFrameworkBool(
			config.GetConnectWhenNetworkNameIsHidden(),
		),
		WifiSecurityType: convert.GraphToFrameworkEnum(config.GetWiFiSecurityType()),
		DisableMacAddressRandomization: convert.GraphToFrameworkBool(
			config.GetDisableMacAddressRandomization(),
		),
		ProxySettings:      convert.GraphToFrameworkEnum(config.GetProxySettings()),
		ProxyManualAddress: convert.GraphToFrameworkString(config.GetProxyManualAddress()),
		ProxyManualPort:    convert.GraphToFrameworkInt32(config.GetProxyManualPort()),
		ProxyAutomaticConfigurationUrl: convert.GraphToFrameworkString(
			config.GetProxyAutomaticConfigurationUrl(),
		),
		EapType:              convert.GraphToFrameworkEnum(config.GetEapType()),
		EapFastConfiguration: convert.GraphToFrameworkEnum(config.GetEapFastConfiguration()),
		AuthenticationMethod: convert.GraphToFrameworkEnum(config.GetAuthenticationMethod()),
		InnerAuthenticationProtocolForEapTtls: convert.GraphToFrameworkEnum(
			config.GetInnerAuthenticationProtocolForEapTtls(),
		),
		OuterIdentityPrivacyTemporaryValue: convert.GraphToFrameworkString(
			config.GetOuterIdentityPrivacyTemporaryValue(),
		),
		UsernameFormatString: convert.GraphToFrameworkString(config.GetUsernameFormatString()),
		PasswordFormatString: passwordFormatString,
		TrustedServerCertificateNames: convert.GraphToFrameworkStringSet(
			ctx,
			config.GetTrustedServerCertificateNames(),
		),
		RootCertificatesForServerValidationOdataBind:        rootCertRefs,
		IdentityCertificateForClientAuthenticationOdataBind: identityCertRef,
	}

	objectValue, diags := types.ObjectValueFrom(ctx, EnterpriseWifiType().AttrTypes, wifiModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create enterprise wifi object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.EnterpriseWifi = objectValue
}

func mapIosEasEmailProfileConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	prior *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosEasEmailProfileConfiguration,
) {
	tflog.Debug(ctx, "Mapping IosEasEmailProfileConfiguration")

	// All three certificate references are navigation properties Graph does not echo back.
	identityCertRef := types.StringNull()
	smimeSigningCertRef := types.StringNull()
	smimeEncryptionCertRef := types.StringNull()

	if !prior.EasEmail.IsNull() && !prior.EasEmail.IsUnknown() {
		var existing EasEmailResourceModel
		diags := prior.EasEmail.As(ctx, &existing, basetypes.ObjectAsOptions{})
		if !diags.HasError() {
			if !existing.IdentityCertificateOdataBind.IsNull() {
				identityCertRef = existing.IdentityCertificateOdataBind
			}
			if !existing.SmimeSigningCertificateOdataBind.IsNull() {
				smimeSigningCertRef = existing.SmimeSigningCertificateOdataBind
			}
			if !existing.SmimeEncryptionCertificateOdataBind.IsNull() {
				smimeEncryptionCertRef = existing.SmimeEncryptionCertificateOdataBind
			}
		}
	}

	// With no prior state (import), fall back to the expanded navigation properties.
	if identityCertRef.IsNull() {
		if cert := config.GetIdentityCertificate(); cert != nil {
			if id := cert.GetId(); id != nil {
				identityCertRef = types.StringValue(odataBindURL(*id))
			}
		}
	}
	if smimeSigningCertRef.IsNull() {
		if cert := config.GetSmimeSigningCertificate(); cert != nil {
			if id := cert.GetId(); id != nil {
				smimeSigningCertRef = types.StringValue(odataBindURL(*id))
			}
		}
	}
	if smimeEncryptionCertRef.IsNull() {
		if cert := config.GetSmimeEncryptionCertificate(); cert != nil {
			if id := cert.GetId(); id != nil {
				smimeEncryptionCertRef = types.StringValue(odataBindURL(*id))
			}
		}
	}

	easModel := EasEmailResourceModel{
		AccountName:          convert.GraphToFrameworkString(config.GetAccountName()),
		HostName:             convert.GraphToFrameworkString(config.GetHostName()),
		AuthenticationMethod: convert.GraphToFrameworkEnum(config.GetAuthenticationMethod()),
		EasServices:          mapEasServicesToSet(ctx, config.GetEasServices()),
		EasServicesUserOverrideEnabled: convert.GraphToFrameworkBool(
			config.GetEasServicesUserOverrideEnabled(),
		),
		DurationOfEmailToSync: convert.GraphToFrameworkEnum(config.GetDurationOfEmailToSync()),
		EmailAddressSource:    convert.GraphToFrameworkEnum(config.GetEmailAddressSource()),
		UsernameSource:        convert.GraphToFrameworkEnum(config.GetUsernameSource()),
		UsernameAADSource:     convert.GraphToFrameworkEnum(config.GetUsernameAADSource()),
		UserDomainNameSource:  convert.GraphToFrameworkEnum(config.GetUserDomainNameSource()),
		CustomDomainName:      convert.GraphToFrameworkString(config.GetCustomDomainName()),
		RequireSsl:            convert.GraphToFrameworkBool(config.GetRequireSsl()),
		UseOAuth:              convert.GraphToFrameworkBool(config.GetUseOAuth()),
		PerAppVPNProfileId:    convert.GraphToFrameworkString(config.GetPerAppVPNProfileId()),
		BlockMovingMessagesToOtherEmailAccounts: convert.GraphToFrameworkBool(
			config.GetBlockMovingMessagesToOtherEmailAccounts(),
		),
		BlockSendingEmailFromThirdPartyApps: convert.GraphToFrameworkBool(
			config.GetBlockSendingEmailFromThirdPartyApps(),
		),
		BlockSyncingRecentlyUsedEmailAddresses: convert.GraphToFrameworkBool(
			config.GetBlockSyncingRecentlyUsedEmailAddresses(),
		),
		RequireSmime: convert.GraphToFrameworkBool(config.GetRequireSmime()),
		SmimeEnablePerMessageSwitch: convert.GraphToFrameworkBool(
			config.GetSmimeEnablePerMessageSwitch(),
		),
		SmimeSigningEnabled: convert.GraphToFrameworkBool(config.GetSmimeSigningEnabled()),
		SmimeSigningUserOverrideEnabled: convert.GraphToFrameworkBool(
			config.GetSmimeSigningUserOverrideEnabled(),
		),
		SmimeSigningCertificateUserOverrideEnabled: convert.GraphToFrameworkBool(
			config.GetSmimeSigningCertificateUserOverrideEnabled(),
		),
		SmimeEncryptByDefaultEnabled: convert.GraphToFrameworkBool(
			config.GetSmimeEncryptByDefaultEnabled(),
		),
		SmimeEncryptByDefaultUserOverrideEnabled: convert.GraphToFrameworkBool(
			config.GetSmimeEncryptByDefaultUserOverrideEnabled(),
		),
		SmimeEncryptionCertificateUserOverrideEnabled: convert.GraphToFrameworkBool(
			config.GetSmimeEncryptionCertificateUserOverrideEnabled(),
		),
		SigningCertificateType: convert.GraphToFrameworkEnum(
			config.GetSigningCertificateType(),
		),
		EncryptionCertificateType: convert.GraphToFrameworkEnum(
			config.GetEncryptionCertificateType(),
		),
		IdentityCertificateOdataBind:        identityCertRef,
		SmimeSigningCertificateOdataBind:    smimeSigningCertRef,
		SmimeEncryptionCertificateOdataBind: smimeEncryptionCertRef,
	}

	objectValue, diags := types.ObjectValueFrom(ctx, EasEmailType().AttrTypes, easModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create EAS email object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.EasEmail = objectValue
}

// mapEasServicesToSet decomposes the EasServices bitmask back into individual values. As with
// SubjectAlternativeNameType, "none" is bit 1 rather than the absence of bits.
func mapEasServicesToSet(ctx context.Context, easServices *graphmodels.EasServices) types.Set {
	if easServices == nil {
		return types.SetNull(types.StringType)
	}

	bits := []struct {
		mask graphmodels.EasServices
		name string
	}{
		{graphmodels.NONE_EASSERVICES, "none"},
		{graphmodels.CALENDARS_EASSERVICES, "calendars"},
		{graphmodels.CONTACTS_EASSERVICES, "contacts"},
		{graphmodels.EMAIL_EASSERVICES, "email"},
		{graphmodels.NOTES_EASSERVICES, "notes"},
		{graphmodels.REMINDERS_EASSERVICES, "reminders"},
	}

	var serviceStrings []string
	for _, b := range bits {
		if (*easServices & b.mask) != 0 {
			serviceStrings = append(serviceStrings, b.name)
		}
	}

	if len(serviceStrings) == 0 {
		return types.SetNull(types.StringType)
	}

	setValue, diags := types.SetValueFrom(ctx, types.StringType, serviceStrings)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create EAS services set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(types.StringType)
	}

	return setValue
}

func mapIosScepCertificateProfile(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	prior *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosScepCertificateProfile,
) {
	tflog.Debug(ctx, "Mapping IosScepCertificateProfile")

	// Preserve the root certificate @odata.bind reference from prior state. Graph does not echo
	// additionalData back, so without this the reference is lost on every read.
	rootCertRef := types.StringNull()
	if !prior.ScepCertificate.IsNull() && !prior.ScepCertificate.IsUnknown() {
		var existingScepData ScepCertificateResourceModel
		diags := prior.ScepCertificate.As(ctx, &existingScepData, basetypes.ObjectAsOptions{})
		if !diags.HasError() && !existingScepData.RootCertificateOdataBind.IsNull() {
			rootCertRef = existingScepData.RootCertificateOdataBind
		}
	}

	// With no prior state (import), fall back to the expanded navigation property if Graph returned
	// one, so the reference is at least discoverable.
	//
	// Confirmed on a live tenant: a plain GET of a SCEP profile returns NEITHER
	// "rootCertificate@odata.bind" NOR an unexpanded "rootCertificate" navigation property — the
	// reference is simply absent. So the prior-state recovery above is load-bearing rather than
	// defensive, and this fallback only fires if the caller adds $expand=rootCertificate. It is kept
	// because it costs nothing and makes import marginally better if that is ever done.
	if rootCertRef.IsNull() {
		if rootCert := config.GetRootCertificate(); rootCert != nil {
			if id := rootCert.GetId(); id != nil {
				rootCertRef = types.StringValue(odataBindURL(*id))
			}
		}
	}

	scepModel := ScepCertificateResourceModel{
		RenewalThresholdPercentage: convert.GraphToFrameworkInt32(
			config.GetRenewalThresholdPercentage(),
		),
		CertificateStore: convert.GraphToFrameworkEnum(config.GetCertificateStore()),
		CertificateValidityPeriodScale: convert.GraphToFrameworkEnum(
			config.GetCertificateValidityPeriodScale(),
		),
		CertificateValidityPeriodValue: convert.GraphToFrameworkInt32(
			config.GetCertificateValidityPeriodValue(),
		),
		SubjectNameFormat: convert.GraphToFrameworkEnum(config.GetSubjectNameFormat()),
		SubjectNameFormatString: convert.GraphToFrameworkString(
			config.GetSubjectNameFormatString(),
		),
		SubjectAlternativeNameType: mapSubjectAlternativeNameTypeToSet(
			ctx,
			config.GetSubjectAlternativeNameType(),
		),
		SubjectAlternativeNameFormatString: convert.GraphToFrameworkString(
			config.GetSubjectAlternativeNameFormatString(),
		),
		RootCertificateOdataBind: rootCertRef,
		KeySize:                  convert.GraphToFrameworkEnum(config.GetKeySize()),
		KeyUsage:                 mapKeyUsageToSet(ctx, config.GetKeyUsage()),
		CustomSubjectAlternativeNames: mapCustomSANsToSet(
			ctx,
			config.GetCustomSubjectAlternativeNames(),
		),
		ExtendedKeyUsages: mapExtendedKeyUsagesToSet(ctx, config.GetExtendedKeyUsages()),
		ScepServerUrls:    convert.GraphToFrameworkStringSet(ctx, config.GetScepServerUrls()),
	}

	objectValue, diags := types.ObjectValueFrom(ctx, ScepCertificateType().AttrTypes, scepModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create SCEP certificate object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.ScepCertificate = objectValue
}

func mapIosPkcsCertificateProfile(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
	config *graphmodels.IosPkcsCertificateProfile,
) {
	tflog.Debug(ctx, "Mapping IosPkcsCertificateProfile")

	pkcsModel := PkcsCertificateResourceModel{
		RenewalThresholdPercentage: convert.GraphToFrameworkInt32(
			config.GetRenewalThresholdPercentage(),
		),
		CertificateStore: convert.GraphToFrameworkEnum(config.GetCertificateStore()),
		CertificateValidityPeriodScale: convert.GraphToFrameworkEnum(
			config.GetCertificateValidityPeriodScale(),
		),
		CertificateValidityPeriodValue: convert.GraphToFrameworkInt32(
			config.GetCertificateValidityPeriodValue(),
		),
		SubjectNameFormat: convert.GraphToFrameworkEnum(config.GetSubjectNameFormat()),
		SubjectNameFormatString: convert.GraphToFrameworkString(
			config.GetSubjectNameFormatString(),
		),
		SubjectAlternativeNameType: mapSubjectAlternativeNameTypeToSet(
			ctx,
			config.GetSubjectAlternativeNameType(),
		),
		SubjectAlternativeNameFormatString: convert.GraphToFrameworkString(
			config.GetSubjectAlternativeNameFormatString(),
		),
		CertificationAuthority: convert.GraphToFrameworkString(
			config.GetCertificationAuthority(),
		),
		CertificationAuthorityName: convert.GraphToFrameworkString(
			config.GetCertificationAuthorityName(),
		),
		CertificateTemplateName: convert.GraphToFrameworkString(
			config.GetCertificateTemplateName(),
		),
		CustomSubjectAlternativeNames: mapCustomSANsToSet(
			ctx,
			config.GetCustomSubjectAlternativeNames(),
		),
	}

	objectValue, diags := types.ObjectValueFrom(ctx, PkcsCertificateType().AttrTypes, pkcsModel)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create PKCS certificate object", map[string]any{
			"errors": diags.Errors(),
		})
		return
	}

	data.PkcsCertificate = objectValue
}

// Helper functions for complex mappings

// mapKeyUsageToSet decomposes the KeyUsages bitmask back into individual values.
func mapKeyUsageToSet(ctx context.Context, keyUsage *graphmodels.KeyUsages) types.Set {
	if keyUsage == nil {
		return types.SetNull(types.StringType)
	}

	var keyUsageStrings []string
	if (*keyUsage & graphmodels.KEYENCIPHERMENT_KEYUSAGES) != 0 {
		keyUsageStrings = append(keyUsageStrings, "keyEncipherment")
	}
	if (*keyUsage & graphmodels.DIGITALSIGNATURE_KEYUSAGES) != 0 {
		keyUsageStrings = append(keyUsageStrings, "digitalSignature")
	}

	if len(keyUsageStrings) == 0 {
		return types.SetNull(types.StringType)
	}

	setValue, diags := types.SetValueFrom(ctx, types.StringType, keyUsageStrings)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create key usage set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(types.StringType)
	}

	return setValue
}

// mapSubjectAlternativeNameTypeToSet decomposes the SubjectAlternativeNameType bitmask back into
// individual values. Note "none" is bit 1 rather than the absence of bits, so it round-trips as a
// real element rather than being treated as empty.
func mapSubjectAlternativeNameTypeToSet(
	ctx context.Context,
	sanType *graphmodels.SubjectAlternativeNameType,
) types.Set {
	if sanType == nil {
		return types.SetNull(types.StringType)
	}

	bits := []struct {
		mask graphmodels.SubjectAlternativeNameType
		name string
	}{
		{graphmodels.NONE_SUBJECTALTERNATIVENAMETYPE, "none"},
		{graphmodels.EMAILADDRESS_SUBJECTALTERNATIVENAMETYPE, "emailAddress"},
		{graphmodels.USERPRINCIPALNAME_SUBJECTALTERNATIVENAMETYPE, "userPrincipalName"},
		{graphmodels.CUSTOMAZUREADATTRIBUTE_SUBJECTALTERNATIVENAMETYPE, "customAzureADAttribute"},
		{graphmodels.DOMAINNAMESERVICE_SUBJECTALTERNATIVENAMETYPE, "domainNameService"},
		{
			graphmodels.UNIVERSALRESOURCEIDENTIFIER_SUBJECTALTERNATIVENAMETYPE,
			"universalResourceIdentifier",
		},
	}

	var sanTypeStrings []string
	for _, b := range bits {
		if (*sanType & b.mask) != 0 {
			sanTypeStrings = append(sanTypeStrings, b.name)
		}
	}

	if len(sanTypeStrings) == 0 {
		return types.SetNull(types.StringType)
	}

	setValue, diags := types.SetValueFrom(ctx, types.StringType, sanTypeStrings)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create subject alternative name type set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(types.StringType)
	}

	return setValue
}

func mapCustomSANsToSet(
	ctx context.Context,
	sans []graphmodels.CustomSubjectAlternativeNameable,
) types.Set {
	if len(sans) == 0 {
		return types.SetNull(customSanAttrTypes())
	}

	sanModels := make([]CustomSubjectAlternativeNameResourceModel, 0, len(sans))
	for _, san := range sans {
		if san == nil {
			continue
		}
		sanModels = append(sanModels, CustomSubjectAlternativeNameResourceModel{
			SanType: convert.GraphToFrameworkEnum(san.GetSanType()),
			Name:    convert.GraphToFrameworkString(san.GetName()),
		})
	}

	setValue, diags := types.SetValueFrom(ctx, customSanAttrTypes(), sanModels)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create custom SANs set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(customSanAttrTypes())
	}

	return setValue
}

func mapExtendedKeyUsagesToSet(
	ctx context.Context,
	ekus []graphmodels.ExtendedKeyUsageable,
) types.Set {
	if len(ekus) == 0 {
		return types.SetNull(extendedKeyUsageAttrTypes())
	}

	ekuModels := make([]ExtendedKeyUsageResourceModel, 0, len(ekus))
	for _, eku := range ekus {
		if eku == nil {
			continue
		}
		ekuModels = append(ekuModels, ExtendedKeyUsageResourceModel{
			Name:             convert.GraphToFrameworkString(eku.GetName()),
			ObjectIdentifier: convert.GraphToFrameworkString(eku.GetObjectIdentifier()),
		})
	}

	setValue, diags := types.SetValueFrom(ctx, extendedKeyUsageAttrTypes(), ekuModels)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to create extended key usages set", map[string]any{
			"errors": diags.Errors(),
		})
		return types.SetNull(extendedKeyUsageAttrTypes())
	}

	return setValue
}
