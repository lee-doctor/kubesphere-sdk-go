# \CustomMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleClusterAdhocQuery**](CustomMetricsAPI.md#HandleClusterAdhocQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/targets/query | Make an ad-hoc query in the whole cluster.
[**HandleClusterMetadataQuery**](CustomMetricsAPI.md#HandleClusterMetadataQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/targets/metadata | Get metadata of metrics in the whole cluster.
[**HandleClusterMetricLabelSetQuery**](CustomMetricsAPI.md#HandleClusterMetricLabelSetQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/targets/labelsets | List all available labels and values of a metric within a specific time span in the whole cluster.
[**HandleMetadataQuery**](CustomMetricsAPI.md#HandleMetadataQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/targets/metadata | Get metadata of metrics for the specific namespace.
[**HandleMetricLabelSetQuery**](CustomMetricsAPI.md#HandleMetricLabelSetQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/targets/labelsets | List all available labels and values of a metric within a specific time span.
[**HandleNamespaceAdhocQuery**](CustomMetricsAPI.md#HandleNamespaceAdhocQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/targets/query | Make an ad-hoc query in the specific namespace.



## HandleClusterAdhocQuery

> MonitoringMetric HandleClusterAdhocQuery(ctx).Expr(expr).Start(start).End(end).Step(step).Time(time).Execute()

Make an ad-hoc query in the whole cluster.

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
	expr := "expr_example" // string | The expression to be evaluated. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomMetricsAPI.HandleClusterAdhocQuery(context.Background()).Expr(expr).Start(start).End(end).Step(step).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomMetricsAPI.HandleClusterAdhocQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleClusterAdhocQuery`: MonitoringMetric
	fmt.Fprintf(os.Stdout, "Response from `CustomMetricsAPI.HandleClusterAdhocQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleClusterAdhocQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expr** | **string** | The expression to be evaluated. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 

### Return type

[**MonitoringMetric**](MonitoringMetric.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleClusterMetadataQuery

> MonitoringMetadata HandleClusterMetadataQuery(ctx).Execute()

Get metadata of metrics in the whole cluster.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomMetricsAPI.HandleClusterMetadataQuery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomMetricsAPI.HandleClusterMetadataQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleClusterMetadataQuery`: MonitoringMetadata
	fmt.Fprintf(os.Stdout, "Response from `CustomMetricsAPI.HandleClusterMetadataQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleClusterMetadataQueryRequest struct via the builder pattern


### Return type

[**MonitoringMetadata**](MonitoringMetadata.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleClusterMetricLabelSetQuery

> MonitoringMetricLabelSet HandleClusterMetricLabelSetQuery(ctx).Metric(metric).Start(start).End(end).Execute()

List all available labels and values of a metric within a specific time span in the whole cluster.

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
	metric := "metric_example" // string | The name of the metric
	start := "start_example" // string | Start time of query. It is a string with Unix time format, eg. 1559347200. 
	end := "end_example" // string | End time of query. It is a string with Unix time format, eg. 1561939200. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomMetricsAPI.HandleClusterMetricLabelSetQuery(context.Background()).Metric(metric).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomMetricsAPI.HandleClusterMetricLabelSetQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleClusterMetricLabelSetQuery`: MonitoringMetricLabelSet
	fmt.Fprintf(os.Stdout, "Response from `CustomMetricsAPI.HandleClusterMetricLabelSetQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleClusterMetricLabelSetQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **string** | The name of the metric | 
 **start** | **string** | Start time of query. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. It is a string with Unix time format, eg. 1561939200.  | 

### Return type

[**MonitoringMetricLabelSet**](MonitoringMetricLabelSet.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleMetadataQuery

> MonitoringMetadata HandleMetadataQuery(ctx, namespace).Execute()

Get metadata of metrics for the specific namespace.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomMetricsAPI.HandleMetadataQuery(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomMetricsAPI.HandleMetadataQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleMetadataQuery`: MonitoringMetadata
	fmt.Fprintf(os.Stdout, "Response from `CustomMetricsAPI.HandleMetadataQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleMetadataQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitoringMetadata**](MonitoringMetadata.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleMetricLabelSetQuery

> MonitoringMetricLabelSet HandleMetricLabelSetQuery(ctx, namespace).Metric(metric).Start(start).End(end).Execute()

List all available labels and values of a metric within a specific time span.

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
	metric := "metric_example" // string | The name of the metric
	start := "start_example" // string | Start time of query. It is a string with Unix time format, eg. 1559347200. 
	end := "end_example" // string | End time of query. It is a string with Unix time format, eg. 1561939200. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomMetricsAPI.HandleMetricLabelSetQuery(context.Background(), namespace).Metric(metric).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomMetricsAPI.HandleMetricLabelSetQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleMetricLabelSetQuery`: MonitoringMetricLabelSet
	fmt.Fprintf(os.Stdout, "Response from `CustomMetricsAPI.HandleMetricLabelSetQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleMetricLabelSetQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metric** | **string** | The name of the metric | 
 **start** | **string** | Start time of query. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. It is a string with Unix time format, eg. 1561939200.  | 

### Return type

[**MonitoringMetricLabelSet**](MonitoringMetricLabelSet.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleNamespaceAdhocQuery

> MonitoringMetric HandleNamespaceAdhocQuery(ctx, namespace).Expr(expr).Start(start).End(end).Step(step).Time(time).Execute()

Make an ad-hoc query in the specific namespace.

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
	expr := "expr_example" // string | The expression to be evaluated. (optional)
	start := "start_example" // string | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  (optional)
	end := "end_example" // string | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  (optional)
	step := "step_example" // string | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). (optional) (default to "10m")
	time := "time_example" // string | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomMetricsAPI.HandleNamespaceAdhocQuery(context.Background(), namespace).Expr(expr).Start(start).End(end).Step(step).Time(time).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomMetricsAPI.HandleNamespaceAdhocQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleNamespaceAdhocQuery`: MonitoringMetric
	fmt.Fprintf(os.Stdout, "Response from `CustomMetricsAPI.HandleNamespaceAdhocQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The name of the namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleNamespaceAdhocQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expr** | **string** | The expression to be evaluated. | 
 **start** | **string** | Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200.  | 
 **end** | **string** | End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200.  | 
 **step** | **string** | Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min). | [default to &quot;10m&quot;]
 **time** | **string** | A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive. | 

### Return type

[**MonitoringMetric**](MonitoringMetric.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

