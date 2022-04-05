package settings

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func SettingsExportCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export {--ids IDs | --namespace NAMESPACE | --type LAYER_TYPE}",
		Short: "export settings with STJ",
		RunE:  cli.CmdRunEWithApi(RunSettingsExportCommand),
	}
	cmd.Flags().StringSlice(Ids, nil, "ids of the items")
	cmd.Flags().StringSlice(Namespace, nil, "namespace of the component")
	cmd.Flags().StringSlice(TypeName, nil, "type of the setting")
	cmd.Flags().StringSlice(AllowReferences, nil, "filter by reference")

	return cmd
}

func RunSettingsExportCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo stackstate_client.ServerInfo) common.CLIError {
	ids, err := cmd.Flags().GetStringSlice(Ids)
	if err != nil {
		return common.NewCLIError(err)
	}
	namespace, err := cmd.Flags().GetStringSlice(Namespace)
	if err != nil {
		return common.NewCLIError(err)
	}
	references, err := cmd.Flags().GetStringSlice(AllowReferences)
	if err != nil {
		return common.NewCLIError(err)
	}
	nodeTypes, err := cmd.Flags().GetStringSlice(TypeName)
	if err != nil {
		return common.NewCLIError(err)
	}
	if len(ids) != 0 && len(namespace) != 0 && len(nodeTypes) != 0 {
		return common.NewCLIArgParseError(fmt.Errorf("can not find \"%s\" or \"%s\" or \"%s\" flags. Pick one or the other", ids, Namespace, TypeName))
	}

	exportArgs := stackstate_client.NewExport()
	if len(ids) != 0 {
		exportArgs.NodesWithIds = &ids
	}
	if len(namespace) != 0 {
		exportArgs.Namespace = &namespace
	}
	if len(nodeTypes) != 0 {
		exportArgs.AllNodesOfTypes = &nodeTypes
	}
	if len(references) != 0 {
		if len(namespace) == 0 {
			return common.NewCLIError(fmt.Errorf("\"%s\" flag is required for \"%s\" flag", Namespace, AllowReferences))
		}
		exportArgs.AllowReferences = &references
	}

	data, resp, err := api.ExportApi.ExportSettings(cli.Context).Export(*exportArgs).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	cli.Printer.PrintLn(data)
	return nil
}
