package settings

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func SettingsListTypesCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-types",
		Short: "list all types of settings available",
		RunE:  di.CmdRunEWithDeps(cli, RunSettingsListTypesCommand),
	}

	return cmd
}

func RunSettingsListTypesCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {

	nodeTypes, resp, err := cli.Client.NodeApi.NodeSettings(cli.Context).Execute()
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
