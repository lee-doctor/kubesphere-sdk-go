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
	"bytes"
	"fmt"
)

// checks if the DevopsPipeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsPipeline{}

// DevopsPipeline struct for DevopsPipeline
type DevopsPipeline struct {
	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class *string `json:"_class,omitempty"`
	Links *Links `json:"_links,omitempty"`
	// the list of all actions.
	Actions []map[string]interface{} `json:"actions,omitempty"`
	// branch names
	BranchNames []string `json:"branchNames,omitempty"`
	Disabled map[string]interface{} `json:"disabled,omitempty"`
	// display name
	DisplayName *string `json:"displayName,omitempty"`
	// estimated duration time, unit is millis
	EstimatedDurationInMillis *int32 `json:"estimatedDurationInMillis,omitempty"`
	// full display name
	FullDisplayName *string `json:"fullDisplayName,omitempty"`
	// full name
	FullName *string `json:"fullName,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// number of failing branches
	NumberOfFailingBranches *int32 `json:"numberOfFailingBranches,omitempty"`
	// number of failing pull requests
	NumberOfFailingPullRequests *int32 `json:"numberOfFailingPullRequests,omitempty"`
	// number of folders
	NumberOfFolders *int32 `json:"numberOfFolders,omitempty"`
	// number of pipelines
	NumberOfPipelines *int32 `json:"numberOfPipelines,omitempty"`
	// number of successful pull requests
	NumberOfSuccessfulBranches *int32 `json:"numberOfSuccessfulBranches,omitempty"`
	// number of successful pull requests
	NumberOfSuccessfulPullRequests *int32 `json:"numberOfSuccessfulPullRequests,omitempty"`
	// the name of organization
	Organization *string `json:"organization,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Permissions *Permissions `json:"permissions,omitempty"`
	// pipeline folder names
	PipelineFolderNames []map[string]interface{} `json:"pipelineFolderNames,omitempty"`
	ScmSource *ScmSource `json:"scmSource,omitempty"`
	// total number of branches
	TotalNumberOfBranches *int32 `json:"totalNumberOfBranches,omitempty"`
	// total number of pull requests
	TotalNumberOfPullRequests *int32 `json:"totalNumberOfPullRequests,omitempty"`
	// the score to description the result of pipeline activity
	WeatherScore int32 `json:"weatherScore"`
}

type _DevopsPipeline DevopsPipeline

// NewDevopsPipeline instantiates a new DevopsPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsPipeline(weatherScore int32) *DevopsPipeline {
	this := DevopsPipeline{}
	this.WeatherScore = weatherScore
	return &this
}

// NewDevopsPipelineWithDefaults instantiates a new DevopsPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsPipelineWithDefaults() *DevopsPipeline {
	this := DevopsPipeline{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *DevopsPipeline) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *DevopsPipeline) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *DevopsPipeline) SetClass(v string) {
	o.Class = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DevopsPipeline) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DevopsPipeline) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *DevopsPipeline) SetLinks(v Links) {
	o.Links = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *DevopsPipeline) GetActions() []map[string]interface{} {
	if o == nil || IsNil(o.Actions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetActionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DevopsPipeline) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []map[string]interface{} and assigns it to the Actions field.
func (o *DevopsPipeline) SetActions(v []map[string]interface{}) {
	o.Actions = v
}

// GetBranchNames returns the BranchNames field value if set, zero value otherwise.
func (o *DevopsPipeline) GetBranchNames() []string {
	if o == nil || IsNil(o.BranchNames) {
		var ret []string
		return ret
	}
	return o.BranchNames
}

// GetBranchNamesOk returns a tuple with the BranchNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetBranchNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.BranchNames) {
		return nil, false
	}
	return o.BranchNames, true
}

// HasBranchNames returns a boolean if a field has been set.
func (o *DevopsPipeline) HasBranchNames() bool {
	if o != nil && !IsNil(o.BranchNames) {
		return true
	}

	return false
}

// SetBranchNames gets a reference to the given []string and assigns it to the BranchNames field.
func (o *DevopsPipeline) SetBranchNames(v []string) {
	o.BranchNames = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *DevopsPipeline) GetDisabled() map[string]interface{} {
	if o == nil || IsNil(o.Disabled) {
		var ret map[string]interface{}
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetDisabledOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Disabled) {
		return map[string]interface{}{}, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *DevopsPipeline) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given map[string]interface{} and assigns it to the Disabled field.
func (o *DevopsPipeline) SetDisabled(v map[string]interface{}) {
	o.Disabled = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DevopsPipeline) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DevopsPipeline) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DevopsPipeline) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field value if set, zero value otherwise.
func (o *DevopsPipeline) GetEstimatedDurationInMillis() int32 {
	if o == nil || IsNil(o.EstimatedDurationInMillis) {
		var ret int32
		return ret
	}
	return *o.EstimatedDurationInMillis
}

// GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetEstimatedDurationInMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedDurationInMillis) {
		return nil, false
	}
	return o.EstimatedDurationInMillis, true
}

// HasEstimatedDurationInMillis returns a boolean if a field has been set.
func (o *DevopsPipeline) HasEstimatedDurationInMillis() bool {
	if o != nil && !IsNil(o.EstimatedDurationInMillis) {
		return true
	}

	return false
}

// SetEstimatedDurationInMillis gets a reference to the given int32 and assigns it to the EstimatedDurationInMillis field.
func (o *DevopsPipeline) SetEstimatedDurationInMillis(v int32) {
	o.EstimatedDurationInMillis = &v
}

// GetFullDisplayName returns the FullDisplayName field value if set, zero value otherwise.
func (o *DevopsPipeline) GetFullDisplayName() string {
	if o == nil || IsNil(o.FullDisplayName) {
		var ret string
		return ret
	}
	return *o.FullDisplayName
}

// GetFullDisplayNameOk returns a tuple with the FullDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetFullDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullDisplayName) {
		return nil, false
	}
	return o.FullDisplayName, true
}

// HasFullDisplayName returns a boolean if a field has been set.
func (o *DevopsPipeline) HasFullDisplayName() bool {
	if o != nil && !IsNil(o.FullDisplayName) {
		return true
	}

	return false
}

// SetFullDisplayName gets a reference to the given string and assigns it to the FullDisplayName field.
func (o *DevopsPipeline) SetFullDisplayName(v string) {
	o.FullDisplayName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *DevopsPipeline) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *DevopsPipeline) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *DevopsPipeline) SetFullName(v string) {
	o.FullName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevopsPipeline) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevopsPipeline) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevopsPipeline) SetName(v string) {
	o.Name = &v
}

// GetNumberOfFailingBranches returns the NumberOfFailingBranches field value if set, zero value otherwise.
func (o *DevopsPipeline) GetNumberOfFailingBranches() int32 {
	if o == nil || IsNil(o.NumberOfFailingBranches) {
		var ret int32
		return ret
	}
	return *o.NumberOfFailingBranches
}

// GetNumberOfFailingBranchesOk returns a tuple with the NumberOfFailingBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetNumberOfFailingBranchesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfFailingBranches) {
		return nil, false
	}
	return o.NumberOfFailingBranches, true
}

// HasNumberOfFailingBranches returns a boolean if a field has been set.
func (o *DevopsPipeline) HasNumberOfFailingBranches() bool {
	if o != nil && !IsNil(o.NumberOfFailingBranches) {
		return true
	}

	return false
}

// SetNumberOfFailingBranches gets a reference to the given int32 and assigns it to the NumberOfFailingBranches field.
func (o *DevopsPipeline) SetNumberOfFailingBranches(v int32) {
	o.NumberOfFailingBranches = &v
}

// GetNumberOfFailingPullRequests returns the NumberOfFailingPullRequests field value if set, zero value otherwise.
func (o *DevopsPipeline) GetNumberOfFailingPullRequests() int32 {
	if o == nil || IsNil(o.NumberOfFailingPullRequests) {
		var ret int32
		return ret
	}
	return *o.NumberOfFailingPullRequests
}

// GetNumberOfFailingPullRequestsOk returns a tuple with the NumberOfFailingPullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetNumberOfFailingPullRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfFailingPullRequests) {
		return nil, false
	}
	return o.NumberOfFailingPullRequests, true
}

// HasNumberOfFailingPullRequests returns a boolean if a field has been set.
func (o *DevopsPipeline) HasNumberOfFailingPullRequests() bool {
	if o != nil && !IsNil(o.NumberOfFailingPullRequests) {
		return true
	}

	return false
}

// SetNumberOfFailingPullRequests gets a reference to the given int32 and assigns it to the NumberOfFailingPullRequests field.
func (o *DevopsPipeline) SetNumberOfFailingPullRequests(v int32) {
	o.NumberOfFailingPullRequests = &v
}

// GetNumberOfFolders returns the NumberOfFolders field value if set, zero value otherwise.
func (o *DevopsPipeline) GetNumberOfFolders() int32 {
	if o == nil || IsNil(o.NumberOfFolders) {
		var ret int32
		return ret
	}
	return *o.NumberOfFolders
}

// GetNumberOfFoldersOk returns a tuple with the NumberOfFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetNumberOfFoldersOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfFolders) {
		return nil, false
	}
	return o.NumberOfFolders, true
}

// HasNumberOfFolders returns a boolean if a field has been set.
func (o *DevopsPipeline) HasNumberOfFolders() bool {
	if o != nil && !IsNil(o.NumberOfFolders) {
		return true
	}

	return false
}

// SetNumberOfFolders gets a reference to the given int32 and assigns it to the NumberOfFolders field.
func (o *DevopsPipeline) SetNumberOfFolders(v int32) {
	o.NumberOfFolders = &v
}

// GetNumberOfPipelines returns the NumberOfPipelines field value if set, zero value otherwise.
func (o *DevopsPipeline) GetNumberOfPipelines() int32 {
	if o == nil || IsNil(o.NumberOfPipelines) {
		var ret int32
		return ret
	}
	return *o.NumberOfPipelines
}

// GetNumberOfPipelinesOk returns a tuple with the NumberOfPipelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetNumberOfPipelinesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfPipelines) {
		return nil, false
	}
	return o.NumberOfPipelines, true
}

// HasNumberOfPipelines returns a boolean if a field has been set.
func (o *DevopsPipeline) HasNumberOfPipelines() bool {
	if o != nil && !IsNil(o.NumberOfPipelines) {
		return true
	}

	return false
}

// SetNumberOfPipelines gets a reference to the given int32 and assigns it to the NumberOfPipelines field.
func (o *DevopsPipeline) SetNumberOfPipelines(v int32) {
	o.NumberOfPipelines = &v
}

// GetNumberOfSuccessfulBranches returns the NumberOfSuccessfulBranches field value if set, zero value otherwise.
func (o *DevopsPipeline) GetNumberOfSuccessfulBranches() int32 {
	if o == nil || IsNil(o.NumberOfSuccessfulBranches) {
		var ret int32
		return ret
	}
	return *o.NumberOfSuccessfulBranches
}

// GetNumberOfSuccessfulBranchesOk returns a tuple with the NumberOfSuccessfulBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetNumberOfSuccessfulBranchesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfSuccessfulBranches) {
		return nil, false
	}
	return o.NumberOfSuccessfulBranches, true
}

// HasNumberOfSuccessfulBranches returns a boolean if a field has been set.
func (o *DevopsPipeline) HasNumberOfSuccessfulBranches() bool {
	if o != nil && !IsNil(o.NumberOfSuccessfulBranches) {
		return true
	}

	return false
}

// SetNumberOfSuccessfulBranches gets a reference to the given int32 and assigns it to the NumberOfSuccessfulBranches field.
func (o *DevopsPipeline) SetNumberOfSuccessfulBranches(v int32) {
	o.NumberOfSuccessfulBranches = &v
}

// GetNumberOfSuccessfulPullRequests returns the NumberOfSuccessfulPullRequests field value if set, zero value otherwise.
func (o *DevopsPipeline) GetNumberOfSuccessfulPullRequests() int32 {
	if o == nil || IsNil(o.NumberOfSuccessfulPullRequests) {
		var ret int32
		return ret
	}
	return *o.NumberOfSuccessfulPullRequests
}

// GetNumberOfSuccessfulPullRequestsOk returns a tuple with the NumberOfSuccessfulPullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetNumberOfSuccessfulPullRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfSuccessfulPullRequests) {
		return nil, false
	}
	return o.NumberOfSuccessfulPullRequests, true
}

// HasNumberOfSuccessfulPullRequests returns a boolean if a field has been set.
func (o *DevopsPipeline) HasNumberOfSuccessfulPullRequests() bool {
	if o != nil && !IsNil(o.NumberOfSuccessfulPullRequests) {
		return true
	}

	return false
}

// SetNumberOfSuccessfulPullRequests gets a reference to the given int32 and assigns it to the NumberOfSuccessfulPullRequests field.
func (o *DevopsPipeline) SetNumberOfSuccessfulPullRequests(v int32) {
	o.NumberOfSuccessfulPullRequests = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DevopsPipeline) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DevopsPipeline) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *DevopsPipeline) SetOrganization(v string) {
	o.Organization = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *DevopsPipeline) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DevopsPipeline) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *DevopsPipeline) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *DevopsPipeline) GetPermissions() Permissions {
	if o == nil || IsNil(o.Permissions) {
		var ret Permissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetPermissionsOk() (*Permissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *DevopsPipeline) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given Permissions and assigns it to the Permissions field.
func (o *DevopsPipeline) SetPermissions(v Permissions) {
	o.Permissions = &v
}

// GetPipelineFolderNames returns the PipelineFolderNames field value if set, zero value otherwise.
func (o *DevopsPipeline) GetPipelineFolderNames() []map[string]interface{} {
	if o == nil || IsNil(o.PipelineFolderNames) {
		var ret []map[string]interface{}
		return ret
	}
	return o.PipelineFolderNames
}

// GetPipelineFolderNamesOk returns a tuple with the PipelineFolderNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetPipelineFolderNamesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.PipelineFolderNames) {
		return nil, false
	}
	return o.PipelineFolderNames, true
}

// HasPipelineFolderNames returns a boolean if a field has been set.
func (o *DevopsPipeline) HasPipelineFolderNames() bool {
	if o != nil && !IsNil(o.PipelineFolderNames) {
		return true
	}

	return false
}

// SetPipelineFolderNames gets a reference to the given []map[string]interface{} and assigns it to the PipelineFolderNames field.
func (o *DevopsPipeline) SetPipelineFolderNames(v []map[string]interface{}) {
	o.PipelineFolderNames = v
}

// GetScmSource returns the ScmSource field value if set, zero value otherwise.
func (o *DevopsPipeline) GetScmSource() ScmSource {
	if o == nil || IsNil(o.ScmSource) {
		var ret ScmSource
		return ret
	}
	return *o.ScmSource
}

// GetScmSourceOk returns a tuple with the ScmSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetScmSourceOk() (*ScmSource, bool) {
	if o == nil || IsNil(o.ScmSource) {
		return nil, false
	}
	return o.ScmSource, true
}

// HasScmSource returns a boolean if a field has been set.
func (o *DevopsPipeline) HasScmSource() bool {
	if o != nil && !IsNil(o.ScmSource) {
		return true
	}

	return false
}

// SetScmSource gets a reference to the given ScmSource and assigns it to the ScmSource field.
func (o *DevopsPipeline) SetScmSource(v ScmSource) {
	o.ScmSource = &v
}

// GetTotalNumberOfBranches returns the TotalNumberOfBranches field value if set, zero value otherwise.
func (o *DevopsPipeline) GetTotalNumberOfBranches() int32 {
	if o == nil || IsNil(o.TotalNumberOfBranches) {
		var ret int32
		return ret
	}
	return *o.TotalNumberOfBranches
}

// GetTotalNumberOfBranchesOk returns a tuple with the TotalNumberOfBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetTotalNumberOfBranchesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNumberOfBranches) {
		return nil, false
	}
	return o.TotalNumberOfBranches, true
}

// HasTotalNumberOfBranches returns a boolean if a field has been set.
func (o *DevopsPipeline) HasTotalNumberOfBranches() bool {
	if o != nil && !IsNil(o.TotalNumberOfBranches) {
		return true
	}

	return false
}

// SetTotalNumberOfBranches gets a reference to the given int32 and assigns it to the TotalNumberOfBranches field.
func (o *DevopsPipeline) SetTotalNumberOfBranches(v int32) {
	o.TotalNumberOfBranches = &v
}

// GetTotalNumberOfPullRequests returns the TotalNumberOfPullRequests field value if set, zero value otherwise.
func (o *DevopsPipeline) GetTotalNumberOfPullRequests() int32 {
	if o == nil || IsNil(o.TotalNumberOfPullRequests) {
		var ret int32
		return ret
	}
	return *o.TotalNumberOfPullRequests
}

// GetTotalNumberOfPullRequestsOk returns a tuple with the TotalNumberOfPullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetTotalNumberOfPullRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNumberOfPullRequests) {
		return nil, false
	}
	return o.TotalNumberOfPullRequests, true
}

// HasTotalNumberOfPullRequests returns a boolean if a field has been set.
func (o *DevopsPipeline) HasTotalNumberOfPullRequests() bool {
	if o != nil && !IsNil(o.TotalNumberOfPullRequests) {
		return true
	}

	return false
}

// SetTotalNumberOfPullRequests gets a reference to the given int32 and assigns it to the TotalNumberOfPullRequests field.
func (o *DevopsPipeline) SetTotalNumberOfPullRequests(v int32) {
	o.TotalNumberOfPullRequests = &v
}

// GetWeatherScore returns the WeatherScore field value
func (o *DevopsPipeline) GetWeatherScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WeatherScore
}

// GetWeatherScoreOk returns a tuple with the WeatherScore field value
// and a boolean to check if the value has been set.
func (o *DevopsPipeline) GetWeatherScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeatherScore, true
}

// SetWeatherScore sets field value
func (o *DevopsPipeline) SetWeatherScore(v int32) {
	o.WeatherScore = v
}

func (o DevopsPipeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsPipeline) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BranchNames) {
		toSerialize["branchNames"] = o.BranchNames
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.EstimatedDurationInMillis) {
		toSerialize["estimatedDurationInMillis"] = o.EstimatedDurationInMillis
	}
	if !IsNil(o.FullDisplayName) {
		toSerialize["fullDisplayName"] = o.FullDisplayName
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumberOfFailingBranches) {
		toSerialize["numberOfFailingBranches"] = o.NumberOfFailingBranches
	}
	if !IsNil(o.NumberOfFailingPullRequests) {
		toSerialize["numberOfFailingPullRequests"] = o.NumberOfFailingPullRequests
	}
	if !IsNil(o.NumberOfFolders) {
		toSerialize["numberOfFolders"] = o.NumberOfFolders
	}
	if !IsNil(o.NumberOfPipelines) {
		toSerialize["numberOfPipelines"] = o.NumberOfPipelines
	}
	if !IsNil(o.NumberOfSuccessfulBranches) {
		toSerialize["numberOfSuccessfulBranches"] = o.NumberOfSuccessfulBranches
	}
	if !IsNil(o.NumberOfSuccessfulPullRequests) {
		toSerialize["numberOfSuccessfulPullRequests"] = o.NumberOfSuccessfulPullRequests
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.PipelineFolderNames) {
		toSerialize["pipelineFolderNames"] = o.PipelineFolderNames
	}
	if !IsNil(o.ScmSource) {
		toSerialize["scmSource"] = o.ScmSource
	}
	if !IsNil(o.TotalNumberOfBranches) {
		toSerialize["totalNumberOfBranches"] = o.TotalNumberOfBranches
	}
	if !IsNil(o.TotalNumberOfPullRequests) {
		toSerialize["totalNumberOfPullRequests"] = o.TotalNumberOfPullRequests
	}
	toSerialize["weatherScore"] = o.WeatherScore
	return toSerialize, nil
}

func (o *DevopsPipeline) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"weatherScore",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDevopsPipeline := _DevopsPipeline{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDevopsPipeline)

	if err != nil {
		return err
	}

	*o = DevopsPipeline(varDevopsPipeline)

	return err
}

type NullableDevopsPipeline struct {
	value *DevopsPipeline
	isSet bool
}

func (v NullableDevopsPipeline) Get() *DevopsPipeline {
	return v.value
}

func (v *NullableDevopsPipeline) Set(val *DevopsPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsPipeline(val *DevopsPipeline) *NullableDevopsPipeline {
	return &NullableDevopsPipeline{value: val, isSet: true}
}

func (v NullableDevopsPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

