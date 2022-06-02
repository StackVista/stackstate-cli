package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
)

const (
	YesFlag      = "yes"
	YesFlagShort = "y"
)

type RunArgs struct {
	ID         int64
	Identifier string
	DoRun      bool
}

func MonitorRunCommand(cli *di.Deps) *cobra.Command {
	args := &RunArgs{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run a monitor",
		Long:  "Run a monitor. Does not save the states of the monitor run, unless `--yes` is specified.",
		RunE:  cli.CmdRunEWithApi(RunMonitorRunCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	cmd.Flags().BoolVarP(&args.DoRun, YesFlag, YesFlagShort, false, "save the state of the monitor run")
	mutex_flags.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "idenfier", true)

	return cmd
}

func RunMonitorRunCommand(args *RunArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		identifier := IdOrIdentifier(args.ID, args.Identifier)
		runResult, resp, err := api.MonitorApi.RunMonitor(cli.Context, identifier).DryRun(!args.DoRun).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"run-result": runResult.Result,
			})
		} else {
			cli.Printer.PrintStruct(runResult.Result)
		}

		return nil
	}
}
