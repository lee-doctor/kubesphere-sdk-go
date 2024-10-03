# DevopsStopPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions. | [optional] 
**ArtifactsZipFile** | Pointer to **map[string]interface{}** |  | [optional] 
**Branch** | Pointer to [**Branch**](Branch.md) |  | [optional] 
**CauseOfBlockage** | Pointer to **map[string]interface{}** |  | [optional] 
**Causes** | Pointer to [**[]DevopsStopPipelineCauses**](DevopsStopPipelineCauses.md) |  | [optional] 
**ChangeSet** | Pointer to **[]map[string]interface{}** | changeset information | [optional] 
**CommitId** | Pointer to **string** | commit id | [optional] 
**CommitUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **map[string]interface{}** |  | [optional] 
**DurationInMillis** | Pointer to **int32** | duration time in millis | [optional] 
**EnQueueTime** | Pointer to **string** | the time of enter the queue | [optional] 
**EndTime** | Pointer to **string** | the time of end | [optional] 
**EstimatedDurationInMillis** | Pointer to **int32** | estimated duration time in millis | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Name** | Pointer to **map[string]interface{}** |  | [optional] 
**Organization** | Pointer to **string** | the name of organization | [optional] 
**Pipeline** | Pointer to **string** | pipeline | [optional] 
**PullRequest** | Pointer to **map[string]interface{}** |  | [optional] 
**Replayable** | Pointer to **bool** | replayable or not | [optional] 
**Result** | Pointer to **string** | the result of pipeline run. e.g. SUCCESS | [optional] 
**RunSummary** | Pointer to **string** | pipeline run summary | [optional] 
**StartTime** | Pointer to **string** | the time of start | [optional] 
**State** | Pointer to **string** | run state. e.g. RUNNING | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewDevopsStopPipeline

`func NewDevopsStopPipeline() *DevopsStopPipeline`

NewDevopsStopPipeline instantiates a new DevopsStopPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsStopPipelineWithDefaults

`func NewDevopsStopPipelineWithDefaults() *DevopsStopPipeline`

NewDevopsStopPipelineWithDefaults instantiates a new DevopsStopPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsStopPipeline) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsStopPipeline) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsStopPipeline) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsStopPipeline) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsStopPipeline) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsStopPipeline) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsStopPipeline) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsStopPipeline) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsStopPipeline) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsStopPipeline) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsStopPipeline) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsStopPipeline) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetArtifactsZipFile

`func (o *DevopsStopPipeline) GetArtifactsZipFile() map[string]interface{}`

GetArtifactsZipFile returns the ArtifactsZipFile field if non-nil, zero value otherwise.

### GetArtifactsZipFileOk

`func (o *DevopsStopPipeline) GetArtifactsZipFileOk() (*map[string]interface{}, bool)`

GetArtifactsZipFileOk returns a tuple with the ArtifactsZipFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsZipFile

`func (o *DevopsStopPipeline) SetArtifactsZipFile(v map[string]interface{})`

SetArtifactsZipFile sets ArtifactsZipFile field to given value.

### HasArtifactsZipFile

`func (o *DevopsStopPipeline) HasArtifactsZipFile() bool`

HasArtifactsZipFile returns a boolean if a field has been set.

### GetBranch

`func (o *DevopsStopPipeline) GetBranch() Branch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DevopsStopPipeline) GetBranchOk() (*Branch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DevopsStopPipeline) SetBranch(v Branch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DevopsStopPipeline) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCauseOfBlockage

`func (o *DevopsStopPipeline) GetCauseOfBlockage() map[string]interface{}`

GetCauseOfBlockage returns the CauseOfBlockage field if non-nil, zero value otherwise.

### GetCauseOfBlockageOk

`func (o *DevopsStopPipeline) GetCauseOfBlockageOk() (*map[string]interface{}, bool)`

GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseOfBlockage

`func (o *DevopsStopPipeline) SetCauseOfBlockage(v map[string]interface{})`

SetCauseOfBlockage sets CauseOfBlockage field to given value.

### HasCauseOfBlockage

`func (o *DevopsStopPipeline) HasCauseOfBlockage() bool`

HasCauseOfBlockage returns a boolean if a field has been set.

### GetCauses

`func (o *DevopsStopPipeline) GetCauses() []DevopsStopPipelineCauses`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *DevopsStopPipeline) GetCausesOk() (*[]DevopsStopPipelineCauses, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *DevopsStopPipeline) SetCauses(v []DevopsStopPipelineCauses)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *DevopsStopPipeline) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetChangeSet

`func (o *DevopsStopPipeline) GetChangeSet() []map[string]interface{}`

GetChangeSet returns the ChangeSet field if non-nil, zero value otherwise.

### GetChangeSetOk

`func (o *DevopsStopPipeline) GetChangeSetOk() (*[]map[string]interface{}, bool)`

GetChangeSetOk returns a tuple with the ChangeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSet

`func (o *DevopsStopPipeline) SetChangeSet(v []map[string]interface{})`

SetChangeSet sets ChangeSet field to given value.

### HasChangeSet

`func (o *DevopsStopPipeline) HasChangeSet() bool`

HasChangeSet returns a boolean if a field has been set.

### GetCommitId

`func (o *DevopsStopPipeline) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *DevopsStopPipeline) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *DevopsStopPipeline) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *DevopsStopPipeline) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetCommitUrl

`func (o *DevopsStopPipeline) GetCommitUrl() map[string]interface{}`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *DevopsStopPipeline) GetCommitUrlOk() (*map[string]interface{}, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *DevopsStopPipeline) SetCommitUrl(v map[string]interface{})`

SetCommitUrl sets CommitUrl field to given value.

### HasCommitUrl

`func (o *DevopsStopPipeline) HasCommitUrl() bool`

HasCommitUrl returns a boolean if a field has been set.

### GetDescription

`func (o *DevopsStopPipeline) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DevopsStopPipeline) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DevopsStopPipeline) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DevopsStopPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurationInMillis

`func (o *DevopsStopPipeline) GetDurationInMillis() int32`

GetDurationInMillis returns the DurationInMillis field if non-nil, zero value otherwise.

### GetDurationInMillisOk

`func (o *DevopsStopPipeline) GetDurationInMillisOk() (*int32, bool)`

GetDurationInMillisOk returns a tuple with the DurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMillis

`func (o *DevopsStopPipeline) SetDurationInMillis(v int32)`

SetDurationInMillis sets DurationInMillis field to given value.

### HasDurationInMillis

`func (o *DevopsStopPipeline) HasDurationInMillis() bool`

HasDurationInMillis returns a boolean if a field has been set.

### GetEnQueueTime

`func (o *DevopsStopPipeline) GetEnQueueTime() string`

GetEnQueueTime returns the EnQueueTime field if non-nil, zero value otherwise.

### GetEnQueueTimeOk

`func (o *DevopsStopPipeline) GetEnQueueTimeOk() (*string, bool)`

GetEnQueueTimeOk returns a tuple with the EnQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnQueueTime

`func (o *DevopsStopPipeline) SetEnQueueTime(v string)`

SetEnQueueTime sets EnQueueTime field to given value.

### HasEnQueueTime

`func (o *DevopsStopPipeline) HasEnQueueTime() bool`

HasEnQueueTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DevopsStopPipeline) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DevopsStopPipeline) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DevopsStopPipeline) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DevopsStopPipeline) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEstimatedDurationInMillis

`func (o *DevopsStopPipeline) GetEstimatedDurationInMillis() int32`

GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field if non-nil, zero value otherwise.

### GetEstimatedDurationInMillisOk

`func (o *DevopsStopPipeline) GetEstimatedDurationInMillisOk() (*int32, bool)`

GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMillis

`func (o *DevopsStopPipeline) SetEstimatedDurationInMillis(v int32)`

SetEstimatedDurationInMillis sets EstimatedDurationInMillis field to given value.

### HasEstimatedDurationInMillis

`func (o *DevopsStopPipeline) HasEstimatedDurationInMillis() bool`

HasEstimatedDurationInMillis returns a boolean if a field has been set.

### GetId

`func (o *DevopsStopPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsStopPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsStopPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsStopPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DevopsStopPipeline) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsStopPipeline) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsStopPipeline) SetName(v map[string]interface{})`

SetName sets Name field to given value.

### HasName

`func (o *DevopsStopPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *DevopsStopPipeline) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DevopsStopPipeline) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DevopsStopPipeline) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DevopsStopPipeline) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPipeline

`func (o *DevopsStopPipeline) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DevopsStopPipeline) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DevopsStopPipeline) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *DevopsStopPipeline) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPullRequest

`func (o *DevopsStopPipeline) GetPullRequest() map[string]interface{}`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *DevopsStopPipeline) GetPullRequestOk() (*map[string]interface{}, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *DevopsStopPipeline) SetPullRequest(v map[string]interface{})`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *DevopsStopPipeline) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetReplayable

`func (o *DevopsStopPipeline) GetReplayable() bool`

GetReplayable returns the Replayable field if non-nil, zero value otherwise.

### GetReplayableOk

`func (o *DevopsStopPipeline) GetReplayableOk() (*bool, bool)`

GetReplayableOk returns a tuple with the Replayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayable

`func (o *DevopsStopPipeline) SetReplayable(v bool)`

SetReplayable sets Replayable field to given value.

### HasReplayable

`func (o *DevopsStopPipeline) HasReplayable() bool`

HasReplayable returns a boolean if a field has been set.

### GetResult

`func (o *DevopsStopPipeline) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevopsStopPipeline) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevopsStopPipeline) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevopsStopPipeline) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRunSummary

`func (o *DevopsStopPipeline) GetRunSummary() string`

GetRunSummary returns the RunSummary field if non-nil, zero value otherwise.

### GetRunSummaryOk

`func (o *DevopsStopPipeline) GetRunSummaryOk() (*string, bool)`

GetRunSummaryOk returns a tuple with the RunSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunSummary

`func (o *DevopsStopPipeline) SetRunSummary(v string)`

SetRunSummary sets RunSummary field to given value.

### HasRunSummary

`func (o *DevopsStopPipeline) HasRunSummary() bool`

HasRunSummary returns a boolean if a field has been set.

### GetStartTime

`func (o *DevopsStopPipeline) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DevopsStopPipeline) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DevopsStopPipeline) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DevopsStopPipeline) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *DevopsStopPipeline) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DevopsStopPipeline) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DevopsStopPipeline) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DevopsStopPipeline) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DevopsStopPipeline) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsStopPipeline) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsStopPipeline) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevopsStopPipeline) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


