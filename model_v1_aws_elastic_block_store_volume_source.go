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

// checks if the V1AWSElasticBlockStoreVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AWSElasticBlockStoreVolumeSource{}

// V1AWSElasticBlockStoreVolumeSource Represents a Persistent Disk resource in AWS.  An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.
type V1AWSElasticBlockStoreVolumeSource struct {
	// fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	FsType *string `json:"fsType,omitempty"`
	// partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty).
	Partition *int32 `json:"partition,omitempty"`
	// readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	ReadOnly *bool `json:"readOnly,omitempty"`
	// volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	VolumeID string `json:"volumeID"`
}

type _V1AWSElasticBlockStoreVolumeSource V1AWSElasticBlockStoreVolumeSource

// NewV1AWSElasticBlockStoreVolumeSource instantiates a new V1AWSElasticBlockStoreVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AWSElasticBlockStoreVolumeSource(volumeID string) *V1AWSElasticBlockStoreVolumeSource {
	this := V1AWSElasticBlockStoreVolumeSource{}
	this.VolumeID = volumeID
	return &this
}

// NewV1AWSElasticBlockStoreVolumeSourceWithDefaults instantiates a new V1AWSElasticBlockStoreVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AWSElasticBlockStoreVolumeSourceWithDefaults() *V1AWSElasticBlockStoreVolumeSource {
	this := V1AWSElasticBlockStoreVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1AWSElasticBlockStoreVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AWSElasticBlockStoreVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1AWSElasticBlockStoreVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1AWSElasticBlockStoreVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V1AWSElasticBlockStoreVolumeSource) GetPartition() int32 {
	if o == nil || IsNil(o.Partition) {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AWSElasticBlockStoreVolumeSource) GetPartitionOk() (*int32, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V1AWSElasticBlockStoreVolumeSource) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *V1AWSElasticBlockStoreVolumeSource) SetPartition(v int32) {
	o.Partition = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1AWSElasticBlockStoreVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AWSElasticBlockStoreVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1AWSElasticBlockStoreVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1AWSElasticBlockStoreVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetVolumeID returns the VolumeID field value
func (o *V1AWSElasticBlockStoreVolumeSource) GetVolumeID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeID
}

// GetVolumeIDOk returns a tuple with the VolumeID field value
// and a boolean to check if the value has been set.
func (o *V1AWSElasticBlockStoreVolumeSource) GetVolumeIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeID, true
}

// SetVolumeID sets field value
func (o *V1AWSElasticBlockStoreVolumeSource) SetVolumeID(v string) {
	o.VolumeID = v
}

func (o V1AWSElasticBlockStoreVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AWSElasticBlockStoreVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	toSerialize["volumeID"] = o.VolumeID
	return toSerialize, nil
}

func (o *V1AWSElasticBlockStoreVolumeSource) UnmarshalJSON(data []byte) (err error) {
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

	varV1AWSElasticBlockStoreVolumeSource := _V1AWSElasticBlockStoreVolumeSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1AWSElasticBlockStoreVolumeSource)

	if err != nil {
		return err
	}

	*o = V1AWSElasticBlockStoreVolumeSource(varV1AWSElasticBlockStoreVolumeSource)

	return err
}

type NullableV1AWSElasticBlockStoreVolumeSource struct {
	value *V1AWSElasticBlockStoreVolumeSource
	isSet bool
}

func (v NullableV1AWSElasticBlockStoreVolumeSource) Get() *V1AWSElasticBlockStoreVolumeSource {
	return v.value
}

func (v *NullableV1AWSElasticBlockStoreVolumeSource) Set(val *V1AWSElasticBlockStoreVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AWSElasticBlockStoreVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AWSElasticBlockStoreVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AWSElasticBlockStoreVolumeSource(val *V1AWSElasticBlockStoreVolumeSource) *NullableV1AWSElasticBlockStoreVolumeSource {
	return &NullableV1AWSElasticBlockStoreVolumeSource{value: val, isSet: true}
}

func (v NullableV1AWSElasticBlockStoreVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AWSElasticBlockStoreVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


