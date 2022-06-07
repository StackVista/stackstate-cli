package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

var (
	stackPackName           = "zabbix"
	stackPackNextVersion    = "2.0.0"
	stackPackCurrentVersion = "1.0.0"
)

func setupStackPackUpgradeCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUpgradeCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = []stackstate_api.Sstackpack{
		{
			Name: &stackPackName,
			NextVersion: &stackstate_api.SstackpackLatestVersion{
				Version: &stackPackNextVersion,
			},
			Version: &stackPackCurrentVersion,
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
			PstackName: "zabbix",
			Punlocked:  &strategyFlag,
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
			PstackName: "zabbix",
			Punlocked:  &strategyFlag,
		}},
		*cli.MockClient.ApiMocks.StackpackApi.UpgradeStackPackCalls)

	expectedJsonCalls := []map[string]interface{}{{
		"success":         true,
		"current-version": stackPackCurrentVersion,
		"next-version":    stackPackNextVersion,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
