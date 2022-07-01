package monitor

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	firstMonitorName        = "aMonitor"
	firstMonitorId          = int64(1)
	firstMonitorIdentifier  = "urn:monitor:aMonitor"
	secondMonitorName       = "zMonitor"
	secondMonitorId         = int64(2)
	secondMonitorIdentifier = "urn:monitor:zMonitor"
	monitorArray            = []sts.Monitor{
		{Id: secondMonitorId,
			Name:            secondMonitorName,
			Identifier:      &secondMonitorIdentifier,
			Description:     nil,
			FunctionId:      237346485409361,
			Arguments:       nil,
			RemediationHint: nil,
			IntervalSeconds: 2},
		{Id: firstMonitorId,
			Name:            firstMonitorName,
			Identifier:      &firstMonitorIdentifier,
			Description:     nil,
			FunctionId:      237346485409361,
			Arguments:       nil,
			RemediationHint: nil,
			IntervalSeconds: 2},
	}
	monitorList = sts.MonitorList{Monitors: monitorArray}

	orderedArray = []sts.Monitor{
		{Id: firstMonitorId,
			Name:            firstMonitorName,
			Identifier:      &firstMonitorIdentifier,
			Description:     nil,
			FunctionId:      237346485409361,
			Arguments:       nil,
			RemediationHint: nil,
			IntervalSeconds: 2},
		{Id: secondMonitorId,
			Name:            secondMonitorName,
			Identifier:      &secondMonitorIdentifier,
			Description:     nil,
			FunctionId:      237346485409361,
			Arguments:       nil,
			RemediationHint: nil,
			IntervalSeconds: 2},
	}
)

func setMonitorListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := MonitorListCommand(&cli.Deps)
	return &cli, cmd
}

func TestMonitorListPrintToTable(t *testing.T) {
	cli, cmd := setMonitorListCmd(t)
	cli.MockClient.ApiMocks.MonitorApi.GetAllMonitorsResponse.Result = monitorList

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")
	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Id", "Identifier", "Name"},
			Data: [][]interface{}{
				{firstMonitorId, firstMonitorIdentifier, firstMonitorName},
				{secondMonitorId, secondMonitorIdentifier, secondMonitorName}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "monitors"},
		},
	}

	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestMonitorListPrintToJson(t *testing.T) {
	cli, cmd := setMonitorListCmd(t)
	cli.MockClient.ApiMocks.MonitorApi.GetAllMonitorsResponse.Result = monitorList

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	expectedJsonCalls := []map[string]interface{}{{
		"monitors": orderedArray,
	}}

	assert.Equal(t, expectedJsonCalls, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}
