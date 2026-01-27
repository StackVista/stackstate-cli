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
		Use:   "confirm-manual-steps --name NAME --id ID",
		Short: "Confirm completion of manual installation steps",
		Long:  `Confirm that manual installation steps have been completed for a StackPack instance. Some StackPacks require manual actions that cannot be verified automatically.`,
		Example: `# confirm manual steps for a StackPack instance
sts stackpack confirm-manual-steps --name kubernetes --id 123456789`,
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
