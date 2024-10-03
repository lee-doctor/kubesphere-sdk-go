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

// checks if the V1alpha2Member type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2Member{}

// V1alpha2Member struct for V1alpha2Member
type V1alpha2Member struct {
	RoleRef string `json:"roleRef"`
	Username string `json:"username"`
}

type _V1alpha2Member V1alpha2Member

// NewV1alpha2Member instantiates a new V1alpha2Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2Member(roleRef string, username string) *V1alpha2Member {
	this := V1alpha2Member{}
	this.RoleRef = roleRef
	this.Username = username
	return &this
}

// NewV1alpha2MemberWithDefaults instantiates a new V1alpha2Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2MemberWithDefaults() *V1alpha2Member {
	this := V1alpha2Member{}
	return &this
}

// GetRoleRef returns the RoleRef field value
func (o *V1alpha2Member) GetRoleRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleRef
}

// GetRoleRefOk returns a tuple with the RoleRef field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Member) GetRoleRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleRef, true
}

// SetRoleRef sets field value
func (o *V1alpha2Member) SetRoleRef(v string) {
	o.RoleRef = v
}

// GetUsername returns the Username field value
func (o *V1alpha2Member) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Member) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *V1alpha2Member) SetUsername(v string) {
	o.Username = v
}

func (o V1alpha2Member) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2Member) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleRef"] = o.RoleRef
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *V1alpha2Member) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roleRef",
		"username",
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

	varV1alpha2Member := _V1alpha2Member{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2Member)

	if err != nil {
		return err
	}

	*o = V1alpha2Member(varV1alpha2Member)

	return err
}

type NullableV1alpha2Member struct {
	value *V1alpha2Member
	isSet bool
}

func (v NullableV1alpha2Member) Get() *V1alpha2Member {
	return v.value
}

func (v *NullableV1alpha2Member) Set(val *V1alpha2Member) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2Member) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2Member) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2Member(val *V1alpha2Member) *NullableV1alpha2Member {
	return &NullableV1alpha2Member{value: val, isSet: true}
}

func (v NullableV1alpha2Member) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2Member) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


