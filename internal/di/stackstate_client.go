package di

import (
	"context"

	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

type StackStateClient interface {
	Connect() (*stackstate_client.APIClient, stackstate_client.ServerInfo, common.CLIError)
}

func NewStackStateClient(client *stackstate_client.APIClient, context context.Context) StackStateClient {
	return StdStackStateClient{
		client:  client,
		Context: context,
	}
}

type StdStackStateClient struct {
	client  *stackstate_client.APIClient
	Context context.Context
}

func (c StdStackStateClient) Connect() (*stackstate_client.APIClient, stackstate_client.ServerInfo, common.CLIError) {
	serverInfo, resp, err := c.client.ServerApi.ServerInfo(c.Context).Execute()
	if err != nil {
		return nil, stackstate_client.ServerInfo{}, common.NewConnectError(err, resp)
	}

	return c.client, serverInfo, nil
}
