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

// checks if the DevopsStopPipelineCauses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsStopPipelineCauses{}

// DevopsStopPipelineCauses struct for DevopsStopPipelineCauses
type DevopsStopPipelineCauses struct {
	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class *string `json:"_class,omitempty"`
	// short description
	ShortDescription *string `json:"shortDescription,omitempty"`
}

// NewDevopsStopPipelineCauses instantiates a new DevopsStopPipelineCauses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsStopPipelineCauses() *DevopsStopPipelineCauses {
	this := DevopsStopPipelineCauses{}
	return &this
}

// NewDevopsStopPipelineCausesWithDefaults instantiates a new DevopsStopPipelineCauses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsStopPipelineCausesWithDefaults() *DevopsStopPipelineCauses {
	this := DevopsStopPipelineCauses{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *DevopsStopPipelineCauses) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipelineCauses) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *DevopsStopPipelineCauses) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *DevopsStopPipelineCauses) SetClass(v string) {
	o.Class = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *DevopsStopPipelineCauses) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsStopPipelineCauses) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *DevopsStopPipelineCauses) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *DevopsStopPipelineCauses) SetShortDescription(v string) {
	o.ShortDescription = &v
}

func (o DevopsStopPipelineCauses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsStopPipelineCauses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["_class"] = o.Class
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	return toSerialize, nil
}

type NullableDevopsStopPipelineCauses struct {
	value *DevopsStopPipelineCauses
	isSet bool
}

func (v NullableDevopsStopPipelineCauses) Get() *DevopsStopPipelineCauses {
	return v.value
}

func (v *NullableDevopsStopPipelineCauses) Set(val *DevopsStopPipelineCauses) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsStopPipelineCauses) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsStopPipelineCauses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsStopPipelineCauses(val *DevopsStopPipelineCauses) *NullableDevopsStopPipelineCauses {
	return &NullableDevopsStopPipelineCauses{value: val, isSet: true}
}

func (v NullableDevopsStopPipelineCauses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsStopPipelineCauses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

