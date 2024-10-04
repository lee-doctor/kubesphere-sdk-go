# \DevOpsProjectMemberAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespaceMembers**](DevOpsProjectMemberAPI.md#CreateNamespaceMembers) | **Post** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/members | Add members to the DevOps project in bulk.
[**DescribeDevOpsProjectNamespaceMember**](DevOpsProjectMemberAPI.md#DescribeDevOpsProjectNamespaceMember) | **Get** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/members/{member} | Retrieve devops project member details.
[**ListAllNamespaceMembers**](DevOpsProjectMemberAPI.md#ListAllNamespaceMembers) | **Get** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/members | List all members in the specified devops project.
[**RemoveDevOpsNamespaceMember**](DevOpsProjectMemberAPI.md#RemoveDevOpsNamespaceMember) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/members/{member} | Delete a member from the DevOps project.
[**UpdateDevOpsNamespaceMember**](DevOpsProjectMemberAPI.md#UpdateDevOpsNamespaceMember) | **Put** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/members/{member} | Update the role bind of the member.



## CreateNamespaceMembers

> []V1alpha2Member CreateNamespaceMembers(ctx, devops).Body(body).Execute()

Add members to the DevOps project in bulk.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | devops project name
	body := []openapiclient.V1alpha2Member{*openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example")} // []V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectMemberAPI.CreateNamespaceMembers(context.Background(), devops).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectMemberAPI.CreateNamespaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespaceMembers`: []V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectMemberAPI.CreateNamespaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceMembersRequest struct via the builder pattern


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


## DescribeDevOpsProjectNamespaceMember

> V1alpha2User DescribeDevOpsProjectNamespaceMember(ctx, devops, member).Execute()

Retrieve devops project member details.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | devops project name
	member := "member_example" // string | devops project member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectMemberAPI.DescribeDevOpsProjectNamespaceMember(context.Background(), devops, member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectMemberAPI.DescribeDevOpsProjectNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDevOpsProjectNamespaceMember`: V1alpha2User
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectMemberAPI.DescribeDevOpsProjectNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**member** | **string** | devops project member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDevOpsProjectNamespaceMemberRequest struct via the builder pattern


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


## ListAllNamespaceMembers

> ApiListResult ListAllNamespaceMembers(ctx, devops).Execute()

List all members in the specified devops project.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | devops project name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectMemberAPI.ListAllNamespaceMembers(context.Background(), devops).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectMemberAPI.ListAllNamespaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllNamespaceMembers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectMemberAPI.ListAllNamespaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllNamespaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RemoveDevOpsNamespaceMember

> ErrorsError RemoveDevOpsNamespaceMember(ctx, devops, member).Execute()

Delete a member from the DevOps project.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | devops project name
	member := "member_example" // string | devops project member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectMemberAPI.RemoveDevOpsNamespaceMember(context.Background(), devops, member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectMemberAPI.RemoveDevOpsNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDevOpsNamespaceMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectMemberAPI.RemoveDevOpsNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**member** | **string** | devops project member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDevOpsNamespaceMemberRequest struct via the builder pattern


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


## UpdateDevOpsNamespaceMember

> V1alpha2Member UpdateDevOpsNamespaceMember(ctx, devops, member).Body(body).Execute()

Update the role bind of the member.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | devops project name
	member := "member_example" // string | devops project member's username
	body := *openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example") // V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectMemberAPI.UpdateDevOpsNamespaceMember(context.Background(), devops, member).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectMemberAPI.UpdateDevOpsNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDevOpsNamespaceMember`: V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectMemberAPI.UpdateDevOpsNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**member** | **string** | devops project member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDevOpsNamespaceMemberRequest struct via the builder pattern


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

