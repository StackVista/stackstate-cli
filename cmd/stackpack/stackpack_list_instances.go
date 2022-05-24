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
	args := &ListArgs{}
	cmd := &cobra.Command{
		Use:   "list-instances",
		Short: "list installed instances of a StackPack",
		Long:  "List all installed instances of a StackPack.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListInstanceCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "name of the instance")
	return cmd
}

func RunStackpackListInstanceCommand(args *ListArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		data := make([][]interface{}, 0)
		instances := make([]stackstate_api.SstackpackConfigurations, 0)
		for _, v := range stackpackList {
			if v.GetName() != args.Name {
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

		if cli.IsJson() {
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
}
