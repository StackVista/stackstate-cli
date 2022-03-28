package di

import "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"

type MockStackStateClient struct {
	apiClient         *stackstate_client.APIClient
	ConnectServerInfo ServerInfo
	ConnectError      error
	ApiMocks          ApiMocks
}

type ApiMocks struct {
	ApiTokenApi                *stackstate_client.ApiTokenApiMock
	EventApi                   *stackstate_client.EventApiMock
	HealthSynchronizationApi   *stackstate_client.HealthSynchronizationApiMock
	ImportApi                  *stackstate_client.ImportApiMock
	MonitorApi                 *stackstate_client.MonitorApiMock
	MonitorUrnApi              *stackstate_client.MonitorUrnApiMock
	NodeApi                    *stackstate_client.NodeApiMock
	ScriptingApi               *stackstate_client.ScriptingApiMock
	TopologySynchronizationApi *stackstate_client.TopologySynchronizationApiMock
	UserProfileApi             *stackstate_client.UserProfileApiMock
}

func NewMockStackStateClient() MockStackStateClient {
	apiTokenApi := stackstate_client.NewApiTokenApiMock()
	eventApi := stackstate_client.NewEventApiMock()
	healthSynchronizationApi := stackstate_client.NewHealthSynchronizationApiMock()
	importApi := stackstate_client.NewImportApiMock()
	monitorApi := stackstate_client.NewMonitorApiMock()
	monitorApiUrn := stackstate_client.NewMonitorUrnApiMock()
	nodeApi := stackstate_client.NewNodeApiMock()
	scriptingApi := stackstate_client.NewScriptingApiMock()
	topologySynchronizationApi := stackstate_client.NewTopologySynchronizationApiMock()
	userProfileApi := stackstate_client.NewUserProfileApiMock()
	apiMocks := ApiMocks{
		ApiTokenApi:                &apiTokenApi,
		EventApi:                   &eventApi,
		HealthSynchronizationApi:   &healthSynchronizationApi,
		ImportApi:                  &importApi,
		MonitorApi:                 &monitorApi,
		MonitorUrnApi:              &monitorApiUrn,
		NodeApi:                    &nodeApi,
		ScriptingApi:               &scriptingApi,
		TopologySynchronizationApi: &topologySynchronizationApi,
		UserProfileApi:             &userProfileApi,
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
