# V1alpha3GitCloneOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | Pointer to **int32** | git clone depth | [optional] 
**Shallow** | Pointer to **bool** | Whether to use git shallow clone | [optional] 
**Timeout** | Pointer to **int32** | git clone timeout mins | [optional] 

## Methods

### NewV1alpha3GitCloneOption

`func NewV1alpha3GitCloneOption() *V1alpha3GitCloneOption`

NewV1alpha3GitCloneOption instantiates a new V1alpha3GitCloneOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3GitCloneOptionWithDefaults

`func NewV1alpha3GitCloneOptionWithDefaults() *V1alpha3GitCloneOption`

NewV1alpha3GitCloneOptionWithDefaults instantiates a new V1alpha3GitCloneOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *V1alpha3GitCloneOption) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *V1alpha3GitCloneOption) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *V1alpha3GitCloneOption) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *V1alpha3GitCloneOption) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetShallow

`func (o *V1alpha3GitCloneOption) GetShallow() bool`

GetShallow returns the Shallow field if non-nil, zero value otherwise.

### GetShallowOk

`func (o *V1alpha3GitCloneOption) GetShallowOk() (*bool, bool)`

GetShallowOk returns a tuple with the Shallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShallow

`func (o *V1alpha3GitCloneOption) SetShallow(v bool)`

SetShallow sets Shallow field to given value.

### HasShallow

`func (o *V1alpha3GitCloneOption) HasShallow() bool`

HasShallow returns a boolean if a field has been set.

### GetTimeout

`func (o *V1alpha3GitCloneOption) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *V1alpha3GitCloneOption) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *V1alpha3GitCloneOption) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *V1alpha3GitCloneOption) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


