# \AppManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoAppAction**](AppManagementAPI.md#DoAppAction) | **Post** /kapis/openpitrix.io/v1/apps/{app}/action | Perform recover or suspend operation on app
[**DoAppVersionAction**](AppManagementAPI.md#DoAppVersionAction) | **Post** /kapis/openpitrix.io/v1/apps/{app}/versions/{version}/action | Perform submit or other operations on app
[**DoRepoAction**](AppManagementAPI.md#DoRepoAction) | **Post** /kapis/openpitrix.io/v1/repos/{repo}/action | Start index repository event
[**DoWorkspaceAppVersionAction**](AppManagementAPI.md#DoWorkspaceAppVersionAction) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/versions/{version}/action | Perform submit or other operations on app
[**DoWorkspaceRepoAction**](AppManagementAPI.md#DoWorkspaceRepoAction) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/repos/{repo}/action | Start index repository event
[**DoWorkspacesAppAction**](AppManagementAPI.md#DoWorkspacesAppAction) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/action | Perform recover or suspend operation on app
[**ListAppVersionAudits**](AppManagementAPI.md#ListAppVersionAudits) | **Get** /kapis/openpitrix.io/v1/apps/{app}/audits | List audits information of the specific app template
[**ListReviews**](AppManagementAPI.md#ListReviews) | **Get** /kapis/openpitrix.io/v1/reviews | Get reviews of version-specific app
[**ListVersionSpecificAppVersionAudits**](AppManagementAPI.md#ListVersionSpecificAppVersionAudits) | **Get** /kapis/openpitrix.io/v1/apps/{app}/versions/{version}/audits | List audits information of version-specific app template
[**ListWorkspaceAppVersionAudits**](AppManagementAPI.md#ListWorkspaceAppVersionAudits) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/apps/{app}/versions/{version}/audits | List audits information of version-specific app template



## DoAppAction

> ErrorsError DoAppAction(ctx, version, app).Execute()

Perform recover or suspend operation on app

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.DoAppAction(context.Background(), version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.DoAppAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoAppAction`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.DoAppAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoAppActionRequest struct via the builder pattern


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


## DoAppVersionAction

> ErrorsError DoAppVersionAction(ctx, version, app).Execute()

Perform submit or other operations on app

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.DoAppVersionAction(context.Background(), version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.DoAppVersionAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoAppVersionAction`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.DoAppVersionAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoAppVersionActionRequest struct via the builder pattern


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


## DoRepoAction

> ErrorsError DoRepoAction(ctx, repo).Body(body).Execute()

Start index repository event

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
	repo := "repo_example" // string | repo id
	body := *openapiclient.NewOpenpitrixRepoActionRequest("Action_example") // OpenpitrixRepoActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.DoRepoAction(context.Background(), repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.DoRepoAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoRepoAction`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.DoRepoAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoRepoActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixRepoActionRequest**](OpenpitrixRepoActionRequest.md) |  | 

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


## DoWorkspaceAppVersionAction

> ErrorsError DoWorkspaceAppVersionAction(ctx, workspace, version, app).Execute()

Perform submit or other operations on app

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
	workspace := "workspace_example" // string | app template workspace id
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.DoWorkspaceAppVersionAction(context.Background(), workspace, version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.DoWorkspaceAppVersionAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoWorkspaceAppVersionAction`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.DoWorkspaceAppVersionAction`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDoWorkspaceAppVersionActionRequest struct via the builder pattern


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


## DoWorkspaceRepoAction

> ErrorsError DoWorkspaceRepoAction(ctx, workspace, repo).Body(body).Execute()

Start index repository event

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
	workspace := "workspace_example" // string | workspace id
	repo := "repo_example" // string | repo id
	body := *openapiclient.NewOpenpitrixRepoActionRequest("Action_example") // OpenpitrixRepoActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.DoWorkspaceRepoAction(context.Background(), workspace, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.DoWorkspaceRepoAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoWorkspaceRepoAction`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.DoWorkspaceRepoAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace id | 
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoWorkspaceRepoActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OpenpitrixRepoActionRequest**](OpenpitrixRepoActionRequest.md) |  | 

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


## DoWorkspacesAppAction

> ErrorsError DoWorkspacesAppAction(ctx, version, workspace, app).Execute()

Perform recover or suspend operation on app

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
	version := "version_example" // string | app template version id
	workspace := "workspace_example" // string | workspace template id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.DoWorkspacesAppAction(context.Background(), version, workspace, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.DoWorkspacesAppAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoWorkspacesAppAction`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.DoWorkspacesAppAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**workspace** | **string** | workspace template id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoWorkspacesAppActionRequest struct via the builder pattern


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


## ListAppVersionAudits

> OpenpitrixAppVersionAudit ListAppVersionAudits(ctx, app).Execute()

List audits information of the specific app template

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
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.ListAppVersionAudits(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.ListAppVersionAudits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppVersionAudits`: OpenpitrixAppVersionAudit
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.ListAppVersionAudits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppVersionAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenpitrixAppVersionAudit**](OpenpitrixAppVersionAudit.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReviews

> OpenpitrixAppVersionReview ListReviews(ctx).Conditions(conditions).Paging(paging).Execute()

Get reviews of version-specific app

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
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.ListReviews(context.Background()).Conditions(conditions).Paging(paging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.ListReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReviews`: OpenpitrixAppVersionReview
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.ListReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditions** | **string** | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a | 
 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]

### Return type

[**OpenpitrixAppVersionReview**](OpenpitrixAppVersionReview.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersionSpecificAppVersionAudits

> OpenpitrixAppVersionAudit ListVersionSpecificAppVersionAudits(ctx, version, app).Execute()

List audits information of version-specific app template

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
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.ListVersionSpecificAppVersionAudits(context.Background(), version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.ListVersionSpecificAppVersionAudits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVersionSpecificAppVersionAudits`: OpenpitrixAppVersionAudit
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.ListVersionSpecificAppVersionAudits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | app template version id | 
**app** | **string** | app template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionSpecificAppVersionAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenpitrixAppVersionAudit**](OpenpitrixAppVersionAudit.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceAppVersionAudits

> OpenpitrixAppVersionAudit ListWorkspaceAppVersionAudits(ctx, workspace, version, app).Execute()

List audits information of version-specific app template

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
	workspace := "workspace_example" // string | app template workspace id
	version := "version_example" // string | app template version id
	app := "app_example" // string | app template id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementAPI.ListWorkspaceAppVersionAudits(context.Background(), workspace, version, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.ListWorkspaceAppVersionAudits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceAppVersionAudits`: OpenpitrixAppVersionAudit
	fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.ListWorkspaceAppVersionAudits`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListWorkspaceAppVersionAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OpenpitrixAppVersionAudit**](OpenpitrixAppVersionAudit.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

