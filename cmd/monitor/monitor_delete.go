package monitor

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func DeleteMonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete -i ID",
		Short: "delete a monitor",
		RunE:  cli.CmdRunEWithApi(RunDeleteMonitorCommand),
	}
	cmd.Flags().StringP(IDFlag, "i", "", IDFlag)
	cmd.MarkFlagRequired(IDFlag) //nolint:errcheck

	return cmd
}

func RunDeleteMonitorCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo *stackstate_client.ServerInfo,
) common.CLIError {
	identifier, err := cmd.Flags().GetString(IDFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	id, err := util.StringToInt64(identifier)
	var resp *http.Response
	if err == nil {
		resp, err = api.MonitorApi.DeleteMonitor(cli.Context, id).Execute()
	} else {
		resp, err = api.MonitorUrnApi.DeleteMonitorByURN(cli.Context, identifier).Execute()
	}
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("Monitor deleted: %s", identifier))
	return nil
}
