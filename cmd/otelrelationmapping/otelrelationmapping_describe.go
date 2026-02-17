package otelrelationmapping

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/util"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type DescribeArgs struct {
	Identifier string
	FilePath   string
}

func OtelRelationMappingDescribeCommand(deps *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe an OTel Relation Mapping by identifier (URN)",
		Long:  "Describe an OTel Relation Mapping by identifier (URN). Optionally write to output file.",
		Example: `# describe an OTel relation mapping by identifier
sts otel-relation-mapping describe --identifier urn:stackpack:stackpack-name:shared:otel-relation-mapping:database

# output relation mapping to a file
sts otel-relation-mapping describe --identifier urn:stackpack:stackpack-name:shared:otel-relation-mapping:database --file exported.yaml`,
		RunE: deps.CmdRunEWithApi(RunDescribeRelationMappingCommand(args)),
	}
	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the Relation Mapping")
	common.AddFileFlagVar(cmd, &args.FilePath, "Path to the output file")
	return cmd
}

func RunDescribeRelationMappingCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		mapping, resp, err := api.OtelMappingApi.GetOtelRelationMapping(cli.Context, args.Identifier).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		yamlData, err := yaml.Marshal(mapping)
		if err != nil {
			return common.NewExecutionError(fmt.Errorf("failed to marshal mapping to YAML: %v", err))
		}
		data := string(yamlData)

		if args.FilePath != "" {
			if err := util.WriteFile(args.FilePath, []byte(data)); err != nil {
				return common.NewWriteFileError(err, args.FilePath)
			}
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{"filepath": args.FilePath})
			} else {
				cli.Printer.Success(fmt.Sprintf("OTel Relation Mapping exported to: %s", args.FilePath))
			}
			return nil
		} else {
			if cli.IsJson() {
				cli.Printer.PrintJson(map[string]interface{}{
					"data":   mapping,
					"format": "json",
				})
			} else {
				cli.Printer.PrintLn(data)
			}
			return nil
		}
	}
}
