package script

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/cobra_util"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func setupCommand(mockScriptingApiService sts.MockScriptingApiService) (*printer.MockPrinter, di.Deps, *cobra.Command) {
	client := sts.APIClient{}
	client.ScriptingApi = mockScriptingApiService
	mockPrinter := printer.NewMockPrinter()
	cli := di.Deps{
		Config:  &conf.Conf{},
		Client:  &client,
		Printer: &mockPrinter,
		Context: context.Background(),
	}
	cmd := ScriptExecuteCommand(&cli)

	return &mockPrinter, cli, cmd
}

func TestExecuteSuccess(t *testing.T) {
	mockApi := sts.NewMockScriptingApiService()
	mockPrinter, cli, cmd := setupCommand(mockApi)

	cobra_util.ExecuteCommandWithContext(cli.Context, cmd, "test script")

	assert.Equal(t,
		&[]sts.ExecuteScriptRequest{{Script: "test script"}},
		mockApi.ExecuteScriptRequests,
	)
	assert.Equal(t, []interface{}{"hello test"}, *mockPrinter.PrintStructCalls)
	assert.Equal(t, []msg.LoadingMsg{msg.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}

func TestExecuteError(t *testing.T) {
	fakeError := fmt.Errorf("bla")
	mockApi := sts.NewMockScriptingApiService()
	mockApi.ReturnFromScriptExecuteExecute.Error = fakeError
	mockPrinter, cli, cmd := setupCommand(mockApi)

	cobra_util.ExecuteCommandWithContext(cli.Context, cmd, "test script")

	assert.Equal(t,
		&[]sts.ExecuteScriptRequest{{Script: "test script"}},
		mockApi.ExecuteScriptRequests,
	)
	assert.Equal(t, []error{fakeError}, *mockPrinter.PrintErrCalls)
	assert.Equal(t, []msg.LoadingMsg{msg.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}

func TestArgumentScriptFlag(t *testing.T) {
	mockApi := sts.NewMockScriptingApiService()
	_, cli, cmd := setupCommand(mockApi)
	cobra_util.ExecuteCommandWithContext(cli.Context, cmd, "-a", "argscript", "test script")

	assert.Equal(t, "argscript", *(*mockApi.ExecuteScriptRequests)[0].ArgumentsScript)
}

func TestTimeoutFlag(t *testing.T) {
	mockApi := sts.NewMockScriptingApiService()
	_, cli, cmd := setupCommand(mockApi)
	cobra_util.ExecuteCommandWithContext(cli.Context, cmd, "-t", "10", "test script")

	assert.Equal(t, int32(10), *(*mockApi.ExecuteScriptRequests)[0].TimeoutMs)
}
