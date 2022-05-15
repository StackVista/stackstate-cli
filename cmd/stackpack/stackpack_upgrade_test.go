package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

var (
	stackPackName        = "zabbix"
	stackPackNextVersion = "1.0.0"
)

func setupStackPackUpgradeCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := StackpackUpgradeCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.StackpackListResponse.Result = []stackstate_api.Sstackpack{
		{
			Name: &stackPackName,
			NextVersion: &stackstate_api.SstackpackLatestVersion{
				Version: &stackPackNextVersion,
			},
		},
	}
	cli.MockClient.ApiMocks.StackpackApi.UpgradeStackPackResponse.Result = "successful"
	return &cli, cmd
}

func TestStackpackUpgradePrintToTable(t *testing.T) {
	cli, cmd := setupStackPackUpgradeCmd()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "upgrade", "--name", "zabbix",
		"--unlocked-strategy", "overwrite",
	)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, []string{"Successfully triggered upgrade of StackPack zabbix"}, *cli.MockPrinter.SuccessCalls)
}

func TestStackpackUpgradePrintToJson(t *testing.T) {
	cli, cmd := setupStackPackUpgradeCmd()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "upgrade", "--name", "zabbix",
		"--unlocked-strategy", "overwrite", "--json",
	)

	expectedJsonCalls := []map[string]interface{}{{
		"upgrade": "Successfully triggered upgrade of StackPack zabbix",
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
