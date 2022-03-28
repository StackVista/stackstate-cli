package di

import (
	"context"

	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
)

type StackStateClient interface {
	Connect() (*stackstate_client.APIClient, ServerInfo, error)
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

type ServerInfo struct{}

func (c StdStackStateClient) Connect() (*stackstate_client.APIClient, ServerInfo, error) {
	_, resp, err := c.client.UserProfileApi.GetCurrentUserProfile(c.Context).Execute()
	if err != nil {
		return nil, ServerInfo{}, common.NewResponseError(err, resp)
	}

	return c.client, ServerInfo{}, nil
}
