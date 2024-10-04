# \AccessManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterMembers**](AccessManagementAPI.md#CreateClusterMembers) | **Post** /kapis/iam.kubesphere.io/v1beta1/clustermembers | Add members to cluster
[**CreateNamespaceMembers**](AccessManagementAPI.md#CreateNamespaceMembers) | **Post** /kapis/iam.kubesphere.io/v1beta1/namespaces/{namespace}/namespacemembers | Add members to the namespace in bulk.
[**CreateSubjectAccessReview**](AccessManagementAPI.md#CreateSubjectAccessReview) | **Post** /kapis/iam.kubesphere.io/v1beta1/subjectaccessreviews | Create subject access review
[**CreateWorkspaceMembers**](AccessManagementAPI.md#CreateWorkspaceMembers) | **Post** /kapis/iam.kubesphere.io/v1beta1/workspaces/{workspace}/workspacemembers | Add members to the specified workspace
[**DescribeWorkspaceMember**](AccessManagementAPI.md#DescribeWorkspaceMember) | **Get** /kapis/iam.kubesphere.io/v1beta1/workspaces/{workspace}/workspacemembers/{workspacemember} | Get workspace member
[**ListClusterMembers**](AccessManagementAPI.md#ListClusterMembers) | **Get** /kapis/iam.kubesphere.io/v1beta1/clustermembers | List all members of cluster
[**ListNamespaceMembers**](AccessManagementAPI.md#ListNamespaceMembers) | **Get** /kapis/iam.kubesphere.io/v1beta1/namespaces/{namespace}/namespacemembers | List all members in the specified namespace
[**ListRoleTemplateOfUser**](AccessManagementAPI.md#ListRoleTemplateOfUser) | **Get** /kapis/iam.kubesphere.io/v1beta1/users/{username}/roletemplates | List all role templates of the specified user
[**ListWorkspaceMembers**](AccessManagementAPI.md#ListWorkspaceMembers) | **Get** /kapis/iam.kubesphere.io/v1beta1/workspaces/{workspace}/workspacemembers | List all members in the specified workspace
[**RemoveClusterMember**](AccessManagementAPI.md#RemoveClusterMember) | **Delete** /kapis/iam.kubesphere.io/v1beta1/clustermembers/{clustermember} | Delete member from cluster
[**RemoveNamespaceMember**](AccessManagementAPI.md#RemoveNamespaceMember) | **Delete** /kapis/iam.kubesphere.io/v1beta1/namespaces/{namespace}/namespacemembers/{member} | Delete a member from the namespace
[**RemoveWorkspaceMember**](AccessManagementAPI.md#RemoveWorkspaceMember) | **Delete** /kapis/iam.kubesphere.io/v1beta1/workspaces/{workspace}/workspacemembers/{workspacemember} | Delete a member from the workspace
[**UpdateClusterMember**](AccessManagementAPI.md#UpdateClusterMember) | **Put** /kapis/iam.kubesphere.io/v1beta1/clustermembers/{clustermember} | Update member from the cluster
[**UpdateNamespaceMember**](AccessManagementAPI.md#UpdateNamespaceMember) | **Put** /kapis/iam.kubesphere.io/v1beta1/namespaces/{namespace}/namespacemembers/{namespacemember} | Update member from the namespace
[**UpdateWorkspaceMember**](AccessManagementAPI.md#UpdateWorkspaceMember) | **Put** /kapis/iam.kubesphere.io/v1beta1/workspaces/{workspace}/workspacemembers/{workspacemember} | Update member from the workspace



## CreateClusterMembers

> []V1beta1Member CreateClusterMembers(ctx).Body(body).Execute()

Add members to cluster

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
	body := []openapiclient.V1beta1Member{*openapiclient.NewV1beta1Member("RoleRef_example", "Username_example")} // []V1beta1Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.CreateClusterMembers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.CreateClusterMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClusterMembers`: []V1beta1Member
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.CreateClusterMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1beta1Member**](V1beta1Member.md) |  | 

### Return type

[**[]V1beta1Member**](V1beta1Member.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNamespaceMembers

> []V1beta1Member CreateNamespaceMembers(ctx, namespace).Body(body).Execute()

Add members to the namespace in bulk.

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
	namespace := "namespace_example" // string | The specified namespace.
	body := []openapiclient.V1beta1Member{*openapiclient.NewV1beta1Member("RoleRef_example", "Username_example")} // []V1beta1Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.CreateNamespaceMembers(context.Background(), namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.CreateNamespaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespaceMembers`: []V1beta1Member
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.CreateNamespaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]V1beta1Member**](V1beta1Member.md) |  | 

### Return type

[**[]V1beta1Member**](V1beta1Member.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubjectAccessReview

> V1beta1SubjectAccessReview CreateSubjectAccessReview(ctx).Body(body).Execute()

Create subject access review



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
	body := *openapiclient.NewV1beta1SubjectAccessReview(*openapiclient.NewV1beta1SubjectAccessReviewSpec()) // V1beta1SubjectAccessReview | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.CreateSubjectAccessReview(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.CreateSubjectAccessReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubjectAccessReview`: V1beta1SubjectAccessReview
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.CreateSubjectAccessReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubjectAccessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1SubjectAccessReview**](V1beta1SubjectAccessReview.md) |  | 

### Return type

[**V1beta1SubjectAccessReview**](V1beta1SubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceMembers

> []V1beta1Member CreateWorkspaceMembers(ctx, workspace).Body(body).Execute()

Add members to the specified workspace

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
	workspace := "workspace_example" // string | The specified workspace.
	body := []openapiclient.V1beta1Member{*openapiclient.NewV1beta1Member("RoleRef_example", "Username_example")} // []V1beta1Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.CreateWorkspaceMembers(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.CreateWorkspaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceMembers`: []V1beta1Member
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.CreateWorkspaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]V1beta1Member**](V1beta1Member.md) |  | 

### Return type

[**[]V1beta1Member**](V1beta1Member.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspaceMember

> V1beta1User DescribeWorkspaceMember(ctx, workspace, workspacemember).Execute()

Get workspace member

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
	workspace := "workspace_example" // string | The specified workspace.
	workspacemember := "workspacemember_example" // string | Workspace member's name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.DescribeWorkspaceMember(context.Background(), workspace, workspacemember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.DescribeWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceMember`: V1beta1User
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.DescribeWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**workspacemember** | **string** | Workspace member&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1beta1User**](V1beta1User.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterMembers

> ApiListResult ListClusterMembers(ctx).Clusterrole(clusterrole).Execute()

List all members of cluster

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
	clusterrole := "clusterrole_example" // string | specific the cluster role name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.ListClusterMembers(context.Background()).Clusterrole(clusterrole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.ListClusterMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterMembers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.ListClusterMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClusterMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterrole** | **string** | specific the cluster role name | 

### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaceMembers

> ApiListResult ListNamespaceMembers(ctx, namespace).Role(role).Execute()

List all members in the specified namespace

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
	namespace := "namespace_example" // string | The specified namespace.
	role := "role_example" // string | specific the role name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.ListNamespaceMembers(context.Background(), namespace).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.ListNamespaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaceMembers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.ListNamespaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | specific the role name | 

### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleTemplateOfUser

> ApiListResult ListRoleTemplateOfUser(ctx, username).Scope(scope).Execute()

List all role templates of the specified user

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
	username := "username_example" // string | the name of the specified user
	scope := "scope_example" // string | the scope of role templates (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.ListRoleTemplateOfUser(context.Background(), username).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.ListRoleTemplateOfUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoleTemplateOfUser`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.ListRoleTemplateOfUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | the name of the specified user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleTemplateOfUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** | the scope of role templates | 

### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceMembers

> ApiListResult ListWorkspaceMembers(ctx, workspace).Workspacerole(workspacerole).Execute()

List all members in the specified workspace

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
	workspace := "workspace_example" // string | The specified workspace.
	workspacerole := "workspacerole_example" // string | specific the workspace role name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.ListWorkspaceMembers(context.Background(), workspace).Workspacerole(workspacerole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.ListWorkspaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceMembers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.ListWorkspaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspacerole** | **string** | specific the workspace role name | 

### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveClusterMember

> ErrorsError RemoveClusterMember(ctx, clustermember).Execute()

Delete member from cluster

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
	clustermember := "clustermember_example" // string | cluster member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.RemoveClusterMember(context.Background(), clustermember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.RemoveClusterMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveClusterMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.RemoveClusterMember`: %v\n", resp)
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

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNamespaceMember

> ErrorsError RemoveNamespaceMember(ctx, namespace, member).Execute()

Delete a member from the namespace

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
	namespace := "namespace_example" // string | The specified namespace.
	member := "member_example" // string | namespace member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.RemoveNamespaceMember(context.Background(), namespace, member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.RemoveNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNamespaceMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.RemoveNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**member** | **string** | namespace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNamespaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWorkspaceMember

> ErrorsError RemoveWorkspaceMember(ctx, workspace, workspacemember).Execute()

Delete a member from the workspace

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
	workspace := "workspace_example" // string | The specified workspace.
	workspacemember := "workspacemember_example" // string | Workspace member's name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.RemoveWorkspaceMember(context.Background(), workspace, workspacemember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.RemoveWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveWorkspaceMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.RemoveWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**workspacemember** | **string** | Workspace member&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWorkspaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterMember

> ErrorsError UpdateClusterMember(ctx, clustermember).Body(body).Execute()

Update member from the cluster

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
	clustermember := "clustermember_example" // string | the member name from cluster
	body := *openapiclient.NewV1beta1Member("RoleRef_example", "Username_example") // V1beta1Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.UpdateClusterMember(context.Background(), clustermember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.UpdateClusterMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.UpdateClusterMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clustermember** | **string** | the member name from cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1beta1Member**](V1beta1Member.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNamespaceMember

> ErrorsError UpdateNamespaceMember(ctx, namespace, namespacemember).Body(body).Execute()

Update member from the namespace

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
	namespace := "namespace_example" // string | The specified namespace.
	namespacemember := "namespacemember_example" // string | the member from namespace
	body := *openapiclient.NewV1beta1Member("RoleRef_example", "Username_example") // V1beta1Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.UpdateNamespaceMember(context.Background(), namespace, namespacemember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.UpdateNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamespaceMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.UpdateNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**namespacemember** | **string** | the member from namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamespaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1beta1Member**](V1beta1Member.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspaceMember

> ErrorsError UpdateWorkspaceMember(ctx, workspace, workspacemember).Body(body).Execute()

Update member from the workspace

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
	workspace := "workspace_example" // string | The specified workspace.
	workspacemember := "workspacemember_example" // string | the member from workspace
	body := *openapiclient.NewV1beta1Member("RoleRef_example", "Username_example") // V1beta1Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessManagementAPI.UpdateWorkspaceMember(context.Background(), workspace, workspacemember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessManagementAPI.UpdateWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AccessManagementAPI.UpdateWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**workspacemember** | **string** | the member from workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1beta1Member**](V1beta1Member.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

