# \WorkspaceMemberAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspaceMembers**](WorkspaceMemberAPI.md#CreateWorkspaceMembers) | **Post** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers | Add members to current cluster in bulk.
[**DescribeWorkspaceMember**](WorkspaceMemberAPI.md#DescribeWorkspaceMember) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers/{workspacemember} | Retrieve the workspace role of the specified member.
[**ListWorkspaceMembers**](WorkspaceMemberAPI.md#ListWorkspaceMembers) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers | List all members in the specified workspace.
[**RemoveWorkspaceMember**](WorkspaceMemberAPI.md#RemoveWorkspaceMember) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers/{workspacemember} | Delete a member from the workspace.
[**UpdateWorkspaceMember**](WorkspaceMemberAPI.md#UpdateWorkspaceMember) | **Put** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers/{workspacemember} | Update the workspace role bind of the member.



## CreateWorkspaceMembers

> []V1alpha2Member CreateWorkspaceMembers(ctx, workspace).Body(body).Execute()

Add members to current cluster in bulk.

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
	workspace := "workspace_example" // string | workspace name
	body := []openapiclient.V1alpha2Member{*openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example")} // []V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberAPI.CreateWorkspaceMembers(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberAPI.CreateWorkspaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceMembers`: []V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberAPI.CreateWorkspaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceMembersRequest struct via the builder pattern


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


## DescribeWorkspaceMember

> V1alpha2User DescribeWorkspaceMember(ctx, workspace, workspacemember).Execute()

Retrieve the workspace role of the specified member.

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
	workspace := "workspace_example" // string | workspace name
	workspacemember := "workspacemember_example" // string | workspace member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberAPI.DescribeWorkspaceMember(context.Background(), workspace, workspacemember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberAPI.DescribeWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceMember`: V1alpha2User
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberAPI.DescribeWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacemember** | **string** | workspace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceMemberRequest struct via the builder pattern


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


## ListWorkspaceMembers

> ApiListResult ListWorkspaceMembers(ctx, workspace).Execute()

List all members in the specified workspace.

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
	workspace := "workspace_example" // string | workspace name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberAPI.ListWorkspaceMembers(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberAPI.ListWorkspaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceMembers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberAPI.ListWorkspaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceMembersRequest struct via the builder pattern


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


## RemoveWorkspaceMember

> ErrorsError RemoveWorkspaceMember(ctx, workspace, workspacemember).Execute()

Delete a member from the workspace.

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
	workspace := "workspace_example" // string | workspace name
	workspacemember := "workspacemember_example" // string | workspace member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberAPI.RemoveWorkspaceMember(context.Background(), workspace, workspacemember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberAPI.RemoveWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveWorkspaceMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberAPI.RemoveWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacemember** | **string** | workspace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWorkspaceMemberRequest struct via the builder pattern


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


## UpdateWorkspaceMember

> V1alpha2Member UpdateWorkspaceMember(ctx, workspace, workspacemember).Body(body).Execute()

Update the workspace role bind of the member.

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
	workspace := "workspace_example" // string | workspace name
	workspacemember := "workspacemember_example" // string | workspace member's username
	body := *openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example") // V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberAPI.UpdateWorkspaceMember(context.Background(), workspace, workspacemember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberAPI.UpdateWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceMember`: V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberAPI.UpdateWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacemember** | **string** | workspace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceMemberRequest struct via the builder pattern


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

