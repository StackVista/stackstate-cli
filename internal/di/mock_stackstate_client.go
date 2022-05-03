package di

import (
	stackstate_client "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

type MockStackStateClient struct {
	apiClient         *stackstate_client.APIClient
	ConnectServerInfo stackstate_client.ServerInfo
	ConnectError      common.CLIError
	ApiMocks          ApiMocks
}

type ApiMocks struct {
	ApiTokenApi                *stackstate_client.ApiTokenApiMock
	EventApi                   *stackstate_client.EventApiMock
	HealthSynchronizationApi   *stackstate_client.HealthSynchronizationApiMock
	ImportApi                  *stackstate_client.ImportApiMock
	ExportApi                  *stackstate_client.ExportApiMock
	MonitorApi                 *stackstate_client.MonitorApiMock
	MonitorUrnApi              *stackstate_client.MonitorUrnApiMock
	NodeApi                    *stackstate_client.NodeApiMock
	ScriptingApi               *stackstate_client.ScriptingApiMock
	TopologySynchronizationApi *stackstate_client.TopologySynchronizationApiMock
	UserProfileApi             *stackstate_client.UserProfileApiMock
	ServerApi                  *stackstate_client.ServerApiMock
	StackpackApi               *stackstate_client.StackpackApiMock
	// MISSING MOCK? You have to manually add new mocks here after generating a new API!
}

func NewMockStackStateClient() MockStackStateClient {
	apiTokenApi := stackstate_client.NewApiTokenApiMock()
	eventApi := stackstate_client.NewEventApiMock()
	healthSynchronizationApi := stackstate_client.NewHealthSynchronizationApiMock()
	importApi := stackstate_client.NewImportApiMock()
	exportApi := stackstate_client.NewExportApiMock()
	monitorApi := stackstate_client.NewMonitorApiMock()
	monitorApiUrn := stackstate_client.NewMonitorUrnApiMock()
	nodeApi := stackstate_client.NewNodeApiMock()
	scriptingApi := stackstate_client.NewScriptingApiMock()
	topologySynchronizationApi := stackstate_client.NewTopologySynchronizationApiMock()
	userProfileApi := stackstate_client.NewUserProfileApiMock()
	serverApi := stackstate_client.NewServerApiMock()
	stackpackApi := stackstate_client.NewStackpackApiMock()
	apiMocks := ApiMocks{
		ApiTokenApi:                &apiTokenApi,
		EventApi:                   &eventApi,
		HealthSynchronizationApi:   &healthSynchronizationApi,
		ImportApi:                  &importApi,
		ExportApi:                  &exportApi,
		MonitorApi:                 &monitorApi,
		MonitorUrnApi:              &monitorApiUrn,
		NodeApi:                    &nodeApi,
		ScriptingApi:               &scriptingApi,
		TopologySynchronizationApi: &topologySynchronizationApi,
		UserProfileApi:             &userProfileApi,
		ServerApi:                  &serverApi,
		StackpackApi:               &stackpackApi,
	}

	apiClient := &stackstate_client.APIClient{
		ApiTokenApi:                apiMocks.ApiTokenApi,
		EventApi:                   apiMocks.EventApi,
		HealthSynchronizationApi:   apiMocks.HealthSynchronizationApi,
		ImportApi:                  apiMocks.ImportApi,
		ExportApi:                  apiMocks.ExportApi,
		MonitorApi:                 apiMocks.MonitorApi,
		MonitorUrnApi:              apiMocks.MonitorUrnApi,
		NodeApi:                    apiMocks.NodeApi,
		ScriptingApi:               apiMocks.ScriptingApi,
		TopologySynchronizationApi: apiMocks.TopologySynchronizationApi,
		UserProfileApi:             apiMocks.UserProfileApi,
		ServerApi:                  apiMocks.ServerApi,
		StackpackApi:               apiMocks.StackpackApi,
	}

	return MockStackStateClient{
		apiClient:         apiClient,
		ApiMocks:          apiMocks,
		ConnectServerInfo: stackstate_client.ServerInfo{},
		ConnectError:      nil,
	}
}

func (c *MockStackStateClient) Connect() (*stackstate_client.APIClient, stackstate_client.ServerInfo, common.CLIError) {
	return c.apiClient, c.ConnectServerInfo, c.ConnectError
}
