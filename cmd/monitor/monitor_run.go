package monitor

import (
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
)

const (
	DryRunFlag = "dry-run"
)

type RunArgs struct {
	ID         int64
	Identifier string
	DryRun     bool
}

func MonitorRunCommand(cli *di.Deps) *cobra.Command {
	args := &RunArgs{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run a monitor",
		Long:  "Run a monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorRunCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	cmd.Flags().BoolVar(&args.DryRun, DryRunFlag, false, "do not save the states of the monitor run")
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
		var resp *http.Response
		var runResult *stackstate_api.MonitorRunResult
		var err error

		switch {
		case args.ID != 0 && args.DryRun:
			runResult, resp, err = api.MonitorApi.DryRunMonitor(cli.Context, args.ID).Execute()
		case args.ID != 0:
			runResult, resp, err = api.MonitorApi.RunMonitor(cli.Context, args.ID).Execute()
		case args.DryRun:
			runResult, resp, err = api.MonitorUrnApi.DryRunMonitorByURN(cli.Context, args.Identifier).Execute()
		default:
			runResult, resp, err = api.MonitorUrnApi.RunMonitorByURN(cli.Context, args.Identifier).Execute()
		}

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"run-result": runResult.Result,
			})
		} else {
			cli.Printer.PrintStruct(runResult.Result)
		}

		return nil
	}
}
