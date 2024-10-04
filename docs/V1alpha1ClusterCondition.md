# V1alpha1ClusterCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **string** |  | [optional] 
**LastUpdateTime** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewV1alpha1ClusterCondition

`func NewV1alpha1ClusterCondition(status string, type_ string, ) *V1alpha1ClusterCondition`

NewV1alpha1ClusterCondition instantiates a new V1alpha1ClusterCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterConditionWithDefaults

`func NewV1alpha1ClusterConditionWithDefaults() *V1alpha1ClusterCondition`

NewV1alpha1ClusterConditionWithDefaults instantiates a new V1alpha1ClusterCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *V1alpha1ClusterCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1alpha1ClusterCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1alpha1ClusterCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1alpha1ClusterCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *V1alpha1ClusterCondition) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *V1alpha1ClusterCondition) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *V1alpha1ClusterCondition) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *V1alpha1ClusterCondition) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMessage

`func (o *V1alpha1ClusterCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1alpha1ClusterCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1alpha1ClusterCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1alpha1ClusterCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *V1alpha1ClusterCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1alpha1ClusterCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1alpha1ClusterCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1alpha1ClusterCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha1ClusterCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha1ClusterCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha1ClusterCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *V1alpha1ClusterCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1ClusterCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1ClusterCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


