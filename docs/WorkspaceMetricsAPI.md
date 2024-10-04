# \WorkspaceMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleAllWorkspaceMetricsQuery**](WorkspaceMetricsAPI.md#HandleAllWorkspaceMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/workspaces | Get workspace-level metric data of all workspaces.
[**HandleWorkspaceMetricsQuery**](WorkspaceMetricsAPI.md#HandleWorkspaceMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/workspaces/{workspace} | Get workspace-level metric data of a specific workspace.



## HandleAllWorkspaceMetricsQuery

> MonitoringMetrics HandleAllWorkspaceMetricsQuery(ctx).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get workspace-level metric data of all workspaces.

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
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both workspace CPU usage and memory usage: `workspace_cpu_usage|workspace_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The workspace filter consists of a regexp pattern. It specifies which workspace data to return. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort workspaces by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMetricsAPI.HandleAllWorkspaceMetricsQuery(context.Background()).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMetricsAPI.HandleAllWorkspaceMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleAllWorkspaceMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMetricsAPI.HandleAllWorkspaceMetricsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleAllWorkspaceMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both workspace CPU usage and memory usage: &#x60;workspace_cpu_usage|workspace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The workspace filter consists of a regexp pattern. It specifies which workspace data to return. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort workspaces by the specified metric. Not applicable if **start** and **end** are provided. | 
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


## HandleWorkspaceMetricsQuery

> MonitoringMetrics HandleWorkspaceMetricsQuery(ctx, workspace).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Type_(type_).Execute()

Get workspace-level metric data of a specific workspace.

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
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both workspace CPU usage and memory usage: `workspace_cpu_usage|workspace_memory_usage`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	type_ := "type__example" // string | Additional operations. Currently available types is statistics. It retrieves the total number of namespaces, devops projects, members and roles in this workspace at the moment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMetricsAPI.HandleWorkspaceMetricsQuery(context.Background(), workspace).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMetricsAPI.HandleWorkspaceMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleWorkspaceMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMetricsAPI.HandleWorkspaceMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | Workspace name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleWorkspaceMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both workspace CPU usage and memory usage: &#x60;workspace_cpu_usage|workspace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **type_** | **string** | Additional operations. Currently available types is statistics. It retrieves the total number of namespaces, devops projects, members and roles in this workspace at the moment. | 

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

