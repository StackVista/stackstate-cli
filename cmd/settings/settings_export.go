package settings

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func SettingsExportCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export settings with STJ",
		RunE:  cli.CmdRunEWithApi(RunSettingsExportCommand),
	}

	return cmd
}

func RunSettingsExportCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo stackstate_client.ServerInfo) common.CLIError {
	exportArgs := stackstate_client.NewExport()

	settingsList, resp, err := api.ExportApi.ExportSettings(cli.Context).Export(*exportArgs).Execute()
	if err != nil {
		return common.NewCLIError(err)
	}
	var respData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return common.NewCLIError(err)
	}

	cli.Printer.PrintStruct(settingsList)
	return nil
}
