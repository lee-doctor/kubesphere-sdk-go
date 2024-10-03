# \AppTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppTemplateAPI.md#CreateApp) | **Post** /kapis/openpitrix.io/v1/apps | Create a new app template
[**CreateAppVersion**](AppTemplateAPI.md#CreateAppVersion) | **Post** /kapis/openpitrix.io/v1/apps/{app}/versions | Create a new app template version
[**CreateWorkspaceApp**](AppTemplateAPI.md#CreateWorkspaceApp) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps | Create a new app template
[**CreateWorkspaceAppVersion**](AppTemplateAPI.md#CreateWorkspaceAppVersion) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/versions | Create a new app template version
[**DeleteApp**](AppTemplateAPI.md#DeleteApp) | **Delete** /kapis/openpitrix.io/v1/apps/{app} | Delete the specified app template
[**DeleteAppVersion**](AppTemplateAPI.md#DeleteAppVersion) | **Delete** /kapis/openpitrix.io/v1/apps/{app}/versions/{version} | Delete the specified app template version
[**DeleteWorkspaceApp**](AppTemplateAPI.md#DeleteWorkspaceApp) | **Delete** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app} | Delete the specified app template
[**DeleteWorkspaceAppVersion**](AppTemplateAPI.md#DeleteWorkspaceAppVersion) | **Delete** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/versions/{version} | Delete the specified app template version
[**DescribeApp**](AppTemplateAPI.md#DescribeApp) | **Get** /kapis/openpitrix.io/v1/apps/{app} | Describe the specified app template
[**DescribeAppVersion**](AppTemplateAPI.md#DescribeAppVersion) | **Get** /kapis/openpitrix.io/v1/apps/{app}/versions/{version} | Describe the specified app template version
[**DescribeWorkspaceApp**](AppTemplateAPI.md#DescribeWorkspaceApp) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app} | Describe the specified app template
[**DescribeWorkspaceAppVersion**](AppTemplateAPI.md#DescribeWorkspaceAppVersion) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/versions/{version} | Describe the specified app template version
[**GetAppVersionFiles**](AppTemplateAPI.md#GetAppVersionFiles) | **Get** /kapis/openpitrix.io/v1/apps/{app}/versions/{version}/files | Get app template package files
[**GetAppVersionPackage**](AppTemplateAPI.md#GetAppVersionPackage) | **Get** /kapis/openpitrix.io/v1/apps/{app}/versions/{version}/package | Get packages of version-specific app
[**ListAppVersions**](AppTemplateAPI.md#ListAppVersions) | **Get** /kapis/openpitrix.io/v1/apps/{app}/versions | Get active versions of app, can filter with these fields(version_id, app_id, name, owner, description, package_name, status, type), default return all active app versions
[**ListApps**](AppTemplateAPI.md#ListApps) | **Get** /kapis/openpitrix.io/v1/apps | List app templates
[**ListWorkspaceAppVersions**](AppTemplateAPI.md#ListWorkspaceAppVersions) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/versions | Get active versions of app, can filter with these fields(version_id, app_id, name, owner, description, package_name, status, type), default return all active app versions
[**ListWorkspaceApps**](AppTemplateAPI.md#ListWorkspaceApps) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps | List app templates in the specified workspace.
[**ModifyApp**](AppTemplateAPI.md#ModifyApp) | **Patch** /kapis/openpitrix.io/v1/apps/{app} | Patch the specified app template
[**ModifyAppVersion**](AppTemplateAPI.md#ModifyAppVersion) | **Patch** /kapis/openpitrix.io/v1/apps/{app}/versions/{version} | Patch the specified app template version
[**ModifyWorkspaceApp**](AppTemplateAPI.md#ModifyWorkspaceApp) | **Patch** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app} | Patch the specified app template
[**ModifyWorkspaceAppVersion**](AppTemplateAPI.md#ModifyWorkspaceAppVersion) | **Patch** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/versions/{version} | Patch the specified app template version



## CreateApp

> OpenpitrixCreateAppResponse CreateApp(ctx, app).Body(body).Execute()

Create a new app template

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
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixCreateAppRequest() // OpenpitrixCreateAppRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.CreateApp(context.Background(), app).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.CreateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApp`: OpenpitrixCreateAppResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.CreateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixCreateAppRequest**](OpenpitrixCreateAppRequest.md) |  | 

### Return type

[**OpenpitrixCreateAppResponse**](OpenpitrixCreateAppResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppVersion

> OpenpitrixCreateAppVersionResponse CreateAppVersion(ctx, app).Body(body).Validate(validate).Execute()

Create a new app template version

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
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixCreateAppVersionRequest() // OpenpitrixCreateAppVersionRequest | 
	validate := "validate_example" // string | Validate format of package(pack by op tool) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.CreateAppVersion(context.Background(), app).Body(body).Validate(validate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.CreateAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppVersion`: OpenpitrixCreateAppVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.CreateAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixCreateAppVersionRequest**](OpenpitrixCreateAppVersionRequest.md) |  | 
 **validate** | **string** | Validate format of package(pack by op tool) | 

### Return type

[**OpenpitrixCreateAppVersionResponse**](OpenpitrixCreateAppVersionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceApp

> OpenpitrixCreateAppResponse CreateWorkspaceApp(ctx, workspace, app).Body(body).Execute()

Create a new app template

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
	workspace := "workspace_example" // string | workspace template id
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixCreateAppRequest() // OpenpitrixCreateAppRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.CreateWorkspaceApp(context.Background(), workspace, app).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.CreateWorkspaceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceApp`: OpenpitrixCreateAppResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.CreateWorkspaceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace template id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OpenpitrixCreateAppRequest**](OpenpitrixCreateAppRequest.md) |  | 

### Return type

[**OpenpitrixCreateAppResponse**](OpenpitrixCreateAppResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceAppVersion

> OpenpitrixCreateAppVersionResponse CreateWorkspaceAppVersion(ctx, workspace, app).Body(body).Validate(validate).Execute()

Create a new app template version

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
	workspace := "workspace_example" // string | workspace template id
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixCreateAppVersionRequest() // OpenpitrixCreateAppVersionRequest | 
	validate := "validate_example" // string | Validate format of package(pack by op tool) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.CreateWorkspaceAppVersion(context.Background(), workspace, app).Body(body).Validate(validate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.CreateWorkspaceAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceAppVersion`: OpenpitrixCreateAppVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.CreateWorkspaceAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace template id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceAppVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OpenpitrixCreateAppVersionRequest**](OpenpitrixCreateAppVersionRequest.md) |  | 
 **validate** | **string** | Validate format of package(pack by op tool) | 

### Return type

[**OpenpitrixCreateAppVersionResponse**](OpenpitrixCreateAppVersionResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> ErrorsError DeleteApp(ctx, app).Execute()

Delete the specified app template

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
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DeleteApp(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DeleteApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApp`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DeleteApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## DeleteAppVersion

> ErrorsError DeleteAppVersion(ctx, version, app).Execute()

Delete the specified app template version

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DeleteAppVersion(context.Background(), version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DeleteAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppVersion`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DeleteAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppVersionRequest struct via the builder pattern


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


## DeleteWorkspaceApp

> ErrorsError DeleteWorkspaceApp(ctx, workspace, app).Execute()

Delete the specified app template

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
	workspace := "workspace_example" // string | workspace template id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DeleteWorkspaceApp(context.Background(), workspace, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DeleteWorkspaceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceApp`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DeleteWorkspaceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace template id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceAppRequest struct via the builder pattern


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


## DeleteWorkspaceAppVersion

> ErrorsError DeleteWorkspaceAppVersion(ctx, workspace, version, app).Execute()

Delete the specified app template version

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
	workspace := "workspace_example" // string | app template workspace id
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DeleteWorkspaceAppVersion(context.Background(), workspace, version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DeleteWorkspaceAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceAppVersion`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DeleteWorkspaceAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | app template workspace id | 
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceAppVersionRequest struct via the builder pattern


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


## DescribeApp

> OpenpitrixAppVersion DescribeApp(ctx, app).Execute()

Describe the specified app template

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
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DescribeApp(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DescribeApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeApp`: OpenpitrixAppVersion
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DescribeApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenpitrixAppVersion**](OpenpitrixAppVersion.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAppVersion

> OpenpitrixAppVersion DescribeAppVersion(ctx, version, app).Execute()

Describe the specified app template version

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DescribeAppVersion(context.Background(), version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DescribeAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAppVersion`: OpenpitrixAppVersion
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DescribeAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAppVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenpitrixAppVersion**](OpenpitrixAppVersion.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspaceApp

> OpenpitrixAppVersion DescribeWorkspaceApp(ctx, workspace, app).Execute()

Describe the specified app template

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
	workspace := "workspace_example" // string | workspace template id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DescribeWorkspaceApp(context.Background(), workspace, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DescribeWorkspaceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceApp`: OpenpitrixAppVersion
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DescribeWorkspaceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace template id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenpitrixAppVersion**](OpenpitrixAppVersion.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspaceAppVersion

> OpenpitrixAppVersion DescribeWorkspaceAppVersion(ctx, workspace, version, app).Execute()

Describe the specified app template version

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
	workspace := "workspace_example" // string | app template workspace id
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.DescribeWorkspaceAppVersion(context.Background(), workspace, version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.DescribeWorkspaceAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceAppVersion`: OpenpitrixAppVersion
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.DescribeWorkspaceAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | app template workspace id | 
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceAppVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OpenpitrixAppVersion**](OpenpitrixAppVersion.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppVersionFiles

> OpenpitrixGetAppVersionPackageFilesResponse GetAppVersionFiles(ctx, version, app).Execute()

Get app template package files

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.GetAppVersionFiles(context.Background(), version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.GetAppVersionFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppVersionFiles`: OpenpitrixGetAppVersionPackageFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.GetAppVersionFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppVersionFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenpitrixGetAppVersionPackageFilesResponse**](OpenpitrixGetAppVersionPackageFilesResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppVersionPackage

> OpenpitrixGetAppVersionPackageResponse GetAppVersionPackage(ctx, version, app).Execute()

Get packages of version-specific app

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.GetAppVersionPackage(context.Background(), version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.GetAppVersionPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppVersionPackage`: OpenpitrixGetAppVersionPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.GetAppVersionPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppVersionPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenpitrixGetAppVersionPackageResponse**](OpenpitrixGetAppVersionPackageResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppVersions

> ModelsPageableResponse ListAppVersions(ctx, app).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

Get active versions of app, can filter with these fields(version_id, app_id, name, owner, description, package_name, status, type), default return all active app versions

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
	app := "app_example" // string | app template id
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ListAppVersions(context.Background(), app).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ListAppVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppVersions`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ListAppVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conditions** | **string** | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]
 **reverse** | **string** | sort parameters, e.g. reverse&#x3D;true | 
 **orderBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> ModelsPageableResponse ListApps(ctx).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

List app templates

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
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ListApps(context.Background()).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ListApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApps`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ListApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditions** | **string** | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]
 **reverse** | **string** | sort parameters, e.g. reverse&#x3D;true | 
 **orderBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceAppVersions

> ModelsPageableResponse ListWorkspaceAppVersions(ctx, workspace, app).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

Get active versions of app, can filter with these fields(version_id, app_id, name, owner, description, package_name, status, type), default return all active app versions

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
	workspace := "workspace_example" // string | workspace template id
	app := "app_example" // string | app template id
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ListWorkspaceAppVersions(context.Background(), workspace, app).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ListWorkspaceAppVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceAppVersions`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ListWorkspaceAppVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace template id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceAppVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conditions** | **string** | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]
 **reverse** | **string** | sort parameters, e.g. reverse&#x3D;true | 
 **orderBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceApps

> ModelsPageableResponse ListWorkspaceApps(ctx, workspace).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

List app templates in the specified workspace.

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
	workspace := "workspace_example" // string | workspace name
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ListWorkspaceApps(context.Background(), workspace).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ListWorkspaceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceApps`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ListWorkspaceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **conditions** | **string** | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]
 **reverse** | **string** | sort parameters, e.g. reverse&#x3D;true | 
 **orderBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**ModelsPageableResponse**](ModelsPageableResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyApp

> ErrorsError ModifyApp(ctx, app).Body(body).Execute()

Patch the specified app template

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
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixModifyAppVersionRequest() // OpenpitrixModifyAppVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ModifyApp(context.Background(), app).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ModifyApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyApp`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ModifyApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixModifyAppVersionRequest**](OpenpitrixModifyAppVersionRequest.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAppVersion

> ErrorsError ModifyAppVersion(ctx, version, app).Body(body).Execute()

Patch the specified app template version

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixModifyAppVersionRequest() // OpenpitrixModifyAppVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ModifyAppVersion(context.Background(), version, app).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ModifyAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAppVersion`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ModifyAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAppVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OpenpitrixModifyAppVersionRequest**](OpenpitrixModifyAppVersionRequest.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWorkspaceApp

> ErrorsError ModifyWorkspaceApp(ctx, workspace, app).Body(body).Execute()

Patch the specified app template

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
	workspace := "workspace_example" // string | workspace template id
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixModifyAppVersionRequest() // OpenpitrixModifyAppVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ModifyWorkspaceApp(context.Background(), workspace, app).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ModifyWorkspaceApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWorkspaceApp`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ModifyWorkspaceApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace template id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWorkspaceAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OpenpitrixModifyAppVersionRequest**](OpenpitrixModifyAppVersionRequest.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWorkspaceAppVersion

> ErrorsError ModifyWorkspaceAppVersion(ctx, workspace, version, app).Body(body).Execute()

Patch the specified app template version

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
	workspace := "workspace_example" // string | app template workspace id
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id
	body := *openapiclient.NewOpenpitrixModifyAppVersionRequest() // OpenpitrixModifyAppVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppTemplateAPI.ModifyWorkspaceAppVersion(context.Background(), workspace, version, app).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppTemplateAPI.ModifyWorkspaceAppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWorkspaceAppVersion`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppTemplateAPI.ModifyWorkspaceAppVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | app template workspace id | 
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWorkspaceAppVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**OpenpitrixModifyAppVersionRequest**](OpenpitrixModifyAppVersionRequest.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

