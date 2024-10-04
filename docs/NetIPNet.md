# NetIPNet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IP** | **string** |  | 
**Mask** | **string** |  | 

## Methods

### NewNetIPNet

`func NewNetIPNet(iP string, mask string, ) *NetIPNet`

NewNetIPNet instantiates a new NetIPNet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetIPNetWithDefaults

`func NewNetIPNetWithDefaults() *NetIPNet`

NewNetIPNetWithDefaults instantiates a new NetIPNet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIP

`func (o *NetIPNet) GetIP() string`

GetIP returns the IP field if non-nil, zero value otherwise.

### GetIPOk

`func (o *NetIPNet) GetIPOk() (*string, bool)`

GetIPOk returns a tuple with the IP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIP

`func (o *NetIPNet) SetIP(v string)`

SetIP sets IP field to given value.


### GetMask

`func (o *NetIPNet) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *NetIPNet) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *NetIPNet) SetMask(v string)`

SetMask sets Mask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


