package settings

import (
	"github.com/spf13/cobra"
	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func SettingsListTypesCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-types",
		Short: "list all types of settings available",
		Long:  "List all types of settings available.",
		RunE:  cli.CmdRunEWithApi(RunSettingsListTypesCommand),
	}

	return cmd
}

func RunSettingsListTypesCommand(cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	settingTypes, resp, err := api.NodeApi.NodeListTypes(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"setting-types": settingTypes,
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
