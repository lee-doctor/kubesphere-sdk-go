# \NamespacedResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDaemonSetRevision**](NamespacedResourcesAPI.md#GetDaemonSetRevision) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/daemonsets/{daemonset}/revisions/{revision} | Get the specified daemonSet revision
[**GetDeploymentRevision**](NamespacedResourcesAPI.md#GetDeploymentRevision) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/deployments/{deployment}/revisions/{revision} | Get the specified deployment revision
[**GetImageConfig**](NamespacedResourcesAPI.md#GetImageConfig) | **Get** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/imageconfig | Get image config
[**GetNamespaceOverview**](NamespacedResourcesAPI.md#GetNamespaceOverview) | **Get** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/metrics | Namespace summary
[**GetNamespaceQuotas**](NamespacedResourcesAPI.md#GetNamespaceQuotas) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/quotas | get specified namespace&#39;s resource quota and usage
[**GetNamespacedAbnormalWorkloads**](NamespacedResourcesAPI.md#GetNamespacedAbnormalWorkloads) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/abnormalworkloads | Get abnormal workloads
[**GetNamespacedResource**](NamespacedResourcesAPI.md#GetNamespacedResource) | **Get** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/{resources}/{name} | Get namespace scope resource
[**GetRepositoryTags**](NamespacedResourcesAPI.md#GetRepositoryTags) | **Get** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/repositorytags | List image tags
[**GetStatefulSetRevision**](NamespacedResourcesAPI.md#GetStatefulSetRevision) | **Get** /kapis/resources.kubesphere.io/v1alpha2/namespaces/{namespace}/statefulsets/{statefulset}/revisions/{revision} | Get the specified statefulSet revision
[**ListNamespacedResources**](NamespacedResourcesAPI.md#ListNamespacedResources) | **Get** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/{resources} | List resources at namespace scope
[**VerifyImageRepositorySecret**](NamespacedResourcesAPI.md#VerifyImageRepositorySecret) | **Post** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/registrysecrets/{secret}/verify | Verify registry credential



## GetDaemonSetRevision

> V1DaemonSet GetDaemonSetRevision(ctx, daemonset, namespace, revision).Execute()

Get the specified daemonSet revision

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
	namespace := "namespace_example" // string | The specified namespace.
	revision := "revision_example" // string | the revision of the daemonset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetDaemonSetRevision(context.Background(), daemonset, namespace, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetDaemonSetRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDaemonSetRevision`: V1DaemonSet
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetDaemonSetRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**daemonset** | **string** | the name of the daemonset | 
**namespace** | **string** | The specified namespace. | 
**revision** | **string** | the revision of the daemonset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDaemonSetRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1DaemonSet**](V1DaemonSet.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRevision

> V1ReplicaSet GetDeploymentRevision(ctx, deployment, namespace, revision).Execute()

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
	namespace := "namespace_example" // string | The specified namespace.
	revision := "revision_example" // string | the revision of the deployment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetDeploymentRevision(context.Background(), deployment, namespace, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetDeploymentRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentRevision`: V1ReplicaSet
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetDeploymentRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deployment** | **string** | the name of deployment | 
**namespace** | **string** | The specified namespace. | 
**revision** | **string** | the revision of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageConfig

> V2ImageConfig GetImageConfig(ctx, namespace).Image(image).Secret(secret).Execute()

Get image config

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
	image := "image_example" // string | Image name to query, e.g. kubesphere/ks-apiserver:v3.1.1
	secret := "secret_example" // string | Secret name of the image repository credential, left empty means anonymous fetch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetImageConfig(context.Background(), namespace).Image(image).Secret(secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetImageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageConfig`: V2ImageConfig
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetImageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | **string** | Image name to query, e.g. kubesphere/ks-apiserver:v3.1.1 | 
 **secret** | **string** | Secret name of the image repository credential, left empty means anonymous fetch. | 

### Return type

[**V2ImageConfig**](V2ImageConfig.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceOverview

> OverviewMetricResults GetNamespaceOverview(ctx, namespace).Execute()

Namespace summary

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetNamespaceOverview(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetNamespaceOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceOverview`: OverviewMetricResults
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetNamespaceOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceOverviewRequest struct via the builder pattern


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


## GetNamespaceQuotas

> ApiResourceQuota GetNamespaceQuotas(ctx, namespace).Execute()

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
	namespace := "namespace_example" // string | The specified namespace.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetNamespaceQuotas(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetNamespaceQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceQuotas`: ApiResourceQuota
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetNamespaceQuotas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResourceQuota**](ApiResourceQuota.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespacedAbnormalWorkloads

> ApiWorkloads GetNamespacedAbnormalWorkloads(ctx, namespace).Execute()

Get abnormal workloads

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetNamespacedAbnormalWorkloads(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetNamespacedAbnormalWorkloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespacedAbnormalWorkloads`: ApiWorkloads
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetNamespacedAbnormalWorkloads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacedAbnormalWorkloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiWorkloads**](ApiWorkloads.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespacedResource

> ApiListResult GetNamespacedResource(ctx, namespace, resources, name).Execute()

Get namespace scope resource

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
	resources := "resources_example" // string | namespace level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | the name of resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetNamespacedResource(context.Background(), namespace, resources, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetNamespacedResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespacedResource`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetNamespacedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**resources** | **string** | namespace level resource type, e.g. pods,jobs,configmaps,services. | 
**name** | **string** | the name of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacedResourceRequest struct via the builder pattern


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


## GetRepositoryTags

> V2RepositoryTags GetRepositoryTags(ctx, namespace).Repository(repository).Secret(secret).Page(page).Limit(limit).Ascending(ascending).Execute()

List image tags



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
	repository := "repository_example" // string | Repository to query, e.g. calico/cni.
	secret := "secret_example" // string | Secret name of the image repository credential, left empty means anonymous fetch. (optional)
	page := "page_example" // string | page (optional) (default to "page=1")
	limit := "limit_example" // string | limit (optional)
	ascending := "ascending_example" // string | sort parameters, e.g. reverse=true (optional) (default to "ascending=false")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetRepositoryTags(context.Background(), namespace).Repository(repository).Secret(secret).Page(page).Limit(limit).Ascending(ascending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetRepositoryTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryTags`: V2RepositoryTags
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetRepositoryTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | **string** | Repository to query, e.g. calico/cni. | 
 **secret** | **string** | Secret name of the image repository credential, left empty means anonymous fetch. | 
 **page** | **string** | page | [default to &quot;page&#x3D;1&quot;]
 **limit** | **string** | limit | 
 **ascending** | **string** | sort parameters, e.g. reverse&#x3D;true | [default to &quot;ascending&#x3D;false&quot;]

### Return type

[**V2RepositoryTags**](V2RepositoryTags.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatefulSetRevision

> V1StatefulSet GetStatefulSetRevision(ctx, statefulset, namespace, revision).Execute()

Get the specified statefulSet revision

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
	namespace := "namespace_example" // string | The specified namespace.
	revision := "revision_example" // string | the revision of the statefulset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.GetStatefulSetRevision(context.Background(), statefulset, namespace, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.GetStatefulSetRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatefulSetRevision`: V1StatefulSet
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.GetStatefulSetRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statefulset** | **string** | the name of the statefulset | 
**namespace** | **string** | The specified namespace. | 
**revision** | **string** | the revision of the statefulset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatefulSetRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1StatefulSet**](V1StatefulSet.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespacedResources

> ApiListResult ListNamespacedResources(ctx, namespace, resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).FieldSelector(fieldSelector).Execute()

List resources at namespace scope

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
	resources := "resources_example" // string | namespace level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | name used to do filtering (optional)
	page := "page_example" // string | page (optional) (default to "page=1")
	limit := "limit_example" // string | limit (optional)
	ascending := "ascending_example" // string | sort parameters, e.g. reverse=true (optional) (default to "ascending=false")
	sortBy := "sortBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)
	fieldSelector := "fieldSelector_example" // string | field selector used for filtering, you can use the = , == and != operators with field selectors( = and == mean the same thing), e.g. fieldSelector=type=kubernetes.io/dockerconfigjson, multiple separated by comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.ListNamespacedResources(context.Background(), namespace, resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).FieldSelector(fieldSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.ListNamespacedResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedResources`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.ListNamespacedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**resources** | **string** | namespace level resource type, e.g. pods,jobs,configmaps,services. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | name used to do filtering | 
 **page** | **string** | page | [default to &quot;page&#x3D;1&quot;]
 **limit** | **string** | limit | 
 **ascending** | **string** | sort parameters, e.g. reverse&#x3D;true | [default to &quot;ascending&#x3D;false&quot;]
 **sortBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 
 **fieldSelector** | **string** | field selector used for filtering, you can use the &#x3D; , &#x3D;&#x3D; and !&#x3D; operators with field selectors( &#x3D; and &#x3D;&#x3D; mean the same thing), e.g. fieldSelector&#x3D;type&#x3D;kubernetes.io/dockerconfigjson, multiple separated by comma | 

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


## VerifyImageRepositorySecret

> V1Secret VerifyImageRepositorySecret(ctx, namespace, secret).Body(body).Execute()

Verify registry credential

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
	secret := "secret_example" // string | Name of the secret.
	body := *openapiclient.NewV1Secret() // V1Secret | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourcesAPI.VerifyImageRepositorySecret(context.Background(), namespace, secret).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourcesAPI.VerifyImageRepositorySecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyImageRepositorySecret`: V1Secret
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourcesAPI.VerifyImageRepositorySecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**secret** | **string** | Name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyImageRepositorySecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Secret**](V1Secret.md) |  | 

### Return type

[**V1Secret**](V1Secret.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

