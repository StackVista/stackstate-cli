package ingestionapikey

import (
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func ListCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Ingestion Api Keys",
		Long:  "Returns only metadata without a key itself.",
		RunE:  deps.CmdRunEWithApi(RunIngestionApiKeyListCommand),
	}

	return cmd
}

func RunIngestionApiKeyListCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	ingestionApiKeys, resp, err := api.IngestionApiKeyAPI.GetIngestionApiKeys(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	sort.SliceStable(ingestionApiKeys, func(i, j int) bool {
		return ingestionApiKeys[i].Name < ingestionApiKeys[j].Name
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"ingestion-api-keys": ingestionApiKeys,
		})
	} else {
		data := make([][]interface{}, 0)
		for _, ingestionApiKey := range ingestionApiKeys {
			sid := ingestionApiKey.Id
			exp := ""
			if ingestionApiKey.Expiration != nil {
				exp = time.UnixMilli(*ingestionApiKey.Expiration).Format(DateFormat)
			}
			data = append(data, []interface{}{sid, ingestionApiKey.Name, exp, ingestionApiKey.Description})
		}

		cli.Printer.Table(printer.TableData{
			Header:              []string{"id", "name", "expiration", "description"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "ingestion api keys"},
		})
	}

	return nil
}
