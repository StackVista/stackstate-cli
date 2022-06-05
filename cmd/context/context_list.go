package context

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list available contexts",
		Long:  "List available contexts in the config file.",
		RunE:  deps.CmdRunEWithApi(RunContextListCommand),
	}

	return cmd
}

func RunContextListCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	ctxs := cli.StsConfig.Contexts

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"contexts": ctxs,
		})
	} else {
		data := make([][]interface{}, 0)
		for _, ctx := range ctxs {
			data = append(data, []interface{}{ctx.Name, ctx.Context.URL})
		}

		cli.Printer.Table(printer.TableData{
			Header:              []string{"name", "url"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "contexts"},
		})
	}

	return nil
}
