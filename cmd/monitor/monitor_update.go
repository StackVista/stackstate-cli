package monitor

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func UpdateMonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update a monitor.",
		RunE:  di.CmdRunEWithDeps(cli, RunUpdateMonitorCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", "The file with the monitor in it. Can either be a YAML or JSON file.")
	cmd.Flags().StringP(IdFlag, "i", "", "The id or identifier of the monitor you wish to update.")
	cmd.MarkFlagRequired(IdFlag)
	cmd.MarkFlagRequired(FileFlag)

	return cmd
}

func RunUpdateMonitorCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	var monitor stackstate_client.UpdateMonitor
	err = util.ReadInputStructFromFile(file, &monitor)
	if err != nil {
		return common.NewCLIError(err)
	}

	identifier, err := cmd.Flags().GetString(IdFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	id, err := strconv.ParseInt(identifier, 0, 64)
	var resp *http.Response
	if err == nil {
		_, resp, err = cli.Client.MonitorApi.UpdateMonitor(cli.Context, id).UpdateMonitor(monitor).Execute()
	} else {
		_, resp, err = cli.Client.MonitorUrnApi.UpdateMonitorByURN(cli.Context, identifier).UpdateMonitor(monitor).Execute()
	}
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("Monitor updated: %s", identifier))
	return nil
}
