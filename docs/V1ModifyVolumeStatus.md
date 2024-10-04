# V1ModifyVolumeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | status is the status of the ControllerModifyVolume operation. It can be in any of following states:  - Pending    Pending indicates that the PersistentVolumeClaim cannot be modified due to unmet requirements, such as    the specified VolumeAttributesClass not existing.  - InProgress    InProgress indicates that the volume is being modified.  - Infeasible   Infeasible indicates that the request has been rejected as invalid by the CSI driver. To    resolve the error, a valid VolumeAttributesClass needs to be specified. Note: New statuses can be added in the future. Consumers should check for unknown statuses and fail appropriately. | 
**TargetVolumeAttributesClassName** | Pointer to **string** | targetVolumeAttributesClassName is the name of the VolumeAttributesClass the PVC currently being reconciled | [optional] 

## Methods

### NewV1ModifyVolumeStatus

`func NewV1ModifyVolumeStatus(status string, ) *V1ModifyVolumeStatus`

NewV1ModifyVolumeStatus instantiates a new V1ModifyVolumeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ModifyVolumeStatusWithDefaults

`func NewV1ModifyVolumeStatusWithDefaults() *V1ModifyVolumeStatus`

NewV1ModifyVolumeStatusWithDefaults instantiates a new V1ModifyVolumeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V1ModifyVolumeStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ModifyVolumeStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ModifyVolumeStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTargetVolumeAttributesClassName

`func (o *V1ModifyVolumeStatus) GetTargetVolumeAttributesClassName() string`

GetTargetVolumeAttributesClassName returns the TargetVolumeAttributesClassName field if non-nil, zero value otherwise.

### GetTargetVolumeAttributesClassNameOk

`func (o *V1ModifyVolumeStatus) GetTargetVolumeAttributesClassNameOk() (*string, bool)`

GetTargetVolumeAttributesClassNameOk returns a tuple with the TargetVolumeAttributesClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVolumeAttributesClassName

`func (o *V1ModifyVolumeStatus) SetTargetVolumeAttributesClassName(v string)`

SetTargetVolumeAttributesClassName sets TargetVolumeAttributesClassName field to given value.

### HasTargetVolumeAttributesClassName

`func (o *V1ModifyVolumeStatus) HasTargetVolumeAttributesClassName() bool`

HasTargetVolumeAttributesClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


