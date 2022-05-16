package stackpack

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func StackpackUninstallCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall instances of a StackPack",
		Long:  "Uninstall StackPack instances by id.",
		RunE:  cli.CmdRunEWithApi(RunStackpackUninstallCommand),
	}
	cmd.Flags().String(NameFlag, "", "name of the StackPack")
	cmd.Flags().Int64(IdFlag, 0, "id of the StackPack instance")
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
	cmd.MarkFlagRequired(IdFlag)   //nolint:errcheck
	return cmd
}

func RunStackpackUninstallCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	name, err := cmd.Flags().GetString(NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	id, err := cmd.Flags().GetInt64(IdFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	result, resp, err := api.StackpackApi.ProvisionUninstall(cli.Context, name, id).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"success": true,
			"result":  result,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Successfully uninstalled StackPack: name=%s id=%d", name, id))
	}

	return nil
}
