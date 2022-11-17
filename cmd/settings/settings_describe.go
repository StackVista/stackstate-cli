package settings

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	stscobra "github.com/stackvista/stackstate-cli/internal/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/util"
)

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
		Short: "Describe settings in STY format",
		Long:  "Describe settings in StackState Templated YAML.",
		RunE:  cli.CmdRunEWithApi(RunSettingsDescribeCommand(args)),
	}
	cmd.Flags().Int64SliceVar(&args.Ids, IdsFlag, nil, "List of ids to describe")
	cmd.Flags().StringVar(&args.Namespace, Namespace, "", "Namespace to describe")
	cmd.Flags().StringSliceVar(&args.NodeTypes, TypeNameFlag, nil, "List of types to describe")
	cmd.Flags().StringSliceVar(&args.AllowReferences, AllowReferencesFlag, nil, "White list of namespaces that are allowed to be referenced (only usable with the --namespace flag)")
	common.AddFileFlagVar(cmd, &args.FilePath, "Path to the output file")
	stscobra.MarkMutexFlags(cmd, []string{IdsFlag, Namespace, TypeNameFlag}, "filter", true)
	return cmd
}

func RunSettingsDescribeCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		if err := stscobra.CheckMutuallyExclusiveFlags(cmd, []string{IdsFlag, Namespace, TypeNameFlag}, true); err != nil {
			return common.NewCLIArgParseError(err)
		}

		if len(args.AllowReferences) != 0 {
			if len(args.Namespace) == 0 {
				return common.NewCLIArgParseError(fmt.Errorf("\"%s\" flag is required for use of the \"%s\" flag", Namespace, AllowReferencesFlag))
			}
		}

		data, resp, err := DoExport(cli.Context, api, args.Ids, args.Namespace, args.NodeTypes, args.AllowReferences)
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if args.FilePath != "" {
			if err := util.WriteFile(args.FilePath, []byte(data)); err != nil {
				return common.NewWriteFileError(err, args.FilePath)
			}

			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"filepath": args.FilePath,
				})
			} else {
				cli.Printer.Success(fmt.Sprintf("settings exported to: %s", args.FilePath))
			}

			return nil
		} else {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"data":   data,
					"format": "sty",
				})
			} else {
				cli.Printer.PrintLn(data)
			}
		}
		return nil
	}
}
