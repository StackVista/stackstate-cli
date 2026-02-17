package otelrelationmapping

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type ApplyArgs struct {
	File string
}

func OtelRelationMappingApplyCommand(deps *di.Deps) *cobra.Command {
	args := &ApplyArgs{}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Create or edit an OTel Relation Mapping from YAML",
		Long:  "Create or edit an OTel Relation Mapping from YAML file.",
		Example: `# create a new OTel relation mapping from a YAML file
sts otel-relation-mapping apply --file new-relation-mapping.yaml

# update an existing relation mapping
sts otel-relation-mapping apply --file updated-relation-mapping.yaml`,
		RunE: deps.CmdRunEWithApi(RunApplyRelationMappingCommand(args)),
	}

	common.AddRequiredFileFlagVar(cmd, &args.File, "Path to .yaml file with the mapping definition")

	return cmd
}

func RunApplyRelationMappingCommand(args *ApplyArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		fileBytes, err := os.ReadFile(args.File)
		if err != nil {
			return common.NewReadFileError(err, args.File)
		}

		ext := strings.ToLower(filepath.Ext(args.File))
		if ext != ".yaml" {
			return common.NewCLIArgParseError(fmt.Errorf("unsupported file type: %s. Only .yaml files are supported", ext))
		}

		return applyYAMLOtelRelationMapping(cli, api, fileBytes)
	}
}

func applyYAMLOtelRelationMapping(cli *di.Deps, api *stackstate_api.APIClient, fileBytes []byte) common.CLIError {
	var mapping stackstate_api.OtelRelationMapping
	if err := yaml.Unmarshal(fileBytes, &mapping); err != nil {
		return common.NewCLIArgParseError(fmt.Errorf("failed to parse YAML: %v", err))
	}

	reqObj := stackstate_api.UpsertOtelRelationMappingsRequest{
		Identifier:  mapping.Identifier,
		Name:        mapping.Name,
		Description: mapping.Description,
		Input:       mapping.Input,
		Output:      mapping.Output,
		Vars:        mapping.Vars,
		ExpireAfter: mapping.ExpireAfter,
	}

	upserted, resp, err := api.OtelMappingApi.UpsertOtelRelationMappings(cli.Context).UpsertOtelRelationMappingsRequest(reqObj).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}
	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{"otel_relation_mapping": upserted})
	} else {
		cli.Printer.Success(fmt.Sprintf("OTel Relation Mapping upserted successfully! Identifier: %s, Name: %s", mapping.GetIdentifier(), mapping.GetName()))
	}
	return nil
}
