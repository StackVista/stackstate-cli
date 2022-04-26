package stackpack

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func StackpackListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list stackpacks",
		Long:  "List available StackPacks.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListCommand),
	}
	cmd.Flags().Bool(InstalledFlag, false, "show only installed StackPacks")
	common.AddJsonFlag(cmd)
	return cmd
}

func RunStackpackListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	json, err := cmd.Flags().GetBool(common.JsonFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	isInstalled, err := cmd.Flags().GetBool(InstalledFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	data := make([][]interface{}, 0)
	respData := make([]stackstate_client.Sstackpack, 0)
	for _, v := range stackpackList {
		if isInstalled && len(v.GetConfigurations()) == 0 {
			continue // skip as this is not installed
		}
		row := []interface{}{
			v.Name,
			v.DisplayName,
			v.Version,
			getVersion(v.NextVersion),
			getVersion(v.LatestVersion),
			len(*v.Configurations),
		}

		data = append(data, row)
		respData = append(respData, v)
	}

	if json {
		cli.Printer.PrintJson(respData)
	} else {
		cli.Printer.Table(printer.TableData{
			Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
		})
	}

	return nil
}

func getVersion(data *stackstate_client.SstackpackLatestVersion) string {
	if data != nil && data.GetVersion() != "" {
		return data.GetVersion()
	}
	return "-"
}
