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

// checks if the V1EnvVarSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EnvVarSource{}

// V1EnvVarSource EnvVarSource represents a source for the value of an EnvVar.
type V1EnvVarSource struct {
	ConfigMapKeyRef *V1ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	FieldRef *V1ObjectFieldSelector `json:"fieldRef,omitempty"`
	ResourceFieldRef *V1ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	SecretKeyRef *V1SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// NewV1EnvVarSource instantiates a new V1EnvVarSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EnvVarSource() *V1EnvVarSource {
	this := V1EnvVarSource{}
	return &this
}

// NewV1EnvVarSourceWithDefaults instantiates a new V1EnvVarSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EnvVarSourceWithDefaults() *V1EnvVarSource {
	this := V1EnvVarSource{}
	return &this
}

// GetConfigMapKeyRef returns the ConfigMapKeyRef field value if set, zero value otherwise.
func (o *V1EnvVarSource) GetConfigMapKeyRef() V1ConfigMapKeySelector {
	if o == nil || IsNil(o.ConfigMapKeyRef) {
		var ret V1ConfigMapKeySelector
		return ret
	}
	return *o.ConfigMapKeyRef
}

// GetConfigMapKeyRefOk returns a tuple with the ConfigMapKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnvVarSource) GetConfigMapKeyRefOk() (*V1ConfigMapKeySelector, bool) {
	if o == nil || IsNil(o.ConfigMapKeyRef) {
		return nil, false
	}
	return o.ConfigMapKeyRef, true
}

// HasConfigMapKeyRef returns a boolean if a field has been set.
func (o *V1EnvVarSource) HasConfigMapKeyRef() bool {
	if o != nil && !IsNil(o.ConfigMapKeyRef) {
		return true
	}

	return false
}

// SetConfigMapKeyRef gets a reference to the given V1ConfigMapKeySelector and assigns it to the ConfigMapKeyRef field.
func (o *V1EnvVarSource) SetConfigMapKeyRef(v V1ConfigMapKeySelector) {
	o.ConfigMapKeyRef = &v
}

// GetFieldRef returns the FieldRef field value if set, zero value otherwise.
func (o *V1EnvVarSource) GetFieldRef() V1ObjectFieldSelector {
	if o == nil || IsNil(o.FieldRef) {
		var ret V1ObjectFieldSelector
		return ret
	}
	return *o.FieldRef
}

// GetFieldRefOk returns a tuple with the FieldRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnvVarSource) GetFieldRefOk() (*V1ObjectFieldSelector, bool) {
	if o == nil || IsNil(o.FieldRef) {
		return nil, false
	}
	return o.FieldRef, true
}

// HasFieldRef returns a boolean if a field has been set.
func (o *V1EnvVarSource) HasFieldRef() bool {
	if o != nil && !IsNil(o.FieldRef) {
		return true
	}

	return false
}

// SetFieldRef gets a reference to the given V1ObjectFieldSelector and assigns it to the FieldRef field.
func (o *V1EnvVarSource) SetFieldRef(v V1ObjectFieldSelector) {
	o.FieldRef = &v
}

// GetResourceFieldRef returns the ResourceFieldRef field value if set, zero value otherwise.
func (o *V1EnvVarSource) GetResourceFieldRef() V1ResourceFieldSelector {
	if o == nil || IsNil(o.ResourceFieldRef) {
		var ret V1ResourceFieldSelector
		return ret
	}
	return *o.ResourceFieldRef
}

// GetResourceFieldRefOk returns a tuple with the ResourceFieldRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnvVarSource) GetResourceFieldRefOk() (*V1ResourceFieldSelector, bool) {
	if o == nil || IsNil(o.ResourceFieldRef) {
		return nil, false
	}
	return o.ResourceFieldRef, true
}

// HasResourceFieldRef returns a boolean if a field has been set.
func (o *V1EnvVarSource) HasResourceFieldRef() bool {
	if o != nil && !IsNil(o.ResourceFieldRef) {
		return true
	}

	return false
}

// SetResourceFieldRef gets a reference to the given V1ResourceFieldSelector and assigns it to the ResourceFieldRef field.
func (o *V1EnvVarSource) SetResourceFieldRef(v V1ResourceFieldSelector) {
	o.ResourceFieldRef = &v
}

// GetSecretKeyRef returns the SecretKeyRef field value if set, zero value otherwise.
func (o *V1EnvVarSource) GetSecretKeyRef() V1SecretKeySelector {
	if o == nil || IsNil(o.SecretKeyRef) {
		var ret V1SecretKeySelector
		return ret
	}
	return *o.SecretKeyRef
}

// GetSecretKeyRefOk returns a tuple with the SecretKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnvVarSource) GetSecretKeyRefOk() (*V1SecretKeySelector, bool) {
	if o == nil || IsNil(o.SecretKeyRef) {
		return nil, false
	}
	return o.SecretKeyRef, true
}

// HasSecretKeyRef returns a boolean if a field has been set.
func (o *V1EnvVarSource) HasSecretKeyRef() bool {
	if o != nil && !IsNil(o.SecretKeyRef) {
		return true
	}

	return false
}

// SetSecretKeyRef gets a reference to the given V1SecretKeySelector and assigns it to the SecretKeyRef field.
func (o *V1EnvVarSource) SetSecretKeyRef(v V1SecretKeySelector) {
	o.SecretKeyRef = &v
}

func (o V1EnvVarSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EnvVarSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigMapKeyRef) {
		toSerialize["configMapKeyRef"] = o.ConfigMapKeyRef
	}
	if !IsNil(o.FieldRef) {
		toSerialize["fieldRef"] = o.FieldRef
	}
	if !IsNil(o.ResourceFieldRef) {
		toSerialize["resourceFieldRef"] = o.ResourceFieldRef
	}
	if !IsNil(o.SecretKeyRef) {
		toSerialize["secretKeyRef"] = o.SecretKeyRef
	}
	return toSerialize, nil
}

type NullableV1EnvVarSource struct {
	value *V1EnvVarSource
	isSet bool
}

func (v NullableV1EnvVarSource) Get() *V1EnvVarSource {
	return v.value
}

func (v *NullableV1EnvVarSource) Set(val *V1EnvVarSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EnvVarSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EnvVarSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EnvVarSource(val *V1EnvVarSource) *NullableV1EnvVarSource {
	return &NullableV1EnvVarSource{value: val, isSet: true}
}

func (v NullableV1EnvVarSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EnvVarSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

