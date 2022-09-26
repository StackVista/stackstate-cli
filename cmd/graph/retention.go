package graph

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

const (
	Set = "set"
	SetShort = "s"
	ScheduleRemoval = "schedule-removal"
)

type RetentionArgs struct {
	Set time.Duration
	ScheduleRemoval bool
}

func RetentionCommand(deps *di.Deps) *cobra.Command {
	args := &RetentionArgs{}
	cmd := &cobra.Command{
		Use:   "retention",
		Short: "Manage the StackState Graph data retention.",
		Long:  "Fetch and set the StackState Graph data retention.",
		RunE:  deps.CmdRunEWithApi(RunRetentionCommand(args)),
	}

	pflags.DurationVarP(cmd.Flags(), &args.Set, Set, SetShort, pflags.Week, "New data retention window.")
	cmd.Flags().BoolVar(&args.ScheduleRemoval, ScheduleRemoval, false, "Schedule removal of expired data.")

	return cmd
}

func RunRetentionCommand(args *RetentionArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {

		// TODO
		window := 691200000
		epoch := 1663493504140000000

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"window": window,
				"epoch": epoch,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Retention window: %d milliseconds", window))
			cli.Printer.Success(fmt.Sprintf("Epoch transactionId: %d", epoch))
		}

		return nil
	}
}
