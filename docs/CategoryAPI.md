# \CategoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategory**](CategoryAPI.md#CreateCategory) | **Post** /kapis/openpitrix.io/v1/categories | Create app template category
[**DeleteCategory**](CategoryAPI.md#DeleteCategory) | **Delete** /kapis/openpitrix.io/v1/categories/{category} | Delete the specified category
[**DescribeCategory**](CategoryAPI.md#DescribeCategory) | **Get** /kapis/openpitrix.io/v1/categories/{category} | Describe the specified category
[**ListCategories**](CategoryAPI.md#ListCategories) | **Get** /kapis/openpitrix.io/v1/categories | List categories
[**ModifyCategory**](CategoryAPI.md#ModifyCategory) | **Patch** /kapis/openpitrix.io/v1/categories/{category} | Patch the specified category



## CreateCategory

> OpenpitrixCreateCategoryResponse CreateCategory(ctx, app).Body(body).Execute()

Create app template category

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
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixCreateCategoryRequest() // OpenpitrixCreateCategoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CreateCategory(context.Background(), app).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CreateCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCategory`: OpenpitrixCreateCategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CreateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixCreateCategoryRequest**](OpenpitrixCreateCategoryRequest.md) |  | 

### Return type

[**OpenpitrixCreateCategoryResponse**](OpenpitrixCreateCategoryResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategory

> ErrorsError DeleteCategory(ctx, category).Execute()

Delete the specified category

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
	category := "category_example" // string | category id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.DeleteCategory(context.Background(), category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.DeleteCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCategory`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.DeleteCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCategory

> OpenpitrixCategory DescribeCategory(ctx, category).Execute()

Describe the specified category

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
	category := "category_example" // string | category id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.DescribeCategory(context.Background(), category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.DescribeCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCategory`: OpenpitrixCategory
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.DescribeCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenpitrixCategory**](OpenpitrixCategory.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategories

> ModelsPageableResponse ListCategories(ctx).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

List categories

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
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.ListCategories(context.Background()).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.ListCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCategories`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.ListCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditions** | **string** | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
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


## ModifyCategory

> ErrorsError ModifyCategory(ctx, category).Body(body).Execute()

Patch the specified category

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
	category := "category_example" // string | category id
	body := *openapiclient.NewOpenpitrixModifyCategoryRequest() // OpenpitrixModifyCategoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.ModifyCategory(context.Background(), category).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.ModifyCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCategory`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.ModifyCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | category id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixModifyCategoryRequest**](OpenpitrixModifyCategoryRequest.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

