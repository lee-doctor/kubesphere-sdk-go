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
	"time"
)

// checks if the RegistriesHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistriesHistory{}

// RegistriesHistory struct for RegistriesHistory
type RegistriesHistory struct {
	// Created time.
	Created *time.Time `json:"created,omitempty"`
	// Created command.
	CreatedBy *string `json:"created_by,omitempty"`
	// Layer empty or not.
	EmptyLayer *bool `json:"empty_layer,omitempty"`
}

// NewRegistriesHistory instantiates a new RegistriesHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistriesHistory() *RegistriesHistory {
	this := RegistriesHistory{}
	return &this
}

// NewRegistriesHistoryWithDefaults instantiates a new RegistriesHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistriesHistoryWithDefaults() *RegistriesHistory {
	this := RegistriesHistory{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RegistriesHistory) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesHistory) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RegistriesHistory) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RegistriesHistory) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RegistriesHistory) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesHistory) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RegistriesHistory) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *RegistriesHistory) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEmptyLayer returns the EmptyLayer field value if set, zero value otherwise.
func (o *RegistriesHistory) GetEmptyLayer() bool {
	if o == nil || IsNil(o.EmptyLayer) {
		var ret bool
		return ret
	}
	return *o.EmptyLayer
}

// GetEmptyLayerOk returns a tuple with the EmptyLayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesHistory) GetEmptyLayerOk() (*bool, bool) {
	if o == nil || IsNil(o.EmptyLayer) {
		return nil, false
	}
	return o.EmptyLayer, true
}

// HasEmptyLayer returns a boolean if a field has been set.
func (o *RegistriesHistory) HasEmptyLayer() bool {
	if o != nil && !IsNil(o.EmptyLayer) {
		return true
	}

	return false
}

// SetEmptyLayer gets a reference to the given bool and assigns it to the EmptyLayer field.
func (o *RegistriesHistory) SetEmptyLayer(v bool) {
	o.EmptyLayer = &v
}

func (o RegistriesHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistriesHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.EmptyLayer) {
		toSerialize["empty_layer"] = o.EmptyLayer
	}
	return toSerialize, nil
}

type NullableRegistriesHistory struct {
	value *RegistriesHistory
	isSet bool
}

func (v NullableRegistriesHistory) Get() *RegistriesHistory {
	return v.value
}

func (v *NullableRegistriesHistory) Set(val *RegistriesHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistriesHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistriesHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistriesHistory(val *RegistriesHistory) *NullableRegistriesHistory {
	return &NullableRegistriesHistory{value: val, isSet: true}
}

func (v NullableRegistriesHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistriesHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


