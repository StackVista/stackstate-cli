package settings

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

var fileMode = 0644

func SettingsDescribeCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe {--ids IDS | --namespace NAMESPACE | --type TYPE}",
		Short: "describe settings in STJ format",
		Long:  "Describe settings in StackState Templated JSON.",
		RunE:  cli.CmdRunEWithApi(RunSettingsDescribeCommand),
	}
	cmd.Flags().Int64Slice(Ids, nil, "list of ids to describe")
	cmd.Flags().String(Namespace, "", "namespace to describe")
	cmd.Flags().StringSlice(TypeName, nil, "list of types to describe")
	cmd.Flags().StringSlice(AllowReferences, nil, "white list of namespaces that are allowed to be referenced (only usable with the --namespace flag)")
	cmd.Flags().StringP(FileFlag, "f", "", "path of the output file")
	common.MarkMutexFlags(cmd, []string{Ids, Namespace, TypeName}, "filter", true)

	return cmd
}

func RunSettingsDescribeCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo stackstate_client.ServerInfo) common.CLIError {
	if err := common.CheckMutuallyExclusiveFlags(cmd, []string{Ids, Namespace, TypeName}, true); err != nil {
		return err
	}

	ids, err := cmd.Flags().GetInt64Slice(Ids)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	namespace, err := cmd.Flags().GetString(Namespace)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	references, err := cmd.Flags().GetStringSlice(AllowReferences)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	nodeTypes, err := cmd.Flags().GetStringSlice(TypeName)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	filepath, err := cmd.Flags().GetString(FileFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
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
			return common.NewCLIArgParseError(fmt.Errorf("\"%s\" flag is required for use of the \"%s\" flag", Namespace, AllowReferences))
		}
		exportArgs.AllowReferences = &references
	}

	data, resp, err := api.ExportApi.ExportSettings(cli.Context).Export(*exportArgs).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if filepath != "" {
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.FileMode(fileMode))
		if err != nil {
			return common.NewWriteFileError(err, filepath)
		}
		defer file.Close()

		if _, err = file.Write([]byte(data)); err != nil {
			return common.NewWriteFileError(err, filepath)
		}
		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"describe-file-path": filepath,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("settings exported to: %s", filepath))
		}

		return nil
	} else {
		if cli.IsJson {
			cli.Printer.PrintJson(map[string]interface{}{
				"data":   data,
				"format": "stj",
			})
		} else {
			cli.Printer.PrintLn(data)
		}
	}
	return nil
}
