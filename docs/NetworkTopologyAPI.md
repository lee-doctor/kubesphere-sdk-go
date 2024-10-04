# \NetworkTopologyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNamespaceNodeTopology**](NetworkTopologyAPI.md#GetNamespaceNodeTopology) | **Get** /kapis/network.kubesphere.io/v1alpha2/namespaces/{namespace}/topology/{node_id} | Get the topology with specifying a node id in the whole topology and specifying a namespace
[**GetNamespaceTopology**](NetworkTopologyAPI.md#GetNamespaceTopology) | **Get** /kapis/network.kubesphere.io/v1alpha2/namespaces/{namespace}/topology | Get the topology with specifying a namespace



## GetNamespaceNodeTopology

> V1alpha2NodeResponse GetNamespaceNodeTopology(ctx, namespace, nodeId).Execute()

Get the topology with specifying a node id in the whole topology and specifying a namespace

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
	namespace := "namespace_example" // string | name of the namespace
	nodeId := "nodeId_example" // string | id of the node in the whole topology

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTopologyAPI.GetNamespaceNodeTopology(context.Background(), namespace, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTopologyAPI.GetNamespaceNodeTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceNodeTopology`: V1alpha2NodeResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkTopologyAPI.GetNamespaceNodeTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of the namespace | 
**nodeId** | **string** | id of the node in the whole topology | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceNodeTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1alpha2NodeResponse**](V1alpha2NodeResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceTopology

> V1alpha2TopologyResponse GetNamespaceTopology(ctx, namespace).Execute()

Get the topology with specifying a namespace

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
	namespace := "namespace_example" // string | name of the namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkTopologyAPI.GetNamespaceTopology(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkTopologyAPI.GetNamespaceTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaceTopology`: V1alpha2TopologyResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkTopologyAPI.GetNamespaceTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | name of the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2TopologyResponse**](V1alpha2TopologyResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

