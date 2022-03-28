package di

import "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"

type MockStackStateClient struct {
	apiClient         *stackstate_client.APIClient
	ConnectServerInfo ServerInfo
	ConnectError      error
	ApiMocks          ApiMocks
}

type ApiMocks struct {
	ApiTokenApi                stackstate_client.ApiTokenApi
	EventApi                   stackstate_client.EventApi
	HealthSynchronizationApi   stackstate_client.HealthSynchronizationApi
	ImportApi                  stackstate_client.ImportApi
	MonitorApi                 stackstate_client.MonitorApi
	MonitorUrnApi              stackstate_client.MonitorUrnApi
	NodeApi                    stackstate_client.NodeApi
	ScriptingApi               stackstate_client.ScriptingApi
	TopologySynchronizationApi stackstate_client.TopologySynchronizationApi
	UserProfileApi             stackstate_client.UserProfileApi
}

func NewMockStackStateClient() MockStackStateClient {
	apiMocks := ApiMocks{
		ApiTokenApi:                stackstate_client.NewApiTokenApiMock(),
		EventApi:                   stackstate_client.NewEventApiMock(),
		HealthSynchronizationApi:   stackstate_client.NewHealthSynchronizationApiMock(),
		ImportApi:                  stackstate_client.NewImportApiMock(),
		MonitorApi:                 stackstate_client.NewMonitorApiMock(),
		MonitorUrnApi:              stackstate_client.NewMonitorUrnApiMock(),
		NodeApi:                    stackstate_client.NewNodeApiMock(),
		ScriptingApi:               stackstate_client.NewScriptingApiMock(),
		TopologySynchronizationApi: stackstate_client.NewTopologySynchronizationApiMock(),
		UserProfileApi:             stackstate_client.NewUserProfileApiMock(),
	}

	apiClient := &stackstate_client.APIClient{
		ApiTokenApi:                apiMocks.ApiTokenApi,
		EventApi:                   apiMocks.EventApi,
		HealthSynchronizationApi:   apiMocks.HealthSynchronizationApi,
		ImportApi:                  apiMocks.ImportApi,
		MonitorApi:                 apiMocks.MonitorApi,
		MonitorUrnApi:              apiMocks.MonitorUrnApi,
		NodeApi:                    apiMocks.NodeApi,
		ScriptingApi:               apiMocks.ScriptingApi,
		TopologySynchronizationApi: apiMocks.TopologySynchronizationApi,
		UserProfileApi:             apiMocks.UserProfileApi,
	}

	return MockStackStateClient{
		apiClient:         apiClient,
		ApiMocks:          apiMocks,
		ConnectServerInfo: ServerInfo{},
		ConnectError:      nil,
	}
}

func (c MockStackStateClient) Connect() (*stackstate_client.APIClient, ServerInfo, error) {
	return c.apiClient, c.ConnectServerInfo, c.ConnectError
}
