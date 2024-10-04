# \NamespaceResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleCreateRouter**](NamespaceResourcesAPI.md#HandleCreateRouter) | **Post** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/router | Create a router for a specified project
[**HandleDeleteRouter**](NamespaceResourcesAPI.md#HandleDeleteRouter) | **Delete** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/router | List router of a specified project
[**HandleGetDaemonSetRevision**](NamespaceResourcesAPI.md#HandleGetDaemonSetRevision) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/daemonsets/{daemonset}/revisions/{revision} | Get the specified daemonset revision
[**HandleGetDeploymentRevision**](NamespaceResourcesAPI.md#HandleGetDeploymentRevision) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/deployments/{deployment}/revisions/{revision} | Get the specified deployment revision
[**HandleGetNamespaceQuotas**](NamespaceResourcesAPI.md#HandleGetNamespaceQuotas) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/quotas | get specified namespace&#39;s resource quota and usage
[**HandleGetNamespacedAbnormalWorkloads**](NamespaceResourcesAPI.md#HandleGetNamespacedAbnormalWorkloads) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/abnormalworkloads | get abnormal workloads&#39; count of specified namespace
[**HandleGetRouter**](NamespaceResourcesAPI.md#HandleGetRouter) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/router | List router of a specified project
[**HandleGetStatefulSetRevision**](NamespaceResourcesAPI.md#HandleGetStatefulSetRevision) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/statefulsets/{statefulset}/revisions/{revision} | Get the specified statefulset revision
[**HandleListNamespaceResources**](NamespaceResourcesAPI.md#HandleListNamespaceResources) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/{resources} | Namespace level resource query
[**HandleUpdateRouter**](NamespaceResourcesAPI.md#HandleUpdateRouter) | **Put** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/router | Update a router for a specified project



## HandleCreateRouter

> V1Service HandleCreateRouter(ctx, namespace).Execute()

Create a router for a specified project

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
	namespace := "namespace_example" // string | the name of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleCreateRouter(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleCreateRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleCreateRouter`: V1Service
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleCreateRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleCreateRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Service**](V1Service.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleDeleteRouter

> V1Service HandleDeleteRouter(ctx, namespace).Execute()

List router of a specified project

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
	namespace := "namespace_example" // string | the name of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleDeleteRouter(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleDeleteRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleDeleteRouter`: V1Service
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleDeleteRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleDeleteRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Service**](V1Service.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetDaemonSetRevision

> V1DaemonSet HandleGetDaemonSetRevision(ctx, daemonset, namespace, revision).Execute()

Get the specified daemonset revision

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
	daemonset := "daemonset_example" // string | the name of the daemonset
	namespace := "namespace_example" // string | the namespace of the daemonset
	revision := "revision_example" // string | the revision of the daemonset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleGetDaemonSetRevision(context.Background(), daemonset, namespace, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleGetDaemonSetRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetDaemonSetRevision`: V1DaemonSet
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleGetDaemonSetRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**daemonset** | **string** | the name of the daemonset | 
**namespace** | **string** | the namespace of the daemonset | 
**revision** | **string** | the revision of the daemonset | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetDaemonSetRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1DaemonSet**](V1DaemonSet.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetDeploymentRevision

> V1ReplicaSet HandleGetDeploymentRevision(ctx, deployment, namespace, revision).Execute()

Get the specified deployment revision

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
	deployment := "deployment_example" // string | the name of deployment
	namespace := "namespace_example" // string | the namespace of the deployment
	revision := "revision_example" // string | the revision of the deployment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleGetDeploymentRevision(context.Background(), deployment, namespace, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleGetDeploymentRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetDeploymentRevision`: V1ReplicaSet
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleGetDeploymentRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deployment** | **string** | the name of deployment | 
**namespace** | **string** | the namespace of the deployment | 
**revision** | **string** | the revision of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetDeploymentRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetNamespaceQuotas

> ApiResourceQuota HandleGetNamespaceQuotas(ctx, namespace).Execute()

get specified namespace's resource quota and usage

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
	namespace := "namespace_example" // string | the name of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleGetNamespaceQuotas(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleGetNamespaceQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetNamespaceQuotas`: ApiResourceQuota
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleGetNamespaceQuotas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetNamespaceQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResourceQuota**](ApiResourceQuota.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetNamespacedAbnormalWorkloads

> ApiWorkloads HandleGetNamespacedAbnormalWorkloads(ctx, namespace).Execute()

get abnormal workloads' count of specified namespace

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
	namespace := "namespace_example" // string | the name of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleGetNamespacedAbnormalWorkloads(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleGetNamespacedAbnormalWorkloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetNamespacedAbnormalWorkloads`: ApiWorkloads
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleGetNamespacedAbnormalWorkloads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetNamespacedAbnormalWorkloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiWorkloads**](ApiWorkloads.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetRouter

> V1Service HandleGetRouter(ctx, namespace).Execute()

List router of a specified project

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
	namespace := "namespace_example" // string | the name of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleGetRouter(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleGetRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetRouter`: V1Service
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleGetRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Service**](V1Service.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetStatefulSetRevision

> V1StatefulSet HandleGetStatefulSetRevision(ctx, statefulset, namespace, revision).Execute()

Get the specified statefulset revision

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
	statefulset := "statefulset_example" // string | the name of the statefulset
	namespace := "namespace_example" // string | the namespace of the statefulset
	revision := "revision_example" // string | the revision of the statefulset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleGetStatefulSetRevision(context.Background(), statefulset, namespace, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleGetStatefulSetRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetStatefulSetRevision`: V1StatefulSet
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleGetStatefulSetRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statefulset** | **string** | the name of the statefulset | 
**namespace** | **string** | the namespace of the statefulset | 
**revision** | **string** | the revision of the statefulset | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetStatefulSetRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1StatefulSet**](V1StatefulSet.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleListNamespaceResources

> ModelsPageableResponse HandleListNamespaceResources(ctx, namespace, resources).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

Namespace level resource query

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
	namespace := "namespace_example" // string | the name of the project
	resources := "resources_example" // string | namespace level resource type, e.g. pods,jobs,configmaps,services.
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleListNamespaceResources(context.Background(), namespace, resources).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleListNamespaceResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleListNamespaceResources`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleListNamespaceResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 
**resources** | **string** | namespace level resource type, e.g. pods,jobs,configmaps,services. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleListNamespaceResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditions** | **string** | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]
 **reverse** | **string** | sort parameters, e.g. reverse&#x3D;true | 
 **orderBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

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


## HandleUpdateRouter

> V1Service HandleUpdateRouter(ctx, namespace).Execute()

Update a router for a specified project

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
	namespace := "namespace_example" // string | the name of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceResourcesAPI.HandleUpdateRouter(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceResourcesAPI.HandleUpdateRouter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleUpdateRouter`: V1Service
	fmt.Fprintf(os.Stdout, "Response from `NamespaceResourcesAPI.HandleUpdateRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleUpdateRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1Service**](V1Service.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

