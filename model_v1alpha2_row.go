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

// checks if the V1alpha2Row type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2Row{}

// V1alpha2Row struct for V1alpha2Row
type V1alpha2Row struct {
	Entries map[string]string `json:"entries"`
	Id string `json:"id"`
}

type _V1alpha2Row V1alpha2Row

// NewV1alpha2Row instantiates a new V1alpha2Row object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2Row(entries map[string]string, id string) *V1alpha2Row {
	this := V1alpha2Row{}
	this.Entries = entries
	this.Id = id
	return &this
}

// NewV1alpha2RowWithDefaults instantiates a new V1alpha2Row object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2RowWithDefaults() *V1alpha2Row {
	this := V1alpha2Row{}
	return &this
}

// GetEntries returns the Entries field value
func (o *V1alpha2Row) GetEntries() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Row) GetEntriesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entries, true
}

// SetEntries sets field value
func (o *V1alpha2Row) SetEntries(v map[string]string) {
	o.Entries = v
}

// GetId returns the Id field value
func (o *V1alpha2Row) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Row) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V1alpha2Row) SetId(v string) {
	o.Id = v
}

func (o V1alpha2Row) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2Row) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entries"] = o.Entries
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *V1alpha2Row) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entries",
		"id",
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

	varV1alpha2Row := _V1alpha2Row{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2Row)

	if err != nil {
		return err
	}

	*o = V1alpha2Row(varV1alpha2Row)

	return err
}

type NullableV1alpha2Row struct {
	value *V1alpha2Row
	isSet bool
}

func (v NullableV1alpha2Row) Get() *V1alpha2Row {
	return v.value
}

func (v *NullableV1alpha2Row) Set(val *V1alpha2Row) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2Row) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2Row) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2Row(val *V1alpha2Row) *NullableV1alpha2Row {
	return &NullableV1alpha2Row{value: val, isSet: true}
}

func (v NullableV1alpha2Row) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2Row) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


