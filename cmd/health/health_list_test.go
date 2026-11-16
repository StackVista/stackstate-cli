package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

const (
	testRepeatSnapshots = "REPEAT_SNAPSHOTS"
	testServerName      = "SUSE Observability Server"
)

func setupHealthListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationStreamsOverviewResponse.Result = stackstate_api.StreamList{
		Items: []stackstate_api.StreamListItem{{
			Urn:              testHealthStreamURN,
			ConsistencyModel: testRepeatSnapshots,
			SubStreams:       1,
		}},
	}
	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverviewResponse.Result = stackstate_api.SubStreamList{
		SubStreams: []stackstate_api.SubStreamListItem{{
			SubStreamId:     testServerName,
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
			Data:                [][]interface{}{{testHealthStreamURN, testRepeatSnapshots, int32(1)}},
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
				"stream_urn":               testHealthStreamURN,
				"stream_consistency_model": testRepeatSnapshots,
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
			Data:                [][]interface{}{{testServerName, int32(36)}},
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
				"sub_stream_id":     testServerName,
				"check_state_count": int32(36),
			}},
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
