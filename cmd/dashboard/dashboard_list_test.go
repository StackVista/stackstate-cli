package dashboard

import (
	"testing"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	firstDashboardName         = "aDashboard"
	firstDashboardId           = int64(1)
	firstDashboardIdentifier   = "urn:custom:dashboard:aDashboard"
	firstDashboardDescription  = "First dashboard description"
	firstDashboardScope        = sts.DASHBOARDSCOPE_PUBLIC_DASHBOARD
	secondDashboardName        = "zDashboard"
	secondDashboardId          = int64(2)
	secondDashboardIdentifier  = "urn:custom:dashboard:zDashboard"
	secondDashboardDescription = "Second dashboard description"
	secondDashboardScope       = sts.DASHBOARDSCOPE_PRIVATE_DASHBOARD

	dashboardArray = []sts.DashboardReadSchema{
		{
			DashboardReadMetadataSchema: &sts.DashboardReadMetadataSchema{
				Id:          secondDashboardId,
				Name:        secondDashboardName,
				Identifier:  secondDashboardIdentifier,
				Description: secondDashboardDescription,
				Scope:       secondDashboardScope,
			},
		},
		{
			DashboardReadMetadataSchema: &sts.DashboardReadMetadataSchema{
				Id:          firstDashboardId,
				Name:        firstDashboardName,
				Identifier:  firstDashboardIdentifier,
				Description: firstDashboardDescription,
				Scope:       firstDashboardScope,
			},
		},
	}
	dashboardList = sts.DashboardList{Dashboards: dashboardArray}

	orderedArray = []sts.DashboardReadSchema{
		{
			DashboardReadMetadataSchema: &sts.DashboardReadMetadataSchema{
				Id:          firstDashboardId,
				Name:        firstDashboardName,
				Identifier:  firstDashboardIdentifier,
				Description: firstDashboardDescription,
				Scope:       firstDashboardScope,
			},
		},
		{
			DashboardReadMetadataSchema: &sts.DashboardReadMetadataSchema{
				Id:          secondDashboardId,
				Name:        secondDashboardName,
				Identifier:  secondDashboardIdentifier,
				Description: secondDashboardDescription,
				Scope:       secondDashboardScope,
			},
		},
	}
)

func setDashboardListCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	cli := di.NewMockDeps(t)
	cmd := DashboardListCommand(&cli.Deps)
	return &cli, cmd
}

func TestDashboardListPrintToTable(t *testing.T) {
	cli, cmd := setDashboardListCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardsResponse.Result = dashboardList

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")
	expectedTableCall := []printer.TableData{
		{
			Header: []string{"Id", "Identifier", "Name", "Description", "Scope"},
			Data: [][]interface{}{
				{firstDashboardId, firstDashboardIdentifier, firstDashboardName, firstDashboardDescription, firstDashboardScope},
				{secondDashboardId, secondDashboardIdentifier, secondDashboardName, secondDashboardDescription, secondDashboardScope}},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "dashboards"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestDashboardListPrintToJson(t *testing.T) {
	cli, cmd := setDashboardListCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardsResponse.Result = dashboardList

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list", "--output", "json")
	expectedJsonCall := []map[string]interface{}{
		{
			"dashboards": orderedArray,
		},
	}
	assert.Equal(t, expectedJsonCall, *cli.MockPrinter.PrintJsonCalls)
	assert.False(t, cli.MockPrinter.HasNonJsonCalls)
}

func TestDashboardListEmptyResults(t *testing.T) {
	cli, cmd := setDashboardListCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardsResponse.Result = sts.DashboardList{Dashboards: []sts.DashboardReadSchema{}}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")
	expectedTableCall := []printer.TableData{
		{
			Header:              []string{"Id", "Identifier", "Name", "Description", "Scope"},
			Data:                [][]interface{}{},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "dashboards"},
		},
	}
	assert.Equal(t, expectedTableCall, *cli.MockPrinter.TableCalls)
}

func TestDashboardListSorting(t *testing.T) {
	cli, cmd := setDashboardListCmd(t)
	cli.MockClient.ApiMocks.DashboardsApi.GetDashboardsResponse.Result = dashboardList

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "list")

	// Verify that dashboards are sorted by identifier (firstDashboardIdentifier comes before secondDashboardIdentifier)
	tableData := (*cli.MockPrinter.TableCalls)[0].Data
	assert.Equal(t, firstDashboardId, tableData[0][0])
	assert.Equal(t, firstDashboardIdentifier, tableData[0][1])
	assert.Equal(t, secondDashboardId, tableData[1][0])
	assert.Equal(t, secondDashboardIdentifier, tableData[1][1])
}
