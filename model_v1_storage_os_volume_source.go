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

// checks if the V1StorageOSVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1StorageOSVolumeSource{}

// V1StorageOSVolumeSource Represents a StorageOS persistent volume resource.
type V1StorageOSVolumeSource struct {
	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	SecretRef *V1LocalObjectReference `json:"secretRef,omitempty"`
	// volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
	VolumeName *string `json:"volumeName,omitempty"`
	// volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to \"default\" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
	VolumeNamespace *string `json:"volumeNamespace,omitempty"`
}

// NewV1StorageOSVolumeSource instantiates a new V1StorageOSVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1StorageOSVolumeSource() *V1StorageOSVolumeSource {
	this := V1StorageOSVolumeSource{}
	return &this
}

// NewV1StorageOSVolumeSourceWithDefaults instantiates a new V1StorageOSVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1StorageOSVolumeSourceWithDefaults() *V1StorageOSVolumeSource {
	this := V1StorageOSVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1StorageOSVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StorageOSVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1StorageOSVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1StorageOSVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1StorageOSVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StorageOSVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1StorageOSVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1StorageOSVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSecretRef returns the SecretRef field value if set, zero value otherwise.
func (o *V1StorageOSVolumeSource) GetSecretRef() V1LocalObjectReference {
	if o == nil || IsNil(o.SecretRef) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.SecretRef
}

// GetSecretRefOk returns a tuple with the SecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StorageOSVolumeSource) GetSecretRefOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.SecretRef) {
		return nil, false
	}
	return o.SecretRef, true
}

// HasSecretRef returns a boolean if a field has been set.
func (o *V1StorageOSVolumeSource) HasSecretRef() bool {
	if o != nil && !IsNil(o.SecretRef) {
		return true
	}

	return false
}

// SetSecretRef gets a reference to the given V1LocalObjectReference and assigns it to the SecretRef field.
func (o *V1StorageOSVolumeSource) SetSecretRef(v V1LocalObjectReference) {
	o.SecretRef = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *V1StorageOSVolumeSource) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StorageOSVolumeSource) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *V1StorageOSVolumeSource) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *V1StorageOSVolumeSource) SetVolumeName(v string) {
	o.VolumeName = &v
}

// GetVolumeNamespace returns the VolumeNamespace field value if set, zero value otherwise.
func (o *V1StorageOSVolumeSource) GetVolumeNamespace() string {
	if o == nil || IsNil(o.VolumeNamespace) {
		var ret string
		return ret
	}
	return *o.VolumeNamespace
}

// GetVolumeNamespaceOk returns a tuple with the VolumeNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StorageOSVolumeSource) GetVolumeNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeNamespace) {
		return nil, false
	}
	return o.VolumeNamespace, true
}

// HasVolumeNamespace returns a boolean if a field has been set.
func (o *V1StorageOSVolumeSource) HasVolumeNamespace() bool {
	if o != nil && !IsNil(o.VolumeNamespace) {
		return true
	}

	return false
}

// SetVolumeNamespace gets a reference to the given string and assigns it to the VolumeNamespace field.
func (o *V1StorageOSVolumeSource) SetVolumeNamespace(v string) {
	o.VolumeNamespace = &v
}

func (o V1StorageOSVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1StorageOSVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.SecretRef) {
		toSerialize["secretRef"] = o.SecretRef
	}
	if !IsNil(o.VolumeName) {
		toSerialize["volumeName"] = o.VolumeName
	}
	if !IsNil(o.VolumeNamespace) {
		toSerialize["volumeNamespace"] = o.VolumeNamespace
	}
	return toSerialize, nil
}

type NullableV1StorageOSVolumeSource struct {
	value *V1StorageOSVolumeSource
	isSet bool
}

func (v NullableV1StorageOSVolumeSource) Get() *V1StorageOSVolumeSource {
	return v.value
}

func (v *NullableV1StorageOSVolumeSource) Set(val *V1StorageOSVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1StorageOSVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1StorageOSVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1StorageOSVolumeSource(val *V1StorageOSVolumeSource) *NullableV1StorageOSVolumeSource {
	return &NullableV1StorageOSVolumeSource{value: val, isSet: true}
}

func (v NullableV1StorageOSVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1StorageOSVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


