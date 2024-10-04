/*
KubeSphere

Testing DevOpsWebhookAPIService

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

func Test_openapi_DevOpsWebhookAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DevOpsWebhookAPIService GetNotifyCommit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DevOpsWebhookAPI.GetNotifyCommit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevOpsWebhookAPIService GithubWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DevOpsWebhookAPI.GithubWebhook(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevOpsWebhookAPIService PostNotifyCommit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DevOpsWebhookAPI.PostNotifyCommit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
