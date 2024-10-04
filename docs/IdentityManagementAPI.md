# \IdentityManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](IdentityManagementAPI.md#CreateUser) | **Post** /kapis/iam.kubesphere.io/v1beta1/users | Create user
[**DeleteUser**](IdentityManagementAPI.md#DeleteUser) | **Delete** /kapis/iam.kubesphere.io/v1beta1/users/{user} | Delete user
[**DescribeUser**](IdentityManagementAPI.md#DescribeUser) | **Get** /kapis/iam.kubesphere.io/v1beta1/users/{user} | Get user
[**ListUserLoginRecords**](IdentityManagementAPI.md#ListUserLoginRecords) | **Get** /kapis/iam.kubesphere.io/v1beta1/users/{user}/loginrecords | List login records
[**ListUsers**](IdentityManagementAPI.md#ListUsers) | **Get** /kapis/iam.kubesphere.io/v1beta1/users | List users
[**ModifyPassword**](IdentityManagementAPI.md#ModifyPassword) | **Put** /kapis/iam.kubesphere.io/v1beta1/users/{user}/password | Reset password
[**UpdateUser**](IdentityManagementAPI.md#UpdateUser) | **Put** /kapis/iam.kubesphere.io/v1beta1/users/{user} | Update user



## CreateUser

> V1beta1User CreateUser(ctx).Body(body).Execute()

Create user

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
	body := *openapiclient.NewV1beta1User(*openapiclient.NewV1beta1UserSpec("Email_example")) // V1beta1User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityManagementAPI.CreateUser(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityManagementAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: V1beta1User
	fmt.Fprintf(os.Stdout, "Response from `IdentityManagementAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1User**](V1beta1User.md) |  | 

### Return type

[**V1beta1User**](V1beta1User.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> ErrorsError DeleteUser(ctx, user).Execute()

Delete user

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
	user := "user_example" // string | username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityManagementAPI.DeleteUser(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityManagementAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `IdentityManagementAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUser

> V1beta1User DescribeUser(ctx, user).Execute()

Get user



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
	user := "user_example" // string | username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityManagementAPI.DescribeUser(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityManagementAPI.DescribeUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeUser`: V1beta1User
	fmt.Fprintf(os.Stdout, "Response from `IdentityManagementAPI.DescribeUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1beta1User**](V1beta1User.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserLoginRecords

> ApiListResult ListUserLoginRecords(ctx, user).Execute()

List login records

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
	user := "user_example" // string | username of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityManagementAPI.ListUserLoginRecords(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityManagementAPI.ListUserLoginRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserLoginRecords`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `IdentityManagementAPI.ListUserLoginRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserLoginRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ApiListResult ListUsers(ctx).Globalrole(globalrole).Execute()

List users

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
	globalrole := "globalrole_example" // string | specific golalrole name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityManagementAPI.ListUsers(context.Background()).Globalrole(globalrole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityManagementAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: ApiListResult
	fmt.Fprintf(os.Stdout, "Response from `IdentityManagementAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalrole** | **string** | specific golalrole name | 

### Return type

[**ApiListResult**](ApiListResult.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPassword

> ErrorsError ModifyPassword(ctx, user).Body(body).Execute()

Reset password

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
	user := "user_example" // string | username
	body := *openapiclient.NewV1beta1PasswordReset("CurrentPassword_example", "Password_example") // V1beta1PasswordReset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityManagementAPI.ModifyPassword(context.Background(), user).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityManagementAPI.ModifyPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyPassword`: ErrorsError
	fmt.Fprintf(os.Stdout, "Response from `IdentityManagementAPI.ModifyPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1beta1PasswordReset**](V1beta1PasswordReset.md) |  | 

### Return type

[**ErrorsError**](ErrorsError.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> V1beta1User UpdateUser(ctx, user).Body(body).Execute()

Update user

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
	user := "user_example" // string | username
	body := *openapiclient.NewV1beta1User(*openapiclient.NewV1beta1UserSpec("Email_example")) // V1beta1User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityManagementAPI.UpdateUser(context.Background(), user).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityManagementAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: V1beta1User
	fmt.Fprintf(os.Stdout, "Response from `IdentityManagementAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1beta1User**](V1beta1User.md) |  | 

### Return type

[**V1beta1User**](V1beta1User.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

