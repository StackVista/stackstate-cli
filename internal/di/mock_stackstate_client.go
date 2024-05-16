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
	ApiTokenAPI                *stackstate_api.ApiTokenAPIMock
	EventAPI                   *stackstate_api.EventAPIMock
	HealthSynchronizationAPI   *stackstate_api.HealthSynchronizationAPIMock
	ImportAPI                  *stackstate_api.ImportAPIMock
	ExportAPI                  *stackstate_api.ExportAPIMock
	MonitorAPI                 *stackstate_api.MonitorAPIMock
	NodeAPI                    *stackstate_api.NodeAPIMock
	ScriptingAPI               *stackstate_api.ScriptingAPIMock
	ServiceTokenAPI            *stackstate_api.ServiceTokenAPIMock
	TopologySynchronizationAPI *stackstate_api.TopologySynchronizationAPIMock
	UserProfileAPI             *stackstate_api.UserProfileAPIMock
	ServerAPI                  *stackstate_api.ServerAPIMock
	StackpackAPI               *stackstate_api.StackpackAPIMock
	SubscriptionAPI            *stackstate_api.SubscriptionAPIMock
	PermissionsAPI             *stackstate_api.PermissionsAPIMock
	SubjectAPI                 *stackstate_api.SubjectAPIMock
	TopicAPI                   *stackstate_api.TopicAPIMock
	IngestionApiKeyAPI         *stackstate_api.IngestionApiKeyAPIMock
	// Admin API:
	RetentionAPI *stackstate_admin_api.RetentionAPIMock
	// MISSING MOCK? You have to manually add new mocks here after generating a new API!
}

func NewMockStackStateClient() MockStackStateClient {
	apiTokenApi := stackstate_api.NewApiTokenAPIMock()
	eventApi := stackstate_api.NewEventAPIMock()
	HealthSynchronizationAPI := stackstate_api.NewHealthSynchronizationAPIMock()
	importApi := stackstate_api.NewImportAPIMock()
	exportApi := stackstate_api.NewExportAPIMock()
	monitorApi := stackstate_api.NewMonitorAPIMock()
	nodeApi := stackstate_api.NewNodeAPIMock()
	scriptingApi := stackstate_api.NewScriptingAPIMock()
	serviceTokenApi := stackstate_api.NewServiceTokenAPIMock()
	topologySynchronizationApi := stackstate_api.NewTopologySynchronizationAPIMock()
	userProfileApi := stackstate_api.NewUserProfileAPIMock()
	serverApi := stackstate_api.NewServerAPIMock()
	stackpackApi := stackstate_api.NewStackpackAPIMock()
	subscriptionApi := stackstate_api.NewSubscriptionAPIMock()
	permissionsApi := stackstate_api.NewPermissionsAPIMock()
	subjectApi := stackstate_api.NewSubjectAPIMock()
	topicApi := stackstate_api.NewTopicAPIMock()
	IngestionApiKeyAPI := stackstate_api.NewIngestionApiKeyAPIMock()
	retentionApi := stackstate_admin_api.NewRetentionAPIMock()

	apiMocks := ApiMocks{
		ApiTokenAPI:                &apiTokenApi,
		EventAPI:                   &eventApi,
		HealthSynchronizationAPI:   &HealthSynchronizationAPI,
		ImportAPI:                  &importApi,
		ExportAPI:                  &exportApi,
		MonitorAPI:                 &monitorApi,
		NodeAPI:                    &nodeApi,
		ScriptingAPI:               &scriptingApi,
		ServiceTokenAPI:            &serviceTokenApi,
		TopologySynchronizationAPI: &topologySynchronizationApi,
		UserProfileAPI:             &userProfileApi,
		ServerAPI:                  &serverApi,
		StackpackAPI:               &stackpackApi,
		SubscriptionAPI:            &subscriptionApi,
		PermissionsAPI:             &permissionsApi,
		SubjectAPI:                 &subjectApi,
		TopicAPI:                   &topicApi,
		IngestionApiKeyAPI:         &IngestionApiKeyAPI,
		RetentionAPI:               &retentionApi,
	}

	apiClient := &stackstate_api.APIClient{
		ApiTokenAPI:                apiMocks.ApiTokenAPI,
		EventAPI:                   apiMocks.EventAPI,
		HealthSynchronizationAPI:   apiMocks.HealthSynchronizationAPI,
		ImportAPI:                  apiMocks.ImportAPI,
		ExportAPI:                  apiMocks.ExportAPI,
		MonitorAPI:                 apiMocks.MonitorAPI,
		NodeAPI:                    apiMocks.NodeAPI,
		ServiceTokenAPI:            apiMocks.ServiceTokenAPI,
		ScriptingAPI:               apiMocks.ScriptingAPI,
		TopologySynchronizationAPI: apiMocks.TopologySynchronizationAPI,
		UserProfileAPI:             apiMocks.UserProfileAPI,
		ServerAPI:                  apiMocks.ServerAPI,
		StackpackAPI:               apiMocks.StackpackAPI,
		SubscriptionAPI:            apiMocks.SubscriptionAPI,
		PermissionsAPI:             apiMocks.PermissionsAPI,
		SubjectAPI:                 apiMocks.SubjectAPI,
		IngestionApiKeyAPI:         apiMocks.IngestionApiKeyAPI,
		TopicAPI:                   apiMocks.TopicAPI,
	}

	adminApiClient := &stackstate_admin_api.APIClient{
		RetentionAPI: apiMocks.RetentionAPI,
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
