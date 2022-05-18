package monitor

import (
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

const (
	DryRunFlag = "dry-run"
)

type RunArgs struct {
	ID     string
	DryRun bool
}

func MonitorRunCommand(cli *di.Deps) *cobra.Command {
	args := &RunArgs{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run a monitor",
		Long:  "Run a monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorRunCommand(args)),
	}

	common.AddRequiredIDFlagVar(cmd, &args.ID, IDFlagUsage)
	cmd.Flags().BoolVar(&args.DryRun, DryRunFlag, false, "do not save the states of the monitor run")

	return cmd
}

func RunMonitorRunCommand(args *RunArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		id, err := util.StringToInt64(args.ID)
		var resp *http.Response
		var runResult *stackstate_api.MonitorRunResult
		if err == nil {
			if args.DryRun {
				runResult, resp, err = api.MonitorApi.DryRunMonitor(cli.Context, id).Execute()
			} else {
				runResult, resp, err = api.MonitorApi.RunMonitor(cli.Context, id).Execute()
			}
		} else {
			if args.DryRun {
				runResult, resp, err = api.MonitorUrnApi.DryRunMonitorByURN(cli.Context, args.ID).Execute()
			} else {
				runResult, resp, err = api.MonitorUrnApi.RunMonitorByURN(cli.Context, args.ID).Execute()
			}
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
