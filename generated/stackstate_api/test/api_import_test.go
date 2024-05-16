/*
StackState API

Testing ImportAPIService

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

func Test_stackstate_api_ImportAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportAPIService ImportSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ImportAPI.ImportSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
