package license

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupShowCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := ShowCommand(&cli.Deps)

	return &cli, cmd
}

func TestShowCommand(t *testing.T) {
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
			cli, cmd := setupShowCmd(t)
			cli.MockClient.ApiMocks.SubscriptionApi.GetSubscriptionResponse.Result = r
			di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

			assert.Len(t, *cli.MockPrinter.PrintJsonCalls, 1)
			assert.Equal(t, &res, (*cli.MockPrinter.PrintJsonCalls)[0]["subscription"])
		})
	}
}

func expiringSubscription(tenant string, expirationDate int64, plan string) *stackstate_api.Subscription {
	s := stackstate_api.NewSubscription(tenant, plan)
	s.ExpiryTimestampMs = &expirationDate
	return s
}
