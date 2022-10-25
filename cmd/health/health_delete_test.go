package health

import (
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupHealthDeleteCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthDeleteCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationApi.DeleteHealthSynchronizationStreamResponse.Response = &http.Response{}

	return &cli, cmd
}

func TestHealthDeletePrintToTable(t *testing.T) {
	cli, cmd := setupHealthDeleteCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete", "--urn", "urn:health:self_monitoring:self_monitoring_scraper")

	assert.Equal(t, []string{"Stream deleted: urn:health:self_monitoring:self_monitoring_scraper"}, *cli.MockPrinter.SuccessCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthDeletePrintToJson(t *testing.T) {
	cli, cmd := setupHealthDeleteCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "delete", "--urn", "urn:health:self_monitoring:self_monitoring_scraper", "-o", "json")
	assert.Equal(t,
		[]map[string]interface{}{{
			"deleted-stream-urn": "urn:health:self_monitoring:self_monitoring_scraper",
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
