# V1beta1UserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Groups** | Pointer to **[]string** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewV1beta1UserSpec

`func NewV1beta1UserSpec(email string, ) *V1beta1UserSpec`

NewV1beta1UserSpec instantiates a new V1beta1UserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1UserSpecWithDefaults

`func NewV1beta1UserSpecWithDefaults() *V1beta1UserSpec`

NewV1beta1UserSpecWithDefaults instantiates a new V1beta1UserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1beta1UserSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1beta1UserSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1beta1UserSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1beta1UserSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *V1beta1UserSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *V1beta1UserSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *V1beta1UserSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *V1beta1UserSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *V1beta1UserSpec) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1beta1UserSpec) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1beta1UserSpec) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGroups

`func (o *V1beta1UserSpec) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V1beta1UserSpec) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V1beta1UserSpec) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V1beta1UserSpec) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLang

`func (o *V1beta1UserSpec) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *V1beta1UserSpec) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *V1beta1UserSpec) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *V1beta1UserSpec) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPassword

`func (o *V1beta1UserSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1beta1UserSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1beta1UserSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1beta1UserSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


