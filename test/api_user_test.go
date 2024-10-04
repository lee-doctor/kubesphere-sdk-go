/*
KubeSphere

Testing UserAPIService

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

func Test_openapi_UserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAPIService CreateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserAPI.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var user string

		resp, httpRes, err := apiClient.UserAPI.DeleteUser(context.Background(), user).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService DescribeUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var user string

		resp, httpRes, err := apiClient.UserAPI.DescribeUser(context.Background(), user).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ListUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserAPI.ListUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService ModifyPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var user string

		resp, httpRes, err := apiClient.UserAPI.ModifyPassword(context.Background(), user).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UpdateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var user string

		resp, httpRes, err := apiClient.UserAPI.UpdateUser(context.Background(), user).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
