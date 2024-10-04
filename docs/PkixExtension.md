# PkixExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Critical** | **bool** |  | 
**Id** | **[]int32** |  | 
**Value** | **string** |  | 

## Methods

### NewPkixExtension

`func NewPkixExtension(critical bool, id []int32, value string, ) *PkixExtension`

NewPkixExtension instantiates a new PkixExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkixExtensionWithDefaults

`func NewPkixExtensionWithDefaults() *PkixExtension`

NewPkixExtensionWithDefaults instantiates a new PkixExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCritical

`func (o *PkixExtension) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *PkixExtension) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *PkixExtension) SetCritical(v bool)`

SetCritical sets Critical field to given value.


### GetId

`func (o *PkixExtension) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkixExtension) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkixExtension) SetId(v []int32)`

SetId sets Id field to given value.


### GetValue

`func (o *PkixExtension) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PkixExtension) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PkixExtension) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


