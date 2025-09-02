package stackpack

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	stackPackName           = "zabbix"
	stackPackNextVersion    = "2.0.0"
	stackPackCurrentVersion = "1.0.0"
	strategyFlag            = "overwrite"
)

func setupStackPackUpgradeCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUpgradeCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{
		{
			Name: stackPackName,
			NextVersion: &stackstate_api.FullStackPack{
				Version: stackPackNextVersion,
			},
			Version: stackPackCurrentVersion,
		},
	}
	cli.MockClient.ApiMocks.StackpackApi.UpgradeStackPackResponse.Result = successfulResponseResult
	return &cli, cmd
}

func TestStackpackUpgradePrintToTable(t *testing.T) {
	cli, cmd := setupStackPackUpgradeCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "upgrade", "--name", "zabbix",
		"--unlocked-strategy", strategyFlag,
	)

	assert.Equal(t,
		[]stackstate_api.UpgradeStackPackCall{{
			PstackPackName: "zabbix",
			Punlocked:      &strategyFlag,
		}},
		*cli.MockClient.ApiMocks.StackpackApi.UpgradeStackPackCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, []string{"Successfully triggered upgrade from 1.0.0 to 2.0.0"}, *cli.MockPrinter.SuccessCalls)
}

func TestStackpackUpgradePrintToJson(t *testing.T) {
	cli, cmd := setupStackPackUpgradeCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "upgrade", "--name", "zabbix",
		"--unlocked-strategy", strategyFlag, "-o", "json",
	)

	assert.Equal(t,
		[]stackstate_api.UpgradeStackPackCall{{
			PstackPackName: "zabbix",
			Punlocked:      &strategyFlag,
		}},
		*cli.MockClient.ApiMocks.StackpackApi.UpgradeStackPackCalls)

	expectedJsonCalls := []map[string]interface{}{{
		"success":         true,
		"current-version": stackPackCurrentVersion,
		"next-version":    stackPackNextVersion,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}

func TestStackpackUpgradeHasWaitFlags(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUpgradeCommand(&cli.Deps)

	// Test that the wait and timeout flags exist
	waitFlag := cmd.Flags().Lookup("wait")
	assert.NotNil(t, waitFlag, "wait flag should exist")
	assert.Equal(t, "false", waitFlag.DefValue, "wait flag default should be false")

	timeoutFlag := cmd.Flags().Lookup("timeout")
	assert.NotNil(t, timeoutFlag, "timeout flag should exist")
	assert.Equal(t, "1m0s", timeoutFlag.DefValue, "timeout flag default should be 1m0s")
}

func TestStackpackUpgradeWithWaitError(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUpgradeCommand(&cli.Deps)

	validateWaitFlags(t, cmd)
}

func TestStackpackUpgradeWithWaitTimeout(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUpgradeCommand(&cli.Deps)

	configID := int64(12345)
	timestamp := int64(1438167001716)

	// Setup initial stackpack for validation and provisioning state for wait
	provisioningStackPack := stackstate_api.FullStackPack{
		Name: stackPackName,
		NextVersion: &stackstate_api.FullStackPack{
			Version: stackPackNextVersion,
		},
		Version: stackPackCurrentVersion,
		Configurations: []stackstate_api.StackPackConfiguration{
			{
				Id:                  &configID,
				Status:              StatusProvisioning,
				StackPackVersion:    stackPackNextVersion,
				LastUpdateTimestamp: &timestamp,
				Config:              map[string]interface{}{},
			},
		},
	}

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{provisioningStackPack}
	cli.MockClient.ApiMocks.StackpackApi.UpgradeStackPackResponse.Result = successfulResponseResult

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "upgrade", "--name", stackPackName,
		"--unlocked-strategy", strategyFlag, "--wait", "--timeout", "50ms")

	// Should return timeout error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout waiting for stackpack")
}

func TestStackpackUpgradeWithWaitJsonFlags(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUpgradeCommand(&cli.Deps)

	// Test that wait and timeout flags can be parsed
	err := cmd.ParseFlags([]string{"--name", "test", "--wait", "--timeout", "100ms"})
	assert.NoError(t, err)

	// Verify flag values
	waitValue, err := cmd.Flags().GetBool("wait")
	assert.NoError(t, err)
	assert.True(t, waitValue)

	timeoutValue, err := cmd.Flags().GetDuration("timeout")
	assert.NoError(t, err)
	assert.Equal(t, 100*time.Millisecond, timeoutValue)

	nameValue, err := cmd.Flags().GetString("name")
	assert.NoError(t, err)
	assert.Equal(t, "test", nameValue)
}
