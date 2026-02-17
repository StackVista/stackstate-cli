package otelcomponentmapping

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

func OtelComponentMappingDescribeCommand(deps *di.Deps) *cobra.Command {
	args := &DescribeArgs{}
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe an Otel Component Mapping",
		Long:  "Describe an Otel Component Mapping by identifier (URN). Optionally write to output file.",
		Example: `# describe an OTel component mapping by identifier
sts otel-component-mapping describe --identifier urn:stackpack:stackpack-name:shared:otel-component-mapping:service

# output mapping to a file
sts otel-component-mapping describe --identifier urn:stackpack:stackpack-name:shared:otel-component-mapping:service --file exported.yaml`,
		RunE: deps.CmdRunEWithApi(RunDescribeComponentMappingCommand(args)),
	}

	common.AddRequiredIdentifierFlagVar(cmd, &args.Identifier, "Identifier (URN) of the Component Mapping")
	common.AddFileFlagVar(cmd, &args.FilePath, "Path to the output file")

	return cmd
}

func RunDescribeComponentMappingCommand(args *DescribeArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		mapping, resp, err := api.OtelMappingApi.GetOtelComponentMapping(cli.Context, args.Identifier).Execute()
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
				cli.Printer.Success(fmt.Sprintf("OTel Component Mapping exported to: %s", args.FilePath))
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
