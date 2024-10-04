# V1beta1SubjectAccessReviewSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extra** | Pointer to **map[string][]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**NonResourceAttributes** | Pointer to [**V1beta1NonResourceAttributes**](V1beta1NonResourceAttributes.md) |  | [optional] 
**ResourceAttributes** | Pointer to [**V1beta1ResourceAttributes**](V1beta1ResourceAttributes.md) |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewV1beta1SubjectAccessReviewSpec

`func NewV1beta1SubjectAccessReviewSpec() *V1beta1SubjectAccessReviewSpec`

NewV1beta1SubjectAccessReviewSpec instantiates a new V1beta1SubjectAccessReviewSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1SubjectAccessReviewSpecWithDefaults

`func NewV1beta1SubjectAccessReviewSpecWithDefaults() *V1beta1SubjectAccessReviewSpec`

NewV1beta1SubjectAccessReviewSpecWithDefaults instantiates a new V1beta1SubjectAccessReviewSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtra

`func (o *V1beta1SubjectAccessReviewSpec) GetExtra() map[string][]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V1beta1SubjectAccessReviewSpec) GetExtraOk() (*map[string][]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V1beta1SubjectAccessReviewSpec) SetExtra(v map[string][]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V1beta1SubjectAccessReviewSpec) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetGroups

`func (o *V1beta1SubjectAccessReviewSpec) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V1beta1SubjectAccessReviewSpec) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V1beta1SubjectAccessReviewSpec) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V1beta1SubjectAccessReviewSpec) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetNonResourceAttributes

`func (o *V1beta1SubjectAccessReviewSpec) GetNonResourceAttributes() V1beta1NonResourceAttributes`

GetNonResourceAttributes returns the NonResourceAttributes field if non-nil, zero value otherwise.

### GetNonResourceAttributesOk

`func (o *V1beta1SubjectAccessReviewSpec) GetNonResourceAttributesOk() (*V1beta1NonResourceAttributes, bool)`

GetNonResourceAttributesOk returns a tuple with the NonResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceAttributes

`func (o *V1beta1SubjectAccessReviewSpec) SetNonResourceAttributes(v V1beta1NonResourceAttributes)`

SetNonResourceAttributes sets NonResourceAttributes field to given value.

### HasNonResourceAttributes

`func (o *V1beta1SubjectAccessReviewSpec) HasNonResourceAttributes() bool`

HasNonResourceAttributes returns a boolean if a field has been set.

### GetResourceAttributes

`func (o *V1beta1SubjectAccessReviewSpec) GetResourceAttributes() V1beta1ResourceAttributes`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *V1beta1SubjectAccessReviewSpec) GetResourceAttributesOk() (*V1beta1ResourceAttributes, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *V1beta1SubjectAccessReviewSpec) SetResourceAttributes(v V1beta1ResourceAttributes)`

SetResourceAttributes sets ResourceAttributes field to given value.

### HasResourceAttributes

`func (o *V1beta1SubjectAccessReviewSpec) HasResourceAttributes() bool`

HasResourceAttributes returns a boolean if a field has been set.

### GetUid

`func (o *V1beta1SubjectAccessReviewSpec) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *V1beta1SubjectAccessReviewSpec) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *V1beta1SubjectAccessReviewSpec) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *V1beta1SubjectAccessReviewSpec) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUser

`func (o *V1beta1SubjectAccessReviewSpec) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V1beta1SubjectAccessReviewSpec) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V1beta1SubjectAccessReviewSpec) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V1beta1SubjectAccessReviewSpec) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


