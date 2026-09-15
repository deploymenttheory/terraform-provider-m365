package graphBetaIosDeviceConfigurationTemplatesJson

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// validate runs the envelope validator against a settings_json value and returns the diagnostics.
func validate(t *testing.T, settingsJSON types.String) validator.StringResponse {
	t.Helper()

	req := validator.StringRequest{
		Path:        path.Root("settings_json"),
		ConfigValue: settingsJSON,
	}
	resp := validator.StringResponse{}

	settingsJSONEnvelope().ValidateString(context.Background(), req, &resp)

	return resp
}

func TestSettingsJSONEnvelopeValidatorRejectsRootEnvelopeKeys(t *testing.T) {
	tests := []struct {
		name         string
		settingsJSON string
		wantInError  []string
	}{
		{
			name:         "@odata.type at root",
			settingsJSON: `{"@odata.type": "#microsoft.graph.iosGeneralDeviceConfiguration", "cameraBlocked": true}`,
			wantInError:  []string{"@odata.type"},
		},
		{
			name:         "displayName at root",
			settingsJSON: `{"displayName": "My Profile", "cameraBlocked": true}`,
			wantInError:  []string{"displayName"},
		},
		{
			name:         "description at root",
			settingsJSON: `{"description": "Something", "cameraBlocked": true}`,
			wantInError:  []string{"description"},
		},
		{
			name:         "roleScopeTagIds at root",
			settingsJSON: `{"roleScopeTagIds": ["0"], "cameraBlocked": true}`,
			wantInError:  []string{"roleScopeTagIds"},
		},
		{
			name: "all four at once are all reported",
			settingsJSON: `{
				"@odata.type": "#microsoft.graph.iosGeneralDeviceConfiguration",
				"displayName": "My Profile",
				"description": "Something",
				"roleScopeTagIds": ["0"],
				"cameraBlocked": true
			}`,
			wantInError: []string{
				"@odata.type",
				"description",
				"displayName",
				"roleScopeTagIds",
			},
		},
		{
			name:         "envelope key with a null value is still rejected",
			settingsJSON: `{"description": null, "cameraBlocked": true}`,
			wantInError:  []string{"description"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := validate(t, types.StringValue(tt.settingsJSON))

			require.True(t, resp.Diagnostics.HasError(), "expected a validation error")

			detail := resp.Diagnostics.Errors()[0].Detail()
			for _, want := range tt.wantInError {
				assert.Contains(t, detail, want,
					"diagnostic should name the offending key")
			}
		})
	}
}

func TestSettingsJSONEnvelopeValidatorAcceptsValidSettings(t *testing.T) {
	tests := []struct {
		name         string
		settingsJSON string
	}{
		{
			name:         "plain settings",
			settingsJSON: `{"cameraBlocked": true, "airDropBlocked": false}`,
		},
		{
			name:         "empty object",
			settingsJSON: `{}`,
		},
		{
			name: "nested @odata.type discriminators survive — this is the critical case",
			settingsJSON: `{
				"homeScreenPages": [
					{
						"@odata.type": "#microsoft.graph.iosHomeScreenPage",
						"displayName": "Page 1",
						"icons": [
							{
								"@odata.type": "#microsoft.graph.iosHomeScreenApp",
								"displayName": "Safari",
								"bundleID": "com.apple.mobilesafari"
							},
							{
								"@odata.type": "#microsoft.graph.iosHomeScreenFolder",
								"displayName": "Utilities",
								"pages": [
									{
										"@odata.type": "#microsoft.graph.iosHomeScreenFolderPage",
										"displayName": "Utilities Page 1",
										"apps": [
											{
												"@odata.type": "#microsoft.graph.iosHomeScreenApp",
												"displayName": "Calculator",
												"bundleID": "com.apple.calculator"
											}
										]
									}
								]
							}
						]
					}
				]
			}`,
		},
		{
			name: "nested singleSignOnExtension discriminator",
			settingsJSON: `{
				"singleSignOnExtension": {
					"@odata.type": "#microsoft.graph.iosKerberosSingleSignOnExtension",
					"realm": "EXAMPLE.COM"
				}
			}`,
		},
		{
			// Both contentFilterSettings implementations, verified against a live tenant. They have
			// entirely different property sets, so the discriminator is the only thing distinguishing
			// them and must survive.
			name: "nested contentFilterSettings autoFilter discriminator",
			settingsJSON: `{
				"contentFilterSettings": {
					"@odata.type": "#microsoft.graph.iosWebContentFilterAutoFilter",
					"allowedUrls": ["https://microsoft.com"],
					"blockedUrls": ["blocked.com"]
				}
			}`,
		},
		{
			name: "nested contentFilterSettings specificWebsitesAccess with iosBookmark entries",
			settingsJSON: `{
				"contentFilterSettings": {
					"@odata.type": "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess",
					"websiteList": [
						{
							"@odata.type": "#microsoft.graph.iosBookmark",
							"url": "https://intranet.example.com",
							"displayName": "Company Intranet",
							"bookmarkFolder": "Work"
						}
					]
				}
			}`,
		},
		{
			// wallpaperImage is a mimeContent {type, value}. Note "type" is a plausible-looking key
			// that must NOT be mistaken for an envelope key — only the four root keys are rejected.
			name: "nested wallpaperImage mimeContent",
			settingsJSON: `{
				"wallpaperDisplayLocation": "lockAndHomeScreens",
				"wallpaperImage": {
					"type": "image/jpeg",
					"value": "/9j/4Q//9k="
				}
			}`,
		},
		{
			name: "nested displayName is fine — only the root is checked",
			settingsJSON: `{
				"homeScreenPages": [{"displayName": "Page 1", "icons": []}]
			}`,
		},
		{
			name: "nested description is fine",
			settingsJSON: `{
				"notificationSettings": [
					{
						"@odata.type": "#microsoft.graph.iosNotificationSettings",
						"bundleID": "com.example.app",
						"description": "Example app notifications"
					}
				]
			}`,
		},
		{
			name:         "similar-but-different root keys are not rejected",
			settingsJSON: `{"displayNameOverride": "x", "descriptionText": "y", "odataType": "z"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := validate(t, types.StringValue(tt.settingsJSON))

			assert.False(t, resp.Diagnostics.HasError(),
				"expected no validation error, got: %v", resp.Diagnostics.Errors())
		})
	}
}

// Malformed and non-object JSON belong to JSONSchemaValidator; this validator must stay silent
// rather than emitting a confusing second diagnostic about envelope keys.
func TestSettingsJSONEnvelopeValidatorIgnoresUnparseableInput(t *testing.T) {
	tests := []struct {
		name         string
		settingsJSON types.String
	}{
		{name: "null value", settingsJSON: types.StringNull()},
		{name: "unknown value", settingsJSON: types.StringUnknown()},
		{name: "malformed JSON", settingsJSON: types.StringValue(`{"cameraBlocked": tru`)},
		{name: "JSON array rather than object", settingsJSON: types.StringValue(`[1, 2, 3]`)},
		{name: "JSON string rather than object", settingsJSON: types.StringValue(`"just a string"`)},
		{name: "empty string", settingsJSON: types.StringValue(``)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := validate(t, tt.settingsJSON)

			assert.False(t, resp.Diagnostics.HasError(),
				"expected no diagnostics, got: %v", resp.Diagnostics.Errors())
		})
	}
}

func TestSettingsJSONEnvelopeValidatorDescription(t *testing.T) {
	ctx := context.Background()
	v := settingsJSONEnvelope()

	description := v.Description(ctx)

	for _, key := range envelopeKeys {
		assert.Contains(t, description, key,
			"description should name every rejected key so the schema docs are complete")
	}
	assert.Equal(t, description, v.MarkdownDescription(ctx))
}
