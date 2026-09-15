package graphBetaIosDeviceConfigurationTemplates

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/client"
	planmodifiers "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/plan_modifiers"
	commonschema "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/schema"
	commonschemagraphbeta "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/schema/graph_beta/device_management"
)

const (
	ResourceName  = "microsoft365_graph_beta_device_management_ios_device_configuration_templates"
	CreateTimeout = 180
	UpdateTimeout = 180
	ReadTimeout   = 180
	DeleteTimeout = 180
)

var (
	// Basic resource interface (CRUD operations)
	_ resource.Resource = &IosDeviceConfigurationTemplatesResource{}

	// Allows the resource to be configured with the provider client
	_ resource.ResourceWithConfigure = &IosDeviceConfigurationTemplatesResource{}

	// Enables import functionality
	_ resource.ResourceWithImportState = &IosDeviceConfigurationTemplatesResource{}

	// Enables plan modification/diff suppression
	_ resource.ResourceWithModifyPlan = &IosDeviceConfigurationTemplatesResource{}

	// Enables identity schema for list resource support
	_ resource.ResourceWithIdentity = &IosDeviceConfigurationTemplatesResource{}
)

func NewIosDeviceConfigurationTemplatesResource() resource.Resource {
	return &IosDeviceConfigurationTemplatesResource{
		ReadPermissions: []string{
			"DeviceManagementConfiguration.Read.All",
		},
		WritePermissions: []string{
			"DeviceManagementConfiguration.ReadWrite.All",
		},
		ResourcePath: "/deviceManagement/deviceConfigurations",
	}
}

type IosDeviceConfigurationTemplatesResource struct {
	client           *msgraphbetasdk.GraphServiceClient
	ReadPermissions  []string
	WritePermissions []string
	ResourcePath     string
}

// Metadata returns the resource type name.
func (r *IosDeviceConfigurationTemplatesResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = ResourceName
}

// Configure sets the client for the resource.
func (r *IosDeviceConfigurationTemplatesResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	r.client = client.SetGraphBetaClientForResource(ctx, req, resp, ResourceName)
}

// ImportState imports the resource state.
func (r *IosDeviceConfigurationTemplatesResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// IdentitySchema defines the identity schema for this resource, used by list operations to uniquely identify instances
func (r *IosDeviceConfigurationTemplatesResource) IdentitySchema(
	ctx context.Context,
	req resource.IdentitySchemaRequest,
	resp *resource.IdentitySchemaResponse,
) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

// Schema defines the schema for the resource.
func (r *IosDeviceConfigurationTemplatesResource) Schema(
	ctx context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages iOS/iPadOS configuration templates in Microsoft Intune. " +
			"This resource creates device configurations for iOS/iPadOS devices including custom configuration " +
			"profiles, trusted root certificates, certificate profiles (SCEP/PKCS), Wi-Fi profiles (personal and " +
			"enterprise), Exchange ActiveSync email profiles, and VPN profiles.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the iOS/iPadOS configuration template.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"display_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The display name for the iOS/iPadOS configuration template.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Optional description of the resource. Maximum length is 1500 characters.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1500),
				},
			},
			"role_scope_tag_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Set of scope tag IDs for this device configuration template.",
				PlanModifiers: []planmodifier.Set{
					planmodifiers.DefaultSetValue(
						[]attr.Value{types.StringValue("0")},
					),
				},
			},
			"general_device_configuration": schema.SingleNestedAttribute{
				Optional: true,
				DeprecationMessage: "Use microsoft365_graph_beta_device_management_ios_device_configuration_templates_json " +
					"with odata_type = \"#microsoft.graph.iosGeneralDeviceConfiguration\" to manage the full property set. " +
					"This block creates the profile shell but cannot track or reconcile device restriction settings configured outside Terraform.",
				MarkdownDescription: "General iOS/iPadOS device restriction policy (`iosGeneralDeviceConfiguration`). " +
					"Set this block (with no inner attributes required) to create a general device restrictions profile. " +
					"Configure the full property set via the JSON resource or Intune portal after creation.",
				Attributes: map[string]schema.Attribute{},
			},
			"custom_configuration": schema.SingleNestedAttribute{
				Optional: true,
				MarkdownDescription: "The custom configuration template allows IT admins to assign settings that aren't built into Intune yet. " +
					"For iOS/iPadOS devices, you can import a .mobileconfig file that you created using Apple Configurator or Profile Manager.",
				Attributes: map[string]schema.Attribute{
					"payload_file_name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The profile name displayed to users.",
					},
					"payload": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The iOS/iPadOS configuration payload (.mobileconfig / .plist) file content.",
					},
					"payload_name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The name of the payload configuration.",
					},
				},
			},
			"trusted_certificate": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Trusted root certificate configuration for iOS/iPadOS devices.",
				Attributes: map[string]schema.Attribute{
					"cert_file_name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The file name of the certificate file (.cer file).",
					},
					"trusted_root_certificate": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The base64-encoded trusted root certificate content. This should be a filebase64() encoded string. e.g filebase64(\"my-root-cert.cer\")",
					},
				},
			},
			"wifi": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Wi-Fi configuration for iOS/iPadOS devices. For enterprise (802.1x) networks use the enterprise Wi-Fi profile instead.",
				Attributes: map[string]schema.Attribute{
					"network_name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The name of the Wi-Fi network shown to users when browsing available networks on the device.",
					},
					"ssid": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The service set identifier (SSID) of the Wi-Fi network the device connects to.",
					},
					"connect_automatically": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether the device connects automatically to this Wi-Fi network when in range.",
					},
					"connect_when_network_name_is_hidden": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether the device connects to the network even when the SSID is not broadcast.",
					},
					"wifi_security_type": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The Wi-Fi security protocol. Possible values are: open, wpaPersonal, wpaEnterprise, wep, " +
							"wpa2Personal, wpa2Enterprise, wpa3Personal. For the personal and WEP variants, set pre_shared_key.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"open",
								"wpaPersonal",
								"wpaEnterprise",
								"wep",
								"wpa2Personal",
								"wpa2Enterprise",
								"wpa3Personal",
							),
						},
					},
					"pre_shared_key": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
						MarkdownDescription: "The pre-shared key (password) for the Wi-Fi network. Applies to the personal and WEP " +
							"security types. Graph does not return this value on read, so it is preserved from configuration.",
					},
					"disable_mac_address_randomization": schema.BoolAttribute{
						Optional: true,
						MarkdownDescription: "Whether to disable the device's private (randomized) Wi-Fi MAC address for this network. " +
							"Set this when the network relies on MAC-based authentication.",
					},
					"proxy_settings": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "How the device obtains its proxy configuration for this network. Possible values are: none, manual, automatic.",
						Validators: []validator.String{
							stringvalidator.OneOf("none", "manual", "automatic"),
						},
					},
					"proxy_manual_address": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The IP address or hostname of the proxy server. Applies when proxy_settings is manual.",
					},
					"proxy_manual_port": schema.Int32Attribute{
						Optional:            true,
						MarkdownDescription: "The port of the proxy server. Applies when proxy_settings is manual.",
						Validators: []validator.Int32{
							int32validator.Between(1, 65535),
						},
					},
					"proxy_automatic_configuration_url": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The URL of the proxy auto-configuration (PAC) file. Applies when proxy_settings is automatic.",
					},
				},
			},
			"scep_certificate": schema.SingleNestedAttribute{
				Optional: true,
				MarkdownDescription: "SCEP certificate profile configuration for iOS/iPadOS devices. Requires a trusted root " +
					"certificate profile to already exist, referenced via root_certificate_odata_bind.",
				Attributes: map[string]schema.Attribute{
					"renewal_threshold_percentage": schema.Int32Attribute{
						Required:            true,
						MarkdownDescription: "The percentage of the certificate lifetime remaining when renewal is attempted (1-99).",
						Validators: []validator.Int32{
							int32validator.Between(1, 99),
						},
					},
					"certificate_store": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The certificate store location. Possible values are: user, machine.",
						Validators: []validator.String{
							stringvalidator.OneOf("user", "machine"),
						},
					},
					"certificate_validity_period_scale": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The unit for the certificate validity period. Possible values are: days, months, years.",
						Validators: []validator.String{
							stringvalidator.OneOf("days", "months", "years"),
						},
					},
					"certificate_validity_period_value": schema.Int32Attribute{
						Required:            true,
						MarkdownDescription: "The certificate validity period, in the unit given by certificate_validity_period_scale.",
						Validators: []validator.Int32{
							int32validator.AtLeast(1),
						},
					},
					"subject_name_format": schema.StringAttribute{
						Required: true,
						MarkdownDescription: "How Intune builds the subject name in the certificate request. Possible values are: " +
							"commonName, commonNameAsEmail, custom, commonNameIncludingEmail, commonNameAsIMEI, commonNameAsSerialNumber. " +
							"Use custom together with subject_name_format_string. See " +
							"https://learn.microsoft.com/en-us/intune/intune-service/protect/certificates-profile-scep",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"commonName",
								"commonNameAsEmail",
								"custom",
								"commonNameIncludingEmail",
								"commonNameAsIMEI",
								"commonNameAsSerialNumber",
							),
						},
					},
					"subject_name_format_string": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The custom subject name format, used when subject_name_format is custom. " +
							"Example: CN={{AAD_Device_ID}},O={{Organization}}",
					},
					"subject_alternative_name_type": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
						MarkdownDescription: "The subject alternative name types to include. Graph models this as a bitmask, so " +
							"multiple values may be combined. Possible values are: none, emailAddress, userPrincipalName, " +
							"customAzureADAttribute, domainNameService, universalResourceIdentifier.",
						Validators: []validator.Set{
							setvalidator.ValueStringsAre(stringvalidator.OneOf(
								"none",
								"emailAddress",
								"userPrincipalName",
								"customAzureADAttribute",
								"domainNameService",
								"universalResourceIdentifier",
							)),
						},
					},
					"subject_alternative_name_format_string": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The custom subject alternative name format string.",
					},
					"root_certificate_odata_bind": schema.StringAttribute{
						Required: true,
						MarkdownDescription: "Reference to a pre-existing iOS trusted root certificate profile. Accepts either a " +
							"bare GUID (e.g. '00000000-0000-0000-0000-000000000000') or the full URL " +
							"\"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('00000000-0000-0000-0000-000000000000')\". " +
							"Graph does not return this reference on read, so it is preserved from configuration.",
					},
					"key_size": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The key size in bits. 2048 is the recommended minimum. Possible values are: size1024, size2048, size4096.",
						Validators: []validator.String{
							stringvalidator.OneOf("size1024", "size2048", "size4096"),
						},
					},
					"key_usage": schema.SetAttribute{
						ElementType: types.StringType,
						Required:    true,
						MarkdownDescription: "Key usage options for the certificate. Graph models this as a bitmask, so both values " +
							"may be combined. Possible values are: keyEncipherment, digitalSignature.",
						Validators: []validator.Set{
							setvalidator.ValueStringsAre(
								stringvalidator.OneOf("keyEncipherment", "digitalSignature"),
							),
						},
					},
					"custom_subject_alternative_names": schema.SetNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Custom Subject Alternative Names for the certificate.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"san_type": schema.StringAttribute{
									Required: true,
									MarkdownDescription: "The SAN type. Possible values are: none, emailAddress, userPrincipalName, " +
										"customAzureADAttribute, domainNameService, universalResourceIdentifier.",
									Validators: []validator.String{
										stringvalidator.OneOf(
											"none",
											"emailAddress",
											"userPrincipalName",
											"customAzureADAttribute",
											"domainNameService",
											"universalResourceIdentifier",
										),
									},
								},
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The SAN value/name.",
								},
							},
						},
					},
					"extended_key_usages": schema.SetNestedAttribute{
						Required:            true,
						MarkdownDescription: "Extended key usage settings for the certificate.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The extended key usage name.",
								},
								"object_identifier": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The extended key usage object identifier (OID).",
								},
							},
						},
					},
					"scep_server_urls": schema.SetAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						MarkdownDescription: "SCEP server URL(s) for certificate enrollment.",
						Validators: []validator.Set{
							setvalidator.SizeAtLeast(1),
						},
					},
				},
			},
			"pkcs_certificate": schema.SingleNestedAttribute{
				Optional: true,
				MarkdownDescription: "PKCS certificate profile configuration for iOS/iPadOS devices. Unlike the macOS equivalent, " +
					"iOS PKCS profiles have no key size, key usage or extended key usage properties.",
				Attributes: map[string]schema.Attribute{
					"renewal_threshold_percentage": schema.Int32Attribute{
						Optional:            true,
						MarkdownDescription: "The percentage of the certificate lifetime remaining when renewal is attempted (1-99).",
						Validators: []validator.Int32{
							int32validator.Between(1, 99),
						},
					},
					"certificate_store": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The certificate store location. Possible values are: user, machine.",
						Validators: []validator.String{
							stringvalidator.OneOf("user", "machine"),
						},
					},
					"certificate_validity_period_scale": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The unit for the certificate validity period. Possible values are: days, months, years.",
						Validators: []validator.String{
							stringvalidator.OneOf("days", "months", "years"),
						},
					},
					"certificate_validity_period_value": schema.Int32Attribute{
						Optional:            true,
						MarkdownDescription: "The certificate validity period, in the unit given by certificate_validity_period_scale.",
						Validators: []validator.Int32{
							int32validator.AtLeast(1),
						},
					},
					"subject_name_format": schema.StringAttribute{
						Required: true,
						MarkdownDescription: "How Intune builds the subject name in the certificate request. Possible values are: " +
							"commonName, commonNameAsEmail, custom, commonNameIncludingEmail, commonNameAsIMEI, commonNameAsSerialNumber.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"commonName",
								"commonNameAsEmail",
								"custom",
								"commonNameIncludingEmail",
								"commonNameAsIMEI",
								"commonNameAsSerialNumber",
							),
						},
					},
					"subject_name_format_string": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The custom subject name format, used when subject_name_format is custom. " +
							"Example: CN={{UserName}},E={{EmailAddress}},O=Example Corp",
					},
					"subject_alternative_name_type": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
						MarkdownDescription: "The subject alternative name types to include. Graph models this as a bitmask, so " +
							"multiple values may be combined. Possible values are: none, emailAddress, userPrincipalName, " +
							"customAzureADAttribute, domainNameService, universalResourceIdentifier.",
						Validators: []validator.Set{
							setvalidator.ValueStringsAre(stringvalidator.OneOf(
								"none",
								"emailAddress",
								"userPrincipalName",
								"customAzureADAttribute",
								"domainNameService",
								"universalResourceIdentifier",
							)),
						},
					},
					"subject_alternative_name_format_string": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The custom subject alternative name format string.",
					},
					"certification_authority": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The fully qualified domain name of the issuing certification authority.",
					},
					"certification_authority_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The name of the issuing certification authority.",
					},
					"certificate_template_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The name of the certificate template on the issuing certification authority.",
					},
					"custom_subject_alternative_names": schema.SetNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Custom Subject Alternative Names for the certificate.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"san_type": schema.StringAttribute{
									Required: true,
									MarkdownDescription: "The SAN type. Possible values are: none, emailAddress, userPrincipalName, " +
										"customAzureADAttribute, domainNameService, universalResourceIdentifier.",
									Validators: []validator.String{
										stringvalidator.OneOf(
											"none",
											"emailAddress",
											"userPrincipalName",
											"customAzureADAttribute",
											"domainNameService",
											"universalResourceIdentifier",
										),
									},
								},
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The SAN value/name.",
								},
							},
						},
					},
				},
			},
			"enterprise_wifi": schema.SingleNestedAttribute{
				Optional: true,
				MarkdownDescription: "Enterprise (802.1x) Wi-Fi configuration for iOS/iPadOS devices. Authentication is via EAP, " +
					"so there is no pre-shared key — use the `wifi` block for personal/WEP networks instead.",
				Attributes: map[string]schema.Attribute{
					"network_name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The name of the Wi-Fi network shown to users when browsing available networks.",
					},
					"ssid": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The service set identifier (SSID) of the Wi-Fi network the device connects to.",
					},
					"connect_automatically": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether the device connects automatically to this Wi-Fi network when in range.",
					},
					"connect_when_network_name_is_hidden": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether the device connects to the network even when the SSID is not broadcast.",
					},
					"wifi_security_type": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The Wi-Fi security protocol. For enterprise networks use wpaEnterprise or " +
							"wpa2Enterprise. Possible values are: open, wpaPersonal, wpaEnterprise, wep, wpa2Personal, " +
							"wpa2Enterprise, wpa3Personal.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"open",
								"wpaPersonal",
								"wpaEnterprise",
								"wep",
								"wpa2Personal",
								"wpa2Enterprise",
								"wpa3Personal",
							),
						},
					},
					"disable_mac_address_randomization": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether to disable the device's private (randomized) Wi-Fi MAC address for this network.",
					},
					"proxy_settings": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "How the device obtains its proxy configuration. Possible values are: none, manual, automatic.",
						Validators: []validator.String{
							stringvalidator.OneOf("none", "manual", "automatic"),
						},
					},
					"proxy_manual_address": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The IP address or hostname of the proxy server. Applies when proxy_settings is manual.",
					},
					"proxy_manual_port": schema.Int32Attribute{
						Optional:            true,
						MarkdownDescription: "The port of the proxy server. Applies when proxy_settings is manual.",
						Validators: []validator.Int32{
							int32validator.Between(1, 65535),
						},
					},
					"proxy_automatic_configuration_url": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The URL of the proxy auto-configuration (PAC) file. Applies when proxy_settings is automatic.",
					},
					"eap_type": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The extensible authentication protocol (EAP) type used for 802.1x authentication. " +
							"Possible values are: eapTls, leap, eapSim, eapTtls, peap, eapFast, teap.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"eapTls",
								"leap",
								"eapSim",
								"eapTtls",
								"peap",
								"eapFast",
								"teap",
							),
						},
					},
					"eap_fast_configuration": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Protected Access Credential (PAC) handling, applicable when eap_type is eapFast. " +
							"Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, " +
							"useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"noProtectedAccessCredential",
								"useProtectedAccessCredential",
								"useProtectedAccessCredentialAndProvision",
								"useProtectedAccessCredentialAndProvisionAnonymously",
							),
						},
					},
					"authentication_method": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "How the device authenticates to the network. Possible values are: certificate, " +
							"usernameAndPassword, derivedCredential.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"certificate",
								"usernameAndPassword",
								"derivedCredential",
							),
						},
					},
					"inner_authentication_protocol_for_eap_ttls": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The non-EAP inner authentication protocol, applicable when eap_type is eapTtls. " +
							"Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, " +
							"microsoftChap, microsoftChapVersionTwo.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"unencryptedPassword",
								"challengeHandshakeAuthenticationProtocol",
								"microsoftChap",
								"microsoftChapVersionTwo",
							),
						},
					},
					"outer_identity_privacy_temporary_value": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The identity sent in the clear during the outer EAP exchange, hiding the real " +
							"user identity until the tunnel is established.",
					},
					"username_format_string": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The username format used for authentication. Example: {{UserPrincipalName}}",
					},
					"password_format_string": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
						MarkdownDescription: "The password format used for authentication. Graph does not return this value on " +
							"read, so it is preserved from configuration.",
					},
					"trusted_server_certificate_names": schema.SetAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						MarkdownDescription: "The common names of the certificates the device should trust for server validation.",
					},
					"root_certificates_for_server_validation_odata_bind": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
						MarkdownDescription: "References to pre-existing iOS trusted root certificate profiles used to validate " +
							"the RADIUS server. Each entry accepts a bare GUID or the full " +
							"\"https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations('...')\" URL. Graph does " +
							"not return these references on read, so they are preserved from configuration.",
					},
					"identity_certificate_for_client_authentication_odata_bind": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Reference to a pre-existing iOS SCEP or PKCS certificate profile used for client " +
							"authentication. Accepts a bare GUID or the full URL. Required when authentication_method is " +
							"certificate. Graph does not return this reference on read.",
					},
				},
			},
			"eas_email": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Exchange ActiveSync email profile configuration for iOS/iPadOS devices.",
				Attributes: map[string]schema.Attribute{
					"account_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The display name of the email account as shown to users on the device.",
					},
					"host_name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The Exchange server hostname that the device connects to.",
					},
					"authentication_method": schema.StringAttribute{
						Required: true,
						MarkdownDescription: "How the device authenticates to Exchange. Possible values are: usernameAndPassword, " +
							"certificate, derivedCredential.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"usernameAndPassword",
								"certificate",
								"derivedCredential",
							),
						},
					},
					"eas_services": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
						MarkdownDescription: "The Exchange data types to synchronise. Graph models this as a bitmask, so multiple " +
							"values may be combined. Possible values are: none, calendars, contacts, email, notes, reminders.",
						Validators: []validator.Set{
							setvalidator.ValueStringsAre(stringvalidator.OneOf(
								"none",
								"calendars",
								"contacts",
								"email",
								"notes",
								"reminders",
							)),
						},
					},
					"eas_services_user_override_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether users may change which Exchange data types are synchronised.",
					},
					"duration_of_email_to_sync": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "How much email history to synchronise. Possible values are: userDefined, oneDay, " +
							"threeDays, oneWeek, twoWeeks, oneMonth, unlimited.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"userDefined",
								"oneDay",
								"threeDays",
								"oneWeek",
								"twoWeeks",
								"oneMonth",
								"unlimited",
							),
						},
					},
					"email_address_source": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Which Entra ID attribute supplies the email address. Possible values are: " +
							"userPrincipalName, primarySmtpAddress.",
						Validators: []validator.String{
							stringvalidator.OneOf("userPrincipalName", "primarySmtpAddress"),
						},
					},
					"username_source": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Which Entra ID attribute supplies the username. Possible values are: " +
							"userPrincipalName, primarySmtpAddress.",
						Validators: []validator.String{
							stringvalidator.OneOf("userPrincipalName", "primarySmtpAddress"),
						},
					},
					"username_aad_source": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Which Entra ID attribute supplies the username. Note this uses a wider set of values " +
							"than username_source. Possible values are: userPrincipalName, primarySmtpAddress, samAccountName.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"userPrincipalName",
								"primarySmtpAddress",
								"samAccountName",
							),
						},
					},
					"user_domain_name_source": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Which form of the domain name to use. Possible values are: fullDomainName, " +
							"netBiosDomainName.",
						Validators: []validator.String{
							stringvalidator.OneOf("fullDomainName", "netBiosDomainName"),
						},
					},
					"custom_domain_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "A custom domain name value used instead of the one derived from Entra ID.",
					},
					"require_ssl": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether SSL is required for connections to the Exchange server.",
					},
					"use_oauth": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether the connection uses OAuth for authentication.",
					},
					"per_app_vpn_profile_id": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The identifier of a per-app VPN profile to associate with this email profile.",
					},
					"block_moving_messages_to_other_email_accounts": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether to prevent users moving messages out of this account into another.",
					},
					"block_sending_email_from_third_party_apps": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether to prevent third-party apps sending email from this account.",
					},
					"block_syncing_recently_used_email_addresses": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether to prevent syncing of recently used email addresses.",
					},
					"require_smime": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether S/MIME is required for outgoing messages.",
					},
					"smime_enable_per_message_switch": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether users may toggle S/MIME on a per-message basis.",
					},
					"smime_signing_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether S/MIME signing is enabled for this account.",
					},
					"smime_signing_user_override_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether users may change the S/MIME signing setting.",
					},
					"smime_signing_certificate_user_override_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether users may select the S/MIME signing certificate.",
					},
					"smime_encrypt_by_default_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether outgoing messages are S/MIME encrypted by default.",
					},
					"smime_encrypt_by_default_user_override_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether users may change the encrypt-by-default setting.",
					},
					"smime_encryption_certificate_user_override_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether users may select the S/MIME encryption certificate.",
					},
					"signing_certificate_type": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The type of certificate used for S/MIME signing. Possible values are: none, " +
							"certificate, derivedCredential.",
						Validators: []validator.String{
							stringvalidator.OneOf("none", "certificate", "derivedCredential"),
						},
					},
					"encryption_certificate_type": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The type of certificate used for S/MIME encryption. Possible values are: none, " +
							"certificate, derivedCredential.",
						Validators: []validator.String{
							stringvalidator.OneOf("none", "certificate", "derivedCredential"),
						},
					},
					"identity_certificate_odata_bind": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Reference to a pre-existing iOS SCEP or PKCS certificate profile used to authenticate " +
							"to Exchange. Accepts a bare GUID or the full URL. Required when authentication_method is certificate. " +
							"Graph does not return this reference on read.",
					},
					"smime_signing_certificate_odata_bind": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Reference to a pre-existing iOS certificate profile used for S/MIME signing. Accepts " +
							"a bare GUID or the full URL. Graph does not return this reference on read.",
					},
					"smime_encryption_certificate_odata_bind": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Reference to a pre-existing iOS certificate profile used for S/MIME encryption. " +
							"Accepts a bare GUID or the full URL. Graph does not return this reference on read.",
					},
				},
			},
			"vpn": schema.SingleNestedAttribute{
				Optional: true,
				MarkdownDescription: "VPN configuration for iOS/iPadOS devices. Note that IKEv2 profiles use a distinct Graph " +
					"type with a substantially different property set and are not supported by this block — `ikEv2` is " +
					"therefore not an accepted connection_type. On-demand rules are also not yet modelled.",
				Attributes: map[string]schema.Attribute{
					"connection_name": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The connection name displayed to users on the device.",
					},
					"connection_type": schema.StringAttribute{
						Required: true,
						MarkdownDescription: "The VPN vendor/provider. Possible values are: ciscoAnyConnect, pulseSecure, " +
							"f5EdgeClient, dellSonicWallMobileConnect, checkPointCapsuleVpn, customVpn, ciscoIPSec, citrix, " +
							"ciscoAnyConnectV2, paloAltoGlobalProtect, zscalerPrivateAccess, f5Access2018, citrixSso, " +
							"paloAltoGlobalProtectV2, alwaysOn, microsoftTunnel, netMotionMobility, microsoftProtect. " +
							"`ikEv2` is intentionally excluded: it maps to a separate Graph type this resource does not model.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"ciscoAnyConnect",
								"pulseSecure",
								"f5EdgeClient",
								"dellSonicWallMobileConnect",
								"checkPointCapsuleVpn",
								"customVpn",
								"ciscoIPSec",
								"citrix",
								"ciscoAnyConnectV2",
								"paloAltoGlobalProtect",
								"zscalerPrivateAccess",
								"f5Access2018",
								"citrixSso",
								"paloAltoGlobalProtectV2",
								"alwaysOn",
								"microsoftTunnel",
								"netMotionMobility",
								"microsoftProtect",
							),
						},
					},
					"authentication_method": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "How the device authenticates to the VPN. Possible values are: certificate, " +
							"usernameAndPassword, sharedSecret, derivedCredential, azureAD.",
						Validators: []validator.String{
							stringvalidator.OneOf(
								"certificate",
								"usernameAndPassword",
								"sharedSecret",
								"derivedCredential",
								"azureAD",
							),
						},
					},
					"identifier": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The vendor-supplied bundle identifier of the VPN app. Required when " +
							"connection_type is customVpn.",
					},
					"provider_type": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The tunnel provider type. Possible values are: notConfigured, appProxy, " +
							"packetTunnel.",
						Validators: []validator.String{
							stringvalidator.OneOf("notConfigured", "appProxy", "packetTunnel"),
						},
					},
					"realm": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The authentication realm, used by some VPN vendors.",
					},
					"role": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The authentication role, used by some VPN vendors.",
					},
					"login_group_or_domain": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The login group or domain, used by some VPN vendors.",
					},
					"enable_split_tunneling": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether only traffic destined for the VPN's networks is tunnelled.",
					},
					"enable_per_app": schema.BoolAttribute{
						Optional: true,
						Computed: true,
						MarkdownDescription: "Whether this is a per-app VPN, tunnelling only traffic from the apps listed in " +
							"targeted_mobile_apps.",
					},
					"include_all_networks": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether all network traffic is routed through the VPN.",
					},
					"exclude_local_networks": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether traffic to local networks bypasses the VPN.",
					},
					"safari_domains": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
						MarkdownDescription: "Domains that trigger the per-app VPN when visited in Safari. Applies when " +
							"enable_per_app is true.",
					},
					"associated_domains": schema.SetAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						MarkdownDescription: "Domains associated with this VPN profile.",
					},
					"excluded_domains": schema.SetAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						MarkdownDescription: "Domains whose traffic bypasses the VPN even while it is connected.",
					},
					"disconnect_on_idle": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether the VPN disconnects after an idle period.",
					},
					"disconnect_on_idle_timer_in_seconds": schema.Int32Attribute{
						Optional:            true,
						MarkdownDescription: "How long the connection may be idle before disconnecting, in seconds.",
						Validators: []validator.Int32{
							int32validator.AtLeast(0),
						},
					},
					"disable_on_demand_user_override": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether users are prevented from disabling on-demand VPN.",
					},
					"opt_in_to_device_id_sharing": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Whether the device identifier is shared with the VPN provider.",
					},
					"microsoft_tunnel_site_id": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "The Microsoft Tunnel site identifier. Applies when connection_type is " +
							"microsoftTunnel.",
					},
					"strict_enforcement": schema.BoolAttribute{
						Optional: true,
						Computed: true,
						MarkdownDescription: "Whether the VPN stays connected and blocks traffic when it cannot be " +
							"established. Applies to Microsoft Tunnel.",
					},
					"cloud_name": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The Zscaler cloud name. Applies when connection_type is zscalerPrivateAccess.",
					},
					"user_domain": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "The Zscaler user domain. Applies when connection_type is zscalerPrivateAccess.",
					},
					"exclude_list": schema.SetAttribute{
						ElementType: types.StringType,
						Optional:    true,
						MarkdownDescription: "Hostnames excluded from the Zscaler tunnel. Applies when connection_type is " +
							"zscalerPrivateAccess.",
					},
					"server": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "The VPN server the device connects to.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Required:            true,
								MarkdownDescription: "The IP address or fully qualified domain name of the VPN server.",
							},
							"description": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "A description of the VPN server.",
							},
							"is_default_server": schema.BoolAttribute{
								Optional:            true,
								Computed:            true,
								MarkdownDescription: "Whether this is the default server for the connection.",
							},
						},
					},
					"proxy_server": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "The proxy server used while the VPN is connected.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "The IP address or hostname of the proxy server.",
							},
							"port": schema.Int32Attribute{
								Optional:            true,
								MarkdownDescription: "The port of the proxy server.",
								Validators: []validator.Int32{
									int32validator.Between(1, 65535),
								},
							},
							"automatic_configuration_script_url": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "The URL of the proxy auto-configuration (PAC) script.",
							},
						},
					},
					"targeted_mobile_apps": schema.SetNestedAttribute{
						Optional: true,
						MarkdownDescription: "The apps whose traffic is tunnelled by this per-app VPN. Applies when " +
							"enable_per_app is true.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "The display name of the app.",
								},
								"app_id": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The bundle identifier of the app. e.g. com.microsoft.Office.Outlook",
								},
								"publisher": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "The publisher of the app.",
								},
								"app_store_url": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "The App Store URL of the app.",
								},
							},
						},
					},
					"custom_data": schema.SetNestedAttribute{
						Optional: true,
						MarkdownDescription: "Vendor-specific key/value configuration. Note this maps to Graph's `customData` " +
							"property, whose entries use a `key` field — distinct from custom_key_value_data below.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The configuration key.",
								},
								"value": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The configuration value.",
								},
							},
						},
					},
					"custom_key_value_data": schema.SetNestedAttribute{
						Optional: true,
						MarkdownDescription: "Vendor-specific name/value configuration. Note this maps to Graph's " +
							"`customKeyValueData` property, whose entries use a `name` field — distinct from custom_data above.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The configuration name.",
								},
								"value": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "The configuration value.",
								},
							},
						},
					},
					"identity_certificate_odata_bind": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: "Reference to a pre-existing iOS SCEP or PKCS certificate profile used to " +
							"authenticate to the VPN. Accepts a bare GUID or the full URL. Required when " +
							"authentication_method is certificate. Graph does not return this reference on read.",
					},
				},
			},
			"assignments": commonschemagraphbeta.DeviceConfigurationWithAllGroupAssignmentsAndFilterSchema(),
			"timeouts":    commonschema.ResourceTimeouts(ctx),
		},
	}
}
