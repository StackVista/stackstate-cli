package context

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupValidateCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	setupConfig(t, &cli)
	cmd := ValidateCommand(&cli.Deps)
	return &cli, cmd
}

func TestValidationFailure(t *testing.T) {
	cli, cmd := setupValidateCmd(t)
	connectError := common.NewConnectError(fmt.Errorf("authentication error"), "https://test", &http.Response{StatusCode: 401})
	cli.MockClient.ConnectError = connectError
	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd)
	assert.Equal(t, connectError, err)
}

func TestValidationToJson(t *testing.T) {
	cli, cmd := setupValidateCmd(t)

	di.ExecuteCommandWithContextUnsafe(
		&cli.Deps,
		cmd,
		"--url",
		"https://test.stackstate.io/",
		"--api-token",
		"blaat",
		"-o", "json",
	)
	assert.Equal(t,
		[]map[string]interface{}{{
			"connected":   true,
			"context":     "foo",
			"server-info": cli.MockClient.ConnectServerInfo,
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
