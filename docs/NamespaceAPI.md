# \NamespaceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespace**](NamespaceAPI.md#CreateNamespace) | **Post** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/namespaces | List the namespaces of the specified workspace for the current user
[**DeleteNamespace**](NamespaceAPI.md#DeleteNamespace) | **Delete** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/namespaces/{namespace} | Delete namespace.
[**DescribeNamespace**](NamespaceAPI.md#DescribeNamespace) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/namespaces/{namespace} | Retrieve namespace details.
[**ListFederatedNamespaces**](NamespaceAPI.md#ListFederatedNamespaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/federatednamespaces | List the federated namespaces for the current user
[**ListNamespaces**](NamespaceAPI.md#ListNamespaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/namespaces | List the namespaces for the current user
[**ListUserWorkspaceNamespaces**](NamespaceAPI.md#ListUserWorkspaceNamespaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/namespaces | List the namespaces of the specified workspace for the current user
[**ListWorkspaceFederatedNamespaces**](NamespaceAPI.md#ListWorkspaceFederatedNamespaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/federatednamespaces | List the federated namespaces of the specified workspace for the current user
[**ListWorkspaceNamespaces**](NamespaceAPI.md#ListWorkspaceNamespaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers/{workspacemember}/namespaces | List the namespaces of the specified workspace for the workspace member
[**PatchNamespace**](NamespaceAPI.md#PatchNamespace) | **Patch** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/namespaces/{namespace} | 
[**UpdateNamespace**](NamespaceAPI.md#UpdateNamespace) | **Put** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/namespaces/{namespace} | 



## CreateNamespace

> V1Namespace CreateNamespace(ctx, workspace).Body(body).Execute()

List the namespaces of the specified workspace for the current user

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
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.CreateNamespace(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.CreateNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.CreateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespace

> ErrorsError DeleteNamespace(ctx, workspace, namespace).Execute()

Delete namespace.

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
	namespace := "namespace_example" // string | project name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.DeleteNamespace(context.Background(), workspace, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.DeleteNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespace`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.DeleteNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**namespace** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceRequest struct via the builder pattern


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


## DescribeNamespace

> V1Namespace DescribeNamespace(ctx, workspace, namespace).Execute()

Retrieve namespace details.

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
	namespace := "namespace_example" // string | project name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.DescribeNamespace(context.Background(), workspace, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.DescribeNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.DescribeNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**namespace** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFederatedNamespaces

> ApiListResult ListFederatedNamespaces(ctx).Execute()

List the federated namespaces for the current user

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.ListFederatedNamespaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.ListFederatedNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFederatedNamespaces`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.ListFederatedNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFederatedNamespacesRequest struct via the builder pattern


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


## ListNamespaces

> ApiListResult ListNamespaces(ctx).Execute()

List the namespaces for the current user

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.ListNamespaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.ListNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaces`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.ListNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesRequest struct via the builder pattern


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


## ListUserWorkspaceNamespaces

> ApiListResult ListUserWorkspaceNamespaces(ctx, workspace).Execute()

List the namespaces of the specified workspace for the current user

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
	resp, r, err := apiClient.NamespaceAPI.ListUserWorkspaceNamespaces(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.ListUserWorkspaceNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserWorkspaceNamespaces`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.ListUserWorkspaceNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserWorkspaceNamespacesRequest struct via the builder pattern


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


## ListWorkspaceFederatedNamespaces

> ApiListResult ListWorkspaceFederatedNamespaces(ctx, workspace).Execute()

List the federated namespaces of the specified workspace for the current user

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
	resp, r, err := apiClient.NamespaceAPI.ListWorkspaceFederatedNamespaces(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.ListWorkspaceFederatedNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceFederatedNamespaces`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.ListWorkspaceFederatedNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceFederatedNamespacesRequest struct via the builder pattern


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


## ListWorkspaceNamespaces

> V1Namespace ListWorkspaceNamespaces(ctx, workspace, workspacemember).Body(body).Execute()

List the namespaces of the specified workspace for the workspace member

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
	workspacemember := "workspacemember_example" // string | workspacemember username
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.ListWorkspaceNamespaces(context.Background(), workspace, workspacemember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.ListWorkspaceNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceNamespaces`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.ListWorkspaceNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacemember** | **string** | workspacemember username | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespace

> V1Namespace PatchNamespace(ctx, workspace, namespace).Body(body).Execute()



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
	namespace := "namespace_example" // string | project name
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.PatchNamespace(context.Background(), workspace, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.PatchNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.PatchNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**namespace** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNamespace

> V1Namespace UpdateNamespace(ctx, workspace, namespace).Body(body).Execute()



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
	namespace := "namespace_example" // string | project name
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceAPI.UpdateNamespace(context.Background(), workspace, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceAPI.UpdateNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `NamespaceAPI.UpdateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**namespace** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

