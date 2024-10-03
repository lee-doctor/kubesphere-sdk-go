# OauthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | **bool** | is authenticated | 
**User** | Pointer to **map[string]interface{}** | user info | [optional] 

## Methods

### NewOauthStatus

`func NewOauthStatus(authenticated bool, ) *OauthStatus`

NewOauthStatus instantiates a new OauthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthStatusWithDefaults

`func NewOauthStatusWithDefaults() *OauthStatus`

NewOauthStatusWithDefaults instantiates a new OauthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *OauthStatus) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *OauthStatus) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *OauthStatus) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.


### GetUser

`func (o *OauthStatus) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OauthStatus) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OauthStatus) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *OauthStatus) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


