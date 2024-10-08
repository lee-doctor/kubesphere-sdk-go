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

// checks if the V1SleepAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SleepAction{}

// V1SleepAction SleepAction describes a \"sleep\" action.
type V1SleepAction struct {
	// Seconds is the number of seconds to sleep.
	Seconds int64 `json:"seconds"`
}

type _V1SleepAction V1SleepAction

// NewV1SleepAction instantiates a new V1SleepAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SleepAction(seconds int64) *V1SleepAction {
	this := V1SleepAction{}
	this.Seconds = seconds
	return &this
}

// NewV1SleepActionWithDefaults instantiates a new V1SleepAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SleepActionWithDefaults() *V1SleepAction {
	this := V1SleepAction{}
	return &this
}

// GetSeconds returns the Seconds field value
func (o *V1SleepAction) GetSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value
// and a boolean to check if the value has been set.
func (o *V1SleepAction) GetSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seconds, true
}

// SetSeconds sets field value
func (o *V1SleepAction) SetSeconds(v int64) {
	o.Seconds = v
}

func (o V1SleepAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SleepAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seconds"] = o.Seconds
	return toSerialize, nil
}

func (o *V1SleepAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seconds",
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

	varV1SleepAction := _V1SleepAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1SleepAction)

	if err != nil {
		return err
	}

	*o = V1SleepAction(varV1SleepAction)

	return err
}

type NullableV1SleepAction struct {
	value *V1SleepAction
	isSet bool
}

func (v NullableV1SleepAction) Get() *V1SleepAction {
	return v.value
}

func (v *NullableV1SleepAction) Set(val *V1SleepAction) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SleepAction) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SleepAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SleepAction(val *V1SleepAction) *NullableV1SleepAction {
	return &NullableV1SleepAction{value: val, isSet: true}
}

func (v NullableV1SleepAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SleepAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


