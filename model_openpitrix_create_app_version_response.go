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

// checks if the OpenpitrixCreateAppVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixCreateAppVersionResponse{}

// OpenpitrixCreateAppVersionResponse struct for OpenpitrixCreateAppVersionResponse
type OpenpitrixCreateAppVersionResponse struct {
	VersionId *string `json:"version_id,omitempty"`
}

// NewOpenpitrixCreateAppVersionResponse instantiates a new OpenpitrixCreateAppVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixCreateAppVersionResponse() *OpenpitrixCreateAppVersionResponse {
	this := OpenpitrixCreateAppVersionResponse{}
	return &this
}

// NewOpenpitrixCreateAppVersionResponseWithDefaults instantiates a new OpenpitrixCreateAppVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixCreateAppVersionResponseWithDefaults() *OpenpitrixCreateAppVersionResponse {
	this := OpenpitrixCreateAppVersionResponse{}
	return &this
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *OpenpitrixCreateAppVersionResponse) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateAppVersionResponse) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *OpenpitrixCreateAppVersionResponse) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *OpenpitrixCreateAppVersionResponse) SetVersionId(v string) {
	o.VersionId = &v
}

func (o OpenpitrixCreateAppVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixCreateAppVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionId) {
		toSerialize["version_id"] = o.VersionId
	}
	return toSerialize, nil
}

type NullableOpenpitrixCreateAppVersionResponse struct {
	value *OpenpitrixCreateAppVersionResponse
	isSet bool
}

func (v NullableOpenpitrixCreateAppVersionResponse) Get() *OpenpitrixCreateAppVersionResponse {
	return v.value
}

func (v *NullableOpenpitrixCreateAppVersionResponse) Set(val *OpenpitrixCreateAppVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixCreateAppVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixCreateAppVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixCreateAppVersionResponse(val *OpenpitrixCreateAppVersionResponse) *NullableOpenpitrixCreateAppVersionResponse {
	return &NullableOpenpitrixCreateAppVersionResponse{value: val, isSet: true}
}

func (v NullableOpenpitrixCreateAppVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixCreateAppVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


