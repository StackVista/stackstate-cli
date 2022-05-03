package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

func setupClient() (StackStateClient, *stackstate_api.ServerApiMock) {
	serverApi := stackstate_api.NewServerApiMock()
	api := stackstate_api.APIClient{
		ServerApi: &serverApi,
	}
	client := StdStackStateClient{
		client:  &api,
		Context: context.Background(),
		apiURL:  "https://test",
	}
	return client, &serverApi
}

func TestConnectError(t *testing.T) {
	client, serverApi := setupClient()
	err := fmt.Errorf("authentication error")
	resp := http.Response{StatusCode: 401}
	serverApi.ServerInfoResponse.Response = &resp
	serverApi.ServerInfoResponse.Error = err

	_, _, actualErr := client.Connect()

	assert.Equal(t, common.NewConnectError(err, "https://test", &resp), actualErr)
}

func TestConnectSuccess(t *testing.T) {
	client, serverApi := setupClient()
	serverApi.ServerInfoResponse.Result = stackstate_api.ServerInfo{DeploymentMode: "test"}

	_, serverInfo, _ := client.Connect()

	assert.Equal(t, "test", serverInfo.DeploymentMode)
}

func TestCombineURLAndApi(t *testing.T) {
	assert.Equal(t, combineURLandPath("https://bla", "api"), "https://bla/api")
	assert.Equal(t, combineURLandPath("https://bla///", "///api"), "https://bla/api")
	assert.Equal(t, combineURLandPath("https://bla", "/api"), "https://bla/api")
	assert.Equal(t, combineURLandPath("/https://bla", "api/"), "/https://bla/api/")
}
