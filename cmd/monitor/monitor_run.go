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

func MonitorRunCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run a monitor",
		Long:  "Run a monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorRunCommand),
	}

	common.AddRequiredIDFlag(cmd, IDFlagUsage)
	cmd.Flags().Bool(DryRunFlag, false, "do not save the states of the monitor run")

	return cmd
}

func RunMonitorRunCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	isDryRun, err := cmd.Flags().GetBool(DryRunFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	identifier, err := cmd.Flags().GetString(common.IDFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	id, err := util.StringToInt64(identifier)
	var resp *http.Response
	var runResult *stackstate_api.MonitorRunResult
	if err == nil {
		if isDryRun {
			runResult, resp, err = api.MonitorApi.DryRunMonitor(cli.Context, id).Execute()
		} else {
			runResult, resp, err = api.MonitorApi.RunMonitor(cli.Context, id).Execute()
		}
	} else {
		if isDryRun {
			runResult, resp, err = api.MonitorUrnApi.DryRunMonitorByURN(cli.Context, identifier).Execute()
		} else {
			runResult, resp, err = api.MonitorUrnApi.RunMonitorByURN(cli.Context, identifier).Execute()
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
