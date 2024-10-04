# \NonResourceAPIsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Version**](NonResourceAPIsAPI.md#Version) | **Get** /version | KubeSphere version info
[**VersionLegacy**](NonResourceAPIsAPI.md#VersionLegacy) | **Get** /kapis/version | KubeSphere version info



## Version

> VersionInfo Version(ctx).Execute()

KubeSphere version info

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
	resp, r, err := apiClient.NonResourceAPIsAPI.Version(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NonResourceAPIsAPI.Version``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Version`: VersionInfo
	fmt.Fprintf(os.Stdout, "Response from `NonResourceAPIsAPI.Version`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionRequest struct via the builder pattern


### Return type

[**VersionInfo**](VersionInfo.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionLegacy

> VersionInfo VersionLegacy(ctx).Execute()

KubeSphere version info



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
	resp, r, err := apiClient.NonResourceAPIsAPI.VersionLegacy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NonResourceAPIsAPI.VersionLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionLegacy`: VersionInfo
	fmt.Fprintf(os.Stdout, "Response from `NonResourceAPIsAPI.VersionLegacy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionLegacyRequest struct via the builder pattern


### Return type

[**VersionInfo**](VersionInfo.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

