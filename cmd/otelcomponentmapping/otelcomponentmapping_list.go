package otelcomponentmapping

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func OtelComponentMappingListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists active Otel Component Mappings",
		Long:  "Lists active Otel Component Mappings.",
		RunE:  deps.CmdRunEWithApi(RunListComponentCommand),
	}

	return cmd
}

func RunListComponentCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	mappingsList, resp, err := api.OtelMappingApi.GetOtelComponentMappings(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	sort.SliceStable(mappingsList, func(i, j int) bool {
		return mappingsList[i].Name < mappingsList[j].Name
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"otel component mappings": mappingsList,
		})
	} else {
		cli.Printer.Table(otelmapping.FormatOtelMappingTable(mappingsList))
	}

	return nil
}
