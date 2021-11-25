package script

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/cobra_util"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

var client = sts.APIClient{}
var cli di.Deps
var mockPrinter *printer.MockPrinter

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	client.ScriptingApi = MockScriptingApiService{}
	printer := printer.NewMockPrinter()
	mockPrinter = &printer
	cli = di.Deps{
		Config:  &config.Config{},
		Client:  &client,
		Printer: mockPrinter,
		Context: context.Background(),
	}
}

func TestExecuteSuccess(t *testing.T) {
	root := ScriptExecuteCommand(&cli)
	cobra_util.ExecuteCommandWithContext(cli.Context, root, "test")

	expected := []interface{}{"hello test"}
	assert.Equal(t, expected, *mockPrinter.PrintStructCalls)
	assert.Equal(t, []msg.LoadingMsg{msg.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}
