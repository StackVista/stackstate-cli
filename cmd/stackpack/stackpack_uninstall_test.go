package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupStackpackUninstallFn(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackUninstallCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackAPI.ProvisionUninstallResponse.Result = "Successfully uninstalled StackPack: name=zabbix id=39206337488300"
	return &cli, cmd
}

func TestStackpackUninstallCommandPrintToConsole(t *testing.T) {
	cli, cmd := setupStackpackUninstallFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "uninstall", "--name", "zabbix", "--id", "39206337488300")

	expectedSuccessMessage := []string{"Successfully uninstalled StackPack: name=zabbix id=39206337488300"}

	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, expectedSuccessMessage, *cli.MockPrinter.SuccessCalls)
}

func TestStackpackUninstallCommandPrintToJson(t *testing.T) {
	cli, cmd := setupStackpackUninstallFn(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "uninstall", "--name", "zabbix", "-i", "39206337488300", "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"success": true,
		"result":  "Successfully uninstalled StackPack: name=zabbix id=39206337488300",
	}}

	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
