package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func MonitorDisableCommand(cli *di.Deps) *cobra.Command {
	args := &IdArgs{}
	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disable a monitor",
		Long:  "Disable a monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorDisableCommand(args, sts.MONITORSTATUSVALUE_DISABLED.Ptr())),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunMonitorDisableCommand(args *IdArgs, status *sts.MonitorStatusValue) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *sts.APIClient,
		serverInfo *sts.ServerInfo,
	) common.CLIError {
		identifier := IdOrIdentifier(args.ID, args.Identifier)

		monitorPatchResult, resp, err := RunMonitorStatusPatch(cli, api, status, identifier)

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
