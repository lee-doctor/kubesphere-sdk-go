# \KubeSphereMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleKubeSphereMetricsQuery**](KubeSphereMetricsAPI.md#HandleKubeSphereMetricsQuery) | **Get** /kapis/monitoring.kubesphere.io/v1alpha3/kubesphere | Get platform-level metric data.



## HandleKubeSphereMetricsQuery

> MonitoringMetrics HandleKubeSphereMetricsQuery(ctx).Execute()

Get platform-level metric data.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KubeSphereMetricsAPI.HandleKubeSphereMetricsQuery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubeSphereMetricsAPI.HandleKubeSphereMetricsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleKubeSphereMetricsQuery`: MonitoringMetrics
	fmt.Fprintf(os.Stdout, "Response from `KubeSphereMetricsAPI.HandleKubeSphereMetricsQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHandleKubeSphereMetricsQueryRequest struct via the builder pattern


### Return type

[**MonitoringMetrics**](MonitoringMetrics.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

