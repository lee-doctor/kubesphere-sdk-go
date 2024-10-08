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

// checks if the V1PortworxVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PortworxVolumeSource{}

// V1PortworxVolumeSource PortworxVolumeSource represents a Portworx volume resource.
type V1PortworxVolumeSource struct {
	// fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// volumeID uniquely identifies a Portworx volume
	VolumeID string `json:"volumeID"`
}

type _V1PortworxVolumeSource V1PortworxVolumeSource

// NewV1PortworxVolumeSource instantiates a new V1PortworxVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PortworxVolumeSource(volumeID string) *V1PortworxVolumeSource {
	this := V1PortworxVolumeSource{}
	this.VolumeID = volumeID
	return &this
}

// NewV1PortworxVolumeSourceWithDefaults instantiates a new V1PortworxVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PortworxVolumeSourceWithDefaults() *V1PortworxVolumeSource {
	this := V1PortworxVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1PortworxVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortworxVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1PortworxVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1PortworxVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1PortworxVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PortworxVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1PortworxVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1PortworxVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetVolumeID returns the VolumeID field value
func (o *V1PortworxVolumeSource) GetVolumeID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeID
}

// GetVolumeIDOk returns a tuple with the VolumeID field value
// and a boolean to check if the value has been set.
func (o *V1PortworxVolumeSource) GetVolumeIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeID, true
}

// SetVolumeID sets field value
func (o *V1PortworxVolumeSource) SetVolumeID(v string) {
	o.VolumeID = v
}

func (o V1PortworxVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PortworxVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	toSerialize["volumeID"] = o.VolumeID
	return toSerialize, nil
}

func (o *V1PortworxVolumeSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"volumeID",
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

	varV1PortworxVolumeSource := _V1PortworxVolumeSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1PortworxVolumeSource)

	if err != nil {
		return err
	}

	*o = V1PortworxVolumeSource(varV1PortworxVolumeSource)

	return err
}

type NullableV1PortworxVolumeSource struct {
	value *V1PortworxVolumeSource
	isSet bool
}

func (v NullableV1PortworxVolumeSource) Get() *V1PortworxVolumeSource {
	return v.value
}

func (v *NullableV1PortworxVolumeSource) Set(val *V1PortworxVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PortworxVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PortworxVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PortworxVolumeSource(val *V1PortworxVolumeSource) *NullableV1PortworxVolumeSource {
	return &NullableV1PortworxVolumeSource{value: val, isSet: true}
}

func (v NullableV1PortworxVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PortworxVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


