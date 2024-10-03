# V1alpha3Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1alpha3PipelineSpec**](V1alpha3PipelineSpec.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewV1alpha3Pipeline

`func NewV1alpha3Pipeline() *V1alpha3Pipeline`

NewV1alpha3Pipeline instantiates a new V1alpha3Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3PipelineWithDefaults

`func NewV1alpha3PipelineWithDefaults() *V1alpha3Pipeline`

NewV1alpha3PipelineWithDefaults instantiates a new V1alpha3Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1alpha3Pipeline) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1alpha3Pipeline) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1alpha3Pipeline) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1alpha3Pipeline) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1alpha3Pipeline) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1alpha3Pipeline) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1alpha3Pipeline) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1alpha3Pipeline) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1alpha3Pipeline) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha3Pipeline) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha3Pipeline) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha3Pipeline) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1alpha3Pipeline) GetSpec() V1alpha3PipelineSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1alpha3Pipeline) GetSpecOk() (*V1alpha3PipelineSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1alpha3Pipeline) SetSpec(v V1alpha3PipelineSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1alpha3Pipeline) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1alpha3Pipeline) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha3Pipeline) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha3Pipeline) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha3Pipeline) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


