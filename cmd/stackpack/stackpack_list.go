package stackpack

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func StackpackListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list stackpacks",
		Long:  "List available StackPacks.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListCommand),
	}
	cmd.Flags().Bool(InstalledFlag, false, "show only installed StackPacks")
	return cmd
}

func RunStackpackListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	isInstalled, err := cmd.Flags().GetBool(InstalledFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	data := make([][]interface{}, 0)
	stackpacks := make([]stackstate_api.Sstackpack, 0)
	for _, stackpack := range stackpackList {
		if isInstalled && len(stackpack.GetConfigurations()) == 0 {
			continue // skip as this is not installed
		}
		row := []interface{}{
			stackpack.Name,
			stackpack.DisplayName,
			stackpack.Version,
			getVersion(stackpack.NextVersion),
			getVersion(stackpack.LatestVersion),
			len(stackpack.Configurations),
		}

		data = append(data, row)
		stackpacks = append(stackpacks, stackpack)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"stackpacks": stackpacks,
		})
	} else {
		cli.Printer.Table(printer.TableData{
			Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		})
	}

	return nil
}

func getVersion(data *stackstate_api.SstackpackLatestVersion) string {
	if data != nil && data.GetVersion() != "" {
		return data.GetVersion()
	}
	return "-"
}
