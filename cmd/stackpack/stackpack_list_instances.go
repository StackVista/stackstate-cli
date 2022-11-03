package stackpack

import (
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func StackpackListInstanceCommand(cli *di.Deps) *cobra.Command {
	args := &ListPropertiesArgs{}
	cmd := &cobra.Command{
		Use:   "list-instances",
		Short: "List installed instances of a StackPack",
		Long:  "List all installed instances of a StackPack.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListInstanceCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the instance")
	return cmd
}

func formatInstancesJson(v stackstate_api.FullStackPack) []stackstate_api.StackPackConfiguration {
	instances := v.GetConfigurations()

	sort.SliceStable(instances, func(i, j int) bool {
		return *instances[i].LastUpdateTimestamp > *instances[j].LastUpdateTimestamp
	})

	return instances
}

func formatInstancesTable(v stackstate_api.FullStackPack) [][]interface{} {
	instances := v.GetConfigurations()

	sort.SliceStable(instances, func(i, j int) bool {
		return *instances[i].LastUpdateTimestamp > *instances[j].LastUpdateTimestamp
	})

	data := make([][]interface{}, 0)
	for _, instance := range instances {
		row := []interface{}{
			instance.Id,
			instance.Status,
			instance.StackPackVersion,
			time.UnixMilli(instance.GetLastUpdateTimestamp()),
		}
		data = append(data, row)
	}

	return data
}

func RunStackpackListInstanceCommand(args *ListPropertiesArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackPackList, resp, err := api.StackpackApi.StackPackList(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		data := [][]interface{}{}
		instances := []stackstate_api.StackPackConfiguration{}

		for _, stackpack := range stackPackList {
			if stackpack.GetName() != args.Name {
				continue
			}
			data = formatInstancesTable(stackpack)
			instances = formatInstancesJson(stackpack)
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
