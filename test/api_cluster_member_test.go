/*
KubeSphere

Testing ClusterMemberAPIService

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

func Test_openapi_ClusterMemberAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClusterMemberAPIService CreateClusterMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ClusterMemberAPI.CreateClusterMembers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterMemberAPIService DescribeClusterMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clustermember string

		resp, httpRes, err := apiClient.ClusterMemberAPI.DescribeClusterMember(context.Background(), clustermember).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterMemberAPIService ListClusterMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ClusterMemberAPI.ListClusterMembers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterMemberAPIService RemoveClusterMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clustermember string

		resp, httpRes, err := apiClient.ClusterMemberAPI.RemoveClusterMember(context.Background(), clustermember).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterMemberAPIService UpdateClusterMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clustermember string

		resp, httpRes, err := apiClient.ClusterMemberAPI.UpdateClusterMember(context.Background(), clustermember).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
