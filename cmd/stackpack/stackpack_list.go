package stackpack

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func StackpackListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list of stackpack",
		RunE:  cli.CmdRunEWithApi(RunStackpackListCommand),
	}
	cmd.Flags().Bool(InstalledFlag, false, "show only installed stackpack")
	return cmd
}

func RunStackpackListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo *stackstate_client.ServerInfo,
) common.CLIError {
	isInstalled, err := cmd.Flags().GetBool(InstalledFlag)
	if err != nil {
		return common.NewCLIError(err)
	}
	stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	data := make([][]interface{}, 0)
	respData := make([]stackstate_client.Stackpack, 0)
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
			len(v.Configurations),
		}

		data = append(data, row)
		respData = append(respData, v)
	}

	cli.Printer.Table(
		[]string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
		data,
		respData,
	)

	return nil
}

func getVersion(data *stackstate_client.StackpackLatestVersion) string {
	if data != nil && data.GetVersion() != "" {
		return data.GetVersion()
	}
	return "-"
}
