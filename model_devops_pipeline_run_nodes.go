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

// checks if the DevopsPipelineRunNodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsPipelineRunNodes{}

// DevopsPipelineRunNodes struct for DevopsPipelineRunNodes
type DevopsPipelineRunNodes struct {
	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class *string `json:"_class,omitempty"`
	Links *Links `json:"_links,omitempty"`
	// the list of all actions
	Actions []map[string]interface{} `json:"actions,omitempty"`
	CauseOfBlockage map[string]interface{} `json:"causeOfBlockage,omitempty"`
	DisplayDescription map[string]interface{} `json:"displayDescription,omitempty"`
	// display name
	DisplayName *string `json:"displayName,omitempty"`
	// duration time in mullis
	DurationInMillis *int32 `json:"durationInMillis,omitempty"`
	// edges
	Edges []map[string]interface{} `json:"edges,omitempty"`
	FirstParent map[string]interface{} `json:"firstParent,omitempty"`
	// id
	Id *string `json:"id,omitempty"`
	Input *DevopsInput `json:"input,omitempty"`
	// restartable or not
	Restartable *bool `json:"restartable,omitempty"`
	// the result of pipeline run. e.g. SUCCESS
	Result *string `json:"result,omitempty"`
	// the time of start
	StartTime *string `json:"startTime,omitempty"`
	// run state. e.g. FINISHED
	State *string `json:"state,omitempty"`
	// type
	Type *string `json:"type,omitempty"`
}

// NewDevopsPipelineRunNodes instantiates a new DevopsPipelineRunNodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsPipelineRunNodes() *DevopsPipelineRunNodes {
	this := DevopsPipelineRunNodes{}
	return &this
}

// NewDevopsPipelineRunNodesWithDefaults instantiates a new DevopsPipelineRunNodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsPipelineRunNodesWithDefaults() *DevopsPipelineRunNodes {
	this := DevopsPipelineRunNodes{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *DevopsPipelineRunNodes) SetClass(v string) {
	o.Class = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *DevopsPipelineRunNodes) SetLinks(v Links) {
	o.Links = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetActions() []map[string]interface{} {
	if o == nil || IsNil(o.Actions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetActionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []map[string]interface{} and assigns it to the Actions field.
func (o *DevopsPipelineRunNodes) SetActions(v []map[string]interface{}) {
	o.Actions = v
}

// GetCauseOfBlockage returns the CauseOfBlockage field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetCauseOfBlockage() map[string]interface{} {
	if o == nil || IsNil(o.CauseOfBlockage) {
		var ret map[string]interface{}
		return ret
	}
	return o.CauseOfBlockage
}

// GetCauseOfBlockageOk returns a tuple with the CauseOfBlockage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetCauseOfBlockageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CauseOfBlockage) {
		return map[string]interface{}{}, false
	}
	return o.CauseOfBlockage, true
}

// HasCauseOfBlockage returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasCauseOfBlockage() bool {
	if o != nil && !IsNil(o.CauseOfBlockage) {
		return true
	}

	return false
}

// SetCauseOfBlockage gets a reference to the given map[string]interface{} and assigns it to the CauseOfBlockage field.
func (o *DevopsPipelineRunNodes) SetCauseOfBlockage(v map[string]interface{}) {
	o.CauseOfBlockage = v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetDisplayDescription() map[string]interface{} {
	if o == nil || IsNil(o.DisplayDescription) {
		var ret map[string]interface{}
		return ret
	}
	return o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetDisplayDescriptionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DisplayDescription) {
		return map[string]interface{}{}, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasDisplayDescription() bool {
	if o != nil && !IsNil(o.DisplayDescription) {
		return true
	}

	return false
}

// SetDisplayDescription gets a reference to the given map[string]interface{} and assigns it to the DisplayDescription field.
func (o *DevopsPipelineRunNodes) SetDisplayDescription(v map[string]interface{}) {
	o.DisplayDescription = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DevopsPipelineRunNodes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDurationInMillis returns the DurationInMillis field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetDurationInMillis() int32 {
	if o == nil || IsNil(o.DurationInMillis) {
		var ret int32
		return ret
	}
	return *o.DurationInMillis
}

// GetDurationInMillisOk returns a tuple with the DurationInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetDurationInMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInMillis) {
		return nil, false
	}
	return o.DurationInMillis, true
}

// HasDurationInMillis returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasDurationInMillis() bool {
	if o != nil && !IsNil(o.DurationInMillis) {
		return true
	}

	return false
}

// SetDurationInMillis gets a reference to the given int32 and assigns it to the DurationInMillis field.
func (o *DevopsPipelineRunNodes) SetDurationInMillis(v int32) {
	o.DurationInMillis = &v
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetEdges() []map[string]interface{} {
	if o == nil || IsNil(o.Edges) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetEdgesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Edges) {
		return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasEdges() bool {
	if o != nil && !IsNil(o.Edges) {
		return true
	}

	return false
}

// SetEdges gets a reference to the given []map[string]interface{} and assigns it to the Edges field.
func (o *DevopsPipelineRunNodes) SetEdges(v []map[string]interface{}) {
	o.Edges = v
}

// GetFirstParent returns the FirstParent field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetFirstParent() map[string]interface{} {
	if o == nil || IsNil(o.FirstParent) {
		var ret map[string]interface{}
		return ret
	}
	return o.FirstParent
}

// GetFirstParentOk returns a tuple with the FirstParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetFirstParentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FirstParent) {
		return map[string]interface{}{}, false
	}
	return o.FirstParent, true
}

// HasFirstParent returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasFirstParent() bool {
	if o != nil && !IsNil(o.FirstParent) {
		return true
	}

	return false
}

// SetFirstParent gets a reference to the given map[string]interface{} and assigns it to the FirstParent field.
func (o *DevopsPipelineRunNodes) SetFirstParent(v map[string]interface{}) {
	o.FirstParent = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevopsPipelineRunNodes) SetId(v string) {
	o.Id = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetInput() DevopsInput {
	if o == nil || IsNil(o.Input) {
		var ret DevopsInput
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetInputOk() (*DevopsInput, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given DevopsInput and assigns it to the Input field.
func (o *DevopsPipelineRunNodes) SetInput(v DevopsInput) {
	o.Input = &v
}

// GetRestartable returns the Restartable field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetRestartable() bool {
	if o == nil || IsNil(o.Restartable) {
		var ret bool
		return ret
	}
	return *o.Restartable
}

// GetRestartableOk returns a tuple with the Restartable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetRestartableOk() (*bool, bool) {
	if o == nil || IsNil(o.Restartable) {
		return nil, false
	}
	return o.Restartable, true
}

// HasRestartable returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasRestartable() bool {
	if o != nil && !IsNil(o.Restartable) {
		return true
	}

	return false
}

// SetRestartable gets a reference to the given bool and assigns it to the Restartable field.
func (o *DevopsPipelineRunNodes) SetRestartable(v bool) {
	o.Restartable = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DevopsPipelineRunNodes) SetResult(v string) {
	o.Result = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *DevopsPipelineRunNodes) SetStartTime(v string) {
	o.StartTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DevopsPipelineRunNodes) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DevopsPipelineRunNodes) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsPipelineRunNodes) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DevopsPipelineRunNodes) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DevopsPipelineRunNodes) SetType(v string) {
	o.Type = &v
}

func (o DevopsPipelineRunNodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsPipelineRunNodes) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CauseOfBlockage) {
		toSerialize["causeOfBlockage"] = o.CauseOfBlockage
	}
	if !IsNil(o.DisplayDescription) {
		toSerialize["displayDescription"] = o.DisplayDescription
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DurationInMillis) {
		toSerialize["durationInMillis"] = o.DurationInMillis
	}
	if !IsNil(o.Edges) {
		toSerialize["edges"] = o.Edges
	}
	if !IsNil(o.FirstParent) {
		toSerialize["firstParent"] = o.FirstParent
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Restartable) {
		toSerialize["restartable"] = o.Restartable
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
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

type NullableDevopsPipelineRunNodes struct {
	value *DevopsPipelineRunNodes
	isSet bool
}

func (v NullableDevopsPipelineRunNodes) Get() *DevopsPipelineRunNodes {
	return v.value
}

func (v *NullableDevopsPipelineRunNodes) Set(val *DevopsPipelineRunNodes) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsPipelineRunNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsPipelineRunNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsPipelineRunNodes(val *DevopsPipelineRunNodes) *NullableDevopsPipelineRunNodes {
	return &NullableDevopsPipelineRunNodes{value: val, isSet: true}
}

func (v NullableDevopsPipelineRunNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsPipelineRunNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


