package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	stackstate_admin_api "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_admin_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

var (
	transactionId = int64(123)
	timestamp = int64(321)
	epochTx = stackstate_admin_api.EpochTx {
		EpochTx: &transactionId,
	}
	windowMs = stackstate_admin_api.WindowMs {
		WindowMs: &timestamp,
	}
)

func TestRetentionFetchesWindowAndEpoch(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := RetentionCommand(&cli.Deps)

	cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowResponse.Result = windowMs
	cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochResponse.Result = epochTx

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 1)

	expectedStrings := []string {
		"Retention window: 321 milliseconds\nEpoch transactionId: 123",
	}
	assert.Equal(t, expectedStrings, *cli.MockPrinter.SuccessCalls)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 2)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 2)

	expectedJsonCall := []map[string]interface{} {
		{
			"retention-window": timestamp,
			"epoch": transactionId,
		},
	}

	assert.Equal(t, expectedJsonCall, *cli.MockPrinter.PrintJsonCalls)
}

func TestRetentionHandlesErrorsWhenFetching(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := RetentionCommand(&cli.Deps)

	windowError := err("Unauthorized")

	cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowResponse.Error = windowError
	cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochResponse.Result = epochTx

	_, resp1 := di.ExecuteCommandWithContext(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 0)

	assert.Equal(t, windowError.Error(), resp1.Error())

	epochError := err("Please stop")

	cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowResponse.Error = nil
	cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowResponse.Result = windowMs
	cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochResponse.Error = epochError

	_, resp2 := di.ExecuteCommandWithContext(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 2)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 1)

	assert.Equal(t, epochError.Error(), resp2.Error())
}

func err(info string) error {
	return fmt.Errorf("Some API error: %s", info)
}
