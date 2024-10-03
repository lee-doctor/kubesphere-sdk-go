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

// checks if the DevopsRunPipeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsRunPipeline{}

// DevopsRunPipeline struct for DevopsRunPipeline
type DevopsRunPipeline struct {
	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class *string `json:"_class,omitempty"`
	Links *Links `json:"_links,omitempty"`
	// the list of all actions
	Actions []map[string]interface{} `json:"actions,omitempty"`
	ArtifactsZipFile map[string]interface{} `json:"artifactsZipFile,omitempty"`
	// the cause of blockage
	CauseOfBlockage *string `json:"causeOfBlockage,omitempty"`
	Causes []DevopsRunPipelineCauses `json:"causes,omitempty"`
	// changeset information
	ChangeSet []map[string]interface{} `json:"changeSet,omitempty"`
	Description map[string]interface{} `json:"description,omitempty"`
	DurationInMillis map[string]interface{} `json:"durationInMillis,omitempty"`
	EnQueueTime map[string]interface{} `json:"enQueueTime,omitempty"`
	EndTime map[string]interface{} `json:"endTime,omitempty"`
	EstimatedDurationInMillis map[string]interface{} `json:"estimatedDurationInMillis,omitempty"`
	// id
	Id *string `json:"id,omitempty"`
	Name map[string]interface{} `json:"name,omitempty"`
	// the name of organization
	Organization *string `json:"organization,omitempty"`
	// pipeline
	Pipeline *string `json:"pipeline,omitempty"`
	// queue id
	QueueId *string `json:"queueId,omitempty"`
	// replayable or not
	Replayable *bool `json:"replayable,omitempty"`
	// the result of pipeline run. e.g. SUCCESS
	Result *string `json:"result,omitempty"`
	RunSummary map[string]interface{} `json:"runSummary,omitempty"`
	StartTime map[string]interface{} `json:"startTime,omitempty"`
	// run state. e.g. RUNNING
	State *string `json:"state,omitempty"`
	// type
	Type *string `json:"type,omitempty"`
}

// NewDevopsRunPipeline instantiates a new DevopsRunPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsRunPipeline() *DevopsRunPipeline {
	this := DevopsRunPipeline{}
	return &this
}

// NewDevopsRunPipelineWithDefaults instantiates a new DevopsRunPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsRunPipelineWithDefaults() *DevopsRunPipeline {
	this := DevopsRunPipeline{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *DevopsRunPipeline) SetClass(v string) {
	o.Class = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *DevopsRunPipeline) SetLinks(v Links) {
	o.Links = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetActions() []map[string]interface{} {
	if o == nil || IsNil(o.Actions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetActionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []map[string]interface{} and assigns it to the Actions field.
func (o *DevopsRunPipeline) SetActions(v []map[string]interface{}) {
	o.Actions = v
}

// GetArtifactsZipFile returns the ArtifactsZipFile field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetArtifactsZipFile() map[string]interface{} {
	if o == nil || IsNil(o.ArtifactsZipFile) {
		var ret map[string]interface{}
		return ret
	}
	return o.ArtifactsZipFile
}

// GetArtifactsZipFileOk returns a tuple with the ArtifactsZipFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetArtifactsZipFileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ArtifactsZipFile) {
		return map[string]interface{}{}, false
	}
	return o.ArtifactsZipFile, true
}

// HasArtifactsZipFile returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasArtifactsZipFile() bool {
	if o != nil && !IsNil(o.ArtifactsZipFile) {
		return true
	}

	return false
}

// SetArtifactsZipFile gets a reference to the given map[string]interface{} and assigns it to the ArtifactsZipFile field.
func (o *DevopsRunPipeline) SetArtifactsZipFile(v map[string]interface{}) {
	o.ArtifactsZipFile = v
}

// GetCauseOfBlockage returns the CauseOfBlockage field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetCauseOfBlockage() string {
	if o == nil || IsNil(o.CauseOfBlockage) {
		var ret string
		return ret
	}
	return *o.CauseOfBlockage
}

// GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetCauseOfBlockageOk() (*string, bool) {
	if o == nil || IsNil(o.CauseOfBlockage) {
		return nil, false
	}
	return o.CauseOfBlockage, true
}

// HasCauseOfBlockage returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasCauseOfBlockage() bool {
	if o != nil && !IsNil(o.CauseOfBlockage) {
		return true
	}

	return false
}

// SetCauseOfBlockage gets a reference to the given string and assigns it to the CauseOfBlockage field.
func (o *DevopsRunPipeline) SetCauseOfBlockage(v string) {
	o.CauseOfBlockage = &v
}

// GetCauses returns the Causes field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetCauses() []DevopsRunPipelineCauses {
	if o == nil || IsNil(o.Causes) {
		var ret []DevopsRunPipelineCauses
		return ret
	}
	return o.Causes
}

// GetCausesOk returns a tuple with the Causes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetCausesOk() ([]DevopsRunPipelineCauses, bool) {
	if o == nil || IsNil(o.Causes) {
		return nil, false
	}
	return o.Causes, true
}

// HasCauses returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasCauses() bool {
	if o != nil && !IsNil(o.Causes) {
		return true
	}

	return false
}

// SetCauses gets a reference to the given []DevopsRunPipelineCauses and assigns it to the Causes field.
func (o *DevopsRunPipeline) SetCauses(v []DevopsRunPipelineCauses) {
	o.Causes = v
}

// GetChangeSet returns the ChangeSet field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetChangeSet() []map[string]interface{} {
	if o == nil || IsNil(o.ChangeSet) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ChangeSet
}

// GetChangeSetOk returns a tuple with the ChangeSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetChangeSetOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ChangeSet) {
		return nil, false
	}
	return o.ChangeSet, true
}

// HasChangeSet returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasChangeSet() bool {
	if o != nil && !IsNil(o.ChangeSet) {
		return true
	}

	return false
}

// SetChangeSet gets a reference to the given []map[string]interface{} and assigns it to the ChangeSet field.
func (o *DevopsRunPipeline) SetChangeSet(v []map[string]interface{}) {
	o.ChangeSet = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetDescription() map[string]interface{} {
	if o == nil || IsNil(o.Description) {
		var ret map[string]interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetDescriptionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return map[string]interface{}{}, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given map[string]interface{} and assigns it to the Description field.
func (o *DevopsRunPipeline) SetDescription(v map[string]interface{}) {
	o.Description = v
}

// GetDurationInMillis returns the DurationInMillis field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetDurationInMillis() map[string]interface{} {
	if o == nil || IsNil(o.DurationInMillis) {
		var ret map[string]interface{}
		return ret
	}
	return o.DurationInMillis
}

// GetDurationInMillisOk returns a tuple with the DurationInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetDurationInMillisOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DurationInMillis) {
		return map[string]interface{}{}, false
	}
	return o.DurationInMillis, true
}

// HasDurationInMillis returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasDurationInMillis() bool {
	if o != nil && !IsNil(o.DurationInMillis) {
		return true
	}

	return false
}

// SetDurationInMillis gets a reference to the given map[string]interface{} and assigns it to the DurationInMillis field.
func (o *DevopsRunPipeline) SetDurationInMillis(v map[string]interface{}) {
	o.DurationInMillis = v
}

// GetEnQueueTime returns the EnQueueTime field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetEnQueueTime() map[string]interface{} {
	if o == nil || IsNil(o.EnQueueTime) {
		var ret map[string]interface{}
		return ret
	}
	return o.EnQueueTime
}

// GetEnQueueTimeOk returns a tuple with the EnQueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetEnQueueTimeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EnQueueTime) {
		return map[string]interface{}{}, false
	}
	return o.EnQueueTime, true
}

// HasEnQueueTime returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasEnQueueTime() bool {
	if o != nil && !IsNil(o.EnQueueTime) {
		return true
	}

	return false
}

// SetEnQueueTime gets a reference to the given map[string]interface{} and assigns it to the EnQueueTime field.
func (o *DevopsRunPipeline) SetEnQueueTime(v map[string]interface{}) {
	o.EnQueueTime = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetEndTime() map[string]interface{} {
	if o == nil || IsNil(o.EndTime) {
		var ret map[string]interface{}
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetEndTimeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EndTime) {
		return map[string]interface{}{}, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given map[string]interface{} and assigns it to the EndTime field.
func (o *DevopsRunPipeline) SetEndTime(v map[string]interface{}) {
	o.EndTime = v
}

// GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetEstimatedDurationInMillis() map[string]interface{} {
	if o == nil || IsNil(o.EstimatedDurationInMillis) {
		var ret map[string]interface{}
		return ret
	}
	return o.EstimatedDurationInMillis
}

// GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetEstimatedDurationInMillisOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.EstimatedDurationInMillis) {
		return map[string]interface{}{}, false
	}
	return o.EstimatedDurationInMillis, true
}

// HasEstimatedDurationInMillis returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasEstimatedDurationInMillis() bool {
	if o != nil && !IsNil(o.EstimatedDurationInMillis) {
		return true
	}

	return false
}

// SetEstimatedDurationInMillis gets a reference to the given map[string]interface{} and assigns it to the EstimatedDurationInMillis field.
func (o *DevopsRunPipeline) SetEstimatedDurationInMillis(v map[string]interface{}) {
	o.EstimatedDurationInMillis = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevopsRunPipeline) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetName() map[string]interface{} {
	if o == nil || IsNil(o.Name) {
		var ret map[string]interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetNameOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return map[string]interface{}{}, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given map[string]interface{} and assigns it to the Name field.
func (o *DevopsRunPipeline) SetName(v map[string]interface{}) {
	o.Name = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *DevopsRunPipeline) SetOrganization(v string) {
	o.Organization = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetPipeline() string {
	if o == nil || IsNil(o.Pipeline) {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetPipelineOk() (*string, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *DevopsRunPipeline) SetPipeline(v string) {
	o.Pipeline = &v
}

// GetQueueId returns the QueueId field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetQueueId() string {
	if o == nil || IsNil(o.QueueId) {
		var ret string
		return ret
	}
	return *o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetQueueIdOk() (*string, bool) {
	if o == nil || IsNil(o.QueueId) {
		return nil, false
	}
	return o.QueueId, true
}

// HasQueueId returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasQueueId() bool {
	if o != nil && !IsNil(o.QueueId) {
		return true
	}

	return false
}

// SetQueueId gets a reference to the given string and assigns it to the QueueId field.
func (o *DevopsRunPipeline) SetQueueId(v string) {
	o.QueueId = &v
}

// GetReplayable returns the Replayable field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetReplayable() bool {
	if o == nil || IsNil(o.Replayable) {
		var ret bool
		return ret
	}
	return *o.Replayable
}

// GetReplayableOk returns a tuple with the Replayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetReplayableOk() (*bool, bool) {
	if o == nil || IsNil(o.Replayable) {
		return nil, false
	}
	return o.Replayable, true
}

// HasReplayable returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasReplayable() bool {
	if o != nil && !IsNil(o.Replayable) {
		return true
	}

	return false
}

// SetReplayable gets a reference to the given bool and assigns it to the Replayable field.
func (o *DevopsRunPipeline) SetReplayable(v bool) {
	o.Replayable = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DevopsRunPipeline) SetResult(v string) {
	o.Result = &v
}

// GetRunSummary returns the RunSummary field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetRunSummary() map[string]interface{} {
	if o == nil || IsNil(o.RunSummary) {
		var ret map[string]interface{}
		return ret
	}
	return o.RunSummary
}

// GetRunSummaryOk returns a tuple with the RunSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetRunSummaryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RunSummary) {
		return map[string]interface{}{}, false
	}
	return o.RunSummary, true
}

// HasRunSummary returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasRunSummary() bool {
	if o != nil && !IsNil(o.RunSummary) {
		return true
	}

	return false
}

// SetRunSummary gets a reference to the given map[string]interface{} and assigns it to the RunSummary field.
func (o *DevopsRunPipeline) SetRunSummary(v map[string]interface{}) {
	o.RunSummary = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetStartTime() map[string]interface{} {
	if o == nil || IsNil(o.StartTime) {
		var ret map[string]interface{}
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetStartTimeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StartTime) {
		return map[string]interface{}{}, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given map[string]interface{} and assigns it to the StartTime field.
func (o *DevopsRunPipeline) SetStartTime(v map[string]interface{}) {
	o.StartTime = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DevopsRunPipeline) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DevopsRunPipeline) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsRunPipeline) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DevopsRunPipeline) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DevopsRunPipeline) SetType(v string) {
	o.Type = &v
}

func (o DevopsRunPipeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsRunPipeline) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CauseOfBlockage) {
		toSerialize["causeOfBlockage"] = o.CauseOfBlockage
	}
	if !IsNil(o.Causes) {
		toSerialize["causes"] = o.Causes
	}
	if !IsNil(o.ChangeSet) {
		toSerialize["changeSet"] = o.ChangeSet
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
	if !IsNil(o.QueueId) {
		toSerialize["queueId"] = o.QueueId
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

type NullableDevopsRunPipeline struct {
	value *DevopsRunPipeline
	isSet bool
}

func (v NullableDevopsRunPipeline) Get() *DevopsRunPipeline {
	return v.value
}

func (v *NullableDevopsRunPipeline) Set(val *DevopsRunPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsRunPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsRunPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsRunPipeline(val *DevopsRunPipeline) *NullableDevopsRunPipeline {
	return &NullableDevopsRunPipeline{value: val, isSet: true}
}

func (v NullableDevopsRunPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsRunPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

