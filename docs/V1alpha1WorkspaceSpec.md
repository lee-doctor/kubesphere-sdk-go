# V1alpha1WorkspaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manager** | Pointer to **string** |  | [optional] 
**NetworkIsolation** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1alpha1WorkspaceSpec

`func NewV1alpha1WorkspaceSpec() *V1alpha1WorkspaceSpec`

NewV1alpha1WorkspaceSpec instantiates a new V1alpha1WorkspaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1WorkspaceSpecWithDefaults

`func NewV1alpha1WorkspaceSpecWithDefaults() *V1alpha1WorkspaceSpec`

NewV1alpha1WorkspaceSpecWithDefaults instantiates a new V1alpha1WorkspaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManager

`func (o *V1alpha1WorkspaceSpec) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *V1alpha1WorkspaceSpec) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *V1alpha1WorkspaceSpec) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *V1alpha1WorkspaceSpec) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNetworkIsolation

`func (o *V1alpha1WorkspaceSpec) GetNetworkIsolation() bool`

GetNetworkIsolation returns the NetworkIsolation field if non-nil, zero value otherwise.

### GetNetworkIsolationOk

`func (o *V1alpha1WorkspaceSpec) GetNetworkIsolationOk() (*bool, bool)`

GetNetworkIsolationOk returns a tuple with the NetworkIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsolation

`func (o *V1alpha1WorkspaceSpec) SetNetworkIsolation(v bool)`

SetNetworkIsolation sets NetworkIsolation field to given value.

### HasNetworkIsolation

`func (o *V1alpha1WorkspaceSpec) HasNetworkIsolation() bool`

HasNetworkIsolation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


