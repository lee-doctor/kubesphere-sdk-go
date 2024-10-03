# V1alpha3Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | default value of param | [optional] 
**Description** | Pointer to **string** | description of pipeline | [optional] 
**Name** | **string** | name of param | 
**Type** | **string** | type of param | 

## Methods

### NewV1alpha3Parameter

`func NewV1alpha3Parameter(name string, type_ string, ) *V1alpha3Parameter`

NewV1alpha3Parameter instantiates a new V1alpha3Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3ParameterWithDefaults

`func NewV1alpha3ParameterWithDefaults() *V1alpha3Parameter`

NewV1alpha3ParameterWithDefaults instantiates a new V1alpha3Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *V1alpha3Parameter) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *V1alpha3Parameter) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *V1alpha3Parameter) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *V1alpha3Parameter) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescription

`func (o *V1alpha3Parameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1alpha3Parameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1alpha3Parameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1alpha3Parameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *V1alpha3Parameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha3Parameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha3Parameter) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *V1alpha3Parameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha3Parameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha3Parameter) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


