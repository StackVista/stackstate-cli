package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	downgradeVersion = "1.2.3"
)

func setupStackPackDowngradeCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDowngradeCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.DowngradeStackPackResponse.Result = successfulResponseResult
	return &cli, cmd
}

func TestStackpackDowngradePrintToTable(t *testing.T) {
	cli, cmd := setupStackPackDowngradeCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "downgrade", "--name", "zabbix",
		"--stackpack-version", downgradeVersion,
	)

	assert.Equal(t,
		[]stackstate_api.DowngradeStackPackCall{{
			PstackPackName: "zabbix",
			Pversion:       &downgradeVersion,
		}},
		*cli.MockClient.ApiMocks.StackpackApi.DowngradeStackPackCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t, []string{"Successfully triggered downgrade of zabbix to 1.2.3"}, *cli.MockPrinter.SuccessCalls)
}

func TestStackpackDowngradePrintToJson(t *testing.T) {
	cli, cmd := setupStackPackDowngradeCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "downgrade", "--name", "zabbix",
		"--stackpack-version", downgradeVersion, "-o", "json",
	)

	assert.Equal(t,
		[]stackstate_api.DowngradeStackPackCall{{
			PstackPackName: "zabbix",
			Pversion:       &downgradeVersion,
		}},
		*cli.MockClient.ApiMocks.StackpackApi.DowngradeStackPackCalls)

	expectedJsonCalls := []map[string]interface{}{{
		"success":        true,
		"target-version": downgradeVersion,
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}

func TestStackpackDowngradeRequiresVersion(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDowngradeCommand(&cli.Deps)

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "downgrade", "--name", "zabbix")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "stackpack-version")
}

func TestStackpackDowngradeHasWaitFlags(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDowngradeCommand(&cli.Deps)

	waitFlag := cmd.Flags().Lookup("wait")
	assert.NotNil(t, waitFlag, "wait flag should exist")
	assert.Equal(t, "false", waitFlag.DefValue, "wait flag default should be false")

	timeoutFlag := cmd.Flags().Lookup("timeout")
	assert.NotNil(t, timeoutFlag, "timeout flag should exist")
	assert.Equal(t, "1m0s", timeoutFlag.DefValue, "timeout flag default should be 1m0s")
}
