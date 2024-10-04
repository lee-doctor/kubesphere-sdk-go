# V1StatefulSetPersistentVolumeClaimRetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WhenDeleted** | Pointer to **string** | WhenDeleted specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is deleted. The default policy of &#x60;Retain&#x60; causes PVCs to not be affected by StatefulSet deletion. The &#x60;Delete&#x60; policy causes those PVCs to be deleted. | [optional] 
**WhenScaled** | Pointer to **string** | WhenScaled specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is scaled down. The default policy of &#x60;Retain&#x60; causes PVCs to not be affected by a scaledown. The &#x60;Delete&#x60; policy causes the associated PVCs for any excess pods above the replica count to be deleted. | [optional] 

## Methods

### NewV1StatefulSetPersistentVolumeClaimRetentionPolicy

`func NewV1StatefulSetPersistentVolumeClaimRetentionPolicy() *V1StatefulSetPersistentVolumeClaimRetentionPolicy`

NewV1StatefulSetPersistentVolumeClaimRetentionPolicy instantiates a new V1StatefulSetPersistentVolumeClaimRetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StatefulSetPersistentVolumeClaimRetentionPolicyWithDefaults

`func NewV1StatefulSetPersistentVolumeClaimRetentionPolicyWithDefaults() *V1StatefulSetPersistentVolumeClaimRetentionPolicy`

NewV1StatefulSetPersistentVolumeClaimRetentionPolicyWithDefaults instantiates a new V1StatefulSetPersistentVolumeClaimRetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhenDeleted

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenDeleted() string`

GetWhenDeleted returns the WhenDeleted field if non-nil, zero value otherwise.

### GetWhenDeletedOk

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenDeletedOk() (*string, bool)`

GetWhenDeletedOk returns a tuple with the WhenDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenDeleted

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) SetWhenDeleted(v string)`

SetWhenDeleted sets WhenDeleted field to given value.

### HasWhenDeleted

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) HasWhenDeleted() bool`

HasWhenDeleted returns a boolean if a field has been set.

### GetWhenScaled

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenScaled() string`

GetWhenScaled returns the WhenScaled field if non-nil, zero value otherwise.

### GetWhenScaledOk

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenScaledOk() (*string, bool)`

GetWhenScaledOk returns a tuple with the WhenScaled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenScaled

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) SetWhenScaled(v string)`

SetWhenScaled sets WhenScaled field to given value.

### HasWhenScaled

`func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) HasWhenScaled() bool`

HasWhenScaled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


