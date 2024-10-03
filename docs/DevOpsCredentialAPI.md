# \DevOpsCredentialAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectCredentialUsage**](DevOpsCredentialAPI.md#GetProjectCredentialUsage) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/credentials/{credential}/usage | Get the specified credential usage of the DevOps project



## GetProjectCredentialUsage

> DevopsCredential GetProjectCredentialUsage(ctx, devops, credential).Execute()

Get the specified credential usage of the DevOps project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	credential := "credential_example" // string | credential's ID, e.g. dockerhub-id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsCredentialAPI.GetProjectCredentialUsage(context.Background(), devops, credential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsCredentialAPI.GetProjectCredentialUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectCredentialUsage`: DevopsCredential
	fmt.Fprintf(os.Stdout, "Response from `DevOpsCredentialAPI.GetProjectCredentialUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**credential** | **string** | credential&#39;s ID, e.g. dockerhub-id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectCredentialUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevopsCredential**](DevopsCredential.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

