package di

import (
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

type MockStackStateClient struct {
	apiClient         *stackstate_api.APIClient
	ConnectServerInfo stackstate_api.ServerInfo
	ConnectError      common.CLIError
	ApiMocks          ApiMocks
}

type ApiMocks struct {
	ApiTokenApi                *stackstate_api.ApiTokenApiMock
	EventApi                   *stackstate_api.EventApiMock
	HealthSynchronizationApi   *stackstate_api.HealthSynchronizationApiMock
	ImportApi                  *stackstate_api.ImportApiMock
	ExportApi                  *stackstate_api.ExportApiMock
	MonitorApi                 *stackstate_api.MonitorApiMock
	MonitorUrnApi              *stackstate_api.MonitorUrnApiMock
	NodeApi                    *stackstate_api.NodeApiMock
	ScriptingApi               *stackstate_api.ScriptingApiMock
	TopologySynchronizationApi *stackstate_api.TopologySynchronizationApiMock
	UserProfileApi             *stackstate_api.UserProfileApiMock
	ServerApi                  *stackstate_api.ServerApiMock
	StackpackApi               *stackstate_api.StackpackApiMock
	// MISSING MOCK? You have to manually add new mocks here after generating a new API!
}

func NewMockStackStateClient() MockStackStateClient {
	apiTokenApi := stackstate_api.NewApiTokenApiMock()
	eventApi := stackstate_api.NewEventApiMock()
	healthSynchronizationApi := stackstate_api.NewHealthSynchronizationApiMock()
	importApi := stackstate_api.NewImportApiMock()
	exportApi := stackstate_api.NewExportApiMock()
	monitorApi := stackstate_api.NewMonitorApiMock()
	monitorApiUrn := stackstate_api.NewMonitorUrnApiMock()
	nodeApi := stackstate_api.NewNodeApiMock()
	scriptingApi := stackstate_api.NewScriptingApiMock()
	topologySynchronizationApi := stackstate_api.NewTopologySynchronizationApiMock()
	userProfileApi := stackstate_api.NewUserProfileApiMock()
	serverApi := stackstate_api.NewServerApiMock()
	stackpackApi := stackstate_api.NewStackpackApiMock()
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

	apiClient := &stackstate_api.APIClient{
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
		ConnectServerInfo: stackstate_api.ServerInfo{},
		ConnectError:      nil,
	}
}

func (c *MockStackStateClient) Connect() (*stackstate_api.APIClient, stackstate_api.ServerInfo, common.CLIError) {
	return c.apiClient, c.ConnectServerInfo, c.ConnectError
}
