package monitor

import (
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

const (
	DryRunFlag = "dry-run"
)

func RunMonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run -i ID",
		Short: "run a monitor",
		RunE:  di.CmdRunEWithDeps(cli, RunRunMonitorCommand),
	}
	cmd.Flags().StringP(IdFlag, "i", "", IdFlag)
	cmd.Flags().Bool(DryRunFlag, false, "do not save the states of the monitor run")
	cmd.MarkFlagRequired(IdFlag)

	return cmd
}

func RunRunMonitorCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	isDryRun, err := cmd.Flags().GetBool(DryRunFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	identifier, err := cmd.Flags().GetString(IdFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	client, _, err := cli.Client.Connect()
	if err != nil {
		return common.NewConnectError(err)
	}

	id, err := strconv.ParseInt(identifier, 0, 64)
	var resp *http.Response
	var runResult stackstate_client.MonitorRunResult
	if err == nil {
		if isDryRun {
			runResult, resp, err = client.MonitorApi.DryRunMonitor(cli.Context, id).Execute()
		} else {
			runResult, resp, err = client.MonitorApi.RunMonitor(cli.Context, id).Execute()
		}
	} else {
		if isDryRun {
			runResult, resp, err = client.MonitorUrnApi.DryRunMonitorByURN(cli.Context, identifier).Execute()
		} else {
			runResult, resp, err = client.MonitorUrnApi.RunMonitorByURN(cli.Context, identifier).Execute()
		}
	}
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.PrintStruct(runResult.Result)

	return nil
}
