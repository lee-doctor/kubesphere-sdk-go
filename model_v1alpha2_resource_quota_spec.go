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

// checks if the V1alpha2ResourceQuotaSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2ResourceQuotaSpec{}

// V1alpha2ResourceQuotaSpec struct for V1alpha2ResourceQuotaSpec
type V1alpha2ResourceQuotaSpec struct {
	Quota V1ResourceQuotaSpec `json:"quota"`
	Selector map[string]string `json:"selector"`
}

type _V1alpha2ResourceQuotaSpec V1alpha2ResourceQuotaSpec

// NewV1alpha2ResourceQuotaSpec instantiates a new V1alpha2ResourceQuotaSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2ResourceQuotaSpec(quota V1ResourceQuotaSpec, selector map[string]string) *V1alpha2ResourceQuotaSpec {
	this := V1alpha2ResourceQuotaSpec{}
	this.Quota = quota
	this.Selector = selector
	return &this
}

// NewV1alpha2ResourceQuotaSpecWithDefaults instantiates a new V1alpha2ResourceQuotaSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2ResourceQuotaSpecWithDefaults() *V1alpha2ResourceQuotaSpec {
	this := V1alpha2ResourceQuotaSpec{}
	return &this
}

// GetQuota returns the Quota field value
func (o *V1alpha2ResourceQuotaSpec) GetQuota() V1ResourceQuotaSpec {
	if o == nil {
		var ret V1ResourceQuotaSpec
		return ret
	}

	return o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ResourceQuotaSpec) GetQuotaOk() (*V1ResourceQuotaSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quota, true
}

// SetQuota sets field value
func (o *V1alpha2ResourceQuotaSpec) SetQuota(v V1ResourceQuotaSpec) {
	o.Quota = v
}

// GetSelector returns the Selector field value
func (o *V1alpha2ResourceQuotaSpec) GetSelector() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ResourceQuotaSpec) GetSelectorOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *V1alpha2ResourceQuotaSpec) SetSelector(v map[string]string) {
	o.Selector = v
}

func (o V1alpha2ResourceQuotaSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2ResourceQuotaSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quota"] = o.Quota
	toSerialize["selector"] = o.Selector
	return toSerialize, nil
}

func (o *V1alpha2ResourceQuotaSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quota",
		"selector",
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

	varV1alpha2ResourceQuotaSpec := _V1alpha2ResourceQuotaSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2ResourceQuotaSpec)

	if err != nil {
		return err
	}

	*o = V1alpha2ResourceQuotaSpec(varV1alpha2ResourceQuotaSpec)

	return err
}

type NullableV1alpha2ResourceQuotaSpec struct {
	value *V1alpha2ResourceQuotaSpec
	isSet bool
}

func (v NullableV1alpha2ResourceQuotaSpec) Get() *V1alpha2ResourceQuotaSpec {
	return v.value
}

func (v *NullableV1alpha2ResourceQuotaSpec) Set(val *V1alpha2ResourceQuotaSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2ResourceQuotaSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2ResourceQuotaSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2ResourceQuotaSpec(val *V1alpha2ResourceQuotaSpec) *NullableV1alpha2ResourceQuotaSpec {
	return &NullableV1alpha2ResourceQuotaSpec{value: val, isSet: true}
}

func (v NullableV1alpha2ResourceQuotaSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2ResourceQuotaSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


