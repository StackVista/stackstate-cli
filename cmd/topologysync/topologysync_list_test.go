package topologysync

import (
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	Topo1 = stackstate_api.NewTopologyStreamListItem(int64(23), "Topo 1", 0, 0, 0, 0, 0)
	Topo2 = stackstate_api.NewTopologyStreamListItem(5, "Topo 2", 23, 5, 23, 5, 0)

	Topo3 = stackstate_api.NewTopologyStreamListItem(13, "Topo 3", 12345, 23, 23, 12345, 1)

	Topo4 = stackstate_api.NewTopologyStreamListItem(7, "Topo 4", 12345, 12345, 12345, 12345, 23)

	AllTopos = []stackstate_api.TopologyStreamListItem{*Topo1, *Topo2, *Topo3, *Topo4}

	SomeTopos = stackstate_api.NewTopologyStreamList(AllTopos)
)

func TestTopologySyncListTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.TopologySynchronizationApi.GetTopologySynchronizationStreamsResponse.Result = *SomeTopos

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Id", "Name", "Identifier", "Components", "Relations", "Errors"},
			Data: [][]interface{}{
				{Topo1.NodeId, Topo1.Name, "-", "+0       -0", "+0       -0", int64(0)},
				{Topo2.NodeId, Topo2.Name, "-", "+23      -5", "+23      -5", int64(0)},
				{Topo3.NodeId, Topo3.Name, "-", "+23   -12345", "+12345   -23", int64(1)},
				{Topo4.NodeId, Topo4.Name, "-", "+12345 -12345", "+12345 -12345", int64(23)},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "synchronizations"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestTopologySyncListJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := ListCommand(&cli.Deps)
	cli.MockClient.ApiMocks.TopologySynchronizationApi.GetTopologySynchronizationStreamsResponse.Result = *SomeTopos

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	expected := []map[string]interface{}{
		{
			"synchronizations": AllTopos,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}
