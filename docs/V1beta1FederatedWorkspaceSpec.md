# V1beta1FederatedWorkspaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | Pointer to [**[]V1beta1GenericOverrideItem**](V1beta1GenericOverrideItem.md) |  | [optional] 
**Placement** | [**V1beta1GenericPlacementFields**](V1beta1GenericPlacementFields.md) |  | 
**Template** | [**V1beta1WorkspaceTemplate**](V1beta1WorkspaceTemplate.md) |  | 

## Methods

### NewV1beta1FederatedWorkspaceSpec

`func NewV1beta1FederatedWorkspaceSpec(placement V1beta1GenericPlacementFields, template V1beta1WorkspaceTemplate, ) *V1beta1FederatedWorkspaceSpec`

NewV1beta1FederatedWorkspaceSpec instantiates a new V1beta1FederatedWorkspaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1FederatedWorkspaceSpecWithDefaults

`func NewV1beta1FederatedWorkspaceSpecWithDefaults() *V1beta1FederatedWorkspaceSpec`

NewV1beta1FederatedWorkspaceSpecWithDefaults instantiates a new V1beta1FederatedWorkspaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *V1beta1FederatedWorkspaceSpec) GetOverrides() []V1beta1GenericOverrideItem`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *V1beta1FederatedWorkspaceSpec) GetOverridesOk() (*[]V1beta1GenericOverrideItem, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *V1beta1FederatedWorkspaceSpec) SetOverrides(v []V1beta1GenericOverrideItem)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *V1beta1FederatedWorkspaceSpec) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetPlacement

`func (o *V1beta1FederatedWorkspaceSpec) GetPlacement() V1beta1GenericPlacementFields`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *V1beta1FederatedWorkspaceSpec) GetPlacementOk() (*V1beta1GenericPlacementFields, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *V1beta1FederatedWorkspaceSpec) SetPlacement(v V1beta1GenericPlacementFields)`

SetPlacement sets Placement field to given value.


### GetTemplate

`func (o *V1beta1FederatedWorkspaceSpec) GetTemplate() V1beta1WorkspaceTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1beta1FederatedWorkspaceSpec) GetTemplateOk() (*V1beta1WorkspaceTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1beta1FederatedWorkspaceSpec) SetTemplate(v V1beta1WorkspaceTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


