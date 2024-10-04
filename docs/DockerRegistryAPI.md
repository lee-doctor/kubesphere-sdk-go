# \DockerRegistryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleGetRegistryEntry**](DockerRegistryAPI.md#HandleGetRegistryEntry) | **Get** /kapis/resources.kubesphere.io/v1alpha2/registry/blob | Retrieve the blob from the registry identified
[**HandleVerifyRegistryCredential**](DockerRegistryAPI.md#HandleVerifyRegistryCredential) | **Post** /kapis/resources.kubesphere.io/v1alpha2/registry/verify | verify if a user has access to the docker registry



## HandleGetRegistryEntry

> RegistriesImageDetails HandleGetRegistryEntry(ctx).Image(image).Namespace(namespace).Secret(secret).Execute()

Retrieve the blob from the registry identified

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
	image := "image_example" // string | query image, condition for filtering.
	namespace := "namespace_example" // string | namespace which secret in. (optional)
	secret := "secret_example" // string | secret name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerRegistryAPI.HandleGetRegistryEntry(context.Background()).Image(image).Namespace(namespace).Secret(secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.HandleGetRegistryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleGetRegistryEntry`: RegistriesImageDetails
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.HandleGetRegistryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleGetRegistryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | **string** | query image, condition for filtering. | 
 **namespace** | **string** | namespace which secret in. | 
 **secret** | **string** | secret name | 

### Return type

[**RegistriesImageDetails**](RegistriesImageDetails.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleVerifyRegistryCredential

> ErrorsError HandleVerifyRegistryCredential(ctx).Body(body).Execute()

verify if a user has access to the docker registry

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
	body := *openapiclient.NewApiRegistryCredential("Password_example", "Serverhost_example", "Username_example") // ApiRegistryCredential | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerRegistryAPI.HandleVerifyRegistryCredential(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.HandleVerifyRegistryCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleVerifyRegistryCredential`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.HandleVerifyRegistryCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleVerifyRegistryCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiRegistryCredential**](ApiRegistryCredential.md) |  | 

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

