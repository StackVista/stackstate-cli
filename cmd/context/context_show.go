package context

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/config"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func ShowCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display the currently active context",
		Long:  "Display the name of the currently active connection context. Use 'sts context list' to see all available contexts.",
		Example: `# show the current context name
sts context show

# show current context details in JSON format
sts context show -o json`,
		RunE: deps.CmdRunEWithConfig(RunContextShowCommand),
	}

	return cmd
}

func RunContextShowCommand(cli *di.Deps, cmd *cobra.Command, cfg *config.Config) common.CLIError {
	ctx, err := cfg.GetContext(cfg.CurrentContext)
	if err != nil {
		return common.NewNotFoundError(err)
	}

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"current-context": ctx.Context,
			"name":            cfg.CurrentContext,
		})
	} else {
		cli.Printer.PrintLn(cfg.CurrentContext)
	}

	return nil
}
