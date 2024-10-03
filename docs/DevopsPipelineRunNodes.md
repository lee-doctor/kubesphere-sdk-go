# DevopsPipelineRunNodes

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
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewDevopsPipelineRunNodes

`func NewDevopsPipelineRunNodes() *DevopsPipelineRunNodes`

NewDevopsPipelineRunNodes instantiates a new DevopsPipelineRunNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsPipelineRunNodesWithDefaults

`func NewDevopsPipelineRunNodesWithDefaults() *DevopsPipelineRunNodes`

NewDevopsPipelineRunNodesWithDefaults instantiates a new DevopsPipelineRunNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsPipelineRunNodes) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsPipelineRunNodes) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsPipelineRunNodes) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsPipelineRunNodes) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsPipelineRunNodes) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsPipelineRunNodes) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsPipelineRunNodes) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsPipelineRunNodes) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsPipelineRunNodes) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsPipelineRunNodes) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsPipelineRunNodes) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsPipelineRunNodes) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCauseOfBlockage

`func (o *DevopsPipelineRunNodes) GetCauseOfBlockage() map[string]interface{}`

GetCauseOfBlockage returns the CauseOfBlockage field if non-nil, zero value otherwise.

### GetCauseOfBlockageOk

`func (o *DevopsPipelineRunNodes) GetCauseOfBlockageOk() (*map[string]interface{}, bool)`

GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfBlockage

`func (o *DevopsPipelineRunNodes) SetCauseOfBlockage(v map[string]interface{})`

SetCauseOfBlockage sets CauseOfBlockage field to given value.

### HasCauseOfBlockage

`func (o *DevopsPipelineRunNodes) HasCauseOfBlockage() bool`

HasCauseOfBlockage returns a boolean if a field has been set.

### GetDisplayDescription

`func (o *DevopsPipelineRunNodes) GetDisplayDescription() map[string]interface{}`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *DevopsPipelineRunNodes) GetDisplayDescriptionOk() (*map[string]interface{}, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *DevopsPipelineRunNodes) SetDisplayDescription(v map[string]interface{})`

SetDisplayDescription sets DisplayDescription field to given value.

### HasDisplayDescription

`func (o *DevopsPipelineRunNodes) HasDisplayDescription() bool`

HasDisplayDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsPipelineRunNodes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsPipelineRunNodes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsPipelineRunNodes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsPipelineRunNodes) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsPipelineRunNodes) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsPipelineRunNodes) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsPipelineRunNodes) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsPipelineRunNodes) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetEdges

`func (o *DevopsPipelineRunNodes) GetEdges() []map[string]interface{}`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *DevopsPipelineRunNodes) GetEdgesOk() (*[]map[string]interface{}, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *DevopsPipelineRunNodes) SetEdges(v []map[string]interface{})`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *DevopsPipelineRunNodes) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetFirstParent

`func (o *DevopsPipelineRunNodes) GetFirstParent() map[string]interface{}`

GetFirstParent returns the FirstParent field if non-nil, zero value otherwise.

### GetFirstParentOk

`func (o *DevopsPipelineRunNodes) GetFirstParentOk() (*map[string]interface{}, bool)`

GetFirstParentOk returns a tuple with the FirstParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstParent

`func (o *DevopsPipelineRunNodes) SetFirstParent(v map[string]interface{})`

SetFirstParent sets FirstParent field to given value.

### HasFirstParent

`func (o *DevopsPipelineRunNodes) HasFirstParent() bool`

HasFirstParent returns a boolean if a field has been set.

### GetId

`func (o *DevopsPipelineRunNodes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsPipelineRunNodes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsPipelineRunNodes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsPipelineRunNodes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *DevopsPipelineRunNodes) GetInput() DevopsInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *DevopsPipelineRunNodes) GetInputOk() (*DevopsInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *DevopsPipelineRunNodes) SetInput(v DevopsInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *DevopsPipelineRunNodes) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetRestartable

`func (o *DevopsPipelineRunNodes) GetRestartable() bool`

GetRestartable returns the Restartable field if non-nil, zero value otherwise.

### GetRestartableOk

`func (o *DevopsPipelineRunNodes) GetRestartableOk() (*bool, bool)`

GetRestartableOk returns a tuple with the Restartable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartable

`func (o *DevopsPipelineRunNodes) SetRestartable(v bool)`

SetRestartable sets Restartable field to given value.

### HasRestartable

`func (o *DevopsPipelineRunNodes) HasRestartable() bool`

HasRestartable returns a boolean if a field has been set.

### GetResult

`func (o *DevopsPipelineRunNodes) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsPipelineRunNodes) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsPipelineRunNodes) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsPipelineRunNodes) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsPipelineRunNodes) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsPipelineRunNodes) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsPipelineRunNodes) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsPipelineRunNodes) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsPipelineRunNodes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsPipelineRunNodes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsPipelineRunNodes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsPipelineRunNodes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DevopsPipelineRunNodes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsPipelineRunNodes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsPipelineRunNodes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsPipelineRunNodes) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


