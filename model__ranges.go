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

// checks if the Ranges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ranges{}

// Ranges struct for Ranges
type Ranges struct {
	Ranges []RangesRanges `json:"ranges,omitempty"`
}

// NewRanges instantiates a new Ranges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRanges() *Ranges {
	this := Ranges{}
	return &this
}

// NewRangesWithDefaults instantiates a new Ranges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangesWithDefaults() *Ranges {
	this := Ranges{}
	return &this
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *Ranges) GetRanges() []RangesRanges {
	if o == nil || IsNil(o.Ranges) {
		var ret []RangesRanges
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ranges) GetRangesOk() ([]RangesRanges, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *Ranges) HasRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []RangesRanges and assigns it to the Ranges field.
func (o *Ranges) SetRanges(v []RangesRanges) {
	o.Ranges = v
}

func (o Ranges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ranges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	return toSerialize, nil
}

type NullableRanges struct {
	value *Ranges
	isSet bool
}

func (v NullableRanges) Get() *Ranges {
	return v.value
}

func (v *NullableRanges) Set(val *Ranges) {
	v.value = val
	v.isSet = true
}

func (v NullableRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRanges(val *Ranges) *NullableRanges {
	return &NullableRanges{value: val, isSet: true}
}

func (v NullableRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


