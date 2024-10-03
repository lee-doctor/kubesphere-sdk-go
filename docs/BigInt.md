# BigInt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abs** | **[]int32** |  | 
**Neg** | **bool** |  | 

## Methods

### NewBigInt

`func NewBigInt(abs []int32, neg bool, ) *BigInt`

NewBigInt instantiates a new BigInt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigIntWithDefaults

`func NewBigIntWithDefaults() *BigInt`

NewBigIntWithDefaults instantiates a new BigInt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbs

`func (o *BigInt) GetAbs() []int32`

GetAbs returns the Abs field if non-nil, zero value otherwise.

### GetAbsOk

`func (o *BigInt) GetAbsOk() (*[]int32, bool)`

GetAbsOk returns a tuple with the Abs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbs

`func (o *BigInt) SetAbs(v []int32)`

SetAbs sets Abs field to given value.


### GetNeg

`func (o *BigInt) GetNeg() bool`

GetNeg returns the Neg field if non-nil, zero value otherwise.

### GetNegOk

`func (o *BigInt) GetNegOk() (*bool, bool)`

GetNegOk returns a tuple with the Neg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeg

`func (o *BigInt) SetNeg(v bool)`

SetNeg sets Neg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


