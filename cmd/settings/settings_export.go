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
		Use:   "export {--ids IDs | --namespace NAMESPACE | --type LAYER_TYPE} --allow-references ALLOW-REFERENCES",
		Short: "export settings with STJ",
		RunE:  cli.CmdRunEWithApi(RunSettingsExportCommand),
	}
	cmd.Flags().StringSlice(Ids, []string{""}, "ids of the items")
	cmd.Flags().StringSlice(Namespace, []string{""}, "namespace of the component")
	cmd.Flags().StringSlice(TypeName, []string{""}, "type of the layer")
	cmd.Flags().StringSlice(AllowReferences, []string{""}, "filter by reference")

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
	if len(ids) != 0 && len(namespace) != 0 {
		return common.NewCLIArgParseError(fmt.Errorf("can not find \"%s\" and \"%s\" flags. Pick one or the other", ids, Namespace))
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

	data, _, err := api.ExportApi.ExportSettings(cli.Context).Export(*exportArgs).Execute()
	if err != nil {
		return common.NewCLIError(err)
	}

	//cli.Printer.PrintLn(`{"nodes": [{ "description": "description-1", "id": -214, "description": "description-1", "name": "name-1", "ownedBy": "urn:stackpack:common", "parameters": [{ "name": "name-param", "type": "LONG"}], "script": { "scriptBody": "script-bdy-1"}}]`)
	cli.Printer.PrintLn(data)
	return nil
}
