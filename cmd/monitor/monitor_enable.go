package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func MonitorEnableCommand(cli *di.Deps) *cobra.Command {
	args := &IdArgs{}
	cmd := &cobra.Command{
		Use:   "enable {--id ID | --identifier URN}",
		Short: "Enable a disabled monitor to resume scheduled execution",
		Long:  "Enable a disabled monitor to resume running on its configured schedule. The monitor will start producing health states again.",
		Example: `# enable a monitor by ID
sts monitor enable --id 123456789

# enable a monitor by identifier
sts monitor enable --identifier urn:stackpack:my-monitor`,
		RunE: cli.CmdRunEWithApi(RunMonitorEnableCommand(args)),
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
