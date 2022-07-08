package health

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func setupHealthStatusCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := HealthStatusCommand(&cli.Deps)
	cli.MockClient.ApiMocks.HealthSynchronizationApi.GetHealthSynchronizationStreamTopologyMatchesResponse.Result = stackstate_api.TopologyMatchResult{
		MultipleMatchesCheckStates: []stackstate_api.MultipleMatchesCheckState{{
			CheckStateId:              "1",
			TopologyElementIdentifier: "id",
			MatchCount:                12,
		}},
	}

	return &cli, cmd
}

func TestHealthStatusPrintToTable(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "--topology-matches", "dummy")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Check state id", "Topology element identifier", "Number of matched topology elements"},
			Data:                [][]interface{}{{"1", "id", int32(12)}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "health status"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
	assert.True(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestHealthStatusPrintToJson(t *testing.T) {
	cli, cmd := setupHealthStatusCmd(t)
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "status", "--urn", "test:urn", "--topology-matches", "dummy", "-o", "json")
	assert.Equal(t,
		[]map[string]interface{}{{
			"status": [][]interface{}{{"1", "id", int32(12)}},
		}},
		*cli.MockPrinter.PrintJsonCalls,
	)
}
