package script

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
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

	util.ExecuteCommandWithContext(cli.Context, cmd, "test script")

	assert.Equal(t,
		&[]sts.ExecuteScriptRequest{{Script: "test script"}},
		mockApi.ExecuteScriptRequests,
	)
	assert.Equal(t, []interface{}{"hello test"}, *mockPrinter.PrintStructCalls)
	assert.Equal(t, []common.LoadingMsg{common.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}

func TestExecuteFromScript(t *testing.T) {
	mockApi := sts.NewMockScriptingApiService()
	mockPrinter, cli, cmd := setupCommand(mockApi)

	tmpFile, err := ioutil.TempFile(os.TempDir(), "script-execute-test-")
	if err != nil {
		t.Fatal("Cannot create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())

	text := []byte("test content")
	if _, err = tmpFile.Write(text); err != nil {
		t.Fatal("Failed to write to temporary file", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	util.ExecuteCommandWithContext(cli.Context, cmd, "--file", tmpFile.Name())

	assert.Equal(t,
		&[]sts.ExecuteScriptRequest{{Script: "test content"}},
		mockApi.ExecuteScriptRequests,
	)
	assert.Equal(t, []interface{}{"hello test"}, *mockPrinter.PrintStructCalls)
	assert.Equal(t, []common.LoadingMsg{common.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}

func TestExecuteResponseError(t *testing.T) {
	fakeError := fmt.Errorf("bla")
	mockApi := sts.NewMockScriptingApiService()
	mockApi.ReturnFromScriptExecuteExecute.Error = fakeError
	mockPrinter, cli, cmd := setupCommand(mockApi)

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "test script")

	assert.Equal(t,
		&[]sts.ExecuteScriptRequest{{Script: "test script"}},
		mockApi.ExecuteScriptRequests,
	)
	assert.IsType(t, common.ResponseError{}, err)
	assert.Equal(t, []common.LoadingMsg{common.AwaitingServer}, *mockPrinter.StartSpinnerCalls)
	assert.Equal(t, 1, *mockPrinter.StopSpinnerCalls)
}

func TestArgumentScriptFlag(t *testing.T) {
	mockApi := sts.NewMockScriptingApiService()
	_, cli, cmd := setupCommand(mockApi)
	util.ExecuteCommandWithContext(cli.Context, cmd, "-a", "argscript", "test script")

	assert.Equal(t, "argscript", *(*mockApi.ExecuteScriptRequests)[0].ArgumentsScript)
}

func TestTimeoutFlag(t *testing.T) {
	mockApi := sts.NewMockScriptingApiService()
	_, cli, cmd := setupCommand(mockApi)
	util.ExecuteCommandWithContext(cli.Context, cmd, "-t", "10", "test script")

	assert.Equal(t, int32(10), *(*mockApi.ExecuteScriptRequests)[0].TimeoutMs)
}
