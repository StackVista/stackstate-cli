package stackpack

import (
	"errors"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

func StackpackInstallCommand(cli *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install --name NAME",
		Short: "install a StackPack",
		Long: "Install an instance of a StackPack. " +
			"Be aware that each StackPack has a different set of parameters." +
			"Run \"sts stackpack list-parameters --name NAME\" to list all of them.",
		Example: "# install an example StackPack with a full name and URL parameter\n" +
			"sts stackpack install --name example -p \"full_name=First Last\" -p URL=https://stackstate.com",
		RunE: cli.CmdRunEWithApi(RunStackpackInstallCommand),
	}
	cmd.Flags().String(NameFlag, "", "name of the StackPack")
	cmd.Flags().StringSliceP(ParameterFlag, "p", nil, "list of parameters of the form \"key=value\"")
	cmd.MarkFlagRequired(NameFlag) //nolint:errcheck
	common.AddJsonFlag(cmd)
	return cmd
}

func RunStackpackInstallCommand(
	cmd *cobra.Command,
	cli *di.Deps,
	api *stackstate_client.APIClient,
	serverInfo stackstate_client.ServerInfo,
) common.CLIError {
	json, err := cmd.Flags().GetBool(common.JsonFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	name, err := cmd.Flags().GetString(NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	params, err := cmd.Flags().GetStringSlice(ParameterFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}
	parameters := make(map[string]string)
	for _, v := range params {
		key, value, err := parseParameter(v)
		if err != nil {
			return common.NewCLIArgParseError(err)
		}
		parameters[key] = value
	}

	instance, resp, err := api.StackpackApi.ProvisionDetails(cli.Context, name).RequestBody(parameters).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if json {
		cli.Printer.PrintJson(instance)
	} else {
		lastUpdateTime := time.UnixMilli(instance.GetLastUpdateTimestamp())

		cli.Printer.Success("StackPack instance installed")
		cli.Printer.Table(
			printer.TableData{
				Header:              []string{"id", "name", "status", "version", "last updated"},
				Data:                [][]interface{}{{instance.Id, instance.Name, instance.Status, instance.StackPackVersion, lastUpdateTime}},
				MissingTableDataMsg: printer.NotFoundMsg{Types: "provision details of " + name},
			},
		)
	}

	return nil
}

func parseParameter(input string) (string, string, error) {
	idx := strings.Index(input, "=")
	if idx <= 0 || idx == len(input)-1 {
		return "", "", errors.New("expected parameter format key=value")
	}
	key := input[0:idx]
	value := input[idx+1:]
	return key, value, nil
}
