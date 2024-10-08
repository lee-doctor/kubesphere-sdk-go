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

// checks if the V1RootFS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RootFS{}

// V1RootFS struct for V1RootFS
type V1RootFS struct {
	DiffIds []V1Hash `json:"diff_ids"`
	Type string `json:"type"`
}

type _V1RootFS V1RootFS

// NewV1RootFS instantiates a new V1RootFS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RootFS(diffIds []V1Hash, type_ string) *V1RootFS {
	this := V1RootFS{}
	this.DiffIds = diffIds
	this.Type = type_
	return &this
}

// NewV1RootFSWithDefaults instantiates a new V1RootFS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RootFSWithDefaults() *V1RootFS {
	this := V1RootFS{}
	return &this
}

// GetDiffIds returns the DiffIds field value
func (o *V1RootFS) GetDiffIds() []V1Hash {
	if o == nil {
		var ret []V1Hash
		return ret
	}

	return o.DiffIds
}

// GetDiffIdsOk returns a tuple with the DiffIds field value
// and a boolean to check if the value has been set.
func (o *V1RootFS) GetDiffIdsOk() ([]V1Hash, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiffIds, true
}

// SetDiffIds sets field value
func (o *V1RootFS) SetDiffIds(v []V1Hash) {
	o.DiffIds = v
}

// GetType returns the Type field value
func (o *V1RootFS) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1RootFS) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1RootFS) SetType(v string) {
	o.Type = v
}

func (o V1RootFS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RootFS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["diff_ids"] = o.DiffIds
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *V1RootFS) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"diff_ids",
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

	varV1RootFS := _V1RootFS{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1RootFS)

	if err != nil {
		return err
	}

	*o = V1RootFS(varV1RootFS)

	return err
}

type NullableV1RootFS struct {
	value *V1RootFS
	isSet bool
}

func (v NullableV1RootFS) Get() *V1RootFS {
	return v.value
}

func (v *NullableV1RootFS) Set(val *V1RootFS) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RootFS) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RootFS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RootFS(val *V1RootFS) *NullableV1RootFS {
	return &NullableV1RootFS{value: val, isSet: true}
}

func (v NullableV1RootFS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RootFS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


