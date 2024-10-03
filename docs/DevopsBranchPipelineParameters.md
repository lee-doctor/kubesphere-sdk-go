# DevopsBranchPipelineParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Choices** | Pointer to **[]map[string]interface{}** | choices | [optional] 
**DefaultParameterValue** | Pointer to [**DefaultParameterValue**](DefaultParameterValue.md) |  | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewDevopsBranchPipelineParameters

`func NewDevopsBranchPipelineParameters() *DevopsBranchPipelineParameters`

NewDevopsBranchPipelineParameters instantiates a new DevopsBranchPipelineParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsBranchPipelineParametersWithDefaults

`func NewDevopsBranchPipelineParametersWithDefaults() *DevopsBranchPipelineParameters`

NewDevopsBranchPipelineParametersWithDefaults instantiates a new DevopsBranchPipelineParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsBranchPipelineParameters) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsBranchPipelineParameters) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsBranchPipelineParameters) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsBranchPipelineParameters) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetChoices

`func (o *DevopsBranchPipelineParameters) GetChoices() []map[string]interface{}`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *DevopsBranchPipelineParameters) GetChoicesOk() (*[]map[string]interface{}, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *DevopsBranchPipelineParameters) SetChoices(v []map[string]interface{})`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *DevopsBranchPipelineParameters) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetDefaultParameterValue

`func (o *DevopsBranchPipelineParameters) GetDefaultParameterValue() DefaultParameterValue`

GetDefaultParameterValue returns the DefaultParameterValue field if non-nil, zero value otherwise.

### GetDefaultParameterValueOk

`func (o *DevopsBranchPipelineParameters) GetDefaultParameterValueOk() (*DefaultParameterValue, bool)`

GetDefaultParameterValueOk returns a tuple with the DefaultParameterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterValue

`func (o *DevopsBranchPipelineParameters) SetDefaultParameterValue(v DefaultParameterValue)`

SetDefaultParameterValue sets DefaultParameterValue field to given value.

### HasDefaultParameterValue

`func (o *DevopsBranchPipelineParameters) HasDefaultParameterValue() bool`

HasDefaultParameterValue returns a boolean if a field has been set.

### GetDescription

`func (o *DevopsBranchPipelineParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DevopsBranchPipelineParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DevopsBranchPipelineParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DevopsBranchPipelineParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DevopsBranchPipelineParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsBranchPipelineParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsBranchPipelineParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevopsBranchPipelineParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DevopsBranchPipelineParameters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsBranchPipelineParameters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsBranchPipelineParameters) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsBranchPipelineParameters) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


