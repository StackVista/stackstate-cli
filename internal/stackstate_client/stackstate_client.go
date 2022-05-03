package stackstate_client

import (
	"context"
	"net/http"

	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

type StackStateClient interface {
	Connect() (*stackstate_api.APIClient, stackstate_api.ServerInfo, common.CLIError)
}

func NewStackStateClient(ctx context.Context, isVerBose bool, pr printer.Printer, URL string, apiPath string, apiToken string) (StackStateClient, context.Context) {
	apiURL := URL + apiPath

	configuration := stackstate_api.NewConfiguration()
	configuration.Servers[0] = stackstate_api.ServerConfiguration{
		URL:         URL + apiPath,
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
	auth["ApiToken"] = stackstate_api.APIKey{
		Key:    apiToken,
		Prefix: "",
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

func (c StdStackStateClient) Connect() (*stackstate_api.APIClient, stackstate_api.ServerInfo, common.CLIError) {
	serverInfo, resp, err := c.client.ServerApi.ServerInfo(c.Context).Execute()
	if err != nil {
		return nil, stackstate_api.ServerInfo{}, common.NewConnectError(err, c.apiURL, resp)
	}

	return c.client, serverInfo, nil
}
