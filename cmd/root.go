package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func RootCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "StackState: topology-powered observability",
	}
	cmd.SetUsageTemplate(cmd.UsageTemplate() +
		"For more information about this CLI visit https://l.stackstate.com/cli\n",
	)

	cmd.AddCommand(VersionCommand(cli))
	cmd.AddCommand(CliCommand(cli))
	if CLIType != "saas" {
		cmd.AddCommand(ScriptCommand(cli))
		cmd.AddCommand(SettingsCommand(cli))
	}
	cmd.AddCommand(MonitorCommand(cli))
	cmd.AddCommand(StackPackCommand(cli))
	common.AddPersistentFlags(cmd)

	return cmd
}
