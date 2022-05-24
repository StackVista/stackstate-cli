package cliconfig

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ValidateCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate connection to the configured StackState",
		Long:  "Validate connection to the configured StackState or to a different StackState by passing flags.",
		Example: "# validate connection to the configured StackState\n" +
			"sts cli-config validate\n" +
			"# validate connection to my.stackstate.com \n" +
			"sts cli-config validate --api-token l9x5g14cMcI97IS4785HWgwEpdPr3KJ4 --api-url \"https://my.stackstate.com/api\"",
		RunE: cli.CmdRunEWithApi(RunTestConnectConfig),
	}
	return cmd
}

func RunTestConnectConfig(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_api.APIClient,
	serverInfo *stackstate_api.ServerInfo,
) common.CLIError {
	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"connected":   true,
			"server-info": serverInfo,
		})
	} else {
		PrintConnectionSuccess(cli.Printer, cli.Config.URL, serverInfo)
	}
	return nil
}
