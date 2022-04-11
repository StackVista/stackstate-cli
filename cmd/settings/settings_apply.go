package settings

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

const (
	FileFlag      = "file"
	NamespaceFlag = "namespace"
)

func SettingsApplyCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply -f FILE",
		Short: "apply settings with STJ",
		RunE:  cli.CmdRunEWithApi(RunSettingsApplyCommand),
	}
	cmd.Flags().StringP(FileFlag, "f", "", ".stj file to import")
	cmd.Flags().StringP(NamespaceFlag, "n", "", "name of the namespace to overwrite"+
		" - WARNING this will overwrite the entire namespace")
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck

	return cmd
}

func RunSettingsApplyCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	file, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}
	namespace, err := cmd.Flags().GetString(NamespaceFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return common.NewCLIError(err)
	}

	req := api.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes))
	if namespace != "" {
		req = req.Namespace(namespace)
	}
	nodes, resp, err := req.Execute()
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

	cli.Printer.Success(fmt.Sprintf("Applied <bold>%d</> setting node(s).\n", len(nodes)))
	cli.Printer.Table([]string{"Type", "Id", "Identifier", "Name"}, tableData, nodes)

	return nil
}
