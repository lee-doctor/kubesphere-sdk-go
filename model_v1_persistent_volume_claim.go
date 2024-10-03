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

// checks if the V1PersistentVolumeClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PersistentVolumeClaim{}

// V1PersistentVolumeClaim PersistentVolumeClaim is a user's request for and claim to a persistent volume
type V1PersistentVolumeClaim struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec *V1PersistentVolumeClaimSpec `json:"spec,omitempty"`
	Status *V1PersistentVolumeClaimStatus `json:"status,omitempty"`
}

// NewV1PersistentVolumeClaim instantiates a new V1PersistentVolumeClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PersistentVolumeClaim() *V1PersistentVolumeClaim {
	this := V1PersistentVolumeClaim{}
	return &this
}

// NewV1PersistentVolumeClaimWithDefaults instantiates a new V1PersistentVolumeClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PersistentVolumeClaimWithDefaults() *V1PersistentVolumeClaim {
	this := V1PersistentVolumeClaim{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1PersistentVolumeClaim) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PersistentVolumeClaim) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1PersistentVolumeClaim) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1PersistentVolumeClaim) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1PersistentVolumeClaim) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PersistentVolumeClaim) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1PersistentVolumeClaim) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1PersistentVolumeClaim) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1PersistentVolumeClaim) GetMetadata() V1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PersistentVolumeClaim) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1PersistentVolumeClaim) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1PersistentVolumeClaim) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *V1PersistentVolumeClaim) GetSpec() V1PersistentVolumeClaimSpec {
	if o == nil || IsNil(o.Spec) {
		var ret V1PersistentVolumeClaimSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PersistentVolumeClaim) GetSpecOk() (*V1PersistentVolumeClaimSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *V1PersistentVolumeClaim) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given V1PersistentVolumeClaimSpec and assigns it to the Spec field.
func (o *V1PersistentVolumeClaim) SetSpec(v V1PersistentVolumeClaimSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1PersistentVolumeClaim) GetStatus() V1PersistentVolumeClaimStatus {
	if o == nil || IsNil(o.Status) {
		var ret V1PersistentVolumeClaimStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PersistentVolumeClaim) GetStatusOk() (*V1PersistentVolumeClaimStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1PersistentVolumeClaim) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1PersistentVolumeClaimStatus and assigns it to the Status field.
func (o *V1PersistentVolumeClaim) SetStatus(v V1PersistentVolumeClaimStatus) {
	o.Status = &v
}

func (o V1PersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PersistentVolumeClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV1PersistentVolumeClaim struct {
	value *V1PersistentVolumeClaim
	isSet bool
}

func (v NullableV1PersistentVolumeClaim) Get() *V1PersistentVolumeClaim {
	return v.value
}

func (v *NullableV1PersistentVolumeClaim) Set(val *V1PersistentVolumeClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PersistentVolumeClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PersistentVolumeClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PersistentVolumeClaim(val *V1PersistentVolumeClaim) *NullableV1PersistentVolumeClaim {
	return &NullableV1PersistentVolumeClaim{value: val, isSet: true}
}

func (v NullableV1PersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PersistentVolumeClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

