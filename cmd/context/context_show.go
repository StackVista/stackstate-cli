package context

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ShowCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show the current context",
		Long:  "Show the current context.",
		RunE:  deps.CmdRunEWithApi(RunContextShowCommand),
	}

	return cmd
}

func RunContextShowCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	cfg := cli.StsConfig

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"current-context": cfg.GetCurrentContext(),
			"name":            cfg.CurrentContext,
		})
	} else {
		cli.Printer.PrintLn(cfg.CurrentContext)
	}

	return nil
}
