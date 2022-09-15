package license

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupUpdateCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := UpdateCommand(&cli.Deps)

	return &cli, cmd
}

func TestUpdateCommand(t *testing.T) {
	results := map[string]stackstate_api.SubscriptionState{
		"MissingUnlicensed": {
			UnlicensedSubscription: stackstate_api.NewUnlicensedSubscription("UnlicensedSubscription", "Missing"),
		},
		"MissingInvalid": {
			UnlicensedSubscription: stackstate_api.NewUnlicensedSubscription("UnlicensedSubscription", "Invalid"),
		},
		"LicensedWithoutExpiration": {
			LicensedSubscription: stackstate_api.NewLicensedSubscription("LicensedSubscription", *stackstate_api.NewSubscription("tenant-1", "SUBSCRIPTION")),
		},
		"LicensedWithExpiration": {
			LicensedSubscription: stackstate_api.NewLicensedSubscription("LicensedSubscription", *expiringSubscription("tenant-1", 213, "SUBSCRIPTION")),
		},
	}

	for n, r := range results {
		t.Run(n, func(t *testing.T) {
			res := r
			cli, cmd := setupUpdateCmd(t)

			cli.MockClient.ApiMocks.SubscriptionApi.PostSubscriptionResponse.Result = r
			di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json", "--key", "the-license-key")

			assert.Len(t, *cli.MockClient.ApiMocks.SubscriptionApi.PostSubscriptionCalls, 1)
			call := (*cli.MockClient.ApiMocks.SubscriptionApi.PostSubscriptionCalls)[0]
			assert.Equal(t, "the-license-key", call.PnewLicense.Key)
			assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
			assert.Equal(t, &res, (*cli.MockPrinter.PrintJsonCalls)[0]["subscription"])
		})
	}
}
