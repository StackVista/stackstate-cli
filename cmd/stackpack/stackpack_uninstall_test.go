package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupStackpackUninstallFn() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := StackpackUninstallCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.ProvisionUninstallResponse.Result = "Successfully uninstalled StackPack: name=zabbix id=39206337488300"
	return &cli, cmd
}

func TestStackpackUninstallCommandPrintToConsole(t *testing.T) {
	cli, cmd := setupStackpackUninstallFn()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "uninstall", "--name", "zabbix", "--id", "39206337488300")

	expectedSuccessMessage := []string{"Successfully uninstalled StackPack: name=zabbix id=39206337488300"}

	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, expectedSuccessMessage, *cli.MockPrinter.SuccessCalls)
}

func TestStackpackUninstallCommandPrintToJson(t *testing.T) {
	cli, cmd := setupStackpackUninstallFn()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "uninstall", "--name", "zabbix", "-i", "39206337488300", "--json")

	expectedJsonCalls := []map[string]interface{}{{
		"success": true,
		"result":  "Successfully uninstalled StackPack: name=zabbix id=39206337488300",
	}}

	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
