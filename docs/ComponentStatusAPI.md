# \ComponentStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentsStatusV1alpha2**](ComponentStatusAPI.md#GetComponentsStatusV1alpha2) | **Get** /kapis/resources.kubesphere.io/v1alpha2/components/{component} | Describe the specified system component
[**GetComponentsStatusV1alpha3**](ComponentStatusAPI.md#GetComponentsStatusV1alpha3) | **Get** /kapis/resources.kubesphere.io/v1alpha3/components/{component} | Describe the specified system component
[**GetComponentsV1alpha2**](ComponentStatusAPI.md#GetComponentsV1alpha2) | **Get** /kapis/resources.kubesphere.io/v1alpha2/components | List the system components
[**GetComponentsV1alpha3**](ComponentStatusAPI.md#GetComponentsV1alpha3) | **Get** /kapis/resources.kubesphere.io/v1alpha3/components | List the system components
[**GetSystemHealthStatusV1alpha2**](ComponentStatusAPI.md#GetSystemHealthStatusV1alpha2) | **Get** /kapis/resources.kubesphere.io/v1alpha2/componenthealth | Get the health status of system components
[**GetSystemHealthStatusV1alpha3**](ComponentStatusAPI.md#GetSystemHealthStatusV1alpha3) | **Get** /kapis/resources.kubesphere.io/v1alpha3/componenthealth | Get the health status of system components



## GetComponentsStatusV1alpha2

> V1alpha2ComponentStatus GetComponentsStatusV1alpha2(ctx, component).Execute()

Describe the specified system component

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
	component := "component_example" // string | component name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentStatusAPI.GetComponentsStatusV1alpha2(context.Background(), component).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.GetComponentsStatusV1alpha2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentsStatusV1alpha2`: V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.GetComponentsStatusV1alpha2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**component** | **string** | component name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsStatusV1alpha2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentsStatusV1alpha3

> V1alpha2ComponentStatus GetComponentsStatusV1alpha3(ctx, component).Execute()

Describe the specified system component

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
	component := "component_example" // string | component name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentStatusAPI.GetComponentsStatusV1alpha3(context.Background(), component).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.GetComponentsStatusV1alpha3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentsStatusV1alpha3`: V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.GetComponentsStatusV1alpha3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**component** | **string** | component name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsStatusV1alpha3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentsV1alpha2

> []V1alpha2ComponentStatus GetComponentsV1alpha2(ctx).Execute()

List the system components

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
	resp, r, err := apiClient.ComponentStatusAPI.GetComponentsV1alpha2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.GetComponentsV1alpha2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentsV1alpha2`: []V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.GetComponentsV1alpha2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsV1alpha2Request struct via the builder pattern


### Return type

[**[]V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentsV1alpha3

> []V1alpha2ComponentStatus GetComponentsV1alpha3(ctx).Execute()

List the system components

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
	resp, r, err := apiClient.ComponentStatusAPI.GetComponentsV1alpha3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.GetComponentsV1alpha3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentsV1alpha3`: []V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.GetComponentsV1alpha3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsV1alpha3Request struct via the builder pattern


### Return type

[**[]V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemHealthStatusV1alpha2

> V1alpha2HealthStatus GetSystemHealthStatusV1alpha2(ctx).Execute()

Get the health status of system components

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
	resp, r, err := apiClient.ComponentStatusAPI.GetSystemHealthStatusV1alpha2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.GetSystemHealthStatusV1alpha2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemHealthStatusV1alpha2`: V1alpha2HealthStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.GetSystemHealthStatusV1alpha2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemHealthStatusV1alpha2Request struct via the builder pattern


### Return type

[**V1alpha2HealthStatus**](V1alpha2HealthStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemHealthStatusV1alpha3

> V1alpha2HealthStatus GetSystemHealthStatusV1alpha3(ctx).Execute()

Get the health status of system components

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
	resp, r, err := apiClient.ComponentStatusAPI.GetSystemHealthStatusV1alpha3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.GetSystemHealthStatusV1alpha3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemHealthStatusV1alpha3`: V1alpha2HealthStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.GetSystemHealthStatusV1alpha3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemHealthStatusV1alpha3Request struct via the builder pattern


### Return type

[**V1alpha2HealthStatus**](V1alpha2HealthStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

