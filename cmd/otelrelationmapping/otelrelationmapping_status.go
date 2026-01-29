package otelrelationmapping

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/otelmapping"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func OtelRelationMappingStatusCommand(deps *di.Deps) *cobra.Command {
	args := &otelmapping.StatusArgs{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of an OTel Relation Mappings",
		Long:  "Get the status of an OTel Relation Mappings.",
		Example: `# get the status of a relation mapping
sts otel-relation-mapping status --identifier urn:stackpack:stackpack-name:shared:otel-relation-mapping:database`,
		RunE: deps.CmdRunEWithApi(otelmapping.RunStatus(args, "relation", otelmapping.FetchRelationStatus)),
	}

	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, "Identifier of the OTel Relation Mapping")

	return cmd
}
