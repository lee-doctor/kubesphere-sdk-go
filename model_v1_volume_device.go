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

// checks if the V1VolumeDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1VolumeDevice{}

// V1VolumeDevice volumeDevice describes a mapping of a raw block device within a container.
type V1VolumeDevice struct {
	// devicePath is the path inside of the container that the device will be mapped to.
	DevicePath string `json:"devicePath"`
	// name must match the name of a persistentVolumeClaim in the pod
	Name string `json:"name"`
}

type _V1VolumeDevice V1VolumeDevice

// NewV1VolumeDevice instantiates a new V1VolumeDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1VolumeDevice(devicePath string, name string) *V1VolumeDevice {
	this := V1VolumeDevice{}
	this.DevicePath = devicePath
	this.Name = name
	return &this
}

// NewV1VolumeDeviceWithDefaults instantiates a new V1VolumeDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1VolumeDeviceWithDefaults() *V1VolumeDevice {
	this := V1VolumeDevice{}
	return &this
}

// GetDevicePath returns the DevicePath field value
func (o *V1VolumeDevice) GetDevicePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePath
}

// GetDevicePathOk returns a tuple with the DevicePath field value
// and a boolean to check if the value has been set.
func (o *V1VolumeDevice) GetDevicePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevicePath, true
}

// SetDevicePath sets field value
func (o *V1VolumeDevice) SetDevicePath(v string) {
	o.DevicePath = v
}

// GetName returns the Name field value
func (o *V1VolumeDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1VolumeDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1VolumeDevice) SetName(v string) {
	o.Name = v
}

func (o V1VolumeDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1VolumeDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devicePath"] = o.DevicePath
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *V1VolumeDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"devicePath",
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

	varV1VolumeDevice := _V1VolumeDevice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1VolumeDevice)

	if err != nil {
		return err
	}

	*o = V1VolumeDevice(varV1VolumeDevice)

	return err
}

type NullableV1VolumeDevice struct {
	value *V1VolumeDevice
	isSet bool
}

func (v NullableV1VolumeDevice) Get() *V1VolumeDevice {
	return v.value
}

func (v *NullableV1VolumeDevice) Set(val *V1VolumeDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VolumeDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VolumeDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VolumeDevice(val *V1VolumeDevice) *NullableV1VolumeDevice {
	return &NullableV1VolumeDevice{value: val, isSet: true}
}

func (v NullableV1VolumeDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VolumeDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


