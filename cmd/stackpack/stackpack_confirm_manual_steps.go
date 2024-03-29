package stackpack

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type ManualStepsArgs struct {
	TypeName string
	ID       int64
}

func StackpackConfirmManualStepsCommand(cli *di.Deps) *cobra.Command {
	args := &ManualStepsArgs{}
	cmd := &cobra.Command{
		Use:   "confirm-manual-steps",
		Short: "Confirm manual steps of StackPack",
		Long: "During some installation of StackPacks manual actions/steps must be taken by the user in order for the installation to complete. " +
			"These steps can not be verified by the system and only require a confirmation from the user. This command sends such a confirmation.",
		RunE: cli.CmdRunEWithApi(RunStackpackConfirmManualStepsCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.TypeName, "Name of the StackPack")
	common.AddRequiredIDFlagVar(cmd, &args.ID, "ID of the StackPack instance")
	return cmd
}

func RunStackpackConfirmManualStepsCommand(args *ManualStepsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		_, resp, err := api.StackpackApi.ConfirmManualSteps(cli.Context, args.TypeName, args.ID).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"success":               true,
				"stackpack-name":        args.TypeName,
				"stackpack-instance-id": args.ID,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Confirmed manual step for StackPack '%s' instance id '%d'", args.TypeName, args.ID))
		}

		return nil
	}
}
