package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/topic"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func TopicCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topic",
		Short: "Inspect Kafka topics for debugging",
		Long:  "Inspect Kafka topics used by SUSE Observability. Useful for debugging data ingestion issues.",
	}

	cmd.AddCommand(topic.ListCommand(deps))
	cmd.AddCommand(topic.DescribeCommand(deps))
	return cmd
}
