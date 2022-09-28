package graph

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_admin_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

const (
	Set             = "set"
	SetShort        = "s"
	ScheduleRemoval = "schedule-removal"
)

type RetentionArgs struct {
	Set             time.Duration
	ScheduleRemoval bool
}

func RetentionCommand(deps *di.Deps) *cobra.Command {
	args := &RetentionArgs{}
	cmd := &cobra.Command{
		Use:   "retention",
		Short: "Manage the StackState Graph data retention.",
		Long: "View and configure how long the StackState data graph retains data.\n" +
			"More info can ben found at http://docs.stackstate.com/setup/retention/.",
		RunE: deps.CmdRunEWithAdminApi(RunRetentionCommand(args)),
	}

	pflags.DurationVarP(cmd.Flags(), &args.Set, Set, SetShort, pflags.Week, "New data retention window")
	cmd.Flags().BoolVar(&args.ScheduleRemoval, ScheduleRemoval, false, "Schedule removal of expired data")

	return cmd
}

func RunRetentionCommand(args *RetentionArgs) di.CmdWithAdminApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
	) common.CLIError {
		window, resp, err := api.RetentionApi.GetRetentionWindow(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		epoch, resp, err := api.RetentionApi.GetRetentionEpoch(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"retention-window": window.WindowMs,
				"epoch":            epoch,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Retention window: %d milliseconds\nEpoch transactionId: %d", *window.WindowMs, *epoch.EpochTx))
		}

		return nil
	}
}
