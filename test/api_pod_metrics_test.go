/*
KubeSphere

Testing PodMetricsAPIService

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

func Test_openapi_PodMetricsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PodMetricsAPIService HandleAllPodMetricsQuery", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var node string

		resp, httpRes, err := apiClient.PodMetricsAPI.HandleAllPodMetricsQuery(context.Background(), node).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PodMetricsAPIService HandleClusterPodMetricsQuery", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PodMetricsAPI.HandleClusterPodMetricsQuery(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PodMetricsAPIService HandleNamespacePodMetricsQuery", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.PodMetricsAPI.HandleNamespacePodMetricsQuery(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PodMetricsAPIService HandlePodMetricsQuery", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var node string
		var pod string

		resp, httpRes, err := apiClient.PodMetricsAPI.HandlePodMetricsQuery(context.Background(), node, pod).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PodMetricsAPIService HandleSpecificPodMetricsQuery", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string
		var pod string

		resp, httpRes, err := apiClient.PodMetricsAPI.HandleSpecificPodMetricsQuery(context.Background(), namespace, pod).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PodMetricsAPIService HandleSpecificWorkloadPodMetricsQuery", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string
		var kind string
		var workload string

		resp, httpRes, err := apiClient.PodMetricsAPI.HandleSpecificWorkloadPodMetricsQuery(context.Background(), namespace, kind, workload).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
