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

// checks if the V1alpha3Parameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha3Parameter{}

// V1alpha3Parameter struct for V1alpha3Parameter
type V1alpha3Parameter struct {
	// default value of param
	DefaultValue *string `json:"default_value,omitempty"`
	// description of pipeline
	Description *string `json:"description,omitempty"`
	// name of param
	Name string `json:"name"`
	// type of param
	Type string `json:"type"`
}

type _V1alpha3Parameter V1alpha3Parameter

// NewV1alpha3Parameter instantiates a new V1alpha3Parameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha3Parameter(name string, type_ string) *V1alpha3Parameter {
	this := V1alpha3Parameter{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewV1alpha3ParameterWithDefaults instantiates a new V1alpha3Parameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha3ParameterWithDefaults() *V1alpha3Parameter {
	this := V1alpha3Parameter{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *V1alpha3Parameter) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha3Parameter) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *V1alpha3Parameter) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *V1alpha3Parameter) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1alpha3Parameter) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha3Parameter) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1alpha3Parameter) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1alpha3Parameter) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *V1alpha3Parameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1alpha3Parameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1alpha3Parameter) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *V1alpha3Parameter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1alpha3Parameter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1alpha3Parameter) SetType(v string) {
	o.Type = v
}

func (o V1alpha3Parameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha3Parameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultValue) {
		toSerialize["default_value"] = o.DefaultValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *V1alpha3Parameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varV1alpha3Parameter := _V1alpha3Parameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha3Parameter)

	if err != nil {
		return err
	}

	*o = V1alpha3Parameter(varV1alpha3Parameter)

	return err
}

type NullableV1alpha3Parameter struct {
	value *V1alpha3Parameter
	isSet bool
}

func (v NullableV1alpha3Parameter) Get() *V1alpha3Parameter {
	return v.value
}

func (v *NullableV1alpha3Parameter) Set(val *V1alpha3Parameter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha3Parameter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha3Parameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha3Parameter(val *V1alpha3Parameter) *NullableV1alpha3Parameter {
	return &NullableV1alpha3Parameter{value: val, isSet: true}
}

func (v NullableV1alpha3Parameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha3Parameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


