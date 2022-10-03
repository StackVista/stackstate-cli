package monitor

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setMonitorEditCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorEditCommand(&cli.Deps)
	return &cli, cmd
}

func TestShouldEditMonitor(t *testing.T) {
	cli, cmd := setMonitorEditCmd(t)
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader("Hello"))}
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = "Hello"
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234")

	assert.Len(t, *cli.MockClient.ApiMocks.NodeApi.UnlockCalls, 0)
	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pbody, "olleH")
}

func TestEditWithUnlock(t *testing.T) {
	cli, cmd := setMonitorEditCmd(t)
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader("Hello"))}
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = "Hello"
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", "1234", "--unlock")

	assert.Len(t, *cli.MockClient.ApiMocks.NodeApi.UnlockCalls, 1)
	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pbody, "olleH")
}
