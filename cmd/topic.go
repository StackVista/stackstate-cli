package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/topic"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func TopicCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topic",
		Short: "manage topic",
		Long:  "Manage and details of topic.",
	}
	cmd.AddCommand(topic.TopicListCommand(cli))
	cmd.AddCommand(topic.TopicShowCommand(cli))

	return cmd
}
