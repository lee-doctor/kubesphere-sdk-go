# \ContainerMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleContainerMetricsQuery**](ContainerMetricsAPI.md#HandleContainerMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/pods/{pod}/containers/{container} | Get container-level metric data of a specific container. Navigate to the container by the pod name and the namespace.
[**HandlePodContainerMetricsQuery**](ContainerMetricsAPI.md#HandlePodContainerMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/pods/{pod}/containers | Get container-level metric data of a specific pod&#39;s containers. Navigate to the pod by the pod&#39;s namespace.



## HandleContainerMetricsQuery

> MonitoringMetrics HandleContainerMetricsQuery(ctx, namespace, pod, container).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()

Get container-level metric data of a specific container. Navigate to the container by the pod name and the namespace.

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
	namespace := "namespace_example" // string | The name of the namespace.
	pod := "pod_example" // string | Pod name.
	container := "container_example" // string | Container name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both container CPU usage and memory usage: `container_cpu_usage|container_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerMetricsAPI.HandleContainerMetricsQuery(context.Background(), namespace, pod, container).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsAPI.HandleContainerMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleContainerMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsAPI.HandleContainerMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 
**pod** | **string** | Pod name. | 
**container** | **string** | Container name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleContainerMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both container CPU usage and memory usage: &#x60;container_cpu_usage|container_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
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


## HandlePodContainerMetricsQuery

> MonitoringMetrics HandlePodContainerMetricsQuery(ctx, namespace, pod).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get container-level metric data of a specific pod's containers. Navigate to the pod by the pod's namespace.

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
	namespace := "namespace_example" // string | The name of the namespace.
	pod := "pod_example" // string | Pod name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both container CPU usage and memory usage: `container_cpu_usage|container_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The container filter consists of a regexp pattern. It specifies which container data to return. For example, the following filter matches container prometheus and prometheus-config-reloader: `prometheus|prometheus-config-reloader`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort containers by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerMetricsAPI.HandlePodContainerMetricsQuery(context.Background(), namespace, pod).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsAPI.HandlePodContainerMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandlePodContainerMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsAPI.HandlePodContainerMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 
**pod** | **string** | Pod name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandlePodContainerMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both container CPU usage and memory usage: &#x60;container_cpu_usage|container_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The container filter consists of a regexp pattern. It specifies which container data to return. For example, the following filter matches container prometheus and prometheus-config-reloader: &#x60;prometheus|prometheus-config-reloader&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort containers by the specified metric. Not applicable if **start** and **end** are provided. | 
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

