package servicetoken

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

type DeleteArgs struct {
	ID string
}

func DeleteCommand(deps *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a service token",
		Long:  "Delete a service token.",
		RunE:  deps.CmdRunEWithApi(RunServiceTokenDeleteCommand(args)),
	}

	common.AddRequiredIDFlagVar(cmd, &args.ID, "id")

	return cmd
}

func RunServiceTokenDeleteCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		id, err := util.StringToInt64(args.ID)
		if err != nil {
			return common.NewCLIArgParseError(fmt.Errorf("invalid id: %s", args.ID))
		}

		resp, err := api.ServiceTokenApi.DeleteServiceToken(cli.Context, id).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted-service-token": args.ID,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Service token deleted: %s", args.ID))
		}

		return nil
	}
}
