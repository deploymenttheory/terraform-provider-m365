package errors_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/constants"
	errors "github.com/deploymenttheory/terraform-provider-microsoft365/internal/services/common/errors/kiota"
)

func TestUnitGraphErrorOptions_SingletonNotFound(t *testing.T) {
	for _, preserve := range []bool{false, true} {
		ctx := context.Background()
		state := tfsdk.State{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{Computed: true},
				},
			},
		}
		require.False(t, state.Set(ctx, struct {
			ID string `tfsdk:"id"`
		}{"conditionalAccess"}).HasError())
		resp := resource.ReadResponse{State: state}
		err := abstractions.NewApiError()
		err.SetStatusCode(404)
		errors.HandleKiotaGraphErrorWithOptions(
			ctx,
			err,
			&resp,
			constants.TfOperationRead,
			nil,
			errors.GraphErrorOptions{PreserveStateOnReadNotFound: preserve},
		)
		require.Equal(t, preserve, resp.Diagnostics.HasError())
		require.Equal(t, !preserve, resp.State.Raw.IsNull())
		if preserve {
			require.True(t, state.Raw.Equal(resp.State.Raw))
		}
	}
}
