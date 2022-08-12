package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type StackStateClient interface {
	Connect() (*stackstate_api.APIClient, *stackstate_api.ServerInfo, common.CLIError)
}

func NewStackStateClient(ctx context.Context,
	isVerBose bool,
	pr printer.Printer,
	url string,
	apiPath string,
	apiToken string,
	serviceToken string,
	k8sServiceAccountToken string) (StackStateClient, context.Context) {
	apiURL := combineURLandPath(url, apiPath)

	configuration := stackstate_api.NewConfiguration()
	configuration.Servers[0] = stackstate_api.ServerConfiguration{
		URL:         apiURL,
		Description: "",
		Variables:   nil,
	}
	configuration.Debug = isVerBose

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
	newCtx := context.WithValue(
		ctx,
		stackstate_api.ContextAPIKeys,
		auth,
	)

	return StdStackStateClient{
		client:  client,
		Context: newCtx,
		apiURL:  apiURL,
	}, newCtx
}

type StdStackStateClient struct {
	client  *stackstate_api.APIClient
	Context context.Context
	apiURL  string
}

func (c StdStackStateClient) Connect() (*stackstate_api.APIClient, *stackstate_api.ServerInfo, common.CLIError) {
	log.Info().Str("api-url", c.apiURL).Msg("Connecting to StackState")

	serverInfo, resp, err := c.client.ServerApi.ServerInfo(c.Context).Execute()
	if err != nil {
		return nil, &stackstate_api.ServerInfo{}, common.NewConnectError(err, c.apiURL, resp)
	}

	return c.client, serverInfo, nil
}

// Returns the concatenated URL and path, stripping any leading and trailing slashes from the result.
func combineURLandPath(url string, apiPath string) string {
	return strings.Trim(fmt.Sprintf("%s/%s", strings.TrimRight(url, "/"), strings.TrimLeft(apiPath, "/")), "/")
}
