// REF: https://learn.microsoft.com/en-us/graph/api/resources/intune-deviceconfig-iosgeneraldeviceconfiguration?view=graph-rest-beta
package graphBetaIosDeviceConfigurationTemplates

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// IosDeviceConfigurationTemplatesResourceModel describes the resource data model.
type IosDeviceConfigurationTemplatesResourceModel struct {
	ID              types.String `tfsdk:"id"`
	DisplayName     types.String `tfsdk:"display_name"`
	Description     types.String `tfsdk:"description"`
	RoleScopeTagIds types.Set    `tfsdk:"role_scope_tag_ids"`
	// Nested configuration blocks (mutually exclusive)
	GeneralDeviceConfiguration types.Object   `tfsdk:"general_device_configuration"`
	CustomConfiguration        types.Object   `tfsdk:"custom_configuration"`
	TrustedCertificate  types.Object   `tfsdk:"trusted_certificate"`
	Wifi                types.Object   `tfsdk:"wifi"`
	ScepCertificate     types.Object   `tfsdk:"scep_certificate"`
	PkcsCertificate     types.Object   `tfsdk:"pkcs_certificate"`
	EnterpriseWifi      types.Object   `tfsdk:"enterprise_wifi"`
	EasEmail            types.Object   `tfsdk:"eas_email"`
	Vpn                 types.Object   `tfsdk:"vpn"`
	Assignments         types.Set      `tfsdk:"assignments"`
	Timeouts            timeouts.Value `tfsdk:"timeouts"`
}

// GeneralDeviceConfigurationResourceModel describes iosGeneralDeviceConfiguration.
//
// This type covers iOS device restriction policies. The full property set is large and overlaps
// significantly with the JSON resource; expose it here as an empty marker block so the structured
// resource can create and manage general device configurations without requiring raw JSON.
type GeneralDeviceConfigurationResourceModel struct{}

// CustomConfigurationResourceModel describes iosCustomConfiguration.
//
// Unlike macOSCustomConfiguration this type has no deploymentChannel property — the iOS SDK model
// exposes only the three payload setters.
type CustomConfigurationResourceModel struct {
	PayloadFileName types.String `tfsdk:"payload_file_name"`
	Payload         types.String `tfsdk:"payload"`
	PayloadName     types.String `tfsdk:"payload_name"`
}

// TrustedCertificateResourceModel describes iosTrustedRootCertificate.
//
// As with custom configurations, iOS has no deploymentChannel here.
type TrustedCertificateResourceModel struct {
	CertFileName           types.String `tfsdk:"cert_file_name"`
	TrustedRootCertificate types.String `tfsdk:"trusted_root_certificate"`
}

// ScepCertificateResourceModel describes iosScepCertificateProfile.
//
// renewal_threshold_percentage, certificate_validity_period_*, subject_name_format and
// subject_alternative_name_type come from the embedded IosCertificateProfileBase; the rest are
// defined directly on the SCEP type. Note there is no allow_all_apps_access on iOS.
type ScepCertificateResourceModel struct {
	RenewalThresholdPercentage         types.Int32  `tfsdk:"renewal_threshold_percentage"`
	CertificateStore                   types.String `tfsdk:"certificate_store"`
	CertificateValidityPeriodScale     types.String `tfsdk:"certificate_validity_period_scale"`
	CertificateValidityPeriodValue     types.Int32  `tfsdk:"certificate_validity_period_value"`
	SubjectNameFormat                  types.String `tfsdk:"subject_name_format"`
	SubjectNameFormatString            types.String `tfsdk:"subject_name_format_string"`
	SubjectAlternativeNameType         types.Set    `tfsdk:"subject_alternative_name_type"`
	SubjectAlternativeNameFormatString types.String `tfsdk:"subject_alternative_name_format_string"`
	RootCertificateOdataBind           types.String `tfsdk:"root_certificate_odata_bind"`
	KeySize                            types.String `tfsdk:"key_size"`
	KeyUsage                           types.Set    `tfsdk:"key_usage"`
	CustomSubjectAlternativeNames      types.Set    `tfsdk:"custom_subject_alternative_names"`
	ExtendedKeyUsages                  types.Set    `tfsdk:"extended_key_usages"`
	ScepServerUrls                     types.Set    `tfsdk:"scep_server_urls"`
}

// PkcsCertificateResourceModel describes iosPkcsCertificateProfile.
//
// Considerably thinner than the macOS equivalent: iOS PKCS has no key_size, key_usage,
// extended_key_usages, scep_server_urls or allow_all_apps_access.
type PkcsCertificateResourceModel struct {
	RenewalThresholdPercentage         types.Int32  `tfsdk:"renewal_threshold_percentage"`
	CertificateStore                   types.String `tfsdk:"certificate_store"`
	CertificateValidityPeriodScale     types.String `tfsdk:"certificate_validity_period_scale"`
	CertificateValidityPeriodValue     types.Int32  `tfsdk:"certificate_validity_period_value"`
	SubjectNameFormat                  types.String `tfsdk:"subject_name_format"`
	SubjectNameFormatString            types.String `tfsdk:"subject_name_format_string"`
	SubjectAlternativeNameType         types.Set    `tfsdk:"subject_alternative_name_type"`
	SubjectAlternativeNameFormatString types.String `tfsdk:"subject_alternative_name_format_string"`
	CertificationAuthority             types.String `tfsdk:"certification_authority"`
	CertificationAuthorityName         types.String `tfsdk:"certification_authority_name"`
	CertificateTemplateName            types.String `tfsdk:"certificate_template_name"`
	CustomSubjectAlternativeNames      types.Set    `tfsdk:"custom_subject_alternative_names"`
}

// EnterpriseWifiResourceModel describes iosEnterpriseWiFiConfiguration.
//
// The type embeds IosWiFiConfiguration, so the connection and proxy fields are inherited. Note that
// pre_shared_key is inherited too but deliberately not exposed: enterprise networks authenticate
// via EAP, and a PSK alongside eap_type would be an invalid combination.
type EnterpriseWifiResourceModel struct {
	// Inherited from IosWiFiConfiguration
	NetworkName                    types.String `tfsdk:"network_name"`
	Ssid                           types.String `tfsdk:"ssid"`
	ConnectAutomatically           types.Bool   `tfsdk:"connect_automatically"`
	ConnectWhenNetworkNameIsHidden types.Bool   `tfsdk:"connect_when_network_name_is_hidden"`
	WifiSecurityType               types.String `tfsdk:"wifi_security_type"`
	DisableMacAddressRandomization types.Bool   `tfsdk:"disable_mac_address_randomization"`
	ProxySettings                  types.String `tfsdk:"proxy_settings"`
	ProxyManualAddress             types.String `tfsdk:"proxy_manual_address"`
	ProxyManualPort                types.Int32  `tfsdk:"proxy_manual_port"`
	ProxyAutomaticConfigurationUrl types.String `tfsdk:"proxy_automatic_configuration_url"`
	// Enterprise-specific
	EapType                                             types.String `tfsdk:"eap_type"`
	EapFastConfiguration                                types.String `tfsdk:"eap_fast_configuration"`
	AuthenticationMethod                                types.String `tfsdk:"authentication_method"`
	InnerAuthenticationProtocolForEapTtls               types.String `tfsdk:"inner_authentication_protocol_for_eap_ttls"`
	OuterIdentityPrivacyTemporaryValue                  types.String `tfsdk:"outer_identity_privacy_temporary_value"`
	UsernameFormatString                                types.String `tfsdk:"username_format_string"`
	PasswordFormatString                                types.String `tfsdk:"password_format_string"`
	TrustedServerCertificateNames                       types.Set    `tfsdk:"trusted_server_certificate_names"`
	RootCertificatesForServerValidationOdataBind        types.Set    `tfsdk:"root_certificates_for_server_validation_odata_bind"`
	IdentityCertificateForClientAuthenticationOdataBind types.String `tfsdk:"identity_certificate_for_client_authentication_odata_bind"`
}

// EasEmailResourceModel describes iosEasEmailProfileConfiguration.
//
// custom_domain_name, user_domain_name_source, username_aad_source and username_source come from
// the embedded EasEmailProfileConfigurationBase. Note username_aad_source uses the UsernameSource
// enum (which adds samAccountName) while username_source and email_address_source use
// UserEmailSource — they are not interchangeable.
type EasEmailResourceModel struct {
	AccountName                    types.String `tfsdk:"account_name"`
	HostName                       types.String `tfsdk:"host_name"`
	AuthenticationMethod           types.String `tfsdk:"authentication_method"`
	EasServices                    types.Set    `tfsdk:"eas_services"`
	EasServicesUserOverrideEnabled types.Bool   `tfsdk:"eas_services_user_override_enabled"`
	DurationOfEmailToSync          types.String `tfsdk:"duration_of_email_to_sync"`
	EmailAddressSource             types.String `tfsdk:"email_address_source"`
	UsernameSource                 types.String `tfsdk:"username_source"`
	UsernameAADSource              types.String `tfsdk:"username_aad_source"`
	UserDomainNameSource           types.String `tfsdk:"user_domain_name_source"`
	CustomDomainName               types.String `tfsdk:"custom_domain_name"`
	RequireSsl                     types.Bool   `tfsdk:"require_ssl"`
	UseOAuth                       types.Bool   `tfsdk:"use_oauth"`
	PerAppVPNProfileId             types.String `tfsdk:"per_app_vpn_profile_id"`
	// Message handling restrictions
	BlockMovingMessagesToOtherEmailAccounts types.Bool `tfsdk:"block_moving_messages_to_other_email_accounts"`
	BlockSendingEmailFromThirdPartyApps     types.Bool `tfsdk:"block_sending_email_from_third_party_apps"`
	BlockSyncingRecentlyUsedEmailAddresses  types.Bool `tfsdk:"block_syncing_recently_used_email_addresses"`
	// S/MIME
	RequireSmime                                  types.Bool   `tfsdk:"require_smime"`
	SmimeEnablePerMessageSwitch                   types.Bool   `tfsdk:"smime_enable_per_message_switch"`
	SmimeSigningEnabled                           types.Bool   `tfsdk:"smime_signing_enabled"`
	SmimeSigningUserOverrideEnabled               types.Bool   `tfsdk:"smime_signing_user_override_enabled"`
	SmimeSigningCertificateUserOverrideEnabled    types.Bool   `tfsdk:"smime_signing_certificate_user_override_enabled"`
	SmimeEncryptByDefaultEnabled                  types.Bool   `tfsdk:"smime_encrypt_by_default_enabled"`
	SmimeEncryptByDefaultUserOverrideEnabled      types.Bool   `tfsdk:"smime_encrypt_by_default_user_override_enabled"`
	SmimeEncryptionCertificateUserOverrideEnabled types.Bool   `tfsdk:"smime_encryption_certificate_user_override_enabled"`
	SigningCertificateType                        types.String `tfsdk:"signing_certificate_type"`
	EncryptionCertificateType                     types.String `tfsdk:"encryption_certificate_type"`
	// Certificate references, set via @odata.bind
	IdentityCertificateOdataBind        types.String `tfsdk:"identity_certificate_odata_bind"`
	SmimeSigningCertificateOdataBind    types.String `tfsdk:"smime_signing_certificate_odata_bind"`
	SmimeEncryptionCertificateOdataBind types.String `tfsdk:"smime_encryption_certificate_odata_bind"`
}

// VpnResourceModel describes iosVpnConfiguration.
//
// Most fields come from the embedded AppleVpnConfiguration rather than the iOS leaf type. Two
// deliberate omissions:
//   - on_demand_rules: 10 fields x N entries; deferred to a later change.
//   - IKEv2: iosikEv2VpnConfiguration is a distinct @odata.type that embeds this one and adds ~23
//     properties. ikEv2 is excluded from the connection_type validator so a profile this resource
//     cannot fully model is never created here.
type VpnResourceModel struct {
	ConnectionName       types.String `tfsdk:"connection_name"`
	ConnectionType       types.String `tfsdk:"connection_type"`
	AuthenticationMethod types.String `tfsdk:"authentication_method"`
	Identifier           types.String `tfsdk:"identifier"`
	ProviderType         types.String `tfsdk:"provider_type"`
	Realm                types.String `tfsdk:"realm"`
	Role                 types.String `tfsdk:"role"`
	LoginGroupOrDomain   types.String `tfsdk:"login_group_or_domain"`
	// Split tunnelling and network scope
	EnableSplitTunneling types.Bool `tfsdk:"enable_split_tunneling"`
	EnablePerApp         types.Bool `tfsdk:"enable_per_app"`
	IncludeAllNetworks   types.Bool `tfsdk:"include_all_networks"`
	ExcludeLocalNetworks types.Bool `tfsdk:"exclude_local_networks"`
	SafariDomains        types.Set  `tfsdk:"safari_domains"`
	AssociatedDomains    types.Set  `tfsdk:"associated_domains"`
	ExcludedDomains      types.Set  `tfsdk:"excluded_domains"`
	// Idle / on-demand behaviour
	DisconnectOnIdle               types.Bool  `tfsdk:"disconnect_on_idle"`
	DisconnectOnIdleTimerInSeconds types.Int32 `tfsdk:"disconnect_on_idle_timer_in_seconds"`
	DisableOnDemandUserOverride    types.Bool  `tfsdk:"disable_on_demand_user_override"`
	OptInToDeviceIdSharing         types.Bool  `tfsdk:"opt_in_to_device_id_sharing"`
	// Microsoft Tunnel
	MicrosoftTunnelSiteId types.String `tfsdk:"microsoft_tunnel_site_id"`
	StrictEnforcement     types.Bool   `tfsdk:"strict_enforcement"`
	// Zscaler-specific
	CloudName   types.String `tfsdk:"cloud_name"`
	UserDomain  types.String `tfsdk:"user_domain"`
	ExcludeList types.Set    `tfsdk:"exclude_list"`
	// Nested objects
	Server      types.Object `tfsdk:"server"`
	ProxyServer types.Object `tfsdk:"proxy_server"`
	// Collections
	TargetedMobileApps types.Set `tfsdk:"targeted_mobile_apps"`
	CustomData         types.Set `tfsdk:"custom_data"`
	CustomKeyValueData types.Set `tfsdk:"custom_key_value_data"`
	// Certificate reference, set via @odata.bind
	IdentityCertificateOdataBind types.String `tfsdk:"identity_certificate_odata_bind"`
}

// VpnServerResourceModel describes the vpnServer object.
type VpnServerResourceModel struct {
	Address         types.String `tfsdk:"address"`
	Description     types.String `tfsdk:"description"`
	IsDefaultServer types.Bool   `tfsdk:"is_default_server"`
}

// VpnProxyServerResourceModel describes the vpnProxyServer object.
type VpnProxyServerResourceModel struct {
	Address                         types.String `tfsdk:"address"`
	Port                            types.Int32  `tfsdk:"port"`
	AutomaticConfigurationScriptUrl types.String `tfsdk:"automatic_configuration_script_url"`
}

// AppListItemResourceModel describes an app in targeted_mobile_apps.
type AppListItemResourceModel struct {
	Name        types.String `tfsdk:"name"`
	AppId       types.String `tfsdk:"app_id"`
	Publisher   types.String `tfsdk:"publisher"`
	AppStoreUrl types.String `tfsdk:"app_store_url"`
}

// KeyValueResourceModel describes a customData entry. Note the field is "key", unlike
// customKeyValueData which uses "name".
type KeyValueResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

// KeyValuePairResourceModel describes a customKeyValueData entry. Note the field is "name", unlike
// customData which uses "key".
type KeyValuePairResourceModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// CustomSubjectAlternativeNameResourceModel describes custom SAN entries.
type CustomSubjectAlternativeNameResourceModel struct {
	SanType types.String `tfsdk:"san_type"`
	Name    types.String `tfsdk:"name"`
}

// ExtendedKeyUsageResourceModel describes extended key usage entries.
type ExtendedKeyUsageResourceModel struct {
	Name             types.String `tfsdk:"name"`
	ObjectIdentifier types.String `tfsdk:"object_identifier"`
}

// WifiResourceModel describes iosWiFiConfiguration.
type WifiResourceModel struct {
	NetworkName                    types.String `tfsdk:"network_name"`
	Ssid                           types.String `tfsdk:"ssid"`
	ConnectAutomatically           types.Bool   `tfsdk:"connect_automatically"`
	ConnectWhenNetworkNameIsHidden types.Bool   `tfsdk:"connect_when_network_name_is_hidden"`
	WifiSecurityType               types.String `tfsdk:"wifi_security_type"`
	PreSharedKey                   types.String `tfsdk:"pre_shared_key"`
	DisableMacAddressRandomization types.Bool   `tfsdk:"disable_mac_address_randomization"`
	ProxySettings                  types.String `tfsdk:"proxy_settings"`
	ProxyManualAddress             types.String `tfsdk:"proxy_manual_address"`
	ProxyManualPort                types.Int32  `tfsdk:"proxy_manual_port"`
	ProxyAutomaticConfigurationUrl types.String `tfsdk:"proxy_automatic_configuration_url"`
}
