# \ClusterResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterAbnormalWorkloads**](ClusterResourcesAPI.md#GetClusterAbnormalWorkloads) | **Get** /kapis/resources.kubesphere.io/v1alpha2/abnormalworkloads | Get abnormal workloads
[**GetClusterOverview**](ClusterResourcesAPI.md#GetClusterOverview) | **Get** /kapis/resources.kubesphere.io/v1alpha3/metrics | Cluster summary
[**GetClusterQuotas**](ClusterResourcesAPI.md#GetClusterQuotas) | **Get** /kapis/resources.kubesphere.io/v1alpha2/quotas | Get whole cluster&#39;s resource usage
[**GetClusterResource**](ClusterResourcesAPI.md#GetClusterResource) | **Get** /kapis/resources.kubesphere.io/v1alpha3/{resources}/{name} | Get cluster scope resource
[**ListClusterResources**](ClusterResourcesAPI.md#ListClusterResources) | **Get** /kapis/resources.kubesphere.io/v1alpha3/{resources} | Cluster level resources



## GetClusterAbnormalWorkloads

> ApiWorkloads GetClusterAbnormalWorkloads(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterResourcesAPI.GetClusterAbnormalWorkloads(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.GetClusterAbnormalWorkloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterAbnormalWorkloads`: ApiWorkloads
	fmt.Fprintf(os.Stdout, "Response from `ClusterResourcesAPI.GetClusterAbnormalWorkloads`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterAbnormalWorkloadsRequest struct via the builder pattern


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


## GetClusterOverview

> OverviewMetricResults GetClusterOverview(ctx).Execute()

Cluster summary

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
	resp, r, err := apiClient.ClusterResourcesAPI.GetClusterOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.GetClusterOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterOverview`: OverviewMetricResults
	fmt.Fprintf(os.Stdout, "Response from `ClusterResourcesAPI.GetClusterOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterOverviewRequest struct via the builder pattern


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


## GetClusterQuotas

> ApiResourceQuota GetClusterQuotas(ctx).Execute()

Get whole cluster's resource usage

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
	resp, r, err := apiClient.ClusterResourcesAPI.GetClusterQuotas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.GetClusterQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterQuotas`: ApiResourceQuota
	fmt.Fprintf(os.Stdout, "Response from `ClusterResourcesAPI.GetClusterQuotas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterQuotasRequest struct via the builder pattern


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


## GetClusterResource

> GetClusterResource(ctx, resources, name).Execute()

Get cluster scope resource

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
	resources := "resources_example" // string | cluster level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | the name of the clustered resources

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterResourcesAPI.GetClusterResource(context.Background(), resources, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.GetClusterResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resources** | **string** | cluster level resource type, e.g. pods,jobs,configmaps,services. | 
**name** | **string** | the name of the clustered resources | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterResources

> ApiListResult ListClusterResources(ctx, resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()

Cluster level resources

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
	resources := "resources_example" // string | cluster level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | name used to do filtering (optional)
	page := "page_example" // string | page (optional) (default to "page=1")
	limit := "limit_example" // string | limit (optional)
	ascending := "ascending_example" // string | sort parameters, e.g. reverse=true (optional) (default to "ascending=false")
	sortBy := "sortBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterResourcesAPI.ListClusterResources(context.Background(), resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.ListClusterResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterResources`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `ClusterResourcesAPI.ListClusterResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resources** | **string** | cluster level resource type, e.g. pods,jobs,configmaps,services. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | name used to do filtering | 
 **page** | **string** | page | [default to &quot;page&#x3D;1&quot;]
 **limit** | **string** | limit | 
 **ascending** | **string** | sort parameters, e.g. reverse&#x3D;true | [default to &quot;ascending&#x3D;false&quot;]
 **sortBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

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

