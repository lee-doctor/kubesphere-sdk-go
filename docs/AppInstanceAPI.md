# \AppInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspaceApplication**](AppInstanceAPI.md#CreateWorkspaceApplication) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/clusters/{cluster}/namespaces/{namespace}/applications | Deploy a new application
[**DeleteApplication**](AppInstanceAPI.md#DeleteApplication) | **Delete** /kapis/openpitrix.io/v1/workspaces/{workspace}/clusters/{cluster}/namespaces/{namespace}/applications/{application} | Delete the specified application
[**DescribeApplication**](AppInstanceAPI.md#DescribeApplication) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/clusters/{cluster}/namespaces/{namespace}/applications/{application} | Describe the specified application of the namespace
[**DescribeWorkspaceApplication**](AppInstanceAPI.md#DescribeWorkspaceApplication) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/namespaces/{namespace}/applications/{application} | Describe the specified application of the namespace
[**ListApplications**](AppInstanceAPI.md#ListApplications) | **Get** /kapis/openpitrix.io/v1/applications | List all applications
[**ListClusterApplications**](AppInstanceAPI.md#ListClusterApplications) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/clusters/{cluster}/applications | List all applications in special cluster
[**ListClusterWorkspaceApplications**](AppInstanceAPI.md#ListClusterWorkspaceApplications) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/clusters/{cluster}/namespaces/{namespace}/applications | List all applications within the specified namespace
[**ListWorkspaceApplications**](AppInstanceAPI.md#ListWorkspaceApplications) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/namespaces/{namespace}/applications | List all applications within the specified namespace
[**ModifyApplication**](AppInstanceAPI.md#ModifyApplication) | **Patch** /kapis/openpitrix.io/v1/workspaces/{workspace}/clusters/{cluster}/namespaces/{namespace}/applications/{application} | Modify application
[**UpgradeApplication**](AppInstanceAPI.md#UpgradeApplication) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/clusters/{cluster}/namespaces/{namespace}/applications/{application} | Upgrade application



## CreateWorkspaceApplication

> ErrorsError CreateWorkspaceApplication(ctx, workspace, cluster, namespace).Body(body).Execute()

Deploy a new application

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
	workspace := "workspace_example" // string | the name of the workspace.
	cluster := "cluster_example" // string | the name of the cluster.
	namespace := "namespace_example" // string | the name of the project
	body := *openapiclient.NewOpenpitrixCreateClusterRequest([]string{"AdvancedParam_example"}) // OpenpitrixCreateClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.CreateWorkspaceApplication(context.Background(), workspace, cluster, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.CreateWorkspaceApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceApplication`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.CreateWorkspaceApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**cluster** | **string** | the name of the cluster. | 
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**OpenpitrixCreateClusterRequest**](OpenpitrixCreateClusterRequest.md) |  | 

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


## DeleteApplication

> ErrorsError DeleteApplication(ctx, workspace, cluster, namespace, application).Execute()

Delete the specified application

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
	workspace := "workspace_example" // string | the name of the workspace.
	cluster := "cluster_example" // string | the name of the cluster.
	namespace := "namespace_example" // string | the name of the project
	application := "application_example" // string | the id of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.DeleteApplication(context.Background(), workspace, cluster, namespace, application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.DeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplication`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.DeleteApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**cluster** | **string** | the name of the cluster. | 
**namespace** | **string** | the name of the project | 
**application** | **string** | the id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## DescribeApplication

> OpenpitrixApplication DescribeApplication(ctx, workspace, cluster, namespace, application).Execute()

Describe the specified application of the namespace

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
	workspace := "workspace_example" // string | the name of the workspace.
	cluster := "cluster_example" // string | the name of the cluster.
	namespace := "namespace_example" // string | the name of the project
	application := "application_example" // string | the id of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.DescribeApplication(context.Background(), workspace, cluster, namespace, application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.DescribeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeApplication`: OpenpitrixApplication
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.DescribeApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**cluster** | **string** | the name of the cluster. | 
**namespace** | **string** | the name of the project | 
**application** | **string** | the id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OpenpitrixApplication**](OpenpitrixApplication.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspaceApplication

> OpenpitrixApplication DescribeWorkspaceApplication(ctx, workspace, namespace, application).Execute()

Describe the specified application of the namespace

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
	workspace := "workspace_example" // string | the name of the workspace
	namespace := "namespace_example" // string | the name of the project
	application := "application_example" // string | the id of the application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.DescribeWorkspaceApplication(context.Background(), workspace, namespace, application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.DescribeWorkspaceApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceApplication`: OpenpitrixApplication
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.DescribeWorkspaceApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace | 
**namespace** | **string** | the name of the project | 
**application** | **string** | the id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OpenpitrixApplication**](OpenpitrixApplication.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> ModelsPageableResponse ListApplications(ctx).Conditions(conditions).Paging(paging).Execute()

List all applications

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
	conditions := "conditions_example" // string | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.ListApplications(context.Background()).Conditions(conditions).Paging(paging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditions** | **string** | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterApplications

> ModelsPageableResponse ListClusterApplications(ctx, workspace, cluster).Conditions(conditions).Paging(paging).Execute()

List all applications in special cluster

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
	workspace := "workspace_example" // string | the name of the workspace.
	cluster := "cluster_example" // string | the name of the cluster.
	conditions := "conditions_example" // string | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.ListClusterApplications(context.Background(), workspace, cluster).Conditions(conditions).Paging(paging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.ListClusterApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterApplications`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.ListClusterApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**cluster** | **string** | the name of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditions** | **string** | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterWorkspaceApplications

> ModelsPageableResponse ListClusterWorkspaceApplications(ctx, workspace, cluster, namespace).Conditions(conditions).Paging(paging).Execute()

List all applications within the specified namespace

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
	workspace := "workspace_example" // string | the name of the workspace.
	cluster := "cluster_example" // string | the name of the cluster.
	namespace := "namespace_example" // string | the name of the project
	conditions := "conditions_example" // string | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.ListClusterWorkspaceApplications(context.Background(), workspace, cluster, namespace).Conditions(conditions).Paging(paging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.ListClusterWorkspaceApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterWorkspaceApplications`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.ListClusterWorkspaceApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**cluster** | **string** | the name of the cluster. | 
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterWorkspaceApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **conditions** | **string** | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceApplications

> ModelsPageableResponse ListWorkspaceApplications(ctx, workspace, namespace).Conditions(conditions).Paging(paging).Execute()

List all applications within the specified namespace

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
	workspace := "workspace_example" // string | the name of the workspace.
	namespace := "namespace_example" // string | the name of the project.
	conditions := "conditions_example" // string | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.ListWorkspaceApplications(context.Background(), workspace, namespace).Conditions(conditions).Paging(paging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.ListWorkspaceApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceApplications`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.ListWorkspaceApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**namespace** | **string** | the name of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditions** | **string** | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApplication

> ErrorsError ModifyApplication(ctx, workspace, cluster, namespace, application).Body(body).Execute()

Modify application

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
	workspace := "workspace_example" // string | the name of the workspace.
	cluster := "cluster_example" // string | the name of the cluster.
	namespace := "namespace_example" // string | the name of the project
	application := "application_example" // string | the id of the application
	body := *openapiclient.NewOpenpitrixModifyClusterAttributesRequest() // OpenpitrixModifyClusterAttributesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.ModifyApplication(context.Background(), workspace, cluster, namespace, application).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.ModifyApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApplication`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.ModifyApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**cluster** | **string** | the name of the cluster. | 
**namespace** | **string** | the name of the project | 
**application** | **string** | the id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**OpenpitrixModifyClusterAttributesRequest**](OpenpitrixModifyClusterAttributesRequest.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeApplication

> ErrorsError UpgradeApplication(ctx, workspace, cluster, namespace, application).Body(body).Execute()

Upgrade application

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
	workspace := "workspace_example" // string | the name of the workspace.
	cluster := "cluster_example" // string | the name of the cluster.
	namespace := "namespace_example" // string | the name of the project
	application := "application_example" // string | the id of the application
	body := *openapiclient.NewOpenpitrixUpgradeClusterRequest([]string{"AdvancedParam_example"}, "ClusterId_example") // OpenpitrixUpgradeClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInstanceAPI.UpgradeApplication(context.Background(), workspace, cluster, namespace, application).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInstanceAPI.UpgradeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeApplication`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppInstanceAPI.UpgradeApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace. | 
**cluster** | **string** | the name of the cluster. | 
**namespace** | **string** | the name of the project | 
**application** | **string** | the id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**OpenpitrixUpgradeClusterRequest**](OpenpitrixUpgradeClusterRequest.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

