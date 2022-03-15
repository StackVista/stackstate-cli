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

func SettingsImportCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import settings with STJ.",
		RunE:  di.CmdRunEWithDeps(cli, RunSettingsImportCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", "The .stj file to import.")
	cmd.MarkFlagRequired(FileFlag)

	return cmd
}

func RunSettingsImportCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return common.NewCLIError(err)
	}

	nodes, resp, err := cli.Client.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes)).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	// TODO: what if len nodes == 0

	data := [][]string{}
	for _, node := range nodes {
		data = append(data, []string{util.ToString(node["_type"]), util.ToString(node["id"]), util.ToString(node["name"])})
	}

	cli.Printer.Success(fmt.Sprintf("Imported %d setting node(s).", len(nodes)))
	cli.Printer.PrintLn("")
	cli.Printer.Table([]string{"Type", "Id", "Name"}, data)

	return nil
}
