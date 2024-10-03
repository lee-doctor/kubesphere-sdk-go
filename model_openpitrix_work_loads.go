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

// checks if the OpenpitrixWorkLoads type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixWorkLoads{}

// OpenpitrixWorkLoads struct for OpenpitrixWorkLoads
type OpenpitrixWorkLoads struct {
	// daemonset list
	Daemonsets []V1DaemonSet `json:"daemonsets,omitempty"`
	// deployment list
	Deployments []V1Deployment `json:"deployments,omitempty"`
	// statefulset list
	Statefulsets []V1StatefulSet `json:"statefulsets,omitempty"`
}

// NewOpenpitrixWorkLoads instantiates a new OpenpitrixWorkLoads object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixWorkLoads() *OpenpitrixWorkLoads {
	this := OpenpitrixWorkLoads{}
	return &this
}

// NewOpenpitrixWorkLoadsWithDefaults instantiates a new OpenpitrixWorkLoads object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixWorkLoadsWithDefaults() *OpenpitrixWorkLoads {
	this := OpenpitrixWorkLoads{}
	return &this
}

// GetDaemonsets returns the Daemonsets field value if set, zero value otherwise.
func (o *OpenpitrixWorkLoads) GetDaemonsets() []V1DaemonSet {
	if o == nil || IsNil(o.Daemonsets) {
		var ret []V1DaemonSet
		return ret
	}
	return o.Daemonsets
}

// GetDaemonsetsOk returns a tuple with the Daemonsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixWorkLoads) GetDaemonsetsOk() ([]V1DaemonSet, bool) {
	if o == nil || IsNil(o.Daemonsets) {
		return nil, false
	}
	return o.Daemonsets, true
}

// HasDaemonsets returns a boolean if a field has been set.
func (o *OpenpitrixWorkLoads) HasDaemonsets() bool {
	if o != nil && !IsNil(o.Daemonsets) {
		return true
	}

	return false
}

// SetDaemonsets gets a reference to the given []V1DaemonSet and assigns it to the Daemonsets field.
func (o *OpenpitrixWorkLoads) SetDaemonsets(v []V1DaemonSet) {
	o.Daemonsets = v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *OpenpitrixWorkLoads) GetDeployments() []V1Deployment {
	if o == nil || IsNil(o.Deployments) {
		var ret []V1Deployment
		return ret
	}
	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixWorkLoads) GetDeploymentsOk() ([]V1Deployment, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *OpenpitrixWorkLoads) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []V1Deployment and assigns it to the Deployments field.
func (o *OpenpitrixWorkLoads) SetDeployments(v []V1Deployment) {
	o.Deployments = v
}

// GetStatefulsets returns the Statefulsets field value if set, zero value otherwise.
func (o *OpenpitrixWorkLoads) GetStatefulsets() []V1StatefulSet {
	if o == nil || IsNil(o.Statefulsets) {
		var ret []V1StatefulSet
		return ret
	}
	return o.Statefulsets
}

// GetStatefulsetsOk returns a tuple with the Statefulsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixWorkLoads) GetStatefulsetsOk() ([]V1StatefulSet, bool) {
	if o == nil || IsNil(o.Statefulsets) {
		return nil, false
	}
	return o.Statefulsets, true
}

// HasStatefulsets returns a boolean if a field has been set.
func (o *OpenpitrixWorkLoads) HasStatefulsets() bool {
	if o != nil && !IsNil(o.Statefulsets) {
		return true
	}

	return false
}

// SetStatefulsets gets a reference to the given []V1StatefulSet and assigns it to the Statefulsets field.
func (o *OpenpitrixWorkLoads) SetStatefulsets(v []V1StatefulSet) {
	o.Statefulsets = v
}

func (o OpenpitrixWorkLoads) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixWorkLoads) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Daemonsets) {
		toSerialize["daemonsets"] = o.Daemonsets
	}
	if !IsNil(o.Deployments) {
		toSerialize["deployments"] = o.Deployments
	}
	if !IsNil(o.Statefulsets) {
		toSerialize["statefulsets"] = o.Statefulsets
	}
	return toSerialize, nil
}

type NullableOpenpitrixWorkLoads struct {
	value *OpenpitrixWorkLoads
	isSet bool
}

func (v NullableOpenpitrixWorkLoads) Get() *OpenpitrixWorkLoads {
	return v.value
}

func (v *NullableOpenpitrixWorkLoads) Set(val *OpenpitrixWorkLoads) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixWorkLoads) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixWorkLoads) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixWorkLoads(val *OpenpitrixWorkLoads) *NullableOpenpitrixWorkLoads {
	return &NullableOpenpitrixWorkLoads{value: val, isSet: true}
}

func (v NullableOpenpitrixWorkLoads) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixWorkLoads) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


