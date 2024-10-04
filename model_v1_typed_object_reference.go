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

// checks if the V1TypedObjectReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1TypedObjectReference{}

// V1TypedObjectReference struct for V1TypedObjectReference
type V1TypedObjectReference struct {
	// APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	ApiGroup string `json:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name"`
	// Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.
	Namespace *string `json:"namespace,omitempty"`
}

type _V1TypedObjectReference V1TypedObjectReference

// NewV1TypedObjectReference instantiates a new V1TypedObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TypedObjectReference(apiGroup string, kind string, name string) *V1TypedObjectReference {
	this := V1TypedObjectReference{}
	this.ApiGroup = apiGroup
	this.Kind = kind
	this.Name = name
	return &this
}

// NewV1TypedObjectReferenceWithDefaults instantiates a new V1TypedObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TypedObjectReferenceWithDefaults() *V1TypedObjectReference {
	this := V1TypedObjectReference{}
	return &this
}

// GetApiGroup returns the ApiGroup field value
func (o *V1TypedObjectReference) GetApiGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value
// and a boolean to check if the value has been set.
func (o *V1TypedObjectReference) GetApiGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGroup, true
}

// SetApiGroup sets field value
func (o *V1TypedObjectReference) SetApiGroup(v string) {
	o.ApiGroup = v
}

// GetKind returns the Kind field value
func (o *V1TypedObjectReference) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *V1TypedObjectReference) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *V1TypedObjectReference) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *V1TypedObjectReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1TypedObjectReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1TypedObjectReference) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1TypedObjectReference) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TypedObjectReference) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1TypedObjectReference) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1TypedObjectReference) SetNamespace(v string) {
	o.Namespace = &v
}

func (o V1TypedObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1TypedObjectReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiGroup"] = o.ApiGroup
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

func (o *V1TypedObjectReference) UnmarshalJSON(data []byte) (err error) {
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

	varV1TypedObjectReference := _V1TypedObjectReference{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1TypedObjectReference)

	if err != nil {
		return err
	}

	*o = V1TypedObjectReference(varV1TypedObjectReference)

	return err
}

type NullableV1TypedObjectReference struct {
	value *V1TypedObjectReference
	isSet bool
}

func (v NullableV1TypedObjectReference) Get() *V1TypedObjectReference {
	return v.value
}

func (v *NullableV1TypedObjectReference) Set(val *V1TypedObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TypedObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TypedObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TypedObjectReference(val *V1TypedObjectReference) *NullableV1TypedObjectReference {
	return &NullableV1TypedObjectReference{value: val, isSet: true}
}

func (v NullableV1TypedObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TypedObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

