package stackpack

import (
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func StackpackListInstanceCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-instances --name NAME",
		Short: "list installed instances of a StackPack",
		Long:  "List all installed instances of a StackPack.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListInstanceCommand),
	}
	cmd.Flags().String(NameFlag, "", "name of the instance")
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
	return cmd
}

func RunStackpackListInstanceCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	name, err := cmd.Flags().GetString(NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	if name == "" {
		return common.NewCLIArgParseError(err)
	}
	stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	data := make([][]interface{}, 0)
	respData := make([]stackstate_client.SstackpackConfigurations, 0)
	for _, v := range stackpackList {
		if v.GetName() != name {
			continue
		}
		for _, conf := range v.GetConfigurations() {
			row := []interface{}{
				conf.Id,
				conf.Status,
				conf.StackPackVersion,
				time.UnixMilli(conf.GetLastUpdateTimestamp()),
			}
			data = append(data, row)
			respData = append(respData, conf)
		}
	}
	cli.Printer.Table(printer.TableData{
		Header:              []string{"id", "status", "version", "last updated"},
		Data:                data,
		StructData:          respData,
		MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
	})
	return nil
}
