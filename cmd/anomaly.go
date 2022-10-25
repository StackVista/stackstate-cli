package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/anomaly"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func AnomalyCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "anomaly",
		Short: "Inspect anomalies",
		Long:  "Inspect anomalies on metric streams.",
	}
	cmd.AddCommand(anomaly.AnomalyCollectFeedback(cli))

	return cmd
}
