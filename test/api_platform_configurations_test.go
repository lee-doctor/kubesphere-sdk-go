/*
KS API

Testing PlatformConfigurationsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_PlatformConfigurationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PlatformConfigurationsAPIService ComponentConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PlatformConfigurationsAPI.ComponentConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlatformConfigurationsAPIService GetThemeConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PlatformConfigurationsAPI.GetThemeConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlatformConfigurationsAPIService OauthConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PlatformConfigurationsAPI.OauthConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlatformConfigurationsAPIService UpdateThemeConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PlatformConfigurationsAPI.UpdateThemeConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
