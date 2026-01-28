package otelcomponentmapping

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type DeleteArgs struct {
	Identifier string
}

func OtelComponentMappingDeleteCommand(deps *di.Deps) *cobra.Command {
	args := &DeleteArgs{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an OTel Component Mapping by identifier (URN)",
		Long:  "Delete an OTel Component Mapping by identifier (URN)",
		RunE:  deps.CmdRunEWithApi(RunDeleteComponentMappingCommand(args)),
	}

	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the Component Mapping to delete")

	return cmd
}

func RunDeleteComponentMappingCommand(args *DeleteArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		if args.Identifier == "" {
			return common.NewCLIArgParseError(fmt.Errorf("--identifier is required"))
		}

		resp, err := api.OtelMappingApi.DeleteOtelComponentMapping(cli.Context, args.Identifier).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		result := map[string]interface{}{"deleted_identifier": args.Identifier}
		if cli.IsJson() {
			cli.Printer.PrintJson(result)
		} else {
			cli.Printer.Success(fmt.Sprintf("OTel Component Mapping deleted: %s", args.Identifier))
		}

		return nil
	}
}
