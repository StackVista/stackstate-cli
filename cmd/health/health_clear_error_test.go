package health

import (
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupHealthClearErrorCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthClearErrorCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationApi.PostHealthSynchronizationStreamClearErrorsResponse.Response = &http.Response{}

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
