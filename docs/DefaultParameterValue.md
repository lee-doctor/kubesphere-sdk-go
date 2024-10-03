# DefaultParameterValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDefaultParameterValue

`func NewDefaultParameterValue() *DefaultParameterValue`

NewDefaultParameterValue instantiates a new DefaultParameterValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultParameterValueWithDefaults

`func NewDefaultParameterValueWithDefaults() *DefaultParameterValue`

NewDefaultParameterValueWithDefaults instantiates a new DefaultParameterValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DefaultParameterValue) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DefaultParameterValue) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DefaultParameterValue) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DefaultParameterValue) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetName

`func (o *DefaultParameterValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefaultParameterValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefaultParameterValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefaultParameterValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *DefaultParameterValue) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DefaultParameterValue) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DefaultParameterValue) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DefaultParameterValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


