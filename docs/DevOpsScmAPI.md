# \DevOpsScmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSCMServers**](DevOpsScmAPI.md#CreateSCMServers) | **Post** /kapis/devops.kubesphere.io/v1alpha2/scms/{scm}/servers | Create scm server in the jenkins.
[**GetOrgRepo**](DevOpsScmAPI.md#GetOrgRepo) | **Get** /kapis/devops.kubesphere.io/v1alpha2/scms/{scm}/organizations/{organization}/repositories | List all repositories in the specified organization.
[**GetSCMOrg**](DevOpsScmAPI.md#GetSCMOrg) | **Get** /kapis/devops.kubesphere.io/v1alpha2/scms/{scm}/organizations | List all organizations of the specified source configuration management (SCM) such as Github.
[**GetSCMServers**](DevOpsScmAPI.md#GetSCMServers) | **Get** /kapis/devops.kubesphere.io/v1alpha2/scms/{scm}/servers | List all servers in the jenkins.
[**Validate**](DevOpsScmAPI.md#Validate) | **Post** /kapis/devops.kubesphere.io/v1alpha2/scms/{scm}/verify | Validate the access token of the specified source configuration management (SCM) such as Github



## CreateSCMServers

> DevopsSCMServer CreateSCMServers(ctx, scm).Body(body).Execute()

Create scm server in the jenkins.

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
	scm := "scm_example" // string | The ID of the source configuration management (SCM).
	body := *openapiclient.NewDevopsCreateScmServerReq() // DevopsCreateScmServerReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsScmAPI.CreateSCMServers(context.Background(), scm).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsScmAPI.CreateSCMServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSCMServers`: DevopsSCMServer
	fmt.Fprintf(os.Stdout, "Response from `DevOpsScmAPI.CreateSCMServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scm** | **string** | The ID of the source configuration management (SCM). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSCMServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DevopsCreateScmServerReq**](DevopsCreateScmServerReq.md) |  | 

### Return type

[**DevopsSCMServer**](DevopsSCMServer.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgRepo

> DevopsOrgRepo GetOrgRepo(ctx, scm, organization).CredentialId(credentialId).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all repositories in the specified organization.

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
	scm := "scm_example" // string | The ID of the source configuration management (SCM).
	organization := "organization_example" // string | organization ID, such as github username.
	credentialId := "credentialId_example" // string | credential ID for SCM.
	pageNumber := "pageNumber_example" // string | page number.
	pageSize := "pageSize_example" // string | the item count of one page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsScmAPI.GetOrgRepo(context.Background(), scm, organization).CredentialId(credentialId).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsScmAPI.GetOrgRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgRepo`: DevopsOrgRepo
	fmt.Fprintf(os.Stdout, "Response from `DevOpsScmAPI.GetOrgRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scm** | **string** | The ID of the source configuration management (SCM). | 
**organization** | **string** | organization ID, such as github username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialId** | **string** | credential ID for SCM. | 
 **pageNumber** | **string** | page number. | 
 **pageSize** | **string** | the item count of one page. | 

### Return type

[**DevopsOrgRepo**](DevopsOrgRepo.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSCMOrg

> []DevopsSCMOrg GetSCMOrg(ctx, scm).CredentialId(credentialId).Execute()

List all organizations of the specified source configuration management (SCM) such as Github.

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
	scm := "scm_example" // string | the ID of the source configuration management (SCM).
	credentialId := "credentialId_example" // string | credential ID for source configuration management (SCM).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsScmAPI.GetSCMOrg(context.Background(), scm).CredentialId(credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsScmAPI.GetSCMOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSCMOrg`: []DevopsSCMOrg
	fmt.Fprintf(os.Stdout, "Response from `DevOpsScmAPI.GetSCMOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scm** | **string** | the ID of the source configuration management (SCM). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSCMOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialId** | **string** | credential ID for source configuration management (SCM). | 

### Return type

[**[]DevopsSCMOrg**](DevopsSCMOrg.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSCMServers

> []DevopsSCMServer GetSCMServers(ctx, scm).Execute()

List all servers in the jenkins.

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
	scm := "scm_example" // string | The ID of the source configuration management (SCM).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsScmAPI.GetSCMServers(context.Background(), scm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsScmAPI.GetSCMServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSCMServers`: []DevopsSCMServer
	fmt.Fprintf(os.Stdout, "Response from `DevOpsScmAPI.GetSCMServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scm** | **string** | The ID of the source configuration management (SCM). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSCMServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DevopsSCMServer**](DevopsSCMServer.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validate

> DevopsValidates Validate(ctx, scm).Execute()

Validate the access token of the specified source configuration management (SCM) such as Github

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
	scm := "scm_example" // string | the ID of the source configuration management (SCM).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsScmAPI.Validate(context.Background(), scm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsScmAPI.Validate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Validate`: DevopsValidates
	fmt.Fprintf(os.Stdout, "Response from `DevOpsScmAPI.Validate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scm** | **string** | the ID of the source configuration management (SCM). | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DevopsValidates**](DevopsValidates.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

