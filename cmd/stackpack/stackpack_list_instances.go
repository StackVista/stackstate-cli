package stackpack

import (
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func StackpackListInstanceCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-instances",
		Short: "list installed instances of a StackPack",
		Long:  "List all installed instances of a StackPack.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListInstanceCommand),
	}
	common.AddRequiredNameFlag(cmd, "name of the instance")
	return cmd
}

func RunStackpackListInstanceCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	name, err := cmd.Flags().GetString(common.NameFlag)
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
	instances := make([]stackstate_api.SstackpackConfigurations, 0)
	for _, v := range stackpackList {
		if v.GetName() != name {
			continue
		}
		for _, instance := range v.GetConfigurations() {
			row := []interface{}{
				instance.Id,
				instance.Status,
				instance.StackPackVersion,
				time.UnixMilli(instance.GetLastUpdateTimestamp()),
			}
			data = append(data, row)
			instances = append(instances, instance)
		}
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"instances": instances,
		})
	} else {
		cli.Printer.Table(printer.TableData{
			Header:              []string{"id", "status", "version", "last updated"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
		})
	}
	return nil
}
