package context

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all configured connection contexts",
		Long:  "List all connection contexts saved in the CLI configuration file, showing the context name and server URL for each.",
		Example: `# list all contexts
sts context list

# list contexts in JSON format
sts context list -o json`,
		RunE: deps.CmdRunEWithConfig(RunContextListCommand),
	}

	return cmd
}

func RunContextListCommand(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
	ctxs := cfg.Contexts

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
