package settings

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func SettingsListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list --type TYPE",
		Short: "list all settings",
		Long:  "List all settings of a certain type. To list all types run \"sts settings list-types\".",
		RunE:  cli.CmdRunEWithApi(RunSettingsListCommand),
	}
	cmd.Flags().StringP(TypeName, "", "", "name of the setting type to list")
	cmd.MarkFlagRequired(TypeName) //nolint:errcheck

	cmd.Flags().StringP(Namespace, "n", "", "filter by namespace")
	cmd.Flags().StringP(OwnedBy, "w", "", "filter by owner")

	return cmd
}

func RunSettingsListCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	typeName, err := cmd.Flags().GetString(TypeName)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	nameSpace, err := cmd.Flags().GetString(Namespace)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	ownedBy, err := cmd.Flags().GetString(OwnedBy)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	apiClient := api.NodeApi.TypeList(cli.Context, typeName)
	if nameSpace != "" {
		apiClient = apiClient.Namespace(nameSpace)
	}
	if ownedBy != "" {
		apiClient = apiClient.OwnedBy(ownedBy)
	}

	typeList, resp, err := apiClient.Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"settings": typeList,
		})
	} else {
		data := make([][]interface{}, 0)
		for _, v := range typeList {
			lastUpdateTime := time.UnixMilli(v.GetLastUpdateTimestamp())
			data = append(data, []interface{}{
				v.GetTypeName(),
				v.GetId(),
				v.GetIdentifier(),
				v.GetName(),
				v.GetOwnedBy(),
				lastUpdateTime,
			})
		}
		cli.Printer.Table(printer.TableData{
			Header:              []string{"Type", "Id", "Identifier", "Name", "owned by", "last updated"},
			Data:                data,
			MissingTableDataMsg: printer.NotFoundMsg{Types: fmt.Sprintf("settings of type \"%s\"", typeName)},
		})
	}
	return nil
}
