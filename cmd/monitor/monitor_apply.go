package monitor

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type ApplyArgs struct {
	File string
}

func MonitorApplyCommand(cli *di.Deps) *cobra.Command {
	args := &ApplyArgs{}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Create or edit a monitor from STY",
		Long:  "Create or edit a monitor from StackState Templated YAML.",
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

		nodes, resp, err := api.ImportAPI.ImportSettings(cli.Context).Body(string(fileBytes)).Execute()
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
				tableData = append(tableData, []interface{}{node["_type"], node["id"], node["status"], node["identifier"], node["name"]})
			}

			cli.Printer.Success(fmt.Sprintf("Applied <bold>%d</> monitor(s).", len(nodes)))
			if len(nodes) > 0 {
				cli.Printer.Table(printer.TableData{
					Header: []string{"Type", "Id", "Status", "Identifier", "Name"},
					Data:   tableData,
				})
			}
		}

		return nil
	}
}
