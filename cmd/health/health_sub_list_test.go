package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func setupHealthSubListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthSubListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationSubStreamOverviewResponse.Result = stackstate_api.SubStreamList{
		SubStreams: []stackstate_api.SubStreamListItem{{
			SubStreamId:     "StackState Server",
			CheckStateCount: 36,
		}},
	}
	return &cli, cmd
}

func TestHealthSubListPrintToTable(t *testing.T) {
	cli, cmd := setupHealthSubListCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "sub-list", "--urn", "dummy:urn")

	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Sub stream id", "Check state count"},
			Data:                [][]interface{}{{"StackState Server", int32(36)}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health sub-stream"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthSubListPrintToJson(t *testing.T) {
	cli, cmd := setupHealthSubListCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "sub-list", "--urn", "dummy:urn", "-o", "json")
	assert.Equal(t,
		[]map[string]interface{}{{
			"sub-stream": [][]interface{}{{"StackState Server", int32(36)}},
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
