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

// checks if the V1NodeSelectorRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1NodeSelectorRequirement{}

// V1NodeSelectorRequirement A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type V1NodeSelectorRequirement struct {
	// The label key that the selector applies to.
	Key string `json:"key"`
	// Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Operator string `json:"operator"`
	// An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
	Values []string `json:"values,omitempty"`
}

type _V1NodeSelectorRequirement V1NodeSelectorRequirement

// NewV1NodeSelectorRequirement instantiates a new V1NodeSelectorRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NodeSelectorRequirement(key string, operator string) *V1NodeSelectorRequirement {
	this := V1NodeSelectorRequirement{}
	this.Key = key
	this.Operator = operator
	return &this
}

// NewV1NodeSelectorRequirementWithDefaults instantiates a new V1NodeSelectorRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NodeSelectorRequirementWithDefaults() *V1NodeSelectorRequirement {
	this := V1NodeSelectorRequirement{}
	return &this
}

// GetKey returns the Key field value
func (o *V1NodeSelectorRequirement) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *V1NodeSelectorRequirement) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *V1NodeSelectorRequirement) SetKey(v string) {
	o.Key = v
}

// GetOperator returns the Operator field value
func (o *V1NodeSelectorRequirement) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *V1NodeSelectorRequirement) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *V1NodeSelectorRequirement) SetOperator(v string) {
	o.Operator = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *V1NodeSelectorRequirement) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NodeSelectorRequirement) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *V1NodeSelectorRequirement) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *V1NodeSelectorRequirement) SetValues(v []string) {
	o.Values = v
}

func (o V1NodeSelectorRequirement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1NodeSelectorRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["operator"] = o.Operator
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

func (o *V1NodeSelectorRequirement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"operator",
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

	varV1NodeSelectorRequirement := _V1NodeSelectorRequirement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1NodeSelectorRequirement)

	if err != nil {
		return err
	}

	*o = V1NodeSelectorRequirement(varV1NodeSelectorRequirement)

	return err
}

type NullableV1NodeSelectorRequirement struct {
	value *V1NodeSelectorRequirement
	isSet bool
}

func (v NullableV1NodeSelectorRequirement) Get() *V1NodeSelectorRequirement {
	return v.value
}

func (v *NullableV1NodeSelectorRequirement) Set(val *V1NodeSelectorRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NodeSelectorRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NodeSelectorRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NodeSelectorRequirement(val *V1NodeSelectorRequirement) *NullableV1NodeSelectorRequirement {
	return &NullableV1NodeSelectorRequirement{value: val, isSet: true}
}

func (v NullableV1NodeSelectorRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NodeSelectorRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


