package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func STSCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sts",
		Short: "SUSE Observability: topology-powered observability", // never actually visible
		Long:  "SUSE Observability: topology-powered observability.",
	}
	cmd.SetUsageTemplate(cmd.UsageTemplate() +
		"For more information about this CLI visit https://documentation.suse.com/cloudnative/suse-observability/latest/en/setup/cli/cli-sts.html\n",
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
	cmd.AddCommand(AgentCommand(cli))
	cmd.AddCommand(UserSessionCommand(cli))
	cmd.AddCommand(DashboardCommand(cli))

	// Experimental commands for otel mapping
	if os.Getenv("STS_EXPERIMENTAL_OTEL_MAPPING") != "" {
		cmd.AddCommand(OtelComponentMappingCommand(cli))
		cmd.AddCommand(OtelRelationtMappingCommand(cli))
	}

	return cmd
}
