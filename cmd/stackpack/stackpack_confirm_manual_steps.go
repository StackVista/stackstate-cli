package stackpack

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func StackpackConfirmManualStepsCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-manual-steps",
		Short: "confirm manual steps of StackPack",
		Long: "During some installation of StackPacks manual actions/step must be taken by the user in order for the installation to complete. " +
			"These steps can not be verified by the system and only require a confirmation from the user. This command sends such a confirmation.",
		RunE: cli.CmdRunEWithApi(RunStackpackConfirmManualStepsCommand),
	}
	common.AddRequiredNameFlag(cmd, "name of the StackPack")
	cmd.Flags().Int64(common.IDFlag, 0, "id of the StackPack instance")
	cmd.MarkFlagRequired(common.IDFlag) //nolint:errcheck
	return cmd
}

func RunStackpackConfirmManualStepsCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	name, err := cmd.Flags().GetString(common.NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	id, err := cmd.Flags().GetInt64(common.IDFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	_, resp, err := api.StackpackApi.ConfirmManualSteps(cli.Context, name, id).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"success":               true,
			"stackpack-name":        name,
			"stackpack-instance-id": id,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("confirmation of manual steps of the provisioning StackPack Name: %s StackPack Id: %d", name, id))
	}

	return nil
}
