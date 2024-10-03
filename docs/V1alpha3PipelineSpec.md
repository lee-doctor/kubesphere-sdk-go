# V1alpha3PipelineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiBranchPipeline** | Pointer to [**V1alpha3MultiBranchPipeline**](V1alpha3MultiBranchPipeline.md) |  | [optional] 
**Pipeline** | Pointer to [**V1alpha3NoScmPipeline**](V1alpha3NoScmPipeline.md) |  | [optional] 
**Type** | **string** | type of devops pipeline, in scm or no scm | 

## Methods

### NewV1alpha3PipelineSpec

`func NewV1alpha3PipelineSpec(type_ string, ) *V1alpha3PipelineSpec`

NewV1alpha3PipelineSpec instantiates a new V1alpha3PipelineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3PipelineSpecWithDefaults

`func NewV1alpha3PipelineSpecWithDefaults() *V1alpha3PipelineSpec`

NewV1alpha3PipelineSpecWithDefaults instantiates a new V1alpha3PipelineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultiBranchPipeline

`func (o *V1alpha3PipelineSpec) GetMultiBranchPipeline() V1alpha3MultiBranchPipeline`

GetMultiBranchPipeline returns the MultiBranchPipeline field if non-nil, zero value otherwise.

### GetMultiBranchPipelineOk

`func (o *V1alpha3PipelineSpec) GetMultiBranchPipelineOk() (*V1alpha3MultiBranchPipeline, bool)`

GetMultiBranchPipelineOk returns a tuple with the MultiBranchPipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiBranchPipeline

`func (o *V1alpha3PipelineSpec) SetMultiBranchPipeline(v V1alpha3MultiBranchPipeline)`

SetMultiBranchPipeline sets MultiBranchPipeline field to given value.

### HasMultiBranchPipeline

`func (o *V1alpha3PipelineSpec) HasMultiBranchPipeline() bool`

HasMultiBranchPipeline returns a boolean if a field has been set.

### GetPipeline

`func (o *V1alpha3PipelineSpec) GetPipeline() V1alpha3NoScmPipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *V1alpha3PipelineSpec) GetPipelineOk() (*V1alpha3NoScmPipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *V1alpha3PipelineSpec) SetPipeline(v V1alpha3NoScmPipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *V1alpha3PipelineSpec) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetType

`func (o *V1alpha3PipelineSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha3PipelineSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha3PipelineSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


