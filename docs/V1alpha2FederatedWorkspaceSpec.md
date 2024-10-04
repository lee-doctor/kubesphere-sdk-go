# V1alpha2FederatedWorkspaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | Pointer to [**[]V1alpha2GenericOverrideItem**](V1alpha2GenericOverrideItem.md) |  | [optional] 
**Placement** | [**V1alpha2GenericPlacementFields**](V1alpha2GenericPlacementFields.md) |  | 
**Template** | [**V1alpha2WorkspaceTemplateSpec**](V1alpha2WorkspaceTemplateSpec.md) |  | 

## Methods

### NewV1alpha2FederatedWorkspaceSpec

`func NewV1alpha2FederatedWorkspaceSpec(placement V1alpha2GenericPlacementFields, template V1alpha2WorkspaceTemplateSpec, ) *V1alpha2FederatedWorkspaceSpec`

NewV1alpha2FederatedWorkspaceSpec instantiates a new V1alpha2FederatedWorkspaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2FederatedWorkspaceSpecWithDefaults

`func NewV1alpha2FederatedWorkspaceSpecWithDefaults() *V1alpha2FederatedWorkspaceSpec`

NewV1alpha2FederatedWorkspaceSpecWithDefaults instantiates a new V1alpha2FederatedWorkspaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *V1alpha2FederatedWorkspaceSpec) GetOverrides() []V1alpha2GenericOverrideItem`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *V1alpha2FederatedWorkspaceSpec) GetOverridesOk() (*[]V1alpha2GenericOverrideItem, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *V1alpha2FederatedWorkspaceSpec) SetOverrides(v []V1alpha2GenericOverrideItem)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *V1alpha2FederatedWorkspaceSpec) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetPlacement

`func (o *V1alpha2FederatedWorkspaceSpec) GetPlacement() V1alpha2GenericPlacementFields`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *V1alpha2FederatedWorkspaceSpec) GetPlacementOk() (*V1alpha2GenericPlacementFields, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *V1alpha2FederatedWorkspaceSpec) SetPlacement(v V1alpha2GenericPlacementFields)`

SetPlacement sets Placement field to given value.


### GetTemplate

`func (o *V1alpha2FederatedWorkspaceSpec) GetTemplate() V1alpha2WorkspaceTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1alpha2FederatedWorkspaceSpec) GetTemplateOk() (*V1alpha2WorkspaceTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1alpha2FederatedWorkspaceSpec) SetTemplate(v V1alpha2WorkspaceTemplateSpec)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


