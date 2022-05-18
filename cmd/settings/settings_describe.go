package settings

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/mutex_flags"
)

var fileMode = 0644

type DescribeArgs struct {
	Ids             []int64
	Namespace       string
	NodeTypes       []string
	AllowReferences []string
	FilePath        string
}

func SettingsDescribeCommand(cli *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "describe settings in STJ format",
		Long:  "Describe settings in StackState Templated JSON.",
		RunE:  cli.CmdRunEWithApi(RunSettingsDescribeCommand(args)),
	}
	cmd.Flags().Int64SliceVar(&args.Ids, IdsFlag, nil, "list of ids to describe")
	cmd.Flags().StringVar(&args.Namespace, Namespace, "", "namespace to describe")
	cmd.Flags().StringSliceVar(&args.NodeTypes, TypeNameFlag, nil, "list of types to describe")
	cmd.Flags().StringSliceVar(&args.AllowReferences, AllowReferencesFlag, nil, "white list of namespaces that are allowed to be referenced (only usable with the --namespace flag)")
	common.AddFileFlagVar(cmd, &args.FilePath, "path to the output file")
	mutex_flags.MarkMutexFlags(cmd, []string{IdsFlag, Namespace, TypeNameFlag}, "filter", true)
	return cmd
}

func RunSettingsDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		if err := mutex_flags.CheckMutuallyExclusiveFlags(cmd, []string{IdsFlag, Namespace, TypeNameFlag}, true); err != nil {
			return common.NewCLIArgParseError(err)
		}

		exportArgs := stackstate_api.NewExport()
		if len(args.Ids) != 0 {
			exportArgs.NodesWithIds = args.Ids
		}
		if len(args.Namespace) != 0 {
			exportArgs.Namespace = &args.Namespace
		}
		if len(args.NodeTypes) != 0 {
			exportArgs.AllNodesOfTypes = args.NodeTypes
		}
		if len(args.AllowReferences) != 0 {
			if len(args.Namespace) == 0 {
				return common.NewCLIArgParseError(fmt.Errorf("\"%s\" flag is required for use of the \"%s\" flag", Namespace, AllowReferencesFlag))
			}
			exportArgs.AllowReferences = args.AllowReferences
		}

		data, resp, err := api.ExportApi.ExportSettings(cli.Context).Export(*exportArgs).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if args.FilePath != "" {
			file, err := os.OpenFile(args.FilePath, os.O_CREATE|os.O_WRONLY, os.FileMode(fileMode))
			if err != nil {
				return common.NewWriteFileError(err, args.FilePath)
			}
			defer file.Close()

			if _, err = file.Write([]byte(data)); err != nil {
				return common.NewWriteFileError(err, args.FilePath)
			}
			if cli.IsJson {
				cli.Printer.PrintJson(map[string]interface{}{
					"filepath": args.FilePath,
				})
			} else {
				cli.Printer.Success(fmt.Sprintf("settings exported to: %s", args.FilePath))
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
}
