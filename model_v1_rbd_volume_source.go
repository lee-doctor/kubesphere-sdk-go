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

// checks if the V1RBDVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RBDVolumeSource{}

// V1RBDVolumeSource Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
type V1RBDVolumeSource struct {
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	FsType *string `json:"fsType,omitempty"`
	// The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Image string `json:"image"`
	// Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring *string `json:"keyring,omitempty"`
	// A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Monitors []string `json:"monitors"`
	// The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool *string `json:"pool,omitempty"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	ReadOnly *bool `json:"readOnly,omitempty"`
	SecretRef *V1LocalObjectReference `json:"secretRef,omitempty"`
	// The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	User *string `json:"user,omitempty"`
}

type _V1RBDVolumeSource V1RBDVolumeSource

// NewV1RBDVolumeSource instantiates a new V1RBDVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RBDVolumeSource(image string, monitors []string) *V1RBDVolumeSource {
	this := V1RBDVolumeSource{}
	this.Image = image
	this.Monitors = monitors
	return &this
}

// NewV1RBDVolumeSourceWithDefaults instantiates a new V1RBDVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RBDVolumeSourceWithDefaults() *V1RBDVolumeSource {
	this := V1RBDVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1RBDVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1RBDVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1RBDVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetImage returns the Image field value
func (o *V1RBDVolumeSource) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *V1RBDVolumeSource) SetImage(v string) {
	o.Image = v
}

// GetKeyring returns the Keyring field value if set, zero value otherwise.
func (o *V1RBDVolumeSource) GetKeyring() string {
	if o == nil || IsNil(o.Keyring) {
		var ret string
		return ret
	}
	return *o.Keyring
}

// GetKeyringOk returns a tuple with the Keyring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetKeyringOk() (*string, bool) {
	if o == nil || IsNil(o.Keyring) {
		return nil, false
	}
	return o.Keyring, true
}

// HasKeyring returns a boolean if a field has been set.
func (o *V1RBDVolumeSource) HasKeyring() bool {
	if o != nil && !IsNil(o.Keyring) {
		return true
	}

	return false
}

// SetKeyring gets a reference to the given string and assigns it to the Keyring field.
func (o *V1RBDVolumeSource) SetKeyring(v string) {
	o.Keyring = &v
}

// GetMonitors returns the Monitors field value
func (o *V1RBDVolumeSource) GetMonitors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetMonitorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Monitors, true
}

// SetMonitors sets field value
func (o *V1RBDVolumeSource) SetMonitors(v []string) {
	o.Monitors = v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *V1RBDVolumeSource) GetPool() string {
	if o == nil || IsNil(o.Pool) {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *V1RBDVolumeSource) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *V1RBDVolumeSource) SetPool(v string) {
	o.Pool = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1RBDVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1RBDVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1RBDVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSecretRef returns the SecretRef field value if set, zero value otherwise.
func (o *V1RBDVolumeSource) GetSecretRef() V1LocalObjectReference {
	if o == nil || IsNil(o.SecretRef) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.SecretRef
}

// GetSecretRefOk returns a tuple with the SecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetSecretRefOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.SecretRef) {
		return nil, false
	}
	return o.SecretRef, true
}

// HasSecretRef returns a boolean if a field has been set.
func (o *V1RBDVolumeSource) HasSecretRef() bool {
	if o != nil && !IsNil(o.SecretRef) {
		return true
	}

	return false
}

// SetSecretRef gets a reference to the given V1LocalObjectReference and assigns it to the SecretRef field.
func (o *V1RBDVolumeSource) SetSecretRef(v V1LocalObjectReference) {
	o.SecretRef = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V1RBDVolumeSource) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RBDVolumeSource) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V1RBDVolumeSource) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *V1RBDVolumeSource) SetUser(v string) {
	o.User = &v
}

func (o V1RBDVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RBDVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	toSerialize["image"] = o.Image
	if !IsNil(o.Keyring) {
		toSerialize["keyring"] = o.Keyring
	}
	toSerialize["monitors"] = o.Monitors
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.SecretRef) {
		toSerialize["secretRef"] = o.SecretRef
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

func (o *V1RBDVolumeSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
		"monitors",
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

	varV1RBDVolumeSource := _V1RBDVolumeSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1RBDVolumeSource)

	if err != nil {
		return err
	}

	*o = V1RBDVolumeSource(varV1RBDVolumeSource)

	return err
}

type NullableV1RBDVolumeSource struct {
	value *V1RBDVolumeSource
	isSet bool
}

func (v NullableV1RBDVolumeSource) Get() *V1RBDVolumeSource {
	return v.value
}

func (v *NullableV1RBDVolumeSource) Set(val *V1RBDVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RBDVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RBDVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RBDVolumeSource(val *V1RBDVolumeSource) *NullableV1RBDVolumeSource {
	return &NullableV1RBDVolumeSource{value: val, isSet: true}
}

func (v NullableV1RBDVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RBDVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

