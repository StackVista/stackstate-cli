package cliconfig

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func DescribeCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "describe the config of the CLI",
		Long:  "Describes the config of the CLI as it is retrieved from the config file, environment variables and/or flags in YAML or JSON format.",
		RunE:  cli.CmdRunE(RunDescribe),
	}

	return cmd
}

func RunDescribe(cli *di.Deps, cmd *cobra.Command) common.CLIError {
	err := cli.LoadConfig(cmd)
	if err != nil {
		return err
	}

	if cli.IsJson {
		cli.Printer.PrintJson(viper.AllSettings())
	} else {
		cli.Printer.PrintStruct(viper.AllSettings())
	}
	return nil
}
