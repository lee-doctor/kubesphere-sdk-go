# V1beta1UserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLoginTime** | Pointer to **string** |  | [optional] 
**LastTransitionTime** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewV1beta1UserStatus

`func NewV1beta1UserStatus() *V1beta1UserStatus`

NewV1beta1UserStatus instantiates a new V1beta1UserStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1UserStatusWithDefaults

`func NewV1beta1UserStatusWithDefaults() *V1beta1UserStatus`

NewV1beta1UserStatusWithDefaults instantiates a new V1beta1UserStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLoginTime

`func (o *V1beta1UserStatus) GetLastLoginTime() string`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *V1beta1UserStatus) GetLastLoginTimeOk() (*string, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *V1beta1UserStatus) SetLastLoginTime(v string)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *V1beta1UserStatus) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *V1beta1UserStatus) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1beta1UserStatus) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1beta1UserStatus) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1beta1UserStatus) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetReason

`func (o *V1beta1UserStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1beta1UserStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1beta1UserStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1beta1UserStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *V1beta1UserStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1beta1UserStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1beta1UserStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V1beta1UserStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


