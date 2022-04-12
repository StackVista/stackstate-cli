package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

var (
	UnlockedStrategyChoices = []string{"fail", "skip", "overwrite"}
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
	cmd.Flags().String(
		UnlockedStrategyFlag,
		"",
		"strategy to use when encountering unlocked settings when applying settings to a namespace"+
			fmt.Sprintf(" (must be { %s })", strings.Join(UnlockedStrategyChoices, " | ")))
	cmd.Flags().IntP(TimeoutFlag, "t", 0, "timeout in seconds")
	cmd.MarkFlagRequired(FileFlag) //nolint:errcheck

	return cmd
}

//nolint:funlen
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
	unlockedStrategy, err := cmd.Flags().GetString(UnlockedStrategyFlag)
	if err != nil {
		return common.NewCLIError(err)
	}
	if unlockedStrategy != "" {
		if err := common.CheckFlagIsValidChoice(UnlockedStrategyFlag, unlockedStrategy, UnlockedStrategyChoices); err != nil {
			return err
		}
	}
	timeout, err := cmd.Flags().GetInt(TimeoutFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return common.NewCLIError(err)
	}

	request := api.ImportApi.ImportSettings(cli.Context).Body(string(fileBytes))
	if namespace != "" {
		request = request.Namespace(namespace)
	}
	if unlockedStrategy != "" {
		request = request.Unlocked(unlockedStrategy)
	}
	if timeout > 0 {
		request = request.TimeoutSeconds(int64(timeout))
	}

	nodes, resp, err := request.Execute()
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
