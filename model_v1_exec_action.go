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

// checks if the V1ExecAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ExecAction{}

// V1ExecAction ExecAction describes a \"run in container\" action.
type V1ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	Command []string `json:"command,omitempty"`
}

// NewV1ExecAction instantiates a new V1ExecAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ExecAction() *V1ExecAction {
	this := V1ExecAction{}
	return &this
}

// NewV1ExecActionWithDefaults instantiates a new V1ExecAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ExecActionWithDefaults() *V1ExecAction {
	this := V1ExecAction{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *V1ExecAction) GetCommand() []string {
	if o == nil || IsNil(o.Command) {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ExecAction) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *V1ExecAction) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *V1ExecAction) SetCommand(v []string) {
	o.Command = v
}

func (o V1ExecAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ExecAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	return toSerialize, nil
}

type NullableV1ExecAction struct {
	value *V1ExecAction
	isSet bool
}

func (v NullableV1ExecAction) Get() *V1ExecAction {
	return v.value
}

func (v *NullableV1ExecAction) Set(val *V1ExecAction) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ExecAction) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ExecAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ExecAction(val *V1ExecAction) *NullableV1ExecAction {
	return &NullableV1ExecAction{value: val, isSet: true}
}

func (v NullableV1ExecAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ExecAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


