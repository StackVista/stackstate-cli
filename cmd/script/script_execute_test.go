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

func setupCommand(mockScriptingApiService sts.ScriptingApiMock) (*printer.MockPrinter, di.Deps, *cobra.Command) {
	client := sts.APIClient{}
	client.ScriptingApi = mockScriptingApiService
	mockPrinter := printer.NewMockPrinter()
	cli := di.Deps{
		Config:  &conf.Conf{},
		Printer: &mockPrinter,
		Context: context.Background(),
		Client:  &client,
	}
	cmd := ScriptExecuteCommand(&cli)

	return &mockPrinter, cli, cmd
}

func TestExecuteSuccess(t *testing.T) {
	mockApi := sts.NewScriptingApiMock()
	mockApi.ScriptExecuteResponse.Result = sts.ExecuteScriptResponse{
		Result: map[string]interface{}{"value": "hello test"},
	}
	mockPrinter, cli, cmd := setupCommand(mockApi)

	util.ExecuteCommandWithContext(cli.Context, cmd, "--script", "test script")

	assert.Equal(t,
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test script"}},
		},
		*mockApi.ScriptExecuteCalls,
	)
	assert.Equal(t, []interface{}{"hello test"}, *mockPrinter.PrintStructCalls)
}

func TestExecuteFromScript(t *testing.T) {
	mockApi := sts.NewScriptingApiMock()
	_, cli, cmd := setupCommand(mockApi)

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
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test content"}},
		},
		*mockApi.ScriptExecuteCalls,
	)
}

func TestExecuteResponseError(t *testing.T) {
	fakeError := fmt.Errorf("bla")
	mockApi := sts.NewScriptingApiMock()
	mockApi.ScriptExecuteResponse.Error = fakeError
	_, cli, cmd := setupCommand(mockApi)

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--script", "test script")

	assert.Equal(t,
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test script"}},
		},
		*mockApi.ScriptExecuteCalls,
	)
	assert.IsType(t, common.ResponseError{}, err)
}

func TestArgumentScriptFlag(t *testing.T) {
	mockApi := sts.NewScriptingApiMock()
	_, cli, cmd := setupCommand(mockApi)
	util.ExecuteCommandWithContext(cli.Context, cmd, "--arguments-script", "argscript", "--script", "test script")

	assert.Equal(t,
		"argscript",
		*(*mockApi.ScriptExecuteCalls)[0].PexecuteScriptRequest.ArgumentsScript,
	)
}

func TestTimeoutFlag(t *testing.T) {
	mockApi := sts.NewScriptingApiMock()
	_, cli, cmd := setupCommand(mockApi)
	util.ExecuteCommandWithContext(cli.Context, cmd, "-t", "10", "--script", "test script")

	assert.Equal(t,
		int32(10),
		*(*mockApi.ScriptExecuteCalls)[0].PexecuteScriptRequest.TimeoutMs,
	)
}

func TestScriptAndFileFlag(t *testing.T) {
	mockApi := sts.NewScriptingApiMock()
	_, cli, cmd := setupCommand(mockApi)
	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--script", "script", "-f", "file")

	assert.Equal(t, common.NewCLIArgParseError(fmt.Errorf("can not load script both from the \"script\" and the \"file\" flags. Pick one or the other")), err)
}
