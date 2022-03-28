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
		Use:   "update -f FILE -i ID",
		Short: "update a monitor",
		RunE:  cli.CmdRunEWithApi(RunUpdateMonitorCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", FileFlagUsage)
	cmd.Flags().StringP(IdFlag, "i", "", IdFlagUsage)
	cmd.MarkFlagRequired(IdFlag)
	cmd.MarkFlagRequired(FileFlag)

	return cmd
}

func RunUpdateMonitorCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo di.ServerInfo) common.CLIError {
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
		_, resp, err = api.MonitorApi.UpdateMonitor(cli.Context, id).UpdateMonitor(monitor).Execute()
	} else {
		_, resp, err = api.MonitorUrnApi.UpdateMonitorByURN(cli.Context, identifier).UpdateMonitor(monitor).Execute()
	}
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("Monitor updated: %s", identifier))
	return nil
}
