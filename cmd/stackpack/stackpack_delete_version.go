package stackpack

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteVersionArgs struct {
	Name    string
	Version string
}

func StackpackDeleteVersionCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteVersionArgs{}
	cmd := &cobra.Command{
		Use:   "delete-version",
		Short: "Delete a specific version of a StackPack",
		Long:  "Delete a specific version of a StackPack from the server. The version cannot be deleted if it is currently in use by an installed instance.",
		Example: `# delete version 1.0.0 of the kubernetes StackPack
sts stackpack delete-version --name kubernetes --version 1.0.0`,
		RunE: cli.CmdRunEWithApi(RunStackpackDeleteVersionCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the StackPack")
	cmd.Flags().StringVar(&args.Version, VersionFlag, "", "Version to delete")
	cmd.MarkFlagRequired(VersionFlag) //nolint:errcheck
	return cmd
}

func RunStackpackDeleteVersionCommand(args *DeleteVersionArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		resp, err := api.StackpackApi.StackPackDeleteVersion(cli.Context, args.Name, args.Version).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted": args.Version,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Successfully deleted version %s of StackPack %s", args.Version, args.Name))
		}

		return nil
	}
}
