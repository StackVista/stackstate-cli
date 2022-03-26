package settings

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"io/ioutil"
)

const (
	TypeName  = "type"
	Namespace = "namespace"
	OwnedBy   = "owned-by"
)

func SettingsListCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list -t TYPE",
		Short: "list all types of settings available",
		RunE:  di.CmdRunEWithDeps(cli, RunSettingsListCommand),
	}
	cmd.Flags().StringP(TypeName, "t", "", "example: ComponentType")
	cmd.MarkFlagRequired(TypeName)

	cmd.Flags().StringP(Namespace, "n", "", "name of the namespace")
	cmd.Flags().StringP(OwnedBy, "w", "", "name of the owner")

	return cmd
}

func RunSettingsListCommand(cli *di.Deps, cmd *cobra.Command, args []string) common.CLIError {
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

	apiClient := cli.Client.NodeApi.TypeList(cli.Context, typeName)
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
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return common.NewCLIError(err)
	}

	data := make([][]string, 0)
	for _, v := range typeList {
		data = append(data, []string{
			fmt.Sprintf("%d", v.GetId()),
			v.GetTypeName(),
			v.GetName(),
			v.GetDescription(),
			v.GetOwnedBy(),
			fmt.Sprintf("%d", v.GetLastUpdateTimestamp()),
		})
	}
	cli.Printer.Table(
		[]string{"id", "type", "name", "description", "owned by", "last updated"},
		data,
		string(respBytes),
	)
	return nil
}
