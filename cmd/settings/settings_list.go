package settings

import (
	"fmt"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type ListArgs struct {
	TypeName  string
	Namespace string
	OwnedBy   string
}

func SettingsListCommand(cli *di.Deps) *cobra.Command {
	args := &ListArgs{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all settings",
		Long:  "List all settings of a certain type. To list all types run \"sts settings list-types\".",
		RunE:  cli.CmdRunEWithApi(RunSettingsListCommand(args)),
	}
	cmd.Flags().StringVar(&args.TypeName, TypeNameFlag, "", "Name of the setting type to list")
	cmd.MarkFlagRequired(TypeNameFlag) //nolint:errcheck

	cmd.Flags().StringVar(&args.Namespace, Namespace, "", "Filter by namespace")
	cmd.Flags().StringVar(&args.OwnedBy, OwnedByFlag, "", "Filter by owner")

	return cmd
}

func RunSettingsListCommand(args *ListArgs) di.CmdWithApiFn {
	return func(
		cmd *cobra.Command,
		cli *di.Deps,
		api *stackstate_api.APIClient,
		serverInfo *stackstate_api.ServerInfo,
	) common.CLIError {
		apiClient := api.NodeApi.TypeList(cli.Context, args.TypeName)
		if args.Namespace != "" {
			apiClient = apiClient.Namespace(args.Namespace)
		}
		if args.OwnedBy != "" {
			apiClient = apiClient.OwnedBy(args.OwnedBy)
		}

		typeList, resp, err := apiClient.Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		sort.SliceStable(typeList, func(i, j int) bool {
			return *typeList[i].Name < *typeList[j].Name
		})

		if cli.IsJson() {
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
				MissingTableDataMsg: printer.NotFoundMsg{Types: fmt.Sprintf("settings of type \"%s\"", args.TypeName)},
			})
		}
		return nil
	}
}
