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

// checks if the V1AzureDiskVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AzureDiskVolumeSource{}

// V1AzureDiskVolumeSource AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type V1AzureDiskVolumeSource struct {
	// cachingMode is the Host Caching mode: None, Read Only, Read Write.
	CachingMode *string `json:"cachingMode,omitempty"`
	// diskName is the Name of the data disk in the blob storage
	DiskName string `json:"diskName"`
	// diskURI is the URI of data disk in the blob storage
	DiskURI string `json:"diskURI"`
	// fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	Kind *string `json:"kind,omitempty"`
	// readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

type _V1AzureDiskVolumeSource V1AzureDiskVolumeSource

// NewV1AzureDiskVolumeSource instantiates a new V1AzureDiskVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AzureDiskVolumeSource(diskName string, diskURI string) *V1AzureDiskVolumeSource {
	this := V1AzureDiskVolumeSource{}
	this.DiskName = diskName
	this.DiskURI = diskURI
	return &this
}

// NewV1AzureDiskVolumeSourceWithDefaults instantiates a new V1AzureDiskVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AzureDiskVolumeSourceWithDefaults() *V1AzureDiskVolumeSource {
	this := V1AzureDiskVolumeSource{}
	return &this
}

// GetCachingMode returns the CachingMode field value if set, zero value otherwise.
func (o *V1AzureDiskVolumeSource) GetCachingMode() string {
	if o == nil || IsNil(o.CachingMode) {
		var ret string
		return ret
	}
	return *o.CachingMode
}

// GetCachingModeOk returns a tuple with the CachingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AzureDiskVolumeSource) GetCachingModeOk() (*string, bool) {
	if o == nil || IsNil(o.CachingMode) {
		return nil, false
	}
	return o.CachingMode, true
}

// HasCachingMode returns a boolean if a field has been set.
func (o *V1AzureDiskVolumeSource) HasCachingMode() bool {
	if o != nil && !IsNil(o.CachingMode) {
		return true
	}

	return false
}

// SetCachingMode gets a reference to the given string and assigns it to the CachingMode field.
func (o *V1AzureDiskVolumeSource) SetCachingMode(v string) {
	o.CachingMode = &v
}

// GetDiskName returns the DiskName field value
func (o *V1AzureDiskVolumeSource) GetDiskName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskName
}

// GetDiskNameOk returns a tuple with the DiskName field value
// and a boolean to check if the value has been set.
func (o *V1AzureDiskVolumeSource) GetDiskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskName, true
}

// SetDiskName sets field value
func (o *V1AzureDiskVolumeSource) SetDiskName(v string) {
	o.DiskName = v
}

// GetDiskURI returns the DiskURI field value
func (o *V1AzureDiskVolumeSource) GetDiskURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskURI
}

// GetDiskURIOk returns a tuple with the DiskURI field value
// and a boolean to check if the value has been set.
func (o *V1AzureDiskVolumeSource) GetDiskURIOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskURI, true
}

// SetDiskURI sets field value
func (o *V1AzureDiskVolumeSource) SetDiskURI(v string) {
	o.DiskURI = v
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1AzureDiskVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AzureDiskVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1AzureDiskVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1AzureDiskVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1AzureDiskVolumeSource) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AzureDiskVolumeSource) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1AzureDiskVolumeSource) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1AzureDiskVolumeSource) SetKind(v string) {
	o.Kind = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1AzureDiskVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AzureDiskVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1AzureDiskVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1AzureDiskVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o V1AzureDiskVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AzureDiskVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CachingMode) {
		toSerialize["cachingMode"] = o.CachingMode
	}
	toSerialize["diskName"] = o.DiskName
	toSerialize["diskURI"] = o.DiskURI
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return toSerialize, nil
}

func (o *V1AzureDiskVolumeSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"diskName",
		"diskURI",
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

	varV1AzureDiskVolumeSource := _V1AzureDiskVolumeSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1AzureDiskVolumeSource)

	if err != nil {
		return err
	}

	*o = V1AzureDiskVolumeSource(varV1AzureDiskVolumeSource)

	return err
}

type NullableV1AzureDiskVolumeSource struct {
	value *V1AzureDiskVolumeSource
	isSet bool
}

func (v NullableV1AzureDiskVolumeSource) Get() *V1AzureDiskVolumeSource {
	return v.value
}

func (v *NullableV1AzureDiskVolumeSource) Set(val *V1AzureDiskVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AzureDiskVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AzureDiskVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AzureDiskVolumeSource(val *V1AzureDiskVolumeSource) *NullableV1AzureDiskVolumeSource {
	return &NullableV1AzureDiskVolumeSource{value: val, isSet: true}
}

func (v NullableV1AzureDiskVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AzureDiskVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


