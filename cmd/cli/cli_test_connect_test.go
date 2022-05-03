package cli

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupTestConnectCmd() (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps()
	cmd := CliTestCommandCommand(&cli.Deps)
	return &cli, cmd
}

func TestConnectionFailure(t *testing.T) {
	cli, cmd := setupTestConnectCmd()
	connectError := common.NewConnectError(fmt.Errorf("authentication error"), &http.Response{StatusCode: 401})
	cli.MockClient.ConnectError = connectError
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)
	assert.Equal(t, connectError, err)
}

func TestConnectionToJson(t *testing.T) {
	cli, cmd := setupTestConnectCmd()

	di.ExecuteCommandWithContextUnsafe(
		&cli.Deps,
		cmd,
		"--url",
		"https://test.stackstate.io/",
		"--api-token",
		"blaat",
		"--json",
	)
	assert.Equal(t,
		[]map[string]interface{}{{
			"connected":   true,
			"server-info": cli.MockClient.ConnectServerInfo,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
