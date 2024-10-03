# ResourceQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**D** | [**ResourceInfDecAmount**](ResourceInfDecAmount.md) |  | 
**I** | [**ResourceInt64Amount**](ResourceInt64Amount.md) |  | 
**S** | **string** |  | 

## Methods

### NewResourceQuantity

`func NewResourceQuantity(format string, d ResourceInfDecAmount, i ResourceInt64Amount, s string, ) *ResourceQuantity`

NewResourceQuantity instantiates a new ResourceQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceQuantityWithDefaults

`func NewResourceQuantityWithDefaults() *ResourceQuantity`

NewResourceQuantityWithDefaults instantiates a new ResourceQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ResourceQuantity) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ResourceQuantity) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ResourceQuantity) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetD

`func (o *ResourceQuantity) GetD() ResourceInfDecAmount`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *ResourceQuantity) GetDOk() (*ResourceInfDecAmount, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *ResourceQuantity) SetD(v ResourceInfDecAmount)`

SetD sets D field to given value.


### GetI

`func (o *ResourceQuantity) GetI() ResourceInt64Amount`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *ResourceQuantity) GetIOk() (*ResourceInt64Amount, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *ResourceQuantity) SetI(v ResourceInt64Amount)`

SetI sets I field to given value.


### GetS

`func (o *ResourceQuantity) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ResourceQuantity) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ResourceQuantity) SetS(v string)`

SetS sets S field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


