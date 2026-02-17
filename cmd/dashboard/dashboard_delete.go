package dashboard

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

// DeleteArgs contains arguments for dashboard delete command
type DeleteArgs struct {
	ID         int64
	Identifier string
}

func DashboardDeleteCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a dashboard permanently",
		Long:  "Delete a dashboard by its ID or identifier. Only user-owned dashboards can be deleted.",
		Example: `# delete a dashboard by ID
sts dashboard delete --id 123456789

# delete a dashboard by identifier
sts dashboard delete --identifier urn:stackpack:my-dashboard`,
		RunE: cli.CmdRunEWithApi(RunDashboardDeleteCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, "ID of the dashboard")
	common.AddIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the dashboard")
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunDashboardDeleteCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		dashboardIdOrUrn, err := ResolveDashboardIdOrUrn(args.ID, args.Identifier)
		if err != nil {
			return common.NewCLIArgParseError(err)
		}

		// Execute the delete request
		resp, err := api.DashboardsApi.DeleteDashboard(cli.Context, dashboardIdOrUrn).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-dashboard-identifier": dashboardIdOrUrn,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Dashboard deleted: %s", dashboardIdOrUrn))
		}
		return nil
	}
}
