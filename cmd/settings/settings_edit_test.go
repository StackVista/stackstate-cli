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
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader("Hello"))}
	cli.MockClient.ApiMocks.ExportApi.ExportSettingsResponse.Result = "Hello"
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--ids", "1234")

	assert.Equal(t, *(*cli.MockClient.ApiMocks.ImportApi.ImportSettingsCalls)[0].Pbody, "olleH")
}
