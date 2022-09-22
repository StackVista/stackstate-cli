package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	stscobra "gitlab.com/stackvista/stackstate-cli2/internal/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type EnableDisableArgs struct {
	ID         int64
	Identifier string
}

func MonitorEnableDisableCommand(cli *di.Deps, use string, description string, status *sts.MonitorStatusValue, message string) *cobra.Command {
	args := &EnableDisableArgs{}
	cmd := &cobra.Command{
		Use:   use,
		Short: description,
		Long:  description,
		RunE:  cli.CmdRunEWithApi(RunMonitorEnableDisableCommand(args, status, message)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunMonitorEnableDisableCommand(args *EnableDisableArgs, status *sts.MonitorStatusValue, message string) di.CmdWithApiFn {
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
