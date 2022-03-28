package di

import (
	"context"

	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

type StackStateClient interface {
	Connect() (*stackstate_client.APIClient, stackstate_client.ServerInfo, error)
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

func (c StdStackStateClient) Connect() (*stackstate_client.APIClient, stackstate_client.ServerInfo, error) {
	serverInfo, resp, err := c.client.ServerApi.ServerInfo(c.Context).Execute()
	if err != nil {
		return nil, stackstate_client.ServerInfo{}, common.NewResponseError(err, resp)
	}

	return c.client, serverInfo, nil
}
