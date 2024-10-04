# \RepositoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepo**](RepositoryAPI.md#CreateRepo) | **Post** /kapis/openpitrix.io/v1/repos | Create repository in the specified workspace, repository used to store package of app
[**CreateWorkspaceRepo**](RepositoryAPI.md#CreateWorkspaceRepo) | **Post** /kapis/openpitrix.io/v1/workspaces/{workspace}/repos | Create repository in the specified workspace, repository used to store package of app
[**DeleteRepo**](RepositoryAPI.md#DeleteRepo) | **Delete** /kapis/openpitrix.io/v1/repos/{repo} | Delete the specified repository in the specified workspace
[**DeleteWorkspaceRepo**](RepositoryAPI.md#DeleteWorkspaceRepo) | **Delete** /kapis/openpitrix.io/v1/workspaces/{workspace}/repos/{repo} | Delete the specified repository in the specified workspace
[**DescribeRepo**](RepositoryAPI.md#DescribeRepo) | **Get** /kapis/openpitrix.io/v1/repos/{repo} | Describe the specified repository in the specified workspace
[**DescribeWorkspaceRepo**](RepositoryAPI.md#DescribeWorkspaceRepo) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/repos/{repo} | Describe the specified repository in the specified workspace
[**ListRepoEvents**](RepositoryAPI.md#ListRepoEvents) | **Get** /kapis/openpitrix.io/v1/repos/{repo}/events | Get repository events
[**ListRepos**](RepositoryAPI.md#ListRepos) | **Get** /kapis/openpitrix.io/v1/repos | List repositories in the specified workspace
[**ListWorkspaceRepoEvents**](RepositoryAPI.md#ListWorkspaceRepoEvents) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/repos/{repo}/events | Get repository events
[**ListWorkspaceRepos**](RepositoryAPI.md#ListWorkspaceRepos) | **Get** /kapis/openpitrix.io/v1/workspaces/{workspace}/repos | List repositories in the specified workspace
[**ModifyRepo**](RepositoryAPI.md#ModifyRepo) | **Patch** /kapis/openpitrix.io/v1/repos/{repo} | Patch the specified repository in the specified workspace
[**ModifyWorkspaceRepo**](RepositoryAPI.md#ModifyWorkspaceRepo) | **Patch** /kapis/openpitrix.io/v1/workspaces/{workspace}/repos/{repo} | Patch the specified repository in the specified workspace



## CreateRepo

> OpenpitrixCreateRepoResponse CreateRepo(ctx).Body(body).Validate(validate).Execute()

Create repository in the specified workspace, repository used to store package of app

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
	body := *openapiclient.NewOpenpitrixCreateRepoRequest([]string{"Providers_example"}) // OpenpitrixCreateRepoRequest | 
	validate := "validate_example" // string | Validate repository (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateRepo(context.Background()).Body(body).Validate(validate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRepo`: OpenpitrixCreateRepoResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateRepo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OpenpitrixCreateRepoRequest**](OpenpitrixCreateRepoRequest.md) |  | 
 **validate** | **string** | Validate repository | 

### Return type

[**OpenpitrixCreateRepoResponse**](OpenpitrixCreateRepoResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceRepo

> OpenpitrixCreateRepoResponse CreateWorkspaceRepo(ctx, workspace).Body(body).Validate(validate).Execute()

Create repository in the specified workspace, repository used to store package of app

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
	workspace := "workspace_example" // string | the name of the workspace
	body := *openapiclient.NewOpenpitrixCreateRepoRequest([]string{"Providers_example"}) // OpenpitrixCreateRepoRequest | 
	validate := "validate_example" // string | Validate repository (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateWorkspaceRepo(context.Background(), workspace).Body(body).Validate(validate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateWorkspaceRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceRepo`: OpenpitrixCreateRepoResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateWorkspaceRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixCreateRepoRequest**](OpenpitrixCreateRepoRequest.md) |  | 
 **validate** | **string** | Validate repository | 

### Return type

[**OpenpitrixCreateRepoResponse**](OpenpitrixCreateRepoResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepo

> ErrorsError DeleteRepo(ctx, repo).Execute()

Delete the specified repository in the specified workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.DeleteRepo(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRepo`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.DeleteRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepoRequest struct via the builder pattern


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


## DeleteWorkspaceRepo

> ErrorsError DeleteWorkspaceRepo(ctx, workspace, repo).Execute()

Delete the specified repository in the specified workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.DeleteWorkspaceRepo(context.Background(), workspace, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteWorkspaceRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceRepo`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.DeleteWorkspaceRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace id | 
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRepoRequest struct via the builder pattern


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


## DescribeRepo

> OpenpitrixRepo DescribeRepo(ctx, repo).Execute()

Describe the specified repository in the specified workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.DescribeRepo(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DescribeRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeRepo`: OpenpitrixRepo
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.DescribeRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenpitrixRepo**](OpenpitrixRepo.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspaceRepo

> OpenpitrixRepo DescribeWorkspaceRepo(ctx, workspace, repo).Execute()

Describe the specified repository in the specified workspace

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.DescribeWorkspaceRepo(context.Background(), workspace, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DescribeWorkspaceRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWorkspaceRepo`: OpenpitrixRepo
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.DescribeWorkspaceRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace id | 
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWorkspaceRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenpitrixRepo**](OpenpitrixRepo.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepoEvents

> ModelsPageableResponse ListRepoEvents(ctx, repo).Execute()

Get repository events

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.ListRepoEvents(context.Background(), repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.ListRepoEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRepoEvents`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.ListRepoEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRepoEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListRepos

> ModelsPageableResponse ListRepos(ctx).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

List repositories in the specified workspace

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
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.ListRepos(context.Background()).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.ListRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRepos`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.ListRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReposRequest struct via the builder pattern


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


## ListWorkspaceRepoEvents

> ModelsPageableResponse ListWorkspaceRepoEvents(ctx, workspace, repo).Execute()

Get repository events

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.ListWorkspaceRepoEvents(context.Background(), workspace, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.ListWorkspaceRepoEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceRepoEvents`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.ListWorkspaceRepoEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace id | 
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceRepoEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ListWorkspaceRepos

> ModelsPageableResponse ListWorkspaceRepos(ctx, workspace).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()

List repositories in the specified workspace

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
	workspace := "workspace_example" // string | the name of the workspace
	conditions := "conditions_example" // string | query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a (optional)
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")
	reverse := "reverse_example" // string | sort parameters, e.g. reverse=true (optional)
	orderBy := "orderBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.ListWorkspaceRepos(context.Background(), workspace).Conditions(conditions).Paging(paging).Reverse(reverse).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.ListWorkspaceRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceRepos`: ModelsPageableResponse
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.ListWorkspaceRepos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | the name of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceReposRequest struct via the builder pattern


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


## ModifyRepo

> ErrorsError ModifyRepo(ctx, repo).Body(body).Execute()

Patch the specified repository in the specified workspace

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
	body := *openapiclient.NewOpenpitrixModifyRepoRequest([]string{"Providers_example"}) // OpenpitrixModifyRepoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.ModifyRepo(context.Background(), repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.ModifyRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRepo`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.ModifyRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OpenpitrixModifyRepoRequest**](OpenpitrixModifyRepoRequest.md) |  | 

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


## ModifyWorkspaceRepo

> ErrorsError ModifyWorkspaceRepo(ctx, workspace, repo).Body(body).Execute()

Patch the specified repository in the specified workspace

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
	body := *openapiclient.NewOpenpitrixModifyRepoRequest([]string{"Providers_example"}) // OpenpitrixModifyRepoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.ModifyWorkspaceRepo(context.Background(), workspace, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.ModifyWorkspaceRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWorkspaceRepo`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.ModifyWorkspaceRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace id | 
**repo** | **string** | repo id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWorkspaceRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OpenpitrixModifyRepoRequest**](OpenpitrixModifyRepoRequest.md) |  | 

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

