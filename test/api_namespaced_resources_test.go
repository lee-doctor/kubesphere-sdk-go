/*
KS API

Testing NamespacedResourcesAPIService

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

func Test_openapi_NamespacedResourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamespacedResourcesAPIService GetDaemonSetRevision", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var daemonset string
		var namespace string
		var revision string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetDaemonSetRevision(context.Background(), daemonset, namespace, revision).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetDeploymentRevision", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deployment string
		var namespace string
		var revision string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetDeploymentRevision(context.Background(), deployment, namespace, revision).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetImageConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetImageConfig(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetNamespaceOverview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetNamespaceOverview(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetNamespaceQuotas", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetNamespaceQuotas(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetNamespacedAbnormalWorkloads", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetNamespacedAbnormalWorkloads(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetNamespacedResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string
		var resources string
		var name string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetNamespacedResource(context.Background(), namespace, resources, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetRepositoryTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetRepositoryTags(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService GetStatefulSetRevision", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var statefulset string
		var namespace string
		var revision string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.GetStatefulSetRevision(context.Background(), statefulset, namespace, revision).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService ListNamespacedResources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string
		var resources string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.ListNamespacedResources(context.Background(), namespace, resources).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacedResourcesAPIService VerifyImageRepositorySecret", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string
		var secret string

		resp, httpRes, err := apiClient.NamespacedResourcesAPI.VerifyImageRepositorySecret(context.Background(), namespace, secret).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
