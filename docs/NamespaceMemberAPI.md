# \NamespaceMemberAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAllNamespaceMembers**](NamespaceMemberAPI.md#CreateAllNamespaceMembers) | **Post** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/members | Add members to the namespace in bulk.
[**DescribeNamespaceMember**](NamespaceMemberAPI.md#DescribeNamespaceMember) | **Get** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/members/{member} | Retrieve the role of the specified member.
[**ListNamespaceMembers**](NamespaceMemberAPI.md#ListNamespaceMembers) | **Get** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/members | List all members in the specified namespace.
[**RemoveNamespaceMember**](NamespaceMemberAPI.md#RemoveNamespaceMember) | **Delete** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/members/{member} | Delete a member from the namespace.
[**UpdateNamespaceMember**](NamespaceMemberAPI.md#UpdateNamespaceMember) | **Put** /kapis/iam.kubesphere.io/v1alpha2/namespaces/{namespace}/members/{member} | Update the role bind of the member.



## CreateAllNamespaceMembers

> []V1alpha2Member CreateAllNamespaceMembers(ctx, namespace).Body(body).Execute()

Add members to the namespace in bulk.

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
	namespace := "namespace_example" // string | namespace
	body := []openapiclient.V1alpha2Member{*openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example")} // []V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMemberAPI.CreateAllNamespaceMembers(context.Background(), namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMemberAPI.CreateAllNamespaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAllNamespaceMembers`: []V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMemberAPI.CreateAllNamespaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllNamespaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]V1alpha2Member**](V1alpha2Member.md) |  | 

### Return type

[**[]V1alpha2Member**](V1alpha2Member.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeNamespaceMember

> V1alpha2User DescribeNamespaceMember(ctx, namespace, member).Execute()

Retrieve the role of the specified member.

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
	namespace := "namespace_example" // string | namespace
	member := "member_example" // string | namespace member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMemberAPI.DescribeNamespaceMember(context.Background(), namespace, member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMemberAPI.DescribeNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeNamespaceMember`: V1alpha2User
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMemberAPI.DescribeNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**member** | **string** | namespace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeNamespaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1alpha2User**](V1alpha2User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaceMembers

> ApiListResult ListNamespaceMembers(ctx, namespace).Execute()

List all members in the specified namespace.

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
	namespace := "namespace_example" // string | namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMemberAPI.ListNamespaceMembers(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMemberAPI.ListNamespaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaceMembers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMemberAPI.ListNamespaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespaceMembersRequest struct via the builder pattern


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


## RemoveNamespaceMember

> ErrorsError RemoveNamespaceMember(ctx, namespace, member).Execute()

Delete a member from the namespace.

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
	namespace := "namespace_example" // string | namespace
	member := "member_example" // string | namespace member's username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMemberAPI.RemoveNamespaceMember(context.Background(), namespace, member).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMemberAPI.RemoveNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNamespaceMember`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMemberAPI.RemoveNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**member** | **string** | namespace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNamespaceMemberRequest struct via the builder pattern


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


## UpdateNamespaceMember

> V1alpha2Member UpdateNamespaceMember(ctx, namespace, member).Body(body).Execute()

Update the role bind of the member.

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
	namespace := "namespace_example" // string | namespace
	member := "member_example" // string | namespace member's username
	body := *openapiclient.NewV1alpha2Member("RoleRef_example", "Username_example") // V1alpha2Member | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamespaceMemberAPI.UpdateNamespaceMember(context.Background(), namespace, member).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespaceMemberAPI.UpdateNamespaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamespaceMember`: V1alpha2Member
	fmt.Fprintf(os.Stdout, "Response from `NamespaceMemberAPI.UpdateNamespaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace | 
**member** | **string** | namespace member&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamespaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1alpha2Member**](V1alpha2Member.md) |  | 

### Return type

[**V1alpha2Member**](V1alpha2Member.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

