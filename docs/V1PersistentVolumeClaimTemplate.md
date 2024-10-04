# V1PersistentVolumeClaimTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | [**V1PersistentVolumeClaimSpec**](V1PersistentVolumeClaimSpec.md) |  | 

## Methods

### NewV1PersistentVolumeClaimTemplate

`func NewV1PersistentVolumeClaimTemplate(spec V1PersistentVolumeClaimSpec, ) *V1PersistentVolumeClaimTemplate`

NewV1PersistentVolumeClaimTemplate instantiates a new V1PersistentVolumeClaimTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PersistentVolumeClaimTemplateWithDefaults

`func NewV1PersistentVolumeClaimTemplateWithDefaults() *V1PersistentVolumeClaimTemplate`

NewV1PersistentVolumeClaimTemplateWithDefaults instantiates a new V1PersistentVolumeClaimTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1PersistentVolumeClaimTemplate) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1PersistentVolumeClaimTemplate) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1PersistentVolumeClaimTemplate) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1PersistentVolumeClaimTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1PersistentVolumeClaimTemplate) GetSpec() V1PersistentVolumeClaimSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1PersistentVolumeClaimTemplate) GetSpecOk() (*V1PersistentVolumeClaimSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1PersistentVolumeClaimTemplate) SetSpec(v V1PersistentVolumeClaimSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


