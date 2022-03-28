package settings

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

const (
	FileFlag = "file"
)

func SettingsApplyCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply -f FILE",
		Short: "apply settings with STJ",
		RunE:  di.CmdRunEWithDeps(cli, RunSettingsApplyCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", ".stj file to import")
	cmd.MarkFlagRequired(FileFlag)

	return cmd
}

func RunSettingsApplyCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return common.NewCLIError(err)
	}

	client, _, err := cli.Client.Connect()
	if err != nil {
		return common.NewConnectError(err)
	}

	nodes, resp, err := client.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes)).Execute()
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
