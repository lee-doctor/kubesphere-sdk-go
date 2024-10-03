/*
KubeSphere

Testing DevOpsProjectMemberAPIService

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

func Test_openapi_DevOpsProjectMemberAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DevOpsProjectMemberAPIService CreateNamespaceMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var devops string

		resp, httpRes, err := apiClient.DevOpsProjectMemberAPI.CreateNamespaceMembers(context.Background(), devops).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevOpsProjectMemberAPIService DescribeDevOpsProjectNamespaceMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var devops string
		var member string

		resp, httpRes, err := apiClient.DevOpsProjectMemberAPI.DescribeDevOpsProjectNamespaceMember(context.Background(), devops, member).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevOpsProjectMemberAPIService ListAllNamespaceMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var devops string

		resp, httpRes, err := apiClient.DevOpsProjectMemberAPI.ListAllNamespaceMembers(context.Background(), devops).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevOpsProjectMemberAPIService RemoveDevOpsNamespaceMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var devops string
		var member string

		resp, httpRes, err := apiClient.DevOpsProjectMemberAPI.RemoveDevOpsNamespaceMember(context.Background(), devops, member).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevOpsProjectMemberAPIService UpdateDevOpsNamespaceMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var devops string
		var member string

		resp, httpRes, err := apiClient.DevOpsProjectMemberAPI.UpdateDevOpsNamespaceMember(context.Background(), devops, member).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}