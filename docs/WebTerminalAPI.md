# \WebTerminalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodeExec**](WebTerminalAPI.md#CreateNodeExec) | **Get** /kapis/terminal.kubesphere.io/v1alpha2/nodes/{nodename}/exec | Create node terminal session
[**CreatePodExec**](WebTerminalAPI.md#CreatePodExec) | **Get** /kapis/terminal.kubesphere.io/v1alpha2/namespaces/{namespace}/pods/{pod}/exec | Create pod terminal session
[**CreateUserKubectlPodExec**](WebTerminalAPI.md#CreateUserKubectlPodExec) | **Get** /kapis/terminal.kubesphere.io/v1alpha2/users/{user}/kubectl | Create kubectl pod terminal session for current user
[**DownloadFileFromPod**](WebTerminalAPI.md#DownloadFileFromPod) | **Get** /kapis/terminal.kubesphere.io/v1alpha2/namespaces/{namespace}/pods/{pod}/file | Download file from pod
[**UploadFileToPod**](WebTerminalAPI.md#UploadFileToPod) | **Post** /kapis/terminal.kubesphere.io/v1alpha2/namespaces/{namespace}/pods/{pod}/file | Upload files to pod



## CreateNodeExec

> CreateNodeExec(ctx, nodename).Execute()

Create node terminal session

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
	nodename := "nodename_example" // string | node name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebTerminalAPI.CreateNodeExec(context.Background(), nodename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebTerminalAPI.CreateNodeExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodename** | **string** | node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePodExec

> CreatePodExec(ctx, namespace, pod).Execute()

Create pod terminal session

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
	namespace := "namespace_example" // string | The specified namespace.
	pod := "pod_example" // string | pod name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebTerminalAPI.CreatePodExec(context.Background(), namespace, pod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebTerminalAPI.CreatePodExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**pod** | **string** | pod name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePodExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserKubectlPodExec

> CreateUserKubectlPodExec(ctx, user).Execute()

Create kubectl pod terminal session for current user

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
	user := "user_example" // string | username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebTerminalAPI.CreateUserKubectlPodExec(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebTerminalAPI.CreateUserKubectlPodExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserKubectlPodExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFileFromPod

> DownloadFileFromPod(ctx, namespace, pod).Container(container).Path(path).Execute()

Download file from pod

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
	namespace := "namespace_example" // string | The specified namespace.
	pod := "pod_example" // string | pod name
	container := "container_example" // string | container name (optional)
	path := "path_example" // string | file path (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebTerminalAPI.DownloadFileFromPod(context.Background(), namespace, pod).Container(container).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebTerminalAPI.DownloadFileFromPod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**pod** | **string** | pod name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileFromPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **container** | **string** | container name | 
 **path** | **string** | file path | 

### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileToPod

> UploadFileToPod(ctx, namespace, pod).Container(container).Path(path).Execute()

Upload files to pod

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
	namespace := "namespace_example" // string | The specified namespace.
	pod := "pod_example" // string | pod name
	container := "container_example" // string | container name (optional)
	path := "path_example" // string | dest dir path (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebTerminalAPI.UploadFileToPod(context.Background(), namespace, pod).Container(container).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebTerminalAPI.UploadFileToPod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The specified namespace. | 
**pod** | **string** | pod name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileToPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **container** | **string** | container name | 
 **path** | **string** | dest dir path | 

### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

