package graph

import (
	"github.com/spf13/cobra"
	stackstate_admin_api "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_admin_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	Immediate = "immediate"
)

type DeleteExpiredDataArgs struct {
	Immediate bool
}

var (
	responses = map[string]string{
		"RemovalInProgress": "Removal in progress.",
		"RemovalSucceeded":  "Removal succeeded.",
		"RemovalFailed":     "Command succeeded but the removal failed. Please consult the logs for more information.",
	}
)

func DeleteExpiredDataCommand(deps *di.Deps) *cobra.Command {
	args := &DeleteExpiredDataArgs{}
	cmd := &cobra.Command{
		Use:   "delete-expired-data",
		Short: "Delete expired data from StackState",
		Long: "Schedule deletion of expired data from StackState.\n" +
			"If --immediate is specified, removes data immediately and restarts StackState.",
		RunE: deps.CmdRunEWithAdminApi(RunDeleteExpiredDataCommand(args)),
	}

	cmd.Flags().BoolVar(&args.Immediate, Immediate, false, "Remove expired data immediately and restart StackState")

	return cmd
}

func RunDeleteExpiredDataCommand(args *DeleteExpiredDataArgs) di.CmdWithAdminApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_admin_api.APIClient,
	) common.CLIError {
		progress, resp, err := api.RetentionApi.RemoveExpiredData(cli.Context).
			ExpireImmediatelyAndRestart(args.Immediate).
			Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"progress": *progress.Progress,
			})
		} else {
			pp := *progress.Progress
			response, ok := responses[pp.Type]

			if ok {
				cli.Printer.Success(response)
			} else {
				cli.Printer.Success("Command executed successfully.")
			}
		}

		return nil
	}
}
