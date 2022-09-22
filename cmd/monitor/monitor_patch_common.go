package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func RunMonitorEnableDisableCommand(args *IdArgs, status *sts.MonitorStatusValue, message string) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *sts.APIClient,
		serverInfo *sts.ServerInfo,
	) common.CLIError {
		identifier := IdOrIdentifier(args.ID, args.Identifier)

		monitorPatch := sts.MonitorPatch{
			Status: status,
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
			cli.Printer.Success(fmt.Sprintf("Monitor %s has been %s", identifier, message))
		}

		return nil
	}
}
