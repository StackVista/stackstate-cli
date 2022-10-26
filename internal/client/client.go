package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	stackstate_admin_api "github.com/stackvista/stackstate-cli/generated/stackstate_admin_api"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/printer"
)

type StackStateClient interface {
	Connect() (*stackstate_api.APIClient, *stackstate_api.ServerInfo, common.CLIError)
	AdminConnect() (*stackstate_admin_api.APIClient, common.CLIError)
}

func NewStackStateClient(ctx context.Context,
	isVerbose bool,
	pr printer.Printer,
	version string,
	url string,
	apiPath string,
	adminApiPath string,
	apiToken string,
	serviceToken string,
	k8sServiceAccountToken string) (StackStateClient, context.Context) {
	userAgent := fmt.Sprintf("StackState CLI v%s", version)
	apiURL := combineURLandPath(url, apiPath)
	client, clientAuth := NewApiClient(isVerbose, pr, userAgent, apiURL, apiToken, serviceToken, k8sServiceAccountToken)

	adminApiURL := combineURLandPath(url, adminApiPath)
	adminClient, adminAuth := NewAdminApiClient(isVerbose, pr, userAgent, adminApiURL, apiToken, serviceToken, k8sServiceAccountToken)

	withClient := context.WithValue(
		ctx,
		stackstate_api.ContextAPIKeys,
		clientAuth,
	)
	newCtx := context.WithValue(
		withClient,
		stackstate_admin_api.ContextAPIKeys,
		adminAuth,
	)

	return StdStackStateClient{
		client:      client,
		adminClient: adminClient,
		Context:     newCtx,
		apiURL:      apiURL,
		adminApiURL: adminApiURL,
	}, newCtx
}

//nolint:dupl
func NewApiClient(
	isVerbose bool,
	pr printer.Printer,
	userAgent string,
	apiURL string,
	apiToken string,
	serviceToken string,
	k8sServiceAccountToken string,
) (*stackstate_api.APIClient, map[string]stackstate_api.APIKey) {
	configuration := stackstate_api.NewConfiguration()
	configuration.UserAgent = userAgent
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

	return client, auth
}

//nolint:dupl
func NewAdminApiClient(
	isVerbose bool,
	pr printer.Printer,
	userAgent string,
	apiURL string,
	apiToken string,
	serviceToken string,
	k8sServiceAccountToken string,
) (*stackstate_admin_api.APIClient, map[string]stackstate_admin_api.APIKey) {
	configuration := stackstate_admin_api.NewConfiguration()
	configuration.UserAgent = userAgent
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

	auth := make(map[string]stackstate_admin_api.APIKey)
	if apiToken != "" {
		auth["ApiToken"] = stackstate_admin_api.APIKey{
			Key:    apiToken,
			Prefix: "",
		}
	}
	if serviceToken != "" {
		auth["ServiceToken"] = stackstate_admin_api.APIKey{
			Key:    serviceToken,
			Prefix: "",
		}
	}
	if k8sServiceAccountToken != "" {
		auth["ServiceBearer"] = stackstate_admin_api.APIKey{
			Key:    k8sServiceAccountToken,
			Prefix: "",
		}
	}

	return client, auth
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
	log.Info().Str("admin-api-url", c.adminApiURL).Msg("Connecting to StackState")

	// NOTE Admin API doesn't have any endpoints that could be used to check connectedness.

	return c.adminClient, nil
}

// Returns the concatenated URL and path, stripping any leading and trailing slashes from the result.
func combineURLandPath(url string, apiPath string) string {
	return strings.Trim(fmt.Sprintf("%s/%s", strings.TrimRight(url, "/"), strings.TrimLeft(apiPath, "/")), "/")
}
