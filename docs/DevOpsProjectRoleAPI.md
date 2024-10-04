# \DevOpsProjectRoleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevOpsProjectNamespaceRole**](DevOpsProjectRoleAPI.md#CreateDevOpsProjectNamespaceRole) | **Post** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/roles | Create role in the specified devops project.
[**DeleteDevOpsProjectRole**](DevOpsProjectRoleAPI.md#DeleteDevOpsProjectRole) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/roles/{role} | Delete role in the specified devops project.
[**DescribeDevOpsProjectNamespaceRole**](DevOpsProjectRoleAPI.md#DescribeDevOpsProjectNamespaceRole) | **Get** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/roles/{role} | Retrieve devops project role details.
[**ListDevOpsProjectRoles**](DevOpsProjectRoleAPI.md#ListDevOpsProjectRoles) | **Get** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/roles | List all roles in the specified devops project.
[**PatchDevOpsProjectNamespaceRole**](DevOpsProjectRoleAPI.md#PatchDevOpsProjectNamespaceRole) | **Patch** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/roles/{role} | Patch devops project role.
[**RetrieveDevOpsProjectMemberRoleTemplates**](DevOpsProjectRoleAPI.md#RetrieveDevOpsProjectMemberRoleTemplates) | **Get** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/members/{member}/roles | Retrieve member&#39;s role templates in devops project.
[**UpdateDevOpsNamespaceRole**](DevOpsProjectRoleAPI.md#UpdateDevOpsNamespaceRole) | **Put** /kapis/iam.kubesphere.io/v1alpha2/devops/{devops}/roles/{role} | Update devops project role.



## CreateDevOpsProjectNamespaceRole

> V1Role CreateDevOpsProjectNamespaceRole(ctx, devops).Body(body).Execute()

Create role in the specified devops project.

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
	devops := "devops_example" // string | devops project name
	body := *openapiclient.NewV1Role([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectRoleAPI.CreateDevOpsProjectNamespaceRole(context.Background(), devops).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectRoleAPI.CreateDevOpsProjectNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDevOpsProjectNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectRoleAPI.CreateDevOpsProjectNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDevOpsProjectNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1Role**](V1Role.md) |  | 

### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevOpsProjectRole

> ErrorsError DeleteDevOpsProjectRole(ctx, devops, role).Execute()

Delete role in the specified devops project.

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
	devops := "devops_example" // string | devops project name
	role := "role_example" // string | role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectRoleAPI.DeleteDevOpsProjectRole(context.Background(), devops, role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectRoleAPI.DeleteDevOpsProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDevOpsProjectRole`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectRoleAPI.DeleteDevOpsProjectRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevOpsProjectRoleRequest struct via the builder pattern


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


## DescribeDevOpsProjectNamespaceRole

> V1Role DescribeDevOpsProjectNamespaceRole(ctx, devops, role).Execute()

Retrieve devops project role details.

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
	devops := "devops_example" // string | devops project name
	role := "role_example" // string | role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectRoleAPI.DescribeDevOpsProjectNamespaceRole(context.Background(), devops, role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectRoleAPI.DescribeDevOpsProjectNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDevOpsProjectNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectRoleAPI.DescribeDevOpsProjectNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDevOpsProjectNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevOpsProjectRoles

> ApiListResult ListDevOpsProjectRoles(ctx, devops).Execute()

List all roles in the specified devops project.

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
	devops := "devops_example" // string | devops project name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectRoleAPI.ListDevOpsProjectRoles(context.Background(), devops).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectRoleAPI.ListDevOpsProjectRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevOpsProjectRoles`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectRoleAPI.ListDevOpsProjectRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDevOpsProjectRolesRequest struct via the builder pattern


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


## PatchDevOpsProjectNamespaceRole

> V1Role PatchDevOpsProjectNamespaceRole(ctx, devops, role).Body(body).Execute()

Patch devops project role.

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
	devops := "devops_example" // string | devops project name
	role := "role_example" // string | role name
	body := *openapiclient.NewV1Role([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectRoleAPI.PatchDevOpsProjectNamespaceRole(context.Background(), devops, role).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectRoleAPI.PatchDevOpsProjectNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDevOpsProjectNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectRoleAPI.PatchDevOpsProjectNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDevOpsProjectNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Role**](V1Role.md) |  | 

### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDevOpsProjectMemberRoleTemplates

> ApiListResult RetrieveDevOpsProjectMemberRoleTemplates(ctx, devops, member).Execute()

Retrieve member's role templates in devops project.

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
	devops := "devops_example" // string | devops project name
	member := "member_example" // string | devops project member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectRoleAPI.RetrieveDevOpsProjectMemberRoleTemplates(context.Background(), devops, member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectRoleAPI.RetrieveDevOpsProjectMemberRoleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDevOpsProjectMemberRoleTemplates`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectRoleAPI.RetrieveDevOpsProjectMemberRoleTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**member** | **string** | devops project member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDevOpsProjectMemberRoleTemplatesRequest struct via the builder pattern


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


## UpdateDevOpsNamespaceRole

> V1Role UpdateDevOpsNamespaceRole(ctx, devops, role).Body(body).Execute()

Update devops project role.

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
	devops := "devops_example" // string | devops project name
	role := "role_example" // string | role name
	body := *openapiclient.NewV1Role([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectRoleAPI.UpdateDevOpsNamespaceRole(context.Background(), devops, role).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectRoleAPI.UpdateDevOpsNamespaceRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDevOpsNamespaceRole`: V1Role
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectRoleAPI.UpdateDevOpsNamespaceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops project name | 
**role** | **string** | role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDevOpsNamespaceRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Role**](V1Role.md) |  | 

### Return type

[**V1Role**](V1Role.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

