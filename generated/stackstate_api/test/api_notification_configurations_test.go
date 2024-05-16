/*
StackState API

Testing NotificationConfigurationsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package stackstate_api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "gitlab.com/stackvista/stackstate-cli2"
)

func Test_stackstate_api_NotificationConfigurationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NotificationConfigurationsAPIService CreateNotificationConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NotificationConfigurationsAPI.CreateNotificationConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationConfigurationsAPIService DeleteNotificationConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var notificationConfigurationIdOrUrn string

		httpRes, err := apiClient.NotificationConfigurationsAPI.DeleteNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationConfigurationsAPIService GetNotificationConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var notificationConfigurationIdOrUrn string

		resp, httpRes, err := apiClient.NotificationConfigurationsAPI.GetNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationConfigurationsAPIService GetNotificationConfigurationChannels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var notificationConfigurationIdOrUrn string

		resp, httpRes, err := apiClient.NotificationConfigurationsAPI.GetNotificationConfigurationChannels(context.Background(), notificationConfigurationIdOrUrn).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationConfigurationsAPIService GetNotificationConfigurations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NotificationConfigurationsAPI.GetNotificationConfigurations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationConfigurationsAPIService UpdateNotificationConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var notificationConfigurationIdOrUrn string

		resp, httpRes, err := apiClient.NotificationConfigurationsAPI.UpdateNotificationConfiguration(context.Background(), notificationConfigurationIdOrUrn).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
