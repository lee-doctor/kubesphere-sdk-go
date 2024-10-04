# \PodMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleAllPodMetricsQuery**](PodMetricsAPI.md#HandleAllPodMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/nodes/{node}/pods | Get pod-level metric data of all pods on a specific node.
[**HandleClusterPodMetricsQuery**](PodMetricsAPI.md#HandleClusterPodMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/pods | Get pod-level metric data of the whole cluster&#39;s pods.
[**HandleNamespacePodMetricsQuery**](PodMetricsAPI.md#HandleNamespacePodMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/pods | Get pod-level metric data of the specific namespace&#39;s pods.
[**HandlePodMetricsQuery**](PodMetricsAPI.md#HandlePodMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/nodes/{node}/pods/{pod} | Get pod-level metric data of a specific pod. Navigate to the pod by the node where it is scheduled.
[**HandleSpecificPodMetricsQuery**](PodMetricsAPI.md#HandleSpecificPodMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/pods/{pod} | Get pod-level metric data of a specific pod. Navigate to the pod by the pod&#39;s namespace.
[**HandleSpecificWorkloadPodMetricsQuery**](PodMetricsAPI.md#HandleSpecificWorkloadPodMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/workloads/{kind}/{workload}/pods | Get pod-level metric data of a specific workload&#39;s pods. Navigate to the workload by the namespace.



## HandleAllPodMetricsQuery

> MonitoringMetrics HandleAllPodMetricsQuery(ctx, node).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get pod-level metric data of all pods on a specific node.

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
	node := "node_example" // string | Node name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: `pod_cpu_usage|pod_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: `redis.*`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodMetricsAPI.HandleAllPodMetricsQuery(context.Background(), node).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodMetricsAPI.HandleAllPodMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleAllPodMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PodMetricsAPI.HandleAllPodMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | Node name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleAllPodMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: &#x60;pod_cpu_usage|pod_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: &#x60;redis.*&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. | 
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


## HandleClusterPodMetricsQuery

> MonitoringMetrics HandleClusterPodMetricsQuery(ctx).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get pod-level metric data of the whole cluster's pods.

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
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: `pod_cpu_usage|pod_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: `redis.*`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodMetricsAPI.HandleClusterPodMetricsQuery(context.Background()).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodMetricsAPI.HandleClusterPodMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleClusterPodMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PodMetricsAPI.HandleClusterPodMetricsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleClusterPodMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: &#x60;pod_cpu_usage|pod_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: &#x60;redis.*&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. | 
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


## HandleNamespacePodMetricsQuery

> MonitoringMetrics HandleNamespacePodMetricsQuery(ctx, namespace).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get pod-level metric data of the specific namespace's pods.

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
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: `pod_cpu_usage|pod_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: `redis.*`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodMetricsAPI.HandleNamespacePodMetricsQuery(context.Background(), namespace).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodMetricsAPI.HandleNamespacePodMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleNamespacePodMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PodMetricsAPI.HandleNamespacePodMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleNamespacePodMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: &#x60;pod_cpu_usage|pod_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: &#x60;redis.*&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. | 
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


## HandlePodMetricsQuery

> MonitoringMetrics HandlePodMetricsQuery(ctx, node, pod).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()

Get pod-level metric data of a specific pod. Navigate to the pod by the node where it is scheduled.

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
	node := "node_example" // string | Node name.
	pod := "pod_example" // string | Pod name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: `pod_cpu_usage|pod_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodMetricsAPI.HandlePodMetricsQuery(context.Background(), node, pod).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodMetricsAPI.HandlePodMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandlePodMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PodMetricsAPI.HandlePodMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | Node name. | 
**pod** | **string** | Pod name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandlePodMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: &#x60;pod_cpu_usage|pod_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
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


## HandleSpecificPodMetricsQuery

> MonitoringMetrics HandleSpecificPodMetricsQuery(ctx, namespace, pod).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()

Get pod-level metric data of a specific pod. Navigate to the pod by the pod's namespace.

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
	pod := "pod_example" // string | Pod name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: `pod_cpu_usage|pod_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodMetricsAPI.HandleSpecificPodMetricsQuery(context.Background(), namespace, pod).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodMetricsAPI.HandleSpecificPodMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleSpecificPodMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PodMetricsAPI.HandleSpecificPodMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 
**pod** | **string** | Pod name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleSpecificPodMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: &#x60;pod_cpu_usage|pod_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
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


## HandleSpecificWorkloadPodMetricsQuery

> MonitoringMetrics HandleSpecificWorkloadPodMetricsQuery(ctx, namespace, kind, workload).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get pod-level metric data of a specific workload's pods. Navigate to the workload by the namespace.

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
	kind := "kind_example" // string | Workload kind. One of deployment, daemonset, statefulset.
	workload := "workload_example" // string | Workload name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: `pod_cpu_usage|pod_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: `redis.*`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodMetricsAPI.HandleSpecificWorkloadPodMetricsQuery(context.Background(), namespace, kind, workload).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodMetricsAPI.HandleSpecificWorkloadPodMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleSpecificWorkloadPodMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PodMetricsAPI.HandleSpecificWorkloadPodMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 
**kind** | **string** | Workload kind. One of deployment, daemonset, statefulset. | 
**workload** | **string** | Workload name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleSpecificWorkloadPodMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both pod CPU usage and memory usage: &#x60;pod_cpu_usage|pod_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The pod filter consists of a regexp pattern. It specifies which pod data to return. For example, the following filter matches any pod whose name begins with redis: &#x60;redis.*&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort pods by the specified metric. Not applicable if **start** and **end** are provided. | 
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

