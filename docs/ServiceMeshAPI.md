# \ServiceMeshAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppHealth**](ServiceMeshAPI.md#GetAppHealth) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/apps/{app}/health | Get app health
[**GetAppMetrics**](ServiceMeshAPI.md#GetAppMetrics) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/apps/{app}/metrics | Get app metrics from a specific namespace
[**GetNamespaceGraph**](ServiceMeshAPI.md#GetNamespaceGraph) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/graph | Get service graph for a specific namespace
[**GetNamespaceHealth**](ServiceMeshAPI.md#GetNamespaceHealth) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/health | Get app/service/workload health of a namespace
[**GetNamespaceMetrics**](ServiceMeshAPI.md#GetNamespaceMetrics) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/metrics | Get metrics from a specific namespace
[**GetServiceHealth**](ServiceMeshAPI.md#GetServiceHealth) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/services/{service}/health | Get service health
[**GetServiceMetrics**](ServiceMeshAPI.md#GetServiceMetrics) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/services/{service}/metrics | Get service metrics from a specific namespace
[**GetServiceTracing**](ServiceMeshAPI.md#GetServiceTracing) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/services/{service}/traces | Get tracing of a service, should have servicemesh enabled first
[**GetWorkloadHealth**](ServiceMeshAPI.md#GetWorkloadHealth) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/workloads/{workload}/health | Get workload health
[**GetWorkloadMetrics**](ServiceMeshAPI.md#GetWorkloadMetrics) | **Get** /kapis/servicemesh.kubesphere.io/v1alpha2/namespaces/{namespace}/workloads/{workload}/metrics | Get workload metrics from a specific namespace



## GetAppHealth

> map[string]interface{} GetAppHealth(ctx, namespace, app).RateInterval(rateInterval).QueryTime(queryTime).Execute()

Get app health

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
	namespace := "namespace_example" // string | name of a namespace
	app := "app_example" // string | app name
	rateInterval := "rateInterval_example" // string | the rate interval used for fetching error rate (default to "10m")
	queryTime := "queryTime_example" // string | the time to use for query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetAppHealth(context.Background(), namespace, app).RateInterval(rateInterval).QueryTime(queryTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetAppHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetAppHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of a namespace | 
**app** | **string** | app name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rateInterval** | **string** | the rate interval used for fetching error rate | [default to &quot;10m&quot;]
 **queryTime** | **string** | the time to use for query | 

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppMetrics

> map[string]interface{} GetAppMetrics(ctx, namespace, app).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()

Get app metrics from a specific namespace

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
	namespace := "namespace_example" // string | name of the namespace
	app := "app_example" // string | name of the app
	filters := "filters_example" // string | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count (optional) (default to "[]")
	queryTime := "queryTime_example" // string | from which UNIX time to extract metrics (optional)
	duration := "duration_example" // string | duration of the query period, in seconds (optional) (default to "1800")
	step := "step_example" // string | step between graph data points, in seconds. (optional) (default to "15")
	rateInterval := "rateInterval_example" // string | metrics rate intervals, e.g. 20s (optional) (default to "1m")
	direction := "direction_example" // string | traffic direction: 'inbound' or 'outbound' (optional) (default to "outbound")
	quantiles := "quantiles_example" // string | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 (optional) (default to "[]")
	byLabels := "byLabels_example" // string | list of labels to use for grouping metrics(via Prometheus 'by' clause), e.g. source_workload, destination_service_name (optional) (default to "[]")
	requestProtocol := "requestProtocol_example" // string | request protocol for the telemetry, e.g. http/tcp/grpc (optional) (default to "all protocols")
	reporter := "reporter_example" // string | istio telemetry reporter, 'source' or 'destination' (optional) (default to "source")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetAppMetrics(context.Background(), namespace, app).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetAppMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetAppMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of the namespace | 
**app** | **string** | name of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count | [default to &quot;[]&quot;]
 **queryTime** | **string** | from which UNIX time to extract metrics | 
 **duration** | **string** | duration of the query period, in seconds | [default to &quot;1800&quot;]
 **step** | **string** | step between graph data points, in seconds. | [default to &quot;15&quot;]
 **rateInterval** | **string** | metrics rate intervals, e.g. 20s | [default to &quot;1m&quot;]
 **direction** | **string** | traffic direction: &#39;inbound&#39; or &#39;outbound&#39; | [default to &quot;outbound&quot;]
 **quantiles** | **string** | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 | [default to &quot;[]&quot;]
 **byLabels** | **string** | list of labels to use for grouping metrics(via Prometheus &#39;by&#39; clause), e.g. source_workload, destination_service_name | [default to &quot;[]&quot;]
 **requestProtocol** | **string** | request protocol for the telemetry, e.g. http/tcp/grpc | [default to &quot;all protocols&quot;]
 **reporter** | **string** | istio telemetry reporter, &#39;source&#39; or &#39;destination&#39; | [default to &quot;source&quot;]

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceGraph

> map[string]interface{} GetNamespaceGraph(ctx, namespace).Duration(duration).GraphType(graphType).GroupBy(groupBy).QueryTime(queryTime).InjectServiceNodes(injectServiceNodes).Execute()

Get service graph for a specific namespace

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
	namespace := "namespace_example" // string | name of a namespace
	duration := "duration_example" // string | duration of the query period, in seconds (optional) (default to "10m")
	graphType := "graphType_example" // string | type of the generated service graph. Available graph types: [app, service, versionedApp, workload]. (optional) (default to "workload")
	groupBy := "groupBy_example" // string | app box grouping characteristic. Available groupings: [app, none, version]. (optional) (default to "none")
	queryTime := "queryTime_example" // string | from which time point in UNIX timestamp, default now (optional)
	injectServiceNodes := "injectServiceNodes_example" // string | flag for injecting the requested service node between source and destination nodes. (optional) (default to "false")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetNamespaceGraph(context.Background(), namespace).Duration(duration).GraphType(graphType).GroupBy(groupBy).QueryTime(queryTime).InjectServiceNodes(injectServiceNodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetNamespaceGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceGraph`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetNamespaceGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of a namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **string** | duration of the query period, in seconds | [default to &quot;10m&quot;]
 **graphType** | **string** | type of the generated service graph. Available graph types: [app, service, versionedApp, workload]. | [default to &quot;workload&quot;]
 **groupBy** | **string** | app box grouping characteristic. Available groupings: [app, none, version]. | [default to &quot;none&quot;]
 **queryTime** | **string** | from which time point in UNIX timestamp, default now | 
 **injectServiceNodes** | **string** | flag for injecting the requested service node between source and destination nodes. | [default to &quot;false&quot;]

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceHealth

> map[string]interface{} GetNamespaceHealth(ctx, namespace).RateInterval(rateInterval).QueryTime(queryTime).Execute()

Get app/service/workload health of a namespace

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
	namespace := "namespace_example" // string | name of a namespace
	rateInterval := "rateInterval_example" // string | the rate interval used for fetching error rate (default to "10m")
	queryTime := "queryTime_example" // string | the time to use for query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetNamespaceHealth(context.Background(), namespace).RateInterval(rateInterval).QueryTime(queryTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetNamespaceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetNamespaceHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of a namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rateInterval** | **string** | the rate interval used for fetching error rate | [default to &quot;10m&quot;]
 **queryTime** | **string** | the time to use for query | 

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceMetrics

> map[string]interface{} GetNamespaceMetrics(ctx, namespace).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()

Get metrics from a specific namespace

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
	namespace := "namespace_example" // string | name of the namespace
	filters := "filters_example" // string | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count (optional) (default to "[]")
	queryTime := "queryTime_example" // string | from which UNIX time to extract metrics (optional)
	duration := "duration_example" // string | duration of the query period, in seconds (optional) (default to "1800")
	step := "step_example" // string | step between graph data points, in seconds. (optional) (default to "15")
	rateInterval := "rateInterval_example" // string | metrics rate intervals, e.g. 20s (optional) (default to "1m")
	direction := "direction_example" // string | traffic direction: 'inbound' or 'outbound' (optional) (default to "outbound")
	quantiles := "quantiles_example" // string | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 (optional) (default to "[]")
	byLabels := "byLabels_example" // string | list of labels to use for grouping metrics(via Prometheus 'by' clause), e.g. source_workload, destination_service_name (optional) (default to "[]")
	requestProtocol := "requestProtocol_example" // string | request protocol for the telemetry, e.g. http/tcp/grpc (optional) (default to "all protocols")
	reporter := "reporter_example" // string | istio telemetry reporter, 'source' or 'destination' (optional) (default to "source")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetNamespaceMetrics(context.Background(), namespace).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetNamespaceMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetNamespaceMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count | [default to &quot;[]&quot;]
 **queryTime** | **string** | from which UNIX time to extract metrics | 
 **duration** | **string** | duration of the query period, in seconds | [default to &quot;1800&quot;]
 **step** | **string** | step between graph data points, in seconds. | [default to &quot;15&quot;]
 **rateInterval** | **string** | metrics rate intervals, e.g. 20s | [default to &quot;1m&quot;]
 **direction** | **string** | traffic direction: &#39;inbound&#39; or &#39;outbound&#39; | [default to &quot;outbound&quot;]
 **quantiles** | **string** | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 | [default to &quot;[]&quot;]
 **byLabels** | **string** | list of labels to use for grouping metrics(via Prometheus &#39;by&#39; clause), e.g. source_workload, destination_service_name | [default to &quot;[]&quot;]
 **requestProtocol** | **string** | request protocol for the telemetry, e.g. http/tcp/grpc | [default to &quot;all protocols&quot;]
 **reporter** | **string** | istio telemetry reporter, &#39;source&#39; or &#39;destination&#39; | [default to &quot;source&quot;]

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceHealth

> map[string]interface{} GetServiceHealth(ctx, namespace, service).RateInterval(rateInterval).QueryTime(queryTime).Execute()

Get service health

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
	namespace := "namespace_example" // string | name of a namespace
	service := "service_example" // string | service name
	rateInterval := "rateInterval_example" // string | the rate interval used for fetching error rate (default to "10m")
	queryTime := "queryTime_example" // string | the time to use for query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetServiceHealth(context.Background(), namespace, service).RateInterval(rateInterval).QueryTime(queryTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetServiceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetServiceHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of a namespace | 
**service** | **string** | service name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rateInterval** | **string** | the rate interval used for fetching error rate | [default to &quot;10m&quot;]
 **queryTime** | **string** | the time to use for query | 

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceMetrics

> map[string]interface{} GetServiceMetrics(ctx, namespace, service).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()

Get service metrics from a specific namespace

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
	namespace := "namespace_example" // string | name of the namespace
	service := "service_example" // string | name of the service
	filters := "filters_example" // string | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count (optional) (default to "[]")
	queryTime := "queryTime_example" // string | from which UNIX time to extract metrics (optional)
	duration := "duration_example" // string | duration of the query period, in seconds (optional) (default to "1800")
	step := "step_example" // string | step between graph data points, in seconds. (optional) (default to "15")
	rateInterval := "rateInterval_example" // string | metrics rate intervals, e.g. 20s (optional) (default to "1m")
	direction := "direction_example" // string | traffic direction: 'inbound' or 'outbound' (optional) (default to "outbound")
	quantiles := "quantiles_example" // string | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 (optional) (default to "[]")
	byLabels := "byLabels_example" // string | list of labels to use for grouping metrics(via Prometheus 'by' clause), e.g. source_workload, destination_service_name (optional) (default to "[]")
	requestProtocol := "requestProtocol_example" // string | request protocol for the telemetry, e.g. http/tcp/grpc (optional) (default to "all protocols")
	reporter := "reporter_example" // string | istio telemetry reporter, 'source' or 'destination' (optional) (default to "source")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetServiceMetrics(context.Background(), namespace, service).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetServiceMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetServiceMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of the namespace | 
**service** | **string** | name of the service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count | [default to &quot;[]&quot;]
 **queryTime** | **string** | from which UNIX time to extract metrics | 
 **duration** | **string** | duration of the query period, in seconds | [default to &quot;1800&quot;]
 **step** | **string** | step between graph data points, in seconds. | [default to &quot;15&quot;]
 **rateInterval** | **string** | metrics rate intervals, e.g. 20s | [default to &quot;1m&quot;]
 **direction** | **string** | traffic direction: &#39;inbound&#39; or &#39;outbound&#39; | [default to &quot;outbound&quot;]
 **quantiles** | **string** | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 | [default to &quot;[]&quot;]
 **byLabels** | **string** | list of labels to use for grouping metrics(via Prometheus &#39;by&#39; clause), e.g. source_workload, destination_service_name | [default to &quot;[]&quot;]
 **requestProtocol** | **string** | request protocol for the telemetry, e.g. http/tcp/grpc | [default to &quot;all protocols&quot;]
 **reporter** | **string** | istio telemetry reporter, &#39;source&#39; or &#39;destination&#39; | [default to &quot;source&quot;]

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTracing

> GetServiceTracing(ctx, namespace, service).Start(start).End(end).Limit(limit).Loopback(loopback).MaxDuration(maxDuration).MinDuration(minDuration).Execute()

Get tracing of a service, should have servicemesh enabled first

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
	namespace := "namespace_example" // string | namespace of service
	service := "service_example" // string | name of service queried
	start := "start_example" // string | start of time range want to query, in unix timestamp (optional)
	end := "end_example" // string | end of time range want to query, in unix timestamp (optional)
	limit := "limit_example" // string | maximum tracing entries returned at one query, default 10 (optional) (default to "10")
	loopback := "loopback_example" // string | loopback of duration want to query, e.g. 30m/1h/2d (optional)
	maxDuration := "maxDuration_example" // string | maximum duration of a request (optional)
	minDuration := "minDuration_example" // string | minimum duration of a request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceMeshAPI.GetServiceTracing(context.Background(), namespace, service).Start(start).End(end).Limit(limit).Loopback(loopback).MaxDuration(maxDuration).MinDuration(minDuration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetServiceTracing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace of service | 
**service** | **string** | name of service queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTracingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **string** | start of time range want to query, in unix timestamp | 
 **end** | **string** | end of time range want to query, in unix timestamp | 
 **limit** | **string** | maximum tracing entries returned at one query, default 10 | [default to &quot;10&quot;]
 **loopback** | **string** | loopback of duration want to query, e.g. 30m/1h/2d | 
 **maxDuration** | **string** | maximum duration of a request | 
 **minDuration** | **string** | minimum duration of a request | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloadHealth

> map[string]interface{} GetWorkloadHealth(ctx, namespace, workload).RateInterval(rateInterval).QueryTime(queryTime).Execute()

Get workload health

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
	namespace := "namespace_example" // string | name of a namespace
	workload := "workload_example" // string | workload name
	rateInterval := "rateInterval_example" // string | the rate interval used for fetching error rate (default to "10m")
	queryTime := "queryTime_example" // string | the time to use for query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetWorkloadHealth(context.Background(), namespace, workload).RateInterval(rateInterval).QueryTime(queryTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetWorkloadHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetWorkloadHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of a namespace | 
**workload** | **string** | workload name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rateInterval** | **string** | the rate interval used for fetching error rate | [default to &quot;10m&quot;]
 **queryTime** | **string** | the time to use for query | 

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloadMetrics

> map[string]interface{} GetWorkloadMetrics(ctx, namespace, workload).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()

Get workload metrics from a specific namespace

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
	namespace := "namespace_example" // string | name of the namespace
	workload := "workload_example" // string | name of the workload
	filters := "filters_example" // string | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count (optional) (default to "[]")
	queryTime := "queryTime_example" // string | from which UNIX time to extract metrics (optional)
	duration := "duration_example" // string | duration of the query period, in seconds (optional) (default to "1800")
	step := "step_example" // string | step between graph data points, in seconds. (optional) (default to "15")
	rateInterval := "rateInterval_example" // string | metrics rate intervals, e.g. 20s (optional) (default to "1m")
	direction := "direction_example" // string | traffic direction: 'inbound' or 'outbound' (optional) (default to "outbound")
	quantiles := "quantiles_example" // string | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 (optional) (default to "[]")
	byLabels := "byLabels_example" // string | list of labels to use for grouping metrics(via Prometheus 'by' clause), e.g. source_workload, destination_service_name (optional) (default to "[]")
	requestProtocol := "requestProtocol_example" // string | request protocol for the telemetry, e.g. http/tcp/grpc (optional) (default to "all protocols")
	reporter := "reporter_example" // string | istio telemetry reporter, 'source' or 'destination' (optional) (default to "source")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMeshAPI.GetWorkloadMetrics(context.Background(), namespace, workload).Filters(filters).QueryTime(queryTime).Duration(duration).Step(step).RateInterval(rateInterval).Direction(direction).Quantiles(quantiles).ByLabels(byLabels).RequestProtocol(requestProtocol).Reporter(reporter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMeshAPI.GetWorkloadMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceMeshAPI.GetWorkloadMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of the namespace | 
**workload** | **string** | name of the workload | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** | type of metrics type, fetch all metrics when empty, e.g. request_count, request_duration, request_error_count | [default to &quot;[]&quot;]
 **queryTime** | **string** | from which UNIX time to extract metrics | 
 **duration** | **string** | duration of the query period, in seconds | [default to &quot;1800&quot;]
 **step** | **string** | step between graph data points, in seconds. | [default to &quot;15&quot;]
 **rateInterval** | **string** | metrics rate intervals, e.g. 20s | [default to &quot;1m&quot;]
 **direction** | **string** | traffic direction: &#39;inbound&#39; or &#39;outbound&#39; | [default to &quot;outbound&quot;]
 **quantiles** | **string** | list of quantiles to fetch, fetch no quantiles when empty. eg. 0.5, 0.9, 0.99 | [default to &quot;[]&quot;]
 **byLabels** | **string** | list of labels to use for grouping metrics(via Prometheus &#39;by&#39; clause), e.g. source_workload, destination_service_name | [default to &quot;[]&quot;]
 **requestProtocol** | **string** | request protocol for the telemetry, e.g. http/tcp/grpc | [default to &quot;all protocols&quot;]
 **reporter** | **string** | istio telemetry reporter, &#39;source&#39; or &#39;destination&#39; | [default to &quot;source&quot;]

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

