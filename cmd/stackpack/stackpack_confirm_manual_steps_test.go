package stackpack

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupStackPacConfirmManualStepsCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := StackpackConfirmManualStepsCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.ConfirmManualStepsResponse.Result = "successful"
	return &cli, cmd
}

func TestStackpackConfirmManualStepsPrintsToTable(t *testing.T) {
	cli, cmd := setupStackPacConfirmManualStepsCmd()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "confirm-manual-steps", "--name", "zabbix",
		"--id", "1234",
	)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t,
		[]string{"confirmation of manual steps of the provisioning StackPack Name: zabbix StackPack Id: 1234"},
		*cli.MockPrinter.SuccessCalls)
}

func TestStackpackConfirmManualStepsPrintsToJson(t *testing.T) {
	cli, cmd := setupStackPacConfirmManualStepsCmd()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "confirm-manual-steps", "--name", "zabbix",
		"--id", "1234", "--json",
	)
	expectedJsonCalls := []map[string]interface{}{{
		"success":               true,
		"stackpack-name":        "zabbix",
		"stackpack-instance-id": int64(1234),
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
}
