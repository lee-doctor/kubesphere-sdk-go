# DevopsPipelineRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions | [optional] 
**ArtifactsZipFile** | Pointer to **map[string]interface{}** |  | [optional] 
**Branch** | Pointer to **map[string]interface{}** |  | [optional] 
**CauseOfBlockage** | Pointer to **map[string]interface{}** |  | [optional] 
**Causes** | Pointer to [**[]DevopsPipelineRunCauses**](DevopsPipelineRunCauses.md) |  | [optional] 
**ChangeSet** | Pointer to **[]map[string]interface{}** | changeset information | [optional] 
**CommitId** | Pointer to **map[string]interface{}** |  | [optional] 
**CommitUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**DurationInMillis** | Pointer to **int32** | duration time in millis | [optional] 
**EnQueueTime** | Pointer to **string** | the time of enter the queue | [optional] 
**EndTime** | Pointer to **string** | the time of end | [optional] 
**EstimatedDurationInMillis** | Pointer to **int32** | estimated duration time in millis | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Name** | Pointer to **map[string]interface{}** |  | [optional] 
**Organization** | Pointer to **string** | the name of organization | [optional] 
**Pipeline** | Pointer to **string** | the name of pipeline | [optional] 
**PullRequest** | Pointer to **map[string]interface{}** |  | [optional] 
**Replayable** | Pointer to **bool** | replayable or not | [optional] 
**Result** | Pointer to **string** | the result of pipeline run. e.g. SUCCESS | [optional] 
**RunSummary** | Pointer to **string** | pipeline run summary | [optional] 
**StartTime** | Pointer to **string** | the time of start | [optional] 
**State** | Pointer to **string** | run state. e.g. RUNNING | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewDevopsPipelineRun

`func NewDevopsPipelineRun() *DevopsPipelineRun`

NewDevopsPipelineRun instantiates a new DevopsPipelineRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsPipelineRunWithDefaults

`func NewDevopsPipelineRunWithDefaults() *DevopsPipelineRun`

NewDevopsPipelineRunWithDefaults instantiates a new DevopsPipelineRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsPipelineRun) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsPipelineRun) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsPipelineRun) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsPipelineRun) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsPipelineRun) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsPipelineRun) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsPipelineRun) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsPipelineRun) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsPipelineRun) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsPipelineRun) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsPipelineRun) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsPipelineRun) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetArtifactsZipFile

`func (o *DevopsPipelineRun) GetArtifactsZipFile() map[string]interface{}`

GetArtifactsZipFile returns the ArtifactsZipFile field if non-nil, zero value otherwise.

### GetArtifactsZipFileOk

`func (o *DevopsPipelineRun) GetArtifactsZipFileOk() (*map[string]interface{}, bool)`

GetArtifactsZipFileOk returns a tuple with the ArtifactsZipFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsZipFile

`func (o *DevopsPipelineRun) SetArtifactsZipFile(v map[string]interface{})`

SetArtifactsZipFile sets ArtifactsZipFile field to given value.

### HasArtifactsZipFile

`func (o *DevopsPipelineRun) HasArtifactsZipFile() bool`

HasArtifactsZipFile returns a boolean if a field has been set.

### GetBranch

`func (o *DevopsPipelineRun) GetBranch() map[string]interface{}`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DevopsPipelineRun) GetBranchOk() (*map[string]interface{}, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DevopsPipelineRun) SetBranch(v map[string]interface{})`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DevopsPipelineRun) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCauseOfBlockage

`func (o *DevopsPipelineRun) GetCauseOfBlockage() map[string]interface{}`

GetCauseOfBlockage returns the CauseOfBlockage field if non-nil, zero value otherwise.

### GetCauseOfBlockageOk

`func (o *DevopsPipelineRun) GetCauseOfBlockageOk() (*map[string]interface{}, bool)`

GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfBlockage

`func (o *DevopsPipelineRun) SetCauseOfBlockage(v map[string]interface{})`

SetCauseOfBlockage sets CauseOfBlockage field to given value.

### HasCauseOfBlockage

`func (o *DevopsPipelineRun) HasCauseOfBlockage() bool`

HasCauseOfBlockage returns a boolean if a field has been set.

### GetCauses

`func (o *DevopsPipelineRun) GetCauses() []DevopsPipelineRunCauses`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *DevopsPipelineRun) GetCausesOk() (*[]DevopsPipelineRunCauses, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *DevopsPipelineRun) SetCauses(v []DevopsPipelineRunCauses)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *DevopsPipelineRun) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetChangeSet

`func (o *DevopsPipelineRun) GetChangeSet() []map[string]interface{}`

GetChangeSet returns the ChangeSet field if non-nil, zero value otherwise.

### GetChangeSetOk

`func (o *DevopsPipelineRun) GetChangeSetOk() (*[]map[string]interface{}, bool)`

GetChangeSetOk returns a tuple with the ChangeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSet

`func (o *DevopsPipelineRun) SetChangeSet(v []map[string]interface{})`

SetChangeSet sets ChangeSet field to given value.

### HasChangeSet

`func (o *DevopsPipelineRun) HasChangeSet() bool`

HasChangeSet returns a boolean if a field has been set.

### GetCommitId

`func (o *DevopsPipelineRun) GetCommitId() map[string]interface{}`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *DevopsPipelineRun) GetCommitIdOk() (*map[string]interface{}, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *DevopsPipelineRun) SetCommitId(v map[string]interface{})`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *DevopsPipelineRun) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetCommitUrl

`func (o *DevopsPipelineRun) GetCommitUrl() map[string]interface{}`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *DevopsPipelineRun) GetCommitUrlOk() (*map[string]interface{}, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *DevopsPipelineRun) SetCommitUrl(v map[string]interface{})`

SetCommitUrl sets CommitUrl field to given value.

### HasCommitUrl

`func (o *DevopsPipelineRun) HasCommitUrl() bool`

HasCommitUrl returns a boolean if a field has been set.

### GetDescription

`func (o *DevopsPipelineRun) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DevopsPipelineRun) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DevopsPipelineRun) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DevopsPipelineRun) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsPipelineRun) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsPipelineRun) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsPipelineRun) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsPipelineRun) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetEnQueueTime

`func (o *DevopsPipelineRun) GetEnQueueTime() string`

GetEnQueueTime returns the EnQueueTime field if non-nil, zero value otherwise.

### GetEnQueueTimeOk

`func (o *DevopsPipelineRun) GetEnQueueTimeOk() (*string, bool)`

GetEnQueueTimeOk returns a tuple with the EnQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnQueueTime

`func (o *DevopsPipelineRun) SetEnQueueTime(v string)`

SetEnQueueTime sets EnQueueTime field to given value.

### HasEnQueueTime

`func (o *DevopsPipelineRun) HasEnQueueTime() bool`

HasEnQueueTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DevopsPipelineRun) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DevopsPipelineRun) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DevopsPipelineRun) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DevopsPipelineRun) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEstimatedDurationInMillis

`func (o *DevopsPipelineRun) GetEstimatedDurationInMillis() int32`

GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field if non-nil, zero value otherwise.

### GetEstimatedDurationInMillisOk

`func (o *DevopsPipelineRun) GetEstimatedDurationInMillisOk() (*int32, bool)`

GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMillis

`func (o *DevopsPipelineRun) SetEstimatedDurationInMillis(v int32)`

SetEstimatedDurationInMillis sets EstimatedDurationInMillis field to given value.

### HasEstimatedDurationInMillis

`func (o *DevopsPipelineRun) HasEstimatedDurationInMillis() bool`

HasEstimatedDurationInMillis returns a boolean if a field has been set.

### GetId

`func (o *DevopsPipelineRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsPipelineRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsPipelineRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsPipelineRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DevopsPipelineRun) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsPipelineRun) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsPipelineRun) SetName(v map[string]interface{})`

SetName sets Name field to given value.

### HasName

`func (o *DevopsPipelineRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *DevopsPipelineRun) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DevopsPipelineRun) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DevopsPipelineRun) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DevopsPipelineRun) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPipeline

`func (o *DevopsPipelineRun) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DevopsPipelineRun) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DevopsPipelineRun) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *DevopsPipelineRun) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPullRequest

`func (o *DevopsPipelineRun) GetPullRequest() map[string]interface{}`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *DevopsPipelineRun) GetPullRequestOk() (*map[string]interface{}, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *DevopsPipelineRun) SetPullRequest(v map[string]interface{})`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *DevopsPipelineRun) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetReplayable

`func (o *DevopsPipelineRun) GetReplayable() bool`

GetReplayable returns the Replayable field if non-nil, zero value otherwise.

### GetReplayableOk

`func (o *DevopsPipelineRun) GetReplayableOk() (*bool, bool)`

GetReplayableOk returns a tuple with the Replayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayable

`func (o *DevopsPipelineRun) SetReplayable(v bool)`

SetReplayable sets Replayable field to given value.

### HasReplayable

`func (o *DevopsPipelineRun) HasReplayable() bool`

HasReplayable returns a boolean if a field has been set.

### GetResult

`func (o *DevopsPipelineRun) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsPipelineRun) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsPipelineRun) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsPipelineRun) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRunSummary

`func (o *DevopsPipelineRun) GetRunSummary() string`

GetRunSummary returns the RunSummary field if non-nil, zero value otherwise.

### GetRunSummaryOk

`func (o *DevopsPipelineRun) GetRunSummaryOk() (*string, bool)`

GetRunSummaryOk returns a tuple with the RunSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunSummary

`func (o *DevopsPipelineRun) SetRunSummary(v string)`

SetRunSummary sets RunSummary field to given value.

### HasRunSummary

`func (o *DevopsPipelineRun) HasRunSummary() bool`

HasRunSummary returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsPipelineRun) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsPipelineRun) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsPipelineRun) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsPipelineRun) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsPipelineRun) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsPipelineRun) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsPipelineRun) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsPipelineRun) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DevopsPipelineRun) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsPipelineRun) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsPipelineRun) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsPipelineRun) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


