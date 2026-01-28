package otelcomponentmapping

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const LongDescription = `Edit an OTel Component Mapping.

The edit command allows you to directly edit any OTel Component Mapping. It will open
the editor defined by your EDITOR environment variables.

The OTel Component Mapping will be presented as JSON format for editing.
`

type EditArgs struct {
	Identifier string
}

func OtelComponentMappingEditCommand(deps *di.Deps) *cobra.Command {
	args := &EditArgs{}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit an OTel Component Mapping using $EDITOR",
		Long:  LongDescription,
		RunE:  deps.CmdRunEWithApi(RunEditComponentMappingCommand(args)),
	}

	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the Component Mapping")

	return cmd
}

func RunEditComponentMappingCommand(args *EditArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		if args.Identifier == "" {
			return common.NewCLIArgParseError(fmt.Errorf("--identifier is required"))
		}

		mapping, resp, err := api.OtelMappingApi.GetOtelComponentMapping(cli.Context, args.Identifier).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		originalYAML, err := yaml.Marshal(mapping)
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to marshal mapping to YAML: %v", err))
		}

		editedContent, err := cli.Editor.Edit("otel-component-mapping-", ".yaml", strings.NewReader(string(originalYAML)))
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to open editor: %v", err))
		}

		// Check if changes were actually made
		if string(originalYAML) == string(editedContent) {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{"message": "No changes made"})
			} else {
				cli.Printer.PrintWarn("No changes made")
			}
			return nil
		}

		var editedMapping stackstate_api.OtelComponentMapping
		if err := yaml.Unmarshal(editedContent, &editedMapping); err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to parse edited YAML: %v", err))
		}

		reqObj := stackstate_api.UpsertOtelComponentMappingsRequest{
			Identifier:  editedMapping.Identifier,
			Name:        editedMapping.Name,
			Description: editedMapping.Description,
			Input:       editedMapping.Input,
			Output:      editedMapping.Output,
			Vars:        editedMapping.Vars,
			ExpireAfter: editedMapping.ExpireAfter,
		}
		upserted, resp, err := api.OtelMappingApi.UpsertOtelComponentMappings(cli.Context).UpsertOtelComponentMappingsRequest(reqObj).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{"otel_component_mapping": upserted})
		} else {
			cli.Printer.Success(fmt.Sprintf("OTel Component Mapping updated successfully! Identifier: %s, Name: %s", upserted.GetIdentifier(), upserted.GetName()))
		}
		return nil
	}
}
