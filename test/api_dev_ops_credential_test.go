/*
KubeSphere

Testing DevOpsCredentialAPIService

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

func Test_openapi_DevOpsCredentialAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DevOpsCredentialAPIService GetProjectCredentialUsage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var devops string
		var credential string

		resp, httpRes, err := apiClient.DevOpsCredentialAPI.GetProjectCredentialUsage(context.Background(), devops, credential).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
