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
	"bytes"
	"fmt"
)

// checks if the PkixAttributeTypeAndValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PkixAttributeTypeAndValue{}

// PkixAttributeTypeAndValue struct for PkixAttributeTypeAndValue
type PkixAttributeTypeAndValue struct {
	Type []int32 `json:"Type"`
	Value map[string]interface{} `json:"Value"`
}

type _PkixAttributeTypeAndValue PkixAttributeTypeAndValue

// NewPkixAttributeTypeAndValue instantiates a new PkixAttributeTypeAndValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixAttributeTypeAndValue(type_ []int32, value map[string]interface{}) *PkixAttributeTypeAndValue {
	this := PkixAttributeTypeAndValue{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewPkixAttributeTypeAndValueWithDefaults instantiates a new PkixAttributeTypeAndValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixAttributeTypeAndValueWithDefaults() *PkixAttributeTypeAndValue {
	this := PkixAttributeTypeAndValue{}
	return &this
}

// GetType returns the Type field value
func (o *PkixAttributeTypeAndValue) GetType() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PkixAttributeTypeAndValue) GetTypeOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *PkixAttributeTypeAndValue) SetType(v []int32) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *PkixAttributeTypeAndValue) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PkixAttributeTypeAndValue) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *PkixAttributeTypeAndValue) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o PkixAttributeTypeAndValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PkixAttributeTypeAndValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Type"] = o.Type
	toSerialize["Value"] = o.Value
	return toSerialize, nil
}

func (o *PkixAttributeTypeAndValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Type",
		"Value",
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

	varPkixAttributeTypeAndValue := _PkixAttributeTypeAndValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPkixAttributeTypeAndValue)

	if err != nil {
		return err
	}

	*o = PkixAttributeTypeAndValue(varPkixAttributeTypeAndValue)

	return err
}

type NullablePkixAttributeTypeAndValue struct {
	value *PkixAttributeTypeAndValue
	isSet bool
}

func (v NullablePkixAttributeTypeAndValue) Get() *PkixAttributeTypeAndValue {
	return v.value
}

func (v *NullablePkixAttributeTypeAndValue) Set(val *PkixAttributeTypeAndValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixAttributeTypeAndValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixAttributeTypeAndValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixAttributeTypeAndValue(val *PkixAttributeTypeAndValue) *NullablePkixAttributeTypeAndValue {
	return &NullablePkixAttributeTypeAndValue{value: val, isSet: true}
}

func (v NullablePkixAttributeTypeAndValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixAttributeTypeAndValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

