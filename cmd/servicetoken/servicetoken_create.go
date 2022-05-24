package servicetoken

import (
	"fmt"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	DateFormat = "2006-01-02"
)

type CreateArgs struct {
	Name       string
	Expiration time.Time
	Roles      []string
}

func CreateCommand(deps *di.Deps) *cobra.Command {
	args := &CreateArgs{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a service token",
		Long:  "Create a new service token for authentication with StackState.",
		RunE:  deps.CmdRunEWithApi(RunServiceTokenCreateCommand(args)),
	}

	common.AddRequiredNameFlagVar(cmd, &args.Name, "name of the service token")
	cmd.Flags().TimeVar(&args.Expiration, "expiration", time.Time{}, []string{DateFormat}, "expiration date of the service token")
	cmd.Flags().StringSliceVar(&args.Roles, "roles", []string{}, "roles assigned to the service token")
	cmd.MarkFlagRequired("roles") //nolint:errcheck
	return cmd
}

func RunServiceTokenCreateCommand(args *CreateArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		req := stackstate_api.NewServiceTokenRequest{
			Name:  args.Name,
			Roles: args.Roles,
		}

		if !args.Expiration.IsZero() {
			m := args.Expiration.UnixMilli()
			req.ExpiryDate = &m
		}

		serviceTokenAPI := api.ServiceTokenApi.CreateNewServiceToken(cli.Context)

		serviceToken, resp, err := serviceTokenAPI.NewServiceTokenRequest(req).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"service-token": serviceToken,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Service token created: %s\n", color.White.Render(serviceToken.Token)))
		}

		return nil
	}
}
