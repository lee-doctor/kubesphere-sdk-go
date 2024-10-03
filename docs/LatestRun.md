# LatestRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions | [optional] 
**ArtifactsZipFile** | Pointer to **string** | the artifacts zip file | [optional] 
**CauseOfBlockage** | Pointer to **map[string]interface{}** |  | [optional] 
**Causes** | Pointer to [**[]LatestRunCauses**](LatestRunCauses.md) |  | [optional] 
**ChangeSet** | Pointer to **[]map[string]interface{}** | changeset information | [optional] 
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**DurationInMillis** | Pointer to **int32** | duration time in millis | [optional] 
**EnQueueTime** | Pointer to **string** | the time of enter the queue | [optional] 
**EndTime** | Pointer to **string** | the time of end | [optional] 
**EstimatedDurationInMillis** | Pointer to **int32** | estimated duration time in millis | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Name** | Pointer to **map[string]interface{}** |  | [optional] 
**Organization** | Pointer to **string** | the name of organization | [optional] 
**Pipeline** | Pointer to **string** | pipeline | [optional] 
**Replayable** | Pointer to **bool** | Replayable or not | [optional] 
**Result** | Pointer to **string** | the result of pipeline run. e.g. SUCCESS | [optional] 
**RunSummary** | Pointer to **string** | pipeline run summary | [optional] 
**StartTime** | Pointer to **string** | the time of start | [optional] 
**State** | Pointer to **string** | run state. e.g. RUNNING | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewLatestRun

`func NewLatestRun() *LatestRun`

NewLatestRun instantiates a new LatestRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestRunWithDefaults

`func NewLatestRunWithDefaults() *LatestRun`

NewLatestRunWithDefaults instantiates a new LatestRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *LatestRun) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *LatestRun) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *LatestRun) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *LatestRun) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *LatestRun) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LatestRun) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LatestRun) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LatestRun) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *LatestRun) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LatestRun) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LatestRun) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *LatestRun) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetArtifactsZipFile

`func (o *LatestRun) GetArtifactsZipFile() string`

GetArtifactsZipFile returns the ArtifactsZipFile field if non-nil, zero value otherwise.

### GetArtifactsZipFileOk

`func (o *LatestRun) GetArtifactsZipFileOk() (*string, bool)`

GetArtifactsZipFileOk returns a tuple with the ArtifactsZipFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsZipFile

`func (o *LatestRun) SetArtifactsZipFile(v string)`

SetArtifactsZipFile sets ArtifactsZipFile field to given value.

### HasArtifactsZipFile

`func (o *LatestRun) HasArtifactsZipFile() bool`

HasArtifactsZipFile returns a boolean if a field has been set.

### GetCauseOfBlockage

`func (o *LatestRun) GetCauseOfBlockage() map[string]interface{}`

GetCauseOfBlockage returns the CauseOfBlockage field if non-nil, zero value otherwise.

### GetCauseOfBlockageOk

`func (o *LatestRun) GetCauseOfBlockageOk() (*map[string]interface{}, bool)`

GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfBlockage

`func (o *LatestRun) SetCauseOfBlockage(v map[string]interface{})`

SetCauseOfBlockage sets CauseOfBlockage field to given value.

### HasCauseOfBlockage

`func (o *LatestRun) HasCauseOfBlockage() bool`

HasCauseOfBlockage returns a boolean if a field has been set.

### GetCauses

`func (o *LatestRun) GetCauses() []LatestRunCauses`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *LatestRun) GetCausesOk() (*[]LatestRunCauses, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *LatestRun) SetCauses(v []LatestRunCauses)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *LatestRun) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetChangeSet

`func (o *LatestRun) GetChangeSet() []map[string]interface{}`

GetChangeSet returns the ChangeSet field if non-nil, zero value otherwise.

### GetChangeSetOk

`func (o *LatestRun) GetChangeSetOk() (*[]map[string]interface{}, bool)`

GetChangeSetOk returns a tuple with the ChangeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSet

`func (o *LatestRun) SetChangeSet(v []map[string]interface{})`

SetChangeSet sets ChangeSet field to given value.

### HasChangeSet

`func (o *LatestRun) HasChangeSet() bool`

HasChangeSet returns a boolean if a field has been set.

### GetDescription

`func (o *LatestRun) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LatestRun) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LatestRun) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LatestRun) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *LatestRun) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *LatestRun) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *LatestRun) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *LatestRun) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetEnQueueTime

`func (o *LatestRun) GetEnQueueTime() string`

GetEnQueueTime returns the EnQueueTime field if non-nil, zero value otherwise.

### GetEnQueueTimeOk

`func (o *LatestRun) GetEnQueueTimeOk() (*string, bool)`

GetEnQueueTimeOk returns a tuple with the EnQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnQueueTime

`func (o *LatestRun) SetEnQueueTime(v string)`

SetEnQueueTime sets EnQueueTime field to given value.

### HasEnQueueTime

`func (o *LatestRun) HasEnQueueTime() bool`

HasEnQueueTime returns a boolean if a field has been set.

### GetEndTime

`func (o *LatestRun) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LatestRun) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LatestRun) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LatestRun) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEstimatedDurationInMillis

`func (o *LatestRun) GetEstimatedDurationInMillis() int32`

GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field if non-nil, zero value otherwise.

### GetEstimatedDurationInMillisOk

`func (o *LatestRun) GetEstimatedDurationInMillisOk() (*int32, bool)`

GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMillis

`func (o *LatestRun) SetEstimatedDurationInMillis(v int32)`

SetEstimatedDurationInMillis sets EstimatedDurationInMillis field to given value.

### HasEstimatedDurationInMillis

`func (o *LatestRun) HasEstimatedDurationInMillis() bool`

HasEstimatedDurationInMillis returns a boolean if a field has been set.

### GetId

`func (o *LatestRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LatestRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LatestRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LatestRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LatestRun) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LatestRun) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LatestRun) SetName(v map[string]interface{})`

SetName sets Name field to given value.

### HasName

`func (o *LatestRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *LatestRun) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *LatestRun) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *LatestRun) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *LatestRun) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPipeline

`func (o *LatestRun) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *LatestRun) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *LatestRun) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *LatestRun) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetReplayable

`func (o *LatestRun) GetReplayable() bool`

GetReplayable returns the Replayable field if non-nil, zero value otherwise.

### GetReplayableOk

`func (o *LatestRun) GetReplayableOk() (*bool, bool)`

GetReplayableOk returns a tuple with the Replayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayable

`func (o *LatestRun) SetReplayable(v bool)`

SetReplayable sets Replayable field to given value.

### HasReplayable

`func (o *LatestRun) HasReplayable() bool`

HasReplayable returns a boolean if a field has been set.

### GetResult

`func (o *LatestRun) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LatestRun) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LatestRun) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *LatestRun) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRunSummary

`func (o *LatestRun) GetRunSummary() string`

GetRunSummary returns the RunSummary field if non-nil, zero value otherwise.

### GetRunSummaryOk

`func (o *LatestRun) GetRunSummaryOk() (*string, bool)`

GetRunSummaryOk returns a tuple with the RunSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunSummary

`func (o *LatestRun) SetRunSummary(v string)`

SetRunSummary sets RunSummary field to given value.

### HasRunSummary

`func (o *LatestRun) HasRunSummary() bool`

HasRunSummary returns a boolean if a field has been set.

### GetStartTime

`func (o *LatestRun) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LatestRun) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LatestRun) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *LatestRun) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *LatestRun) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LatestRun) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LatestRun) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LatestRun) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *LatestRun) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LatestRun) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LatestRun) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LatestRun) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


