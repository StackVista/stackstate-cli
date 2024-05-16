package graph

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	stackstate_admin_api "github.com/stackvista/stackstate-cli/generated/stackstate_admin_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/pkg/pflags"
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
		Short: "Manage the StackState Graph data retention",
		Long: "View and configure how long the StackState data graph retains data.\n" +
			"More info can ben found at http://docs.stackstate.com/setup/retention/.",
		RunE: deps.CmdRunEWithAdminApi(RunRetentionCommand(args)),
	}

	pflags.DurationVarP(cmd.Flags(), &args.Set, Set, SetShort, (time.Duration)(0), "New data retention window")
	cmd.Flags().BoolVar(&args.ScheduleRemoval, ScheduleRemoval, false, "Schedule removal of expired data")

	return cmd
}

func getOrSetRetentionWindow(cli *di.Deps, api *stackstate_admin_api.APIClient, args *RetentionArgs) (*stackstate_admin_api.WindowMs, *http.Response, error) {
	if args.Set == 0 {
		return api.RetentionAPI.GetRetentionWindow(cli.Context).Execute()
	} else {
		newWindow := stackstate_admin_api.WindowMs{
			WindowMs: args.Set.Milliseconds(),
		}

		return api.RetentionAPI.SetRetentionWindow(cli.Context).
			WindowMs(newWindow).
			ScheduleRemoval(args.ScheduleRemoval).
			Execute()
	}
}

func RunRetentionCommand(args *RetentionArgs) di.CmdWithAdminApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_admin_api.APIClient,
	) common.CLIError {
		window, resp, err := getOrSetRetentionWindow(cli, api, args)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		epoch, resp, err := api.RetentionAPI.GetRetentionEpoch(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"retention-window": window.WindowMs,
				"epoch":            epoch.EpochTx,
			})
		} else {
			cli.Printer.Successf("Retention window: %d milliseconds\nEpoch transactionId: %d", window.WindowMs, epoch.EpochTx)
		}

		return nil
	}
}
