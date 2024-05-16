package settings

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupEditCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := SettingsEditCommand(&cli.Deps)
	return &cli, cmd
}

func TestShouldEditSettings(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader("Hello"))}
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Result = "Hello"
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "1234")

	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.UnlockCalls, 0)
	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportAPI.ImportSettingsCalls)[0].Pbody, "olleH")
}

func TestShouldEditWithUnlock(t *testing.T) {
	cli, cmd := setupEditCmd(t)
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader("Hello"))}
	cli.MockClient.ApiMocks.ExportAPI.ExportSettingsResponse.Result = "Hello"
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "1234,2345,3456", "--unlock")

	assert.Len(t, *cli.MockClient.ApiMocks.NodeAPI.UnlockCalls, 3)
	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportAPI.ImportSettingsCalls)[0].Pbody, "olleH")
}
