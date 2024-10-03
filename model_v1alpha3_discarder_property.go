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

// checks if the V1alpha3DiscarderProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha3DiscarderProperty{}

// V1alpha3DiscarderProperty struct for V1alpha3DiscarderProperty
type V1alpha3DiscarderProperty struct {
	// days to keep pipeline
	DaysToKeep *string `json:"days_to_keep,omitempty"`
	// nums to keep pipeline
	NumToKeep *string `json:"num_to_keep,omitempty"`
}

// NewV1alpha3DiscarderProperty instantiates a new V1alpha3DiscarderProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha3DiscarderProperty() *V1alpha3DiscarderProperty {
	this := V1alpha3DiscarderProperty{}
	return &this
}

// NewV1alpha3DiscarderPropertyWithDefaults instantiates a new V1alpha3DiscarderProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha3DiscarderPropertyWithDefaults() *V1alpha3DiscarderProperty {
	this := V1alpha3DiscarderProperty{}
	return &this
}

// GetDaysToKeep returns the DaysToKeep field value if set, zero value otherwise.
func (o *V1alpha3DiscarderProperty) GetDaysToKeep() string {
	if o == nil || IsNil(o.DaysToKeep) {
		var ret string
		return ret
	}
	return *o.DaysToKeep
}

// GetDaysToKeepOk returns a tuple with the DaysToKeep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha3DiscarderProperty) GetDaysToKeepOk() (*string, bool) {
	if o == nil || IsNil(o.DaysToKeep) {
		return nil, false
	}
	return o.DaysToKeep, true
}

// HasDaysToKeep returns a boolean if a field has been set.
func (o *V1alpha3DiscarderProperty) HasDaysToKeep() bool {
	if o != nil && !IsNil(o.DaysToKeep) {
		return true
	}

	return false
}

// SetDaysToKeep gets a reference to the given string and assigns it to the DaysToKeep field.
func (o *V1alpha3DiscarderProperty) SetDaysToKeep(v string) {
	o.DaysToKeep = &v
}

// GetNumToKeep returns the NumToKeep field value if set, zero value otherwise.
func (o *V1alpha3DiscarderProperty) GetNumToKeep() string {
	if o == nil || IsNil(o.NumToKeep) {
		var ret string
		return ret
	}
	return *o.NumToKeep
}

// GetNumToKeepOk returns a tuple with the NumToKeep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha3DiscarderProperty) GetNumToKeepOk() (*string, bool) {
	if o == nil || IsNil(o.NumToKeep) {
		return nil, false
	}
	return o.NumToKeep, true
}

// HasNumToKeep returns a boolean if a field has been set.
func (o *V1alpha3DiscarderProperty) HasNumToKeep() bool {
	if o != nil && !IsNil(o.NumToKeep) {
		return true
	}

	return false
}

// SetNumToKeep gets a reference to the given string and assigns it to the NumToKeep field.
func (o *V1alpha3DiscarderProperty) SetNumToKeep(v string) {
	o.NumToKeep = &v
}

func (o V1alpha3DiscarderProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha3DiscarderProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DaysToKeep) {
		toSerialize["days_to_keep"] = o.DaysToKeep
	}
	if !IsNil(o.NumToKeep) {
		toSerialize["num_to_keep"] = o.NumToKeep
	}
	return toSerialize, nil
}

type NullableV1alpha3DiscarderProperty struct {
	value *V1alpha3DiscarderProperty
	isSet bool
}

func (v NullableV1alpha3DiscarderProperty) Get() *V1alpha3DiscarderProperty {
	return v.value
}

func (v *NullableV1alpha3DiscarderProperty) Set(val *V1alpha3DiscarderProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha3DiscarderProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha3DiscarderProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha3DiscarderProperty(val *V1alpha3DiscarderProperty) *NullableV1alpha3DiscarderProperty {
	return &NullableV1alpha3DiscarderProperty{value: val, isSet: true}
}

func (v NullableV1alpha3DiscarderProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha3DiscarderProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

