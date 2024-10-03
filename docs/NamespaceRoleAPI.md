# \NamespaceRoleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespaceRole**](NamespaceRoleAPI.md#CreateNamespaceRole) | **Post** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/roles | Create role in the specified namespace.
[**CreateRoleBinding**](NamespaceRoleAPI.md#CreateRoleBinding) | **Post** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/rolebindings | Create rolebinding in the specified namespace.
[**DeleteNamespaceRole**](NamespaceRoleAPI.md#DeleteNamespaceRole) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/roles/{role} | Delete role in the specified namespace.
[**DescribeNamespaceRole**](NamespaceRoleAPI.md#DescribeNamespaceRole) | **Get** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/roles/{role} | Retrieve role details.
[**ListNamespaceRoles**](NamespaceRoleAPI.md#ListNamespaceRoles) | **Get** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/roles | List all roles in the specified namespace.
[**PatchNamespaceRole**](NamespaceRoleAPI.md#PatchNamespaceRole) | **Patch** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/roles/{role} | Patch namespace role.
[**RetrieveNamespaceMemberRoleTemplates**](NamespaceRoleAPI.md#RetrieveNamespaceMemberRoleTemplates) | **Get** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/members/{member}/roles | Retrieve member&#39;s role templates in namespace.
[**UpdateNamespaceRole**](NamespaceRoleAPI.md#UpdateNamespaceRole) | **Put** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/roles/{role} | Update namespace role.



## CreateNamespaceRole

> V1Role CreateNamespaceRole(ctx, namespace).Body(body).Execute()

Create role in the specified namespace.

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
	namespace := "namespace_example" // string | namespace
	body := *openapiclient.NewV1Role([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.CreateNamespaceRole(context.Background(), namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.CreateNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.CreateNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1Role**](V1Role.md) |  | 

### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoleBinding

> []V1RoleBinding CreateRoleBinding(ctx, namespace).Body(body).Execute()

Create rolebinding in the specified namespace.

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
	namespace := "namespace_example" // string | namespace
	body := []openapiclient.V1RoleBinding{*openapiclient.NewV1RoleBinding(*openapiclient.NewV1RoleRef("ApiGroup_example", "Kind_example", "Name_example"))} // []V1RoleBinding | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.CreateRoleBinding(context.Background(), namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.CreateRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoleBinding`: []V1RoleBinding
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.CreateRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]V1RoleBinding**](V1RoleBinding.md) |  | 

### Return type

[**[]V1RoleBinding**](V1RoleBinding.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespaceRole

> ErrorsError DeleteNamespaceRole(ctx, namespace, role).Execute()

Delete role in the specified namespace.

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
	namespace := "namespace_example" // string | namespace
	role := "role_example" // string | role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.DeleteNamespaceRole(context.Background(), namespace, role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.DeleteNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespaceRole`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.DeleteNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceRoleRequest struct via the builder pattern


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


## DescribeNamespaceRole

> V1Role DescribeNamespaceRole(ctx, namespace, role).Execute()

Retrieve role details.

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
	namespace := "namespace_example" // string | namespace
	role := "role_example" // string | role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.DescribeNamespaceRole(context.Background(), namespace, role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.DescribeNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.DescribeNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaceRoles

> ApiListResult ListNamespaceRoles(ctx, namespace).Execute()

List all roles in the specified namespace.

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
	namespace := "namespace_example" // string | namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.ListNamespaceRoles(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.ListNamespaceRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaceRoles`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.ListNamespaceRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespaceRolesRequest struct via the builder pattern


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


## PatchNamespaceRole

> V1Role PatchNamespaceRole(ctx, namespace, role).Body(body).Execute()

Patch namespace role.

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
	namespace := "namespace_example" // string | namespace
	role := "role_example" // string | role name
	body := *openapiclient.NewV1Role([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.PatchNamespaceRole(context.Background(), namespace, role).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.PatchNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.PatchNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Role**](V1Role.md) |  | 

### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNamespaceMemberRoleTemplates

> ApiListResult RetrieveNamespaceMemberRoleTemplates(ctx, namespace, member).Execute()

Retrieve member's role templates in namespace.

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
	namespace := "namespace_example" // string | namespace
	member := "member_example" // string | namespace member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.RetrieveNamespaceMemberRoleTemplates(context.Background(), namespace, member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.RetrieveNamespaceMemberRoleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveNamespaceMemberRoleTemplates`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.RetrieveNamespaceMemberRoleTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**member** | **string** | namespace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNamespaceMemberRoleTemplatesRequest struct via the builder pattern


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


## UpdateNamespaceRole

> V1Role UpdateNamespaceRole(ctx, namespace, role).Body(body).Execute()

Update namespace role.

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
	namespace := "namespace_example" // string | namespace
	role := "role_example" // string | role name
	body := *openapiclient.NewV1Role([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceRoleAPI.UpdateNamespaceRole(context.Background(), namespace, role).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceRoleAPI.UpdateNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `NamespaceRoleAPI.UpdateNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Role**](V1Role.md) |  | 

### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

