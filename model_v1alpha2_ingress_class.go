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

// checks if the V1alpha2IngressClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2IngressClass{}

// V1alpha2IngressClass struct for V1alpha2IngressClass
type V1alpha2IngressClass struct {
	Default *bool `json:"default,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1alpha2IngressClass instantiates a new V1alpha2IngressClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2IngressClass() *V1alpha2IngressClass {
	this := V1alpha2IngressClass{}
	return &this
}

// NewV1alpha2IngressClassWithDefaults instantiates a new V1alpha2IngressClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2IngressClassWithDefaults() *V1alpha2IngressClass {
	this := V1alpha2IngressClass{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *V1alpha2IngressClass) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2IngressClass) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *V1alpha2IngressClass) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *V1alpha2IngressClass) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1alpha2IngressClass) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2IngressClass) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1alpha2IngressClass) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1alpha2IngressClass) SetName(v string) {
	o.Name = &v
}

func (o V1alpha2IngressClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2IngressClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1alpha2IngressClass struct {
	value *V1alpha2IngressClass
	isSet bool
}

func (v NullableV1alpha2IngressClass) Get() *V1alpha2IngressClass {
	return v.value
}

func (v *NullableV1alpha2IngressClass) Set(val *V1alpha2IngressClass) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2IngressClass) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2IngressClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2IngressClass(val *V1alpha2IngressClass) *NullableV1alpha2IngressClass {
	return &NullableV1alpha2IngressClass{value: val, isSet: true}
}

func (v NullableV1alpha2IngressClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2IngressClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

