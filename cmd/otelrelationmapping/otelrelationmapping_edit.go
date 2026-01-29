package otelrelationmapping

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const LongDescription = `Edit an OTel Relation Mapping.

The edit command allows you to directly edit any OTel Relation Mapping. It will open
the editor defined by your EDITOR environment variables.

The OTel Relation Mapping will be presented as JSON format for editing.
`

type EditArgs struct {
	Identifier string
}

func OtelRelationMappingEditCommand(deps *di.Deps) *cobra.Command {
	args := &EditArgs{}
	cmd := &cobra.Command{
		Use:   "edit --identifier URN",
		Short: "Edit an OTel Relation Mapping using $EDITOR",
		Long:  LongDescription,
		Example: `# edit a relation mapping using your editor
sts otel-relation-mapping edit --identifier urn:stackpack:stackpack-name:shared:otel-relation-mapping:database`,
		RunE: deps.CmdRunEWithApi(RunEditRelationMappingCommand(args)),
	}
	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the Relation Mapping")
	return cmd
}

func RunEditRelationMappingCommand(args *EditArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		mapping, resp, err := api.OtelMappingApi.GetOtelRelationMapping(cli.Context, args.Identifier).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		originalYAML, err := yaml.Marshal(mapping)
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to marshal mapping to YAML: %v", err))
		}

		editedContent, err := cli.Editor.Edit("otel-relation-mapping-", ".yaml", strings.NewReader(string(originalYAML)))
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to open editor: %v", err))
		}

		if string(originalYAML) == string(editedContent) {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{"message": "No changes made"})
			} else {
				cli.Printer.PrintWarn("No changes made")
			}
			return nil
		}

		var editedMapping stackstate_api.OtelRelationMapping
		if err := yaml.Unmarshal(editedContent, &editedMapping); err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to parse edited YAML: %v", err))
		}

		reqObj := stackstate_api.UpsertOtelRelationMappingsRequest{
			Identifier:  editedMapping.Identifier,
			Name:        editedMapping.Name,
			Description: editedMapping.Description,
			Input:       editedMapping.Input,
			Output:      editedMapping.Output,
			Vars:        editedMapping.Vars,
			ExpireAfter: editedMapping.ExpireAfter,
		}
		upserted, resp, err := api.OtelMappingApi.UpsertOtelRelationMappings(cli.Context).UpsertOtelRelationMappingsRequest(reqObj).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{"otel_relation_mapping": upserted})
		} else {
			cli.Printer.Success(fmt.Sprintf("OTel Relation Mapping updated successfully! Identifier: %s, Name: %s", upserted.GetIdentifier(), upserted.GetName()))
		}
		return nil
	}
}
