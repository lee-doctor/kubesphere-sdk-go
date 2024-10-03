# DevopsBranchPipelineRunNodesSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to [**[]DevopsBranchPipelineRunNodesStepsActions**](DevopsBranchPipelineRunNodesStepsActions.md) | references the reachable path to this resource | [optional] 
**DisplayDescription** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayName** | Pointer to **string** | display name | [optional] 
**DurationInMillis** | Pointer to **int32** | duration time in millis | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Input** | Pointer to [**DevopsInput**](DevopsInput.md) |  | [optional] 
**Result** | Pointer to **string** | result | [optional] 
**StartTime** | Pointer to **string** | the time of start | [optional] 
**State** | Pointer to **string** | run state. e.g. RUNNING | [optional] 
**Type** | Pointer to **string** | source type | [optional] 

## Methods

### NewDevopsBranchPipelineRunNodesSteps

`func NewDevopsBranchPipelineRunNodesSteps() *DevopsBranchPipelineRunNodesSteps`

NewDevopsBranchPipelineRunNodesSteps instantiates a new DevopsBranchPipelineRunNodesSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsBranchPipelineRunNodesStepsWithDefaults

`func NewDevopsBranchPipelineRunNodesStepsWithDefaults() *DevopsBranchPipelineRunNodesSteps`

NewDevopsBranchPipelineRunNodesStepsWithDefaults instantiates a new DevopsBranchPipelineRunNodesSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsBranchPipelineRunNodesSteps) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsBranchPipelineRunNodesSteps) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsBranchPipelineRunNodesSteps) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsBranchPipelineRunNodesSteps) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsBranchPipelineRunNodesSteps) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsBranchPipelineRunNodesSteps) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsBranchPipelineRunNodesSteps) GetActions() []DevopsBranchPipelineRunNodesStepsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetActionsOk() (*[]DevopsBranchPipelineRunNodesStepsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsBranchPipelineRunNodesSteps) SetActions(v []DevopsBranchPipelineRunNodesStepsActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsBranchPipelineRunNodesSteps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDisplayDescription

`func (o *DevopsBranchPipelineRunNodesSteps) GetDisplayDescription() map[string]interface{}`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetDisplayDescriptionOk() (*map[string]interface{}, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *DevopsBranchPipelineRunNodesSteps) SetDisplayDescription(v map[string]interface{})`

SetDisplayDescription sets DisplayDescription field to given value.

### HasDisplayDescription

`func (o *DevopsBranchPipelineRunNodesSteps) HasDisplayDescription() bool`

HasDisplayDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsBranchPipelineRunNodesSteps) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsBranchPipelineRunNodesSteps) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsBranchPipelineRunNodesSteps) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsBranchPipelineRunNodesSteps) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsBranchPipelineRunNodesSteps) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsBranchPipelineRunNodesSteps) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetId

`func (o *DevopsBranchPipelineRunNodesSteps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsBranchPipelineRunNodesSteps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsBranchPipelineRunNodesSteps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *DevopsBranchPipelineRunNodesSteps) GetInput() DevopsInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetInputOk() (*DevopsInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DevopsBranchPipelineRunNodesSteps) SetInput(v DevopsInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *DevopsBranchPipelineRunNodesSteps) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetResult

`func (o *DevopsBranchPipelineRunNodesSteps) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsBranchPipelineRunNodesSteps) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsBranchPipelineRunNodesSteps) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsBranchPipelineRunNodesSteps) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsBranchPipelineRunNodesSteps) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsBranchPipelineRunNodesSteps) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsBranchPipelineRunNodesSteps) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsBranchPipelineRunNodesSteps) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsBranchPipelineRunNodesSteps) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DevopsBranchPipelineRunNodesSteps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsBranchPipelineRunNodesSteps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsBranchPipelineRunNodesSteps) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsBranchPipelineRunNodesSteps) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


