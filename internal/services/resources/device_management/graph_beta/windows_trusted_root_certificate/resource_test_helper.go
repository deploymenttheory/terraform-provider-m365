package graphBetaWindowsTrustedRootCertificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	msgraphbetasdk "github.com/microsoftgraph/msgraph-beta-sdk-go"

	"github.com/deploymenttheory/terraform-provider-microsoft365/internal/acceptance/exists"
)

type WindowsTrustedRootCertificateTestResource struct{}

func (r WindowsTrustedRootCertificateTestResource) Exists(
	ctx context.Context,
	_ any,
	state *terraform.InstanceState,
) (*bool, error) {
	//nolint:wrapcheck // The shared helper supplies contextual errors.
	return exists.CheckResourceExists(
		ctx,
		state,
		func(client *msgraphbetasdk.GraphServiceClient, ctx context.Context, state *terraform.InstanceState) error {
			_, err := client.DeviceManagement().
				DeviceConfigurations().
				ByDeviceConfigurationId(state.ID).
				Get(ctx, nil)
			return err
		},
	)
}
