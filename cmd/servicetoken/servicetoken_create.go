package servicetoken

import (
	"fmt"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
)

const (
	DateFormat = "2006-01-02"
)

type CreateArgs struct {
	Name             string
	Expiration       time.Time
	Roles            []string
	DedicatedSubject string
}

func CreateCommand(deps *di.Deps) *cobra.Command {
	args := &CreateArgs{}
	cmd := &cobra.Command{
		Use:   "create --name NAME --roles ROLES",
		Short: "Create a new service token",
		Long:  "Create a new service token for API authentication. The token is shown once upon creation and cannot be retrieved later.",
		Example: `# create a token with admin role
sts service-token create --name my-token --roles stackstate-admin

# create a token with expiration date
sts service-token create --name my-token --roles stackstate-power-user --expiration 2025-12-31`,
		RunE: deps.CmdRunEWithApi(RunServiceTokenCreateCommand(args)),
	}

	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the service token")
	cmd.Flags().TimeVar(&args.Expiration, "expiration", time.Time{}, []string{DateFormat}, "Expiration date of the service token")
	cmd.Flags().StringSliceVar(&args.Roles, "roles", []string{}, "Roles assigned to the service token")
	cmd.Flags().StringVar(&args.DedicatedSubject, "dedicatedSubject", "", "Subject solely created for usage with this token. The dedicated subject is cleaned after the token is deleted")
	cmd.MarkFlagRequired("roles") //nolint:errcheck
	return cmd
}

func RunServiceTokenCreateCommand(args *CreateArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		req := stackstate_api.NewServiceTokenRequest{
			Name:             args.Name,
			Roles:            args.Roles,
			DedicatedSubject: &args.DedicatedSubject,
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
