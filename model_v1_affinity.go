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

// checks if the V1Affinity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Affinity{}

// V1Affinity Affinity is a group of affinity scheduling rules.
type V1Affinity struct {
	NodeAffinity *V1NodeAffinity `json:"nodeAffinity,omitempty"`
	PodAffinity *V1PodAffinity `json:"podAffinity,omitempty"`
	PodAntiAffinity *V1PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

// NewV1Affinity instantiates a new V1Affinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Affinity() *V1Affinity {
	this := V1Affinity{}
	return &this
}

// NewV1AffinityWithDefaults instantiates a new V1Affinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AffinityWithDefaults() *V1Affinity {
	this := V1Affinity{}
	return &this
}

// GetNodeAffinity returns the NodeAffinity field value if set, zero value otherwise.
func (o *V1Affinity) GetNodeAffinity() V1NodeAffinity {
	if o == nil || IsNil(o.NodeAffinity) {
		var ret V1NodeAffinity
		return ret
	}
	return *o.NodeAffinity
}

// GetNodeAffinityOk returns a tuple with the NodeAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Affinity) GetNodeAffinityOk() (*V1NodeAffinity, bool) {
	if o == nil || IsNil(o.NodeAffinity) {
		return nil, false
	}
	return o.NodeAffinity, true
}

// HasNodeAffinity returns a boolean if a field has been set.
func (o *V1Affinity) HasNodeAffinity() bool {
	if o != nil && !IsNil(o.NodeAffinity) {
		return true
	}

	return false
}

// SetNodeAffinity gets a reference to the given V1NodeAffinity and assigns it to the NodeAffinity field.
func (o *V1Affinity) SetNodeAffinity(v V1NodeAffinity) {
	o.NodeAffinity = &v
}

// GetPodAffinity returns the PodAffinity field value if set, zero value otherwise.
func (o *V1Affinity) GetPodAffinity() V1PodAffinity {
	if o == nil || IsNil(o.PodAffinity) {
		var ret V1PodAffinity
		return ret
	}
	return *o.PodAffinity
}

// GetPodAffinityOk returns a tuple with the PodAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Affinity) GetPodAffinityOk() (*V1PodAffinity, bool) {
	if o == nil || IsNil(o.PodAffinity) {
		return nil, false
	}
	return o.PodAffinity, true
}

// HasPodAffinity returns a boolean if a field has been set.
func (o *V1Affinity) HasPodAffinity() bool {
	if o != nil && !IsNil(o.PodAffinity) {
		return true
	}

	return false
}

// SetPodAffinity gets a reference to the given V1PodAffinity and assigns it to the PodAffinity field.
func (o *V1Affinity) SetPodAffinity(v V1PodAffinity) {
	o.PodAffinity = &v
}

// GetPodAntiAffinity returns the PodAntiAffinity field value if set, zero value otherwise.
func (o *V1Affinity) GetPodAntiAffinity() V1PodAntiAffinity {
	if o == nil || IsNil(o.PodAntiAffinity) {
		var ret V1PodAntiAffinity
		return ret
	}
	return *o.PodAntiAffinity
}

// GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Affinity) GetPodAntiAffinityOk() (*V1PodAntiAffinity, bool) {
	if o == nil || IsNil(o.PodAntiAffinity) {
		return nil, false
	}
	return o.PodAntiAffinity, true
}

// HasPodAntiAffinity returns a boolean if a field has been set.
func (o *V1Affinity) HasPodAntiAffinity() bool {
	if o != nil && !IsNil(o.PodAntiAffinity) {
		return true
	}

	return false
}

// SetPodAntiAffinity gets a reference to the given V1PodAntiAffinity and assigns it to the PodAntiAffinity field.
func (o *V1Affinity) SetPodAntiAffinity(v V1PodAntiAffinity) {
	o.PodAntiAffinity = &v
}

func (o V1Affinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Affinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodeAffinity) {
		toSerialize["nodeAffinity"] = o.NodeAffinity
	}
	if !IsNil(o.PodAffinity) {
		toSerialize["podAffinity"] = o.PodAffinity
	}
	if !IsNil(o.PodAntiAffinity) {
		toSerialize["podAntiAffinity"] = o.PodAntiAffinity
	}
	return toSerialize, nil
}

type NullableV1Affinity struct {
	value *V1Affinity
	isSet bool
}

func (v NullableV1Affinity) Get() *V1Affinity {
	return v.value
}

func (v *NullableV1Affinity) Set(val *V1Affinity) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Affinity) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Affinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Affinity(val *V1Affinity) *NullableV1Affinity {
	return &NullableV1Affinity{value: val, isSet: true}
}

func (v NullableV1Affinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Affinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


