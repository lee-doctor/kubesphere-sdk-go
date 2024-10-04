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

// checks if the V1NodeAffinity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1NodeAffinity{}

// V1NodeAffinity Node affinity is a group of node affinity scheduling rules.
type V1NodeAffinity struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []V1PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	RequiredDuringSchedulingIgnoredDuringExecution *V1NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NewV1NodeAffinity instantiates a new V1NodeAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NodeAffinity() *V1NodeAffinity {
	this := V1NodeAffinity{}
	return &this
}

// NewV1NodeAffinityWithDefaults instantiates a new V1NodeAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NodeAffinityWithDefaults() *V1NodeAffinity {
	this := V1NodeAffinity{}
	return &this
}

// GetPreferredDuringSchedulingIgnoredDuringExecution returns the PreferredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *V1NodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []V1PreferredSchedulingTerm {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		var ret []V1PreferredSchedulingTerm
		return ret
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution
}

// GetPreferredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the PreferredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecutionOk() ([]V1PreferredSchedulingTerm, bool) {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution, true
}

// HasPreferredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *V1NodeAffinity) HasPreferredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetPreferredDuringSchedulingIgnoredDuringExecution gets a reference to the given []V1PreferredSchedulingTerm and assigns it to the PreferredDuringSchedulingIgnoredDuringExecution field.
func (o *V1NodeAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []V1PreferredSchedulingTerm) {
	o.PreferredDuringSchedulingIgnoredDuringExecution = v
}

// GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *V1NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() V1NodeSelector {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		var ret V1NodeSelector
		return ret
	}
	return *o.RequiredDuringSchedulingIgnoredDuringExecution
}

// GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() (*V1NodeSelector, bool) {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.RequiredDuringSchedulingIgnoredDuringExecution, true
}

// HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *V1NodeAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetRequiredDuringSchedulingIgnoredDuringExecution gets a reference to the given V1NodeSelector and assigns it to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (o *V1NodeAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v V1NodeSelector) {
	o.RequiredDuringSchedulingIgnoredDuringExecution = &v
}

func (o V1NodeAffinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1NodeAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["preferredDuringSchedulingIgnoredDuringExecution"] = o.PreferredDuringSchedulingIgnoredDuringExecution
	}
	if !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["requiredDuringSchedulingIgnoredDuringExecution"] = o.RequiredDuringSchedulingIgnoredDuringExecution
	}
	return toSerialize, nil
}

type NullableV1NodeAffinity struct {
	value *V1NodeAffinity
	isSet bool
}

func (v NullableV1NodeAffinity) Get() *V1NodeAffinity {
	return v.value
}

func (v *NullableV1NodeAffinity) Set(val *V1NodeAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NodeAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NodeAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NodeAffinity(val *V1NodeAffinity) *NullableV1NodeAffinity {
	return &NullableV1NodeAffinity{value: val, isSet: true}
}

func (v NullableV1NodeAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NodeAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


