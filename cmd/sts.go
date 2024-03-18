package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
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

	cmd.AddCommand(ContextCommand(cli))
	cmd.AddCommand(VersionCommand(cli))
	cmd.AddCommand(ScriptCommand(cli))
	cmd.AddCommand(SettingsCommand(cli))
	cmd.AddCommand(StackPackCommand(cli))
	cmd.AddCommand(MonitorCommand(cli))
	cmd.AddCommand(ServiceTokenCommand(cli))
	cmd.AddCommand(HealthCommand(cli))
	cmd.AddCommand(LicenseCommand(cli))
	cmd.AddCommand(GraphCommand(cli))
	cmd.AddCommand(RbacCommand(cli))
	cmd.AddCommand(TopicCommand(cli))
	cmd.AddCommand(TopologySyncCommand(cli))
	cmd.AddCommand(IngestionApiKeyCommand(cli))

	return cmd
}
