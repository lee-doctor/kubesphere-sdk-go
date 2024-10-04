# \ClusterResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleClusterGetNamespacedAbnormalWorkloads**](ClusterResourcesAPI.md#HandleClusterGetNamespacedAbnormalWorkloads) | **Get** /kapis/resources.kubesphere.io/v1alpha2/abnormalworkloads | get abnormal workloads&#39; count of whole cluster
[**HandleGetClusterQuotas**](ClusterResourcesAPI.md#HandleGetClusterQuotas) | **Get** /kapis/resources.kubesphere.io/v1alpha2/quotas | get whole cluster&#39;s resource usage
[**HandleListResources**](ClusterResourcesAPI.md#HandleListResources) | **Get** /kapis/resources.kubesphere.io/v1alpha2/{resources} | Cluster level resources



## HandleClusterGetNamespacedAbnormalWorkloads

> ApiWorkloads HandleClusterGetNamespacedAbnormalWorkloads(ctx).Execute()

get abnormal workloads' count of whole cluster

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterResourcesAPI.HandleClusterGetNamespacedAbnormalWorkloads(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.HandleClusterGetNamespacedAbnormalWorkloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleClusterGetNamespacedAbnormalWorkloads`: ApiWorkloads
	fmt.Fprintf(os.Stdout, "Response from `ClusterResourcesAPI.HandleClusterGetNamespacedAbnormalWorkloads`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleClusterGetNamespacedAbnormalWorkloadsRequest struct via the builder pattern


### Return type

[**ApiWorkloads**](ApiWorkloads.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetClusterQuotas

> ApiResourceQuota HandleGetClusterQuotas(ctx).Execute()

get whole cluster's resource usage

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterResourcesAPI.HandleGetClusterQuotas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.HandleGetClusterQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetClusterQuotas`: ApiResourceQuota
	fmt.Fprintf(os.Stdout, "Response from `ClusterResourcesAPI.HandleGetClusterQuotas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetClusterQuotasRequest struct via the builder pattern


### Return type

[**ApiResourceQuota**](ApiResourceQuota.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleListResources

> ModelsPageableResponse HandleListResources(ctx, resources).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

Cluster level resources

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
	resources := "resources_example" // string | cluster level resource type, e.g. nodes,workspaces,storageclasses,clusterrole.
	conditions := "conditions_example" // string | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterResourcesAPI.HandleListResources(context.Background(), resources).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterResourcesAPI.HandleListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleListResources`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `ClusterResourcesAPI.HandleListResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resources** | **string** | cluster level resource type, e.g. nodes,workspaces,storageclasses,clusterrole. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleListResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conditions** | **string** | query conditions, connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]
 **reverse** | **string** | sort parameters, e.g. reverse&#x3D;true | 
 **orderBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

