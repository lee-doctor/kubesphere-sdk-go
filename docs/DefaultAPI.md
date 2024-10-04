# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleJobReRun**](DefaultAPI.md#HandleJobReRun) | **Post** /kapis/operations.kubesphere.io/v1alpha2/namespaces/{namespace}/jobs/{job} | Rerun job whether the job is complete or not



## HandleJobReRun

> ErrorsError HandleJobReRun(ctx, job, namespace).ResourceVersion(resourceVersion).Action(action).Execute()

Rerun job whether the job is complete or not

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
	namespace := "namespace_example" // string | the name of the namespace where the job runs in
	resourceVersion := "resourceVersion_example" // string | version of job, rerun when the version matches
	action := "action_example" // string | action must be \"rerun\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HandleJobReRun(context.Background(), job, namespace).ResourceVersion(resourceVersion).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HandleJobReRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleJobReRun`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HandleJobReRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**job** | **string** | job name | 
**namespace** | **string** | the name of the namespace where the job runs in | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleJobReRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceVersion** | **string** | version of job, rerun when the version matches | 
 **action** | **string** | action must be \&quot;rerun\&quot; | 

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

