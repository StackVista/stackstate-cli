package graph

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	Immediate = "immediate"
)

type DeleteExpiredDataArgs struct {
	Immediate bool
}

func DeleteExpiredDataCommand(deps *di.Deps) *cobra.Command {
	args := &DeleteExpiredDataArgs{}
	cmd := &cobra.Command{
		Use:   "delete-expired-data",
		Short: "Delete expired data from StackState.",
		Long:  "Schedule deletion of expired data from StackState. If --immediate is specified, removes data immediately and restarts StackState.",
		RunE:  deps.CmdRunEWithApi(RunDeleteExpiredDataCommand(args)),
	}

	cmd.Flags().BoolVar(&args.Immediate, Immediate, false, "Remove expired data immediately and restart StackState.")

	return cmd
}

func RunDeleteExpiredDataCommand(args *DeleteExpiredDataArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {

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
}
