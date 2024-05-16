package rbac

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func ListPermissionsCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-permissions",
		Short: "List available permissions",
		Long:  "List available permissions.",
		RunE:  deps.CmdRunEWithApi(RunListPermissionsCommand),
	}

	return cmd
}

func RunListPermissionsCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	permissions, resp, err := api.PermissionsAPI.GetPermissions(cli.Context).Execute()

	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"permissions": permissions.Permissions,
		})
	} else {
		cli.Printer.Successf("Available permissions: %s", strings.Join(permissions.Permissions, ", "))
	}

	return nil
}
