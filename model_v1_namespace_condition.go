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

// checks if the V1NamespaceCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1NamespaceCondition{}

// V1NamespaceCondition NamespaceCondition contains details about state of namespace.
type V1NamespaceCondition struct {
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	Message *string `json:"message,omitempty"`
	Reason *string `json:"reason,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// Type of namespace controller condition.
	Type string `json:"type"`
}

type _V1NamespaceCondition V1NamespaceCondition

// NewV1NamespaceCondition instantiates a new V1NamespaceCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NamespaceCondition(status string, type_ string) *V1NamespaceCondition {
	this := V1NamespaceCondition{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewV1NamespaceConditionWithDefaults instantiates a new V1NamespaceCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NamespaceConditionWithDefaults() *V1NamespaceCondition {
	this := V1NamespaceCondition{}
	return &this
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *V1NamespaceCondition) GetLastTransitionTime() string {
	if o == nil || IsNil(o.LastTransitionTime) {
		var ret string
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamespaceCondition) GetLastTransitionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastTransitionTime) {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *V1NamespaceCondition) HasLastTransitionTime() bool {
	if o != nil && !IsNil(o.LastTransitionTime) {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given string and assigns it to the LastTransitionTime field.
func (o *V1NamespaceCondition) SetLastTransitionTime(v string) {
	o.LastTransitionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1NamespaceCondition) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamespaceCondition) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1NamespaceCondition) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1NamespaceCondition) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1NamespaceCondition) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NamespaceCondition) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1NamespaceCondition) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1NamespaceCondition) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value
func (o *V1NamespaceCondition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V1NamespaceCondition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V1NamespaceCondition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *V1NamespaceCondition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1NamespaceCondition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1NamespaceCondition) SetType(v string) {
	o.Type = v
}

func (o V1NamespaceCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1NamespaceCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastTransitionTime) {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *V1NamespaceCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"type",
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

	varV1NamespaceCondition := _V1NamespaceCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1NamespaceCondition)

	if err != nil {
		return err
	}

	*o = V1NamespaceCondition(varV1NamespaceCondition)

	return err
}

type NullableV1NamespaceCondition struct {
	value *V1NamespaceCondition
	isSet bool
}

func (v NullableV1NamespaceCondition) Get() *V1NamespaceCondition {
	return v.value
}

func (v *NullableV1NamespaceCondition) Set(val *V1NamespaceCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NamespaceCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NamespaceCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NamespaceCondition(val *V1NamespaceCondition) *NullableV1NamespaceCondition {
	return &NullableV1NamespaceCondition{value: val, isSet: true}
}

func (v NullableV1NamespaceCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NamespaceCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


