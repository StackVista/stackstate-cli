package script

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/cobra_util"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func setupCommand(mockScriptingApiService MockScriptingApiService) (*printer.MockPrinter, di.Deps, *cobra.Command) {
	client := sts.APIClient{}

	client.ScriptingApi = mockScriptingApiService
	mockPrinter := printer.NewMockPrinter()
	cli := di.Deps{
		Config:  &config.Config{},
		Client:  &client,
		Printer: &mockPrinter,
		Context: context.Background(),
	}
	command := ScriptExecuteCommand(&cli)

	return &mockPrinter, cli, command
}

func TestExecuteSuccess(t *testing.T) {
	mockPrinter, cli, root := setupCommand(NewMockScriptingApiService())
	cobra_util.ExecuteCommandWithContext(cli.Context, root, "test")

	expected := []interface{}{"hello test"}
	assert.Equal(t, expected, *mockPrinter.PrintStructCalls)
	assert.Equal(t, []msg.LoadingMsg{msg.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}

func TestExecuteError(t *testing.T) {
	fakeError := fmt.Errorf("bla")
	mockScriptingApiService := NewMockScriptingApiService()
	mockScriptingApiService.ReturnFromScriptExecuteExecute.Error = fakeError

	mockPrinter, cli, root := setupCommand(mockScriptingApiService)

	cobra_util.ExecuteCommandWithContext(cli.Context, root, "test")

	assert.Equal(t, []error{fakeError}, *mockPrinter.PrintErrCalls)
	assert.Equal(t, []msg.LoadingMsg{msg.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}
