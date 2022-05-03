package cli

import (
	"github.com/spf13/cobra"
	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func CliTestCommandCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-connect",
		Short: "test connection to the configured StackState",
		Long:  "Test connection to the configured StackState or to a different StackState by passing flags.",
		Example: "# test connection to the configured StackState\n" +
			"sts cli test-connect\n" +
			"# test connection to my.stackstate.com \n" +
			"sts cli test-connect --api-token l9x5g14cMcI97IS4785HWgwEpdPr3KJ4 --api-url \"https://my.stackstate.com/api\"",
		RunE: cli.CmdRunEWithApi(RunTestConnectConfig),
	}
	return cmd
}

func RunTestConnectConfig(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"connected":   true,
			"server-info": serverInfo,
		})
	} else {
		PrintConnectionSuccess(cli.Printer, cli.Config.URL, serverInfo)
	}
	return nil
}
