# \NamespacedResourceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleNamespaceGetResources**](NamespacedResourceAPI.md#HandleNamespaceGetResources) | **Get** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/{resources}/{name} | Namespace level get resource query
[**HandleNamespaceListResources**](NamespacedResourceAPI.md#HandleNamespaceListResources) | **Get** /kapis/resources.kubesphere.io/v1alpha3/namespaces/{namespace}/{resources} | Namespace level resource query



## HandleNamespaceGetResources

> ApiListResult HandleNamespaceGetResources(ctx, namespace, resources, name).Execute()

Namespace level get resource query

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
	namespace := "namespace_example" // string | the name of the project
	resources := "resources_example" // string | namespace level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | the name of resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourceAPI.HandleNamespaceGetResources(context.Background(), namespace, resources, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourceAPI.HandleNamespaceGetResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleNamespaceGetResources`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourceAPI.HandleNamespaceGetResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 
**resources** | **string** | namespace level resource type, e.g. pods,jobs,configmaps,services. | 
**name** | **string** | the name of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleNamespaceGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## HandleNamespaceListResources

> ApiListResult HandleNamespaceListResources(ctx, namespace, resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()

Namespace level resource query

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
	namespace := "namespace_example" // string | the name of the project
	resources := "resources_example" // string | namespace level resource type, e.g. pods,jobs,configmaps,services.
	name := "name_example" // string | name used to do filtering (optional)
	page := "page_example" // string | page (optional) (default to "page=1")
	limit := "limit_example" // string | limit (optional)
	ascending := "ascending_example" // string | sort parameters, e.g. reverse=true (optional) (default to "ascending=false")
	sortBy := "sortBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespacedResourceAPI.HandleNamespaceListResources(context.Background(), namespace, resources).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacedResourceAPI.HandleNamespaceListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleNamespaceListResources`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespacedResourceAPI.HandleNamespaceListResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | the name of the project | 
**resources** | **string** | namespace level resource type, e.g. pods,jobs,configmaps,services. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleNamespaceListResourcesRequest struct via the builder pattern


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

