# \WorkspaceRoleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspaceRole**](WorkspaceRoleAPI.md#CreateWorkspaceRole) | **Post** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspaceroles | Create workspace role.
[**DeleteWorkspaceRole**](WorkspaceRoleAPI.md#DeleteWorkspaceRole) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspaceroles/{workspacerole} | Delete workspace role.
[**DescribeWorkspaceRole**](WorkspaceRoleAPI.md#DescribeWorkspaceRole) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspaceroles/{workspacerole} | Retrieve workspace role details.
[**ListWorkspaceRoles**](WorkspaceRoleAPI.md#ListWorkspaceRoles) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspaceroles | List all workspace roles.
[**PatchWorkspaceRole**](WorkspaceRoleAPI.md#PatchWorkspaceRole) | **Patch** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspaceroles/{workspacerole} | Patch workspace role.
[**RetrieveWorkspaceMemberRoleTemplates**](WorkspaceRoleAPI.md#RetrieveWorkspaceMemberRoleTemplates) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers/{workspacemember}/workspaceroles | Retrieve member&#39;s role templates in workspace.
[**UpdateWorkspaceRole**](WorkspaceRoleAPI.md#UpdateWorkspaceRole) | **Put** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspaceroles/{workspacerole} | Update workspace role.



## CreateWorkspaceRole

> V1alpha2WorkspaceRole CreateWorkspaceRole(ctx, workspace).Body(body).Execute()

Create workspace role.

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
	workspace := "workspace_example" // string | workspace name
	body := *openapiclient.NewV1alpha2WorkspaceRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1alpha2WorkspaceRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceRoleAPI.CreateWorkspaceRole(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceRoleAPI.CreateWorkspaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceRole`: V1alpha2WorkspaceRole
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceRoleAPI.CreateWorkspaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha2WorkspaceRole**](V1alpha2WorkspaceRole.md) |  | 

### Return type

[**V1alpha2WorkspaceRole**](V1alpha2WorkspaceRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspaceRole

> ErrorsError DeleteWorkspaceRole(ctx, workspace, workspacerole).Execute()

Delete workspace role.

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
	workspace := "workspace_example" // string | workspace name
	workspacerole := "workspacerole_example" // string | workspace role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceRoleAPI.DeleteWorkspaceRole(context.Background(), workspace, workspacerole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceRoleAPI.DeleteWorkspaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceRole`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceRoleAPI.DeleteWorkspaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacerole** | **string** | workspace role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRoleRequest struct via the builder pattern


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


## DescribeWorkspaceRole

> V1alpha2WorkspaceRole DescribeWorkspaceRole(ctx, workspace, workspacerole).Execute()

Retrieve workspace role details.

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
	workspace := "workspace_example" // string | workspace name
	workspacerole := "workspacerole_example" // string | workspace role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceRoleAPI.DescribeWorkspaceRole(context.Background(), workspace, workspacerole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceRoleAPI.DescribeWorkspaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceRole`: V1alpha2WorkspaceRole
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceRoleAPI.DescribeWorkspaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacerole** | **string** | workspace role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1alpha2WorkspaceRole**](V1alpha2WorkspaceRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceRoles

> ApiListResult ListWorkspaceRoles(ctx, workspace).Execute()

List all workspace roles.

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
	workspace := "workspace_example" // string | workspace name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceRoleAPI.ListWorkspaceRoles(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceRoleAPI.ListWorkspaceRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceRoles`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceRoleAPI.ListWorkspaceRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceRolesRequest struct via the builder pattern


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


## PatchWorkspaceRole

> ErrorsError PatchWorkspaceRole(ctx, workspace, workspacerole).Body(body).Execute()

Patch workspace role.

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
	workspace := "workspace_example" // string | workspace name
	workspacerole := "workspacerole_example" // string | workspace role name
	body := *openapiclient.NewV1alpha2WorkspaceRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1alpha2WorkspaceRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceRoleAPI.PatchWorkspaceRole(context.Background(), workspace, workspacerole).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceRoleAPI.PatchWorkspaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkspaceRole`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceRoleAPI.PatchWorkspaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacerole** | **string** | workspace role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkspaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1alpha2WorkspaceRole**](V1alpha2WorkspaceRole.md) |  | 

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


## RetrieveWorkspaceMemberRoleTemplates

> ApiListResult RetrieveWorkspaceMemberRoleTemplates(ctx, workspace, workspacemember).Execute()

Retrieve member's role templates in workspace.

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
	workspace := "workspace_example" // string | workspace
	workspacemember := "workspacemember_example" // string | workspace member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceRoleAPI.RetrieveWorkspaceMemberRoleTemplates(context.Background(), workspace, workspacemember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceRoleAPI.RetrieveWorkspaceMemberRoleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWorkspaceMemberRoleTemplates`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceRoleAPI.RetrieveWorkspaceMemberRoleTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace | 
**workspacemember** | **string** | workspace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWorkspaceMemberRoleTemplatesRequest struct via the builder pattern


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


## UpdateWorkspaceRole

> V1alpha2WorkspaceRole UpdateWorkspaceRole(ctx, workspace, workspacerole).Body(body).Execute()

Update workspace role.

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
	workspace := "workspace_example" // string | workspace name
	workspacerole := "workspacerole_example" // string | workspace role name
	body := *openapiclient.NewV1alpha2WorkspaceRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1alpha2WorkspaceRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceRoleAPI.UpdateWorkspaceRole(context.Background(), workspace, workspacerole).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceRoleAPI.UpdateWorkspaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceRole`: V1alpha2WorkspaceRole
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceRoleAPI.UpdateWorkspaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacerole** | **string** | workspace role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1alpha2WorkspaceRole**](V1alpha2WorkspaceRole.md) |  | 

### Return type

[**V1alpha2WorkspaceRole**](V1alpha2WorkspaceRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

