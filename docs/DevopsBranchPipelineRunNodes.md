# DevopsBranchPipelineRunNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions | [optional] 
**CauseOfBlockage** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayDescription** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayName** | Pointer to **string** | display name | [optional] 
**DurationInMillis** | Pointer to **int32** | duration time in millis | [optional] 
**Edges** | Pointer to [**[]DevopsBranchPipelineRunNodesEdges**](DevopsBranchPipelineRunNodesEdges.md) |  | [optional] 
**FirstParent** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Input** | Pointer to [**DevopsInput**](DevopsInput.md) |  | [optional] 
**Restartable** | Pointer to **bool** | restartable or not | [optional] 
**Result** | Pointer to **string** | the result of pipeline run. e.g. SUCCESS. e.g. SUCCESS | [optional] 
**StartTime** | Pointer to **string** | the time of start | [optional] 
**State** | Pointer to **string** | run state. e.g. RUNNING | [optional] 
**Steps** | Pointer to [**[]DevopsBranchPipelineRunNodesSteps**](DevopsBranchPipelineRunNodesSteps.md) |  | [optional] 
**Type** | Pointer to **string** | source type, e.g. \&quot;WorkflowRun\&quot; | [optional] 

## Methods

### NewDevopsBranchPipelineRunNodes

`func NewDevopsBranchPipelineRunNodes() *DevopsBranchPipelineRunNodes`

NewDevopsBranchPipelineRunNodes instantiates a new DevopsBranchPipelineRunNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsBranchPipelineRunNodesWithDefaults

`func NewDevopsBranchPipelineRunNodesWithDefaults() *DevopsBranchPipelineRunNodes`

NewDevopsBranchPipelineRunNodesWithDefaults instantiates a new DevopsBranchPipelineRunNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsBranchPipelineRunNodes) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsBranchPipelineRunNodes) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsBranchPipelineRunNodes) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsBranchPipelineRunNodes) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsBranchPipelineRunNodes) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsBranchPipelineRunNodes) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsBranchPipelineRunNodes) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsBranchPipelineRunNodes) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsBranchPipelineRunNodes) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsBranchPipelineRunNodes) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsBranchPipelineRunNodes) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsBranchPipelineRunNodes) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCauseOfBlockage

`func (o *DevopsBranchPipelineRunNodes) GetCauseOfBlockage() map[string]interface{}`

GetCauseOfBlockage returns the CauseOfBlockage field if non-nil, zero value otherwise.

### GetCauseOfBlockageOk

`func (o *DevopsBranchPipelineRunNodes) GetCauseOfBlockageOk() (*map[string]interface{}, bool)`

GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfBlockage

`func (o *DevopsBranchPipelineRunNodes) SetCauseOfBlockage(v map[string]interface{})`

SetCauseOfBlockage sets CauseOfBlockage field to given value.

### HasCauseOfBlockage

`func (o *DevopsBranchPipelineRunNodes) HasCauseOfBlockage() bool`

HasCauseOfBlockage returns a boolean if a field has been set.

### GetDisplayDescription

`func (o *DevopsBranchPipelineRunNodes) GetDisplayDescription() map[string]interface{}`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *DevopsBranchPipelineRunNodes) GetDisplayDescriptionOk() (*map[string]interface{}, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *DevopsBranchPipelineRunNodes) SetDisplayDescription(v map[string]interface{})`

SetDisplayDescription sets DisplayDescription field to given value.

### HasDisplayDescription

`func (o *DevopsBranchPipelineRunNodes) HasDisplayDescription() bool`

HasDisplayDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsBranchPipelineRunNodes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsBranchPipelineRunNodes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsBranchPipelineRunNodes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsBranchPipelineRunNodes) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsBranchPipelineRunNodes) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsBranchPipelineRunNodes) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsBranchPipelineRunNodes) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsBranchPipelineRunNodes) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetEdges

`func (o *DevopsBranchPipelineRunNodes) GetEdges() []DevopsBranchPipelineRunNodesEdges`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *DevopsBranchPipelineRunNodes) GetEdgesOk() (*[]DevopsBranchPipelineRunNodesEdges, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *DevopsBranchPipelineRunNodes) SetEdges(v []DevopsBranchPipelineRunNodesEdges)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *DevopsBranchPipelineRunNodes) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetFirstParent

`func (o *DevopsBranchPipelineRunNodes) GetFirstParent() map[string]interface{}`

GetFirstParent returns the FirstParent field if non-nil, zero value otherwise.

### GetFirstParentOk

`func (o *DevopsBranchPipelineRunNodes) GetFirstParentOk() (*map[string]interface{}, bool)`

GetFirstParentOk returns a tuple with the FirstParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstParent

`func (o *DevopsBranchPipelineRunNodes) SetFirstParent(v map[string]interface{})`

SetFirstParent sets FirstParent field to given value.

### HasFirstParent

`func (o *DevopsBranchPipelineRunNodes) HasFirstParent() bool`

HasFirstParent returns a boolean if a field has been set.

### GetId

`func (o *DevopsBranchPipelineRunNodes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsBranchPipelineRunNodes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsBranchPipelineRunNodes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsBranchPipelineRunNodes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *DevopsBranchPipelineRunNodes) GetInput() DevopsInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DevopsBranchPipelineRunNodes) GetInputOk() (*DevopsInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DevopsBranchPipelineRunNodes) SetInput(v DevopsInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *DevopsBranchPipelineRunNodes) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetRestartable

`func (o *DevopsBranchPipelineRunNodes) GetRestartable() bool`

GetRestartable returns the Restartable field if non-nil, zero value otherwise.

### GetRestartableOk

`func (o *DevopsBranchPipelineRunNodes) GetRestartableOk() (*bool, bool)`

GetRestartableOk returns a tuple with the Restartable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartable

`func (o *DevopsBranchPipelineRunNodes) SetRestartable(v bool)`

SetRestartable sets Restartable field to given value.

### HasRestartable

`func (o *DevopsBranchPipelineRunNodes) HasRestartable() bool`

HasRestartable returns a boolean if a field has been set.

### GetResult

`func (o *DevopsBranchPipelineRunNodes) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsBranchPipelineRunNodes) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsBranchPipelineRunNodes) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsBranchPipelineRunNodes) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsBranchPipelineRunNodes) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsBranchPipelineRunNodes) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsBranchPipelineRunNodes) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsBranchPipelineRunNodes) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsBranchPipelineRunNodes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsBranchPipelineRunNodes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsBranchPipelineRunNodes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsBranchPipelineRunNodes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *DevopsBranchPipelineRunNodes) GetSteps() []DevopsBranchPipelineRunNodesSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DevopsBranchPipelineRunNodes) GetStepsOk() (*[]DevopsBranchPipelineRunNodesSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DevopsBranchPipelineRunNodes) SetSteps(v []DevopsBranchPipelineRunNodesSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *DevopsBranchPipelineRunNodes) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetType

`func (o *DevopsBranchPipelineRunNodes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsBranchPipelineRunNodes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsBranchPipelineRunNodes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsBranchPipelineRunNodes) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


