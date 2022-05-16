package servicetoken

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func DeleteCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a service token",
		Long:  "Delete a service token.",
		RunE:  deps.CmdRunEWithApi(RunServiceTokenDeleteCommand),
	}

	common.AddRequiredIDFlag(cmd, "id")

	return cmd
}

func RunServiceTokenDeleteCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	sid, err := cmd.Flags().GetString("id")
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	id, err := util.StringToInt64(sid)
	if err != nil {
		return common.NewCLIArgParseError(fmt.Errorf("invalid id: %s", sid))
	}

	resp, err := api.ServiceTokenApi.DeleteServiceToken(cli.Context, id).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"deleted-service-token": sid,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Service token deleted: %s", sid))
	}

	return nil
}
