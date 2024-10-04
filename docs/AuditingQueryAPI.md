# \AuditingQueryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Auditing**](AuditingQueryAPI.md#Auditing) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/auditing/events | Query auditing events against the cluster



## Auditing

> V1alpha1APIResponse Auditing(ctx).Operation(operation).WorkspaceFilter(workspaceFilter).WorkspaceSearch(workspaceSearch).ObjectrefNamespaceFilter(objectrefNamespaceFilter).ObjectrefNamespaceSearch(objectrefNamespaceSearch).ObjectrefNameFilter(objectrefNameFilter).ObjectrefNameSearch(objectrefNameSearch).LevelFilter(levelFilter).VerbFilter(verbFilter).UserFilter(userFilter).UserSearch(userSearch).GroupSearch(groupSearch).SourceIpSearch(sourceIpSearch).ObjectrefResourceFilter(objectrefResourceFilter).ObjectrefSubresourceFilter(objectrefSubresourceFilter).ResponseCodeFilter(responseCodeFilter).ResponseStatusFilter(responseStatusFilter).StartTime(startTime).EndTime(endTime).Interval(interval).Sort(sort).From(from).Size(size).Execute()

Query auditing events against the cluster

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
	objectrefNamespaceFilter := "objectrefNamespaceFilter_example" // string | A comma-separated list of namespaces. This field restricts the query to specified `ObjectRef.Namespace`. (optional)
	objectrefNamespaceSearch := "objectrefNamespaceSearch_example" // string | A comma-separated list of keywords. Differing from **objectref_namespace_filter**, this field performs fuzzy matching on `ObjectRef.Namespace`. (optional)
	objectrefNameFilter := "objectrefNameFilter_example" // string | A comma-separated list of names. This field restricts the query to specified `ObjectRef.Name`. (optional)
	objectrefNameSearch := "objectrefNameSearch_example" // string | A comma-separated list of keywords. Differing from **objectref_name_filter**, this field performs fuzzy matching on `ObjectRef.Name`. (optional)
	levelFilter := "levelFilter_example" // string | A comma-separated list of levels. This know values are Metadata, Request, RequestResponse. (optional)
	verbFilter := "verbFilter_example" // string | A comma-separated list of verbs. This field restricts the query to specified verb. This field restricts the query to specified `Verb`. (optional)
	userFilter := "userFilter_example" // string | A comma-separated list of user. This field restricts the query to specified user. For example, the following filter matches the user user1 and user2: `user1,user2`. (optional)
	userSearch := "userSearch_example" // string | A comma-separated list of keywords. Differing from **user_filter**, this field performs fuzzy matching on 'User.username'. For example, the following value limits the query to user whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`. (optional)
	groupSearch := "groupSearch_example" // string | A comma-separated list of keywords. This field performs fuzzy matching on 'User.Groups'. For example, the following value limits the query to group which contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): `my,demo`. (optional)
	sourceIpSearch := "sourceIpSearch_example" // string | A comma-separated list of keywords. This field performs fuzzy matching on 'SourceIPs'. For example, the following value limits the query to SourceIPs which contains 127.0 *OR* 192.168.: `127.0,192.168.`. (optional)
	objectrefResourceFilter := "objectrefResourceFilter_example" // string | A comma-separated list of resource. This field restricts the query to specified ip. This field restricts the query to specified `ObjectRef.Resource`. (optional)
	objectrefSubresourceFilter := "objectrefSubresourceFilter_example" // string | A comma-separated list of subresource. This field restricts the query to specified subresource. This field restricts the query to specified `ObjectRef.Subresource`. (optional)
	responseCodeFilter := "responseCodeFilter_example" // string | A comma-separated list of response status code. This field restricts the query to specified response status code. This field restricts the query to specified `ResponseStatus.code`. (optional)
	responseStatusFilter := "responseStatusFilter_example" // string | A comma-separated list of response status. This field restricts the query to specified response status. This field restricts the query to specified `ResponseStatus.status`. (optional)
	startTime := "startTime_example" // string | Start time of query (limits `RequestReceivedTimestamp`). The format is a string representing seconds since the epoch, eg. 1136214245. (optional)
	endTime := "endTime_example" // string | End time of query (limits `RequestReceivedTimestamp`). The format is a string representing seconds since the epoch, eg. 1136214245. (optional)
	interval := "interval_example" // string | Time interval. It requires **operation** is set to `histogram`. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min). (optional) (default to "15m")
	sort := "sort_example" // string | Sort order. One of asc, desc. This field sorts events by `RequestReceivedTimestamp`. (optional) (default to "desc")
	from := int32(56) // int32 | The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to `query`. Defaults to 0 (i.e. from the beginning of the result set). (optional) (default to 0)
	size := int32(56) // int32 | Size of result set to return. It requires **operation** is set to `query`. Defaults to 10 (i.e. 10 event records). (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditingQueryAPI.Auditing(context.Background()).Operation(operation).WorkspaceFilter(workspaceFilter).WorkspaceSearch(workspaceSearch).ObjectrefNamespaceFilter(objectrefNamespaceFilter).ObjectrefNamespaceSearch(objectrefNamespaceSearch).ObjectrefNameFilter(objectrefNameFilter).ObjectrefNameSearch(objectrefNameSearch).LevelFilter(levelFilter).VerbFilter(verbFilter).UserFilter(userFilter).UserSearch(userSearch).GroupSearch(groupSearch).SourceIpSearch(sourceIpSearch).ObjectrefResourceFilter(objectrefResourceFilter).ObjectrefSubresourceFilter(objectrefSubresourceFilter).ResponseCodeFilter(responseCodeFilter).ResponseStatusFilter(responseStatusFilter).StartTime(startTime).EndTime(endTime).Interval(interval).Sort(sort).From(from).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditingQueryAPI.Auditing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Auditing`: V1alpha1APIResponse
	fmt.Fprintf(os.Stdout, "Response from `AuditingQueryAPI.Auditing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operation** | **string** | Operation type. This can be one of three types: &#x60;query&#x60; (for querying events), &#x60;statistics&#x60; (for retrieving statistical data), &#x60;histogram&#x60; (for displaying events count by time interval). Defaults to query. | [default to &quot;query&quot;]
 **workspaceFilter** | **string** | A comma-separated list of workspaces. This field restricts the query to specified workspaces. For example, the following filter matches the workspace my-ws and demo-ws: &#x60;my-ws,demo-ws&#x60;. | 
 **workspaceSearch** | **string** | A comma-separated list of keywords. Differing from **workspace_filter**, this field performs fuzzy matching on workspaces. For example, the following value limits the query to workspaces whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **objectrefNamespaceFilter** | **string** | A comma-separated list of namespaces. This field restricts the query to specified &#x60;ObjectRef.Namespace&#x60;. | 
 **objectrefNamespaceSearch** | **string** | A comma-separated list of keywords. Differing from **objectref_namespace_filter**, this field performs fuzzy matching on &#x60;ObjectRef.Namespace&#x60;. | 
 **objectrefNameFilter** | **string** | A comma-separated list of names. This field restricts the query to specified &#x60;ObjectRef.Name&#x60;. | 
 **objectrefNameSearch** | **string** | A comma-separated list of keywords. Differing from **objectref_name_filter**, this field performs fuzzy matching on &#x60;ObjectRef.Name&#x60;. | 
 **levelFilter** | **string** | A comma-separated list of levels. This know values are Metadata, Request, RequestResponse. | 
 **verbFilter** | **string** | A comma-separated list of verbs. This field restricts the query to specified verb. This field restricts the query to specified &#x60;Verb&#x60;. | 
 **userFilter** | **string** | A comma-separated list of user. This field restricts the query to specified user. For example, the following filter matches the user user1 and user2: &#x60;user1,user2&#x60;. | 
 **userSearch** | **string** | A comma-separated list of keywords. Differing from **user_filter**, this field performs fuzzy matching on &#39;User.username&#39;. For example, the following value limits the query to user whose name contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **groupSearch** | **string** | A comma-separated list of keywords. This field performs fuzzy matching on &#39;User.Groups&#39;. For example, the following value limits the query to group which contains the word my(My,MY,...) *OR* demo(Demo,DemO,...): &#x60;my,demo&#x60;. | 
 **sourceIpSearch** | **string** | A comma-separated list of keywords. This field performs fuzzy matching on &#39;SourceIPs&#39;. For example, the following value limits the query to SourceIPs which contains 127.0 *OR* 192.168.: &#x60;127.0,192.168.&#x60;. | 
 **objectrefResourceFilter** | **string** | A comma-separated list of resource. This field restricts the query to specified ip. This field restricts the query to specified &#x60;ObjectRef.Resource&#x60;. | 
 **objectrefSubresourceFilter** | **string** | A comma-separated list of subresource. This field restricts the query to specified subresource. This field restricts the query to specified &#x60;ObjectRef.Subresource&#x60;. | 
 **responseCodeFilter** | **string** | A comma-separated list of response status code. This field restricts the query to specified response status code. This field restricts the query to specified &#x60;ResponseStatus.code&#x60;. | 
 **responseStatusFilter** | **string** | A comma-separated list of response status. This field restricts the query to specified response status. This field restricts the query to specified &#x60;ResponseStatus.status&#x60;. | 
 **startTime** | **string** | Start time of query (limits &#x60;RequestReceivedTimestamp&#x60;). The format is a string representing seconds since the epoch, eg. 1136214245. | 
 **endTime** | **string** | End time of query (limits &#x60;RequestReceivedTimestamp&#x60;). The format is a string representing seconds since the epoch, eg. 1136214245. | 
 **interval** | **string** | Time interval. It requires **operation** is set to &#x60;histogram&#x60;. The format is [0-9]+[smhdwMqy]. Defaults to 15m (i.e. 15 min). | [default to &quot;15m&quot;]
 **sort** | **string** | Sort order. One of asc, desc. This field sorts events by &#x60;RequestReceivedTimestamp&#x60;. | [default to &quot;desc&quot;]
 **from** | **int32** | The offset from the result set. This field returns query results from the specified offset. It requires **operation** is set to &#x60;query&#x60;. Defaults to 0 (i.e. from the beginning of the result set). | [default to 0]
 **size** | **int32** | Size of result set to return. It requires **operation** is set to &#x60;query&#x60;. Defaults to 10 (i.e. 10 event records). | [default to 10]

### Return type

[**V1alpha1APIResponse**](V1alpha1APIResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

