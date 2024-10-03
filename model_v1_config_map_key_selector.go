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
	"bytes"
	"fmt"
)

// checks if the V1ConfigMapKeySelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ConfigMapKeySelector{}

// V1ConfigMapKeySelector Selects a key from a ConfigMap.
type V1ConfigMapKeySelector struct {
	// The key to select.
	Key string `json:"key"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`
	// Specify whether the ConfigMap or its key must be defined
	Optional *bool `json:"optional,omitempty"`
}

type _V1ConfigMapKeySelector V1ConfigMapKeySelector

// NewV1ConfigMapKeySelector instantiates a new V1ConfigMapKeySelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ConfigMapKeySelector(key string) *V1ConfigMapKeySelector {
	this := V1ConfigMapKeySelector{}
	this.Key = key
	return &this
}

// NewV1ConfigMapKeySelectorWithDefaults instantiates a new V1ConfigMapKeySelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ConfigMapKeySelectorWithDefaults() *V1ConfigMapKeySelector {
	this := V1ConfigMapKeySelector{}
	return &this
}

// GetKey returns the Key field value
func (o *V1ConfigMapKeySelector) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *V1ConfigMapKeySelector) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *V1ConfigMapKeySelector) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ConfigMapKeySelector) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigMapKeySelector) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ConfigMapKeySelector) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ConfigMapKeySelector) SetName(v string) {
	o.Name = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *V1ConfigMapKeySelector) GetOptional() bool {
	if o == nil || IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigMapKeySelector) GetOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *V1ConfigMapKeySelector) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *V1ConfigMapKeySelector) SetOptional(v bool) {
	o.Optional = &v
}

func (o V1ConfigMapKeySelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ConfigMapKeySelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	return toSerialize, nil
}

func (o *V1ConfigMapKeySelector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV1ConfigMapKeySelector := _V1ConfigMapKeySelector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ConfigMapKeySelector)

	if err != nil {
		return err
	}

	*o = V1ConfigMapKeySelector(varV1ConfigMapKeySelector)

	return err
}

type NullableV1ConfigMapKeySelector struct {
	value *V1ConfigMapKeySelector
	isSet bool
}

func (v NullableV1ConfigMapKeySelector) Get() *V1ConfigMapKeySelector {
	return v.value
}

func (v *NullableV1ConfigMapKeySelector) Set(val *V1ConfigMapKeySelector) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ConfigMapKeySelector) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ConfigMapKeySelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ConfigMapKeySelector(val *V1ConfigMapKeySelector) *NullableV1ConfigMapKeySelector {
	return &NullableV1ConfigMapKeySelector{value: val, isSet: true}
}

func (v NullableV1ConfigMapKeySelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ConfigMapKeySelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


