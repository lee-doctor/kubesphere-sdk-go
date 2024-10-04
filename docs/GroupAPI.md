# \GroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupAPI.md#CreateGroup) | **Post** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups | Create Group
[**CreateGroupBinding**](GroupAPI.md#CreateGroupBinding) | **Post** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groupbindings | Create GroupBinding to add a user to the group
[**CreateWorkspaceRoleBinding**](GroupAPI.md#CreateWorkspaceRoleBinding) | **Post** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacerolebindings | Create group&#39;s workspacerolebindings of the workspace.
[**DeleteGroup**](GroupAPI.md#DeleteGroup) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups/{group} | Delete group.
[**DeleteGroupBinding**](GroupAPI.md#DeleteGroupBinding) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groupbindings/{groupbinding} | Delete GroupBinding to remove user from the group.
[**DeleteRoleBinding**](GroupAPI.md#DeleteRoleBinding) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/namespace/{namespace}/rolebindings/{rolebinding} | Delete rolebinding under namespace.
[**DeleteWorkspaceRoleBinding**](GroupAPI.md#DeleteWorkspaceRoleBinding) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacerolebindings/{rolebinding} | Delete workspacerolebinding.
[**DescribeGroup**](GroupAPI.md#DescribeGroup) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups/{group} | Retrieve group details.
[**ListGroupBindings**](GroupAPI.md#ListGroupBindings) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups/{group}/groupbindings | Retrieve group&#39;s members in the workspace.
[**ListGroupDevOpsRoleBindings**](GroupAPI.md#ListGroupDevOpsRoleBindings) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups/{group}/devopsrolebindings | Retrieve group&#39;s rolebindings of all devops projects in the workspace.
[**ListGroupRoleBindings**](GroupAPI.md#ListGroupRoleBindings) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups/{group}/rolebindings | Retrieve group&#39;s rolebindings of all projects in the workspace.
[**ListGroupWorkspaceRoleBindings**](GroupAPI.md#ListGroupWorkspaceRoleBindings) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups/{group}/workspacerolebindings | Retrieve group&#39;s workspacerolebindings of the workspace.
[**ListWorkspaceGroups**](GroupAPI.md#ListWorkspaceGroups) | **Get** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups | List groups of the specified workspace.
[**UpdateGroup**](GroupAPI.md#UpdateGroup) | **Put** /kapis/iam.kubesphere.io/v1alpha2/workspaces/{workspace}/groups/{group} | Update Group



## CreateGroup

> V1alpha2Group CreateGroup(ctx, workspace).Body(body).Execute()

Create Group

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
	body := *openapiclient.NewV1alpha2Group() // V1alpha2Group | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.CreateGroup(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: V1alpha2Group
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha2Group**](V1alpha2Group.md) |  | 

### Return type

[**V1alpha2Group**](V1alpha2Group.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupBinding

> V1alpha2GroupBinding CreateGroupBinding(ctx, workspace).Body(body).Execute()

Create GroupBinding to add a user to the group

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
	body := []openapiclient.V1alpha2GroupMember{*openapiclient.NewV1alpha2GroupMember("GroupName_example", "UserName_example")} // []V1alpha2GroupMember | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.CreateGroupBinding(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateGroupBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupBinding`: V1alpha2GroupBinding
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.CreateGroupBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]V1alpha2GroupMember**](V1alpha2GroupMember.md) |  | 

### Return type

[**V1alpha2GroupBinding**](V1alpha2GroupBinding.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceRoleBinding

> []V1alpha2WorkspaceRoleBinding CreateWorkspaceRoleBinding(ctx, workspace).Body(body).Execute()

Create group's workspacerolebindings of the workspace.

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
	body := []openapiclient.V1alpha2WorkspaceRoleBinding{*openapiclient.NewV1alpha2WorkspaceRoleBinding(*openapiclient.NewV1RoleRef("ApiGroup_example", "Kind_example", "Name_example"))} // []V1alpha2WorkspaceRoleBinding | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.CreateWorkspaceRoleBinding(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateWorkspaceRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceRoleBinding`: []V1alpha2WorkspaceRoleBinding
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.CreateWorkspaceRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]V1alpha2WorkspaceRoleBinding**](V1alpha2WorkspaceRoleBinding.md) |  | 

### Return type

[**[]V1alpha2WorkspaceRoleBinding**](V1alpha2WorkspaceRoleBinding.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> ErrorsError DeleteGroup(ctx, workspace, group).Execute()

Delete group.

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
	group := "group_example" // string | group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.DeleteGroup(context.Background(), workspace, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroup`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**group** | **string** | group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteGroupBinding

> ErrorsError DeleteGroupBinding(ctx, workspace, groupbinding).Execute()

Delete GroupBinding to remove user from the group.

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
	groupbinding := "groupbinding_example" // string | groupbinding name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.DeleteGroupBinding(context.Background(), workspace, groupbinding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteGroupBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupBinding`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.DeleteGroupBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**groupbinding** | **string** | groupbinding name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupBindingRequest struct via the builder pattern


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


## DeleteRoleBinding

> ErrorsError DeleteRoleBinding(ctx, workspace, namespace, rolebinding).Execute()

Delete rolebinding under namespace.

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
	namespace := "namespace_example" // string | groupbinding name
	rolebinding := "rolebinding_example" // string | groupbinding name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.DeleteRoleBinding(context.Background(), workspace, namespace, rolebinding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoleBinding`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.DeleteRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**namespace** | **string** | groupbinding name | 
**rolebinding** | **string** | groupbinding name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleBindingRequest struct via the builder pattern


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


## DeleteWorkspaceRoleBinding

> ErrorsError DeleteWorkspaceRoleBinding(ctx, workspace, rolebinding).Execute()

Delete workspacerolebinding.

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
	rolebinding := "rolebinding_example" // string | groupbinding name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.DeleteWorkspaceRoleBinding(context.Background(), workspace, rolebinding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteWorkspaceRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceRoleBinding`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.DeleteWorkspaceRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**rolebinding** | **string** | groupbinding name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRoleBindingRequest struct via the builder pattern


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


## DescribeGroup

> V1alpha2Group DescribeGroup(ctx, workspace, group).Execute()

Retrieve group details.

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
	group := "group_example" // string | group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.DescribeGroup(context.Background(), workspace, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DescribeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGroup`: V1alpha2Group
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.DescribeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**group** | **string** | group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1alpha2Group**](V1alpha2Group.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupBindings

> ApiListResult ListGroupBindings(ctx, workspace, group).Execute()

Retrieve group's members in the workspace.

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
	group := "group_example" // string | group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ListGroupBindings(context.Background(), workspace, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroupBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupBindings`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListGroupBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**group** | **string** | group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupBindingsRequest struct via the builder pattern


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


## ListGroupDevOpsRoleBindings

> ApiListResult ListGroupDevOpsRoleBindings(ctx, workspace, group).Execute()

Retrieve group's rolebindings of all devops projects in the workspace.

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
	group := "group_example" // string | group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ListGroupDevOpsRoleBindings(context.Background(), workspace, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroupDevOpsRoleBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupDevOpsRoleBindings`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListGroupDevOpsRoleBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**group** | **string** | group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupDevOpsRoleBindingsRequest struct via the builder pattern


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


## ListGroupRoleBindings

> ApiListResult ListGroupRoleBindings(ctx, workspace, group).Execute()

Retrieve group's rolebindings of all projects in the workspace.

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
	group := "group_example" // string | group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ListGroupRoleBindings(context.Background(), workspace, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroupRoleBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupRoleBindings`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListGroupRoleBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**group** | **string** | group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupRoleBindingsRequest struct via the builder pattern


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


## ListGroupWorkspaceRoleBindings

> ApiListResult ListGroupWorkspaceRoleBindings(ctx, workspace, group).Execute()

Retrieve group's workspacerolebindings of the workspace.

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
	group := "group_example" // string | group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ListGroupWorkspaceRoleBindings(context.Background(), workspace, group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListGroupWorkspaceRoleBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupWorkspaceRoleBindings`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListGroupWorkspaceRoleBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**group** | **string** | group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupWorkspaceRoleBindingsRequest struct via the builder pattern


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


## ListWorkspaceGroups

> ApiListResult ListWorkspaceGroups(ctx, workspace).Execute()

List groups of the specified workspace.

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
	resp, r, err := apiClient.GroupAPI.ListWorkspaceGroups(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ListWorkspaceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceGroups`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ListWorkspaceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceGroupsRequest struct via the builder pattern


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


## UpdateGroup

> V1alpha2Group UpdateGroup(ctx, workspace, group).Body(body).Execute()

Update Group

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
	group := "group_example" // string | group name
	body := *openapiclient.NewV1alpha2Group() // V1alpha2Group | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.UpdateGroup(context.Background(), workspace, group).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroup`: V1alpha2Group
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**group** | **string** | group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1alpha2Group**](V1alpha2Group.md) |  | 

### Return type

[**V1alpha2Group**](V1alpha2Group.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

