# V1alpha2WorkspaceTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1beta1WorkspaceSpec**](V1beta1WorkspaceSpec.md) |  | [optional] 

## Methods

### NewV1alpha2WorkspaceTemplateSpec

`func NewV1alpha2WorkspaceTemplateSpec() *V1alpha2WorkspaceTemplateSpec`

NewV1alpha2WorkspaceTemplateSpec instantiates a new V1alpha2WorkspaceTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2WorkspaceTemplateSpecWithDefaults

`func NewV1alpha2WorkspaceTemplateSpecWithDefaults() *V1alpha2WorkspaceTemplateSpec`

NewV1alpha2WorkspaceTemplateSpecWithDefaults instantiates a new V1alpha2WorkspaceTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1alpha2WorkspaceTemplateSpec) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha2WorkspaceTemplateSpec) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha2WorkspaceTemplateSpec) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha2WorkspaceTemplateSpec) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1alpha2WorkspaceTemplateSpec) GetSpec() V1beta1WorkspaceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1alpha2WorkspaceTemplateSpec) GetSpecOk() (*V1beta1WorkspaceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1alpha2WorkspaceTemplateSpec) SetSpec(v V1beta1WorkspaceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1alpha2WorkspaceTemplateSpec) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


