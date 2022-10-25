package graph

import (
	"fmt"
	"testing"

	stackstate_admin_api "github.com/stackvista/stackstate-cli/generated/stackstate_admin_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	transactionId = int64(123)
	timestamp     = int64(321)
	duration      = "5h"
	newTimestamp  = int64(5 * 3600 * 1000)
	epochTx       = *stackstate_admin_api.NewEpochTx(transactionId)
	windowMs      = *stackstate_admin_api.NewWindowMs(timestamp)
	newWindowMs   = *stackstate_admin_api.NewWindowMs(newTimestamp)
)

func TestRetentionFetchesWindowAndEpoch(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := RetentionCommand(&cli.Deps)

	cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowResponse.Result = windowMs
	cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochResponse.Result = epochTx

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 1)

	expectedStrings := []string{
		fmt.Sprintf("Retention window: %d milliseconds\nEpoch transactionId: %d", timestamp, transactionId),
	}
	assert.Equal(t, expectedStrings, *cli.MockPrinter.SuccessCalls)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 2)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 2)

	expectedJsonCall := []map[string]interface{}{
		{
			"retention-window": timestamp,
			"epoch":            transactionId,
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

func TestRetentionSetsWindow(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := RetentionCommand(&cli.Deps)

	cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowResponse.Result = windowMs
	cli.MockClient.ApiMocks.RetentionApi.SetRetentionWindowResponse.Result = newWindowMs
	cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochResponse.Result = epochTx

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.SetRetentionWindowCalls, 0)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 1)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-s", duration, "--schedule-removal")

	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionEpochCalls, 2)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.SetRetentionWindowCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.RetentionApi.GetRetentionWindowCalls, 1)

	setCalls := *cli.MockClient.ApiMocks.RetentionApi.SetRetentionWindowCalls

	assert.Equal(t, true, *setCalls[0].PscheduleRemoval)
	assert.Equal(t, newWindowMs, *setCalls[0].PwindowMs)

	expectedStrings := []string{
		fmt.Sprintf("Retention window: %d milliseconds\nEpoch transactionId: %d", timestamp, transactionId),
		fmt.Sprintf("Retention window: %d milliseconds\nEpoch transactionId: %d", newTimestamp, transactionId),
	}
	assert.Equal(t, expectedStrings, *cli.MockPrinter.SuccessCalls)
}

func err(info string) error {
	return fmt.Errorf("Some API error: %s", info)
}
