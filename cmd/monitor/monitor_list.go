package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ListMonitorsCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all monitors.",
		RunE:  di.CmdRunEWithDeps(cli, RunListMonitorsCommand),
	}
	return cmd
}

func RunListMonitorsCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	client, ctx := cli.NewStsClient()

	monitors, resp, err := client.MonitorApi.GetAllMonitors(ctx).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.PrintStruct(monitors)
	return nil
}
