package stackpack

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func StackpackListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List stackpacks",
		Long:  "List available StackPacks.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListCommand),
	}
	cmd.Flags().Bool(InstalledFlag, false, "Show only installed StackPacks")
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
	stackPackList, resp, err := api.StackpackApi.StackPackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	sort.SliceStable(stackPackList, func(i, j int) bool {
		return stackPackList[i].Name < stackPackList[j].Name
	})

	data := make([][]interface{}, 0)
	stackpacks := make([]stackstate_api.FullStackPack, 0)
	for _, stackpack := range stackPackList {
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

func getVersion(data *stackstate_api.FullStackPack) string {
	if data != nil && data.GetVersion() != "" {
		return data.GetVersion()
	}
	return "-"
}
