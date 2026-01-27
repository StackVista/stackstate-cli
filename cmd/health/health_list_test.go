package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
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
	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverviewResponse.Result = stackstate_api.SubStreamList{
		SubStreams: []stackstate_api.SubStreamListItem{{
			SubStreamId:     "SUSE Observability Server",
			CheckStateCount: 36,
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
			"streams": []map[string]interface{}{{
				"stream_urn":               "urn:health:self_monitoring:self_monitoring_scraper",
				"stream_consistency_model": "REPEAT_SNAPSHOTS",
				"sub_stream_count":         int32(1),
			}},
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}

func TestHealthListUrnPrintToTable(t *testing.T) {
	cli, cmd := setupHealthListCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "--urn", "dummy:urn")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Sub stream id", "Check state count"},
			Data:                [][]interface{}{{"SUSE Observability Server", int32(36)}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health sub-stream"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthListUrnPrintToJson(t *testing.T) {
	cli, cmd := setupHealthListCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "--urn", "dummy:urn", "-o", "json")
	assert.Equal(t,
		[]map[string]interface{}{{
			"sub-stream": []map[string]interface{}{{
				"sub_stream_id":     "SUSE Observability Server",
				"check_state_count": int32(36),
			}},
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
