# \TerminalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleTerminalSession**](TerminalAPI.md#HandleTerminalSession) | **Get** /kapis/terminal.kubesphere.io/v1alpha2/namespaces/{namespace}/pods/{pod}/exec | create terminal session



## HandleTerminalSession

> HandleTerminalSession(ctx, namespace, pod).Execute()

create terminal session

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
	namespace := "namespace_example" // string | namespace of which the pod located in
	pod := "pod_example" // string | name of the pod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TerminalAPI.HandleTerminalSession(context.Background(), namespace, pod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerminalAPI.HandleTerminalSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace of which the pod located in | 
**pod** | **string** | name of the pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleTerminalSessionRequest struct via the builder pattern


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

