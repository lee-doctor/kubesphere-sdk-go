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

// checks if the V1StatefulSetPersistentVolumeClaimRetentionPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1StatefulSetPersistentVolumeClaimRetentionPolicy{}

// V1StatefulSetPersistentVolumeClaimRetentionPolicy StatefulSetPersistentVolumeClaimRetentionPolicy describes the policy used for PVCs created from the StatefulSet VolumeClaimTemplates.
type V1StatefulSetPersistentVolumeClaimRetentionPolicy struct {
	// WhenDeleted specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is deleted. The default policy of `Retain` causes PVCs to not be affected by StatefulSet deletion. The `Delete` policy causes those PVCs to be deleted.
	WhenDeleted *string `json:"whenDeleted,omitempty"`
	// WhenScaled specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is scaled down. The default policy of `Retain` causes PVCs to not be affected by a scaledown. The `Delete` policy causes the associated PVCs for any excess pods above the replica count to be deleted.
	WhenScaled *string `json:"whenScaled,omitempty"`
}

// NewV1StatefulSetPersistentVolumeClaimRetentionPolicy instantiates a new V1StatefulSetPersistentVolumeClaimRetentionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1StatefulSetPersistentVolumeClaimRetentionPolicy() *V1StatefulSetPersistentVolumeClaimRetentionPolicy {
	this := V1StatefulSetPersistentVolumeClaimRetentionPolicy{}
	return &this
}

// NewV1StatefulSetPersistentVolumeClaimRetentionPolicyWithDefaults instantiates a new V1StatefulSetPersistentVolumeClaimRetentionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1StatefulSetPersistentVolumeClaimRetentionPolicyWithDefaults() *V1StatefulSetPersistentVolumeClaimRetentionPolicy {
	this := V1StatefulSetPersistentVolumeClaimRetentionPolicy{}
	return &this
}

// GetWhenDeleted returns the WhenDeleted field value if set, zero value otherwise.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenDeleted() string {
	if o == nil || IsNil(o.WhenDeleted) {
		var ret string
		return ret
	}
	return *o.WhenDeleted
}

// GetWhenDeletedOk returns a tuple with the WhenDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenDeletedOk() (*string, bool) {
	if o == nil || IsNil(o.WhenDeleted) {
		return nil, false
	}
	return o.WhenDeleted, true
}

// HasWhenDeleted returns a boolean if a field has been set.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) HasWhenDeleted() bool {
	if o != nil && !IsNil(o.WhenDeleted) {
		return true
	}

	return false
}

// SetWhenDeleted gets a reference to the given string and assigns it to the WhenDeleted field.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) SetWhenDeleted(v string) {
	o.WhenDeleted = &v
}

// GetWhenScaled returns the WhenScaled field value if set, zero value otherwise.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenScaled() string {
	if o == nil || IsNil(o.WhenScaled) {
		var ret string
		return ret
	}
	return *o.WhenScaled
}

// GetWhenScaledOk returns a tuple with the WhenScaled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) GetWhenScaledOk() (*string, bool) {
	if o == nil || IsNil(o.WhenScaled) {
		return nil, false
	}
	return o.WhenScaled, true
}

// HasWhenScaled returns a boolean if a field has been set.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) HasWhenScaled() bool {
	if o != nil && !IsNil(o.WhenScaled) {
		return true
	}

	return false
}

// SetWhenScaled gets a reference to the given string and assigns it to the WhenScaled field.
func (o *V1StatefulSetPersistentVolumeClaimRetentionPolicy) SetWhenScaled(v string) {
	o.WhenScaled = &v
}

func (o V1StatefulSetPersistentVolumeClaimRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1StatefulSetPersistentVolumeClaimRetentionPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WhenDeleted) {
		toSerialize["whenDeleted"] = o.WhenDeleted
	}
	if !IsNil(o.WhenScaled) {
		toSerialize["whenScaled"] = o.WhenScaled
	}
	return toSerialize, nil
}

type NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy struct {
	value *V1StatefulSetPersistentVolumeClaimRetentionPolicy
	isSet bool
}

func (v NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy) Get() *V1StatefulSetPersistentVolumeClaimRetentionPolicy {
	return v.value
}

func (v *NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy) Set(val *V1StatefulSetPersistentVolumeClaimRetentionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1StatefulSetPersistentVolumeClaimRetentionPolicy(val *V1StatefulSetPersistentVolumeClaimRetentionPolicy) *NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy {
	return &NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy{value: val, isSet: true}
}

func (v NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1StatefulSetPersistentVolumeClaimRetentionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


