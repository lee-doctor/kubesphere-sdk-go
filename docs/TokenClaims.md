# TokenClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Exp** | Pointer to **string** |  | [optional] 
**Extra** | Pointer to **map[string][]string** |  | [optional] 
**Iat** | Pointer to **string** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**Jti** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Nbf** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**PreferredUsername** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenClaims

`func NewTokenClaims() *TokenClaims`

NewTokenClaims instantiates a new TokenClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenClaimsWithDefaults

`func NewTokenClaimsWithDefaults() *TokenClaims`

NewTokenClaimsWithDefaults instantiates a new TokenClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *TokenClaims) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *TokenClaims) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *TokenClaims) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *TokenClaims) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetEmail

`func (o *TokenClaims) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TokenClaims) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TokenClaims) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TokenClaims) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExp

`func (o *TokenClaims) GetExp() string`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *TokenClaims) GetExpOk() (*string, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *TokenClaims) SetExp(v string)`

SetExp sets Exp field to given value.

### HasExp

`func (o *TokenClaims) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetExtra

`func (o *TokenClaims) GetExtra() map[string][]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *TokenClaims) GetExtraOk() (*map[string][]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *TokenClaims) SetExtra(v map[string][]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *TokenClaims) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetIat

`func (o *TokenClaims) GetIat() string`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *TokenClaims) GetIatOk() (*string, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *TokenClaims) SetIat(v string)`

SetIat sets Iat field to given value.

### HasIat

`func (o *TokenClaims) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *TokenClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *TokenClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *TokenClaims) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *TokenClaims) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetJti

`func (o *TokenClaims) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *TokenClaims) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *TokenClaims) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *TokenClaims) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetLocale

`func (o *TokenClaims) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TokenClaims) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TokenClaims) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TokenClaims) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetNbf

`func (o *TokenClaims) GetNbf() string`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *TokenClaims) GetNbfOk() (*string, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *TokenClaims) SetNbf(v string)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *TokenClaims) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetNonce

`func (o *TokenClaims) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TokenClaims) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TokenClaims) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TokenClaims) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPreferredUsername

`func (o *TokenClaims) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *TokenClaims) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *TokenClaims) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *TokenClaims) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### GetScopes

`func (o *TokenClaims) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenClaims) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenClaims) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenClaims) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSub

`func (o *TokenClaims) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *TokenClaims) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *TokenClaims) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *TokenClaims) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenClaims) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenClaims) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenClaims) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenClaims) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUrl

`func (o *TokenClaims) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TokenClaims) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TokenClaims) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TokenClaims) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *TokenClaims) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenClaims) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenClaims) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TokenClaims) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


