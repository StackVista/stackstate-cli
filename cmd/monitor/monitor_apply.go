package monitor

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func MonitorApplyCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply -f FILE",
		Short: "apply a monitor with STJ",
		RunE:  cli.CmdRunEWithApi(RunMonitorApplyCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", FileFlagUsage)
	cmd.MarkFlagRequired(FileFlag)

	return cmd
}

func RunMonitorApplyCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo stackstate_client.ServerInfo) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return common.NewCLIError(err)
	}

	nodes, resp, err := api.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes)).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if len(nodes) == 0 {
		cli.Printer.PrintWarn("Nothing was imported.")
		return nil
	}

	tableData := [][]string{}
	for _, node := range nodes {
		tableData = append(tableData, []string{
			util.ToString(node["_type"]),
			util.ToString(node["id"]),
			util.ToString(node["identifier"]),
			util.ToString(node["name"]),
		})
	}

	cli.Printer.Success(fmt.Sprintf("Applied <bold>%d</> setting node(s).", len(nodes)))
	cli.Printer.PrintLn("")
	cli.Printer.Table([]string{"Type", "Id", "Identifier", "Name"}, tableData, nodes)
	return nil
}
