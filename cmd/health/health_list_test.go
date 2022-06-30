package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func setupHealthListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationStreamsOverviewResponse.Result = stackstate_api.StreamList{
		Items: []stackstate_api.StreamListItem{{
			Urn:              "urn:health:self_monitoring:self_monitoring_scraper",
			ConsistencyModel: "REPEAT_SNAPSHOTS",
			SubStreams:       1,
		}},
	}
	return &cli, cmd
}

func TestHealthListPrintToTable(t *testing.T) {
	cli, cmd := setupHealthListCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Stream urn", "Stream consistency model", "Sub stream count"},
			Data:                [][]interface{}{{"urn:health:self_monitoring:self_monitoring_scraper", "REPEAT_SNAPSHOTS", int32(1)}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health stream"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthListPrintToJson(t *testing.T) {
	cli, cmd := setupHealthListCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "-o", "json")
	assert.Equal(t,
		[]map[string]interface{}{{
			"streams": [][]interface{}{{"urn:health:self_monitoring:self_monitoring_scraper", "REPEAT_SNAPSHOTS", int32(1)}},
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
