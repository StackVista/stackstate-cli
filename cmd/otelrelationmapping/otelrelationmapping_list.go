package otelrelationmapping

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

type ListArgs struct{}

func OtelRelationMappingListCommand(deps *di.Deps) *cobra.Command {
	args := &ListArgs{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists active OTel Relation Mappings",
		Long:  "Lists active OTel Relation Mappings.",
		Example: `# list all OTel relation mappings
sts otel-relation-mapping list

# list all relation mappings in JSON format
sts otel-relation-mapping list -o json`,
		RunE: deps.CmdRunEWithApi(RunListRelationMappingCommand(args)),
	}

	return cmd
}

func RunListRelationMappingCommand(_ *ListArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		mappingsList, resp, err := api.OtelMappingApi.GetOtelRelationMappings(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		sort.SliceStable(mappingsList, func(i, j int) bool {
			return mappingsList[i].Name < mappingsList[j].Name
		})

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"otel_relation_mappings": mappingsList,
			})
		} else {
			cli.Printer.Table(otelmapping.FormatOtelMappingTable(mappingsList))
		}

		return nil
	}
}
