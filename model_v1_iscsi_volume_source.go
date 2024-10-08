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

// checks if the V1ISCSIVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ISCSIVolumeSource{}

// V1ISCSIVolumeSource Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
type V1ISCSIVolumeSource struct {
	// chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
	ChapAuthDiscovery *bool `json:"chapAuthDiscovery,omitempty"`
	// chapAuthSession defines whether support iSCSI Session CHAP authentication
	ChapAuthSession *bool `json:"chapAuthSession,omitempty"`
	// fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	FsType *string `json:"fsType,omitempty"`
	// initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	InitiatorName *string `json:"initiatorName,omitempty"`
	// iqn is the target iSCSI Qualified Name.
	Iqn string `json:"iqn"`
	// iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	IscsiInterface *string `json:"iscsiInterface,omitempty"`
	// lun represents iSCSI Target Lun number.
	Lun int32 `json:"lun"`
	// portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Portals []string `json:"portals,omitempty"`
	// readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	ReadOnly *bool `json:"readOnly,omitempty"`
	SecretRef *V1LocalObjectReference `json:"secretRef,omitempty"`
	// targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal"`
}

type _V1ISCSIVolumeSource V1ISCSIVolumeSource

// NewV1ISCSIVolumeSource instantiates a new V1ISCSIVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ISCSIVolumeSource(iqn string, lun int32, targetPortal string) *V1ISCSIVolumeSource {
	this := V1ISCSIVolumeSource{}
	this.Iqn = iqn
	this.Lun = lun
	this.TargetPortal = targetPortal
	return &this
}

// NewV1ISCSIVolumeSourceWithDefaults instantiates a new V1ISCSIVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ISCSIVolumeSourceWithDefaults() *V1ISCSIVolumeSource {
	this := V1ISCSIVolumeSource{}
	return &this
}

// GetChapAuthDiscovery returns the ChapAuthDiscovery field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetChapAuthDiscovery() bool {
	if o == nil || IsNil(o.ChapAuthDiscovery) {
		var ret bool
		return ret
	}
	return *o.ChapAuthDiscovery
}

// GetChapAuthDiscoveryOk returns a tuple with the ChapAuthDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetChapAuthDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.ChapAuthDiscovery) {
		return nil, false
	}
	return o.ChapAuthDiscovery, true
}

// HasChapAuthDiscovery returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasChapAuthDiscovery() bool {
	if o != nil && !IsNil(o.ChapAuthDiscovery) {
		return true
	}

	return false
}

// SetChapAuthDiscovery gets a reference to the given bool and assigns it to the ChapAuthDiscovery field.
func (o *V1ISCSIVolumeSource) SetChapAuthDiscovery(v bool) {
	o.ChapAuthDiscovery = &v
}

// GetChapAuthSession returns the ChapAuthSession field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetChapAuthSession() bool {
	if o == nil || IsNil(o.ChapAuthSession) {
		var ret bool
		return ret
	}
	return *o.ChapAuthSession
}

// GetChapAuthSessionOk returns a tuple with the ChapAuthSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetChapAuthSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.ChapAuthSession) {
		return nil, false
	}
	return o.ChapAuthSession, true
}

// HasChapAuthSession returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasChapAuthSession() bool {
	if o != nil && !IsNil(o.ChapAuthSession) {
		return true
	}

	return false
}

// SetChapAuthSession gets a reference to the given bool and assigns it to the ChapAuthSession field.
func (o *V1ISCSIVolumeSource) SetChapAuthSession(v bool) {
	o.ChapAuthSession = &v
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1ISCSIVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetInitiatorName returns the InitiatorName field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetInitiatorName() string {
	if o == nil || IsNil(o.InitiatorName) {
		var ret string
		return ret
	}
	return *o.InitiatorName
}

// GetInitiatorNameOk returns a tuple with the InitiatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetInitiatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorName) {
		return nil, false
	}
	return o.InitiatorName, true
}

// HasInitiatorName returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasInitiatorName() bool {
	if o != nil && !IsNil(o.InitiatorName) {
		return true
	}

	return false
}

// SetInitiatorName gets a reference to the given string and assigns it to the InitiatorName field.
func (o *V1ISCSIVolumeSource) SetInitiatorName(v string) {
	o.InitiatorName = &v
}

// GetIqn returns the Iqn field value
func (o *V1ISCSIVolumeSource) GetIqn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iqn
}

// GetIqnOk returns a tuple with the Iqn field value
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetIqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iqn, true
}

// SetIqn sets field value
func (o *V1ISCSIVolumeSource) SetIqn(v string) {
	o.Iqn = v
}

// GetIscsiInterface returns the IscsiInterface field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetIscsiInterface() string {
	if o == nil || IsNil(o.IscsiInterface) {
		var ret string
		return ret
	}
	return *o.IscsiInterface
}

// GetIscsiInterfaceOk returns a tuple with the IscsiInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetIscsiInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.IscsiInterface) {
		return nil, false
	}
	return o.IscsiInterface, true
}

// HasIscsiInterface returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasIscsiInterface() bool {
	if o != nil && !IsNil(o.IscsiInterface) {
		return true
	}

	return false
}

// SetIscsiInterface gets a reference to the given string and assigns it to the IscsiInterface field.
func (o *V1ISCSIVolumeSource) SetIscsiInterface(v string) {
	o.IscsiInterface = &v
}

// GetLun returns the Lun field value
func (o *V1ISCSIVolumeSource) GetLun() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Lun
}

// GetLunOk returns a tuple with the Lun field value
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetLunOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lun, true
}

// SetLun sets field value
func (o *V1ISCSIVolumeSource) SetLun(v int32) {
	o.Lun = v
}

// GetPortals returns the Portals field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetPortals() []string {
	if o == nil || IsNil(o.Portals) {
		var ret []string
		return ret
	}
	return o.Portals
}

// GetPortalsOk returns a tuple with the Portals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetPortalsOk() ([]string, bool) {
	if o == nil || IsNil(o.Portals) {
		return nil, false
	}
	return o.Portals, true
}

// HasPortals returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasPortals() bool {
	if o != nil && !IsNil(o.Portals) {
		return true
	}

	return false
}

// SetPortals gets a reference to the given []string and assigns it to the Portals field.
func (o *V1ISCSIVolumeSource) SetPortals(v []string) {
	o.Portals = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1ISCSIVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSecretRef returns the SecretRef field value if set, zero value otherwise.
func (o *V1ISCSIVolumeSource) GetSecretRef() V1LocalObjectReference {
	if o == nil || IsNil(o.SecretRef) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.SecretRef
}

// GetSecretRefOk returns a tuple with the SecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetSecretRefOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.SecretRef) {
		return nil, false
	}
	return o.SecretRef, true
}

// HasSecretRef returns a boolean if a field has been set.
func (o *V1ISCSIVolumeSource) HasSecretRef() bool {
	if o != nil && !IsNil(o.SecretRef) {
		return true
	}

	return false
}

// SetSecretRef gets a reference to the given V1LocalObjectReference and assigns it to the SecretRef field.
func (o *V1ISCSIVolumeSource) SetSecretRef(v V1LocalObjectReference) {
	o.SecretRef = &v
}

// GetTargetPortal returns the TargetPortal field value
func (o *V1ISCSIVolumeSource) GetTargetPortal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPortal
}

// GetTargetPortalOk returns a tuple with the TargetPortal field value
// and a boolean to check if the value has been set.
func (o *V1ISCSIVolumeSource) GetTargetPortalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPortal, true
}

// SetTargetPortal sets field value
func (o *V1ISCSIVolumeSource) SetTargetPortal(v string) {
	o.TargetPortal = v
}

func (o V1ISCSIVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ISCSIVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChapAuthDiscovery) {
		toSerialize["chapAuthDiscovery"] = o.ChapAuthDiscovery
	}
	if !IsNil(o.ChapAuthSession) {
		toSerialize["chapAuthSession"] = o.ChapAuthSession
	}
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.InitiatorName) {
		toSerialize["initiatorName"] = o.InitiatorName
	}
	toSerialize["iqn"] = o.Iqn
	if !IsNil(o.IscsiInterface) {
		toSerialize["iscsiInterface"] = o.IscsiInterface
	}
	toSerialize["lun"] = o.Lun
	if !IsNil(o.Portals) {
		toSerialize["portals"] = o.Portals
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.SecretRef) {
		toSerialize["secretRef"] = o.SecretRef
	}
	toSerialize["targetPortal"] = o.TargetPortal
	return toSerialize, nil
}

func (o *V1ISCSIVolumeSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iqn",
		"lun",
		"targetPortal",
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

	varV1ISCSIVolumeSource := _V1ISCSIVolumeSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ISCSIVolumeSource)

	if err != nil {
		return err
	}

	*o = V1ISCSIVolumeSource(varV1ISCSIVolumeSource)

	return err
}

type NullableV1ISCSIVolumeSource struct {
	value *V1ISCSIVolumeSource
	isSet bool
}

func (v NullableV1ISCSIVolumeSource) Get() *V1ISCSIVolumeSource {
	return v.value
}

func (v *NullableV1ISCSIVolumeSource) Set(val *V1ISCSIVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ISCSIVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ISCSIVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ISCSIVolumeSource(val *V1ISCSIVolumeSource) *NullableV1ISCSIVolumeSource {
	return &NullableV1ISCSIVolumeSource{value: val, isSet: true}
}

func (v NullableV1ISCSIVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ISCSIVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


