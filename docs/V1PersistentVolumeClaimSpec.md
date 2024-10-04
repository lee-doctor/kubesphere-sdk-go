# V1PersistentVolumeClaimSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**DataSource** | Pointer to [**V1TypedLocalObjectReference**](V1TypedLocalObjectReference.md) |  | [optional] 
**DataSourceRef** | Pointer to [**V1TypedObjectReference**](V1TypedObjectReference.md) |  | [optional] 
**Resources** | Pointer to [**V1VolumeResourceRequirements**](V1VolumeResourceRequirements.md) |  | [optional] 
**Selector** | Pointer to [**V1LabelSelector**](V1LabelSelector.md) |  | [optional] 
**StorageClassName** | Pointer to **string** | storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | [optional] 
**VolumeAttributesClassName** | Pointer to **string** | volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it&#39;s not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass (Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled. | [optional] 
**VolumeMode** | Pointer to **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. | [optional] 
**VolumeName** | Pointer to **string** | volumeName is the binding reference to the PersistentVolume backing this claim. | [optional] 

## Methods

### NewV1PersistentVolumeClaimSpec

`func NewV1PersistentVolumeClaimSpec() *V1PersistentVolumeClaimSpec`

NewV1PersistentVolumeClaimSpec instantiates a new V1PersistentVolumeClaimSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PersistentVolumeClaimSpecWithDefaults

`func NewV1PersistentVolumeClaimSpecWithDefaults() *V1PersistentVolumeClaimSpec`

NewV1PersistentVolumeClaimSpecWithDefaults instantiates a new V1PersistentVolumeClaimSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *V1PersistentVolumeClaimSpec) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *V1PersistentVolumeClaimSpec) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *V1PersistentVolumeClaimSpec) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *V1PersistentVolumeClaimSpec) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetDataSource

`func (o *V1PersistentVolumeClaimSpec) GetDataSource() V1TypedLocalObjectReference`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *V1PersistentVolumeClaimSpec) GetDataSourceOk() (*V1TypedLocalObjectReference, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *V1PersistentVolumeClaimSpec) SetDataSource(v V1TypedLocalObjectReference)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *V1PersistentVolumeClaimSpec) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceRef

`func (o *V1PersistentVolumeClaimSpec) GetDataSourceRef() V1TypedObjectReference`

GetDataSourceRef returns the DataSourceRef field if non-nil, zero value otherwise.

### GetDataSourceRefOk

`func (o *V1PersistentVolumeClaimSpec) GetDataSourceRefOk() (*V1TypedObjectReference, bool)`

GetDataSourceRefOk returns a tuple with the DataSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRef

`func (o *V1PersistentVolumeClaimSpec) SetDataSourceRef(v V1TypedObjectReference)`

SetDataSourceRef sets DataSourceRef field to given value.

### HasDataSourceRef

`func (o *V1PersistentVolumeClaimSpec) HasDataSourceRef() bool`

HasDataSourceRef returns a boolean if a field has been set.

### GetResources

`func (o *V1PersistentVolumeClaimSpec) GetResources() V1VolumeResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1PersistentVolumeClaimSpec) GetResourcesOk() (*V1VolumeResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1PersistentVolumeClaimSpec) SetResources(v V1VolumeResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1PersistentVolumeClaimSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSelector

`func (o *V1PersistentVolumeClaimSpec) GetSelector() V1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1PersistentVolumeClaimSpec) GetSelectorOk() (*V1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1PersistentVolumeClaimSpec) SetSelector(v V1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1PersistentVolumeClaimSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetStorageClassName

`func (o *V1PersistentVolumeClaimSpec) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *V1PersistentVolumeClaimSpec) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *V1PersistentVolumeClaimSpec) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *V1PersistentVolumeClaimSpec) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetVolumeAttributesClassName

`func (o *V1PersistentVolumeClaimSpec) GetVolumeAttributesClassName() string`

GetVolumeAttributesClassName returns the VolumeAttributesClassName field if non-nil, zero value otherwise.

### GetVolumeAttributesClassNameOk

`func (o *V1PersistentVolumeClaimSpec) GetVolumeAttributesClassNameOk() (*string, bool)`

GetVolumeAttributesClassNameOk returns a tuple with the VolumeAttributesClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttributesClassName

`func (o *V1PersistentVolumeClaimSpec) SetVolumeAttributesClassName(v string)`

SetVolumeAttributesClassName sets VolumeAttributesClassName field to given value.

### HasVolumeAttributesClassName

`func (o *V1PersistentVolumeClaimSpec) HasVolumeAttributesClassName() bool`

HasVolumeAttributesClassName returns a boolean if a field has been set.

### GetVolumeMode

`func (o *V1PersistentVolumeClaimSpec) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *V1PersistentVolumeClaimSpec) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *V1PersistentVolumeClaimSpec) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *V1PersistentVolumeClaimSpec) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVolumeName

`func (o *V1PersistentVolumeClaimSpec) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *V1PersistentVolumeClaimSpec) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *V1PersistentVolumeClaimSpec) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *V1PersistentVolumeClaimSpec) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


