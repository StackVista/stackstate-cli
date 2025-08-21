package dashboard

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func DashboardListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all dashboards",
		Long:  "List all dashboards.",
		RunE:  cli.CmdRunEWithApi(RunDashboardListCommand),
	}
	return cmd
}

func RunDashboardListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	dashboards, resp, err := api.DashboardsApi.GetDashboards(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	// Sort dashboards by identifier for consistent output
	sort.SliceStable(dashboards.Dashboards, func(i, j int) bool {
		identifierI := getDashboardIdentifier(dashboards.Dashboards[i])
		identifierJ := getDashboardIdentifier(dashboards.Dashboards[j])
		return identifierI < identifierJ
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"dashboards": dashboards.Dashboards,
		})
	} else {
		tableData := [][]interface{}{}
		for _, dashboard := range dashboards.Dashboards {
			id := getDashboardId(dashboard)
			identifier := getDashboardIdentifier(dashboard)
			name := getDashboardName(dashboard)
			description := getDashboardDescription(dashboard)
			scope := getDashboardScope(dashboard)

			tableData = append(tableData, []interface{}{id, identifier, name, description, scope})
		}
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Id", "Identifier", "Name", "Description", "Scope"},
			Data:                tableData,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "dashboards"},
		})
	}

	return nil
}

// Helper functions to extract data from the union type DashboardReadSchema
// These handle both DashboardReadMetadataSchema and DashboardReadFullSchema variants
func getDashboardId(dashboard stackstate_api.DashboardReadSchema) interface{} {
	if dashboard.DashboardReadMetadataSchema != nil {
		return dashboard.DashboardReadMetadataSchema.Id
	}
	if dashboard.DashboardReadFullSchema != nil {
		return dashboard.DashboardReadFullSchema.Id
	}
	return ""
}

func getDashboardIdentifier(dashboard stackstate_api.DashboardReadSchema) string {
	if dashboard.DashboardReadMetadataSchema != nil {
		return dashboard.DashboardReadMetadataSchema.Identifier
	}
	if dashboard.DashboardReadFullSchema != nil {
		return dashboard.DashboardReadFullSchema.Identifier
	}
	return ""
}

func getDashboardName(dashboard stackstate_api.DashboardReadSchema) string {
	if dashboard.DashboardReadMetadataSchema != nil {
		return dashboard.DashboardReadMetadataSchema.Name
	}
	if dashboard.DashboardReadFullSchema != nil {
		return dashboard.DashboardReadFullSchema.Name
	}
	return ""
}

func getDashboardDescription(dashboard stackstate_api.DashboardReadSchema) string {
	if dashboard.DashboardReadMetadataSchema != nil {
		return dashboard.DashboardReadMetadataSchema.Description
	}
	if dashboard.DashboardReadFullSchema != nil {
		return dashboard.DashboardReadFullSchema.Description
	}
	return ""
}

func getDashboardScope(dashboard stackstate_api.DashboardReadSchema) interface{} {
	if dashboard.DashboardReadMetadataSchema != nil {
		return dashboard.DashboardReadMetadataSchema.Scope
	}
	if dashboard.DashboardReadFullSchema != nil {
		return dashboard.DashboardReadFullSchema.Scope
	}
	return ""
}
