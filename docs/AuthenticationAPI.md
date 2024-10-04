# \AuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Logout**](AuthenticationAPI.md#Logout) | **Get** /oauth/logout | Logout
[**OauthCallback**](AuthenticationAPI.md#OauthCallback) | **Get** /oauth/callback/{callback} | OAuth2 callback
[**OpenidAuthorizeGet**](AuthenticationAPI.md#OpenidAuthorizeGet) | **Get** /oauth/authorize | Authorization endpoint
[**OpenidAuthorizePost**](AuthenticationAPI.md#OpenidAuthorizePost) | **Post** /oauth/authorize | Authorization endpoint
[**OpenidConfiguration**](AuthenticationAPI.md#OpenidConfiguration) | **Get** /.well-known/openid-configuration | OpenID provider configuration information
[**OpenidKeys**](AuthenticationAPI.md#OpenidKeys) | **Get** /oauth/keys | JSON Web Key Set
[**OpenidToken**](AuthenticationAPI.md#OpenidToken) | **Post** /oauth/token | Token endpoint
[**OpenidUserinfo**](AuthenticationAPI.md#OpenidUserinfo) | **Get** /oauth/userinfo | User info endpoint
[**TokenReview**](AuthenticationAPI.md#TokenReview) | **Post** /oauth/authenticate | Token review



## Logout

> string Logout(ctx).IdTokenHint(idTokenHint).PostLogoutRedirectUri(postLogoutRedirectUri).State(state).Execute()

Logout



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
	idTokenHint := "idTokenHint_example" // string | ID Token previously issued by the OP to the RP passed to the Logout Endpoint as a hint about the End-User's current authenticated session with the Client. This is used as an indication of the identity of the End-User that the RP is requesting be logged out by the OP. (optional)
	postLogoutRedirectUri := "postLogoutRedirectUri_example" // string | URL to which the RP is requesting that the End-User's User Agent be redirected after a logout has been performed.  (optional)
	state := "state_example" // string | Opaque value used by the RP to maintain state between the logout request and the callback to the endpoint specified by the post_logout_redirect_uri parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Logout(context.Background()).IdTokenHint(idTokenHint).PostLogoutRedirectUri(postLogoutRedirectUri).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Logout`: string
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Logout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idTokenHint** | **string** | ID Token previously issued by the OP to the RP passed to the Logout Endpoint as a hint about the End-User&#39;s current authenticated session with the Client. This is used as an indication of the identity of the End-User that the RP is requesting be logged out by the OP. | 
 **postLogoutRedirectUri** | **string** | URL to which the RP is requesting that the End-User&#39;s User Agent be redirected after a logout has been performed.  | 
 **state** | **string** | Opaque value used by the RP to maintain state between the logout request and the callback to the endpoint specified by the post_logout_redirect_uri parameter. | 

### Return type

**string**

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthCallback

> OauthToken OauthCallback(ctx, callback).AccessToken(accessToken).TokenType(tokenType).State(state).ExpiresIn(expiresIn).Scope(scope).Execute()

OAuth2 callback

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
	callback := "callback_example" // string | The identity provider name.
	accessToken := "accessToken_example" // string | The access token issued by the authorization server.
	tokenType := "tokenType_example" // string | The type of the token issued as described in [RFC6479] Section 7.1. Value is case insensitive.
	state := "state_example" // string | if the \"state\" parameter was present in the client authorization request.The exact value received from the client.
	expiresIn := "expiresIn_example" // string | The lifetime in seconds of the access token.  For example, the value \"3600\" denotes that the access token will expire in one hour from the time the response was generated.If omitted, the authorization server SHOULD provide the expiration time via other means or document the default value. (optional)
	scope := "scope_example" // string | if identical to the scope requested by the client;otherwise, REQUIRED.  The scope of the access token as described by [RFC6479] Section 3.3. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.OauthCallback(context.Background(), callback).AccessToken(accessToken).TokenType(tokenType).State(state).ExpiresIn(expiresIn).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OauthCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthCallback`: OauthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.OauthCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callback** | **string** | The identity provider name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthCallbackRequest struct via the builder pattern


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

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenidAuthorizeGet

> OpenidAuthorizeGet(ctx).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Scope(scope).State(state).Execute()

Authorization endpoint



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
	responseType := "responseType_example" // string | The value MUST be one of \"code\" for requesting an authorization code as described by [RFC6749] Section 4.1.1, \"token\" for requesting an access token (implicit grant) as described by [RFC6749] Section 4.2.2.
	clientId := "clientId_example" // string | OAuth 2.0 Client Identifier valid at the Authorization Server.
	redirectUri := "redirectUri_example" // string | Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider.
	scope := "scope_example" // string | OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. (optional)
	state := "state_example" // string | Opaque value used to maintain state between the request and the callback. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.OpenidAuthorizeGet(context.Background()).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Scope(scope).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OpenidAuthorizeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenidAuthorizeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **string** | The value MUST be one of \&quot;code\&quot; for requesting an authorization code as described by [RFC6749] Section 4.1.1, \&quot;token\&quot; for requesting an access token (implicit grant) as described by [RFC6749] Section 4.2.2. | 
 **clientId** | **string** | OAuth 2.0 Client Identifier valid at the Authorization Server. | 
 **redirectUri** | **string** | Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider. | 
 **scope** | **string** | OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. | 
 **state** | **string** | Opaque value used to maintain state between the request and the callback. | 

### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenidAuthorizePost

> OpenidAuthorizePost(ctx).ClientId(clientId).RedirectUri(redirectUri).ResponseType(responseType).Scope(scope).State(state).Execute()

Authorization endpoint



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
	clientId := "clientId_example" // string | OAuth 2.0 Client Identifier valid at the Authorization Server.
	redirectUri := "redirectUri_example" // string | Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider.
	responseType := "responseType_example" // string | The value MUST be one of \\\"code\\\" for requesting an authorization code as described by [RFC6749] Section 4.1.1, \\\"token\\\" for requesting an access token (implicit grant) as described by [RFC6749] Section 4.2.2. (optional)
	scope := "scope_example" // string | OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. (optional)
	state := "state_example" // string | Opaque value used to maintain state between the request and the callback. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.OpenidAuthorizePost(context.Background()).ClientId(clientId).RedirectUri(redirectUri).ResponseType(responseType).Scope(scope).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OpenidAuthorizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenidAuthorizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | OAuth 2.0 Client Identifier valid at the Authorization Server. | 
 **redirectUri** | **string** | Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider. | 
 **responseType** | **string** | The value MUST be one of \\\&quot;code\\\&quot; for requesting an authorization code as described by [RFC6749] Section 4.1.1, \\\&quot;token\\\&quot; for requesting an access token (implicit grant) as described by [RFC6749] Section 4.2.2. | 
 **scope** | **string** | OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. | 
 **state** | **string** | Opaque value used to maintain state between the request and the callback. | 

### Return type

 (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenidConfiguration

> OauthProviderMetadata OpenidConfiguration(ctx).Execute()

OpenID provider configuration information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.OpenidConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OpenidConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenidConfiguration`: OauthProviderMetadata
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.OpenidConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenidConfigurationRequest struct via the builder pattern


### Return type

[**OauthProviderMetadata**](OauthProviderMetadata.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenidKeys

> JoseJSONWebKeySet OpenidKeys(ctx).Execute()

JSON Web Key Set



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.OpenidKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OpenidKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenidKeys`: JoseJSONWebKeySet
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.OpenidKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenidKeysRequest struct via the builder pattern


### Return type

[**JoseJSONWebKeySet**](JoseJSONWebKeySet.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenidToken

> OauthToken OpenidToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Username(username).Password(password).Code(code).Execute()

Token endpoint



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
	grantType := "grantType_example" // string | OAuth defines four grant types: authorization code, implicit, resource owner password credentials, and client credentials.
	clientId := "clientId_example" // string | Valid client credential.
	clientSecret := "clientSecret_example" // string | Valid client credential.
	username := "username_example" // string | The resource owner username. (optional)
	password := "password_example" // string | The resource owner password. (optional)
	code := "code_example" // string | Valid authorization code. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.OpenidToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Username(username).Password(password).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OpenidToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenidToken`: OauthToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.OpenidToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenidTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | OAuth defines four grant types: authorization code, implicit, resource owner password credentials, and client credentials. | 
 **clientId** | **string** | Valid client credential. | 
 **clientSecret** | **string** | Valid client credential. | 
 **username** | **string** | The resource owner username. | 
 **password** | **string** | The resource owner password. | 
 **code** | **string** | Valid authorization code. | 

### Return type

[**OauthToken**](OauthToken.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenidUserinfo

> TokenClaims OpenidUserinfo(ctx).Execute()

User info endpoint



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.OpenidUserinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.OpenidUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenidUserinfo`: TokenClaims
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.OpenidUserinfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenidUserinfoRequest struct via the builder pattern


### Return type

[**TokenClaims**](TokenClaims.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenReview

> OauthTokenReview TokenReview(ctx).Body(body).Execute()

Token review



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

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

