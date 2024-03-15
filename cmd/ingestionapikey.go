package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd/ingestionapikey"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func IngestionApiKeyCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingestion-api-key",
		Short: "Manage Ingestion API Keys",
		Long:  "Manage API Keys used by ingestion pipelines, means data (spans, metrics, logs an so on) send by STS Agent, OTel and so on.",
	}

	cmd.AddCommand(ingestionapikey.CreateCommand(deps))
	cmd.AddCommand(ingestionapikey.ListCommand(deps))
	cmd.AddCommand(ingestionapikey.DeleteCommand(deps))
	return cmd
}
