package stackpack

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func StackpackListParameterCommand(cli *di.Deps) *cobra.Command {
	args := &ListArgs{}
	cmd := &cobra.Command{
		Use:   "list-parameters",
		Short: "list-parameters",
		Long:  "List all parameters required for a StackPack's installation.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListParameterCommand(args)),
	}
	common.AddRequiredNameFlagVar(cmd, &args.Name, "name of the StackPack")
	return cmd
}

func RunStackpackListParameterCommand(args *ListArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		data := make([][]interface{}, 0)
		steps := make([]stackstate_api.StackPackStep, 0)
		for _, stack := range stackpackList {
			if stack.GetName() != args.Name {
				continue
			}
			for _, step := range stack.GetSteps() {
				data = append(data, []interface{}{step.Name, step.Display, step.GetValue().Type})
				steps = append(steps, step)
			}
		}
		if cli.IsJson {
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
