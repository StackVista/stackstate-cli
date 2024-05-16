package health

import (
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupHealthClearErrorCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthClearErrorCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationAPI.PostHealthSynchronizationStreamClearErrorsResponse.Response = &http.Response{}

	return &cli, cmd
}

func TestHealthClearErrorPrintToTable(t *testing.T) {
	cli, cmd := setupHealthClearErrorCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "clear-error", "--urn", "urn:health:dummy_error")

	expectedSuccessCalls := []string{"Stream error clear: urn:health:dummy_error"}
	assert.Equal(t, expectedSuccessCalls, *cli.MockPrinter.SuccessCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthClearErrorPrintToJson(t *testing.T) {
	cli, cmd := setupHealthClearErrorCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "clear-error", "--urn", "urn:health:dummy_error", "-o", "json")
	assert.Equal(t,
		[]map[string]interface{}{{
			"stream-error-clear": "urn:health:dummy_error",
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
