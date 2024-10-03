# \DevOpsJenkinsfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ToJenkinsfile**](DevOpsJenkinsfileAPI.md#ToJenkinsfile) | **Post** /kapis/devops.kubesphere.io/v1alpha2/tojenkinsfile | Convert json to jenkinsfile format.
[**ToJson**](DevOpsJenkinsfileAPI.md#ToJson) | **Post** /kapis/devops.kubesphere.io/v1alpha2/tojson | Convert jenkinsfile to json format. Usually the frontend uses json to show or edit pipeline



## ToJenkinsfile

> DevopsResJenkinsfile ToJenkinsfile(ctx).Json(json).Execute()

Convert json to jenkinsfile format.

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
	json := "json_example" // string | json data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsJenkinsfileAPI.ToJenkinsfile(context.Background()).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsJenkinsfileAPI.ToJenkinsfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToJenkinsfile`: DevopsResJenkinsfile
	fmt.Fprintf(os.Stdout, "Response from `DevOpsJenkinsfileAPI.ToJenkinsfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToJenkinsfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **json** | **string** | json data | 

### Return type

[**DevopsResJenkinsfile**](DevopsResJenkinsfile.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToJson

> DevopsResJson ToJson(ctx).Jenkinsfile(jenkinsfile).Execute()

Convert jenkinsfile to json format. Usually the frontend uses json to show or edit pipeline

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
	jenkinsfile := "jenkinsfile_example" // string | jenkinsfile (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsJenkinsfileAPI.ToJson(context.Background()).Jenkinsfile(jenkinsfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsJenkinsfileAPI.ToJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToJson`: DevopsResJson
	fmt.Fprintf(os.Stdout, "Response from `DevOpsJenkinsfileAPI.ToJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jenkinsfile** | **string** | jenkinsfile | 

### Return type

[**DevopsResJson**](DevopsResJson.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

