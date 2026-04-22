package stackpack

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type ListVersionsArgs struct {
	Name string
}

func StackpackListVersionsCommand(cli *di.Deps) *cobra.Command {
	args := &ListVersionsArgs{}
	cmd := &cobra.Command{
		Use:   "list-versions",
		Short: "List all available versions of a StackPack",
		Long:  "List all available versions of a StackPack. Shows version string, whether it is a development version, and whether it is currently in use by an installed instance.",
		Example: `# list all versions of the kubernetes StackPack
sts stackpack list-versions --name kubernetes

# list versions as JSON
sts stackpack list-versions --name kubernetes -o json`,
		RunE: cli.CmdRunEWithApi(RunStackpackListVersionsCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the StackPack")
	return cmd
}

func RunStackpackListVersionsCommand(args *ListVersionsArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		versions, resp, err := api.StackpackApi.StackPackListVersions(cli.Context, args.Name).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		sort.SliceStable(versions, func(i, j int) bool {
			return versions[i].Version < versions[j].Version
		})

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"versions": versions,
			})
		} else {
			data := make([][]interface{}, 0, len(versions))
			for _, v := range versions {
				data = append(data, []interface{}{
					v.Version,
					v.IsDev,
					v.IsInUse,
				})
			}
			cli.Printer.Table(printer.TableData{
				Header:              []string{"version", "dev", "in use"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: "versions"},
			})
		}

		return nil
	}
}
