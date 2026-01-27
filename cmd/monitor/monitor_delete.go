package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteArgs struct {
	ID         int64
	Identifier string
}

func MonitorDeleteCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete {--id ID | --identifier URN}",
		Short: "Delete a monitor permanently",
		Long:  "Delete a monitor by its ID or identifier. This removes the monitor and all its associated health states.",
		Example: `# delete a monitor by ID
sts monitor delete --id 123456789

# delete a monitor by identifier
sts monitor delete --identifier urn:stackpack:my-monitor`,
		RunE: cli.CmdRunEWithApi(RunDeleteMonitorCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunDeleteMonitorCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		identifier := IdOrIdentifier(args.ID, args.Identifier)
		resp, err := api.MonitorApi.DeleteMonitor(cli.Context, identifier).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-monitor-identifier": identifier,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Monitor deleted: %s", identifier))
		}
		return nil
	}
}
