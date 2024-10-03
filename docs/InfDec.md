# InfDec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scale** | **int32** |  | 
**Unscaled** | [**BigInt**](BigInt.md) |  | 

## Methods

### NewInfDec

`func NewInfDec(scale int32, unscaled BigInt, ) *InfDec`

NewInfDec instantiates a new InfDec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfDecWithDefaults

`func NewInfDecWithDefaults() *InfDec`

NewInfDecWithDefaults instantiates a new InfDec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScale

`func (o *InfDec) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *InfDec) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *InfDec) SetScale(v int32)`

SetScale sets Scale field to given value.


### GetUnscaled

`func (o *InfDec) GetUnscaled() BigInt`

GetUnscaled returns the Unscaled field if non-nil, zero value otherwise.

### GetUnscaledOk

`func (o *InfDec) GetUnscaledOk() (*BigInt, bool)`

GetUnscaledOk returns a tuple with the Unscaled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnscaled

`func (o *InfDec) SetUnscaled(v BigInt)`

SetUnscaled sets Unscaled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


