package cli

import (
	"github.com/spf13/cobra"
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
		RunE: cli.CmdRunE(RunTestConnectConfig),
	}
	return cmd
}

func RunTestConnectConfig(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	return testConect(cli)
}
