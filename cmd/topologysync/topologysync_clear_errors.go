package topologysync

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const (
	DefaultIdentifierType = "NodeId"
)

type ClearErrorsArgs struct {
	Id int64
}

func ClearErrorsCommand(deps *di.Deps) *cobra.Command {
	args := &ClearErrorsArgs{}
	cmd := &cobra.Command{
		Use:   "clear-errors",
		Short: "Clear the errors from an active topology synchronization",
		Long:  "Clear the errors from an active topology synchronization.",
		RunE:  deps.CmdRunEWithApi(RunClearErrorsCommand(args)),
	}
	cmd.Flags().Int64VarP(&args.Id, common.IDFlag, common.IDFlagShort, 0, "The ID of topology synchronization")
	cmd.MarkFlagRequired(common.IDFlag) //nolint:errcheck

	return cmd
}

func RunClearErrorsCommand(args *ClearErrorsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		resp, err := api.TopologySynchronizationApi.
			PostTopologySynchronizationStreamClearErrors(cli.Context).
			IdentifierType(DefaultIdentifierType).
			Identifier(fmt.Sprintf("%d", args.Id)).
			Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"cleared": args.Id,
			})
		} else {
			cli.Printer.Successf("Errors cleared from topology synchronization: %d", args.Id)
		}
		return nil
	}
}
