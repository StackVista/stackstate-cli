package monitor

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

var (
	hello = "Hello"
)

func setMonitorEditCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEditCommand(&cli.Deps)
	return &cli, cmd
}

func TestShouldEditMonitor(t *testing.T) {
	cli, cmd := setMonitorEditCmd(t)
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader("Hello"))}
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Result = "Hello"
	cli.MockClient.ApiMocks.NodeAPI.LockResponse.Result = stackstate_api.NodeUnlockedAsLockedResponse(stackstate_api.NewNodeUnlocked("NodeUnlocked"))
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234")

	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.LockCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.UnlockCalls, 0)
	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportAPI.ImportSettingsCalls)[0].Pbody, "olleH")
}

func TestShouldOfferCloneMonitor(t *testing.T) {
	cli, cmd := setMonitorEditCmd(t)
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader(hello))}
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Result = hello
	cli.MockClient.ApiMocks.NodeAPI.LockResponse.Result = stackstate_api.NodeLockedAsLockedResponse(stackstate_api.NewNodeLocked("NodeLocked", "stackpack"))
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234")

	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.LockCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.UnlockCalls, 0)
	assert.Len(t, *cli.MockClient.ApiMocks.ImportAPI.ImportSettingsCalls, 0)
	assert.Equal(t, "The monitor  that you are trying is locked (StackState specific), it cannot be edited", (*cli.MockPrinter.PrintLnCalls)[0])
}

func TestEditWithUnlock(t *testing.T) {
	cli, cmd := setMonitorEditCmd(t)
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader("Hello"))}
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Result = hello
	cli.MockClient.ApiMocks.NodeAPI.LockResponse.Result = stackstate_api.NodeUnlockedAsLockedResponse(stackstate_api.NewNodeUnlocked("NodeUnlocked"))
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234", "--unlock")

	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.LockCalls, 1)
	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.UnlockCalls, 1)
	assert.Equal(t, "olleH", *(*cli.MockClient.ApiMocks.ImportAPI.ImportSettingsCalls)[0].Pbody)
}
