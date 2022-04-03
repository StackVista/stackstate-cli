package settings

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"time"
)

func SettingsListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list --type TYPE",
		Short: "list all types of settings available",
		RunE:  cli.CmdRunEWithApi(RunSettingsListCommand),
	}
	cmd.Flags().StringP(TypeName, "", "", "example: ComponentType")
	cmd.MarkFlagRequired(TypeName)

	cmd.Flags().StringP(Namespace, "n", "", "name of the namespace")
	cmd.Flags().StringP(OwnedBy, "w", "", "name of the owner")

	return cmd
}

func RunSettingsListCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_client.APIClient, serverInfo stackstate_client.ServerInfo) common.CLIError {
	typeName, err := cmd.Flags().GetString(TypeName)
	if err != nil {
		return common.NewCLIError(err)
	}
	nameSpace, err := cmd.Flags().GetString(Namespace)
	if err != nil {
		return common.NewCLIError(err)
	}
	ownedBy, err := cmd.Flags().GetString(OwnedBy)
	if err != nil {
		return common.NewCLIError(err)
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
	var respData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return common.NewCLIError(err)
	}

	data := make([][]interface{}, 0)
	for _, v := range typeList {
		lastUpdateTime := time.UnixMilli(v.GetLastUpdateTimestamp())
		data = append(data, []interface{}{
			v.GetTypeName(),
			v.GetId(),
			v.GetIdentifier(),
			v.GetName(),
			v.GetDescription(),
			v.GetOwnedBy(),
			lastUpdateTime,
		})
	}
	cli.Printer.Table(
		[]string{"Type", "Id", "Identifier", "Name", "description", "owned by", "last updated"},
		data,
		respData,
	)
	return nil
}
