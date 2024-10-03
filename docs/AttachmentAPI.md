# \AttachmentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeAttachment**](AttachmentAPI.md#DescribeAttachment) | **Get** /kapis/openpitrix.io/v1/attachments/{attachment} | Get attachment by attachment id



## DescribeAttachment

> OpenpitrixAttachment DescribeAttachment(ctx, attachment).Execute()

Get attachment by attachment id

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
	attachment := "attachment_example" // string | attachment id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentAPI.DescribeAttachment(context.Background(), attachment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAPI.DescribeAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAttachment`: OpenpitrixAttachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentAPI.DescribeAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachment** | **string** | attachment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenpitrixAttachment**](OpenpitrixAttachment.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

