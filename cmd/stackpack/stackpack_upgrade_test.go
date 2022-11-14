package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	stackPackName           = "zabbix"
	stackPackNextVersion    = "2.0.0"
	stackPackCurrentVersion = "1.0.0"
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
	cli.MockClient.ApiMocks.StackpackApi.UpgradeStackPackResponse.Result = "successful"
	return &cli, cmd
}

func TestStackpackUpgradePrintToTable(t *testing.T) {
	strategyFlag := "overwrite"
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
	strategyFlag := "overwrite"
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
