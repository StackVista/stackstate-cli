package monitor

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
)

type DeleteArgs struct {
	ID         int64
	Identifier string
}

func MonitorDeleteCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a monitor",
		Long:  "Delete a monitor.",
		RunE:  cli.CmdRunEWithApi(RunDeleteMonitorCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	mutex_flags.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunDeleteMonitorCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		var resp *http.Response
		var err error

		if args.ID != 0 {
			resp, err = api.MonitorApi.DeleteMonitor(cli.Context, args.ID).Execute()
		} else {
			resp, err = api.MonitorUrnApi.DeleteMonitorByURN(cli.Context, args.Identifier).Execute()
		}
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		identifier := args.Identifier
		if identifier == "" {
			identifier = fmt.Sprintf("%d", args.ID)
		}
		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-monitor-identifier": identifier,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Monitor deleted: %s", identifier))
		}
		return nil
	}
}
