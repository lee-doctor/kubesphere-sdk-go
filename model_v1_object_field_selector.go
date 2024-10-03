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

// checks if the V1ObjectFieldSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ObjectFieldSelector{}

// V1ObjectFieldSelector ObjectFieldSelector selects an APIVersioned field of an object.
type V1ObjectFieldSelector struct {
	// Version of the schema the FieldPath is written in terms of, defaults to \"v1\".
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Path of the field to select in the specified API version.
	FieldPath string `json:"fieldPath"`
}

type _V1ObjectFieldSelector V1ObjectFieldSelector

// NewV1ObjectFieldSelector instantiates a new V1ObjectFieldSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ObjectFieldSelector(fieldPath string) *V1ObjectFieldSelector {
	this := V1ObjectFieldSelector{}
	this.FieldPath = fieldPath
	return &this
}

// NewV1ObjectFieldSelectorWithDefaults instantiates a new V1ObjectFieldSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ObjectFieldSelectorWithDefaults() *V1ObjectFieldSelector {
	this := V1ObjectFieldSelector{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1ObjectFieldSelector) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ObjectFieldSelector) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1ObjectFieldSelector) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1ObjectFieldSelector) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetFieldPath returns the FieldPath field value
func (o *V1ObjectFieldSelector) GetFieldPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldPath
}

// GetFieldPathOk returns a tuple with the FieldPath field value
// and a boolean to check if the value has been set.
func (o *V1ObjectFieldSelector) GetFieldPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldPath, true
}

// SetFieldPath sets field value
func (o *V1ObjectFieldSelector) SetFieldPath(v string) {
	o.FieldPath = v
}

func (o V1ObjectFieldSelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ObjectFieldSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	toSerialize["fieldPath"] = o.FieldPath
	return toSerialize, nil
}

func (o *V1ObjectFieldSelector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldPath",
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

	varV1ObjectFieldSelector := _V1ObjectFieldSelector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ObjectFieldSelector)

	if err != nil {
		return err
	}

	*o = V1ObjectFieldSelector(varV1ObjectFieldSelector)

	return err
}

type NullableV1ObjectFieldSelector struct {
	value *V1ObjectFieldSelector
	isSet bool
}

func (v NullableV1ObjectFieldSelector) Get() *V1ObjectFieldSelector {
	return v.value
}

func (v *NullableV1ObjectFieldSelector) Set(val *V1ObjectFieldSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ObjectFieldSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ObjectFieldSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ObjectFieldSelector(val *V1ObjectFieldSelector) *NullableV1ObjectFieldSelector {
	return &NullableV1ObjectFieldSelector{value: val, isSet: true}
}

func (v NullableV1ObjectFieldSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ObjectFieldSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


