# V1alpha2UserSpec

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

### NewV1alpha2UserSpec

`func NewV1alpha2UserSpec(email string, ) *V1alpha2UserSpec`

NewV1alpha2UserSpec instantiates a new V1alpha2UserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2UserSpecWithDefaults

`func NewV1alpha2UserSpecWithDefaults() *V1alpha2UserSpec`

NewV1alpha2UserSpecWithDefaults instantiates a new V1alpha2UserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1alpha2UserSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1alpha2UserSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1alpha2UserSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1alpha2UserSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *V1alpha2UserSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *V1alpha2UserSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *V1alpha2UserSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *V1alpha2UserSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *V1alpha2UserSpec) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *V1alpha2UserSpec) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *V1alpha2UserSpec) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGroups

`func (o *V1alpha2UserSpec) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V1alpha2UserSpec) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V1alpha2UserSpec) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V1alpha2UserSpec) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLang

`func (o *V1alpha2UserSpec) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *V1alpha2UserSpec) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *V1alpha2UserSpec) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *V1alpha2UserSpec) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPassword

`func (o *V1alpha2UserSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1alpha2UserSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1alpha2UserSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1alpha2UserSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


