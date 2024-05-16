package stackpack

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type ListArgs struct {
	Installed bool
}

func StackpackListCommand(cli *di.Deps) *cobra.Command {
	args := &ListArgs{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List stackpacks",
		Long:  "List available StackPacks.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListCommand(args)),
	}
	cmd.Flags().BoolVar(&args.Installed, InstalledFlag, false, "Show only installed StackPacks")
	return cmd
}

func fetchAllStackPacks(cli *di.Deps, api *stackstate_api.APIClient) ([]stackstate_api.FullStackPack, common.CLIError) {
	stackPackList, resp, err := api.StackpackAPI.StackPackList(cli.Context).Execute()
	if err != nil {
		return nil, common.NewResponseError(err, resp)
	}

	sort.SliceStable(stackPackList, func(i, j int) bool {
		return stackPackList[i].Name < stackPackList[j].Name
	})

	return stackPackList, nil
}

func RunStackpackListCommand(args *ListArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackPackList, err := fetchAllStackPacks(cli, api)
		if err != nil {
			return err
		}

		data := make([][]interface{}, 0)
		stackpacks := make([]stackstate_api.FullStackPack, 0)
		for _, stackpack := range stackPackList {
			if args.Installed && len(stackpack.GetConfigurations()) == 0 {
				continue // skip as this is not installed
			}
			row := []interface{}{
				stackpack.Name,
				stackpack.DisplayName,
				stackpack.Version,
				getVersion(stackpack.NextVersion),
				getVersion(stackpack.LatestVersion),
				len(stackpack.Configurations),
			}

			data = append(data, row)
			stackpacks = append(stackpacks, stackpack)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"stackpacks": stackpacks,
			})
		} else {
			cli.Printer.Table(printer.TableData{
				Header:              []string{"name", "display name", "installed version", "next version", "latest version", "instance count"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: "StackPacks"},
			})
		}

		return nil
	}
}

func getVersion(data *stackstate_api.FullStackPack) string {
	if data != nil && data.GetVersion() != "" {
		return data.GetVersion()
	}
	return "-"
}
