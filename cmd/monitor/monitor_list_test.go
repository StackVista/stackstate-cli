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
	firstMonitorTags        = []string{"app", "namespace:prod"}
	secondMonitorName       = "zMonitor"
	secondMonitorId         = int64(2)
	secondMonitorIdentifier = "urn:monitor:zMonitor"
	secondMonitorTags       []string
	monitorArray            = []sts.Monitor{
		{Id: secondMonitorId,
			Name:            secondMonitorName,
			Identifier:      &secondMonitorIdentifier,
			Description:     nil,
			FunctionId:      237346485409361,
			Arguments:       nil,
			RemediationHint: nil,
			IntervalSeconds: 2,
			Tags:            secondMonitorTags,
			Status:          sts.MONITORSTATUSVALUE_ENABLED},
		{Id: firstMonitorId,
			Name:            firstMonitorName,
			Identifier:      &firstMonitorIdentifier,
			Description:     nil,
			FunctionId:      237346485409361,
			Arguments:       nil,
			RemediationHint: nil,
			IntervalSeconds: 2,
			Tags:            firstMonitorTags,
			Status:          sts.MONITORSTATUSVALUE_ENABLED},
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
			IntervalSeconds: 2,
			Tags:            firstMonitorTags,
			Status:          sts.MONITORSTATUSVALUE_ENABLED},
		{Id: secondMonitorId,
			Name:            secondMonitorName,
			Identifier:      &secondMonitorIdentifier,
			Description:     nil,
			FunctionId:      237346485409361,
			Arguments:       nil,
			RemediationHint: nil,
			IntervalSeconds: 2,
			Tags:            secondMonitorTags,
			Status:          sts.MONITORSTATUSVALUE_ENABLED},
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
			Header: []string{"Id", "Identifier", "Name", "Tags", "Status"},
			Data: [][]interface{}{
				{firstMonitorId, firstMonitorIdentifier, firstMonitorName, firstMonitorTags, sts.MONITORSTATUSVALUE_ENABLED},
				{secondMonitorId, secondMonitorIdentifier, secondMonitorName, secondMonitorTags, sts.MONITORSTATUSVALUE_ENABLED}},
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
