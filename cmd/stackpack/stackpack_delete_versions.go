package stackpack

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

const (
	FromFlag = "from"
	ToFlag   = "to"
)

type DeleteVersionsArgs struct {
	Name string
	From string
	To   string
	All  bool
	Dev  bool
}

func StackpackDeleteVersionsCommand(cli *di.Deps) *cobra.Command {
	args := &DeleteVersionsArgs{}
	cmd := &cobra.Command{
		Use:   "delete-versions",
		Short: "Delete multiple versions of a StackPack",
		Long:  "Delete multiple versions of a StackPack by specifying a version range or by deleting all versions. Versions currently in use by installed instances are skipped automatically.",
		Example: `# delete all versions up to and including 1.5.0
sts stackpack delete-versions --name kubernetes --to 1.5.0

# delete versions in a specific range
sts stackpack delete-versions --name kubernetes --from 1.0.0 --to 1.5.0

# delete all development (SNAPSHOT) versions
sts stackpack delete-versions --name kubernetes --all --dev`,
		RunE: cli.CmdRunEWithApi(RunStackpackDeleteVersionsCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the StackPack")
	cmd.Flags().StringVar(&args.From, FromFlag, "", "Inclusive lower bound: delete versions >= this version")
	cmd.Flags().StringVar(&args.To, ToFlag, "", "Inclusive upper bound: delete versions <= this version")
	cmd.Flags().BoolVar(&args.All, AllFlag, false, "Delete all versions (cannot be combined with --from or --to)")
	cmd.Flags().BoolVar(&args.Dev, DevFlag, false, "Restrict to development versions only (e.g. SNAPSHOT)")
	return cmd
}

func RunStackpackDeleteVersionsCommand(args *DeleteVersionsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		if !args.All && args.From == "" && args.To == "" {
			return common.NewCLIArgParseError(fmt.Errorf("at least one of --all, --from, or --to must be specified"))
		}
		if args.All && (args.From != "" || args.To != "") {
			return common.NewCLIArgParseError(fmt.Errorf("--all cannot be combined with --from or --to"))
		}

		req := api.StackpackApi.StackPackDeleteVersions(cli.Context, args.Name)
		if args.From != "" {
			req = req.From(args.From)
		}
		if args.To != "" {
			req = req.To(args.To)
		}
		if args.All {
			req = req.All(true)
		}
		if args.Dev {
			req = req.Dev(true)
		}

		result, resp, err := req.Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"deleted":      result.Deleted,
				"skippedInUse": result.SkippedInUse,
			})
		} else {
			if len(result.Deleted) > 0 {
				cli.Printer.Success(fmt.Sprintf("Deleted versions: %s", strings.Join(result.Deleted, ", ")))
			}
			if len(result.SkippedInUse) > 0 {
				cli.Printer.Table(printer.TableData{
					Header: []string{"skipped (in use)"},
					Data: func() [][]interface{} {
						rows := make([][]interface{}, 0, len(result.SkippedInUse))
						for _, v := range result.SkippedInUse {
							rows = append(rows, []interface{}{v})
						}
						return rows
					}(),
				})
			}
			if len(result.Deleted) == 0 && len(result.SkippedInUse) == 0 {
				cli.Printer.Success("No versions matched the specified criteria")
			}
		}

		return nil
	}
}
