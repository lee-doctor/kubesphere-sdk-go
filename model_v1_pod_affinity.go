/*
KS API

KubeSphere OpenAPI

API version: v4.1.1
Contact: support@kubesphere.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V1PodAffinity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PodAffinity{}

// V1PodAffinity Pod affinity is a group of inter pod affinity scheduling rules.
type V1PodAffinity struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []V1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	// If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
	RequiredDuringSchedulingIgnoredDuringExecution []V1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NewV1PodAffinity instantiates a new V1PodAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PodAffinity() *V1PodAffinity {
	this := V1PodAffinity{}
	return &this
}

// NewV1PodAffinityWithDefaults instantiates a new V1PodAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PodAffinityWithDefaults() *V1PodAffinity {
	this := V1PodAffinity{}
	return &this
}

// GetPreferredDuringSchedulingIgnoredDuringExecution returns the PreferredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *V1PodAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []V1WeightedPodAffinityTerm {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		var ret []V1WeightedPodAffinityTerm
		return ret
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution
}

// GetPreferredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the PreferredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodAffinity) GetPreferredDuringSchedulingIgnoredDuringExecutionOk() ([]V1WeightedPodAffinityTerm, bool) {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution, true
}

// HasPreferredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *V1PodAffinity) HasPreferredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetPreferredDuringSchedulingIgnoredDuringExecution gets a reference to the given []V1WeightedPodAffinityTerm and assigns it to the PreferredDuringSchedulingIgnoredDuringExecution field.
func (o *V1PodAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []V1WeightedPodAffinityTerm) {
	o.PreferredDuringSchedulingIgnoredDuringExecution = v
}

// GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *V1PodAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() []V1PodAffinityTerm {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		var ret []V1PodAffinityTerm
		return ret
	}
	return o.RequiredDuringSchedulingIgnoredDuringExecution
}

// GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() ([]V1PodAffinityTerm, bool) {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.RequiredDuringSchedulingIgnoredDuringExecution, true
}

// HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *V1PodAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetRequiredDuringSchedulingIgnoredDuringExecution gets a reference to the given []V1PodAffinityTerm and assigns it to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (o *V1PodAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []V1PodAffinityTerm) {
	o.RequiredDuringSchedulingIgnoredDuringExecution = v
}

func (o V1PodAffinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PodAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["preferredDuringSchedulingIgnoredDuringExecution"] = o.PreferredDuringSchedulingIgnoredDuringExecution
	}
	if !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["requiredDuringSchedulingIgnoredDuringExecution"] = o.RequiredDuringSchedulingIgnoredDuringExecution
	}
	return toSerialize, nil
}

type NullableV1PodAffinity struct {
	value *V1PodAffinity
	isSet bool
}

func (v NullableV1PodAffinity) Get() *V1PodAffinity {
	return v.value
}

func (v *NullableV1PodAffinity) Set(val *V1PodAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PodAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PodAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PodAffinity(val *V1PodAffinity) *NullableV1PodAffinity {
	return &NullableV1PodAffinity{value: val, isSet: true}
}

func (v NullableV1PodAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PodAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


