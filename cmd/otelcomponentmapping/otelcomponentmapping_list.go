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
		Short: "Lists active OTel Component Mappings",
		Long:  "Lists active OTel Component Mappings.",
		Example: `# list all OTel component mappings
sts otel-component-mapping list

# list all component mappings in JSON format
sts otel-component-mapping list -o json`,
		RunE: deps.CmdRunEWithApi(RunListComponentCommand),
	}

	return cmd
}

func RunListComponentCommand(_ *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, _ *stackstate_api.ServerInfo) common.CLIError {
	mappingsList, resp, err := api.OtelMappingApi.GetOtelComponentMappings(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	sort.SliceStable(mappingsList, func(i, j int) bool {
		return mappingsList[i].Name < mappingsList[j].Name
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"otel_component_mappings": mappingsList,
		})
	} else {
		cli.Printer.Table(otelmapping.FormatOtelMappingTable(mappingsList))
	}

	return nil
}
