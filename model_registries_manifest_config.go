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
)

// checks if the RegistriesManifestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistriesManifestConfig{}

// RegistriesManifestConfig struct for RegistriesManifestConfig
type RegistriesManifestConfig struct {
	// The digest of the content, as defined by the Registry V2 HTTP API Specificiation. Reference https://docs.docker.com/registry/spec/api/#digest-parameter
	Digest *string `json:"digest,omitempty"`
	// The MIME type of the image.
	MediaType *string `json:"mediaType,omitempty"`
	// The size in bytes of the image.
	Size *int32 `json:"size,omitempty"`
}

// NewRegistriesManifestConfig instantiates a new RegistriesManifestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistriesManifestConfig() *RegistriesManifestConfig {
	this := RegistriesManifestConfig{}
	return &this
}

// NewRegistriesManifestConfigWithDefaults instantiates a new RegistriesManifestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistriesManifestConfigWithDefaults() *RegistriesManifestConfig {
	this := RegistriesManifestConfig{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *RegistriesManifestConfig) GetDigest() string {
	if o == nil || IsNil(o.Digest) {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesManifestConfig) GetDigestOk() (*string, bool) {
	if o == nil || IsNil(o.Digest) {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *RegistriesManifestConfig) HasDigest() bool {
	if o != nil && !IsNil(o.Digest) {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *RegistriesManifestConfig) SetDigest(v string) {
	o.Digest = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *RegistriesManifestConfig) GetMediaType() string {
	if o == nil || IsNil(o.MediaType) {
		var ret string
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesManifestConfig) GetMediaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaType) {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *RegistriesManifestConfig) HasMediaType() bool {
	if o != nil && !IsNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given string and assigns it to the MediaType field.
func (o *RegistriesManifestConfig) SetMediaType(v string) {
	o.MediaType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RegistriesManifestConfig) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesManifestConfig) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RegistriesManifestConfig) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *RegistriesManifestConfig) SetSize(v int32) {
	o.Size = &v
}

func (o RegistriesManifestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistriesManifestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Digest) {
		toSerialize["digest"] = o.Digest
	}
	if !IsNil(o.MediaType) {
		toSerialize["mediaType"] = o.MediaType
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableRegistriesManifestConfig struct {
	value *RegistriesManifestConfig
	isSet bool
}

func (v NullableRegistriesManifestConfig) Get() *RegistriesManifestConfig {
	return v.value
}

func (v *NullableRegistriesManifestConfig) Set(val *RegistriesManifestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistriesManifestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistriesManifestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistriesManifestConfig(val *RegistriesManifestConfig) *NullableRegistriesManifestConfig {
	return &NullableRegistriesManifestConfig{value: val, isSet: true}
}

func (v NullableRegistriesManifestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistriesManifestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


