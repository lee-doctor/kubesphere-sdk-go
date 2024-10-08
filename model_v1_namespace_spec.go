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

// checks if the V1NamespaceSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1NamespaceSpec{}

// V1NamespaceSpec NamespaceSpec describes the attributes on a Namespace.
type V1NamespaceSpec struct {
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
	Finalizers []string `json:"finalizers,omitempty"`
}

// NewV1NamespaceSpec instantiates a new V1NamespaceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NamespaceSpec() *V1NamespaceSpec {
	this := V1NamespaceSpec{}
	return &this
}

// NewV1NamespaceSpecWithDefaults instantiates a new V1NamespaceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NamespaceSpecWithDefaults() *V1NamespaceSpec {
	this := V1NamespaceSpec{}
	return &this
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *V1NamespaceSpec) GetFinalizers() []string {
	if o == nil || IsNil(o.Finalizers) {
		var ret []string
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamespaceSpec) GetFinalizersOk() ([]string, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *V1NamespaceSpec) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []string and assigns it to the Finalizers field.
func (o *V1NamespaceSpec) SetFinalizers(v []string) {
	o.Finalizers = v
}

func (o V1NamespaceSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1NamespaceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Finalizers) {
		toSerialize["finalizers"] = o.Finalizers
	}
	return toSerialize, nil
}

type NullableV1NamespaceSpec struct {
	value *V1NamespaceSpec
	isSet bool
}

func (v NullableV1NamespaceSpec) Get() *V1NamespaceSpec {
	return v.value
}

func (v *NullableV1NamespaceSpec) Set(val *V1NamespaceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NamespaceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NamespaceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NamespaceSpec(val *V1NamespaceSpec) *NullableV1NamespaceSpec {
	return &NullableV1NamespaceSpec{value: val, isSet: true}
}

func (v NullableV1NamespaceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NamespaceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


