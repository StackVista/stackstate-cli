package settings

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

var fileMode = 0644

func SettingsExportCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe {--ids IDs | --namespace NAMESPACE | --type TYPE}",
		Short: "describe settings with STJ",
		RunE:  cli.CmdRunEWithApi(RunSettingsExportCommand),
	}
	cmd.Flags().Int64Slice(Ids, nil, "list of ids to describe")
	cmd.Flags().String(Namespace, "", "namespace to describe")
	cmd.Flags().StringSlice(TypeName, nil, "list of types to describe")
	cmd.Flags().StringSlice(AllowReferences, nil, "white list of namespaces that are allowed to be referenced (only usable with the --namespace flag)")
	cmd.Flags().StringP(FileFlag, "f", "", "path of the output file")

	return cmd
}

func RunSettingsExportCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo stackstate_client.ServerInfo) common.CLIError {
	if err := common.CheckRequiredMutuallyExclusiveFlags(cmd, []string{Ids, Namespace, TypeName}); err != nil {
		return err
	}

	//TODO check output type!!

	ids, err := cmd.Flags().GetInt64Slice(Ids)
	if err != nil {
		return common.NewCLIError(err)
	}
	namespace, err := cmd.Flags().GetString(Namespace)
	if err != nil {
		return common.NewCLIError(err)
	}
	references, err := cmd.Flags().GetStringSlice(AllowReferences)
	if err != nil {
		return common.NewCLIError(err)
	}
	nodeTypes, err := cmd.Flags().GetStringSlice(TypeName)
	if err != nil {
		return common.NewCLIError(err)
	}

	filePath, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIError(err)
	}

	exportArgs := stackstate_client.NewExport()
	if len(ids) != 0 {
		exportArgs.NodesWithIds = &ids
	}
	if len(namespace) != 0 {
		exportArgs.Namespace = &namespace
	}
	if len(nodeTypes) != 0 {
		exportArgs.AllNodesOfTypes = &nodeTypes
	}
	if len(references) != 0 {
		if len(namespace) == 0 {
			return common.NewCLIError(fmt.Errorf("\"%s\" flag is required for use of the \"%s\" flag", Namespace, AllowReferences))
		}
		exportArgs.AllowReferences = &references
	}

	data, resp, err := api.ExportApi.ExportSettings(cli.Context).Export(*exportArgs).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.FileMode(fileMode))
		if err != nil {
			return common.NewCLIError(err)
		}
		defer file.Close()

		if _, err = file.Write([]byte(data)); err != nil {
			return common.NewCLIError(err)
		}
		cli.Printer.Success(fmt.Sprintf("settings exported to: %s", filePath))
		return nil
	}
	cli.Printer.PrintLn(data)
	return nil
}
