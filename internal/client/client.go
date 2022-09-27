package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	stackstate_admin_api "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_admin_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type StackStateClient interface {
	Connect() (*stackstate_api.APIClient, *stackstate_api.ServerInfo, common.CLIError)
	AdminConnect() (*stackstate_admin_api.APIClient, common.CLIError)
}

func NewStackStateClient(ctx context.Context,
	isVerbose bool,
	pr printer.Printer,
	url string,
	apiPath string,
	adminApiPath string,
	apiToken string,
	serviceToken string,
	k8sServiceAccountToken string) (StackStateClient, context.Context) {

	apiURL := combineURLandPath(url, apiPath)
	client := NewApiClient(isVerbose, pr, apiURL)

	adminApiURL := combineURLandPath(url, adminApiPath)
	adminClient := NewAdminApiClient(isVerbose, pr, adminApiURL)

	auth := make(map[string]stackstate_api.APIKey)
	if apiToken != "" {
		auth["ApiToken"] = stackstate_api.APIKey{
			Key:    apiToken,
			Prefix: "",
		}
	}
	if serviceToken != "" {
		auth["ServiceToken"] = stackstate_api.APIKey{
			Key:    serviceToken,
			Prefix: "",
		}
	}
	if k8sServiceAccountToken != "" {
		auth["ServiceBearer"] = stackstate_api.APIKey{
			Key:    k8sServiceAccountToken,
			Prefix: "",
		}
	}
	newCtx := context.WithValue(
		ctx,
		// NOTE The Admin API uses the same naming convention, so this kinda, sorta works out of the box for both clients.
		// NOTE Nonetheless, it's utterly terrible. :shrug:
		stackstate_api.ContextAPIKeys,
		auth,
	)

	return StdStackStateClient{
		client:	     client,
		adminClient: adminClient,
		Context:     newCtx,
		apiURL:      apiURL,
		adminApiURL: adminApiURL,
	}, newCtx
}

func NewApiClient(isVerbose bool, pr printer.Printer, apiURL string) *stackstate_api.APIClient {
	configuration := stackstate_api.NewConfiguration()
	configuration.Servers[0] = stackstate_api.ServerConfiguration{
		URL:         apiURL,
		Description: "",
		Variables:   nil,
	}
	configuration.Debug = isVerbose

	stopSpinner := func() {}
	configuration.OnPreCallAPI = func(r *http.Request) {
		stopSpinner = pr.StartSpinner(printer.AwaitingServer)
	}
	configuration.OnPostCallAPI = func(r *http.Request) {
		stopSpinner()
	}
	client := stackstate_api.NewAPIClient(configuration)

	return client
}

func NewAdminApiClient(isVerbose bool, pr printer.Printer, apiURL string) *stackstate_admin_api.APIClient {
	configuration := stackstate_admin_api.NewConfiguration()
	configuration.Servers[0] = stackstate_admin_api.ServerConfiguration{
		URL:         apiURL,
		Description: "",
		Variables:   nil,
	}
	configuration.Debug = isVerbose

	stopSpinner := func() {}
	configuration.OnPreCallAPI = func(r *http.Request) {
		stopSpinner = pr.StartSpinner(printer.AwaitingServer)
	}
	configuration.OnPostCallAPI = func(r *http.Request) {
		stopSpinner()
	}
	client := stackstate_admin_api.NewAPIClient(configuration)

	return client
}

type StdStackStateClient struct {
	client      *stackstate_api.APIClient
	adminClient *stackstate_admin_api.APIClient
	Context     context.Context
	apiURL      string
	adminApiURL string
}

func (c StdStackStateClient) Connect() (*stackstate_api.APIClient, *stackstate_api.ServerInfo, common.CLIError) {
	log.Info().Str("api-url", c.apiURL).Msg("Connecting to StackState")

	serverInfo, resp, err := c.client.ServerApi.ServerInfo(c.Context).Execute()
	if err != nil {
		return nil, &stackstate_api.ServerInfo{}, common.NewConnectError(err, c.apiURL, resp)
	}

	return c.client, serverInfo, nil
}

func (c StdStackStateClient) AdminConnect() (*stackstate_admin_api.APIClient, common.CLIError) {
	log.Info().Str("admin-api-url", c.apiURL).Msg("Connecting to StackState")

	// NOTE Admin API doesn't have any endpoints that could be used to check connectedness.

	return c.adminClient, nil
}

// Returns the concatenated URL and path, stripping any leading and trailing slashes from the result.
func combineURLandPath(url string, apiPath string) string {
	return strings.Trim(fmt.Sprintf("%s/%s", strings.TrimRight(url, "/"), strings.TrimLeft(apiPath, "/")), "/")
}
