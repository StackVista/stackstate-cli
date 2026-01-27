package settings

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func SettingsListTypesCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-types",
		Short: "List all available setting types",
		Long:  "List all available setting types. Use these type names with 'sts settings list --type TYPE' to list settings of that type.",
		Example: `# list all available types
sts settings list-types`,
		RunE: cli.CmdRunEWithApi(RunSettingsListTypesCommand),
	}

	return cmd
}

func RunSettingsListTypesCommand(cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	settingTypes, resp, err := api.NodeApi.NodeListTypes(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	sort.SliceStable(settingTypes.NodeTypes, func(i, j int) bool {
		return settingTypes.NodeTypes[i].TypeName < settingTypes.NodeTypes[j].TypeName
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"setting-types": *settingTypes,
		})
	} else {
		data := make([][]interface{}, 0)
		for _, nodeType := range settingTypes.NodeTypes {
			data = append(data, []interface{}{nodeType.TypeName, nodeType.Description})
		}

		cli.Printer.Table(printer.TableData{
			Header:              []string{"name", "description"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "setting types"},
		})
	}

	return nil
}
