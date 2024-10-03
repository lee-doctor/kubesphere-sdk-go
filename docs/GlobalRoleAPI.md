# \GlobalRoleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalRole**](GlobalRoleAPI.md#CreateGlobalRole) | **Post** /kapis/iam.kubesphere.io/v1alpha2/globalroles | Create global role.
[**DeleteGlobalRole**](GlobalRoleAPI.md#DeleteGlobalRole) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/globalroles/{globalrole} | Delete global role.
[**DescribeGlobalRole**](GlobalRoleAPI.md#DescribeGlobalRole) | **Get** /kapis/iam.kubesphere.io/v1alpha2/globalroles/{globalrole} | Retrieve global role details.
[**ListGlobalRoles**](GlobalRoleAPI.md#ListGlobalRoles) | **Get** /kapis/iam.kubesphere.io/v1alpha2/globalroles | List all global roles.
[**PatchGlobalRole**](GlobalRoleAPI.md#PatchGlobalRole) | **Patch** /kapis/iam.kubesphere.io/v1alpha2/globalroles/{globalrole} | Patch global role.
[**RetrieveMemberRoleTemplates**](GlobalRoleAPI.md#RetrieveMemberRoleTemplates) | **Get** /kapis/iam.kubesphere.io/v1alpha2/users/{user}/globalroles | Retrieve user&#39;s global role templates.
[**UpdateGlobalRole**](GlobalRoleAPI.md#UpdateGlobalRole) | **Put** /kapis/iam.kubesphere.io/v1alpha2/globalroles/{globalrole} | Update global role.



## CreateGlobalRole

> V1alpha2GlobalRole CreateGlobalRole(ctx).Body(body).Execute()

Create global role.

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
	body := *openapiclient.NewV1alpha2GlobalRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1alpha2GlobalRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalRoleAPI.CreateGlobalRole(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleAPI.CreateGlobalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGlobalRole`: V1alpha2GlobalRole
	fmt.Fprintf(os.Stdout, "Response from `GlobalRoleAPI.CreateGlobalRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha2GlobalRole**](V1alpha2GlobalRole.md) |  | 

### Return type

[**V1alpha2GlobalRole**](V1alpha2GlobalRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalRole

> ErrorsError DeleteGlobalRole(ctx, globalrole).Execute()

Delete global role.

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
	globalrole := "globalrole_example" // string | global role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalRoleAPI.DeleteGlobalRole(context.Background(), globalrole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleAPI.DeleteGlobalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlobalRole`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `GlobalRoleAPI.DeleteGlobalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalrole** | **string** | global role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalRoleRequest struct via the builder pattern


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


## DescribeGlobalRole

> V1alpha2GlobalRole DescribeGlobalRole(ctx, globalrole).Execute()

Retrieve global role details.

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
	globalrole := "globalrole_example" // string | global role name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalRoleAPI.DescribeGlobalRole(context.Background(), globalrole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleAPI.DescribeGlobalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeGlobalRole`: V1alpha2GlobalRole
	fmt.Fprintf(os.Stdout, "Response from `GlobalRoleAPI.DescribeGlobalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalrole** | **string** | global role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeGlobalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha2GlobalRole**](V1alpha2GlobalRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalRoles

> ApiListResult ListGlobalRoles(ctx).Execute()

List all global roles.

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
	resp, r, err := apiClient.GlobalRoleAPI.ListGlobalRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleAPI.ListGlobalRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalRoles`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `GlobalRoleAPI.ListGlobalRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalRolesRequest struct via the builder pattern


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


## PatchGlobalRole

> V1alpha2GlobalRole PatchGlobalRole(ctx, globalrole).Body(body).Execute()

Patch global role.

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
	globalrole := "globalrole_example" // string | global role name
	body := *openapiclient.NewV1alpha2GlobalRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1alpha2GlobalRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalRoleAPI.PatchGlobalRole(context.Background(), globalrole).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleAPI.PatchGlobalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchGlobalRole`: V1alpha2GlobalRole
	fmt.Fprintf(os.Stdout, "Response from `GlobalRoleAPI.PatchGlobalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalrole** | **string** | global role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchGlobalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha2GlobalRole**](V1alpha2GlobalRole.md) |  | 

### Return type

[**V1alpha2GlobalRole**](V1alpha2GlobalRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMemberRoleTemplates

> ApiListResult RetrieveMemberRoleTemplates(ctx, user).Execute()

Retrieve user's global role templates.

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
	user := "user_example" // string | username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalRoleAPI.RetrieveMemberRoleTemplates(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleAPI.RetrieveMemberRoleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveMemberRoleTemplates`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `GlobalRoleAPI.RetrieveMemberRoleTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMemberRoleTemplatesRequest struct via the builder pattern


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


## UpdateGlobalRole

> V1alpha2GlobalRole UpdateGlobalRole(ctx, globalrole).Body(body).Execute()

Update global role.

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
	globalrole := "globalrole_example" // string | global role name
	body := *openapiclient.NewV1alpha2GlobalRole([]openapiclient.V1PolicyRule{*openapiclient.NewV1PolicyRule([]string{"Verbs_example"})}) // V1alpha2GlobalRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalRoleAPI.UpdateGlobalRole(context.Background(), globalrole).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalRoleAPI.UpdateGlobalRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGlobalRole`: V1alpha2GlobalRole
	fmt.Fprintf(os.Stdout, "Response from `GlobalRoleAPI.UpdateGlobalRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalrole** | **string** | global role name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1alpha2GlobalRole**](V1alpha2GlobalRole.md) |  | 

### Return type

[**V1alpha2GlobalRole**](V1alpha2GlobalRole.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

