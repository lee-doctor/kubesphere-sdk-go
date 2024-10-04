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

// checks if the V1ResourceFieldSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ResourceFieldSelector{}

// V1ResourceFieldSelector ResourceFieldSelector represents container resources (cpu, memory) and their output format
type V1ResourceFieldSelector struct {
	// Container name: required for volumes, optional for env vars
	ContainerName *string `json:"containerName,omitempty"`
	// Specifies the output format of the exposed resources, defaults to \"1\"
	Divisor *string `json:"divisor,omitempty"`
	// Required: resource to select
	Resource string `json:"resource"`
}

type _V1ResourceFieldSelector V1ResourceFieldSelector

// NewV1ResourceFieldSelector instantiates a new V1ResourceFieldSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ResourceFieldSelector(resource string) *V1ResourceFieldSelector {
	this := V1ResourceFieldSelector{}
	this.Resource = resource
	return &this
}

// NewV1ResourceFieldSelectorWithDefaults instantiates a new V1ResourceFieldSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ResourceFieldSelectorWithDefaults() *V1ResourceFieldSelector {
	this := V1ResourceFieldSelector{}
	return &this
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *V1ResourceFieldSelector) GetContainerName() string {
	if o == nil || IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ResourceFieldSelector) GetContainerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *V1ResourceFieldSelector) HasContainerName() bool {
	if o != nil && !IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *V1ResourceFieldSelector) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetDivisor returns the Divisor field value if set, zero value otherwise.
func (o *V1ResourceFieldSelector) GetDivisor() string {
	if o == nil || IsNil(o.Divisor) {
		var ret string
		return ret
	}
	return *o.Divisor
}

// GetDivisorOk returns a tuple with the Divisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ResourceFieldSelector) GetDivisorOk() (*string, bool) {
	if o == nil || IsNil(o.Divisor) {
		return nil, false
	}
	return o.Divisor, true
}

// HasDivisor returns a boolean if a field has been set.
func (o *V1ResourceFieldSelector) HasDivisor() bool {
	if o != nil && !IsNil(o.Divisor) {
		return true
	}

	return false
}

// SetDivisor gets a reference to the given string and assigns it to the Divisor field.
func (o *V1ResourceFieldSelector) SetDivisor(v string) {
	o.Divisor = &v
}

// GetResource returns the Resource field value
func (o *V1ResourceFieldSelector) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *V1ResourceFieldSelector) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *V1ResourceFieldSelector) SetResource(v string) {
	o.Resource = v
}

func (o V1ResourceFieldSelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ResourceFieldSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerName) {
		toSerialize["containerName"] = o.ContainerName
	}
	if !IsNil(o.Divisor) {
		toSerialize["divisor"] = o.Divisor
	}
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

func (o *V1ResourceFieldSelector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource",
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

	varV1ResourceFieldSelector := _V1ResourceFieldSelector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ResourceFieldSelector)

	if err != nil {
		return err
	}

	*o = V1ResourceFieldSelector(varV1ResourceFieldSelector)

	return err
}

type NullableV1ResourceFieldSelector struct {
	value *V1ResourceFieldSelector
	isSet bool
}

func (v NullableV1ResourceFieldSelector) Get() *V1ResourceFieldSelector {
	return v.value
}

func (v *NullableV1ResourceFieldSelector) Set(val *V1ResourceFieldSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ResourceFieldSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ResourceFieldSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ResourceFieldSelector(val *V1ResourceFieldSelector) *NullableV1ResourceFieldSelector {
	return &NullableV1ResourceFieldSelector{value: val, isSet: true}
}

func (v NullableV1ResourceFieldSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ResourceFieldSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


