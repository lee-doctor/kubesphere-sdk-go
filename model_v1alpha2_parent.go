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

// checks if the V1alpha2Parent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2Parent{}

// V1alpha2Parent struct for V1alpha2Parent
type V1alpha2Parent struct {
	Id string `json:"id"`
	Label string `json:"label"`
	TopologyId string `json:"topologyId"`
}

type _V1alpha2Parent V1alpha2Parent

// NewV1alpha2Parent instantiates a new V1alpha2Parent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2Parent(id string, label string, topologyId string) *V1alpha2Parent {
	this := V1alpha2Parent{}
	this.Id = id
	this.Label = label
	this.TopologyId = topologyId
	return &this
}

// NewV1alpha2ParentWithDefaults instantiates a new V1alpha2Parent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2ParentWithDefaults() *V1alpha2Parent {
	this := V1alpha2Parent{}
	return &this
}

// GetId returns the Id field value
func (o *V1alpha2Parent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Parent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V1alpha2Parent) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *V1alpha2Parent) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Parent) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *V1alpha2Parent) SetLabel(v string) {
	o.Label = v
}

// GetTopologyId returns the TopologyId field value
func (o *V1alpha2Parent) GetTopologyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyId
}

// GetTopologyIdOk returns a tuple with the TopologyId field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Parent) GetTopologyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyId, true
}

// SetTopologyId sets field value
func (o *V1alpha2Parent) SetTopologyId(v string) {
	o.TopologyId = v
}

func (o V1alpha2Parent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2Parent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	toSerialize["topologyId"] = o.TopologyId
	return toSerialize, nil
}

func (o *V1alpha2Parent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"label",
		"topologyId",
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

	varV1alpha2Parent := _V1alpha2Parent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2Parent)

	if err != nil {
		return err
	}

	*o = V1alpha2Parent(varV1alpha2Parent)

	return err
}

type NullableV1alpha2Parent struct {
	value *V1alpha2Parent
	isSet bool
}

func (v NullableV1alpha2Parent) Get() *V1alpha2Parent {
	return v.value
}

func (v *NullableV1alpha2Parent) Set(val *V1alpha2Parent) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2Parent) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2Parent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2Parent(val *V1alpha2Parent) *NullableV1alpha2Parent {
	return &NullableV1alpha2Parent{value: val, isSet: true}
}

func (v NullableV1alpha2Parent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2Parent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

