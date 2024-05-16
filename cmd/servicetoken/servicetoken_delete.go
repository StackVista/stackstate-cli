package servicetoken

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteArgs struct {
	ID int64
}

func DeleteCommand(deps *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a service token",
		Long:  "Delete a service token.",
		RunE:  deps.CmdRunEWithApi(RunServiceTokenDeleteCommand(args)),
	}

	common.AddRequiredIDFlagVar(cmd, &args.ID, "ID of the service token to delete")

	return cmd
}

func RunServiceTokenDeleteCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		resp, err := api.ServiceTokenAPI.DeleteServiceToken(cli.Context, args.ID).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-service-token": args.ID,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Service token deleted: %d", args.ID))
		}

		return nil
	}
}
