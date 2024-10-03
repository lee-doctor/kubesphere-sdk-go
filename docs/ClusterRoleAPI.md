# \ClusterRoleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterRole**](ClusterRoleAPI.md#CreateClusterRole) | **Post** /kapis/iam.kubesphere.io/v1alpha2/clusterroles | Create cluster role.
[**DeleteClusterRole**](ClusterRoleAPI.md#DeleteClusterRole) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/clusterroles/{clusterrole} | Delete cluster role.
[**DescribeClusterRole**](ClusterRoleAPI.md#DescribeClusterRole) | **Get** /kapis/iam.kubesphere.io/v1alpha2/clusterroles/{clusterrole} | Retrieve cluster role details.
[**ListClusterRoles**](ClusterRoleAPI.md#ListClusterRoles) | **Get** /kapis/iam.kubesphere.io/v1alpha2/clusterroles | List all cluster roles.
[**PatchClusterRole**](ClusterRoleAPI.md#PatchClusterRole) | **Patch** /kapis/iam.kubesphere.io/v1alpha2/clusterroles/{clusterrole} | Patch cluster role.
[**RetrieveClusterMemberRoleTemplates**](ClusterRoleAPI.md#RetrieveClusterMemberRoleTemplates) | **Get** /kapis/iam.kubesphere.io/v1alpha2/clustermembers/{clustermember}/clusterroles | Retrieve user&#39;s role templates in cluster.
[**UpdateClusterRole**](ClusterRoleAPI.md#UpdateClusterRole) | **Put** /kapis/iam.kubesphere.io/v1alpha2/clusterroles/{clusterrole} | Update cluster role.



## CreateClusterRole

> V1ClusterRole CreateClusterRole(ctx).Body(body).Execute()

Create cluster role.

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
	body := *openapiclient.NewV1ClusterRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1ClusterRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterRoleAPI.CreateClusterRole(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterRoleAPI.CreateClusterRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClusterRole`: V1ClusterRole
	fmt.Fprintf(os.Stdout, "Response from `ClusterRoleAPI.CreateClusterRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ClusterRole**](V1ClusterRole.md) |  | 

### Return type

[**V1ClusterRole**](V1ClusterRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterRole

> ErrorsError DeleteClusterRole(ctx, clusterrole).Execute()

Delete cluster role.

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
	clusterrole := "clusterrole_example" // string | cluster role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterRoleAPI.DeleteClusterRole(context.Background(), clusterrole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterRoleAPI.DeleteClusterRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClusterRole`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `ClusterRoleAPI.DeleteClusterRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterrole** | **string** | cluster role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DescribeClusterRole

> V1ClusterRole DescribeClusterRole(ctx, clusterrole).Execute()

Retrieve cluster role details.

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
	clusterrole := "clusterrole_example" // string | cluster role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterRoleAPI.DescribeClusterRole(context.Background(), clusterrole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterRoleAPI.DescribeClusterRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeClusterRole`: V1ClusterRole
	fmt.Fprintf(os.Stdout, "Response from `ClusterRoleAPI.DescribeClusterRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterrole** | **string** | cluster role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeClusterRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ClusterRole**](V1ClusterRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterRoles

> ApiListResult ListClusterRoles(ctx).Execute()

List all cluster roles.

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
	resp, r, err := apiClient.ClusterRoleAPI.ListClusterRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterRoleAPI.ListClusterRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterRoles`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `ClusterRoleAPI.ListClusterRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterRolesRequest struct via the builder pattern


### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchClusterRole

> V1ClusterRole PatchClusterRole(ctx, clusterrole).Body(body).Execute()

Patch cluster role.

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
	clusterrole := "clusterrole_example" // string | cluster role name
	body := *openapiclient.NewV1ClusterRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1ClusterRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterRoleAPI.PatchClusterRole(context.Background(), clusterrole).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterRoleAPI.PatchClusterRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchClusterRole`: V1ClusterRole
	fmt.Fprintf(os.Stdout, "Response from `ClusterRoleAPI.PatchClusterRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterrole** | **string** | cluster role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchClusterRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1ClusterRole**](V1ClusterRole.md) |  | 

### Return type

[**V1ClusterRole**](V1ClusterRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveClusterMemberRoleTemplates

> ApiListResult RetrieveClusterMemberRoleTemplates(ctx, clustermember).Execute()

Retrieve user's role templates in cluster.

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
	clustermember := "clustermember_example" // string | cluster member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterRoleAPI.RetrieveClusterMemberRoleTemplates(context.Background(), clustermember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterRoleAPI.RetrieveClusterMemberRoleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveClusterMemberRoleTemplates`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `ClusterRoleAPI.RetrieveClusterMemberRoleTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clustermember** | **string** | cluster member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveClusterMemberRoleTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterRole

> V1ClusterRole UpdateClusterRole(ctx, clusterrole).Body(body).Execute()

Update cluster role.

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
	clusterrole := "clusterrole_example" // string | cluster role name
	body := *openapiclient.NewV1ClusterRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1ClusterRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterRoleAPI.UpdateClusterRole(context.Background(), clusterrole).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterRoleAPI.UpdateClusterRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterRole`: V1ClusterRole
	fmt.Fprintf(os.Stdout, "Response from `ClusterRoleAPI.UpdateClusterRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterrole** | **string** | cluster role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1ClusterRole**](V1ClusterRole.md) |  | 

### Return type

[**V1ClusterRole**](V1ClusterRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

