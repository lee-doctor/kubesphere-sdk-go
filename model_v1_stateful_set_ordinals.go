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

// checks if the V1StatefulSetOrdinals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1StatefulSetOrdinals{}

// V1StatefulSetOrdinals StatefulSetOrdinals describes the policy used for replica ordinal assignment in this StatefulSet.
type V1StatefulSetOrdinals struct {
	// start is the number representing the first replica's index. It may be used to number replicas from an alternate index (eg: 1-indexed) over the default 0-indexed names, or to orchestrate progressive movement of replicas from one StatefulSet to another. If set, replica indices will be in the range:   [.spec.ordinals.start, .spec.ordinals.start + .spec.replicas). If unset, defaults to 0. Replica indices will be in the range:   [0, .spec.replicas).
	Start int32 `json:"start"`
}

type _V1StatefulSetOrdinals V1StatefulSetOrdinals

// NewV1StatefulSetOrdinals instantiates a new V1StatefulSetOrdinals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1StatefulSetOrdinals(start int32) *V1StatefulSetOrdinals {
	this := V1StatefulSetOrdinals{}
	this.Start = start
	return &this
}

// NewV1StatefulSetOrdinalsWithDefaults instantiates a new V1StatefulSetOrdinals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1StatefulSetOrdinalsWithDefaults() *V1StatefulSetOrdinals {
	this := V1StatefulSetOrdinals{}
	return &this
}

// GetStart returns the Start field value
func (o *V1StatefulSetOrdinals) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *V1StatefulSetOrdinals) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *V1StatefulSetOrdinals) SetStart(v int32) {
	o.Start = v
}

func (o V1StatefulSetOrdinals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1StatefulSetOrdinals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	return toSerialize, nil
}

func (o *V1StatefulSetOrdinals) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
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

	varV1StatefulSetOrdinals := _V1StatefulSetOrdinals{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1StatefulSetOrdinals)

	if err != nil {
		return err
	}

	*o = V1StatefulSetOrdinals(varV1StatefulSetOrdinals)

	return err
}

type NullableV1StatefulSetOrdinals struct {
	value *V1StatefulSetOrdinals
	isSet bool
}

func (v NullableV1StatefulSetOrdinals) Get() *V1StatefulSetOrdinals {
	return v.value
}

func (v *NullableV1StatefulSetOrdinals) Set(val *V1StatefulSetOrdinals) {
	v.value = val
	v.isSet = true
}

func (v NullableV1StatefulSetOrdinals) IsSet() bool {
	return v.isSet
}

func (v *NullableV1StatefulSetOrdinals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1StatefulSetOrdinals(val *V1StatefulSetOrdinals) *NullableV1StatefulSetOrdinals {
	return &NullableV1StatefulSetOrdinals{value: val, isSet: true}
}

func (v NullableV1StatefulSetOrdinals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1StatefulSetOrdinals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

