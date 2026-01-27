package stackpack

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func StackpackListParameterCommand(cli *di.Deps) *cobra.Command {
	args := &ListPropertiesArgs{}
	cmd := &cobra.Command{
		Use:   "list-parameters --name NAME",
		Short: "List installation parameters for a StackPack",
		Long:  "List all parameters required to install a StackPack. Use these parameter names with 'sts stackpack install -p key=value'.",
		Example: `# list parameters for the kubernetes StackPack
sts stackpack list-parameters --name kubernetes`,
		RunE: cli.CmdRunEWithApi(RunStackpackListParameterCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the StackPack")
	return cmd
}

func formatParametersTable(stackpack stackstate_api.FullStackPack) [][]interface{} {
	steps := stackpack.GetSteps()

	sort.SliceStable(steps, func(i, j int) bool {
		return *steps[i].Name < *steps[j].Name
	})

	data := make([][]interface{}, 0)
	for _, step := range steps {
		data = append(data, []interface{}{step.Name, step.Display, step.GetValue().Type})
	}

	return data
}

func formatParametersJson(stackpack stackstate_api.FullStackPack) []stackstate_api.StackPackStep {
	steps := stackpack.GetSteps()

	sort.SliceStable(steps, func(i, j int) bool {
		return *steps[i].Name < *steps[j].Name
	})

	return steps
}

func RunStackpackListParameterCommand(args *ListPropertiesArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackPackList, resp, err := api.StackpackApi.StackPackList(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		data := make([][]interface{}, 0)
		steps := make([]stackstate_api.StackPackStep, 0)
		for _, stack := range stackPackList {
			if stack.GetName() != args.Name {
				continue
			}

			data = formatParametersTable(stack)
			steps = formatParametersJson(stack)
		}
		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"parameters": steps,
			})
		} else {
			cli.Printer.Table(printer.TableData{
				Header:              []string{"name", "display name", "type"},
				Data:                data,
				MissingTableDataMsg: printer.NotFoundMsg{Types: fmt.Sprintf("StackPack \"%s\"", args.Name)},
			})
		}
		return nil
	}
}
