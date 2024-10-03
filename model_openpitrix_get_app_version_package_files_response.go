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

// checks if the OpenpitrixGetAppVersionPackageFilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixGetAppVersionPackageFilesResponse{}

// OpenpitrixGetAppVersionPackageFilesResponse struct for OpenpitrixGetAppVersionPackageFilesResponse
type OpenpitrixGetAppVersionPackageFilesResponse struct {
	Files *map[string]string `json:"files,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
}

// NewOpenpitrixGetAppVersionPackageFilesResponse instantiates a new OpenpitrixGetAppVersionPackageFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixGetAppVersionPackageFilesResponse() *OpenpitrixGetAppVersionPackageFilesResponse {
	this := OpenpitrixGetAppVersionPackageFilesResponse{}
	return &this
}

// NewOpenpitrixGetAppVersionPackageFilesResponseWithDefaults instantiates a new OpenpitrixGetAppVersionPackageFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixGetAppVersionPackageFilesResponseWithDefaults() *OpenpitrixGetAppVersionPackageFilesResponse {
	this := OpenpitrixGetAppVersionPackageFilesResponse{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) GetFiles() map[string]string {
	if o == nil || IsNil(o.Files) {
		var ret map[string]string
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) GetFilesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given map[string]string and assigns it to the Files field.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) SetFiles(v map[string]string) {
	o.Files = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *OpenpitrixGetAppVersionPackageFilesResponse) SetVersionId(v string) {
	o.VersionId = &v
}

func (o OpenpitrixGetAppVersionPackageFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixGetAppVersionPackageFilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.VersionId) {
		toSerialize["version_id"] = o.VersionId
	}
	return toSerialize, nil
}

type NullableOpenpitrixGetAppVersionPackageFilesResponse struct {
	value *OpenpitrixGetAppVersionPackageFilesResponse
	isSet bool
}

func (v NullableOpenpitrixGetAppVersionPackageFilesResponse) Get() *OpenpitrixGetAppVersionPackageFilesResponse {
	return v.value
}

func (v *NullableOpenpitrixGetAppVersionPackageFilesResponse) Set(val *OpenpitrixGetAppVersionPackageFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixGetAppVersionPackageFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixGetAppVersionPackageFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixGetAppVersionPackageFilesResponse(val *OpenpitrixGetAppVersionPackageFilesResponse) *NullableOpenpitrixGetAppVersionPackageFilesResponse {
	return &NullableOpenpitrixGetAppVersionPackageFilesResponse{value: val, isSet: true}
}

func (v NullableOpenpitrixGetAppVersionPackageFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixGetAppVersionPackageFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

