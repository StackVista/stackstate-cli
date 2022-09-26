package graph

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func DeleteExpiredDataCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-expired-data",
		Short: "Delete expired data from StackState.",
		Long:  "Delete expired data from StackState.",
		RunE:  deps.CmdRunEWithApi(runDeleteExpiredDataCommand),
	}

	return cmd
}

func runDeleteExpiredDataCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {

	progress := "RemovalInProgress" // TODO

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"progress": progress,
		})
	} else {
		cli.Printer.Success("Command executed successfully.")
	}

	return nil
}
