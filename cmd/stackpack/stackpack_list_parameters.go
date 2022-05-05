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
	cmd := &cobra.Command{
		Use:   "list-parameters",
		Short: "list-parameters",
		Long:  "List available parameters.",
		RunE:  cli.CmdRunEWithApi(RunStackpackListParameterCommand),
	}
	cmd.Flags().String(NameFlag, "", "name of the parameter")
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
	return cmd
}

func RunStackpackListParameterCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo stackstate_api.ServerInfo,
) common.CLIError {
	name, err := cmd.Flags().GetString(NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	stackpackList, resp, err := api.StackpackApi.StackpackList(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	data := make([][]interface{}, 0)
	steps := make([]stackstate_api.StackPackStep, 0)
	for _, stack := range stackpackList {
		if stack.GetName() != name {
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
			MissingTableDataMsg: printer.NotFoundMsg{Types: fmt.Sprintf("StackPack %s does not exist", name)},
		})
	}
	return nil
}
