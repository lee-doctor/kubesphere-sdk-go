# DevopsNodesDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions | [optional] 
**CauseOfBlockage** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayDescription** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayName** | Pointer to **string** | display name | [optional] 
**DurationInMillis** | Pointer to **int32** | duration time in mullis | [optional] 
**Edges** | Pointer to **[]map[string]interface{}** | edges | [optional] 
**FirstParent** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Input** | Pointer to [**DevopsInput**](DevopsInput.md) |  | [optional] 
**Restartable** | Pointer to **bool** | restartable or not | [optional] 
**Result** | Pointer to **string** | the result of pipeline run. e.g. SUCCESS | [optional] 
**StartTime** | Pointer to **string** | the time of start | [optional] 
**State** | Pointer to **string** | run state. e.g. FINISHED | [optional] 
**Steps** | Pointer to [**[]DevopsNodeSteps**](DevopsNodeSteps.md) | steps | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewDevopsNodesDetail

`func NewDevopsNodesDetail() *DevopsNodesDetail`

NewDevopsNodesDetail instantiates a new DevopsNodesDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsNodesDetailWithDefaults

`func NewDevopsNodesDetailWithDefaults() *DevopsNodesDetail`

NewDevopsNodesDetailWithDefaults instantiates a new DevopsNodesDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsNodesDetail) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsNodesDetail) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsNodesDetail) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsNodesDetail) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsNodesDetail) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsNodesDetail) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsNodesDetail) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsNodesDetail) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsNodesDetail) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsNodesDetail) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsNodesDetail) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsNodesDetail) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCauseOfBlockage

`func (o *DevopsNodesDetail) GetCauseOfBlockage() map[string]interface{}`

GetCauseOfBlockage returns the CauseOfBlockage field if non-nil, zero value otherwise.

### GetCauseOfBlockageOk

`func (o *DevopsNodesDetail) GetCauseOfBlockageOk() (*map[string]interface{}, bool)`

GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfBlockage

`func (o *DevopsNodesDetail) SetCauseOfBlockage(v map[string]interface{})`

SetCauseOfBlockage sets CauseOfBlockage field to given value.

### HasCauseOfBlockage

`func (o *DevopsNodesDetail) HasCauseOfBlockage() bool`

HasCauseOfBlockage returns a boolean if a field has been set.

### GetDisplayDescription

`func (o *DevopsNodesDetail) GetDisplayDescription() map[string]interface{}`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *DevopsNodesDetail) GetDisplayDescriptionOk() (*map[string]interface{}, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *DevopsNodesDetail) SetDisplayDescription(v map[string]interface{})`

SetDisplayDescription sets DisplayDescription field to given value.

### HasDisplayDescription

`func (o *DevopsNodesDetail) HasDisplayDescription() bool`

HasDisplayDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsNodesDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsNodesDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsNodesDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsNodesDetail) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsNodesDetail) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsNodesDetail) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsNodesDetail) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsNodesDetail) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetEdges

`func (o *DevopsNodesDetail) GetEdges() []map[string]interface{}`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *DevopsNodesDetail) GetEdgesOk() (*[]map[string]interface{}, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *DevopsNodesDetail) SetEdges(v []map[string]interface{})`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *DevopsNodesDetail) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetFirstParent

`func (o *DevopsNodesDetail) GetFirstParent() map[string]interface{}`

GetFirstParent returns the FirstParent field if non-nil, zero value otherwise.

### GetFirstParentOk

`func (o *DevopsNodesDetail) GetFirstParentOk() (*map[string]interface{}, bool)`

GetFirstParentOk returns a tuple with the FirstParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstParent

`func (o *DevopsNodesDetail) SetFirstParent(v map[string]interface{})`

SetFirstParent sets FirstParent field to given value.

### HasFirstParent

`func (o *DevopsNodesDetail) HasFirstParent() bool`

HasFirstParent returns a boolean if a field has been set.

### GetId

`func (o *DevopsNodesDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsNodesDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsNodesDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsNodesDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *DevopsNodesDetail) GetInput() DevopsInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DevopsNodesDetail) GetInputOk() (*DevopsInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DevopsNodesDetail) SetInput(v DevopsInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *DevopsNodesDetail) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetRestartable

`func (o *DevopsNodesDetail) GetRestartable() bool`

GetRestartable returns the Restartable field if non-nil, zero value otherwise.

### GetRestartableOk

`func (o *DevopsNodesDetail) GetRestartableOk() (*bool, bool)`

GetRestartableOk returns a tuple with the Restartable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartable

`func (o *DevopsNodesDetail) SetRestartable(v bool)`

SetRestartable sets Restartable field to given value.

### HasRestartable

`func (o *DevopsNodesDetail) HasRestartable() bool`

HasRestartable returns a boolean if a field has been set.

### GetResult

`func (o *DevopsNodesDetail) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsNodesDetail) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsNodesDetail) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsNodesDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsNodesDetail) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsNodesDetail) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsNodesDetail) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsNodesDetail) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsNodesDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsNodesDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsNodesDetail) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsNodesDetail) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *DevopsNodesDetail) GetSteps() []DevopsNodeSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DevopsNodesDetail) GetStepsOk() (*[]DevopsNodeSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DevopsNodesDetail) SetSteps(v []DevopsNodeSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *DevopsNodesDetail) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetType

`func (o *DevopsNodesDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsNodesDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsNodesDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsNodesDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


