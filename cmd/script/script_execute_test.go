package script

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func setupScriptExecuteCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := ScriptExecuteCommand(&cli.Deps)
	return &cli, cmd
}

func TestExecuteSuccess(t *testing.T) {
	cli, cmd := setupScriptExecuteCmd()
	cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteResponse.Result = sts.ExecuteScriptResponse{
		Result: map[string]interface{}{"value": "hello test"},
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--script", "test script")

	assert.Equal(t,
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test script"}},
		},
		*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls,
	)
	assert.Equal(t, []interface{}{"hello test"}, *cli.MockPrinter.PrintStructCalls)
}

func TestExecuteFromScript(t *testing.T) {
	cli, cmd := setupScriptExecuteCmd()

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

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--file", tmpFile.Name())

	assert.Equal(t,
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test content"}},
		},
		*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls,
	)
	assert.Equal(t, []string{"script executed (no response)"}, *cli.MockPrinter.SuccessCalls)
}

func TestExecuteResponseError(t *testing.T) {
	fakeError := fmt.Errorf("bla")
	fakeErrorResp := &http.Response{StatusCode: 403}
	cli, cmd := setupScriptExecuteCmd()
	cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteResponse.Error = fakeError
	cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteResponse.Response = fakeErrorResp

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--script", "test script")

	assert.Equal(t,
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test script"}},
		},
		*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls,
	)
	assert.Equal(t, common.NewResponseError(fakeError, fakeErrorResp), err)
}

func TestArgumentScriptFlag(t *testing.T) {
	cli, cmd := setupScriptExecuteCmd()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--arguments-script", "argscript", "--script", "test script")

	assert.Equal(t,
		"argscript",
		*(*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls)[0].PexecuteScriptRequest.ArgumentsScript,
	)
}

func TestTimeoutFlag(t *testing.T) {
	cli, cmd := setupScriptExecuteCmd()
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-t", "10", "--script", "test script")

	assert.Equal(t,
		int32(10),
		*(*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls)[0].PexecuteScriptRequest.TimeoutMs,
	)
}

func TestScriptAndFileFlag(t *testing.T) {
	cli, cmd := setupScriptExecuteCmd()
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--script", "script", "-f", "file")

	assert.Equal(t,
		common.NewCLIArgParseError(
			fmt.Errorf("can not load script both from the \"script\" and the \"file\" flags."+
				" Pick one or the other"),
		),
		err,
	)
}

func TestConnectionError(t *testing.T) {
	cli, cmd := setupScriptExecuteCmd()
	respError := common.NewConnectError(fmt.Errorf("authentication error"), &http.Response{StatusCode: 401})
	cli.MockClient.ConnectError = respError
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--script", "script")
	assert.Equal(t, respError, err)
}
