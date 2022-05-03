package di

import (
	"context"

	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

type StackStateClient interface {
	Connect() (*stackstate_api.APIClient, stackstate_api.ServerInfo, common.CLIError)
}

func NewStackStateClient(client *stackstate_api.APIClient, context context.Context) StackStateClient {
	return StdStackStateClient{
		client:  client,
		Context: context,
	}
}

type StdStackStateClient struct {
	client  *stackstate_api.APIClient
	Context context.Context
}

func (c StdStackStateClient) Connect() (*stackstate_api.APIClient, stackstate_api.ServerInfo, common.CLIError) {
	serverInfo, resp, err := c.client.ServerApi.ServerInfo(c.Context).Execute()
	if err != nil {
		return nil, stackstate_api.ServerInfo{}, common.NewConnectError(err, resp)
	}

	return c.client, serverInfo, nil
}
