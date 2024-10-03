# \PVCMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleNamespacePVCMetricsQuery**](PVCMetricsAPI.md#HandleNamespacePVCMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/persistentvolumeclaims | Get PVC-level metric data of the specific namespace&#39;s PVCs.
[**HandlePVCMetricsQuery**](PVCMetricsAPI.md#HandlePVCMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/storageclasses/{storageclass}/persistentvolumeclaims | Get PVC-level metric data of the specific storageclass&#39;s PVCs.
[**HandleSpecificPVCMetricsQuery**](PVCMetricsAPI.md#HandleSpecificPVCMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/persistentvolumeclaims/{pvc} | Get PVC-level metric data of a specific PVC. Navigate to the PVC by the PVC&#39;s namespace.



## HandleNamespacePVCMetricsQuery

> MonitoringMetrics HandleNamespacePVCMetricsQuery(ctx, namespace).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get PVC-level metric data of the specific namespace's PVCs.

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
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both PVC available and used inodes: `pvc_inodes_available|pvc_inodes_used`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The PVC filter consists of a regexp pattern. It specifies which PVC data to return. For example, the following filter matches any pod whose name begins with redis: `redis.*`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort PVCs by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PVCMetricsAPI.HandleNamespacePVCMetricsQuery(context.Background(), namespace).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PVCMetricsAPI.HandleNamespacePVCMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleNamespacePVCMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PVCMetricsAPI.HandleNamespacePVCMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleNamespacePVCMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both PVC available and used inodes: &#x60;pvc_inodes_available|pvc_inodes_used&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The PVC filter consists of a regexp pattern. It specifies which PVC data to return. For example, the following filter matches any pod whose name begins with redis: &#x60;redis.*&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort PVCs by the specified metric. Not applicable if **start** and **end** are provided. | 
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


## HandlePVCMetricsQuery

> MonitoringMetrics HandlePVCMetricsQuery(ctx, storageclass).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()

Get PVC-level metric data of the specific storageclass's PVCs.

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
	storageclass := "storageclass_example" // string | The name of the storageclass.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both PVC available and used inodes: `pvc_inodes_available|pvc_inodes_used`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	resourcesFilter := "resourcesFilter_example" // string | The PVC filter consists of a regexp pattern. It specifies which PVC data to return. For example, the following filter matches any pod whose name begins with redis: `redis.*`. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)
	sortMetric := "sortMetric_example" // string | Sort PVCs by the specified metric. Not applicable if **start** and **end** are provided. (optional)
	sortType := "sortType_example" // string | Sort order. One of asc, desc. (optional) (default to "desc.")
	page := int32(56) // int32 | The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data. (optional)
	limit := int32(56) // int32 | Page size, the maximum number of results in a single page. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PVCMetricsAPI.HandlePVCMetricsQuery(context.Background(), storageclass).MetricsFilter(metricsFilter).ResourcesFilter(resourcesFilter).Start(start).End(end).Step(step).Time(time).SortMetric(sortMetric).SortType(sortType).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PVCMetricsAPI.HandlePVCMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandlePVCMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PVCMetricsAPI.HandlePVCMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageclass** | **string** | The name of the storageclass. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandlePVCMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both PVC available and used inodes: &#x60;pvc_inodes_available|pvc_inodes_used&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
 **resourcesFilter** | **string** | The PVC filter consists of a regexp pattern. It specifies which PVC data to return. For example, the following filter matches any pod whose name begins with redis: &#x60;redis.*&#x60;. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 
 **sortMetric** | **string** | Sort PVCs by the specified metric. Not applicable if **start** and **end** are provided. | 
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


## HandleSpecificPVCMetricsQuery

> MonitoringMetrics HandleSpecificPVCMetricsQuery(ctx, namespace, pvc).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()

Get PVC-level metric data of a specific PVC. Navigate to the PVC by the PVC's namespace.

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
	pvc := "pvc_example" // string | PVC name.
	metricsFilter := "metricsFilter_example" // string | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both PVC available and used inodes: `pvc_inodes_available|pvc_inodes_used`. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PVCMetricsAPI.HandleSpecificPVCMetricsQuery(context.Background(), namespace, pvc).MetricsFilter(metricsFilter).Start(start).End(end).Step(step).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PVCMetricsAPI.HandleSpecificPVCMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleSpecificPVCMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `PVCMetricsAPI.HandleSpecificPVCMetricsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 
**pvc** | **string** | PVC name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleSpecificPVCMetricsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricsFilter** | **string** | The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both PVC available and used inodes: &#x60;pvc_inodes_available|pvc_inodes_used&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/). | 
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

