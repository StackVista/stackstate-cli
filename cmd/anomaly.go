package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/cmd/anomaly"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
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
