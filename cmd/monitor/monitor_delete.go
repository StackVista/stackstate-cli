package monitor

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func DeleteMonitorCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete -i ID",
		Short: "delete a monitor",
		RunE:  di.CmdRunEWithDeps(cli, RunDeleteMonitorCommand),
	}
	cmd.Flags().StringP(IdFlag, "i", "", IdFlag)
	cmd.MarkFlagRequired(IdFlag)

	return cmd
}

func RunDeleteMonitorCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	identifier, err := cmd.Flags().GetString(IdFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	id, err := strconv.ParseInt(identifier, 0, 64)
	var resp *http.Response
	if err == nil {
		resp, err = cli.Client.MonitorApi.DeleteMonitor(cli.Context, id).Execute()
	} else {
		resp, err = cli.Client.MonitorUrnApi.DeleteMonitorByURN(cli.Context, identifier).Execute()
	}
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.Success(fmt.Sprintf("Monitor deleted: %s", identifier))
	return nil
}
