# DevopsRunPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions | [optional] 
**ArtifactsZipFile** | Pointer to **map[string]interface{}** |  | [optional] 
**CauseOfBlockage** | Pointer to **string** | the cause of blockage | [optional] 
**Causes** | Pointer to [**[]DevopsRunPipelineCauses**](DevopsRunPipelineCauses.md) |  | [optional] 
**ChangeSet** | Pointer to **[]map[string]interface{}** | changeset information | [optional] 
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**DurationInMillis** | Pointer to **map[string]interface{}** |  | [optional] 
**EnQueueTime** | Pointer to **map[string]interface{}** |  | [optional] 
**EndTime** | Pointer to **map[string]interface{}** |  | [optional] 
**EstimatedDurationInMillis** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Name** | Pointer to **map[string]interface{}** |  | [optional] 
**Organization** | Pointer to **string** | the name of organization | [optional] 
**Pipeline** | Pointer to **string** | pipeline | [optional] 
**QueueId** | Pointer to **string** | queue id | [optional] 
**Replayable** | Pointer to **bool** | replayable or not | [optional] 
**Result** | Pointer to **string** | the result of pipeline run. e.g. SUCCESS | [optional] 
**RunSummary** | Pointer to **map[string]interface{}** |  | [optional] 
**StartTime** | Pointer to **map[string]interface{}** |  | [optional] 
**State** | Pointer to **string** | run state. e.g. RUNNING | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewDevopsRunPipeline

`func NewDevopsRunPipeline() *DevopsRunPipeline`

NewDevopsRunPipeline instantiates a new DevopsRunPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsRunPipelineWithDefaults

`func NewDevopsRunPipelineWithDefaults() *DevopsRunPipeline`

NewDevopsRunPipelineWithDefaults instantiates a new DevopsRunPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsRunPipeline) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsRunPipeline) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsRunPipeline) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsRunPipeline) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsRunPipeline) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsRunPipeline) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsRunPipeline) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsRunPipeline) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsRunPipeline) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsRunPipeline) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsRunPipeline) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsRunPipeline) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetArtifactsZipFile

`func (o *DevopsRunPipeline) GetArtifactsZipFile() map[string]interface{}`

GetArtifactsZipFile returns the ArtifactsZipFile field if non-nil, zero value otherwise.

### GetArtifactsZipFileOk

`func (o *DevopsRunPipeline) GetArtifactsZipFileOk() (*map[string]interface{}, bool)`

GetArtifactsZipFileOk returns a tuple with the ArtifactsZipFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsZipFile

`func (o *DevopsRunPipeline) SetArtifactsZipFile(v map[string]interface{})`

SetArtifactsZipFile sets ArtifactsZipFile field to given value.

### HasArtifactsZipFile

`func (o *DevopsRunPipeline) HasArtifactsZipFile() bool`

HasArtifactsZipFile returns a boolean if a field has been set.

### GetCauseOfBlockage

`func (o *DevopsRunPipeline) GetCauseOfBlockage() string`

GetCauseOfBlockage returns the CauseOfBlockage field if non-nil, zero value otherwise.

### GetCauseOfBlockageOk

`func (o *DevopsRunPipeline) GetCauseOfBlockageOk() (*string, bool)`

GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfBlockage

`func (o *DevopsRunPipeline) SetCauseOfBlockage(v string)`

SetCauseOfBlockage sets CauseOfBlockage field to given value.

### HasCauseOfBlockage

`func (o *DevopsRunPipeline) HasCauseOfBlockage() bool`

HasCauseOfBlockage returns a boolean if a field has been set.

### GetCauses

`func (o *DevopsRunPipeline) GetCauses() []DevopsRunPipelineCauses`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *DevopsRunPipeline) GetCausesOk() (*[]DevopsRunPipelineCauses, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *DevopsRunPipeline) SetCauses(v []DevopsRunPipelineCauses)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *DevopsRunPipeline) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetChangeSet

`func (o *DevopsRunPipeline) GetChangeSet() []map[string]interface{}`

GetChangeSet returns the ChangeSet field if non-nil, zero value otherwise.

### GetChangeSetOk

`func (o *DevopsRunPipeline) GetChangeSetOk() (*[]map[string]interface{}, bool)`

GetChangeSetOk returns a tuple with the ChangeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSet

`func (o *DevopsRunPipeline) SetChangeSet(v []map[string]interface{})`

SetChangeSet sets ChangeSet field to given value.

### HasChangeSet

`func (o *DevopsRunPipeline) HasChangeSet() bool`

HasChangeSet returns a boolean if a field has been set.

### GetDescription

`func (o *DevopsRunPipeline) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DevopsRunPipeline) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DevopsRunPipeline) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DevopsRunPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsRunPipeline) GetDurationInMillis() map[string]interface{}`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsRunPipeline) GetDurationInMillisOk() (*map[string]interface{}, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsRunPipeline) SetDurationInMillis(v map[string]interface{})`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsRunPipeline) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetEnQueueTime

`func (o *DevopsRunPipeline) GetEnQueueTime() map[string]interface{}`

GetEnQueueTime returns the EnQueueTime field if non-nil, zero value otherwise.

### GetEnQueueTimeOk

`func (o *DevopsRunPipeline) GetEnQueueTimeOk() (*map[string]interface{}, bool)`

GetEnQueueTimeOk returns a tuple with the EnQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnQueueTime

`func (o *DevopsRunPipeline) SetEnQueueTime(v map[string]interface{})`

SetEnQueueTime sets EnQueueTime field to given value.

### HasEnQueueTime

`func (o *DevopsRunPipeline) HasEnQueueTime() bool`

HasEnQueueTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DevopsRunPipeline) GetEndTime() map[string]interface{}`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DevopsRunPipeline) GetEndTimeOk() (*map[string]interface{}, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DevopsRunPipeline) SetEndTime(v map[string]interface{})`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DevopsRunPipeline) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEstimatedDurationInMillis

`func (o *DevopsRunPipeline) GetEstimatedDurationInMillis() map[string]interface{}`

GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field if non-nil, zero value otherwise.

### GetEstimatedDurationInMillisOk

`func (o *DevopsRunPipeline) GetEstimatedDurationInMillisOk() (*map[string]interface{}, bool)`

GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMillis

`func (o *DevopsRunPipeline) SetEstimatedDurationInMillis(v map[string]interface{})`

SetEstimatedDurationInMillis sets EstimatedDurationInMillis field to given value.

### HasEstimatedDurationInMillis

`func (o *DevopsRunPipeline) HasEstimatedDurationInMillis() bool`

HasEstimatedDurationInMillis returns a boolean if a field has been set.

### GetId

`func (o *DevopsRunPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsRunPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsRunPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsRunPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DevopsRunPipeline) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsRunPipeline) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsRunPipeline) SetName(v map[string]interface{})`

SetName sets Name field to given value.

### HasName

`func (o *DevopsRunPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *DevopsRunPipeline) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DevopsRunPipeline) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DevopsRunPipeline) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DevopsRunPipeline) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPipeline

`func (o *DevopsRunPipeline) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DevopsRunPipeline) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DevopsRunPipeline) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *DevopsRunPipeline) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetQueueId

`func (o *DevopsRunPipeline) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *DevopsRunPipeline) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *DevopsRunPipeline) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *DevopsRunPipeline) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetReplayable

`func (o *DevopsRunPipeline) GetReplayable() bool`

GetReplayable returns the Replayable field if non-nil, zero value otherwise.

### GetReplayableOk

`func (o *DevopsRunPipeline) GetReplayableOk() (*bool, bool)`

GetReplayableOk returns a tuple with the Replayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayable

`func (o *DevopsRunPipeline) SetReplayable(v bool)`

SetReplayable sets Replayable field to given value.

### HasReplayable

`func (o *DevopsRunPipeline) HasReplayable() bool`

HasReplayable returns a boolean if a field has been set.

### GetResult

`func (o *DevopsRunPipeline) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsRunPipeline) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsRunPipeline) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsRunPipeline) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRunSummary

`func (o *DevopsRunPipeline) GetRunSummary() map[string]interface{}`

GetRunSummary returns the RunSummary field if non-nil, zero value otherwise.

### GetRunSummaryOk

`func (o *DevopsRunPipeline) GetRunSummaryOk() (*map[string]interface{}, bool)`

GetRunSummaryOk returns a tuple with the RunSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunSummary

`func (o *DevopsRunPipeline) SetRunSummary(v map[string]interface{})`

SetRunSummary sets RunSummary field to given value.

### HasRunSummary

`func (o *DevopsRunPipeline) HasRunSummary() bool`

HasRunSummary returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsRunPipeline) GetStartTime() map[string]interface{}`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsRunPipeline) GetStartTimeOk() (*map[string]interface{}, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsRunPipeline) SetStartTime(v map[string]interface{})`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsRunPipeline) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsRunPipeline) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsRunPipeline) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsRunPipeline) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsRunPipeline) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DevopsRunPipeline) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsRunPipeline) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsRunPipeline) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsRunPipeline) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


