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

// checks if the V1alpha2HealthStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2HealthStatus{}

// V1alpha2HealthStatus struct for V1alpha2HealthStatus
type V1alpha2HealthStatus struct {
	// kubesphere components status
	KubesphereStatus []V1alpha2ComponentStatus `json:"kubesphereStatus"`
	NodeStatus V1alpha2NodeStatus `json:"nodeStatus"`
}

type _V1alpha2HealthStatus V1alpha2HealthStatus

// NewV1alpha2HealthStatus instantiates a new V1alpha2HealthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2HealthStatus(kubesphereStatus []V1alpha2ComponentStatus, nodeStatus V1alpha2NodeStatus) *V1alpha2HealthStatus {
	this := V1alpha2HealthStatus{}
	this.KubesphereStatus = kubesphereStatus
	this.NodeStatus = nodeStatus
	return &this
}

// NewV1alpha2HealthStatusWithDefaults instantiates a new V1alpha2HealthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2HealthStatusWithDefaults() *V1alpha2HealthStatus {
	this := V1alpha2HealthStatus{}
	return &this
}

// GetKubesphereStatus returns the KubesphereStatus field value
func (o *V1alpha2HealthStatus) GetKubesphereStatus() []V1alpha2ComponentStatus {
	if o == nil {
		var ret []V1alpha2ComponentStatus
		return ret
	}

	return o.KubesphereStatus
}

// GetKubesphereStatusOk returns a tuple with the KubesphereStatus field value
// and a boolean to check if the value has been set.
func (o *V1alpha2HealthStatus) GetKubesphereStatusOk() ([]V1alpha2ComponentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubesphereStatus, true
}

// SetKubesphereStatus sets field value
func (o *V1alpha2HealthStatus) SetKubesphereStatus(v []V1alpha2ComponentStatus) {
	o.KubesphereStatus = v
}

// GetNodeStatus returns the NodeStatus field value
func (o *V1alpha2HealthStatus) GetNodeStatus() V1alpha2NodeStatus {
	if o == nil {
		var ret V1alpha2NodeStatus
		return ret
	}

	return o.NodeStatus
}

// GetNodeStatusOk returns a tuple with the NodeStatus field value
// and a boolean to check if the value has been set.
func (o *V1alpha2HealthStatus) GetNodeStatusOk() (*V1alpha2NodeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeStatus, true
}

// SetNodeStatus sets field value
func (o *V1alpha2HealthStatus) SetNodeStatus(v V1alpha2NodeStatus) {
	o.NodeStatus = v
}

func (o V1alpha2HealthStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2HealthStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kubesphereStatus"] = o.KubesphereStatus
	toSerialize["nodeStatus"] = o.NodeStatus
	return toSerialize, nil
}

func (o *V1alpha2HealthStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kubesphereStatus",
		"nodeStatus",
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

	varV1alpha2HealthStatus := _V1alpha2HealthStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2HealthStatus)

	if err != nil {
		return err
	}

	*o = V1alpha2HealthStatus(varV1alpha2HealthStatus)

	return err
}

type NullableV1alpha2HealthStatus struct {
	value *V1alpha2HealthStatus
	isSet bool
}

func (v NullableV1alpha2HealthStatus) Get() *V1alpha2HealthStatus {
	return v.value
}

func (v *NullableV1alpha2HealthStatus) Set(val *V1alpha2HealthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2HealthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2HealthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2HealthStatus(val *V1alpha2HealthStatus) *NullableV1alpha2HealthStatus {
	return &NullableV1alpha2HealthStatus{value: val, isSet: true}
}

func (v NullableV1alpha2HealthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2HealthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


