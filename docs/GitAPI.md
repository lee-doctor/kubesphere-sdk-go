# \GitAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleVerifyGitCredential**](GitAPI.md#HandleVerifyGitCredential) | **Post** /kapis/resources.kubesphere.io/v1alpha2/git/verify | Verify if the kubernetes secret has read access to the git repository



## HandleVerifyGitCredential

> ErrorsError HandleVerifyGitCredential(ctx).Body(body).Execute()

Verify if the kubernetes secret has read access to the git repository

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
	resp, r, err := apiClient.GitAPI.HandleVerifyGitCredential(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.HandleVerifyGitCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleVerifyGitCredential`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.HandleVerifyGitCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleVerifyGitCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitAuthInfo**](GitAuthInfo.md) |  | 

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

