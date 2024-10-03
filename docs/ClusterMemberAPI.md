# \ClusterMemberAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterMembers**](ClusterMemberAPI.md#CreateClusterMembers) | **Post** /kapis/iam.kubesphere.io/v1alpha2/clustermembers | Add members to current cluster in bulk.
[**DescribeClusterMember**](ClusterMemberAPI.md#DescribeClusterMember) | **Get** /kapis/iam.kubesphere.io/v1alpha2/clustermembers/{clustermember} | Retrieve the cluster role of the specified member.
[**ListClusterMembers**](ClusterMemberAPI.md#ListClusterMembers) | **Get** /kapis/iam.kubesphere.io/v1alpha2/clustermembers | List all members in cluster.
[**RemoveClusterMember**](ClusterMemberAPI.md#RemoveClusterMember) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/clustermembers/{clustermember} | Delete a member from current cluster.
[**UpdateClusterMember**](ClusterMemberAPI.md#UpdateClusterMember) | **Put** /kapis/iam.kubesphere.io/v1alpha2/clustermembers/{clustermember} | Update the cluster role bind of the member.



## CreateClusterMembers

> []V1alpha2Member CreateClusterMembers(ctx).Body(body).Execute()

Add members to current cluster in bulk.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := []openapiclient.V1alpha2Member{*openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example")} // []V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterMemberAPI.CreateClusterMembers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterMemberAPI.CreateClusterMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClusterMembers`: []V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `ClusterMemberAPI.CreateClusterMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1alpha2Member**](V1alpha2Member.md) |  | 

### Return type

[**[]V1alpha2Member**](V1alpha2Member.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeClusterMember

> V1alpha2User DescribeClusterMember(ctx, clustermember).Execute()

Retrieve the cluster role of the specified member.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clustermember := "clustermember_example" // string | cluster member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterMemberAPI.DescribeClusterMember(context.Background(), clustermember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterMemberAPI.DescribeClusterMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeClusterMember`: V1alpha2User
	fmt.Fprintf(os.Stdout, "Response from `ClusterMemberAPI.DescribeClusterMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clustermember** | **string** | cluster member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeClusterMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2User**](V1alpha2User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterMembers

> ApiListResult ListClusterMembers(ctx).Execute()

List all members in cluster.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterMemberAPI.ListClusterMembers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterMemberAPI.ListClusterMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterMembers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `ClusterMemberAPI.ListClusterMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterMembersRequest struct via the builder pattern


### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveClusterMember

> ErrorsError RemoveClusterMember(ctx, clustermember).Execute()

Delete a member from current cluster.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clustermember := "clustermember_example" // string | cluster member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterMemberAPI.RemoveClusterMember(context.Background(), clustermember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterMemberAPI.RemoveClusterMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveClusterMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `ClusterMemberAPI.RemoveClusterMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clustermember** | **string** | cluster member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveClusterMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterMember

> V1alpha2Member UpdateClusterMember(ctx, clustermember).Body(body).Execute()

Update the cluster role bind of the member.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clustermember := "clustermember_example" // string | cluster member's username
	body := *openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example") // V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterMemberAPI.UpdateClusterMember(context.Background(), clustermember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterMemberAPI.UpdateClusterMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterMember`: V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `ClusterMemberAPI.UpdateClusterMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clustermember** | **string** | cluster member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha2Member**](V1alpha2Member.md) |  | 

### Return type

[**V1alpha2Member**](V1alpha2Member.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

