package servicetoken

import (
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

func ListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List service tokens",
		Long:  "List all service tokens.",
		RunE:  cli.CmdRunEWithApi(RunServiceTokenListCommand),
	}

	return cmd
}

func RunServiceTokenListCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	serviceTokens, resp, err := api.ServiceTokenApi.GetServiceTokens(cli.Context).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	sort.SliceStable(serviceTokens, func(i, j int) bool {
		return serviceTokens[i].Name < serviceTokens[j].Name
	})

	if cli.IsJson() {
		cli.Printer.PrintJson(map[string]interface{}{
			"service-tokens": serviceTokens,
		})
	} else {
		data := make([][]interface{}, 0)
		for _, serviceToken := range serviceTokens {
			sid := *serviceToken.Id
			exp := ""
			if serviceToken.Expiration != nil {
				exp = time.UnixMilli(*serviceToken.Expiration).Format(DateFormat)
			}
			data = append(data, []interface{}{sid, serviceToken.Name, exp, serviceToken.Roles, serviceToken.DedicatedSubject})
		}

		cli.Printer.Table(printer.TableData{
			Header:              []string{"id", "name", "expiration", "roles", "dedicated_Subject"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: "service tokens"},
		})
	}

	return nil
}
