package stackpack

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupStackPacConfirmManualStepsCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := StackpackConfirmManualStepsCommand(&cli.Deps)
	cli.MockClient.ApiMocks.StackpackApi.ConfirmManualStepsResponse.Result = successfulResponseResult
	return &cli, cmd
}

func TestStackpackConfirmManualStepsPrintsToTable(t *testing.T) {
	cli, cmd := setupStackPacConfirmManualStepsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "confirm-manual-steps", "--name", "zabbix",
		"--id", "1234",
	)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
	assert.Equal(t,
		[]string{"Confirmed manual step for StackPack 'zabbix' instance id '1234'"},
		*cli.MockPrinter.SuccessCalls)
	assert.Equal(t,
		*cli.MockClient.ApiMocks.StackpackApi.ConfirmManualStepsCalls,
		[]stackstate_api.ConfirmManualStepsCall{{
			PstackPackName:       name,
			PstackPackInstanceId: int64(1234),
		}},
	)
}

func TestStackpackConfirmManualStepsPrintsToJson(t *testing.T) {
	cli, cmd := setupStackPacConfirmManualStepsCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "confirm-manual-steps", "--name", "zabbix",
		"--id", "1234", "-o", "json",
	)
	expectedJsonCalls := []map[string]interface{}{{
		"success":               true,
		"stackpack-name":        "zabbix",
		"stackpack-instance-id": int64(1234),
	}}
	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.Equal(t,
		*cli.MockClient.ApiMocks.StackpackApi.ConfirmManualStepsCalls,
		[]stackstate_api.ConfirmManualStepsCall{{
			PstackPackName:       name,
			PstackPackInstanceId: int64(1234),
		}},
	)
}
