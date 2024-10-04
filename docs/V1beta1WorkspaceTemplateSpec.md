# V1beta1WorkspaceTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placement** | [**V1beta1GenericPlacement**](V1beta1GenericPlacement.md) |  | 
**Template** | [**V1beta1Template**](V1beta1Template.md) |  | 

## Methods

### NewV1beta1WorkspaceTemplateSpec

`func NewV1beta1WorkspaceTemplateSpec(placement V1beta1GenericPlacement, template V1beta1Template, ) *V1beta1WorkspaceTemplateSpec`

NewV1beta1WorkspaceTemplateSpec instantiates a new V1beta1WorkspaceTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1WorkspaceTemplateSpecWithDefaults

`func NewV1beta1WorkspaceTemplateSpecWithDefaults() *V1beta1WorkspaceTemplateSpec`

NewV1beta1WorkspaceTemplateSpecWithDefaults instantiates a new V1beta1WorkspaceTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacement

`func (o *V1beta1WorkspaceTemplateSpec) GetPlacement() V1beta1GenericPlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *V1beta1WorkspaceTemplateSpec) GetPlacementOk() (*V1beta1GenericPlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *V1beta1WorkspaceTemplateSpec) SetPlacement(v V1beta1GenericPlacement)`

SetPlacement sets Placement field to given value.


### GetTemplate

`func (o *V1beta1WorkspaceTemplateSpec) GetTemplate() V1beta1Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1beta1WorkspaceTemplateSpec) GetTemplateOk() (*V1beta1Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1beta1WorkspaceTemplateSpec) SetTemplate(v V1beta1Template)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


