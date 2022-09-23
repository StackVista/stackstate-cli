package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	stscobra "gitlab.com/stackvista/stackstate-cli2/internal/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func MonitorEnableCommand(cli *di.Deps) *cobra.Command {
	args := &IdArgs{}
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable a monitor",
		Long:  "Enables a monitor to run based on its `intervalSeconds` schedule.",
		RunE:  cli.CmdRunEWithApi(RunMonitorEnableCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunMonitorEnableCommand(args *IdArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *sts.APIClient,
		serverInfo *sts.ServerInfo,
	) common.CLIError {
		identifier := IdOrIdentifier(args.ID, args.Identifier)
		monitorPatchResult, resp, err := RunMonitorStatusPatch(cli, api, sts.MONITORSTATUSVALUE_ENABLED.Ptr(), identifier)

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"monitor": monitorPatchResult,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Monitor %s has been enabled", identifier))
		}

		return nil
	}
}
