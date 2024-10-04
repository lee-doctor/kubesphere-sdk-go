# \MultiClusterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindingClusters**](MultiClusterAPI.md#BindingClusters) | **Post** /kapis/cluster.kubesphere.io/v1alpha1/labelbindings | Binding clusters.
[**CreateLabels**](MultiClusterAPI.md#CreateLabels) | **Post** /kapis/cluster.kubesphere.io/v1alpha1/labels | Create cluster labels.
[**DeleteLabels**](MultiClusterAPI.md#DeleteLabels) | **Delete** /kapis/cluster.kubesphere.io/v1alpha1/labels | Delete cluster labels.
[**ListLabelGroups**](MultiClusterAPI.md#ListLabelGroups) | **Get** /kapis/cluster.kubesphere.io/v1alpha1/labels | List labels.
[**UpdateKubeConfig**](MultiClusterAPI.md#UpdateKubeConfig) | **Put** /kapis/cluster.kubesphere.io/v1alpha1/clusters/{cluster}/kubeconfig | Update kubeconfig
[**UpdateLabel**](MultiClusterAPI.md#UpdateLabel) | **Put** /kapis/cluster.kubesphere.io/v1alpha1/labels/{label} | Update a label.
[**ValidateCluster**](MultiClusterAPI.md#ValidateCluster) | **Post** /kapis/cluster.kubesphere.io/v1alpha1/clusters/validation | Cluster validation



## BindingClusters

> BindingClusters(ctx).Body(body).Execute()

Binding clusters.

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
	body := []openapiclient.V1alpha1BindingClustersRequest{*openapiclient.NewV1alpha1BindingClustersRequest([]string{"Clusters_example"}, []string{"Labels_example"})} // []V1alpha1BindingClustersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.BindingClusters(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.BindingClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindingClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1alpha1BindingClustersRequest**](V1alpha1BindingClustersRequest.md) |  | 

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


## CreateLabels

> ApiListResult CreateLabels(ctx).Body(body).Execute()

Create cluster labels.

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
	body := []openapiclient.V1alpha1CreateLabelRequest{*openapiclient.NewV1alpha1CreateLabelRequest("Key_example", "Value_example")} // []V1alpha1CreateLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiClusterAPI.CreateLabels(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.CreateLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLabels`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `MultiClusterAPI.CreateLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1alpha1CreateLabelRequest**](V1alpha1CreateLabelRequest.md) |  | 

### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLabels

> DeleteLabels(ctx).Body(body).Execute()

Delete cluster labels.

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
	body := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.DeleteLabels(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.DeleteLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** |  | 

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


## ListLabelGroups

> map[string][]V1alpha1LabelValue ListLabelGroups(ctx).Execute()

List labels.

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
	resp, r, err := apiClient.MultiClusterAPI.ListLabelGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ListLabelGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLabelGroups`: map[string][]V1alpha1LabelValue
	fmt.Fprintf(os.Stdout, "Response from `MultiClusterAPI.ListLabelGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelGroupsRequest struct via the builder pattern


### Return type

[**map[string][]V1alpha1LabelValue**](array.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKubeConfig

> UpdateKubeConfig(ctx, cluster).Execute()

Update kubeconfig

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
	cluster := "cluster_example" // string | The specified cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.UpdateKubeConfig(context.Background(), cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.UpdateKubeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** | The specified cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKubeConfigRequest struct via the builder pattern


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


## UpdateLabel

> V1alpha1Label UpdateLabel(ctx, label).Body(body).Execute()

Update a label.

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
	label := "label_example" // string | Name of the label.
	body := *openapiclient.NewV1alpha1CreateLabelRequest("Key_example", "Value_example") // V1alpha1CreateLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiClusterAPI.UpdateLabel(context.Background(), label).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.UpdateLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLabel`: V1alpha1Label
	fmt.Fprintf(os.Stdout, "Response from `MultiClusterAPI.UpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string** | Name of the label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha1CreateLabelRequest**](V1alpha1CreateLabelRequest.md) |  | 

### Return type

[**V1alpha1Label**](V1alpha1Label.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCluster

> ValidateCluster(ctx).Body(body).Execute()

Cluster validation

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
	body := *openapiclient.NewV1alpha1Cluster() // V1alpha1Cluster | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.ValidateCluster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ValidateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1Cluster**](V1alpha1Cluster.md) |  | 

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

