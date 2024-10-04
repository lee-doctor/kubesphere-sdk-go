# \ClusteredResourceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleGetResources**](ClusteredResourceAPI.md#HandleGetResources) | **Get** /kapis/resources.kubesphere.io/v1alpha3/{resources}/{name} | Cluster level resource
[**HandleListResourcesV3**](ClusteredResourceAPI.md#HandleListResourcesV3) | **Get** /kapis/resources.kubesphere.io/v1alpha3/{resources} | Cluster level resources



## HandleGetResources

> HandleGetResources(ctx, resources, name).Execute()

Cluster level resource

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
	resources := "resources_example" // string | cluster level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | the name of the clustered resources

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusteredResourceAPI.HandleGetResources(context.Background(), resources, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusteredResourceAPI.HandleGetResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resources** | **string** | cluster level resource type, e.g. pods,jobs,configmaps,services. | 
**name** | **string** | the name of the clustered resources | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## HandleListResourcesV3

> ApiListResult HandleListResourcesV3(ctx, resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()

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
	resources := "resources_example" // string | cluster level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | name used to do filtering (optional)
	page := "page_example" // string | page (optional) (default to "page=1")
	limit := "limit_example" // string | limit (optional)
	ascending := "ascending_example" // string | sort parameters, e.g. reverse=true (optional) (default to "ascending=false")
	sortBy := "sortBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusteredResourceAPI.HandleListResourcesV3(context.Background(), resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusteredResourceAPI.HandleListResourcesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleListResourcesV3`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `ClusteredResourceAPI.HandleListResourcesV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resources** | **string** | cluster level resource type, e.g. pods,jobs,configmaps,services. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleListResourcesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | name used to do filtering | 
 **page** | **string** | page | [default to &quot;page&#x3D;1&quot;]
 **limit** | **string** | limit | 
 **ascending** | **string** | sort parameters, e.g. reverse&#x3D;true | [default to &quot;ascending&#x3D;false&quot;]
 **sortBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

