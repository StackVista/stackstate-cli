package script

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommand() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := ScriptExecuteCommand(&mockCli.Deps)

	return mockCli, cmd
}

func TestExecuteSuccess(t *testing.T) {
	cli, cmd := setupCommand()
	cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteResponse.Result = sts.ExecuteScriptResponse{
		Result: map[string]interface{}{"value": "hello test"},
	}

	util.ExecuteCommandWithContext(cli.Context, cmd, "--script", "test script")

	assert.Equal(t,
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test script"}},
		},
		*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls,
	)
	assert.Equal(t, []interface{}{"hello test"}, *cli.MockPrinter.PrintStructCalls)
}

func TestExecuteFromScript(t *testing.T) {
	cli, cmd := setupCommand()

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
		*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls,
	)
}

func TestExecuteResponseError(t *testing.T) {
	fakeError := fmt.Errorf("bla")
	cli, cmd := setupCommand()
	cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteResponse.Error = fakeError

	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--script", "test script")

	assert.Equal(t,
		[]sts.ScriptExecuteCall{
			{PexecuteScriptRequest: &sts.ExecuteScriptRequest{Script: "test script"}},
		},
		*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls,
	)
	assert.IsType(t, common.ResponseError{}, err)
}

func TestArgumentScriptFlag(t *testing.T) {
	cli, cmd := setupCommand()
	util.ExecuteCommandWithContext(cli.Context, cmd, "--arguments-script", "argscript", "--script", "test script")

	assert.Equal(t,
		"argscript",
		*(*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls)[0].PexecuteScriptRequest.ArgumentsScript,
	)
}

func TestTimeoutFlag(t *testing.T) {
	cli, cmd := setupCommand()
	util.ExecuteCommandWithContext(cli.Context, cmd, "-t", "10", "--script", "test script")

	assert.Equal(t,
		int32(10),
		*(*cli.MockClient.ApiMocks.ScriptingApi.ScriptExecuteCalls)[0].PexecuteScriptRequest.TimeoutMs,
	)
}

func TestScriptAndFileFlag(t *testing.T) {
	cli, cmd := setupCommand()
	_, err := util.ExecuteCommandWithContext(cli.Context, cmd, "--script", "script", "-f", "file")

	assert.Equal(t, common.NewCLIArgParseError(fmt.Errorf("can not load script both from the \"script\" and the \"file\" flags. Pick one or the other")), err)
}
