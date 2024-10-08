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

// checks if the OverviewMetricValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverviewMetricValue{}

// OverviewMetricValue struct for OverviewMetricValue
type OverviewMetricValue struct {
	Value []map[string]interface{} `json:"value"`
}

type _OverviewMetricValue OverviewMetricValue

// NewOverviewMetricValue instantiates a new OverviewMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverviewMetricValue(value []map[string]interface{}) *OverviewMetricValue {
	this := OverviewMetricValue{}
	this.Value = value
	return &this
}

// NewOverviewMetricValueWithDefaults instantiates a new OverviewMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverviewMetricValueWithDefaults() *OverviewMetricValue {
	this := OverviewMetricValue{}
	return &this
}

// GetValue returns the Value field value
func (o *OverviewMetricValue) GetValue() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OverviewMetricValue) GetValueOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *OverviewMetricValue) SetValue(v []map[string]interface{}) {
	o.Value = v
}

func (o OverviewMetricValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverviewMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *OverviewMetricValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varOverviewMetricValue := _OverviewMetricValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOverviewMetricValue)

	if err != nil {
		return err
	}

	*o = OverviewMetricValue(varOverviewMetricValue)

	return err
}

type NullableOverviewMetricValue struct {
	value *OverviewMetricValue
	isSet bool
}

func (v NullableOverviewMetricValue) Get() *OverviewMetricValue {
	return v.value
}

func (v *NullableOverviewMetricValue) Set(val *OverviewMetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullableOverviewMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullableOverviewMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverviewMetricValue(val *OverviewMetricValue) *NullableOverviewMetricValue {
	return &NullableOverviewMetricValue{value: val, isSet: true}
}

func (v NullableOverviewMetricValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverviewMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


