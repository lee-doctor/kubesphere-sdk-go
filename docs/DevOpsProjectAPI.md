# \DevOpsProjectAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredential**](DevOpsProjectAPI.md#CreateCredential) | **Post** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/credentials | create the credential of the specified devops for the current user
[**CreateDevOpsProject**](DevOpsProjectAPI.md#CreateDevOpsProject) | **Post** /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/devops | Create the devopsproject of the specified workspace for the current user
[**CreatePipeline**](DevOpsProjectAPI.md#CreatePipeline) | **Post** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/pipelines | create the pipeline of the specified devops for the current user
[**DeleteDevOpsProject**](DevOpsProjectAPI.md#DeleteDevOpsProject) | **Delete** /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/devops/{devops} | Get the devopsproject of the specified workspace for the current user
[**GetCredential**](DevOpsProjectAPI.md#GetCredential) | **Get** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/credentials/{credential} | get the credential of the specified devops for the current user
[**GetDevOpsProject**](DevOpsProjectAPI.md#GetDevOpsProject) | **Get** /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/devops/{devops} | Get the devopsproject of the specified workspace for the current user
[**GetPipelineByName**](DevOpsProjectAPI.md#GetPipelineByName) | **Get** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/pipelines/{pipeline} | get the pipeline of the specified devops for the current user
[**ListCredential**](DevOpsProjectAPI.md#ListCredential) | **Get** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/credentials | list the credentials of the specified devops for the current user
[**ListDevOpsProject**](DevOpsProjectAPI.md#ListDevOpsProject) | **Get** /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/devops | List the devopsproject of the specified workspace for the current user
[**ListDevOpsProjects**](DevOpsProjectAPI.md#ListDevOpsProjects) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/workspacemembers/{workspacemember}/devops | List the devops projects of specified workspace for the workspace member
[**ListPipeline**](DevOpsProjectAPI.md#ListPipeline) | **Get** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/pipelines | list the pipelines of the specified devops for the current user
[**ListUserDevOpsProjects**](DevOpsProjectAPI.md#ListUserDevOpsProjects) | **Get** /kapis/tenant.kubesphere.io/v1alpha2/workspaces/{workspace}/devops | List the devops projects of the specified workspace for the current user
[**UpdateCredential**](DevOpsProjectAPI.md#UpdateCredential) | **Put** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/credentials/{credential} | put the credential of the specified devops for the current user
[**UpdateDevOpsProject**](DevOpsProjectAPI.md#UpdateDevOpsProject) | **Put** /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/devops/{devops} | Put the devopsproject of the specified workspace for the current user
[**UpdatePipeline**](DevOpsProjectAPI.md#UpdatePipeline) | **Put** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/pipelines/{pipeline} | put the pipeline of the specified devops for the current user



## CreateCredential

> []V1alpha3Pipeline CreateCredential(ctx, devops).Execute()

create the credential of the specified devops for the current user

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
	devops := "devops_example" // string | devops name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.CreateCredential(context.Background(), devops).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.CreateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredential`: []V1alpha3Pipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.CreateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]V1alpha3Pipeline**](V1alpha3Pipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevOpsProject

> []V1alpha3DevOpsProject CreateDevOpsProject(ctx, workspace).Execute()

Create the devopsproject of the specified workspace for the current user

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.CreateDevOpsProject(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.CreateDevOpsProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDevOpsProject`: []V1alpha3DevOpsProject
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.CreateDevOpsProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDevOpsProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]V1alpha3DevOpsProject**](V1alpha3DevOpsProject.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipeline

> []V1alpha3Pipeline CreatePipeline(ctx, devops).Execute()

create the pipeline of the specified devops for the current user

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
	devops := "devops_example" // string | devops name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.CreatePipeline(context.Background(), devops).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.CreatePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePipeline`: []V1alpha3Pipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.CreatePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]V1alpha3Pipeline**](V1alpha3Pipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevOpsProject

> []V1alpha3DevOpsProject DeleteDevOpsProject(ctx, workspace, devops).Execute()

Get the devopsproject of the specified workspace for the current user

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
	devops := "devops_example" // string | project name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.DeleteDevOpsProject(context.Background(), workspace, devops).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.DeleteDevOpsProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDevOpsProject`: []V1alpha3DevOpsProject
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.DeleteDevOpsProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**devops** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevOpsProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1alpha3DevOpsProject**](V1alpha3DevOpsProject.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredential

> []V1Secret GetCredential(ctx, devops, credential).Execute()

get the credential of the specified devops for the current user

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
	devops := "devops_example" // string | project name
	credential := "credential_example" // string | pipeline name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.GetCredential(context.Background(), devops, credential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.GetCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredential`: []V1Secret
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.GetCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | project name | 
**credential** | **string** | pipeline name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1Secret**](V1Secret.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevOpsProject

> []V1alpha3DevOpsProject GetDevOpsProject(ctx, workspace, devops).Execute()

Get the devopsproject of the specified workspace for the current user

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
	devops := "devops_example" // string | project name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.GetDevOpsProject(context.Background(), workspace, devops).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.GetDevOpsProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevOpsProject`: []V1alpha3DevOpsProject
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.GetDevOpsProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**devops** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevOpsProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1alpha3DevOpsProject**](V1alpha3DevOpsProject.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineByName

> []V1alpha3Pipeline GetPipelineByName(ctx, devops, pipeline).Execute()

get the pipeline of the specified devops for the current user

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
	devops := "devops_example" // string | project name
	pipeline := "pipeline_example" // string | pipeline name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.GetPipelineByName(context.Background(), devops, pipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.GetPipelineByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineByName`: []V1alpha3Pipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.GetPipelineByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | project name | 
**pipeline** | **string** | pipeline name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1alpha3Pipeline**](V1alpha3Pipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> ApiListResult ListCredential(ctx, devops).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()

list the credentials of the specified devops for the current user

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
	devops := "devops_example" // string | devops name
	name := "name_example" // string | name used to do filtering (optional)
	page := "page_example" // string | page (optional) (default to "page=1")
	limit := "limit_example" // string | limit (optional)
	ascending := "ascending_example" // string | sort parameters, e.g. ascending=false (optional) (default to "ascending=false")
	sortBy := "sortBy_example" // string | sort parameters, e.g. orderBy=createTime (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.ListCredential(context.Background(), devops).Name(name).Page(page).Limit(limit).Ascending(ascending).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.ListCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCredential`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.ListCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | name used to do filtering | 
 **page** | **string** | page | [default to &quot;page&#x3D;1&quot;]
 **limit** | **string** | limit | 
 **ascending** | **string** | sort parameters, e.g. ascending&#x3D;false | [default to &quot;ascending&#x3D;false&quot;]
 **sortBy** | **string** | sort parameters, e.g. orderBy&#x3D;createTime | 

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


## ListDevOpsProject

> ApiListResult ListDevOpsProject(ctx, workspace).Paging(paging).Execute()

List the devopsproject of the specified workspace for the current user

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
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.ListDevOpsProject(context.Background(), workspace).Paging(paging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.ListDevOpsProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevOpsProject`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.ListDevOpsProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDevOpsProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]

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


## ListDevOpsProjects

> V1Namespace ListDevOpsProjects(ctx, workspace, workspacemember).Body(body).Execute()

List the devops projects of specified workspace for the workspace member

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
	workspacemember := "workspacemember_example" // string | workspacemember username
	body := *openapiclient.NewV1Namespace() // V1Namespace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.ListDevOpsProjects(context.Background(), workspace, workspacemember).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.ListDevOpsProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevOpsProjects`: V1Namespace
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.ListDevOpsProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**workspacemember** | **string** | workspacemember username | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDevOpsProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Namespace**](V1Namespace.md) |  | 

### Return type

[**V1Namespace**](V1Namespace.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipeline

> ApiListResult ListPipeline(ctx, devops).Paging(paging).Execute()

list the pipelines of the specified devops for the current user

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
	devops := "devops_example" // string | devops name
	paging := "paging_example" // string | paging query, e.g. limit=100,page=1 (optional) (default to "limit=10,page=1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.ListPipeline(context.Background(), devops).Paging(paging).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.ListPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPipeline`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.ListPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | devops name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paging** | **string** | paging query, e.g. limit&#x3D;100,page&#x3D;1 | [default to &quot;limit&#x3D;10,page&#x3D;1&quot;]

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


## ListUserDevOpsProjects

> ApiListResult ListUserDevOpsProjects(ctx, workspace).Execute()

List the devops projects of the specified workspace for the current user

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.ListUserDevOpsProjects(context.Background(), workspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.ListUserDevOpsProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserDevOpsProjects`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.ListUserDevOpsProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserDevOpsProjectsRequest struct via the builder pattern


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


## UpdateCredential

> []V1Secret UpdateCredential(ctx, devops, credential).Execute()

put the credential of the specified devops for the current user

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
	devops := "devops_example" // string | project name
	credential := "credential_example" // string | credential name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.UpdateCredential(context.Background(), devops, credential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.UpdateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredential`: []V1Secret
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | project name | 
**credential** | **string** | credential name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1Secret**](V1Secret.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevOpsProject

> []V1alpha3DevOpsProject UpdateDevOpsProject(ctx, workspace, devops).Execute()

Put the devopsproject of the specified workspace for the current user

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
	devops := "devops_example" // string | project name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.UpdateDevOpsProject(context.Background(), workspace, devops).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.UpdateDevOpsProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDevOpsProject`: []V1alpha3DevOpsProject
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.UpdateDevOpsProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | workspace name | 
**devops** | **string** | project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDevOpsProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1alpha3DevOpsProject**](V1alpha3DevOpsProject.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipeline

> []V1alpha3Pipeline UpdatePipeline(ctx, devops, pipeline).Execute()

put the pipeline of the specified devops for the current user

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
	devops := "devops_example" // string | project name
	pipeline := "pipeline_example" // string | pipeline name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsProjectAPI.UpdatePipeline(context.Background(), devops, pipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsProjectAPI.UpdatePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePipeline`: []V1alpha3Pipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsProjectAPI.UpdatePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | project name | 
**pipeline** | **string** | pipeline name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1alpha3Pipeline**](V1alpha3Pipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

