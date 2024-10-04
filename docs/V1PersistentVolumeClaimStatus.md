# V1PersistentVolumeClaimStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**AllocatedResourceStatuses** | Pointer to **map[string]string** | allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax. Valid values are either:  * Un-prefixed keys:   - storage - the capacity of the volume.  * Custom resources must use implementation-defined prefixed names such as \&quot;example.com/my-custom-resource\&quot; Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used.  ClaimResourceStatus can be in any of following states:  - ControllerResizeInProgress:   State set when resize controller starts resizing the volume in control-plane.  - ControllerResizeFailed:   State set when resize has failed in resize controller with a terminal error.  - NodeResizePending:   State set when resize controller has finished resizing the volume but further resizing of   volume is needed on the node.  - NodeResizeInProgress:   State set when kubelet starts resizing the volume.  - NodeResizeFailed:   State set when resizing has failed in kubelet with a terminal error. Transient errors don&#39;t set   NodeResizeFailed. For example: if expanding a PVC for more capacity - this field can be one of the following states:  - pvc.status.allocatedResourceStatus[&#39;storage&#39;] &#x3D; \&quot;ControllerResizeInProgress\&quot;      - pvc.status.allocatedResourceStatus[&#39;storage&#39;] &#x3D; \&quot;ControllerResizeFailed\&quot;      - pvc.status.allocatedResourceStatus[&#39;storage&#39;] &#x3D; \&quot;NodeResizePending\&quot;      - pvc.status.allocatedResourceStatus[&#39;storage&#39;] &#x3D; \&quot;NodeResizeInProgress\&quot;      - pvc.status.allocatedResourceStatus[&#39;storage&#39;] &#x3D; \&quot;NodeResizeFailed\&quot; When this field is not set, it means that no resize operation is in progress for the given PVC.  A controller that receives PVC update with previously unknown resourceName or ClaimResourceStatus should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC.  This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature. | [optional] 
**AllocatedResources** | Pointer to [**map[string]ResourceQuantity**](ResourceQuantity.md) | allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax. Valid values are either:  * Un-prefixed keys:   - storage - the capacity of the volume.  * Custom resources must use implementation-defined prefixed names such as \&quot;example.com/my-custom-resource\&quot; Apart from above values - keys that are unprefixed or have kubernetes.io prefix are considered reserved and hence may not be used.  Capacity reported here may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity.  A controller that receives PVC update with previously unknown resourceName should ignore the update for the purpose it was designed. For example - a controller that only is responsible for resizing capacity of the volume, should ignore PVC updates that change other valid resources associated with PVC.  This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature. | [optional] 
**Capacity** | Pointer to [**map[string]ResourceQuantity**](ResourceQuantity.md) | capacity represents the actual resources of the underlying volume. | [optional] 
**Conditions** | Pointer to [**[]V1PersistentVolumeClaimCondition**](V1PersistentVolumeClaimCondition.md) | conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#39;ResizeStarted&#39;. | [optional] 
**CurrentVolumeAttributesClassName** | Pointer to **string** | currentVolumeAttributesClassName is the current name of the VolumeAttributesClass the PVC is using. When unset, there is no VolumeAttributeClass applied to this PersistentVolumeClaim This is an alpha field and requires enabling VolumeAttributesClass feature. | [optional] 
**ModifyVolumeStatus** | Pointer to [**V1ModifyVolumeStatus**](V1ModifyVolumeStatus.md) |  | [optional] 
**Phase** | Pointer to **string** | phase represents the current phase of PersistentVolumeClaim. | [optional] 

## Methods

### NewV1PersistentVolumeClaimStatus

`func NewV1PersistentVolumeClaimStatus() *V1PersistentVolumeClaimStatus`

NewV1PersistentVolumeClaimStatus instantiates a new V1PersistentVolumeClaimStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PersistentVolumeClaimStatusWithDefaults

`func NewV1PersistentVolumeClaimStatusWithDefaults() *V1PersistentVolumeClaimStatus`

NewV1PersistentVolumeClaimStatusWithDefaults instantiates a new V1PersistentVolumeClaimStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *V1PersistentVolumeClaimStatus) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *V1PersistentVolumeClaimStatus) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *V1PersistentVolumeClaimStatus) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *V1PersistentVolumeClaimStatus) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetAllocatedResourceStatuses

`func (o *V1PersistentVolumeClaimStatus) GetAllocatedResourceStatuses() map[string]string`

GetAllocatedResourceStatuses returns the AllocatedResourceStatuses field if non-nil, zero value otherwise.

### GetAllocatedResourceStatusesOk

`func (o *V1PersistentVolumeClaimStatus) GetAllocatedResourceStatusesOk() (*map[string]string, bool)`

GetAllocatedResourceStatusesOk returns a tuple with the AllocatedResourceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResourceStatuses

`func (o *V1PersistentVolumeClaimStatus) SetAllocatedResourceStatuses(v map[string]string)`

SetAllocatedResourceStatuses sets AllocatedResourceStatuses field to given value.

### HasAllocatedResourceStatuses

`func (o *V1PersistentVolumeClaimStatus) HasAllocatedResourceStatuses() bool`

HasAllocatedResourceStatuses returns a boolean if a field has been set.

### GetAllocatedResources

`func (o *V1PersistentVolumeClaimStatus) GetAllocatedResources() map[string]ResourceQuantity`

GetAllocatedResources returns the AllocatedResources field if non-nil, zero value otherwise.

### GetAllocatedResourcesOk

`func (o *V1PersistentVolumeClaimStatus) GetAllocatedResourcesOk() (*map[string]ResourceQuantity, bool)`

GetAllocatedResourcesOk returns a tuple with the AllocatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResources

`func (o *V1PersistentVolumeClaimStatus) SetAllocatedResources(v map[string]ResourceQuantity)`

SetAllocatedResources sets AllocatedResources field to given value.

### HasAllocatedResources

`func (o *V1PersistentVolumeClaimStatus) HasAllocatedResources() bool`

HasAllocatedResources returns a boolean if a field has been set.

### GetCapacity

`func (o *V1PersistentVolumeClaimStatus) GetCapacity() map[string]ResourceQuantity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *V1PersistentVolumeClaimStatus) GetCapacityOk() (*map[string]ResourceQuantity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *V1PersistentVolumeClaimStatus) SetCapacity(v map[string]ResourceQuantity)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *V1PersistentVolumeClaimStatus) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConditions

`func (o *V1PersistentVolumeClaimStatus) GetConditions() []V1PersistentVolumeClaimCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1PersistentVolumeClaimStatus) GetConditionsOk() (*[]V1PersistentVolumeClaimCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1PersistentVolumeClaimStatus) SetConditions(v []V1PersistentVolumeClaimCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1PersistentVolumeClaimStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCurrentVolumeAttributesClassName

`func (o *V1PersistentVolumeClaimStatus) GetCurrentVolumeAttributesClassName() string`

GetCurrentVolumeAttributesClassName returns the CurrentVolumeAttributesClassName field if non-nil, zero value otherwise.

### GetCurrentVolumeAttributesClassNameOk

`func (o *V1PersistentVolumeClaimStatus) GetCurrentVolumeAttributesClassNameOk() (*string, bool)`

GetCurrentVolumeAttributesClassNameOk returns a tuple with the CurrentVolumeAttributesClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVolumeAttributesClassName

`func (o *V1PersistentVolumeClaimStatus) SetCurrentVolumeAttributesClassName(v string)`

SetCurrentVolumeAttributesClassName sets CurrentVolumeAttributesClassName field to given value.

### HasCurrentVolumeAttributesClassName

`func (o *V1PersistentVolumeClaimStatus) HasCurrentVolumeAttributesClassName() bool`

HasCurrentVolumeAttributesClassName returns a boolean if a field has been set.

### GetModifyVolumeStatus

`func (o *V1PersistentVolumeClaimStatus) GetModifyVolumeStatus() V1ModifyVolumeStatus`

GetModifyVolumeStatus returns the ModifyVolumeStatus field if non-nil, zero value otherwise.

### GetModifyVolumeStatusOk

`func (o *V1PersistentVolumeClaimStatus) GetModifyVolumeStatusOk() (*V1ModifyVolumeStatus, bool)`

GetModifyVolumeStatusOk returns a tuple with the ModifyVolumeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyVolumeStatus

`func (o *V1PersistentVolumeClaimStatus) SetModifyVolumeStatus(v V1ModifyVolumeStatus)`

SetModifyVolumeStatus sets ModifyVolumeStatus field to given value.

### HasModifyVolumeStatus

`func (o *V1PersistentVolumeClaimStatus) HasModifyVolumeStatus() bool`

HasModifyVolumeStatus returns a boolean if a field has been set.

### GetPhase

`func (o *V1PersistentVolumeClaimStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1PersistentVolumeClaimStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1PersistentVolumeClaimStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1PersistentVolumeClaimStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


