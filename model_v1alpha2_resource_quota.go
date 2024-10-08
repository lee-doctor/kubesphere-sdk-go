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

// checks if the V1alpha2ResourceQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2ResourceQuota{}

// V1alpha2ResourceQuota struct for V1alpha2ResourceQuota
type V1alpha2ResourceQuota struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec V1alpha2ResourceQuotaSpec `json:"spec"`
	Status *V1alpha2ResourceQuotaStatus `json:"status,omitempty"`
}

type _V1alpha2ResourceQuota V1alpha2ResourceQuota

// NewV1alpha2ResourceQuota instantiates a new V1alpha2ResourceQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2ResourceQuota(spec V1alpha2ResourceQuotaSpec) *V1alpha2ResourceQuota {
	this := V1alpha2ResourceQuota{}
	this.Spec = spec
	return &this
}

// NewV1alpha2ResourceQuotaWithDefaults instantiates a new V1alpha2ResourceQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2ResourceQuotaWithDefaults() *V1alpha2ResourceQuota {
	this := V1alpha2ResourceQuota{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1alpha2ResourceQuota) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2ResourceQuota) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1alpha2ResourceQuota) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1alpha2ResourceQuota) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha2ResourceQuota) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2ResourceQuota) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha2ResourceQuota) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha2ResourceQuota) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha2ResourceQuota) GetMetadata() V1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2ResourceQuota) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha2ResourceQuota) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1alpha2ResourceQuota) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value
func (o *V1alpha2ResourceQuota) GetSpec() V1alpha2ResourceQuotaSpec {
	if o == nil {
		var ret V1alpha2ResourceQuotaSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ResourceQuota) GetSpecOk() (*V1alpha2ResourceQuotaSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *V1alpha2ResourceQuota) SetSpec(v V1alpha2ResourceQuotaSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1alpha2ResourceQuota) GetStatus() V1alpha2ResourceQuotaStatus {
	if o == nil || IsNil(o.Status) {
		var ret V1alpha2ResourceQuotaStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2ResourceQuota) GetStatusOk() (*V1alpha2ResourceQuotaStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1alpha2ResourceQuota) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1alpha2ResourceQuotaStatus and assigns it to the Status field.
func (o *V1alpha2ResourceQuota) SetStatus(v V1alpha2ResourceQuotaStatus) {
	o.Status = &v
}

func (o V1alpha2ResourceQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2ResourceQuota) ToMap() (map[string]interface{}, error) {
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
	toSerialize["spec"] = o.Spec
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *V1alpha2ResourceQuota) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"spec",
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

	varV1alpha2ResourceQuota := _V1alpha2ResourceQuota{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2ResourceQuota)

	if err != nil {
		return err
	}

	*o = V1alpha2ResourceQuota(varV1alpha2ResourceQuota)

	return err
}

type NullableV1alpha2ResourceQuota struct {
	value *V1alpha2ResourceQuota
	isSet bool
}

func (v NullableV1alpha2ResourceQuota) Get() *V1alpha2ResourceQuota {
	return v.value
}

func (v *NullableV1alpha2ResourceQuota) Set(val *V1alpha2ResourceQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2ResourceQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2ResourceQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2ResourceQuota(val *V1alpha2ResourceQuota) *NullableV1alpha2ResourceQuota {
	return &NullableV1alpha2ResourceQuota{value: val, isSet: true}
}

func (v NullableV1alpha2ResourceQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2ResourceQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


