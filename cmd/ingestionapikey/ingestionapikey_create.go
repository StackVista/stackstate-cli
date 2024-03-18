package ingestionapikey

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
	Name        string
	Expiration  time.Time
	Description string
}

func CreateCommand(deps *di.Deps) *cobra.Command {
	args := &CreateArgs{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Ingestion Api Key",
		Long:  "Creates a token and then returns it in the response, the token can't be obtained any more after that so store it in the safe space.",
		RunE:  deps.CmdRunEWithApi(RunIngestionApiKeyGenerationCommand(args)),
	}

	common.AddRequiredNameFlagVar(cmd, &args.Name, "Name of the API Key")
	cmd.Flags().TimeVar(&args.Expiration, "expiration", time.Time{}, []string{DateFormat}, "Expiration date of the API Key")
	cmd.Flags().StringVar(&args.Description, "description", "", "Optional description of the API Key")
	return cmd
}

func RunIngestionApiKeyGenerationCommand(args *CreateArgs) di.CmdWithApiFn {
	return func(cmd *cobra.Command, cli *di.Deps, api *stackstate_api.APIClient, serverInfo *stackstate_api.ServerInfo) common.CLIError {
		req := stackstate_api.GenerateIngestionApiKeyRequest{
			Name: args.Name,
		}

		if len(args.Description) > 0 {
			req.Description = &args.Description
		}

		if !args.Expiration.IsZero() {
			m := args.Expiration.UnixMilli()
			req.Expiration = &m
		}

		ingestionApiKeyAPI := api.IngestionApiKeyApi.GenerateIngestionApiKey(cli.Context)

		serviceToken, resp, err := ingestionApiKeyAPI.GenerateIngestionApiKeyRequest(req).Execute()
		if err != nil {
			return common.NewResponseError(err, resp)
		}

		if cli.IsJson() {
			cli.Printer.PrintJson(map[string]interface{}{
				"ingestion-api-key": serviceToken,
			})
		} else {
			cli.Printer.Success(fmt.Sprintf("Ingestion API Key generated: %s\n", color.White.Render(serviceToken.ApiKey)))
		}

		return nil
	}
}
