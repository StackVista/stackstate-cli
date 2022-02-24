package monitor

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	IdFlag = "id"
)

func DeleteMonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a monitor.",
		RunE:  di.CmdRunEWithDeps(cli, RunDeleteMonitorCommand),
	}
	cmd.Flags().Int64P(IdFlag, "i", 0, "The id of the monitor you wish to delete.")
	cmd.MarkFlagRequired(IdFlag)

	return cmd
}

func RunDeleteMonitorCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	monitorId, err := cmd.Flags().GetInt64(IdFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	client, ctx := cli.NewStsClient()
	resp, err := client.MonitorApi.DeleteMonitor(ctx, monitorId).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success("Monitor deleted.")
	return nil
}
