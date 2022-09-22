package monitor

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	stscobra "gitlab.com/stackvista/stackstate-cli2/internal/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type DisableArgs struct {
	ID         int64
	Identifier string
}

func MonitorDisableCommand(cli *di.Deps) *cobra.Command {
	args := &EnableArgs{}
	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable a monitor",
		Long:  "Disable a monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorDisableCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunMonitorDisableCommand(args *EnableArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		identifier := IdOrIdentifier(args.ID, args.Identifier)

		monitorPatch := stackstate_api.MonitorPatch{
			Status: stackstate_api.MONITORSTATUSVALUE_DISABLED.Ptr(),
		}
		monitorPatchResult, resp, err := api.MonitorApi.PatchMonitor(cli.Context, identifier).MonitorPatch(monitorPatch).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"monitor": monitorPatchResult,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Monitor %s has been disabled", identifier))
		}

		return nil
	}
}
