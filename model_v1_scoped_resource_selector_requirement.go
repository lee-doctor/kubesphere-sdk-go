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

// checks if the V1ScopedResourceSelectorRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ScopedResourceSelectorRequirement{}

// V1ScopedResourceSelectorRequirement A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates the scope name and values.
type V1ScopedResourceSelectorRequirement struct {
	// Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
	Operator string `json:"operator"`
	// The name of the scope that the selector applies to.
	ScopeName string `json:"scopeName"`
	// An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values []string `json:"values,omitempty"`
}

type _V1ScopedResourceSelectorRequirement V1ScopedResourceSelectorRequirement

// NewV1ScopedResourceSelectorRequirement instantiates a new V1ScopedResourceSelectorRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ScopedResourceSelectorRequirement(operator string, scopeName string) *V1ScopedResourceSelectorRequirement {
	this := V1ScopedResourceSelectorRequirement{}
	this.Operator = operator
	this.ScopeName = scopeName
	return &this
}

// NewV1ScopedResourceSelectorRequirementWithDefaults instantiates a new V1ScopedResourceSelectorRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ScopedResourceSelectorRequirementWithDefaults() *V1ScopedResourceSelectorRequirement {
	this := V1ScopedResourceSelectorRequirement{}
	return &this
}

// GetOperator returns the Operator field value
func (o *V1ScopedResourceSelectorRequirement) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *V1ScopedResourceSelectorRequirement) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *V1ScopedResourceSelectorRequirement) SetOperator(v string) {
	o.Operator = v
}

// GetScopeName returns the ScopeName field value
func (o *V1ScopedResourceSelectorRequirement) GetScopeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScopeName
}

// GetScopeNameOk returns a tuple with the ScopeName field value
// and a boolean to check if the value has been set.
func (o *V1ScopedResourceSelectorRequirement) GetScopeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeName, true
}

// SetScopeName sets field value
func (o *V1ScopedResourceSelectorRequirement) SetScopeName(v string) {
	o.ScopeName = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *V1ScopedResourceSelectorRequirement) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ScopedResourceSelectorRequirement) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *V1ScopedResourceSelectorRequirement) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *V1ScopedResourceSelectorRequirement) SetValues(v []string) {
	o.Values = v
}

func (o V1ScopedResourceSelectorRequirement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ScopedResourceSelectorRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operator"] = o.Operator
	toSerialize["scopeName"] = o.ScopeName
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

func (o *V1ScopedResourceSelectorRequirement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operator",
		"scopeName",
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

	varV1ScopedResourceSelectorRequirement := _V1ScopedResourceSelectorRequirement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ScopedResourceSelectorRequirement)

	if err != nil {
		return err
	}

	*o = V1ScopedResourceSelectorRequirement(varV1ScopedResourceSelectorRequirement)

	return err
}

type NullableV1ScopedResourceSelectorRequirement struct {
	value *V1ScopedResourceSelectorRequirement
	isSet bool
}

func (v NullableV1ScopedResourceSelectorRequirement) Get() *V1ScopedResourceSelectorRequirement {
	return v.value
}

func (v *NullableV1ScopedResourceSelectorRequirement) Set(val *V1ScopedResourceSelectorRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ScopedResourceSelectorRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ScopedResourceSelectorRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ScopedResourceSelectorRequirement(val *V1ScopedResourceSelectorRequirement) *NullableV1ScopedResourceSelectorRequirement {
	return &NullableV1ScopedResourceSelectorRequirement{value: val, isSet: true}
}

func (v NullableV1ScopedResourceSelectorRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ScopedResourceSelectorRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


