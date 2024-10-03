/*
KubeSphere

Testing WorkspaceMetricsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_WorkspaceMetricsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkspaceMetricsAPIService HandleAllWorkspaceMetricsQuery", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkspaceMetricsAPI.HandleAllWorkspaceMetricsQuery(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspaceMetricsAPIService HandleWorkspaceMetricsQuery", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string

		resp, httpRes, err := apiClient.WorkspaceMetricsAPI.HandleWorkspaceMetricsQuery(context.Background(), workspace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}