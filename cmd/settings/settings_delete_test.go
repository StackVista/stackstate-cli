package settings

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestSettingsDelete(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := SettingsDeleteCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "1,2,3")

	calls := *cli.MockClient.ApiMocks.NodeApi.DeleteCalls
	assert.Len(t, calls, 3)
	assert.Equal(t, int64(1), calls[0].PnodeId)
	assert.Equal(t, int64(15), *calls[0].PtimeoutSeconds)
	assert.Equal(t, int64(2), calls[1].PnodeId)
	assert.Equal(t, int64(15), *calls[1].PtimeoutSeconds)
	assert.Equal(t, int64(3), calls[2].PnodeId)
	assert.Equal(t, int64(15), *calls[2].PtimeoutSeconds)

	expectedSuccessCalls := []string{
		"The following nodes were deleted: 1, 2, 3",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.SuccessCalls)
}

func TestSettingsDeleteJustOne(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := SettingsDeleteCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "23")

	calls := *cli.MockClient.ApiMocks.NodeApi.DeleteCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, int64(23), calls[0].PnodeId)
	assert.Equal(t, int64(15), *calls[0].PtimeoutSeconds)

	expectedSuccessCalls := []string{
		"The following nodes were deleted: 23",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.SuccessCalls)
}

func TestSettingsDeleteTimeout(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := SettingsDeleteCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "5", "--timeout", "30")

	calls := *cli.MockClient.ApiMocks.NodeApi.DeleteCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, int64(5), calls[0].PnodeId)
	assert.Equal(t, int64(30), *calls[0].PtimeoutSeconds)

	expectedSuccessCalls := []string{
		"The following nodes were deleted: 5",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.SuccessCalls)
}

func TestSettingsDeleteJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := SettingsDeleteCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "1,2,3", "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"deleted": []int64{1, 2, 3},
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
