package topologysync

import (
	"fmt"
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	Error1 = stackstate_api.NewTopologyStreamError(stackstate_api.MESSAGELEVEL_WARN, "Oops")
	Error2 = stackstate_api.NewTopologyStreamError(stackstate_api.MESSAGELEVEL_ERROR, "Welp, better leave the premises quickly.")
	Error3 = (func() *stackstate_api.TopologyStreamError {
		t := stackstate_api.NewTopologyStreamError(stackstate_api.MESSAGELEVEL_INFO, "I don't think this one will happen at all.")
		t.SetExternalId("SomeId")
		return t
	})()

	SomeErrors = []stackstate_api.TopologyStreamError{
		*Error1,
		*Error2,
		*Error3,
	}

	SomeTopoDetails = stackstate_api.NewTopologyStreamListItemWithErrorDetails(*Topo1, SomeErrors)
)

func TestTopologyDescribeJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)
	cli.MockClient.ApiMocks.TopologySynchronizationApi.GetTopologySynchronizationStreamByIdResponse.Result = *SomeTopoDetails

	id := SomeTopoDetails.Item.NodeId

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", fmt.Sprintf("%d", id), "-o", "json")

	expected := []map[string]interface{}{
		{
			"synchronization": SomeTopoDetails,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestTopologySyncDescribeErrors(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeCommand(&cli.Deps)
	cli.MockClient.ApiMocks.TopologySynchronizationApi.GetTopologySynchronizationStreamByIdResponse.Result = *SomeTopoDetails

	id := SomeTopoDetails.Item.NodeId

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--id", fmt.Sprintf("%d", id))

	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Id", "Name", "Identifier", "Status", "Components", "Relations", "Errors"},
			Data: [][]interface{}{
				{Topo1.NodeId, Topo1.Name, "-", Topo1.Status, "+0       -0", "+0       -0", int64(0)},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "synchronizations"},
		},
		{
			Header: []string{"External ID", "Level", "Message"},
			Data: [][]interface{}{
				{"-", Error1.Level, Error1.Message},
				{"-", Error2.Level, Error2.Message},
				{*Error3.ExternalId, Error3.Level, Error3.Message},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "synchronization errors"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}
