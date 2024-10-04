# \EventsQueryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Events**](EventsQueryAPI.md#Events) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/events | Query events against the cluster



## Events

> V1alpha1APIResponse Events(ctx).Operation(operation).WorkspaceFilter(workspaceFilter).WorkspaceSearch(workspaceSearch).InvolvedObjectNamespaceFilter(involvedObjectNamespaceFilter).InvolvedObjectNamespaceSearch(involvedObjectNamespaceSearch).InvolvedObjectNameFilter(involvedObjectNameFilter).InvolvedObjectNameSearch(involvedObjectNameSearch).InvolvedObjectKindFilter(involvedObjectKindFilter).ReasonFilter(reasonFilter).ReasonSearch(reasonSearch).MessageSearch(messageSearch).TypeFilter(typeFilter).StartTime(startTime).EndTime(endTime).Interval(interval).Sort(sort).From(from).Size(size).Execute()

Query events against the cluster

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
	operation := "operation_example" // string | Operation type. This can be one of three types: `query` (for querying events), `statistics` (for retrieving statistical data), `histogram` (for displaying events count by time interval). Defaults to query. (optional) (default to "query")
	workspaceFilter := "workspaceFilter_example" // string | A comma-separated list of workspaces. This field restricts the query to specified workspaces. For example, the following filter matches the workspace my-ws and demo-ws: `my-ws,demo-ws`. (optional)
	workspaceSearch := "workspaceSearch_example" // string | A comma-separated list of keywords. Differing from **workspace_filter**, this field performs fuzzy matching on workspaces. For example, the following value limits the query to workspaces whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`. (optional)
	involvedObjectNamespaceFilter := "involvedObjectNamespaceFilter_example" // string | A comma-separated list of namespaces. This field restricts the query to specified `involvedObject.namespace`. (optional)
	involvedObjectNamespaceSearch := "involvedObjectNamespaceSearch_example" // string | A comma-separated list of keywords. Differing from **involved_object_namespace_filter**, this field performs fuzzy matching on `involvedObject.namespace` (optional)
	involvedObjectNameFilter := "involvedObjectNameFilter_example" // string | A comma-separated list of names. This field restricts the query to specified `involvedObject.name`. (optional)
	involvedObjectNameSearch := "involvedObjectNameSearch_example" // string | A comma-separated list of keywords. Differing from **involved_object_name_filter**, this field performs fuzzy matching on `involvedObject.name`. (optional)
	involvedObjectKindFilter := "involvedObjectKindFilter_example" // string | A comma-separated list of kinds. This field restricts the query to specified `involvedObject.kind`. (optional)
	reasonFilter := "reasonFilter_example" // string | A comma-separated list of reasons. This field restricts the query to specified `reason`. (optional)
	reasonSearch := "reasonSearch_example" // string | A comma-separated list of keywords. Differing from **reason_filter**, this field performs fuzzy matching on `reason`. (optional)
	messageSearch := "messageSearch_example" // string | A comma-separated list of keywords. This field performs fuzzy matching on `message`. (optional)
	typeFilter := "typeFilter_example" // string | Type of event matching on `type`. This can be one of two types: `Warning`, `Normal` (optional)
	startTime := "startTime_example" // string | Start time of query (limits `lastTimestamp`). The format is a string representing seconds since the epoch, eg. 1136214245. (optional)
	endTime := "endTime_example" // string | End time of query (limits `lastTimestamp`). The format is a string representing seconds since the epoch, eg. 1136214245. (optional)
	interval := "interval_example" // string | Time interval. It requires **operation** is set to `histogram`. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min). (optional) (default to "15m")
	sort := "sort_example" // string | Sort order. One of asc, desc. This field sorts events by `lastTimestamp`. (optional) (default to "desc")
	from := int32(56) // int32 | The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to `query`. Defaults to 0 (i.e. from the beginning of the result set). (optional) (default to 0)
	size := int32(56) // int32 | Size of result set to return. It requires **operation** is set to `query`. Defaults to 10 (i.e. 10 event records). (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsQueryAPI.Events(context.Background()).Operation(operation).WorkspaceFilter(workspaceFilter).WorkspaceSearch(workspaceSearch).InvolvedObjectNamespaceFilter(involvedObjectNamespaceFilter).InvolvedObjectNamespaceSearch(involvedObjectNamespaceSearch).InvolvedObjectNameFilter(involvedObjectNameFilter).InvolvedObjectNameSearch(involvedObjectNameSearch).InvolvedObjectKindFilter(involvedObjectKindFilter).ReasonFilter(reasonFilter).ReasonSearch(reasonSearch).MessageSearch(messageSearch).TypeFilter(typeFilter).StartTime(startTime).EndTime(endTime).Interval(interval).Sort(sort).From(from).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsQueryAPI.Events``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Events`: V1alpha1APIResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsQueryAPI.Events`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operation** | **string** | Operation type. This can be one of three types: &#x60;query&#x60; (for querying events), &#x60;statistics&#x60; (for retrieving statistical data), &#x60;histogram&#x60; (for displaying events count by time interval). Defaults to query. | [default to &quot;query&quot;]
 **workspaceFilter** | **string** | A comma-separated list of workspaces. This field restricts the query to specified workspaces. For example, the following filter matches the workspace my-ws and demo-ws: &#x60;my-ws,demo-ws&#x60;. | 
 **workspaceSearch** | **string** | A comma-separated list of keywords. Differing from **workspace_filter**, this field performs fuzzy matching on workspaces. For example, the following value limits the query to workspaces whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **involvedObjectNamespaceFilter** | **string** | A comma-separated list of namespaces. This field restricts the query to specified &#x60;involvedObject.namespace&#x60;. | 
 **involvedObjectNamespaceSearch** | **string** | A comma-separated list of keywords. Differing from **involved_object_namespace_filter**, this field performs fuzzy matching on &#x60;involvedObject.namespace&#x60; | 
 **involvedObjectNameFilter** | **string** | A comma-separated list of names. This field restricts the query to specified &#x60;involvedObject.name&#x60;. | 
 **involvedObjectNameSearch** | **string** | A comma-separated list of keywords. Differing from **involved_object_name_filter**, this field performs fuzzy matching on &#x60;involvedObject.name&#x60;. | 
 **involvedObjectKindFilter** | **string** | A comma-separated list of kinds. This field restricts the query to specified &#x60;involvedObject.kind&#x60;. | 
 **reasonFilter** | **string** | A comma-separated list of reasons. This field restricts the query to specified &#x60;reason&#x60;. | 
 **reasonSearch** | **string** | A comma-separated list of keywords. Differing from **reason_filter**, this field performs fuzzy matching on &#x60;reason&#x60;. | 
 **messageSearch** | **string** | A comma-separated list of keywords. This field performs fuzzy matching on &#x60;message&#x60;. | 
 **typeFilter** | **string** | Type of event matching on &#x60;type&#x60;. This can be one of two types: &#x60;Warning&#x60;, &#x60;Normal&#x60; | 
 **startTime** | **string** | Start time of query (limits &#x60;lastTimestamp&#x60;). The format is a string representing seconds since the epoch, eg. 1136214245. | 
 **endTime** | **string** | End time of query (limits &#x60;lastTimestamp&#x60;). The format is a string representing seconds since the epoch, eg. 1136214245. | 
 **interval** | **string** | Time interval. It requires **operation** is set to &#x60;histogram&#x60;. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min). | [default to &quot;15m&quot;]
 **sort** | **string** | Sort order. One of asc, desc. This field sorts events by &#x60;lastTimestamp&#x60;. | [default to &quot;desc&quot;]
 **from** | **int32** | The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to &#x60;query&#x60;. Defaults to 0 (i.e. from the beginning of the result set). | [default to 0]
 **size** | **int32** | Size of result set to return. It requires **operation** is set to &#x60;query&#x60;. Defaults to 10 (i.e. 10 event records). | [default to 10]

### Return type

[**V1alpha1APIResponse**](V1alpha1APIResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

