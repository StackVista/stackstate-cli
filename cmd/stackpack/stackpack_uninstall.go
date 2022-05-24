package stackpack

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

type UninstallArgs struct {
	Name string
	Id   int64
}

func StackpackUninstallCommand(cli *di.Deps) *cobra.Command {
	args := &UninstallArgs{}
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall instances of a StackPack",
		Long:  "Uninstall StackPack instances by id.",
		RunE:  cli.CmdRunEWithApi(RunStackpackUninstallCommand(args)),
	}
	cmd.Flags().Int64VarP(&args.Id, IdFlag, common.IDFlagShort, 0, "id of the StackPack instance")
	common.AddRequiredNameFlagVar(cmd, &args.Name, "name of the StackPack")
	cmd.MarkFlagRequired(IdFlag) //nolint:errcheck
	return cmd
}

func RunStackpackUninstallCommand(args *UninstallArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		result, resp, err := api.StackpackApi.ProvisionUninstall(cli.Context, args.Name, args.Id).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"success": true,
				"result":  result,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Successfully uninstalled StackPack: name=%s id=%d", args.Name, args.Id))
		}

		return nil
	}
}
