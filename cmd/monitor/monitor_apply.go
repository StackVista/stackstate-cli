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

func MonitorApplyCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply -f FILE",
		Short: "apply a monitor with STJ",
		Long:  "Apply a monitor with StackState Templated JSON.",
		RunE:  cli.CmdRunEWithApi(RunMonitorApplyCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", FileFlagUsage)
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck

	return cmd
}

func RunMonitorApplyCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo stackstate_api.ServerInfo,
) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	nodes, resp, err := api.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes)).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if len(nodes) == 0 {
		cli.Printer.PrintWarn("Nothing was imported.")
		return nil
	}

	tableData := make([][]interface{}, 0)
	for _, node := range nodes {
		tableData = append(tableData, []interface{}{node["_type"], node["id"], node["identifier"], node["name"]})
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"nodes": nodes,
		})
	} else {
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
