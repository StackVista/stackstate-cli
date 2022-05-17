package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func STSCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "StackState: topology-powered observability", // never actually visible
		Long:  "StackState: topology-powered observability.",
	}
	cmd.SetUsageTemplate(cmd.UsageTemplate() +
		"For more information about this CLI visit https://l.stackstate.com/cli\n",
	)

	cmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return common.NewCLIArgParseError(err)
	})

	cmd.AddCommand(VersionCommand(cli))
	cmd.AddCommand(CliConfigCommand(cli))
	if cli.CLIType != "saas" {
		cmd.AddCommand(ScriptCommand(cli))
		cmd.AddCommand(SettingsCommand(cli))
		cmd.AddCommand(StackPackCommand(cli))
	}
	cmd.AddCommand(MonitorCommand(cli))
	cmd.AddCommand(AnomalyCommand(cli))

	cmd.AddCommand(ServiceTokenCommand(cli))
	return cmd
}
