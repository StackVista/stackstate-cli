package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type CloneArgs struct {
	IdArgs
	NewName string
}

func MonitorCloneCommand(cli *di.Deps) *cobra.Command {
	args := &CloneArgs{}
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone a monitor",
		Long:  "Clones a monitor.",
		RunE:  cli.CmdRunEWithApi(RunMonitorCloneCommand(args)),
	}

	common.AddIDFlagVar(cmd, &args.ID, IDFlagUsage)
	common.AddIdentifierFlagVar(cmd, &args.Identifier, IdentifierFlagUsage)
	common.AddRequiredNameFlagVar(cmd, &args.NewName, NameFlagUsage)
	stscobra.MarkMutexFlags(cmd, []string{common.IDFlag, common.IdentifierFlag}, "identifier", true)

	return cmd
}

func RunMonitorCloneCommand(args *CloneArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *sts.APIClient,
		serverInfo *sts.ServerInfo,
	) common.CLIError {
		identifier := IdOrIdentifier(args.ID, args.Identifier)
		clonedMonitor, resp, err := api.NodeApi.Clone(cli.Context, "Monitor", identifier).NodeName(sts.NodeName{Name: args.NewName}).Execute()

		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"monitor": clonedMonitor,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Monitor %d has been created", clonedMonitor.GetId()))
		}

		return nil
	}
}
