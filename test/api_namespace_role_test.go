/*
KubeSphere

Testing NamespaceRoleAPIService

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

func Test_openapi_NamespaceRoleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamespaceRoleAPIService CreateNamespaceRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.CreateNamespaceRole(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceRoleAPIService CreateRoleBinding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.CreateRoleBinding(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceRoleAPIService DeleteNamespaceRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var role string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.DeleteNamespaceRole(context.Background(), namespace, role).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceRoleAPIService DescribeNamespaceRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var role string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.DescribeNamespaceRole(context.Background(), namespace, role).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceRoleAPIService ListNamespaceRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.ListNamespaceRoles(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceRoleAPIService PatchNamespaceRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var role string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.PatchNamespaceRole(context.Background(), namespace, role).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceRoleAPIService RetrieveNamespaceMemberRoleTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var member string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.RetrieveNamespaceMemberRoleTemplates(context.Background(), namespace, member).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespaceRoleAPIService UpdateNamespaceRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var role string

		resp, httpRes, err := apiClient.NamespaceRoleAPI.UpdateNamespaceRole(context.Background(), namespace, role).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}