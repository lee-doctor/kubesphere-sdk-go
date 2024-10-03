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

// checks if the V1TypedLocalObjectReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1TypedLocalObjectReference{}

// V1TypedLocalObjectReference TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
type V1TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	ApiGroup string `json:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name"`
}

type _V1TypedLocalObjectReference V1TypedLocalObjectReference

// NewV1TypedLocalObjectReference instantiates a new V1TypedLocalObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TypedLocalObjectReference(apiGroup string, kind string, name string) *V1TypedLocalObjectReference {
	this := V1TypedLocalObjectReference{}
	this.ApiGroup = apiGroup
	this.Kind = kind
	this.Name = name
	return &this
}

// NewV1TypedLocalObjectReferenceWithDefaults instantiates a new V1TypedLocalObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TypedLocalObjectReferenceWithDefaults() *V1TypedLocalObjectReference {
	this := V1TypedLocalObjectReference{}
	return &this
}

// GetApiGroup returns the ApiGroup field value
func (o *V1TypedLocalObjectReference) GetApiGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value
// and a boolean to check if the value has been set.
func (o *V1TypedLocalObjectReference) GetApiGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGroup, true
}

// SetApiGroup sets field value
func (o *V1TypedLocalObjectReference) SetApiGroup(v string) {
	o.ApiGroup = v
}

// GetKind returns the Kind field value
func (o *V1TypedLocalObjectReference) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *V1TypedLocalObjectReference) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *V1TypedLocalObjectReference) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *V1TypedLocalObjectReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1TypedLocalObjectReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1TypedLocalObjectReference) SetName(v string) {
	o.Name = v
}

func (o V1TypedLocalObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1TypedLocalObjectReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiGroup"] = o.ApiGroup
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *V1TypedLocalObjectReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiGroup",
		"kind",
		"name",
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

	varV1TypedLocalObjectReference := _V1TypedLocalObjectReference{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1TypedLocalObjectReference)

	if err != nil {
		return err
	}

	*o = V1TypedLocalObjectReference(varV1TypedLocalObjectReference)

	return err
}

type NullableV1TypedLocalObjectReference struct {
	value *V1TypedLocalObjectReference
	isSet bool
}

func (v NullableV1TypedLocalObjectReference) Get() *V1TypedLocalObjectReference {
	return v.value
}

func (v *NullableV1TypedLocalObjectReference) Set(val *V1TypedLocalObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TypedLocalObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TypedLocalObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TypedLocalObjectReference(val *V1TypedLocalObjectReference) *NullableV1TypedLocalObjectReference {
	return &NullableV1TypedLocalObjectReference{value: val, isSet: true}
}

func (v NullableV1TypedLocalObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TypedLocalObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


