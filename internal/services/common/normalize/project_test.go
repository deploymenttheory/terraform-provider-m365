package normalize

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectOntoShape(t *testing.T) {
	tests := []struct {
		name     string
		shape    string
		response string
		expected string
	}{
		{
			name:     "response extras are dropped",
			shape:    `{"cameraBlocked": true}`,
			response: `{"cameraBlocked": true, "airDropBlocked": false, "siriBlocked": false}`,
			expected: `{"cameraBlocked": true}`,
		},
		{
			name:     "response value wins over shape value",
			shape:    `{"cameraBlocked": true}`,
			response: `{"cameraBlocked": false}`,
			expected: `{"cameraBlocked": false}`,
		},
		{
			name:     "shape key missing from response is omitted so the diff surfaces",
			shape:    `{"cameraBlocked": true, "retiredProperty": "x"}`,
			response: `{"cameraBlocked": true}`,
			expected: `{"cameraBlocked": true}`,
		},
		{
			name:     "shape key present with null response keeps the null",
			shape:    `{"passcodeMinimumLength": 6}`,
			response: `{"passcodeMinimumLength": null}`,
			expected: `{"passcodeMinimumLength": null}`,
		},
		{
			name:     "null in both is preserved",
			shape:    `{"description": null}`,
			response: `{"description": null}`,
			expected: `{"description": null}`,
		},
		{
			name:     "three levels of nesting",
			shape:    `{"a": {"b": {"c": 1}}}`,
			response: `{"a": {"b": {"c": 2, "extra": 3}, "extra": 4}, "extra": 5}`,
			expected: `{"a": {"b": {"c": 2}}}`,
		},
		{
			name:     "nested object pruned to empty when shape is empty",
			shape:    `{"a": {}}`,
			response: `{"a": {"x": 1, "y": 2}}`,
			expected: `{"a": {}}`,
		},
		{
			name:     "empty shape prunes everything",
			shape:    `{}`,
			response: `{"a": 1, "b": 2}`,
			expected: `{}`,
		},
		{
			name:     "arrays of equal length project element-wise",
			shape:    `{"pages": [{"name": "one"}, {"name": "two"}]}`,
			response: `{"pages": [{"name": "one", "extra": 1}, {"name": "two", "extra": 2}]}`,
			expected: `{"pages": [{"name": "one"}, {"name": "two"}]}`,
		},
		{
			name:     "response array longer than shape keeps extra elements verbatim",
			shape:    `{"pages": [{"name": "one"}]}`,
			response: `{"pages": [{"name": "one", "extra": 1}, {"name": "two", "extra": 2}]}`,
			expected: `{"pages": [{"name": "one"}, {"name": "two", "extra": 2}]}`,
		},
		{
			name:     "response array shorter than shape yields the shorter array",
			shape:    `{"pages": [{"name": "one"}, {"name": "two"}]}`,
			response: `{"pages": [{"name": "one", "extra": 1}]}`,
			expected: `{"pages": [{"name": "one"}]}`,
		},
		{
			name:     "empty response array yields empty",
			shape:    `{"pages": [{"name": "one"}]}`,
			response: `{"pages": []}`,
			expected: `{"pages": []}`,
		},
		{
			name:     "array of scalars is taken verbatim",
			shape:    `{"roleScopeTagIds": ["0"]}`,
			response: `{"roleScopeTagIds": ["0", "1", "2"]}`,
			expected: `{"roleScopeTagIds": ["0", "1", "2"]}`,
		},
		{
			name:     "empty shape array keeps the whole response array",
			shape:    `{"pages": []}`,
			response: `{"pages": [{"name": "one"}, {"name": "two"}]}`,
			expected: `{"pages": [{"name": "one"}, {"name": "two"}]}`,
		},
		{
			name:     "type mismatch object shape vs array response takes response verbatim",
			shape:    `{"a": {"b": 1}}`,
			response: `{"a": [1, 2, 3]}`,
			expected: `{"a": [1, 2, 3]}`,
		},
		{
			name:     "type mismatch array shape vs object response takes response verbatim",
			shape:    `{"a": [1]}`,
			response: `{"a": {"b": 1, "c": 2}}`,
			expected: `{"a": {"b": 1, "c": 2}}`,
		},
		{
			name:     "shape leaf scalar vs response object takes response verbatim",
			shape:    `{"a": 1}`,
			response: `{"a": {"b": 1, "c": 2}}`,
			expected: `{"a": {"b": 1, "c": 2}}`,
		},
		{
			name:     "shape leaf null vs response object takes response verbatim",
			shape:    `{"a": null}`,
			response: `{"a": {"b": 1, "c": 2}}`,
			expected: `{"a": {"b": 1, "c": 2}}`,
		},
		{
			// Graph echoes @odata.type only where the enclosing collection's element type is abstract.
			// Where it is concrete the discriminator is redundant and dropped — but the operator still
			// had to supply it, so dropping it from state would break every apply.
			name:     "@odata.type is carried over when the response omits it",
			shape:    `{"@odata.type": "#microsoft.graph.iosHomeScreenPage", "displayName": "Page 1"}`,
			response: `{"displayName": "Page 1"}`,
			expected: `{"@odata.type": "#microsoft.graph.iosHomeScreenPage", "displayName": "Page 1"}`,
		},
		{
			name:     "a response @odata.type still wins over the shape's",
			shape:    `{"@odata.type": "#microsoft.graph.iosHomeScreenApp", "bundleID": "x"}`,
			response: `{"@odata.type": "#microsoft.graph.iosHomeScreenFolder", "bundleID": "x"}`,
			expected: `{"@odata.type": "#microsoft.graph.iosHomeScreenFolder", "bundleID": "x"}`,
		},
		{
			// The exception is scoped to @odata.type alone: a genuine setting the server drops must
			// still surface as a diff.
			name:     "other keys the response omits are still dropped",
			shape:    `{"@odata.type": "#microsoft.graph.iosHomeScreenPage", "retiredSetting": true}`,
			response: `{}`,
			expected: `{"@odata.type": "#microsoft.graph.iosHomeScreenPage"}`,
		},
		{
			// The real shape that surfaced this: only `icons` (abstract []iosHomeScreenItem) keeps its
			// discriminators; homeScreenPages, pages and apps are concrete and lose theirs.
			name: "home screen tree where Graph drops the concrete-collection discriminators",
			shape: `{
				"homeScreenPages": [
					{
						"@odata.type": "#microsoft.graph.iosHomeScreenPage",
						"displayName": "Productivity",
						"icons": [
							{
								"@odata.type": "#microsoft.graph.iosHomeScreenFolder",
								"displayName": "Utilities",
								"pages": [
									{
										"@odata.type": "#microsoft.graph.iosHomeScreenFolderPage",
										"displayName": "Utilities Page 1",
										"apps": [
											{"@odata.type": "#microsoft.graph.iosHomeScreenApp", "bundleID": "com.apple.mobilesafari"}
										]
									}
								]
							}
						]
					}
				]
			}`,
			response: `{
				"homeScreenPages": [
					{
						"displayName": "Productivity",
						"icons": [
							{
								"@odata.type": "#microsoft.graph.iosHomeScreenFolder",
								"displayName": "Utilities",
								"pages": [
									{
										"displayName": "Utilities Page 1",
										"apps": [
											{"bundleID": "com.apple.mobilesafari"}
										]
									}
								]
							}
						]
					}
				]
			}`,
			expected: `{
				"homeScreenPages": [
					{
						"@odata.type": "#microsoft.graph.iosHomeScreenPage",
						"displayName": "Productivity",
						"icons": [
							{
								"@odata.type": "#microsoft.graph.iosHomeScreenFolder",
								"displayName": "Utilities",
								"pages": [
									{
										"@odata.type": "#microsoft.graph.iosHomeScreenFolderPage",
										"displayName": "Utilities Page 1",
										"apps": [
											{"@odata.type": "#microsoft.graph.iosHomeScreenApp", "bundleID": "com.apple.mobilesafari"}
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
			name:     "nested @odata.type discriminators survive projection",
			shape:    `{"homeScreenPages": [{"icons": [{"@odata.type": "#microsoft.graph.iosHomeScreenApp", "bundleID": "com.apple.mobilesafari"}]}]}`,
			response: `{"homeScreenPages": [{"displayName": null, "icons": [{"@odata.type": "#microsoft.graph.iosHomeScreenApp", "bundleID": "com.apple.mobilesafari", "isWebClip": false}]}]}`,
			expected: `{"homeScreenPages": [{"icons": [{"@odata.type": "#microsoft.graph.iosHomeScreenApp", "bundleID": "com.apple.mobilesafari"}]}]}`,
		},
		{
			name:     "mixed-type array projects element-wise by position",
			shape:    `{"items": [{"a": 1}, "scalar", [1]]}`,
			response: `{"items": [{"a": 2, "b": 3}, "other", [4, 5]]}`,
			expected: `{"items": [{"a": 2}, "other", [4, 5]]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var shape, response, expected any
			require.NoError(t, json.Unmarshal([]byte(tt.shape), &shape))
			require.NoError(t, json.Unmarshal([]byte(tt.response), &response))
			require.NoError(t, json.Unmarshal([]byte(tt.expected), &expected))

			actual := ProjectOntoShape(shape, response)

			assert.Equal(t, expected, actual)
		})
	}
}

// The function must not mutate either argument, since callers hold the prior configuration and
// the parsed response for other purposes.
func TestProjectOntoShapeDoesNotMutateArguments(t *testing.T) {
	shapeJSON := `{"a": {"b": 1}, "pages": [{"name": "one"}]}`
	responseJSON := `{"a": {"b": 2, "extra": 3}, "pages": [{"name": "two", "extra": 4}], "dropped": 5}`

	var shape, response any
	require.NoError(t, json.Unmarshal([]byte(shapeJSON), &shape))
	require.NoError(t, json.Unmarshal([]byte(responseJSON), &response))

	_ = ProjectOntoShape(shape, response)

	var shapeAfter, responseAfter any
	require.NoError(t, json.Unmarshal([]byte(shapeJSON), &shapeAfter))
	require.NoError(t, json.Unmarshal([]byte(responseJSON), &responseAfter))

	assert.Equal(t, shapeAfter, shape, "shape was mutated")
	assert.Equal(t, responseAfter, response, "response was mutated")
}

// Top-level scalars and nulls are handled without a surrounding object, since ProjectOntoShape
// recurses on itself and must terminate cleanly at any node type.
func TestProjectOntoShapeTopLevelNonObjects(t *testing.T) {
	assert.Equal(t, "response", ProjectOntoShape("shape", "response"))
	assert.Equal(t, nil, ProjectOntoShape(nil, nil))
	assert.Equal(t, "response", ProjectOntoShape(nil, "response"))
	assert.Equal(t, nil, ProjectOntoShape("shape", nil))
	assert.Equal(t, []any{1, 2}, ProjectOntoShape([]any{1}, []any{1, 2}))
	assert.Equal(t, map[string]any{}, ProjectOntoShape(map[string]any{}, map[string]any{"a": 1}))
}

// The realistic case that motivates the function: a configuration declaring one property against
// a response carrying the full property surface.
func TestProjectOntoShapeFullSurfaceResponse(t *testing.T) {
	shape := map[string]any{"cameraBlocked": true}

	response := map[string]any{
		"cameraBlocked":                 true,
		"airDropBlocked":                false,
		"appStoreBlocked":               false,
		"siriBlocked":                   false,
		"passcodeMinimumLength":         nil,
		"mediaContentRatingApps":        "allAllowed",
		"emailInDomainSuffixes":         []any{},
		"kioskModeAllowAutoLock":        false,
		"kioskModeBlockAutoLock":        false,
		"wiFiConnectOnlyToConfiguredNe": false,
	}

	actual := ProjectOntoShape(shape, response)

	assert.Equal(t, map[string]any{"cameraBlocked": true}, actual,
		"a single-property configuration must not pick up server defaults")
}
