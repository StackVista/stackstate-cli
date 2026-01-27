package usersession

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func GetUserSessionRolesCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles",
		Short: "List available roles for the current session",
		Long:  "List all roles available to the current user session.",
		Example: `# list roles for current session
sts user-session roles`,
		RunE: cli.CmdRunEWithApi(RunUserSessionRolesCommand),
	}

	return cmd
}

func RunUserSessionRolesCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	roles, resp, err := api.UserSessionApi.GetUserSessionAvailableRoles(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"roles": roles.Roles,
		})
	} else {
		for _, role := range roles.Roles {
			cli.Printer.PrintLn(role)
		}
	}

	return nil
}
