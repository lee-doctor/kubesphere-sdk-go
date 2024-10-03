# \AuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthenticationAPI.md#Authorize) | **Get** /oauth/authorize | All requests for OAuth tokens involve a request to /oauth/authorize.
[**Login**](AuthenticationAPI.md#Login) | **Post** /kapis/iam.kubesphere.io/v1alpha2/login | KubeSphere APIs support token-based authentication via the Authtoken request header. The POST Login API is used to retrieve the authentication token. After the authentication token is obtained, it must be inserted into the Authtoken header for all requests.
[**OauthCallBack**](AuthenticationAPI.md#OauthCallBack) | **Get** /oauth/callback/{callback} | OAuth callback API, the path param callback is config by identity provider
[**Token**](AuthenticationAPI.md#Token) | **Post** /oauth/token | The resource owner password credentials grant type is suitable in cases where the resource owner has a trust relationship with the client, such as the device operating system or a highly privileged application.
[**TokenReview**](AuthenticationAPI.md#TokenReview) | **Post** /oauth/authenticate | TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.



## Authorize

> Authorize(ctx).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Execute()

All requests for OAuth tokens involve a request to /oauth/authorize.

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
	responseType := "responseType_example" // string | The value MUST be one of \"code\" for requesting an authorization code as described by [RFC6749] Section 4.1.1, \"token\" for requesting an access token (implicit grant) as described by [RFC6749] Section 4.2.2.
	clientId := "clientId_example" // string | The client identifier issued to the client during the registration process described by [RFC6749] Section 2.2.
	redirectUri := "redirectUri_example" // string | After completing its interaction with the resource owner, the authorization server directs the resource owner's user-agent back to the client.The redirection endpoint URI MUST be an absolute URI as defined by [RFC3986] Section 4.3. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.Authorize(context.Background()).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Authorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **string** | The value MUST be one of \&quot;code\&quot; for requesting an authorization code as described by [RFC6749] Section 4.1.1, \&quot;token\&quot; for requesting an access token (implicit grant) as described by [RFC6749] Section 4.2.2. | 
 **clientId** | **string** | The client identifier issued to the client during the registration process described by [RFC6749] Section 2.2. | 
 **redirectUri** | **string** | After completing its interaction with the resource owner, the authorization server directs the resource owner&#39;s user-agent back to the client.The redirection endpoint URI MUST be an absolute URI as defined by [RFC3986] Section 4.3. | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> OauthToken Login(ctx).Body(body).Execute()

KubeSphere APIs support token-based authentication via the Authtoken request header. The POST Login API is used to retrieve the authentication token. After the authentication token is obtained, it must be inserted into the Authtoken header for all requests.

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
	body := *openapiclient.NewOauthLoginRequest("Password_example", "Username_example") // OauthLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Login(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: OauthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OauthLoginRequest**](OauthLoginRequest.md) |  | 

### Return type

[**OauthToken**](OauthToken.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthCallBack

> OauthToken OauthCallBack(ctx, callback).AccessToken(accessToken).TokenType(tokenType).State(state).ExpiresIn(expiresIn).Scope(scope).Execute()

OAuth callback API, the path param callback is config by identity provider

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
	callback := "callback_example" // string | callback name.
	accessToken := "accessToken_example" // string | The access token issued by the authorization server.
	tokenType := "tokenType_example" // string | The type of the token issued as described in [RFC6479] Section 7.1. Value is case insensitive.
	state := "state_example" // string | if the \"state\" parameter was present in the client authorization request.The exact value received from the client.
	expiresIn := "expiresIn_example" // string | The lifetime in seconds of the access token.  For example, the value \"3600\" denotes that the access token will expire in one hour from the time the response was generated.If omitted, the authorization server SHOULD provide the expiration time via other means or document the default value. (optional)
	scope := "scope_example" // string | if identical to the scope requested by the client;otherwise, REQUIRED.  The scope of the access token as described by [RFC6479] Section 3.3. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.OauthCallBack(context.Background(), callback).AccessToken(accessToken).TokenType(tokenType).State(state).ExpiresIn(expiresIn).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OauthCallBack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthCallBack`: OauthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.OauthCallBack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callback** | **string** | callback name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthCallBackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **string** | The access token issued by the authorization server. | 
 **tokenType** | **string** | The type of the token issued as described in [RFC6479] Section 7.1. Value is case insensitive. | 
 **state** | **string** | if the \&quot;state\&quot; parameter was present in the client authorization request.The exact value received from the client. | 
 **expiresIn** | **string** | The lifetime in seconds of the access token.  For example, the value \&quot;3600\&quot; denotes that the access token will expire in one hour from the time the response was generated.If omitted, the authorization server SHOULD provide the expiration time via other means or document the default value. | 
 **scope** | **string** | if identical to the scope requested by the client;otherwise, REQUIRED.  The scope of the access token as described by [RFC6479] Section 3.3. | 

### Return type

[**OauthToken**](OauthToken.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Token

> OauthToken Token(ctx).GrantType(grantType).Username(username).Password(password).Execute()

The resource owner password credentials grant type is suitable in cases where the resource owner has a trust relationship with the client, such as the device operating system or a highly privileged application.

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
	grantType := "grantType_example" // string | Value MUST be set to \\\"password\\\".
	username := "username_example" // string | The resource owner username.
	password := "password_example" // string | The resource owner password.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Token(context.Background()).GrantType(grantType).Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Token``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Token`: OauthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Token`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | Value MUST be set to \\\&quot;password\\\&quot;. | 
 **username** | **string** | The resource owner username. | 
 **password** | **string** | The resource owner password. | 

### Return type

[**OauthToken**](OauthToken.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenReview

> OauthTokenReview TokenReview(ctx).Body(body).Execute()

TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.

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
	body := *openapiclient.NewOauthTokenReview("ApiVersion_example", "Kind_example") // OauthTokenReview | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.TokenReview(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.TokenReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenReview`: OauthTokenReview
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.TokenReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OauthTokenReview**](OauthTokenReview.md) |  | 

### Return type

[**OauthTokenReview**](OauthTokenReview.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

