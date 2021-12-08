package cli

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/conf"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	TestConnectFlagName = "test-connect"
)

func CliSaveConfigCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-config",
		Short: "Save config to file.",
		RunE:  di.CmdRunEWithDeps(cli, RunCliSaveConfig),
	}
	cmd.Flags().BoolP(TestConnectFlagName, "t", false, "Test connection to StackState after the config file has been saved.")

	// same as peristent flags on RootCommand.
	cmd.Flags().String("api-url", "", "StackState API URL.")
	cmd.Flags().String("api-token", "", "StackState API Token.")
	cmd.MarkFlagRequired("api-url")
	cmd.MarkFlagRequired("api-token")

	return cmd
}

func RunCliSaveConfig(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
	testConnect, err := cmd.Flags().GetBool(TestConnectFlagName)
	if err != nil {
		return common.NewCLIError(err)
	}

	filename, err := conf.WriteConf(*cli.Config)
	if err != nil {
		return common.NewCLIError(err)
	}
	cli.Printer.Success("Config saved to: " + filename)

	if testConnect {
		testConect(cli)
	}

	return nil
}
