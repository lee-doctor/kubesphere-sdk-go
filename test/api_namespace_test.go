/*
KubeSphere

Testing NamespaceAPIService

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

func Test_openapi_NamespaceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamespaceAPIService CreateNamespace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string

		resp, httpRes, err := apiClient.NamespaceAPI.CreateNamespace(context.Background(), workspace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService DeleteNamespace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string
		var namespace string

		resp, httpRes, err := apiClient.NamespaceAPI.DeleteNamespace(context.Background(), workspace, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService DescribeNamespace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string
		var namespace string

		resp, httpRes, err := apiClient.NamespaceAPI.DescribeNamespace(context.Background(), workspace, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService ListFederatedNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamespaceAPI.ListFederatedNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService ListNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamespaceAPI.ListNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService ListUserWorkspaceNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string

		resp, httpRes, err := apiClient.NamespaceAPI.ListUserWorkspaceNamespaces(context.Background(), workspace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService ListWorkspaceFederatedNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string

		resp, httpRes, err := apiClient.NamespaceAPI.ListWorkspaceFederatedNamespaces(context.Background(), workspace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService ListWorkspaceNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string
		var workspacemember string

		resp, httpRes, err := apiClient.NamespaceAPI.ListWorkspaceNamespaces(context.Background(), workspace, workspacemember).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService PatchNamespace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string
		var namespace string

		resp, httpRes, err := apiClient.NamespaceAPI.PatchNamespace(context.Background(), workspace, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceAPIService UpdateNamespace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspace string
		var namespace string

		resp, httpRes, err := apiClient.NamespaceAPI.UpdateNamespace(context.Background(), workspace, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}