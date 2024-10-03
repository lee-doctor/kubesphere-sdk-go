/*
KubeSphere

KubeSphere OpenAPI

API version: v3.1.0
Contact: info@kubesphere.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DevopsStopPipeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsStopPipeline{}

// DevopsStopPipeline struct for DevopsStopPipeline
type DevopsStopPipeline struct {
	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class *string `json:"_class,omitempty"`
	Links *Links `json:"_links,omitempty"`
	// the list of all actions.
	Actions []map[string]interface{} `json:"actions,omitempty"`
	ArtifactsZipFile map[string]interface{} `json:"artifactsZipFile,omitempty"`
	Branch *Branch `json:"branch,omitempty"`
	CauseOfBlockage map[string]interface{} `json:"causeOfBlockage,omitempty"`
	Causes []DevopsStopPipelineCauses `json:"causes,omitempty"`
	// changeset information
	ChangeSet []map[string]interface{} `json:"changeSet,omitempty"`
	// commit id
	CommitId *string `json:"commitId,omitempty"`
	CommitUrl map[string]interface{} `json:"commitUrl,omitempty"`
	Description map[string]interface{} `json:"description,omitempty"`
	// duration time in millis
	DurationInMillis *int32 `json:"durationInMillis,omitempty"`
	// the time of enter the queue
	EnQueueTime *string `json:"enQueueTime,omitempty"`
	// the time of end
	EndTime *string `json:"endTime,omitempty"`
	// estimated duration time in millis
	EstimatedDurationInMillis *int32 `json:"estimatedDurationInMillis,omitempty"`
	// id
	Id *string `json:"id,omitempty"`
	Name map[string]interface{} `json:"name,omitempty"`
	// the name of organization
	Organization *string `json:"organization,omitempty"`
	// pipeline
	Pipeline *string `json:"pipeline,omitempty"`
	PullRequest map[string]interface{} `json:"pullRequest,omitempty"`
	// replayable or not
	Replayable *bool `json:"replayable,omitempty"`
	// the result of pipeline run. e.g. SUCCESS
	Result *string `json:"result,omitempty"`
	// pipeline run summary
	RunSummary *string `json:"runSummary,omitempty"`
	// the time of start
	StartTime *string `json:"startTime,omitempty"`
	// run state. e.g. RUNNING
	State *string `json:"state,omitempty"`
	// type
	Type *string `json:"type,omitempty"`
}

// NewDevopsStopPipeline instantiates a new DevopsStopPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsStopPipeline() *DevopsStopPipeline {
	this := DevopsStopPipeline{}
	return &this
}

// NewDevopsStopPipelineWithDefaults instantiates a new DevopsStopPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsStopPipelineWithDefaults() *DevopsStopPipeline {
	this := DevopsStopPipeline{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *DevopsStopPipeline) SetClass(v string) {
	o.Class = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *DevopsStopPipeline) SetLinks(v Links) {
	o.Links = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetActions() []map[string]interface{} {
	if o == nil || IsNil(o.Actions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetActionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []map[string]interface{} and assigns it to the Actions field.
func (o *DevopsStopPipeline) SetActions(v []map[string]interface{}) {
	o.Actions = v
}

// GetArtifactsZipFile returns the ArtifactsZipFile field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetArtifactsZipFile() map[string]interface{} {
	if o == nil || IsNil(o.ArtifactsZipFile) {
		var ret map[string]interface{}
		return ret
	}
	return o.ArtifactsZipFile
}

// GetArtifactsZipFileOk returns a tuple with the ArtifactsZipFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetArtifactsZipFileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ArtifactsZipFile) {
		return map[string]interface{}{}, false
	}
	return o.ArtifactsZipFile, true
}

// HasArtifactsZipFile returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasArtifactsZipFile() bool {
	if o != nil && !IsNil(o.ArtifactsZipFile) {
		return true
	}

	return false
}

// SetArtifactsZipFile gets a reference to the given map[string]interface{} and assigns it to the ArtifactsZipFile field.
func (o *DevopsStopPipeline) SetArtifactsZipFile(v map[string]interface{}) {
	o.ArtifactsZipFile = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetBranch() Branch {
	if o == nil || IsNil(o.Branch) {
		var ret Branch
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetBranchOk() (*Branch, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given Branch and assigns it to the Branch field.
func (o *DevopsStopPipeline) SetBranch(v Branch) {
	o.Branch = &v
}

// GetCauseOfBlockage returns the CauseOfBlockage field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetCauseOfBlockage() map[string]interface{} {
	if o == nil || IsNil(o.CauseOfBlockage) {
		var ret map[string]interface{}
		return ret
	}
	return o.CauseOfBlockage
}

// GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetCauseOfBlockageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CauseOfBlockage) {
		return map[string]interface{}{}, false
	}
	return o.CauseOfBlockage, true
}

// HasCauseOfBlockage returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasCauseOfBlockage() bool {
	if o != nil && !IsNil(o.CauseOfBlockage) {
		return true
	}

	return false
}

// SetCauseOfBlockage gets a reference to the given map[string]interface{} and assigns it to the CauseOfBlockage field.
func (o *DevopsStopPipeline) SetCauseOfBlockage(v map[string]interface{}) {
	o.CauseOfBlockage = v
}

// GetCauses returns the Causes field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetCauses() []DevopsStopPipelineCauses {
	if o == nil || IsNil(o.Causes) {
		var ret []DevopsStopPipelineCauses
		return ret
	}
	return o.Causes
}

// GetCausesOk returns a tuple with the Causes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetCausesOk() ([]DevopsStopPipelineCauses, bool) {
	if o == nil || IsNil(o.Causes) {
		return nil, false
	}
	return o.Causes, true
}

// HasCauses returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasCauses() bool {
	if o != nil && !IsNil(o.Causes) {
		return true
	}

	return false
}

// SetCauses gets a reference to the given []DevopsStopPipelineCauses and assigns it to the Causes field.
func (o *DevopsStopPipeline) SetCauses(v []DevopsStopPipelineCauses) {
	o.Causes = v
}

// GetChangeSet returns the ChangeSet field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetChangeSet() []map[string]interface{} {
	if o == nil || IsNil(o.ChangeSet) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ChangeSet
}

// GetChangeSetOk returns a tuple with the ChangeSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetChangeSetOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ChangeSet) {
		return nil, false
	}
	return o.ChangeSet, true
}

// HasChangeSet returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasChangeSet() bool {
	if o != nil && !IsNil(o.ChangeSet) {
		return true
	}

	return false
}

// SetChangeSet gets a reference to the given []map[string]interface{} and assigns it to the ChangeSet field.
func (o *DevopsStopPipeline) SetChangeSet(v []map[string]interface{}) {
	o.ChangeSet = v
}

// GetCommitId returns the CommitId field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetCommitId() string {
	if o == nil || IsNil(o.CommitId) {
		var ret string
		return ret
	}
	return *o.CommitId
}

// GetCommitIdOk returns a tuple with the CommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetCommitIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommitId) {
		return nil, false
	}
	return o.CommitId, true
}

// HasCommitId returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasCommitId() bool {
	if o != nil && !IsNil(o.CommitId) {
		return true
	}

	return false
}

// SetCommitId gets a reference to the given string and assigns it to the CommitId field.
func (o *DevopsStopPipeline) SetCommitId(v string) {
	o.CommitId = &v
}

// GetCommitUrl returns the CommitUrl field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetCommitUrl() map[string]interface{} {
	if o == nil || IsNil(o.CommitUrl) {
		var ret map[string]interface{}
		return ret
	}
	return o.CommitUrl
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetCommitUrlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CommitUrl) {
		return map[string]interface{}{}, false
	}
	return o.CommitUrl, true
}

// HasCommitUrl returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasCommitUrl() bool {
	if o != nil && !IsNil(o.CommitUrl) {
		return true
	}

	return false
}

// SetCommitUrl gets a reference to the given map[string]interface{} and assigns it to the CommitUrl field.
func (o *DevopsStopPipeline) SetCommitUrl(v map[string]interface{}) {
	o.CommitUrl = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetDescription() map[string]interface{} {
	if o == nil || IsNil(o.Description) {
		var ret map[string]interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetDescriptionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return map[string]interface{}{}, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given map[string]interface{} and assigns it to the Description field.
func (o *DevopsStopPipeline) SetDescription(v map[string]interface{}) {
	o.Description = v
}

// GetDurationInMillis returns the DurationInMillis field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetDurationInMillis() int32 {
	if o == nil || IsNil(o.DurationInMillis) {
		var ret int32
		return ret
	}
	return *o.DurationInMillis
}

// GetDurationInMillisOk returns a tuple with the DurationInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetDurationInMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInMillis) {
		return nil, false
	}
	return o.DurationInMillis, true
}

// HasDurationInMillis returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasDurationInMillis() bool {
	if o != nil && !IsNil(o.DurationInMillis) {
		return true
	}

	return false
}

// SetDurationInMillis gets a reference to the given int32 and assigns it to the DurationInMillis field.
func (o *DevopsStopPipeline) SetDurationInMillis(v int32) {
	o.DurationInMillis = &v
}

// GetEnQueueTime returns the EnQueueTime field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetEnQueueTime() string {
	if o == nil || IsNil(o.EnQueueTime) {
		var ret string
		return ret
	}
	return *o.EnQueueTime
}

// GetEnQueueTimeOk returns a tuple with the EnQueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetEnQueueTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EnQueueTime) {
		return nil, false
	}
	return o.EnQueueTime, true
}

// HasEnQueueTime returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasEnQueueTime() bool {
	if o != nil && !IsNil(o.EnQueueTime) {
		return true
	}

	return false
}

// SetEnQueueTime gets a reference to the given string and assigns it to the EnQueueTime field.
func (o *DevopsStopPipeline) SetEnQueueTime(v string) {
	o.EnQueueTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *DevopsStopPipeline) SetEndTime(v string) {
	o.EndTime = &v
}

// GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetEstimatedDurationInMillis() int32 {
	if o == nil || IsNil(o.EstimatedDurationInMillis) {
		var ret int32
		return ret
	}
	return *o.EstimatedDurationInMillis
}

// GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetEstimatedDurationInMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedDurationInMillis) {
		return nil, false
	}
	return o.EstimatedDurationInMillis, true
}

// HasEstimatedDurationInMillis returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasEstimatedDurationInMillis() bool {
	if o != nil && !IsNil(o.EstimatedDurationInMillis) {
		return true
	}

	return false
}

// SetEstimatedDurationInMillis gets a reference to the given int32 and assigns it to the EstimatedDurationInMillis field.
func (o *DevopsStopPipeline) SetEstimatedDurationInMillis(v int32) {
	o.EstimatedDurationInMillis = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevopsStopPipeline) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetName() map[string]interface{} {
	if o == nil || IsNil(o.Name) {
		var ret map[string]interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetNameOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return map[string]interface{}{}, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given map[string]interface{} and assigns it to the Name field.
func (o *DevopsStopPipeline) SetName(v map[string]interface{}) {
	o.Name = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *DevopsStopPipeline) SetOrganization(v string) {
	o.Organization = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetPipeline() string {
	if o == nil || IsNil(o.Pipeline) {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetPipelineOk() (*string, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *DevopsStopPipeline) SetPipeline(v string) {
	o.Pipeline = &v
}

// GetPullRequest returns the PullRequest field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetPullRequest() map[string]interface{} {
	if o == nil || IsNil(o.PullRequest) {
		var ret map[string]interface{}
		return ret
	}
	return o.PullRequest
}

// GetPullRequestOk returns a tuple with the PullRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetPullRequestOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PullRequest) {
		return map[string]interface{}{}, false
	}
	return o.PullRequest, true
}

// HasPullRequest returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasPullRequest() bool {
	if o != nil && !IsNil(o.PullRequest) {
		return true
	}

	return false
}

// SetPullRequest gets a reference to the given map[string]interface{} and assigns it to the PullRequest field.
func (o *DevopsStopPipeline) SetPullRequest(v map[string]interface{}) {
	o.PullRequest = v
}

// GetReplayable returns the Replayable field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetReplayable() bool {
	if o == nil || IsNil(o.Replayable) {
		var ret bool
		return ret
	}
	return *o.Replayable
}

// GetReplayableOk returns a tuple with the Replayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetReplayableOk() (*bool, bool) {
	if o == nil || IsNil(o.Replayable) {
		return nil, false
	}
	return o.Replayable, true
}

// HasReplayable returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasReplayable() bool {
	if o != nil && !IsNil(o.Replayable) {
		return true
	}

	return false
}

// SetReplayable gets a reference to the given bool and assigns it to the Replayable field.
func (o *DevopsStopPipeline) SetReplayable(v bool) {
	o.Replayable = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DevopsStopPipeline) SetResult(v string) {
	o.Result = &v
}

// GetRunSummary returns the RunSummary field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetRunSummary() string {
	if o == nil || IsNil(o.RunSummary) {
		var ret string
		return ret
	}
	return *o.RunSummary
}

// GetRunSummaryOk returns a tuple with the RunSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetRunSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.RunSummary) {
		return nil, false
	}
	return o.RunSummary, true
}

// HasRunSummary returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasRunSummary() bool {
	if o != nil && !IsNil(o.RunSummary) {
		return true
	}

	return false
}

// SetRunSummary gets a reference to the given string and assigns it to the RunSummary field.
func (o *DevopsStopPipeline) SetRunSummary(v string) {
	o.RunSummary = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *DevopsStopPipeline) SetStartTime(v string) {
	o.StartTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DevopsStopPipeline) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DevopsStopPipeline) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipeline) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DevopsStopPipeline) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DevopsStopPipeline) SetType(v string) {
	o.Type = &v
}

func (o DevopsStopPipeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsStopPipeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["_class"] = o.Class
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.ArtifactsZipFile) {
		toSerialize["artifactsZipFile"] = o.ArtifactsZipFile
	}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.CauseOfBlockage) {
		toSerialize["causeOfBlockage"] = o.CauseOfBlockage
	}
	if !IsNil(o.Causes) {
		toSerialize["causes"] = o.Causes
	}
	if !IsNil(o.ChangeSet) {
		toSerialize["changeSet"] = o.ChangeSet
	}
	if !IsNil(o.CommitId) {
		toSerialize["commitId"] = o.CommitId
	}
	if !IsNil(o.CommitUrl) {
		toSerialize["commitUrl"] = o.CommitUrl
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DurationInMillis) {
		toSerialize["durationInMillis"] = o.DurationInMillis
	}
	if !IsNil(o.EnQueueTime) {
		toSerialize["enQueueTime"] = o.EnQueueTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.EstimatedDurationInMillis) {
		toSerialize["estimatedDurationInMillis"] = o.EstimatedDurationInMillis
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Pipeline) {
		toSerialize["pipeline"] = o.Pipeline
	}
	if !IsNil(o.PullRequest) {
		toSerialize["pullRequest"] = o.PullRequest
	}
	if !IsNil(o.Replayable) {
		toSerialize["replayable"] = o.Replayable
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.RunSummary) {
		toSerialize["runSummary"] = o.RunSummary
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDevopsStopPipeline struct {
	value *DevopsStopPipeline
	isSet bool
}

func (v NullableDevopsStopPipeline) Get() *DevopsStopPipeline {
	return v.value
}

func (v *NullableDevopsStopPipeline) Set(val *DevopsStopPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsStopPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsStopPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsStopPipeline(val *DevopsStopPipeline) *NullableDevopsStopPipeline {
	return &NullableDevopsStopPipeline{value: val, isSet: true}
}

func (v NullableDevopsStopPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsStopPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

