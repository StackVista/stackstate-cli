package script

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/config"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	msg "gitlab.com/stackvista/stackstate-cli2/internal/messages"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func TestExecute(t *testing.T) {
	client := sts.APIClient{}
	client.ScriptingApi = MockScriptingApiService{}
	printer := printer.NewMockPrinter()
	ctx := di.Deps{
		Config:  &config.Config{},
		Client:  &client,
		Printer: printer,
	}

	cmd := ScriptExecuteCommand(&ctx)
	context := context.Background()
	RunScriptExecuteCommand(&ctx, &context, cmd, []string{""})

	expected := []interface{}{"hello test"}
	assert.Equal(t, expected, *printer.PrintStructCalls)
	assert.Equal(t, []msg.LoadingMsg{msg.AwaitingServer}, *printer.StartSpinnerCalls)
	assert.Equal(t, 1, *printer.StopSpinnerCalls)
}
