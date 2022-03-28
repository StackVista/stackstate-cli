package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func CreateMonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create -f FILE",
		Short: "create a monitor",
		RunE:  cli.CmdRunEWithApi(RunCreateMonitorCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", FileFlagUsage)
	cmd.MarkFlagRequired(FileFlag)

	return cmd
}

func RunCreateMonitorCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo stackstate_client.ServerInfo) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	var monitor stackstate_client.CreateMonitor
	err = util.ReadInputStructFromFile(file, &monitor)
	if err != nil {
		return common.NewCLIError(err)
	}

	createdMonitor, resp, err := api.MonitorApi.CreateMonitor(cli.Context).CreateMonitor(monitor).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("Monitor created: %s (%d)", util.ToString(createdMonitor.Identifier), createdMonitor.Id))
	return nil
}
