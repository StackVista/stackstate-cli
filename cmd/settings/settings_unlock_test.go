package settings

import (
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestSettingsUnlock(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := SettingsUnlockCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "1,2,3")

	calls := *cli.MockClient.ApiMocks.NodeAPI.UnlockCalls
	assert.Len(t, calls, 3)
	assert.Equal(t, int64(1), calls[0].PnodeId)
	assert.Equal(t, int64(2), calls[1].PnodeId)
	assert.Equal(t, int64(3), calls[2].PnodeId)

	expectedSuccessCalls := []string{
		"The following nodes were unlocked: 1, 2, 3",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.SuccessCalls)
}

func TestSettingsUnlockJustOne(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := SettingsUnlockCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "23")

	calls := *cli.MockClient.ApiMocks.NodeAPI.UnlockCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, int64(23), calls[0].PnodeId)

	expectedSuccessCalls := []string{
		"The following nodes were unlocked: 23",
	}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.SuccessCalls)
}

func TestSettingsUnlockJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := SettingsUnlockCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "1,2,3", "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"unlocked": []int64{1, 2, 3},
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
