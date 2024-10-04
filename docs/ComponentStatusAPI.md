# \ComponentStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleGetComponentStatus**](ComponentStatusAPI.md#HandleGetComponentStatus) | **Get** /kapis/resources.kubesphere.io/v1alpha2/components/{component} | Describe the specified system component.
[**HandleGetComponentStatusV3**](ComponentStatusAPI.md#HandleGetComponentStatusV3) | **Get** /kapis/resources.kubesphere.io/v1alpha3/components/{component} | Describe the specified system component.
[**HandleGetComponents**](ComponentStatusAPI.md#HandleGetComponents) | **Get** /kapis/resources.kubesphere.io/v1alpha2/components | List the system components.
[**HandleGetComponentsV3**](ComponentStatusAPI.md#HandleGetComponentsV3) | **Get** /kapis/resources.kubesphere.io/v1alpha3/components | List the system components.
[**HandleGetSystemHealthStatus**](ComponentStatusAPI.md#HandleGetSystemHealthStatus) | **Get** /kapis/resources.kubesphere.io/v1alpha2/componenthealth | Get the health status of system components.
[**HandleGetSystemHealthStatusV3**](ComponentStatusAPI.md#HandleGetSystemHealthStatusV3) | **Get** /kapis/resources.kubesphere.io/v1alpha3/componenthealth | Get the health status of system components.



## HandleGetComponentStatus

> V1alpha2ComponentStatus HandleGetComponentStatus(ctx, component).Execute()

Describe the specified system component.

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
	resp, r, err := apiClient.ComponentStatusAPI.HandleGetComponentStatus(context.Background(), component).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.HandleGetComponentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetComponentStatus`: V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.HandleGetComponentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**component** | **string** | component name | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetComponentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetComponentStatusV3

> V1alpha2ComponentStatus HandleGetComponentStatusV3(ctx, component).Execute()

Describe the specified system component.

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
	resp, r, err := apiClient.ComponentStatusAPI.HandleGetComponentStatusV3(context.Background(), component).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.HandleGetComponentStatusV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetComponentStatusV3`: V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.HandleGetComponentStatusV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**component** | **string** | component name | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetComponentStatusV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetComponents

> []V1alpha2ComponentStatus HandleGetComponents(ctx).Execute()

List the system components.

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
	resp, r, err := apiClient.ComponentStatusAPI.HandleGetComponents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.HandleGetComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetComponents`: []V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.HandleGetComponents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetComponentsRequest struct via the builder pattern


### Return type

[**[]V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetComponentsV3

> []V1alpha2ComponentStatus HandleGetComponentsV3(ctx).Execute()

List the system components.

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
	resp, r, err := apiClient.ComponentStatusAPI.HandleGetComponentsV3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.HandleGetComponentsV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetComponentsV3`: []V1alpha2ComponentStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.HandleGetComponentsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetComponentsV3Request struct via the builder pattern


### Return type

[**[]V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetSystemHealthStatus

> V1alpha2HealthStatus HandleGetSystemHealthStatus(ctx).Execute()

Get the health status of system components.

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
	resp, r, err := apiClient.ComponentStatusAPI.HandleGetSystemHealthStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.HandleGetSystemHealthStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetSystemHealthStatus`: V1alpha2HealthStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.HandleGetSystemHealthStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetSystemHealthStatusRequest struct via the builder pattern


### Return type

[**V1alpha2HealthStatus**](V1alpha2HealthStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleGetSystemHealthStatusV3

> V1alpha2HealthStatus HandleGetSystemHealthStatusV3(ctx).Execute()

Get the health status of system components.

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
	resp, r, err := apiClient.ComponentStatusAPI.HandleGetSystemHealthStatusV3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentStatusAPI.HandleGetSystemHealthStatusV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetSystemHealthStatusV3`: V1alpha2HealthStatus
	fmt.Fprintf(os.Stdout, "Response from `ComponentStatusAPI.HandleGetSystemHealthStatusV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetSystemHealthStatusV3Request struct via the builder pattern


### Return type

[**V1alpha2HealthStatus**](V1alpha2HealthStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

