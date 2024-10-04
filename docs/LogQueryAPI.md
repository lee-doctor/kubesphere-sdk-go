# \LogQueryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryLogs**](LogQueryAPI.md#QueryLogs) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/logs | Query logs against the cluster.



## QueryLogs

> V1alpha2APIResponse QueryLogs(ctx).Operation(operation).Namespaces(namespaces).NamespaceQuery(namespaceQuery).Workloads(workloads).WorkloadQuery(workloadQuery).Pods(pods).PodQuery(podQuery).Containers(containers).ContainerQuery(containerQuery).LogQuery(logQuery).Interval(interval).StartTime(startTime).EndTime(endTime).Sort(sort).From(from).Size(size).Execute()

Query logs against the cluster.

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
	operation := "operation_example" // string | Operation type. This can be one of four types: query (for querying logs), statistics (for retrieving statistical data), histogram (for displaying log count by time interval) and export (for exporting logs). Defaults to query. (optional) (default to "query")
	namespaces := "namespaces_example" // string | A comma-separated list of namespaces. This field restricts the query to specified namespaces. For example, the following filter matches the namespace my-ns and demo-ns: `my-ns,demo-ns` (optional)
	namespaceQuery := "namespaceQuery_example" // string | A comma-separated list of keywords. Differing from **namespaces**, this field performs fuzzy matching on namespaces. For example, the following value limits the query to namespaces whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`. (optional)
	workloads := "workloads_example" // string | A comma-separated list of workloads. This field restricts the query to specified workloads. For example, the following filter matches the workload my-wl and demo-wl: `my-wl,demo-wl` (optional)
	workloadQuery := "workloadQuery_example" // string | A comma-separated list of keywords. Differing from **workloads**, this field performs fuzzy matching on workloads. For example, the following value limits the query to workloads whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`. (optional)
	pods := "pods_example" // string | A comma-separated list of pods. This field restricts the query to specified pods. For example, the following filter matches the pod my-po and demo-po: `my-po,demo-po` (optional)
	podQuery := "podQuery_example" // string | A comma-separated list of keywords. Differing from **pods**, this field performs fuzzy matching on pods. For example, the following value limits the query to pods whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`. (optional)
	containers := "containers_example" // string | A comma-separated list of containers. This field restricts the query to specified containers. For example, the following filter matches the container my-cont and demo-cont: `my-cont,demo-cont` (optional)
	containerQuery := "containerQuery_example" // string | A comma-separated list of keywords. Differing from **containers**, this field performs fuzzy matching on containers. For example, the following value limits the query to containers whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`. (optional)
	logQuery := "logQuery_example" // string | A comma-separated list of keywords. The query returns logs which contain at least one keyword. Case-insensitive matching. For example, if the field is set to `err,INFO`, the query returns any log containing err(ERR,Err,...) *OR* INFO(info,InFo,...). (optional)
	interval := "interval_example" // string | Time interval. It requires **operation** is set to histogram. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min). (optional) (default to "15m")
	startTime := "startTime_example" // string | Start time of query. Default to 0. The format is a string representing seconds since the epoch, eg. 1559664000. (optional)
	endTime := "endTime_example" // string | End time of query. Default to now. The format is a string representing seconds since the epoch, eg. 1559664000. (optional)
	sort := "sort_example" // string | Sort order. One of asc, desc. This field sorts logs by timestamp. (optional) (default to "desc")
	from := int32(56) // int32 | The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to query. Defaults to 0 (i.e. from the beginning of the result set). (optional) (default to 0)
	size := int32(56) // int32 | Size of result to return. It requires **operation** is set to query. Defaults to 10 (i.e. 10 log records). (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogQueryAPI.QueryLogs(context.Background()).Operation(operation).Namespaces(namespaces).NamespaceQuery(namespaceQuery).Workloads(workloads).WorkloadQuery(workloadQuery).Pods(pods).PodQuery(podQuery).Containers(containers).ContainerQuery(containerQuery).LogQuery(logQuery).Interval(interval).StartTime(startTime).EndTime(endTime).Sort(sort).From(from).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogQueryAPI.QueryLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryLogs`: V1alpha2APIResponse
	fmt.Fprintf(os.Stdout, "Response from `LogQueryAPI.QueryLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operation** | **string** | Operation type. This can be one of four types: query (for querying logs), statistics (for retrieving statistical data), histogram (for displaying log count by time interval) and export (for exporting logs). Defaults to query. | [default to &quot;query&quot;]
 **namespaces** | **string** | A comma-separated list of namespaces. This field restricts the query to specified namespaces. For example, the following filter matches the namespace my-ns and demo-ns: &#x60;my-ns,demo-ns&#x60; | 
 **namespaceQuery** | **string** | A comma-separated list of keywords. Differing from **namespaces**, this field performs fuzzy matching on namespaces. For example, the following value limits the query to namespaces whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **workloads** | **string** | A comma-separated list of workloads. This field restricts the query to specified workloads. For example, the following filter matches the workload my-wl and demo-wl: &#x60;my-wl,demo-wl&#x60; | 
 **workloadQuery** | **string** | A comma-separated list of keywords. Differing from **workloads**, this field performs fuzzy matching on workloads. For example, the following value limits the query to workloads whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **pods** | **string** | A comma-separated list of pods. This field restricts the query to specified pods. For example, the following filter matches the pod my-po and demo-po: &#x60;my-po,demo-po&#x60; | 
 **podQuery** | **string** | A comma-separated list of keywords. Differing from **pods**, this field performs fuzzy matching on pods. For example, the following value limits the query to pods whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **containers** | **string** | A comma-separated list of containers. This field restricts the query to specified containers. For example, the following filter matches the container my-cont and demo-cont: &#x60;my-cont,demo-cont&#x60; | 
 **containerQuery** | **string** | A comma-separated list of keywords. Differing from **containers**, this field performs fuzzy matching on containers. For example, the following value limits the query to containers whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **logQuery** | **string** | A comma-separated list of keywords. The query returns logs which contain at least one keyword. Case-insensitive matching. For example, if the field is set to &#x60;err,INFO&#x60;, the query returns any log containing err(ERR,Err,...) *OR* INFO(info,InFo,...). | 
 **interval** | **string** | Time interval. It requires **operation** is set to histogram. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min). | [default to &quot;15m&quot;]
 **startTime** | **string** | Start time of query. Default to 0. The format is a string representing seconds since the epoch, eg. 1559664000. | 
 **endTime** | **string** | End time of query. Default to now. The format is a string representing seconds since the epoch, eg. 1559664000. | 
 **sort** | **string** | Sort order. One of asc, desc. This field sorts logs by timestamp. | [default to &quot;desc&quot;]
 **from** | **int32** | The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to query. Defaults to 0 (i.e. from the beginning of the result set). | [default to 0]
 **size** | **int32** | Size of result to return. It requires **operation** is set to query. Defaults to 10 (i.e. 10 log records). | [default to 10]

### Return type

[**V1alpha2APIResponse**](V1alpha2APIResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

