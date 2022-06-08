package context

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ShowCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show the current context",
		Long:  "Show the current context.",
		RunE:  deps.CmdRunE(RunContextShowCommand),
	}

	return cmd
}

func RunContextShowCommand(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	cfg := cli.StsConfig

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
