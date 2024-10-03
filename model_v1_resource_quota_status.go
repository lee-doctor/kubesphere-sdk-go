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

// checks if the V1ResourceQuotaStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ResourceQuotaStatus{}

// V1ResourceQuotaStatus ResourceQuotaStatus defines the enforced hard limits and observed use.
type V1ResourceQuotaStatus struct {
	// Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Hard *map[string]ResourceQuantity `json:"hard,omitempty"`
	// Used is the current observed total usage of the resource in the namespace.
	Used *map[string]ResourceQuantity `json:"used,omitempty"`
}

// NewV1ResourceQuotaStatus instantiates a new V1ResourceQuotaStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ResourceQuotaStatus() *V1ResourceQuotaStatus {
	this := V1ResourceQuotaStatus{}
	return &this
}

// NewV1ResourceQuotaStatusWithDefaults instantiates a new V1ResourceQuotaStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ResourceQuotaStatusWithDefaults() *V1ResourceQuotaStatus {
	this := V1ResourceQuotaStatus{}
	return &this
}

// GetHard returns the Hard field value if set, zero value otherwise.
func (o *V1ResourceQuotaStatus) GetHard() map[string]ResourceQuantity {
	if o == nil || IsNil(o.Hard) {
		var ret map[string]ResourceQuantity
		return ret
	}
	return *o.Hard
}

// GetHardOk returns a tuple with the Hard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ResourceQuotaStatus) GetHardOk() (*map[string]ResourceQuantity, bool) {
	if o == nil || IsNil(o.Hard) {
		return nil, false
	}
	return o.Hard, true
}

// HasHard returns a boolean if a field has been set.
func (o *V1ResourceQuotaStatus) HasHard() bool {
	if o != nil && !IsNil(o.Hard) {
		return true
	}

	return false
}

// SetHard gets a reference to the given map[string]ResourceQuantity and assigns it to the Hard field.
func (o *V1ResourceQuotaStatus) SetHard(v map[string]ResourceQuantity) {
	o.Hard = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *V1ResourceQuotaStatus) GetUsed() map[string]ResourceQuantity {
	if o == nil || IsNil(o.Used) {
		var ret map[string]ResourceQuantity
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ResourceQuotaStatus) GetUsedOk() (*map[string]ResourceQuantity, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *V1ResourceQuotaStatus) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given map[string]ResourceQuantity and assigns it to the Used field.
func (o *V1ResourceQuotaStatus) SetUsed(v map[string]ResourceQuantity) {
	o.Used = &v
}

func (o V1ResourceQuotaStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ResourceQuotaStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hard) {
		toSerialize["hard"] = o.Hard
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	return toSerialize, nil
}

type NullableV1ResourceQuotaStatus struct {
	value *V1ResourceQuotaStatus
	isSet bool
}

func (v NullableV1ResourceQuotaStatus) Get() *V1ResourceQuotaStatus {
	return v.value
}

func (v *NullableV1ResourceQuotaStatus) Set(val *V1ResourceQuotaStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ResourceQuotaStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ResourceQuotaStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ResourceQuotaStatus(val *V1ResourceQuotaStatus) *NullableV1ResourceQuotaStatus {
	return &NullableV1ResourceQuotaStatus{value: val, isSet: true}
}

func (v NullableV1ResourceQuotaStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ResourceQuotaStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

