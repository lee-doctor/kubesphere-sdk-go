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

// checks if the Links type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Links{}

// Links struct for Links
type Links struct {
	Actions *Actions `json:"actions,omitempty"`
	Branches *Branches `json:"branches,omitempty"`
	Queue *Queue `json:"queue,omitempty"`
	Runs *Runs `json:"runs,omitempty"`
	Scm *Scm `json:"scm,omitempty"`
	Self *Self `json:"self,omitempty"`
	Trends *Trends `json:"trends,omitempty"`
}

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks() *Links {
	this := Links{}
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *Links) GetActions() Actions {
	if o == nil || IsNil(o.Actions) {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetActionsOk() (*Actions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *Links) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *Links) SetActions(v Actions) {
	o.Actions = &v
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *Links) GetBranches() Branches {
	if o == nil || IsNil(o.Branches) {
		var ret Branches
		return ret
	}
	return *o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetBranchesOk() (*Branches, bool) {
	if o == nil || IsNil(o.Branches) {
		return nil, false
	}
	return o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *Links) HasBranches() bool {
	if o != nil && !IsNil(o.Branches) {
		return true
	}

	return false
}

// SetBranches gets a reference to the given Branches and assigns it to the Branches field.
func (o *Links) SetBranches(v Branches) {
	o.Branches = &v
}

// GetQueue returns the Queue field value if set, zero value otherwise.
func (o *Links) GetQueue() Queue {
	if o == nil || IsNil(o.Queue) {
		var ret Queue
		return ret
	}
	return *o.Queue
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetQueueOk() (*Queue, bool) {
	if o == nil || IsNil(o.Queue) {
		return nil, false
	}
	return o.Queue, true
}

// HasQueue returns a boolean if a field has been set.
func (o *Links) HasQueue() bool {
	if o != nil && !IsNil(o.Queue) {
		return true
	}

	return false
}

// SetQueue gets a reference to the given Queue and assigns it to the Queue field.
func (o *Links) SetQueue(v Queue) {
	o.Queue = &v
}

// GetRuns returns the Runs field value if set, zero value otherwise.
func (o *Links) GetRuns() Runs {
	if o == nil || IsNil(o.Runs) {
		var ret Runs
		return ret
	}
	return *o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetRunsOk() (*Runs, bool) {
	if o == nil || IsNil(o.Runs) {
		return nil, false
	}
	return o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *Links) HasRuns() bool {
	if o != nil && !IsNil(o.Runs) {
		return true
	}

	return false
}

// SetRuns gets a reference to the given Runs and assigns it to the Runs field.
func (o *Links) SetRuns(v Runs) {
	o.Runs = &v
}

// GetScm returns the Scm field value if set, zero value otherwise.
func (o *Links) GetScm() Scm {
	if o == nil || IsNil(o.Scm) {
		var ret Scm
		return ret
	}
	return *o.Scm
}

// GetScmOk returns a tuple with the Scm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetScmOk() (*Scm, bool) {
	if o == nil || IsNil(o.Scm) {
		return nil, false
	}
	return o.Scm, true
}

// HasScm returns a boolean if a field has been set.
func (o *Links) HasScm() bool {
	if o != nil && !IsNil(o.Scm) {
		return true
	}

	return false
}

// SetScm gets a reference to the given Scm and assigns it to the Scm field.
func (o *Links) SetScm(v Scm) {
	o.Scm = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Links) GetSelf() Self {
	if o == nil || IsNil(o.Self) {
		var ret Self
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetSelfOk() (*Self, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Links) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Self and assigns it to the Self field.
func (o *Links) SetSelf(v Self) {
	o.Self = &v
}

// GetTrends returns the Trends field value if set, zero value otherwise.
func (o *Links) GetTrends() Trends {
	if o == nil || IsNil(o.Trends) {
		var ret Trends
		return ret
	}
	return *o.Trends
}

// GetTrendsOk returns a tuple with the Trends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetTrendsOk() (*Trends, bool) {
	if o == nil || IsNil(o.Trends) {
		return nil, false
	}
	return o.Trends, true
}

// HasTrends returns a boolean if a field has been set.
func (o *Links) HasTrends() bool {
	if o != nil && !IsNil(o.Trends) {
		return true
	}

	return false
}

// SetTrends gets a reference to the given Trends and assigns it to the Trends field.
func (o *Links) SetTrends(v Trends) {
	o.Trends = &v
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Links) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Branches) {
		toSerialize["branches"] = o.Branches
	}
	if !IsNil(o.Queue) {
		toSerialize["queue"] = o.Queue
	}
	if !IsNil(o.Runs) {
		toSerialize["runs"] = o.Runs
	}
	if !IsNil(o.Scm) {
		toSerialize["scm"] = o.Scm
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Trends) {
		toSerialize["trends"] = o.Trends
	}
	return toSerialize, nil
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


