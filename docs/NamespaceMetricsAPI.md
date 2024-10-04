# \NamespaceMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleAllNamespaceMetricsQuery**](NamespaceMetricsAPI.md#HandleAllNamespaceMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces | Get namespace-level metric data of all namespaces.
[**HandleNamespaceMetricsQuery**](NamespaceMetricsAPI.md#HandleNamespaceMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace} | Get namespace-level metric data of the specific namespace.
[**HandleWorkspaceNamespaceMetricsQuery**](NamespaceMetricsAPI.md#HandleWorkspaceNamespaceMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/workspaces/{workspace}/namespaces | Get namespace-level metric data of a specific workspace.



## HandleAllNamespaceMetricsQuery

> MonitoringMetrics HandleAllNamespaceMetricsQuery(ctx).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get namespace-level metric data of all namespaces.

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
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: `namespace_cpu_usage|namespace_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The namespace filter consists of a regexp pattern. It specifies which namespace data to return. For example, the following filter matches both namespace test and kube-system: `test|kube-system`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort namespaces by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMetricsAPI.HandleAllNamespaceMetricsQuery(context.Background()).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMetricsAPI.HandleAllNamespaceMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleAllNamespaceMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMetricsAPI.HandleAllNamespaceMetricsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleAllNamespaceMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: &#x60;namespace_cpu_usage|namespace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The namespace filter consists of a regexp pattern. It specifies which namespace data to return. For example, the following filter matches both namespace test and kube-system: &#x60;test|kube-system&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort namespaces by the specified metric. Not applicable if **start** and **end** are provided. | 
 **sortType** | **string** | Sort order. One of asc, desc. | [default to &quot;desc.&quot;]
 **page** | **int32** | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. | 
 **limit** | **int32** | Page size, the maximum number of results in a single page. Defaults to 5. | [default to 5]

### Return type

[**MonitoringMetrics**](MonitoringMetrics.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleNamespaceMetricsQuery

> MonitoringMetrics HandleNamespaceMetricsQuery(ctx, namespace).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()

Get namespace-level metric data of the specific namespace.

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
	namespace := "namespace_example" // string | The name of the namespace.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: `namespace_cpu_usage|namespace_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMetricsAPI.HandleNamespaceMetricsQuery(context.Background(), namespace).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMetricsAPI.HandleNamespaceMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleNamespaceMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMetricsAPI.HandleNamespaceMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleNamespaceMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: &#x60;namespace_cpu_usage|namespace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 

### Return type

[**MonitoringMetrics**](MonitoringMetrics.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleWorkspaceNamespaceMetricsQuery

> MonitoringMetrics HandleWorkspaceNamespaceMetricsQuery(ctx, workspace).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get namespace-level metric data of a specific workspace.

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
	workspace := "workspace_example" // string | Workspace name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: `namespace_cpu_usage|namespace_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The namespace filter consists of a regexp pattern. It specifies which namespace data to return. For example, the following filter matches both namespace test and kube-system: `test|kube-system`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort namespaces by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMetricsAPI.HandleWorkspaceNamespaceMetricsQuery(context.Background(), workspace).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMetricsAPI.HandleWorkspaceNamespaceMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleWorkspaceNamespaceMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMetricsAPI.HandleWorkspaceNamespaceMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | Workspace name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleWorkspaceNamespaceMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: &#x60;namespace_cpu_usage|namespace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The namespace filter consists of a regexp pattern. It specifies which namespace data to return. For example, the following filter matches both namespace test and kube-system: &#x60;test|kube-system&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort namespaces by the specified metric. Not applicable if **start** and **end** are provided. | 
 **sortType** | **string** | Sort order. One of asc, desc. | [default to &quot;desc.&quot;]
 **page** | **int32** | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. | 
 **limit** | **int32** | Page size, the maximum number of results in a single page. Defaults to 5. | [default to 5]

### Return type

[**MonitoringMetrics**](MonitoringMetrics.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

