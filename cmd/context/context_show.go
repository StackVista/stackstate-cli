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
		Short: "Show the current context",
		Long:  "Show the current context.",
		RunE:  deps.CmdRunEWithConfig(RunContextShowCommand),
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
