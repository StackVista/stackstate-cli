package dashboard

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

// CloneArgs contains arguments for dashboard clone command
type CloneArgs struct {
	ID          int64
	Identifier  string
	NewName     string
	Description string
	Scope       string
}

func DashboardCloneCommand(cli *di.Deps) *cobra.Command {
	args := &CloneArgs{}
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone an existing dashboard to create a new copy",
		Long:  "Clone an existing dashboard to create a new copy with a different name. Optionally set a new description and scope for the cloned dashboard.",
		Example: `# clone a dashboard by ID
sts dashboard clone --id 123456789 --name My Dashboard Copy

# clone a dashboard as private
sts dashboard clone --identifier urn:stackpack:my-dashboard --name "Private Copy" --scope privateDashboard`,
		RunE: cli.CmdRunEWithApi(RunDashboardCloneCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, "ID of the dashboard")
	common.AddIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the dashboard")
	common.AddRequiredNameFlagVar(cmd, &args.NewName, "Name for the new dashboard")
	cmd.Flags().StringVar(&args.Description, "description", "", "Description for the new dashboard")
	cmd.Flags().StringVar(&args.Scope, "scope", "", "Scope for the new dashboard (publicDashboard or privateDashboard)")
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunDashboardCloneCommand(args *CloneArgs) di.CmdWithApiFn {
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

		// Create the clone schema with required name
		cloneSchema := stackstate_api.NewDashboardCloneSchema(args.NewName)

		// Add optional fields
		if args.Description != "" {
			cloneSchema.SetDescription(args.Description)
		}

		// Validate scope if provided - must be public or private
		if args.Scope != "" {
			switch args.Scope {
			case "publicDashboard":
				scope := stackstate_api.DASHBOARDSCOPE_PUBLIC_DASHBOARD
				cloneSchema.SetScope(scope)
			case "privateDashboard":
				scope := stackstate_api.DASHBOARDSCOPE_PRIVATE_DASHBOARD
				cloneSchema.SetScope(scope)
			default:
				return common.NewCLIArgParseError(fmt.Errorf("invalid scope: %s. Must be 'publicDashboard' or 'privateDashboard'", args.Scope))
			}
		}

		// Execute the clone request
		clonedDashboard, resp, err := api.DashboardsApi.CloneDashboard(cli.Context, dashboardIdOrUrn).DashboardCloneSchema(*cloneSchema).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"dashboard": clonedDashboard,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Dashboard cloned successfully! New dashboard ID: %d, Name: %s", clonedDashboard.GetId(), clonedDashboard.GetName()))
		}

		return nil
	}
}
