package cli

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func CliTestCommandCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-connect",
		Short: "Test connection with StackState.",
		RunE:  di.CmdRunEWithDeps(cli, RunTestConnectConfig),
	}
	return cmd
}

func RunTestConnectConfig(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	return testConect(cli)
}
