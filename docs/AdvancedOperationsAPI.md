# \AdvancedOperationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegistryEntry**](AdvancedOperationsAPI.md#GetRegistryEntry) | **Get** /kapis/resources.kubesphere.io/v1alpha2/registry/blob | Retrieve the blob from the registry
[**JobReRun**](AdvancedOperationsAPI.md#JobReRun) | **Post** /kapis/operations.kubesphere.io/v1alpha2/namespaces/{namespace}/jobs/{job} | Job rerun
[**VerifyGitCredential**](AdvancedOperationsAPI.md#VerifyGitCredential) | **Post** /kapis/resources.kubesphere.io/v1alpha2/git/verify | Verify the git credential
[**VerifyRegistryCredential**](AdvancedOperationsAPI.md#VerifyRegistryCredential) | **Post** /kapis/resources.kubesphere.io/v1alpha2/registry/verify | Verify registry credential



## GetRegistryEntry

> RegistriesImageDetails GetRegistryEntry(ctx).Image(image).Namespace(namespace).Secret(secret).Insecure(insecure).Execute()

Retrieve the blob from the registry

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
	insecure := "insecure_example" // string | whether verify cert if using https repo (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOperationsAPI.GetRegistryEntry(context.Background()).Image(image).Namespace(namespace).Secret(secret).Insecure(insecure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOperationsAPI.GetRegistryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegistryEntry`: RegistriesImageDetails
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOperationsAPI.GetRegistryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | **string** | query image, condition for filtering. | 
 **namespace** | **string** | namespace which secret in. | 
 **secret** | **string** | secret name | 
 **insecure** | **string** | whether verify cert if using https repo | 

### Return type

[**RegistriesImageDetails**](RegistriesImageDetails.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobReRun

> ErrorsError JobReRun(ctx, job, namespace).ResourceVersion(resourceVersion).Action(action).Execute()

Job rerun



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
	job := "job_example" // string | job name
	namespace := "namespace_example" // string | The specified namespace.
	resourceVersion := "resourceVersion_example" // string | version of job, rerun when the version matches
	action := "action_example" // string | action must be \"rerun\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOperationsAPI.JobReRun(context.Background(), job, namespace).ResourceVersion(resourceVersion).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOperationsAPI.JobReRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobReRun`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOperationsAPI.JobReRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**job** | **string** | job name | 
**namespace** | **string** | The specified namespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobReRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceVersion** | **string** | version of job, rerun when the version matches | 
 **action** | **string** | action must be \&quot;rerun\&quot; | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyGitCredential

> ErrorsError VerifyGitCredential(ctx).Body(body).Execute()

Verify the git credential

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
	body := *openapiclient.NewGitAuthInfo("RemoteUrl_example") // GitAuthInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOperationsAPI.VerifyGitCredential(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOperationsAPI.VerifyGitCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyGitCredential`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOperationsAPI.VerifyGitCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyGitCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitAuthInfo**](GitAuthInfo.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyRegistryCredential

> ErrorsError VerifyRegistryCredential(ctx).Body(body).Execute()

Verify registry credential

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
	resp, r, err := apiClient.AdvancedOperationsAPI.VerifyRegistryCredential(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOperationsAPI.VerifyRegistryCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyRegistryCredential`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOperationsAPI.VerifyRegistryCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyRegistryCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiRegistryCredential**](ApiRegistryCredential.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

