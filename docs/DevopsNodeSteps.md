# DevopsNodeSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to [**[]DevopsNodeStepsActions**](DevopsNodeStepsActions.md) |  | [optional] 
**Aprovable** | **bool** | indicate if this step can be approved by current user | 
**DisplayDescription** | Pointer to **string** | display description | [optional] 
**DisplayName** | Pointer to **string** | display name | [optional] 
**DurationInMillis** | Pointer to **int32** | duration time in mullis | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Input** | Pointer to [**DevopsInput**](DevopsInput.md) |  | [optional] 
**Result** | Pointer to **string** | the result of pipeline run. e.g. SUCCESS | [optional] 
**StartTime** | Pointer to **string** | the time of starts | [optional] 
**State** | Pointer to **string** | run state. e.g. SKIPPED | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewDevopsNodeSteps

`func NewDevopsNodeSteps(aprovable bool, ) *DevopsNodeSteps`

NewDevopsNodeSteps instantiates a new DevopsNodeSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsNodeStepsWithDefaults

`func NewDevopsNodeStepsWithDefaults() *DevopsNodeSteps`

NewDevopsNodeStepsWithDefaults instantiates a new DevopsNodeSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsNodeSteps) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsNodeSteps) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsNodeSteps) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsNodeSteps) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsNodeSteps) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsNodeSteps) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsNodeSteps) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsNodeSteps) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsNodeSteps) GetActions() []DevopsNodeStepsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsNodeSteps) GetActionsOk() (*[]DevopsNodeStepsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsNodeSteps) SetActions(v []DevopsNodeStepsActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsNodeSteps) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAprovable

`func (o *DevopsNodeSteps) GetAprovable() bool`

GetAprovable returns the Aprovable field if non-nil, zero value otherwise.

### GetAprovableOk

`func (o *DevopsNodeSteps) GetAprovableOk() (*bool, bool)`

GetAprovableOk returns a tuple with the Aprovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprovable

`func (o *DevopsNodeSteps) SetAprovable(v bool)`

SetAprovable sets Aprovable field to given value.


### GetDisplayDescription

`func (o *DevopsNodeSteps) GetDisplayDescription() string`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *DevopsNodeSteps) GetDisplayDescriptionOk() (*string, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *DevopsNodeSteps) SetDisplayDescription(v string)`

SetDisplayDescription sets DisplayDescription field to given value.

### HasDisplayDescription

`func (o *DevopsNodeSteps) HasDisplayDescription() bool`

HasDisplayDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsNodeSteps) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsNodeSteps) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsNodeSteps) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsNodeSteps) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsNodeSteps) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsNodeSteps) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsNodeSteps) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsNodeSteps) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetId

`func (o *DevopsNodeSteps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsNodeSteps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsNodeSteps) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsNodeSteps) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *DevopsNodeSteps) GetInput() DevopsInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DevopsNodeSteps) GetInputOk() (*DevopsInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DevopsNodeSteps) SetInput(v DevopsInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *DevopsNodeSteps) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetResult

`func (o *DevopsNodeSteps) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsNodeSteps) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsNodeSteps) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsNodeSteps) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsNodeSteps) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsNodeSteps) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsNodeSteps) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsNodeSteps) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsNodeSteps) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsNodeSteps) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsNodeSteps) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsNodeSteps) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DevopsNodeSteps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsNodeSteps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsNodeSteps) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsNodeSteps) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


