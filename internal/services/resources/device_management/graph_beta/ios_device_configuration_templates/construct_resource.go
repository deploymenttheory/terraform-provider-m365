package graphBetaIosDeviceConfigurationTemplates

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	graphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/constructors"
	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/convert"
)

var (
	errNoConfigurationType = errors.New(
		"no configuration type specified: exactly one of general_device_configuration, " +
			"custom_configuration, trusted_certificate, wifi, scep_certificate, pkcs_certificate, " +
			"enterprise_wifi, eas_email, or vpn is required",
	)
	errConstructConfiguration    = errors.New("failed to construct configuration")
	errSetRoleScopeTagIds        = errors.New("failed to set role scope tags")
	errExtractKeyUsage           = errors.New("failed to extract key usage values")
	errInvalidKeyUsage           = errors.New("invalid key usage value")
	errExtractSanType            = errors.New("failed to extract subject alternative name types")
	errInvalidSanType            = errors.New("invalid subject alternative name type")
	errExtractCustomSans         = errors.New("failed to extract custom subject alternative names")
	errExtractExtendedKeyUsages  = errors.New("failed to extract extended key usages")
	errExtractEasServices        = errors.New("failed to extract EAS services")
	errInvalidEasService         = errors.New("invalid EAS service value")
	errExtractVpnServer          = errors.New("failed to extract VPN server")
	errExtractVpnProxyServer     = errors.New("failed to extract VPN proxy server")
	errExtractTargetedApps       = errors.New("failed to extract targeted mobile apps")
	errExtractCustomData         = errors.New("failed to extract custom data")
	errExtractCustomKeyValueData = errors.New("failed to extract custom key value data")
)

// Main entry point to construct the iOS/iPadOS configuration template resource for the Terraform provider.
func constructResource(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, fmt.Sprintf("Constructing %s resource from model", ResourceName))

	var requestBody graphmodels.DeviceConfigurationable

	// Determine which configuration type to construct. ModifyPlan already guarantees exactly one
	// block is set, so the default arm is a safety net rather than the normal validation path.
	var err error
	switch {
	case !data.GeneralDeviceConfiguration.IsNull() && !data.GeneralDeviceConfiguration.IsUnknown():
		requestBody = constructIosGeneralDeviceConfiguration(ctx)
	case !data.CustomConfiguration.IsNull() && !data.CustomConfiguration.IsUnknown():
		requestBody, err = constructIosCustomConfiguration(ctx, data)
	case !data.TrustedCertificate.IsNull() && !data.TrustedCertificate.IsUnknown():
		requestBody, err = constructIosTrustedRootCertificate(ctx, data)
	case !data.Wifi.IsNull() && !data.Wifi.IsUnknown():
		requestBody, err = constructIosWiFiConfiguration(ctx, data)
	case !data.ScepCertificate.IsNull() && !data.ScepCertificate.IsUnknown():
		requestBody, err = constructIosScepCertificateProfile(ctx, data)
	case !data.PkcsCertificate.IsNull() && !data.PkcsCertificate.IsUnknown():
		requestBody, err = constructIosPkcsCertificateProfile(ctx, data)
	case !data.EnterpriseWifi.IsNull() && !data.EnterpriseWifi.IsUnknown():
		requestBody, err = constructIosEnterpriseWiFiConfiguration(ctx, data)
	case !data.EasEmail.IsNull() && !data.EasEmail.IsUnknown():
		requestBody, err = constructIosEasEmailProfileConfiguration(ctx, data)
	case !data.Vpn.IsNull() && !data.Vpn.IsUnknown():
		requestBody, err = constructIosVpnConfiguration(ctx, data)
	default:
		return nil, errNoConfigurationType
	}

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errConstructConfiguration, err)
	}

	// Set common properties
	convert.FrameworkToGraphString(data.DisplayName, requestBody.SetDisplayName)
	convert.FrameworkToGraphString(data.Description, requestBody.SetDescription)

	if err := convert.FrameworkToGraphStringSet(
		ctx,
		data.RoleScopeTagIds,
		requestBody.SetRoleScopeTagIds,
	); err != nil {
		return nil, fmt.Errorf("%w: %w", errSetRoleScopeTagIds, err)
	}

	if err := constructors.DebugLogGraphObject(
		ctx,
		fmt.Sprintf("Final JSON to be sent to Graph API for resource %s", ResourceName),
		requestBody,
	); err != nil {
		tflog.Error(ctx, "Failed to debug log object", map[string]any{
			"error": err.Error(),
		})
	}

	tflog.Debug(ctx, fmt.Sprintf("Finished constructing %s resource", ResourceName))

	return requestBody, nil
}

// constructIosGeneralDeviceConfiguration constructs an IosGeneralDeviceConfiguration.
// The base fields (display_name, description, roleScopeTagIds) are set by the shared header in
// constructResource; the type-specific properties are intentionally left unset here and can be
// managed via the JSON resource or the Intune portal.
func constructIosGeneralDeviceConfiguration(
	ctx context.Context,
) graphmodels.DeviceConfigurationable {
	tflog.Debug(ctx, "Constructing IosGeneralDeviceConfiguration")
	return graphmodels.NewIosGeneralDeviceConfiguration()
}

// constructIosCustomConfiguration constructs an IosCustomConfiguration
func constructIosCustomConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosCustomConfiguration")

	customConfig := graphmodels.NewIosCustomConfiguration()

	var customConfigData CustomConfigurationResourceModel
	diags := data.CustomConfiguration.As(ctx, &customConfigData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract custom configuration data: %v", diags.Errors())
	}

	convert.FrameworkToGraphString(
		customConfigData.PayloadFileName,
		customConfig.SetPayloadFileName,
	)
	convert.FrameworkToGraphBytes(customConfigData.Payload, customConfig.SetPayload)
	convert.FrameworkToGraphString(customConfigData.PayloadName, customConfig.SetPayloadName)

	return customConfig, nil
}

// constructIosTrustedRootCertificate constructs an IosTrustedRootCertificate
func constructIosTrustedRootCertificate(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosTrustedRootCertificate")

	certConfig := graphmodels.NewIosTrustedRootCertificate()

	var certData TrustedCertificateResourceModel
	diags := data.TrustedCertificate.As(ctx, &certData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract trusted certificate data: %v", diags.Errors())
	}

	convert.FrameworkToGraphString(certData.CertFileName, certConfig.SetCertFileName)

	// Handle base64-encoded certificate data from filebase64()
	if !certData.TrustedRootCertificate.IsNull() && !certData.TrustedRootCertificate.IsUnknown() {
		certBase64 := certData.TrustedRootCertificate.ValueString()
		certBytes, err := base64.StdEncoding.DecodeString(certBase64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64 certificate data: %w", err)
		}
		certConfig.SetTrustedRootCertificate(certBytes)
	}

	return certConfig, nil
}

// constructIosWiFiConfiguration constructs an IosWiFiConfiguration
func constructIosWiFiConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosWiFiConfiguration")

	wifiConfig := graphmodels.NewIosWiFiConfiguration()

	var wifiData WifiResourceModel
	diags := data.Wifi.As(ctx, &wifiData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract wifi data: %v", diags.Errors())
	}

	convert.FrameworkToGraphString(wifiData.NetworkName, wifiConfig.SetNetworkName)
	convert.FrameworkToGraphString(wifiData.Ssid, wifiConfig.SetSsid)
	convert.FrameworkToGraphBool(wifiData.ConnectAutomatically, wifiConfig.SetConnectAutomatically)
	convert.FrameworkToGraphBool(
		wifiData.ConnectWhenNetworkNameIsHidden,
		wifiConfig.SetConnectWhenNetworkNameIsHidden,
	)

	if err := convert.FrameworkToGraphEnum(
		wifiData.WifiSecurityType,
		graphmodels.ParseWiFiSecurityType,
		wifiConfig.SetWiFiSecurityType,
	); err != nil {
		return nil, fmt.Errorf("failed to set wifi security type: %w", err)
	}

	convert.FrameworkToGraphString(wifiData.PreSharedKey, wifiConfig.SetPreSharedKey)
	convert.FrameworkToGraphBool(
		wifiData.DisableMacAddressRandomization,
		wifiConfig.SetDisableMacAddressRandomization,
	)

	if err := convert.FrameworkToGraphEnum(
		wifiData.ProxySettings,
		graphmodels.ParseWiFiProxySetting,
		wifiConfig.SetProxySettings,
	); err != nil {
		return nil, fmt.Errorf("failed to set proxy settings: %w", err)
	}

	convert.FrameworkToGraphString(wifiData.ProxyManualAddress, wifiConfig.SetProxyManualAddress)
	convert.FrameworkToGraphInt32(wifiData.ProxyManualPort, wifiConfig.SetProxyManualPort)
	convert.FrameworkToGraphString(
		wifiData.ProxyAutomaticConfigurationUrl,
		wifiConfig.SetProxyAutomaticConfigurationUrl,
	)

	return wifiConfig, nil
}

// constructIosScepCertificateProfile constructs an IosScepCertificateProfile
func constructIosScepCertificateProfile(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosScepCertificateProfile")

	scepConfig := graphmodels.NewIosScepCertificateProfile()

	var scepData ScepCertificateResourceModel
	diags := data.ScepCertificate.As(ctx, &scepData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract SCEP certificate data: %v", diags.Errors())
	}

	convert.FrameworkToGraphInt32(
		scepData.RenewalThresholdPercentage,
		scepConfig.SetRenewalThresholdPercentage,
	)

	if err := convert.FrameworkToGraphEnum(
		scepData.CertificateStore,
		graphmodels.ParseCertificateStore,
		scepConfig.SetCertificateStore,
	); err != nil {
		return nil, fmt.Errorf("failed to set certificate store: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		scepData.CertificateValidityPeriodScale,
		graphmodels.ParseCertificateValidityPeriodScale,
		scepConfig.SetCertificateValidityPeriodScale,
	); err != nil {
		return nil, fmt.Errorf("failed to set certificate validity period scale: %w", err)
	}

	convert.FrameworkToGraphInt32(
		scepData.CertificateValidityPeriodValue,
		scepConfig.SetCertificateValidityPeriodValue,
	)

	if err := convert.FrameworkToGraphEnum(
		scepData.SubjectNameFormat,
		graphmodels.ParseAppleSubjectNameFormat,
		scepConfig.SetSubjectNameFormat,
	); err != nil {
		return nil, fmt.Errorf("failed to set subject name format: %w", err)
	}

	convert.FrameworkToGraphString(
		scepData.SubjectNameFormatString,
		scepConfig.SetSubjectNameFormatString,
	)

	if err := setSubjectAlternativeNameType(
		ctx,
		scepData.SubjectAlternativeNameType,
		scepConfig.SetSubjectAlternativeNameType,
	); err != nil {
		return nil, fmt.Errorf("failed to set subject alternative name type: %w", err)
	}

	convert.FrameworkToGraphString(
		scepData.SubjectAlternativeNameFormatString,
		scepConfig.SetSubjectAlternativeNameFormatString,
	)

	// The root certificate is a navigation property, set as an @odata.bind reference rather than a
	// nested object. SetAdditionalData replaces the whole map, so this must be the only additional
	// data this block contributes.
	if !scepData.RootCertificateOdataBind.IsNull() &&
		!scepData.RootCertificateOdataBind.IsUnknown() {
		scepConfig.SetAdditionalData(map[string]any{
			"rootCertificate@odata.bind": normalizeODataBind(
				scepData.RootCertificateOdataBind.ValueString(),
			),
		})
	}

	if err := convert.FrameworkToGraphEnum(
		scepData.KeySize,
		graphmodels.ParseKeySize,
		scepConfig.SetKeySize,
	); err != nil {
		return nil, fmt.Errorf("failed to set key size: %w", err)
	}

	if err := setKeyUsage(ctx, scepData.KeyUsage, scepConfig.SetKeyUsage); err != nil {
		return nil, fmt.Errorf("failed to set key usage: %w", err)
	}

	if err := convertCustomSubjectAlternativeNames(
		ctx,
		scepData.CustomSubjectAlternativeNames,
		scepConfig.SetCustomSubjectAlternativeNames,
	); err != nil {
		return nil, fmt.Errorf("failed to set custom subject alternative names: %w", err)
	}

	if err := convertExtendedKeyUsages(
		ctx,
		scepData.ExtendedKeyUsages,
		scepConfig.SetExtendedKeyUsages,
	); err != nil {
		return nil, fmt.Errorf("failed to set extended key usages: %w", err)
	}

	if err := convert.FrameworkToGraphStringSet(
		ctx,
		scepData.ScepServerUrls,
		scepConfig.SetScepServerUrls,
	); err != nil {
		return nil, fmt.Errorf("failed to set SCEP server URLs: %w", err)
	}

	return scepConfig, nil
}

// constructIosPkcsCertificateProfile constructs an IosPkcsCertificateProfile
func constructIosPkcsCertificateProfile(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosPkcsCertificateProfile")

	pkcsConfig := graphmodels.NewIosPkcsCertificateProfile()

	var pkcsData PkcsCertificateResourceModel
	diags := data.PkcsCertificate.As(ctx, &pkcsData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract PKCS certificate data: %v", diags.Errors())
	}

	convert.FrameworkToGraphInt32(
		pkcsData.RenewalThresholdPercentage,
		pkcsConfig.SetRenewalThresholdPercentage,
	)

	if err := convert.FrameworkToGraphEnum(
		pkcsData.CertificateStore,
		graphmodels.ParseCertificateStore,
		pkcsConfig.SetCertificateStore,
	); err != nil {
		return nil, fmt.Errorf("failed to set certificate store: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		pkcsData.CertificateValidityPeriodScale,
		graphmodels.ParseCertificateValidityPeriodScale,
		pkcsConfig.SetCertificateValidityPeriodScale,
	); err != nil {
		return nil, fmt.Errorf("failed to set certificate validity period scale: %w", err)
	}

	convert.FrameworkToGraphInt32(
		pkcsData.CertificateValidityPeriodValue,
		pkcsConfig.SetCertificateValidityPeriodValue,
	)

	if err := convert.FrameworkToGraphEnum(
		pkcsData.SubjectNameFormat,
		graphmodels.ParseAppleSubjectNameFormat,
		pkcsConfig.SetSubjectNameFormat,
	); err != nil {
		return nil, fmt.Errorf("failed to set subject name format: %w", err)
	}

	convert.FrameworkToGraphString(
		pkcsData.SubjectNameFormatString,
		pkcsConfig.SetSubjectNameFormatString,
	)

	if err := setSubjectAlternativeNameType(
		ctx,
		pkcsData.SubjectAlternativeNameType,
		pkcsConfig.SetSubjectAlternativeNameType,
	); err != nil {
		return nil, fmt.Errorf("failed to set subject alternative name type: %w", err)
	}

	convert.FrameworkToGraphString(
		pkcsData.SubjectAlternativeNameFormatString,
		pkcsConfig.SetSubjectAlternativeNameFormatString,
	)

	convert.FrameworkToGraphString(
		pkcsData.CertificationAuthority,
		pkcsConfig.SetCertificationAuthority,
	)
	convert.FrameworkToGraphString(
		pkcsData.CertificationAuthorityName,
		pkcsConfig.SetCertificationAuthorityName,
	)
	convert.FrameworkToGraphString(
		pkcsData.CertificateTemplateName,
		pkcsConfig.SetCertificateTemplateName,
	)

	if err := convertCustomSubjectAlternativeNames(
		ctx,
		pkcsData.CustomSubjectAlternativeNames,
		pkcsConfig.SetCustomSubjectAlternativeNames,
	); err != nil {
		return nil, fmt.Errorf("failed to set custom subject alternative names: %w", err)
	}

	return pkcsConfig, nil
}

// constructIosEnterpriseWiFiConfiguration constructs an IosEnterpriseWiFiConfiguration
func constructIosEnterpriseWiFiConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosEnterpriseWiFiConfiguration")

	wifiConfig := graphmodels.NewIosEnterpriseWiFiConfiguration()

	var wifiData EnterpriseWifiResourceModel
	diags := data.EnterpriseWifi.As(ctx, &wifiData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract enterprise wifi data: %v", diags.Errors())
	}

	// Inherited from IosWiFiConfiguration. preSharedKey is inherited too but intentionally not
	// exposed: enterprise networks authenticate via EAP.
	convert.FrameworkToGraphString(wifiData.NetworkName, wifiConfig.SetNetworkName)
	convert.FrameworkToGraphString(wifiData.Ssid, wifiConfig.SetSsid)
	convert.FrameworkToGraphBool(wifiData.ConnectAutomatically, wifiConfig.SetConnectAutomatically)
	convert.FrameworkToGraphBool(
		wifiData.ConnectWhenNetworkNameIsHidden,
		wifiConfig.SetConnectWhenNetworkNameIsHidden,
	)

	if err := convert.FrameworkToGraphEnum(
		wifiData.WifiSecurityType,
		graphmodels.ParseWiFiSecurityType,
		wifiConfig.SetWiFiSecurityType,
	); err != nil {
		return nil, fmt.Errorf("failed to set wifi security type: %w", err)
	}

	convert.FrameworkToGraphBool(
		wifiData.DisableMacAddressRandomization,
		wifiConfig.SetDisableMacAddressRandomization,
	)

	if err := convert.FrameworkToGraphEnum(
		wifiData.ProxySettings,
		graphmodels.ParseWiFiProxySetting,
		wifiConfig.SetProxySettings,
	); err != nil {
		return nil, fmt.Errorf("failed to set proxy settings: %w", err)
	}

	convert.FrameworkToGraphString(wifiData.ProxyManualAddress, wifiConfig.SetProxyManualAddress)
	convert.FrameworkToGraphInt32(wifiData.ProxyManualPort, wifiConfig.SetProxyManualPort)
	convert.FrameworkToGraphString(
		wifiData.ProxyAutomaticConfigurationUrl,
		wifiConfig.SetProxyAutomaticConfigurationUrl,
	)

	// Enterprise-specific
	if err := convert.FrameworkToGraphEnum(
		wifiData.EapType,
		graphmodels.ParseEapType,
		wifiConfig.SetEapType,
	); err != nil {
		return nil, fmt.Errorf("failed to set EAP type: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		wifiData.EapFastConfiguration,
		graphmodels.ParseEapFastConfiguration,
		wifiConfig.SetEapFastConfiguration,
	); err != nil {
		return nil, fmt.Errorf("failed to set EAP-FAST configuration: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		wifiData.AuthenticationMethod,
		graphmodels.ParseWiFiAuthenticationMethod,
		wifiConfig.SetAuthenticationMethod,
	); err != nil {
		return nil, fmt.Errorf("failed to set authentication method: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		wifiData.InnerAuthenticationProtocolForEapTtls,
		graphmodels.ParseNonEapAuthenticationMethodForEapTtlsType,
		wifiConfig.SetInnerAuthenticationProtocolForEapTtls,
	); err != nil {
		return nil, fmt.Errorf("failed to set inner authentication protocol: %w", err)
	}

	convert.FrameworkToGraphString(
		wifiData.OuterIdentityPrivacyTemporaryValue,
		wifiConfig.SetOuterIdentityPrivacyTemporaryValue,
	)
	convert.FrameworkToGraphString(
		wifiData.UsernameFormatString,
		wifiConfig.SetUsernameFormatString,
	)
	convert.FrameworkToGraphString(
		wifiData.PasswordFormatString,
		wifiConfig.SetPasswordFormatString,
	)

	if err := convert.FrameworkToGraphStringSet(
		ctx,
		wifiData.TrustedServerCertificateNames,
		wifiConfig.SetTrustedServerCertificateNames,
	); err != nil {
		return nil, fmt.Errorf("failed to set trusted server certificate names: %w", err)
	}

	// Both certificate references are navigation properties set via @odata.bind. Accumulate them
	// into one map: SetAdditionalData replaces the whole map rather than merging.
	additionalData := map[string]any{}

	if !wifiData.RootCertificatesForServerValidationOdataBind.IsNull() &&
		!wifiData.RootCertificatesForServerValidationOdataBind.IsUnknown() {
		var rootCertRefs []string
		if diags := wifiData.RootCertificatesForServerValidationOdataBind.ElementsAs(
			ctx, &rootCertRefs, false,
		); diags.HasError() {
			return nil, fmt.Errorf("failed to extract root certificate references: %v", diags.Errors())
		}
		if len(rootCertRefs) > 0 {
			normalized := make([]string, 0, len(rootCertRefs))
			for _, ref := range rootCertRefs {
				normalized = append(normalized, normalizeODataBind(ref))
			}
			// Kiota's WriteAdditionalData has an explicit []string case, so this serializes as a
			// JSON array rather than being dropped.
			additionalData["rootCertificatesForServerValidation@odata.bind"] = normalized
		}
	}

	if !wifiData.IdentityCertificateForClientAuthenticationOdataBind.IsNull() &&
		!wifiData.IdentityCertificateForClientAuthenticationOdataBind.IsUnknown() {
		additionalData["identityCertificateForClientAuthentication@odata.bind"] = normalizeODataBind(
			wifiData.IdentityCertificateForClientAuthenticationOdataBind.ValueString(),
		)
	}

	if len(additionalData) > 0 {
		wifiConfig.SetAdditionalData(additionalData)
	}

	return wifiConfig, nil
}

// constructIosEasEmailProfileConfiguration constructs an IosEasEmailProfileConfiguration
func constructIosEasEmailProfileConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosEasEmailProfileConfiguration")

	easConfig := graphmodels.NewIosEasEmailProfileConfiguration()

	var easData EasEmailResourceModel
	diags := data.EasEmail.As(ctx, &easData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract EAS email data: %v", diags.Errors())
	}

	convert.FrameworkToGraphString(easData.AccountName, easConfig.SetAccountName)
	convert.FrameworkToGraphString(easData.HostName, easConfig.SetHostName)

	if err := convert.FrameworkToGraphEnum(
		easData.AuthenticationMethod,
		graphmodels.ParseEasAuthenticationMethod,
		easConfig.SetAuthenticationMethod,
	); err != nil {
		return nil, fmt.Errorf("failed to set authentication method: %w", err)
	}

	if err := setEasServices(ctx, easData.EasServices, easConfig.SetEasServices); err != nil {
		return nil, fmt.Errorf("failed to set EAS services: %w", err)
	}

	convert.FrameworkToGraphBool(
		easData.EasServicesUserOverrideEnabled,
		easConfig.SetEasServicesUserOverrideEnabled,
	)

	if err := convert.FrameworkToGraphEnum(
		easData.DurationOfEmailToSync,
		graphmodels.ParseEmailSyncDuration,
		easConfig.SetDurationOfEmailToSync,
	); err != nil {
		return nil, fmt.Errorf("failed to set email sync duration: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		easData.EmailAddressSource,
		graphmodels.ParseUserEmailSource,
		easConfig.SetEmailAddressSource,
	); err != nil {
		return nil, fmt.Errorf("failed to set email address source: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		easData.UsernameSource,
		graphmodels.ParseUserEmailSource,
		easConfig.SetUsernameSource,
	); err != nil {
		return nil, fmt.Errorf("failed to set username source: %w", err)
	}

	// Note this is a different enum from UsernameSource above: UsernameSource (AAD) additionally
	// allows samAccountName.
	if err := convert.FrameworkToGraphEnum(
		easData.UsernameAADSource,
		graphmodels.ParseUsernameSource,
		easConfig.SetUsernameAADSource,
	); err != nil {
		return nil, fmt.Errorf("failed to set username AAD source: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		easData.UserDomainNameSource,
		graphmodels.ParseDomainNameSource,
		easConfig.SetUserDomainNameSource,
	); err != nil {
		return nil, fmt.Errorf("failed to set user domain name source: %w", err)
	}

	convert.FrameworkToGraphString(easData.CustomDomainName, easConfig.SetCustomDomainName)
	convert.FrameworkToGraphBool(easData.RequireSsl, easConfig.SetRequireSsl)
	convert.FrameworkToGraphBool(easData.UseOAuth, easConfig.SetUseOAuth)
	convert.FrameworkToGraphString(easData.PerAppVPNProfileId, easConfig.SetPerAppVPNProfileId)

	convert.FrameworkToGraphBool(
		easData.BlockMovingMessagesToOtherEmailAccounts,
		easConfig.SetBlockMovingMessagesToOtherEmailAccounts,
	)
	convert.FrameworkToGraphBool(
		easData.BlockSendingEmailFromThirdPartyApps,
		easConfig.SetBlockSendingEmailFromThirdPartyApps,
	)
	convert.FrameworkToGraphBool(
		easData.BlockSyncingRecentlyUsedEmailAddresses,
		easConfig.SetBlockSyncingRecentlyUsedEmailAddresses,
	)

	convert.FrameworkToGraphBool(easData.RequireSmime, easConfig.SetRequireSmime)
	convert.FrameworkToGraphBool(
		easData.SmimeEnablePerMessageSwitch,
		easConfig.SetSmimeEnablePerMessageSwitch,
	)
	convert.FrameworkToGraphBool(easData.SmimeSigningEnabled, easConfig.SetSmimeSigningEnabled)
	convert.FrameworkToGraphBool(
		easData.SmimeSigningUserOverrideEnabled,
		easConfig.SetSmimeSigningUserOverrideEnabled,
	)
	convert.FrameworkToGraphBool(
		easData.SmimeSigningCertificateUserOverrideEnabled,
		easConfig.SetSmimeSigningCertificateUserOverrideEnabled,
	)
	convert.FrameworkToGraphBool(
		easData.SmimeEncryptByDefaultEnabled,
		easConfig.SetSmimeEncryptByDefaultEnabled,
	)
	convert.FrameworkToGraphBool(
		easData.SmimeEncryptByDefaultUserOverrideEnabled,
		easConfig.SetSmimeEncryptByDefaultUserOverrideEnabled,
	)
	convert.FrameworkToGraphBool(
		easData.SmimeEncryptionCertificateUserOverrideEnabled,
		easConfig.SetSmimeEncryptionCertificateUserOverrideEnabled,
	)

	if err := convert.FrameworkToGraphEnum(
		easData.SigningCertificateType,
		graphmodels.ParseEmailCertificateType,
		easConfig.SetSigningCertificateType,
	); err != nil {
		return nil, fmt.Errorf("failed to set signing certificate type: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		easData.EncryptionCertificateType,
		graphmodels.ParseEmailCertificateType,
		easConfig.SetEncryptionCertificateType,
	); err != nil {
		return nil, fmt.Errorf("failed to set encryption certificate type: %w", err)
	}

	// All three certificate references are navigation properties. Accumulate before setting, since
	// SetAdditionalData replaces the whole map.
	additionalData := map[string]any{}
	if !easData.IdentityCertificateOdataBind.IsNull() &&
		!easData.IdentityCertificateOdataBind.IsUnknown() {
		additionalData["identityCertificate@odata.bind"] = normalizeODataBind(
			easData.IdentityCertificateOdataBind.ValueString(),
		)
	}
	if !easData.SmimeSigningCertificateOdataBind.IsNull() &&
		!easData.SmimeSigningCertificateOdataBind.IsUnknown() {
		additionalData["smimeSigningCertificate@odata.bind"] = normalizeODataBind(
			easData.SmimeSigningCertificateOdataBind.ValueString(),
		)
	}
	if !easData.SmimeEncryptionCertificateOdataBind.IsNull() &&
		!easData.SmimeEncryptionCertificateOdataBind.IsUnknown() {
		additionalData["smimeEncryptionCertificate@odata.bind"] = normalizeODataBind(
			easData.SmimeEncryptionCertificateOdataBind.ValueString(),
		)
	}
	if len(additionalData) > 0 {
		easConfig.SetAdditionalData(additionalData)
	}

	return easConfig, nil
}

// constructIosVpnConfiguration constructs an IosVpnConfiguration
func constructIosVpnConfiguration(
	ctx context.Context,
	data *IosDeviceConfigurationTemplatesResourceModel,
) (graphmodels.DeviceConfigurationable, error) {
	tflog.Debug(ctx, "Constructing IosVpnConfiguration")

	vpnConfig := graphmodels.NewIosVpnConfiguration()

	var vpnData VpnResourceModel
	diags := data.Vpn.As(ctx, &vpnData, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return nil, fmt.Errorf("failed to extract VPN data: %v", diags.Errors())
	}

	convert.FrameworkToGraphString(vpnData.ConnectionName, vpnConfig.SetConnectionName)

	if err := convert.FrameworkToGraphEnum(
		vpnData.ConnectionType,
		graphmodels.ParseAppleVpnConnectionType,
		vpnConfig.SetConnectionType,
	); err != nil {
		return nil, fmt.Errorf("failed to set connection type: %w", err)
	}

	if err := convert.FrameworkToGraphEnum(
		vpnData.AuthenticationMethod,
		graphmodels.ParseVpnAuthenticationMethod,
		vpnConfig.SetAuthenticationMethod,
	); err != nil {
		return nil, fmt.Errorf("failed to set authentication method: %w", err)
	}

	convert.FrameworkToGraphString(vpnData.Identifier, vpnConfig.SetIdentifier)

	if err := convert.FrameworkToGraphEnum(
		vpnData.ProviderType,
		graphmodels.ParseVpnProviderType,
		vpnConfig.SetProviderType,
	); err != nil {
		return nil, fmt.Errorf("failed to set provider type: %w", err)
	}

	convert.FrameworkToGraphString(vpnData.Realm, vpnConfig.SetRealm)
	convert.FrameworkToGraphString(vpnData.Role, vpnConfig.SetRole)
	convert.FrameworkToGraphString(vpnData.LoginGroupOrDomain, vpnConfig.SetLoginGroupOrDomain)

	convert.FrameworkToGraphBool(vpnData.EnableSplitTunneling, vpnConfig.SetEnableSplitTunneling)
	convert.FrameworkToGraphBool(vpnData.EnablePerApp, vpnConfig.SetEnablePerApp)
	convert.FrameworkToGraphBool(vpnData.IncludeAllNetworks, vpnConfig.SetIncludeAllNetworks)
	convert.FrameworkToGraphBool(vpnData.ExcludeLocalNetworks, vpnConfig.SetExcludeLocalNetworks)

	for _, set := range []struct {
		value  types.Set
		setter func([]string)
		label  string
	}{
		{vpnData.SafariDomains, vpnConfig.SetSafariDomains, "safari domains"},
		{vpnData.AssociatedDomains, vpnConfig.SetAssociatedDomains, "associated domains"},
		{vpnData.ExcludedDomains, vpnConfig.SetExcludedDomains, "excluded domains"},
		{vpnData.ExcludeList, vpnConfig.SetExcludeList, "exclude list"},
	} {
		if err := convert.FrameworkToGraphStringSet(ctx, set.value, set.setter); err != nil {
			return nil, fmt.Errorf("failed to set %s: %w", set.label, err)
		}
	}

	convert.FrameworkToGraphBool(vpnData.DisconnectOnIdle, vpnConfig.SetDisconnectOnIdle)
	convert.FrameworkToGraphInt32(
		vpnData.DisconnectOnIdleTimerInSeconds,
		vpnConfig.SetDisconnectOnIdleTimerInSeconds,
	)
	convert.FrameworkToGraphBool(
		vpnData.DisableOnDemandUserOverride,
		vpnConfig.SetDisableOnDemandUserOverride,
	)
	convert.FrameworkToGraphBool(
		vpnData.OptInToDeviceIdSharing,
		vpnConfig.SetOptInToDeviceIdSharing,
	)

	convert.FrameworkToGraphString(
		vpnData.MicrosoftTunnelSiteId,
		vpnConfig.SetMicrosoftTunnelSiteId,
	)
	convert.FrameworkToGraphBool(vpnData.StrictEnforcement, vpnConfig.SetStrictEnforcement)

	convert.FrameworkToGraphString(vpnData.CloudName, vpnConfig.SetCloudName)
	convert.FrameworkToGraphString(vpnData.UserDomain, vpnConfig.SetUserDomain)

	if err := setVpnServer(ctx, vpnData.Server, vpnConfig.SetServer); err != nil {
		return nil, fmt.Errorf("failed to set VPN server: %w", err)
	}

	if err := setVpnProxyServer(ctx, vpnData.ProxyServer, vpnConfig.SetProxyServer); err != nil {
		return nil, fmt.Errorf("failed to set VPN proxy server: %w", err)
	}

	if err := setTargetedMobileApps(
		ctx,
		vpnData.TargetedMobileApps,
		vpnConfig.SetTargetedMobileApps,
	); err != nil {
		return nil, fmt.Errorf("failed to set targeted mobile apps: %w", err)
	}

	if err := setCustomData(ctx, vpnData.CustomData, vpnConfig.SetCustomData); err != nil {
		return nil, fmt.Errorf("failed to set custom data: %w", err)
	}

	if err := setCustomKeyValueData(
		ctx,
		vpnData.CustomKeyValueData,
		vpnConfig.SetCustomKeyValueData,
	); err != nil {
		return nil, fmt.Errorf("failed to set custom key value data: %w", err)
	}

	if !vpnData.IdentityCertificateOdataBind.IsNull() &&
		!vpnData.IdentityCertificateOdataBind.IsUnknown() {
		vpnConfig.SetAdditionalData(map[string]any{
			"identityCertificate@odata.bind": normalizeODataBind(
				vpnData.IdentityCertificateOdataBind.ValueString(),
			),
		})
	}

	return vpnConfig, nil
}

// Helper functions for complex conversions

// setVpnServer builds the nested vpnServer object.
func setVpnServer(
	ctx context.Context,
	serverObj types.Object,
	setter func(graphmodels.VpnServerable),
) error {
	if serverObj.IsNull() || serverObj.IsUnknown() {
		return nil
	}

	var serverData VpnServerResourceModel
	if diags := serverObj.As(ctx, &serverData, basetypes.ObjectAsOptions{}); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractVpnServer, diags.Errors())
	}

	server := graphmodels.NewVpnServer()
	convert.FrameworkToGraphString(serverData.Address, server.SetAddress)
	convert.FrameworkToGraphString(serverData.Description, server.SetDescription)
	convert.FrameworkToGraphBool(serverData.IsDefaultServer, server.SetIsDefaultServer)

	setter(server)
	return nil
}

// setVpnProxyServer builds the nested vpnProxyServer object.
func setVpnProxyServer(
	ctx context.Context,
	proxyObj types.Object,
	setter func(graphmodels.VpnProxyServerable),
) error {
	if proxyObj.IsNull() || proxyObj.IsUnknown() {
		return nil
	}

	var proxyData VpnProxyServerResourceModel
	if diags := proxyObj.As(ctx, &proxyData, basetypes.ObjectAsOptions{}); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractVpnProxyServer, diags.Errors())
	}

	proxy := graphmodels.NewVpnProxyServer()
	convert.FrameworkToGraphString(proxyData.Address, proxy.SetAddress)
	convert.FrameworkToGraphInt32(proxyData.Port, proxy.SetPort)
	convert.FrameworkToGraphString(
		proxyData.AutomaticConfigurationScriptUrl,
		proxy.SetAutomaticConfigurationScriptUrl,
	)

	setter(proxy)
	return nil
}

func setTargetedMobileApps(
	ctx context.Context,
	appSet types.Set,
	setter func([]graphmodels.AppListItemable),
) error {
	if appSet.IsNull() || appSet.IsUnknown() {
		return nil
	}

	var appModels []AppListItemResourceModel
	if diags := appSet.ElementsAs(ctx, &appModels, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractTargetedApps, diags.Errors())
	}

	apps := make([]graphmodels.AppListItemable, 0, len(appModels))
	for _, appModel := range appModels {
		app := graphmodels.NewAppListItem()
		convert.FrameworkToGraphString(appModel.Name, app.SetName)
		convert.FrameworkToGraphString(appModel.AppId, app.SetAppId)
		convert.FrameworkToGraphString(appModel.Publisher, app.SetPublisher)
		convert.FrameworkToGraphString(appModel.AppStoreUrl, app.SetAppStoreUrl)
		apps = append(apps, app)
	}

	setter(apps)
	return nil
}

// setCustomData populates Graph's customData property, whose entries use a "key" field.
func setCustomData(
	ctx context.Context,
	kvSet types.Set,
	setter func([]graphmodels.KeyValueable),
) error {
	if kvSet.IsNull() || kvSet.IsUnknown() {
		return nil
	}

	var kvModels []KeyValueResourceModel
	if diags := kvSet.ElementsAs(ctx, &kvModels, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractCustomData, diags.Errors())
	}

	kvs := make([]graphmodels.KeyValueable, 0, len(kvModels))
	for _, kvModel := range kvModels {
		kv := graphmodels.NewKeyValue()
		convert.FrameworkToGraphString(kvModel.Key, kv.SetKey)
		convert.FrameworkToGraphString(kvModel.Value, kv.SetValue)
		kvs = append(kvs, kv)
	}

	setter(kvs)
	return nil
}

// setCustomKeyValueData populates Graph's customKeyValueData property, whose entries use a "name"
// field rather than "key" — a distinct property from customData above.
func setCustomKeyValueData(
	ctx context.Context,
	kvSet types.Set,
	setter func([]graphmodels.KeyValuePairable),
) error {
	if kvSet.IsNull() || kvSet.IsUnknown() {
		return nil
	}

	var kvModels []KeyValuePairResourceModel
	if diags := kvSet.ElementsAs(ctx, &kvModels, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractCustomKeyValueData, diags.Errors())
	}

	kvs := make([]graphmodels.KeyValuePairable, 0, len(kvModels))
	for _, kvModel := range kvModels {
		kv := graphmodels.NewKeyValuePair()
		convert.FrameworkToGraphString(kvModel.Name, kv.SetName)
		convert.FrameworkToGraphString(kvModel.Value, kv.SetValue)
		kvs = append(kvs, kv)
	}

	setter(kvs)
	return nil
}

// normalizeODataBind accepts either a bare GUID or a full Graph URL and returns the full URL form
// required by an @odata.bind reference.
func normalizeODataBind(value string) string {
	if strings.HasPrefix(value, "https://") {
		return value
	}
	return fmt.Sprintf(
		"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('%s')",
		value,
	)
}

// setKeyUsage combines the configured key usages into the single bitmask value Graph expects.
// KeyUsages is a flags enum (keyEncipherment=1, digitalSignature=2) serialized as a comma-joined
// string, so the individual values must be OR-ed together rather than sent as a list.
func setKeyUsage(
	ctx context.Context,
	keyUsageSet types.Set,
	setter func(*graphmodels.KeyUsages),
) error {
	if keyUsageSet.IsNull() || keyUsageSet.IsUnknown() {
		return nil
	}

	var keyUsageStrings []string
	if diags := keyUsageSet.ElementsAs(ctx, &keyUsageStrings, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractKeyUsage, diags.Errors())
	}

	if len(keyUsageStrings) == 0 {
		return nil
	}

	var combined graphmodels.KeyUsages
	for _, usage := range keyUsageStrings {
		parsed, err := graphmodels.ParseKeyUsages(usage)
		if err != nil {
			return fmt.Errorf("%w: %s", errInvalidKeyUsage, usage)
		}
		// Parse returns (nil, nil) for values outside the enum, so a nil result is a bad value
		// rather than an error condition.
		value, ok := parsed.(*graphmodels.KeyUsages)
		if !ok || value == nil {
			return fmt.Errorf("%w: %s", errInvalidKeyUsage, usage)
		}
		combined |= *value
	}

	setter(&combined)
	return nil
}

// setSubjectAlternativeNameType combines the configured SAN types into the single bitmask value
// Graph expects. Like KeyUsages this is a flags enum (values 1/2/4/8/16/32) serialized as a
// comma-joined string.
func setSubjectAlternativeNameType(
	ctx context.Context,
	sanTypeSet types.Set,
	setter func(*graphmodels.SubjectAlternativeNameType),
) error {
	if sanTypeSet.IsNull() || sanTypeSet.IsUnknown() {
		return nil
	}

	var sanTypeStrings []string
	if diags := sanTypeSet.ElementsAs(ctx, &sanTypeStrings, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractSanType, diags.Errors())
	}

	if len(sanTypeStrings) == 0 {
		return nil
	}

	var combined graphmodels.SubjectAlternativeNameType
	for _, sanType := range sanTypeStrings {
		parsed, err := graphmodels.ParseSubjectAlternativeNameType(sanType)
		if err != nil {
			return fmt.Errorf("%w: %s", errInvalidSanType, sanType)
		}
		value, ok := parsed.(*graphmodels.SubjectAlternativeNameType)
		if !ok || value == nil {
			return fmt.Errorf("%w: %s", errInvalidSanType, sanType)
		}
		combined |= *value
	}

	setter(&combined)
	return nil
}

// setEasServices combines the configured Exchange services into the single bitmask value Graph
// expects. EasServices is a flags enum (none=1, calendars=2, contacts=4, email=8, notes=16,
// reminders=32) serialized as a comma-joined string — the third such bitmask in this resource,
// alongside keyUsage and subjectAlternativeNameType.
func setEasServices(
	ctx context.Context,
	easServicesSet types.Set,
	setter func(*graphmodels.EasServices),
) error {
	if easServicesSet.IsNull() || easServicesSet.IsUnknown() {
		return nil
	}

	var serviceStrings []string
	if diags := easServicesSet.ElementsAs(ctx, &serviceStrings, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractEasServices, diags.Errors())
	}

	if len(serviceStrings) == 0 {
		return nil
	}

	var combined graphmodels.EasServices
	for _, service := range serviceStrings {
		parsed, err := graphmodels.ParseEasServices(service)
		if err != nil {
			return fmt.Errorf("%w: %s", errInvalidEasService, service)
		}
		value, ok := parsed.(*graphmodels.EasServices)
		if !ok || value == nil {
			return fmt.Errorf("%w: %s", errInvalidEasService, service)
		}
		combined |= *value
	}

	setter(&combined)
	return nil
}

func convertCustomSubjectAlternativeNames(
	ctx context.Context,
	sanSet types.Set,
	setter func([]graphmodels.CustomSubjectAlternativeNameable),
) error {
	if sanSet.IsNull() || sanSet.IsUnknown() {
		return nil
	}

	var sanModels []CustomSubjectAlternativeNameResourceModel
	if diags := sanSet.ElementsAs(ctx, &sanModels, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractCustomSans, diags.Errors())
	}

	sans := make([]graphmodels.CustomSubjectAlternativeNameable, 0, len(sanModels))
	for _, sanModel := range sanModels {
		san := graphmodels.NewCustomSubjectAlternativeName()

		if err := convert.FrameworkToGraphEnum(
			sanModel.SanType,
			graphmodels.ParseSubjectAlternativeNameType,
			san.SetSanType,
		); err != nil {
			return fmt.Errorf("%w: %s", errInvalidSanType, sanModel.SanType.ValueString())
		}

		convert.FrameworkToGraphString(sanModel.Name, san.SetName)
		sans = append(sans, san)
	}

	setter(sans)
	return nil
}

func convertExtendedKeyUsages(
	ctx context.Context,
	ekuSet types.Set,
	setter func([]graphmodels.ExtendedKeyUsageable),
) error {
	if ekuSet.IsNull() || ekuSet.IsUnknown() {
		return nil
	}

	var ekuModels []ExtendedKeyUsageResourceModel
	if diags := ekuSet.ElementsAs(ctx, &ekuModels, false); diags.HasError() {
		return fmt.Errorf("%w: %v", errExtractExtendedKeyUsages, diags.Errors())
	}

	ekus := make([]graphmodels.ExtendedKeyUsageable, 0, len(ekuModels))
	for _, ekuModel := range ekuModels {
		eku := graphmodels.NewExtendedKeyUsage()
		convert.FrameworkToGraphString(ekuModel.Name, eku.SetName)
		convert.FrameworkToGraphString(ekuModel.ObjectIdentifier, eku.SetObjectIdentifier)
		ekus = append(ekus, eku)
	}

	setter(ekus)
	return nil
}
