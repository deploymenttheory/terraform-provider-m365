package graphBetaNetworkConditionalAccessSettings

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	kiotahttp "github.com/microsoft/kiota-http-go"
	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
	"github.com/stretchr/testify/require"

	providerclient "github.com/deploymenttheory/terraform-provider-microsoft365/internal/client"
)

func testModel(status types.String) NetworkConditionalAccessSettingsResourceModel {
	return NetworkConditionalAccessSettingsResourceModel{
		ID:              types.StringValue(singletonID),
		SignalingStatus: status,
		Timeouts: timeouts.Value{
			Object: types.ObjectValueMust(
				map[string]attr.Type{
					"create": types.StringType,
					"read":   types.StringType,
					"update": types.StringType,
					"delete": types.StringType,
				},
				map[string]attr.Value{
					"create": types.StringValue("4s"),
					"read":   types.StringValue("4s"),
					"update": types.StringValue("4s"),
					"delete": types.StringNull(),
				},
			),
		},
	}
}

func testState(
	t *testing.T,
	r *NetworkConditionalAccessSettingsResource,
	model NetworkConditionalAccessSettingsResourceModel,
) tfsdk.State {
	t.Helper()
	var schema resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &schema)
	state := tfsdk.State{Schema: schema.Schema}
	require.False(t, state.Set(context.Background(), model).HasError())
	return state
}

func testClient(t *testing.T, h http.HandlerFunc) *NetworkConditionalAccessSettingsResource {
	t.Helper()
	server := httptest.NewServer(h)
	t.Cleanup(server.Close)
	t.Setenv("TF_ACC", "")
	require.NoError(t, os.Unsetenv("TF_ACC"))
	httpClient, err := providerclient.ConfigureGraphClientOptions(
		context.Background(),
		&providerclient.ProviderData{
			ClientOptions: &providerclient.ClientOptions{
				EnableRetry:       true,
				MaxRetries:        2,
				RetryDelaySeconds: 1,
				EnableCompression: true,
			},
		},
	)
	require.NoError(t, err)
	adapter, err := kiotahttp.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		&authentication.AnonymousAuthenticationProvider{},
		nil,
		nil,
		httpClient,
	)
	require.NoError(t, err)
	adapter.SetBaseUrl(server.URL)
	r := NewNetworkConditionalAccessSettingsResource().(*NetworkConditionalAccessSettingsResource)
	r.client = msgraphbetasdk.NewGraphServiceClient(adapter)
	return r
}

func respond(t *testing.T, w http.ResponseWriter, code int, body string) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if body != "" {
		_, err := w.Write([]byte(body))
		require.NoError(t, err)
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_10_ReadErrorsPreserveState(t *testing.T) {
	for _, code := range []int{400, 401, 403, 404, 500} {
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			r := testClient(t, func(w http.ResponseWriter, q *http.Request) {
				respond(
					t,
					w,
					code,
					`{"error":{"code":"TestFailure","message":"Synthetic failure"}}`,
				)
			})
			state := testState(t, r, testModel(types.StringValue("enabled")))
			resp := resource.ReadResponse{State: state}
			r.Read(context.Background(), resource.ReadRequest{State: state}, &resp)
			require.True(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
			require.True(t, state.Raw.Equal(resp.State.Raw), "read error discarded state")
		})
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_11_MalformedReadPreservesState(t *testing.T) {
	for _, body := range []string{`{}`, `null`, `{"signalingStatus":null}`, `{"signalingStatus":"unknownFutureValue"}`} {
		t.Run(body, func(t *testing.T) {
			r := testClient(
				t,
				func(w http.ResponseWriter, q *http.Request) { respond(t, w, 200, body) },
			)
			state := testState(t, r, testModel(types.StringValue("enabled")))
			resp := resource.ReadResponse{State: state}
			r.Read(context.Background(), resource.ReadRequest{State: state}, &resp)
			require.True(t, resp.Diagnostics.HasError())
			require.True(t, state.Raw.Equal(resp.State.Raw))
		})
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_12_PatchErrors(t *testing.T) {
	for _, code := range []int{400, 403, 404} {
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			r := testClient(t, func(w http.ResponseWriter, q *http.Request) {
				require.Equal(t, "PATCH", q.Method)
				respond(t, w, code, `{"error":{"code":"TestFailure","message":"Write failed"}}`)
			})
			before := testState(t, r, testModel(types.StringValue("enabled")))
			desired := testState(t, r, testModel(types.StringValue("disabled")))
			plan := tfsdk.Plan{Schema: desired.Schema, Raw: desired.Raw}
			config := tfsdk.Config{Schema: desired.Schema, Raw: desired.Raw}
			update := resource.UpdateResponse{State: before}
			r.Update(
				context.Background(),
				resource.UpdateRequest{State: before, Plan: plan, Config: config},
				&update,
			)
			require.True(t, update.Diagnostics.HasError())
			require.True(t, before.Raw.Equal(update.State.Raw))
			empty := testState(t, r, testModel(types.StringValue("disabled")))
			empty.RemoveResource(context.Background())
			create := resource.CreateResponse{State: empty}
			r.Create(
				context.Background(),
				resource.CreateRequest{Plan: plan, Config: config},
				&create,
			)
			require.True(t, create.Diagnostics.HasError())
			require.True(t, create.State.Raw.IsNull())
		})
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_13_ReadbackFailureRetainsID(t *testing.T) {
	for _, operation := range []string{"create", "update"} {
		for _, code := range []int{403, 404, 200} {
			t.Run(fmt.Sprintf("%s/%d", operation, code), func(t *testing.T) {
				writes := 0
				r := testClient(t, func(w http.ResponseWriter, q *http.Request) {
					if q.Method == "PATCH" {
						writes++
						respond(t, w, 204, "")
						return
					}
					respond(
						t,
						w,
						code,
						`{"error":{"code":"TestFailure","message":"Readback failed"}}`,
					)
				})
				before := testState(t, r, testModel(types.StringValue("enabled")))
				desired := testState(t, r, testModel(types.StringValue("disabled")))
				plan := tfsdk.Plan{Schema: desired.Schema, Raw: desired.Raw}
				config := tfsdk.Config{Schema: desired.Schema, Raw: desired.Raw}
				var result tfsdk.State
				if operation == "create" {
					resp := resource.CreateResponse{State: before}
					r.Create(
						context.Background(),
						resource.CreateRequest{Plan: plan, Config: config},
						&resp,
					)
					require.True(t, resp.Diagnostics.HasError())
					result = resp.State
				} else {
					resp := resource.UpdateResponse{State: before}
					r.Update(
						context.Background(),
						resource.UpdateRequest{State: before, Plan: plan, Config: config},
						&resp,
					)
					require.True(t, resp.Diagnostics.HasError())
					result = resp.State
				}
				var actual NetworkConditionalAccessSettingsResourceModel
				require.False(t, result.Get(context.Background(), &actual).HasError())
				require.Equal(t, singletonID, actual.ID.ValueString())
				require.Equal(t, "disabled", actual.SignalingStatus.ValueString())
				require.Equal(t, 1, writes)
			})
		}
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_14_OmittedAndUnchangedUpdatesNeverWrite(
	t *testing.T,
) {
	for _, omitted := range []bool{true, false} {
		t.Run(fmt.Sprint(omitted), func(t *testing.T) {
			r := testClient(t, func(w http.ResponseWriter, q *http.Request) {
				t.Errorf("unexpected request %s", q.Method)
				respond(t, w, 500, `{}`)
			})
			before := testState(t, r, testModel(types.StringValue("disabled")))
			configModel := testModel(types.StringValue("disabled"))
			if omitted {
				configModel.SignalingStatus = types.StringNull()
			}
			config := testState(t, r, configModel)
			resp := resource.UpdateResponse{State: before}
			r.Update(
				context.Background(),
				resource.UpdateRequest{
					State:  before,
					Plan:   tfsdk.Plan{Schema: before.Schema, Raw: before.Raw},
					Config: tfsdk.Config{Schema: config.Schema, Raw: config.Raw},
				},
				&resp,
			)
			require.False(t, resp.Diagnostics.HasError())
			require.True(t, before.Raw.Equal(resp.State.Raw))
			deleted := resource.DeleteResponse{State: resp.State}
			r.Delete(context.Background(), resource.DeleteRequest{State: resp.State}, &deleted)
			require.False(t, deleted.Diagnostics.HasError())
			require.True(t, deleted.State.Raw.IsNull())
		})
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_15_PartialPatchAndConsistency(t *testing.T) {
	writes, reads := 0, 0
	r := testClient(t, func(w http.ResponseWriter, q *http.Request) {
		require.Equal(t, "/networkAccess/settings/conditionalAccess", q.URL.Path)
		switch q.Method {
		case "PATCH":
			require.Empty(t, q.Header.Get("Content-Encoding"))
			var body map[string]any
			require.NoError(t, json.NewDecoder(q.Body).Decode(&body))
			require.Equal(t, map[string]any{"signalingStatus": "disabled"}, body)
			writes++
			respond(t, w, 204, "")
		case "GET":
			reads++
			status := "disabled"
			if reads == 1 {
				status = "enabled"
			}
			respond(
				t,
				w,
				200,
				strings.Replace(successFixture(t), `"enabled"`, fmt.Sprintf("%q", status), 1),
			)
		default:
			t.Errorf("unexpected %s", q.Method)
		}
	})
	model := testModel(types.StringValue("disabled"))
	model.Timeouts.Object = types.ObjectValueMust(
		model.Timeouts.AttributeTypes(context.Background()),
		map[string]attr.Value{
			"create": types.StringValue("10s"),
			"read":   types.StringValue("10s"),
			"update": types.StringValue("10s"),
			"delete": types.StringNull(),
		},
	)
	state := testState(t, r, model)
	resp := resource.CreateResponse{State: state}
	r.Create(
		context.Background(),
		resource.CreateRequest{
			Plan:   tfsdk.Plan{Schema: state.Schema, Raw: state.Raw},
			Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw},
		},
		&resp,
	)
	require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
	require.Equal(t, 1, writes)
	require.Equal(t, 2, reads)
}

func TestUnitResourceNetworkConditionalAccessSettings_16_ImportValidation(t *testing.T) {
	r := NewNetworkConditionalAccessSettingsResource().(*NetworkConditionalAccessSettingsResource)
	for _, id := range []string{"conditionalAccess", "", "invalid", "11111111-1111-1111-1111-111111111111"} {
		t.Run(id, func(t *testing.T) {
			state := testState(t, r, testModel(types.StringNull()))
			resp := resource.ImportStateResponse{State: state}
			r.ImportState(context.Background(), resource.ImportStateRequest{ID: id}, &resp)
			require.Equal(t, id != "conditionalAccess", resp.Diagnostics.HasError())
			if id == "conditionalAccess" {
				var value string
				require.False(
					t,
					resp.State.GetAttribute(context.Background(), path.Root("id"), &value).
						HasError(),
				)
				require.Equal(t, id, value)
			}
		})
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_17_UnknownIsNotDefaulted(t *testing.T) {
	r := NewNetworkConditionalAccessSettingsResource().(*NetworkConditionalAccessSettingsResource)
	state := testState(t, r, testModel(types.StringValue("disabled")))
	for _, unknown := range []bool{true, false} {
		configValue := types.StringNull()
		if unknown {
			configValue = types.StringUnknown()
		}
		resp := planmodifier.StringResponse{PlanValue: types.StringUnknown()}
		stringplanmodifier.UseStateForUnknown().
			PlanModifyString(context.Background(), planmodifier.StringRequest{State: state, StateValue: types.StringValue("disabled"), PlanValue: types.StringUnknown(), ConfigValue: configValue}, &resp)
		require.Equal(t, unknown, resp.PlanValue.IsUnknown())
		if !unknown {
			require.Equal(t, "disabled", resp.PlanValue.ValueString())
		}
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_18_ConstructRejectsUnresolved(t *testing.T) {
	for _, v := range []types.String{types.StringNull(), types.StringUnknown(), types.StringValue("invalid")} {
		_, err := constructResource(v)
		require.Error(t, err)
	}
}

func TestUnitResourceNetworkConditionalAccessSettings_19_IdentityImport(t *testing.T) {
	ctx := context.Background()
	r := NewNetworkConditionalAccessSettingsResource().(*NetworkConditionalAccessSettingsResource)
	var schema resource.IdentitySchemaResponse
	r.IdentitySchema(ctx, resource.IdentitySchemaRequest{}, &schema)
	identity := tfsdk.ResourceIdentity{Schema: schema.IdentitySchema}
	require.False(t, identity.Set(ctx, struct {
		ID string `tfsdk:"id"`
	}{singletonID}).HasError())
	state := testState(t, r, testModel(types.StringNull()))
	resp := resource.ImportStateResponse{State: state, Identity: &identity}
	r.ImportState(ctx, resource.ImportStateRequest{Identity: &identity}, &resp)
	require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)
	var id string
	require.False(t, resp.State.GetAttribute(ctx, path.Root("id"), &id).HasError())
	require.Equal(t, singletonID, id)
}

func successFixture(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	body, err := os.ReadFile(
		filepath.Join(
			filepath.Dir(filename),
			"tests/responses/validate_get/get_conditional_access_settings_success.json",
		),
	)
	require.NoError(t, err)
	return string(body)
}
