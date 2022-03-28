package settings

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func SettingsListTypesCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-types",
		Short: "list all types of settings available",
		RunE:  cli.CmdRunEWithApi(RunSettingsListTypesCommand),
	}

	return cmd
}

func RunSettingsListTypesCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo di.ServerInfo) common.CLIError {
	nodeTypes, resp, err := api.NodeApi.NodeListTypes(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	data := make([][]string, 0)
	for _, nodeType := range nodeTypes.NodeTypes {
		data = append(data, []string{nodeType.TypeName, nodeType.Description})
	}

	cli.Printer.Table(
		[]string{"Name", "Description"},
		data,
		nodeTypes,
	)

	return nil
}
