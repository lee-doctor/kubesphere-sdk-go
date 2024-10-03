# DevopsCronData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | **string** | Cron script data. | 
**PipelineName** | Pointer to **string** | Pipeline name, if pipeline haven&#39;t created, not required&#39; | [optional] 

## Methods

### NewDevopsCronData

`func NewDevopsCronData(cron string, ) *DevopsCronData`

NewDevopsCronData instantiates a new DevopsCronData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsCronDataWithDefaults

`func NewDevopsCronDataWithDefaults() *DevopsCronData`

NewDevopsCronDataWithDefaults instantiates a new DevopsCronData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *DevopsCronData) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *DevopsCronData) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *DevopsCronData) SetCron(v string)`

SetCron sets Cron field to given value.


### GetPipelineName

`func (o *DevopsCronData) GetPipelineName() string`

GetPipelineName returns the PipelineName field if non-nil, zero value otherwise.

### GetPipelineNameOk

`func (o *DevopsCronData) GetPipelineNameOk() (*string, bool)`

GetPipelineNameOk returns a tuple with the PipelineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineName

`func (o *DevopsCronData) SetPipelineName(v string)`

SetPipelineName sets PipelineName field to given value.

### HasPipelineName

`func (o *DevopsCronData) HasPipelineName() bool`

HasPipelineName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


