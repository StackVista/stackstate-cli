package monitor

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type ApplyArgs struct {
	File string
}

func MonitorApplyCommand(cli *di.Deps) *cobra.Command {
	args := &ApplyArgs{}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "apply a monitor with STJ",
		Long:  "Apply a monitor with StackState Templated JSON.",
		RunE:  cli.CmdRunEWithApi(RunMonitorApplyCommand(args)),
	}

	common.AddRequiredFileFlagVar(cmd, &args.File, FileFlagUsage)

	return cmd
}

func RunMonitorApplyCommand(args *ApplyArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		fileBytes, err := os.ReadFile(args.File)
		if err != nil {
			return common.NewCLIArgParseError(err)
		}

		nodes, resp, err := api.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes)).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"nodes": nodes,
			})
		} else {
			if len(nodes) == 0 {
				cli.Printer.PrintWarn("Nothing was imported.")
				return nil
			}

			tableData := make([][]interface{}, 0)
			for _, node := range nodes {
				tableData = append(tableData, []interface{}{node["_type"], node["id"], node["identifier"], node["name"]})
			}

			cli.Printer.Success(fmt.Sprintf("Applied <bold>%d</> monitor(s).", len(nodes)))
			if len(nodes) > 0 {
				cli.Printer.Table(printer.TableData{
					Header: []string{"Type", "Id", "Identifier", "Name"},
					Data:   tableData,
				})
			}
		}
		return nil
	}
}
