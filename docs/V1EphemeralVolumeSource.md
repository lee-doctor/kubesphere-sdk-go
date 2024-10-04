# V1EphemeralVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeClaimTemplate** | Pointer to [**V1PersistentVolumeClaimTemplate**](V1PersistentVolumeClaimTemplate.md) |  | [optional] 

## Methods

### NewV1EphemeralVolumeSource

`func NewV1EphemeralVolumeSource() *V1EphemeralVolumeSource`

NewV1EphemeralVolumeSource instantiates a new V1EphemeralVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EphemeralVolumeSourceWithDefaults

`func NewV1EphemeralVolumeSourceWithDefaults() *V1EphemeralVolumeSource`

NewV1EphemeralVolumeSourceWithDefaults instantiates a new V1EphemeralVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeClaimTemplate

`func (o *V1EphemeralVolumeSource) GetVolumeClaimTemplate() V1PersistentVolumeClaimTemplate`

GetVolumeClaimTemplate returns the VolumeClaimTemplate field if non-nil, zero value otherwise.

### GetVolumeClaimTemplateOk

`func (o *V1EphemeralVolumeSource) GetVolumeClaimTemplateOk() (*V1PersistentVolumeClaimTemplate, bool)`

GetVolumeClaimTemplateOk returns a tuple with the VolumeClaimTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimTemplate

`func (o *V1EphemeralVolumeSource) SetVolumeClaimTemplate(v V1PersistentVolumeClaimTemplate)`

SetVolumeClaimTemplate sets VolumeClaimTemplate field to given value.

### HasVolumeClaimTemplate

`func (o *V1EphemeralVolumeSource) HasVolumeClaimTemplate() bool`

HasVolumeClaimTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


