package servicetoken

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/time_flags"
)

func CreateCommand(deps *di.Deps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a service token",
		Long:  "Create a new service token for authentication with StackState.",
		RunE:  deps.CmdRunEWithApi(RunServiceTokenCreateCommand),
	}

	common.AddRequiredNameFlag(cmd, "name of the service token")
	cmd.Flags().String("expiration", "", "expiration date of the service token")
	cmd.Flags().StringSlice("roles", []string{}, "roles assigned to the service token")
	cmd.MarkFlagRequired("roles") //nolint:errcheck
	return cmd
}

func RunServiceTokenCreateCommand(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
	name, err := cmd.Flags().GetString(common.NameFlag)
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	expiration, cliErr := time_flags.GetDateFlag(cmd, "expiration")
	if cliErr != nil {
		return cliErr
	}

	roles, err := cmd.Flags().GetStringSlice("roles")
	if err != nil {
		return common.NewCLIArgParseError(err)
	}

	req := stackstate_api.NewServiceTokenRequest{
		Name:  name,
		Roles: roles,
	}

	if expiration != nil {
		m := expiration.UnixMilli()
		req.ExpiryDate = &m
	}

	serviceTokenAPI := api.ServiceTokenApi.CreateNewServiceToken(cli.Context)

	serviceToken, resp, err := serviceTokenAPI.NewServiceTokenRequest(req).Execute()
	if err != nil {
		return common.NewResponseError(err, resp)
	}

	if cli.IsJson {
		cli.Printer.PrintJson(map[string]interface{}{
			"service-token": serviceToken,
		})
	} else {
		cli.Printer.Success(fmt.Sprintf("Service token created: %s\n", color.White.Render(serviceToken.Token)))
	}

	return nil
}
