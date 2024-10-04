# \UserRelatedResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespace**](UserRelatedResourcesAPI.md#CreateNamespace) | **Post** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/namespaces | Create namespace in workspace
[**CreateWorkspace**](UserRelatedResourcesAPI.md#CreateWorkspace) | **Post** /kapis/tenant.kubesphere.io/v1beta1/workspaces | Create workspace
[**CreateWorkspaceResourceQuota**](UserRelatedResourcesAPI.md#CreateWorkspaceResourceQuota) | **Post** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/resourcequotas | Create workspace resource quota
[**CreateWorkspaceTemplate**](UserRelatedResourcesAPI.md#CreateWorkspaceTemplate) | **Post** /kapis/tenant.kubesphere.io/v1beta1/workspacetemplates | Create workspace template
[**DeleteNamespace**](UserRelatedResourcesAPI.md#DeleteNamespace) | **Delete** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/namespaces/{namespace} | Delete namespace from workspace
[**DeleteWorkspace**](UserRelatedResourcesAPI.md#DeleteWorkspace) | **Delete** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace} | Delete workspace
[**DeleteWorkspaceResourceQuota**](UserRelatedResourcesAPI.md#DeleteWorkspaceResourceQuota) | **Delete** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/resourcequotas/{resourcequota} | Delete workspace resource quota.
[**DeleteWorkspaceTemplate**](UserRelatedResourcesAPI.md#DeleteWorkspaceTemplate) | **Delete** /kapis/tenant.kubesphere.io/v1beta1/workspacetemplates/{workspace} | Delete workspace template
[**DescribeNamespace**](UserRelatedResourcesAPI.md#DescribeNamespace) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/namespaces/{namespace} | Get namespace
[**DescribeNamespaceV3**](UserRelatedResourcesAPI.md#DescribeNamespaceV3) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/workspaces/{workspace}/namespaces/{namespace} | Retrieve namespace details.
[**DescribeWorkspaceResourceQuota**](UserRelatedResourcesAPI.md#DescribeWorkspaceResourceQuota) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/resourcequotas/{resourcequota} | Get workspace resource quota
[**DescribeWorkspaceTemplate**](UserRelatedResourcesAPI.md#DescribeWorkspaceTemplate) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/workspacetemplates/{workspace} | Describe workspace.
[**GetKubeconfig**](UserRelatedResourcesAPI.md#GetKubeconfig) | **Get** /kapis/resources.kubesphere.io/v1alpha2/users/{user}/kubeconfig | Get users&#39; kubeconfig
[**GetPlatformMetrics**](UserRelatedResourcesAPI.md#GetPlatformMetrics) | **Get** /kapis/tenant.kubesphere.io/v1beta1/metrics | Get platform metrics
[**GetWorkspace**](UserRelatedResourcesAPI.md#GetWorkspace) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/workspaces/{workspace} | Get workspace.
[**GetWorkspaceMetrics**](UserRelatedResourcesAPI.md#GetWorkspaceMetrics) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/metrics | Get workspace metrics
[**GetWorkspaceTemplate**](UserRelatedResourcesAPI.md#GetWorkspaceTemplate) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspacetemplates/{workspace} | Get workspace template
[**GetWorkspace_0**](UserRelatedResourcesAPI.md#GetWorkspace_0) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace} | Get workspace
[**ListClusters**](UserRelatedResourcesAPI.md#ListClusters) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/clusters | List clusters available to users
[**ListNamespaces**](UserRelatedResourcesAPI.md#ListNamespaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/namespaces | List the namespaces for the current user
[**ListNamespacesWorkspace**](UserRelatedResourcesAPI.md#ListNamespacesWorkspace) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/namespaces | List the namespaces in workspace
[**ListNamespacesWorkspaceMember**](UserRelatedResourcesAPI.md#ListNamespacesWorkspaceMember) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/workspacemembers/{workspacemember}/namespaces | List namespaces in workspace of the member
[**ListNamespaces_0**](UserRelatedResourcesAPI.md#ListNamespaces_0) | **Get** /kapis/tenant.kubesphere.io/v1beta1/namespaces | List the namespaces for the current user
[**ListWorkspaceClusters**](UserRelatedResourcesAPI.md#ListWorkspaceClusters) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/clusters | List clusters authorized to the specified workspace
[**ListWorkspaceClustersV3**](UserRelatedResourcesAPI.md#ListWorkspaceClustersV3) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/workspaces/{workspace}/clusters | List clusters authorized to the specified workspace.
[**ListWorkspaceNamespaces**](UserRelatedResourcesAPI.md#ListWorkspaceNamespaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/workspaces/{workspace}/namespaces | List the namespaces of the specified workspace for the current user
[**ListWorkspaceTemplates**](UserRelatedResourcesAPI.md#ListWorkspaceTemplates) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/workspacetemplates | List all workspaces that belongs to the current user
[**ListWorkspaceTemplates_0**](UserRelatedResourcesAPI.md#ListWorkspaceTemplates_0) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspacetemplates | List all workspace templates
[**ListWorkspaces**](UserRelatedResourcesAPI.md#ListWorkspaces) | **Get** /kapis/tenant.kubesphere.io/v1alpha3/workspaces | List all workspaces that belongs to the current user
[**ListWorkspaces_0**](UserRelatedResourcesAPI.md#ListWorkspaces_0) | **Get** /kapis/tenant.kubesphere.io/v1beta1/workspaces | List workspaces
[**PatchNamespace**](UserRelatedResourcesAPI.md#PatchNamespace) | **Patch** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/namespaces/{namespace} | Patch namespace
[**PatchWorkspace**](UserRelatedResourcesAPI.md#PatchWorkspace) | **Patch** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace} | Update workspace
[**PatchWorkspaceTemplate**](UserRelatedResourcesAPI.md#PatchWorkspaceTemplate) | **Patch** /kapis/tenant.kubesphere.io/v1beta1/workspacetemplates/{workspace} | Patch workspace template
[**PatchWorkspaceTemplateClustersVisibility**](UserRelatedResourcesAPI.md#PatchWorkspaceTemplateClustersVisibility) | **Post** /kapis/cluster.kubesphere.io/v1alpha1/clusters/{cluster}/grantrequests | Patch workspace template&#39;s visibility in different clusters
[**UpdateNamespace**](UserRelatedResourcesAPI.md#UpdateNamespace) | **Put** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/namespaces/{namespace} | Update namespace
[**UpdateWorkspace**](UserRelatedResourcesAPI.md#UpdateWorkspace) | **Put** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace} | Update workspace
[**UpdateWorkspaceResourceQuota**](UserRelatedResourcesAPI.md#UpdateWorkspaceResourceQuota) | **Put** /kapis/tenant.kubesphere.io/v1beta1/workspaces/{workspace}/resourcequotas/{resourcequota} | Update workspace resource quota
[**UpdateWorkspaceTemplate**](UserRelatedResourcesAPI.md#UpdateWorkspaceTemplate) | **Put** /kapis/tenant.kubesphere.io/v1beta1/workspacetemplates/{workspace} | Update workspace template
[**UserRelatedClusters**](UserRelatedResourcesAPI.md#UserRelatedClusters) | **Get** /kapis/tenant.kubesphere.io/v1beta1/clusters | List clusters available to users



## CreateNamespace

> V1Namespace CreateNamespace(ctx, workspace).Body(body).Execute()

Create namespace in workspace

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
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.CreateNamespace(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.CreateNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.CreateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> V1beta1WorkspaceTemplate CreateWorkspace(ctx).Body(body).Execute()

Create workspace

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
	body := *openapiclient.NewV1beta1WorkspaceTemplate() // V1beta1WorkspaceTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.CreateWorkspace(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.CreateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspace`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md) |  | 

### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceResourceQuota

> V1alpha2ResourceQuota CreateWorkspaceResourceQuota(ctx, workspace).Body(body).Execute()

Create workspace resource quota

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
	body := *openapiclient.NewV1alpha2ResourceQuota(*openapiclient.NewV1alpha2ResourceQuotaSpec(*openapiclient.NewV1ResourceQuotaSpec(), map[string]string{"key": "Inner_example"})) // V1alpha2ResourceQuota | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.CreateWorkspaceResourceQuota(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.CreateWorkspaceResourceQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceResourceQuota`: V1alpha2ResourceQuota
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.CreateWorkspaceResourceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceResourceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha2ResourceQuota**](V1alpha2ResourceQuota.md) |  | 

### Return type

[**V1alpha2ResourceQuota**](V1alpha2ResourceQuota.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceTemplate

> V1beta1WorkspaceTemplate CreateWorkspaceTemplate(ctx).Body(body).Execute()

Create workspace template

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
	body := *openapiclient.NewV1beta1WorkspaceTemplate() // V1beta1WorkspaceTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.CreateWorkspaceTemplate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.CreateWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceTemplate`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.CreateWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md) |  | 

### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespace

> ErrorsError DeleteNamespace(ctx, workspace, namespace).Execute()

Delete namespace from workspace

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
	namespace := "namespace_example" // string | The specified namespace.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.DeleteNamespace(context.Background(), workspace, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DeleteNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespace`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DeleteNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceRequest struct via the builder pattern


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


## DeleteWorkspace

> ErrorsError DeleteWorkspace(ctx, workspace).Execute()

Delete workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.DeleteWorkspace(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DeleteWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspace`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DeleteWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRequest struct via the builder pattern


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


## DeleteWorkspaceResourceQuota

> ErrorsError DeleteWorkspaceResourceQuota(ctx, workspace, resourcequota).Execute()

Delete workspace resource quota.

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
	resourcequota := "resourcequota_example" // string | resource quota name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.DeleteWorkspaceResourceQuota(context.Background(), workspace, resourcequota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DeleteWorkspaceResourceQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceResourceQuota`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DeleteWorkspaceResourceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**resourcequota** | **string** | resource quota name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceResourceQuotaRequest struct via the builder pattern


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


## DeleteWorkspaceTemplate

> ErrorsError DeleteWorkspaceTemplate(ctx, workspace).Execute()

Delete workspace template

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.DeleteWorkspaceTemplate(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DeleteWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceTemplate`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DeleteWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceTemplateRequest struct via the builder pattern


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


## DescribeNamespace

> V1Namespace DescribeNamespace(ctx, workspace, namespace).Execute()

Get namespace

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
	namespace := "namespace_example" // string | The specified namespace.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.DescribeNamespace(context.Background(), workspace, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DescribeNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DescribeNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeNamespaceV3

> V1Namespace DescribeNamespaceV3(ctx, workspace, namespace).Execute()

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.DescribeNamespaceV3(context.Background(), workspace, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DescribeNamespaceV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeNamespaceV3`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DescribeNamespaceV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**namespace** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeNamespaceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspaceResourceQuota

> V1alpha2ResourceQuota DescribeWorkspaceResourceQuota(ctx, workspace, resourcequota).Execute()

Get workspace resource quota

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
	resourcequota := "resourcequota_example" // string | Resource quota name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.DescribeWorkspaceResourceQuota(context.Background(), workspace, resourcequota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DescribeWorkspaceResourceQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceResourceQuota`: V1alpha2ResourceQuota
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DescribeWorkspaceResourceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**resourcequota** | **string** | Resource quota name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceResourceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1alpha2ResourceQuota**](V1alpha2ResourceQuota.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspaceTemplate

> V1alpha2WorkspaceTemplate DescribeWorkspaceTemplate(ctx, workspace).Execute()

Describe workspace.

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.DescribeWorkspaceTemplate(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.DescribeWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceTemplate`: V1alpha2WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.DescribeWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2WorkspaceTemplate**](V1alpha2WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubeconfig

> string GetKubeconfig(ctx, user).Execute()

Get users' kubeconfig

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
	user := "user_example" // string | username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.GetKubeconfig(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.GetKubeconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubeconfig`: string
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.GetKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformMetrics

> OverviewMetricResults GetPlatformMetrics(ctx).Execute()

Get platform metrics

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.GetPlatformMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.GetPlatformMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformMetrics`: OverviewMetricResults
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.GetPlatformMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformMetricsRequest struct via the builder pattern


### Return type

[**OverviewMetricResults**](OverviewMetricResults.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspace

> V1alpha1Workspace GetWorkspace(ctx, workspace).Execute()

Get workspace.

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.GetWorkspace(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.GetWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspace`: V1alpha1Workspace
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.GetWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha1Workspace**](V1alpha1Workspace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceMetrics

> OverviewMetricResults GetWorkspaceMetrics(ctx, workspace).Execute()

Get workspace metrics

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.GetWorkspaceMetrics(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.GetWorkspaceMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceMetrics`: OverviewMetricResults
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.GetWorkspaceMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OverviewMetricResults**](OverviewMetricResults.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceTemplate

> V1beta1WorkspaceTemplate GetWorkspaceTemplate(ctx, workspace).Execute()

Get workspace template

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.GetWorkspaceTemplate(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.GetWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceTemplate`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.GetWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspace_0

> V1beta1WorkspaceTemplate GetWorkspace_0(ctx, workspace).Execute()

Get workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.GetWorkspace_0(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.GetWorkspace_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspace_0`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.GetWorkspace_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspace_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> ApiListResult ListClusters(ctx).Execute()

List clusters available to users

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListNamespaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaces`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesRequest struct via the builder pattern


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


## ListNamespacesWorkspace

> ApiListResult ListNamespacesWorkspace(ctx, workspace).Execute()

List the namespaces in workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListNamespacesWorkspace(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListNamespacesWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacesWorkspace`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListNamespacesWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListNamespacesWorkspaceMember

> V1Namespace ListNamespacesWorkspaceMember(ctx, workspace, workspacemember).Body(body).Execute()

List namespaces in workspace of the member

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
	workspacemember := "workspacemember_example" // string | workspacemember username
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListNamespacesWorkspaceMember(context.Background(), workspace, workspacemember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListNamespacesWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacesWorkspaceMember`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListNamespacesWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**workspacemember** | **string** | workspacemember username | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesWorkspaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaces_0

> ApiListResult ListNamespaces_0(ctx).Execute()

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListNamespaces_0(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListNamespaces_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaces_0`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListNamespaces_0`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespaces_2Request struct via the builder pattern


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


## ListWorkspaceClusters

> ApiListResult ListWorkspaceClusters(ctx, workspace).Execute()

List clusters authorized to the specified workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListWorkspaceClusters(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListWorkspaceClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceClusters`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListWorkspaceClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListWorkspaceClustersV3

> ApiListResult ListWorkspaceClustersV3(ctx, workspace).Execute()

List clusters authorized to the specified workspace.

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListWorkspaceClustersV3(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListWorkspaceClustersV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceClustersV3`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListWorkspaceClustersV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceClustersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListWorkspaceNamespaces

> ApiListResult ListWorkspaceNamespaces(ctx, workspace).Execute()

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListWorkspaceNamespaces(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListWorkspaceNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceNamespaces`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListWorkspaceNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListWorkspaceTemplates

> ModelsPageableResponse ListWorkspaceTemplates(ctx).Execute()

List all workspaces that belongs to the current user

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListWorkspaceTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListWorkspaceTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceTemplates`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListWorkspaceTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceTemplatesRequest struct via the builder pattern


### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceTemplates_0

> ModelsPageableResponse ListWorkspaceTemplates_0(ctx).Execute()

List all workspace templates

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListWorkspaceTemplates_0(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListWorkspaceTemplates_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceTemplates_0`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListWorkspaceTemplates_0`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceTemplates_3Request struct via the builder pattern


### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> ModelsPageableResponse ListWorkspaces(ctx).Execute()

List all workspaces that belongs to the current user

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListWorkspaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaces`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces_0

> ModelsPageableResponse ListWorkspaces_0(ctx).Execute()

List workspaces

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.ListWorkspaces_0(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.ListWorkspaces_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaces_0`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.ListWorkspaces_0`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaces_4Request struct via the builder pattern


### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespace

> V1Namespace PatchNamespace(ctx, workspace, namespace).Body(body).Execute()

Patch namespace



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
	namespace := "namespace_example" // string | The specified namespace.
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.PatchNamespace(context.Background(), workspace, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.PatchNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.PatchNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkspace

> V1beta1WorkspaceTemplate PatchWorkspace(ctx, workspace).Body(body).Execute()

Update workspace

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
	body := *openapiclient.NewV1beta1WorkspaceTemplate() // V1beta1WorkspaceTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.PatchWorkspace(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.PatchWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkspace`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.PatchWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md) |  | 

### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkspaceTemplate

> V1beta1WorkspaceTemplate PatchWorkspaceTemplate(ctx, workspace).Body(body).Execute()

Patch workspace template

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
	body := *openapiclient.NewV1beta1WorkspaceTemplate() // V1beta1WorkspaceTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.PatchWorkspaceTemplate(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.PatchWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkspaceTemplate`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.PatchWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md) |  | 

### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkspaceTemplateClustersVisibility

> V1beta1WorkspaceTemplate PatchWorkspaceTemplateClustersVisibility(ctx, cluster).Body(body).Execute()

Patch workspace template's visibility in different clusters

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
	cluster := "cluster_example" // string | The specified cluster.
	body := []openapiclient.V1alpha1UpdateVisibilityRequest{*openapiclient.NewV1alpha1UpdateVisibilityRequest("Op_example", "Workspace_example")} // []V1alpha1UpdateVisibilityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.PatchWorkspaceTemplateClustersVisibility(context.Background(), cluster).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.PatchWorkspaceTemplateClustersVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkspaceTemplateClustersVisibility`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.PatchWorkspaceTemplateClustersVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The specified cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkspaceTemplateClustersVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]V1alpha1UpdateVisibilityRequest**](V1alpha1UpdateVisibilityRequest.md) |  | 

### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNamespace

> V1Namespace UpdateNamespace(ctx, workspace, namespace).Body(body).Execute()

Update namespace



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
	namespace := "namespace_example" // string | The specified namespace.
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.UpdateNamespace(context.Background(), workspace, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.UpdateNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamespace`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.UpdateNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspace

> V1beta1WorkspaceTemplate UpdateWorkspace(ctx, workspace).Body(body).Execute()

Update workspace

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
	body := *openapiclient.NewV1beta1WorkspaceTemplate() // V1beta1WorkspaceTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.UpdateWorkspace(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.UpdateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspace`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.UpdateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md) |  | 

### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspaceResourceQuota

> V1alpha2ResourceQuota UpdateWorkspaceResourceQuota(ctx, workspace, resourcequota).Body(body).Execute()

Update workspace resource quota

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
	resourcequota := "resourcequota_example" // string | Resource quota name
	body := *openapiclient.NewV1alpha2ResourceQuota(*openapiclient.NewV1alpha2ResourceQuotaSpec(*openapiclient.NewV1ResourceQuotaSpec(), map[string]string{"key": "Inner_example"})) // V1alpha2ResourceQuota | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.UpdateWorkspaceResourceQuota(context.Background(), workspace, resourcequota).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.UpdateWorkspaceResourceQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceResourceQuota`: V1alpha2ResourceQuota
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.UpdateWorkspaceResourceQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 
**resourcequota** | **string** | Resource quota name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceResourceQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1alpha2ResourceQuota**](V1alpha2ResourceQuota.md) |  | 

### Return type

[**V1alpha2ResourceQuota**](V1alpha2ResourceQuota.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspaceTemplate

> V1beta1WorkspaceTemplate UpdateWorkspaceTemplate(ctx, workspace).Body(body).Execute()

Update workspace template

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
	body := *openapiclient.NewV1beta1WorkspaceTemplate() // V1beta1WorkspaceTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRelatedResourcesAPI.UpdateWorkspaceTemplate(context.Background(), workspace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.UpdateWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceTemplate`: V1beta1WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.UpdateWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | The specified workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md) |  | 

### Return type

[**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserRelatedClusters

> ApiListResult UserRelatedClusters(ctx).Execute()

List clusters available to users

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
	resp, r, err := apiClient.UserRelatedResourcesAPI.UserRelatedClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRelatedResourcesAPI.UserRelatedClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserRelatedClusters`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `UserRelatedResourcesAPI.UserRelatedClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserRelatedClustersRequest struct via the builder pattern


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

