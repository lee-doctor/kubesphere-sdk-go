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

// checks if the V1EnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EnvVar{}

// V1EnvVar EnvVar represents an environment variable present in a Container.
type V1EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name"`
	// Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. \"$$(VAR_NAME)\" will produce the string literal \"$(VAR_NAME)\". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\".
	Value *string `json:"value,omitempty"`
	ValueFrom *V1EnvVarSource `json:"valueFrom,omitempty"`
}

type _V1EnvVar V1EnvVar

// NewV1EnvVar instantiates a new V1EnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EnvVar(name string) *V1EnvVar {
	this := V1EnvVar{}
	this.Name = name
	return &this
}

// NewV1EnvVarWithDefaults instantiates a new V1EnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EnvVarWithDefaults() *V1EnvVar {
	this := V1EnvVar{}
	return &this
}

// GetName returns the Name field value
func (o *V1EnvVar) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1EnvVar) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1EnvVar) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1EnvVar) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnvVar) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1EnvVar) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V1EnvVar) SetValue(v string) {
	o.Value = &v
}

// GetValueFrom returns the ValueFrom field value if set, zero value otherwise.
func (o *V1EnvVar) GetValueFrom() V1EnvVarSource {
	if o == nil || IsNil(o.ValueFrom) {
		var ret V1EnvVarSource
		return ret
	}
	return *o.ValueFrom
}

// GetValueFromOk returns a tuple with the ValueFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnvVar) GetValueFromOk() (*V1EnvVarSource, bool) {
	if o == nil || IsNil(o.ValueFrom) {
		return nil, false
	}
	return o.ValueFrom, true
}

// HasValueFrom returns a boolean if a field has been set.
func (o *V1EnvVar) HasValueFrom() bool {
	if o != nil && !IsNil(o.ValueFrom) {
		return true
	}

	return false
}

// SetValueFrom gets a reference to the given V1EnvVarSource and assigns it to the ValueFrom field.
func (o *V1EnvVar) SetValueFrom(v V1EnvVarSource) {
	o.ValueFrom = &v
}

func (o V1EnvVar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueFrom) {
		toSerialize["valueFrom"] = o.ValueFrom
	}
	return toSerialize, nil
}

func (o *V1EnvVar) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varV1EnvVar := _V1EnvVar{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1EnvVar)

	if err != nil {
		return err
	}

	*o = V1EnvVar(varV1EnvVar)

	return err
}

type NullableV1EnvVar struct {
	value *V1EnvVar
	isSet bool
}

func (v NullableV1EnvVar) Get() *V1EnvVar {
	return v.value
}

func (v *NullableV1EnvVar) Set(val *V1EnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EnvVar(val *V1EnvVar) *NullableV1EnvVar {
	return &NullableV1EnvVar{value: val, isSet: true}
}

func (v NullableV1EnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


