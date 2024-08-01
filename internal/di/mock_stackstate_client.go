package di

import (
	stackstate_admin_api "github.com/stackvista/stackstate-cli/generated/stackstate_admin_api"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
)

type MockStackStateClient struct {
	apiClient         *stackstate_api.APIClient
	adminApiClient    *stackstate_admin_api.APIClient
	ConnectServerInfo *stackstate_api.ServerInfo
	ConnectError      common.CLIError
	ApiMocks          ApiMocks
}

type ApiMocks struct {
	// Main API:
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
	SubscriptionApi            *stackstate_api.SubscriptionApiMock
	PermissionsApi             *stackstate_api.PermissionsApiMock
	SubjectApi                 *stackstate_api.SubjectApiMock
	TopicApi                   *stackstate_api.TopicApiMock
	AgentRegistrationsApi      *stackstate_api.AgentRegistrationsApiMock
	IngestionApiKeyApi         *stackstate_api.IngestionApiKeyApiMock
	// Admin API:
	RetentionApi *stackstate_admin_api.RetentionApiMock
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
	subscriptionApi := stackstate_api.NewSubscriptionApiMock()
	permissionsApi := stackstate_api.NewPermissionsApiMock()
	subjectApi := stackstate_api.NewSubjectApiMock()
	topicApi := stackstate_api.NewTopicApiMock()
	ingestionApiKeyApi := stackstate_api.NewIngestionApiKeyApiMock()
	retentionApi := stackstate_admin_api.NewRetentionApiMock()
	agentRegistrationsApi := stackstate_api.NewAgentRegistrationsApiMock()

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
		SubscriptionApi:            &subscriptionApi,
		PermissionsApi:             &permissionsApi,
		SubjectApi:                 &subjectApi,
		TopicApi:                   &topicApi,
		IngestionApiKeyApi:         &ingestionApiKeyApi,
		RetentionApi:               &retentionApi,
		AgentRegistrationsApi:      &agentRegistrationsApi,
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
		SubscriptionApi:            apiMocks.SubscriptionApi,
		PermissionsApi:             apiMocks.PermissionsApi,
		SubjectApi:                 apiMocks.SubjectApi,
		IngestionApiKeyApi:         apiMocks.IngestionApiKeyApi,
		TopicApi:                   apiMocks.TopicApi,
		AgentRegistrationsApi:      apiMocks.AgentRegistrationsApi,
	}

	adminApiClient := &stackstate_admin_api.APIClient{
		RetentionApi: apiMocks.RetentionApi,
	}

	// NOTE Used for min version checks.
	mockAPIVersion := stackstate_api.ServerVersion{Major: 5, Minor: 1, Patch: 0} //nolint:gomnd
	mockServerInfo := stackstate_api.NewServerInfo(mockAPIVersion, "SaaS")

	return MockStackStateClient{
		apiClient:         apiClient,
		adminApiClient:    adminApiClient,
		ApiMocks:          apiMocks,
		ConnectServerInfo: mockServerInfo,
		ConnectError:      nil,
	}
}

func (c *MockStackStateClient) Connect() (*stackstate_api.APIClient, *stackstate_api.ServerInfo, common.CLIError) {
	return c.apiClient, c.ConnectServerInfo, c.ConnectError
}

func (c *MockStackStateClient) AdminConnect() (*stackstate_admin_api.APIClient, common.CLIError) {
	return c.adminApiClient, c.ConnectError
}
