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

// checks if the Trends type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Trends{}

// Trends struct for Trends
type Trends struct {
	Class *string `json:"_class,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewTrends instantiates a new Trends object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrends() *Trends {
	this := Trends{}
	return &this
}

// NewTrendsWithDefaults instantiates a new Trends object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrendsWithDefaults() *Trends {
	this := Trends{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Trends) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trends) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Trends) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Trends) SetClass(v string) {
	o.Class = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Trends) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trends) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Trends) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Trends) SetHref(v string) {
	o.Href = &v
}

func (o Trends) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trends) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["_class"] = o.Class
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableTrends struct {
	value *Trends
	isSet bool
}

func (v NullableTrends) Get() *Trends {
	return v.value
}

func (v *NullableTrends) Set(val *Trends) {
	v.value = val
	v.isSet = true
}

func (v NullableTrends) IsSet() bool {
	return v.isSet
}

func (v *NullableTrends) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrends(val *Trends) *NullableTrends {
	return &NullableTrends{value: val, isSet: true}
}

func (v NullableTrends) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrends) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

