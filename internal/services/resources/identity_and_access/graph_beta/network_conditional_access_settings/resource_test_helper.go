package graphBetaNetworkConditionalAccessSettings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/exists"
)

// NetworkConditionalAccessSettingsTestResource checks the singleton's availability.
type NetworkConditionalAccessSettingsTestResource struct{}

func (r NetworkConditionalAccessSettingsTestResource) Exists(
	ctx context.Context,
	_ any,
	state *terraform.InstanceState,
) (*bool, error) {
	//nolint:wrapcheck // The generic helper adds operation context.
	return exists.CheckResourceExists(
		ctx,
		state,
		func(client *msgraphbetasdk.GraphServiceClient, ctx context.Context, _ *terraform.InstanceState) error {
			r := &NetworkConditionalAccessSettingsResource{client: client}
			_, err := r.getSettings(ctx)
			return err
		},
	)
}
