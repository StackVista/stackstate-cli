package di

import (
	stackstate_admin_api "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_admin_api"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

type MockStackStateClient struct {
	apiClient         *stackstate_api.APIClient
	adminApiClient    *stackstate_admin_api.APIClient
	ConnectServerInfo *stackstate_api.ServerInfo
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
	NodeApi                    *stackstate_api.NodeApiMock
	ScriptingApi               *stackstate_api.ScriptingApiMock
	ServiceTokenApi            *stackstate_api.ServiceTokenApiMock
	TopologySynchronizationApi *stackstate_api.TopologySynchronizationApiMock
	UserProfileApi             *stackstate_api.UserProfileApiMock
	ServerApi                  *stackstate_api.ServerApiMock
	StackpackApi               *stackstate_api.StackpackApiMock
	AnomalyFeedbackApi         *stackstate_api.ExportAnomalyApiMock
	SubscriptionApi            *stackstate_api.SubscriptionApiMock
	RetentionApi               *stackstate_admin_api.RetentionApiMock
	// MISSING MOCK? You have to manually add new mocks here after generating a new API!
}

func NewMockStackStateClient() MockStackStateClient {
	apiTokenApi := stackstate_api.NewApiTokenApiMock()
	eventApi := stackstate_api.NewEventApiMock()
	healthSynchronizationApi := stackstate_api.NewHealthSynchronizationApiMock()
	importApi := stackstate_api.NewImportApiMock()
	exportApi := stackstate_api.NewExportApiMock()
	monitorApi := stackstate_api.NewMonitorApiMock()
	nodeApi := stackstate_api.NewNodeApiMock()
	scriptingApi := stackstate_api.NewScriptingApiMock()
	serviceTokenApi := stackstate_api.NewServiceTokenApiMock()
	topologySynchronizationApi := stackstate_api.NewTopologySynchronizationApiMock()
	userProfileApi := stackstate_api.NewUserProfileApiMock()
	serverApi := stackstate_api.NewServerApiMock()
	stackpackApi := stackstate_api.NewStackpackApiMock()
	anomalyFeedbackApi := stackstate_api.NewExportAnomalyApiMock()
	subscriptionApi := stackstate_api.NewSubscriptionApiMock()
	retentionApi := stackstate_admin_api.NewRetentionApiMock()

	apiMocks := ApiMocks{
		ApiTokenApi:                &apiTokenApi,
		EventApi:                   &eventApi,
		HealthSynchronizationApi:   &healthSynchronizationApi,
		ImportApi:                  &importApi,
		ExportApi:                  &exportApi,
		MonitorApi:                 &monitorApi,
		NodeApi:                    &nodeApi,
		ScriptingApi:               &scriptingApi,
		ServiceTokenApi:            &serviceTokenApi,
		TopologySynchronizationApi: &topologySynchronizationApi,
		UserProfileApi:             &userProfileApi,
		ServerApi:                  &serverApi,
		StackpackApi:               &stackpackApi,
		AnomalyFeedbackApi:         &anomalyFeedbackApi,
		SubscriptionApi:            &subscriptionApi,
		RetentionApi:               &retentionApi,
	}

	apiClient := &stackstate_api.APIClient{
		ApiTokenApi:                apiMocks.ApiTokenApi,
		EventApi:                   apiMocks.EventApi,
		HealthSynchronizationApi:   apiMocks.HealthSynchronizationApi,
		ImportApi:                  apiMocks.ImportApi,
		ExportApi:                  apiMocks.ExportApi,
		MonitorApi:                 apiMocks.MonitorApi,
		NodeApi:                    apiMocks.NodeApi,
		ServiceTokenApi:            apiMocks.ServiceTokenApi,
		ScriptingApi:               apiMocks.ScriptingApi,
		TopologySynchronizationApi: apiMocks.TopologySynchronizationApi,
		UserProfileApi:             apiMocks.UserProfileApi,
		ServerApi:                  apiMocks.ServerApi,
		StackpackApi:               apiMocks.StackpackApi,
		ExportAnomalyApi:           apiMocks.AnomalyFeedbackApi,
		SubscriptionApi:            apiMocks.SubscriptionApi,
	}

	adminApiClient := &stackstate_admin_api.APIClient{
		RetentionApi: apiMocks.RetentionApi,
	}

	return MockStackStateClient{
		apiClient:         apiClient,
		adminApiClient:    adminApiClient,
		ApiMocks:          apiMocks,
		ConnectServerInfo: &stackstate_api.ServerInfo{},
		ConnectError:      nil,
	}
}

func (c *MockStackStateClient) Connect() (*stackstate_api.APIClient, *stackstate_api.ServerInfo, common.CLIError) {
	return c.apiClient, c.ConnectServerInfo, c.ConnectError
}

func (c *MockStackStateClient) AdminConnect() (*stackstate_admin_api.APIClient, common.CLIError) {
	return c.adminApiClient, c.ConnectError
}
