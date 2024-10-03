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

// checks if the V1beta1GenericPlacementFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1GenericPlacementFields{}

// V1beta1GenericPlacementFields struct for V1beta1GenericPlacementFields
type V1beta1GenericPlacementFields struct {
	ClusterSelector *V1LabelSelector `json:"clusterSelector,omitempty"`
	Clusters []V1beta1GenericClusterReference `json:"clusters,omitempty"`
}

// NewV1beta1GenericPlacementFields instantiates a new V1beta1GenericPlacementFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1GenericPlacementFields() *V1beta1GenericPlacementFields {
	this := V1beta1GenericPlacementFields{}
	return &this
}

// NewV1beta1GenericPlacementFieldsWithDefaults instantiates a new V1beta1GenericPlacementFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1GenericPlacementFieldsWithDefaults() *V1beta1GenericPlacementFields {
	this := V1beta1GenericPlacementFields{}
	return &this
}

// GetClusterSelector returns the ClusterSelector field value if set, zero value otherwise.
func (o *V1beta1GenericPlacementFields) GetClusterSelector() V1LabelSelector {
	if o == nil || IsNil(o.ClusterSelector) {
		var ret V1LabelSelector
		return ret
	}
	return *o.ClusterSelector
}

// GetClusterSelectorOk returns a tuple with the ClusterSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1GenericPlacementFields) GetClusterSelectorOk() (*V1LabelSelector, bool) {
	if o == nil || IsNil(o.ClusterSelector) {
		return nil, false
	}
	return o.ClusterSelector, true
}

// HasClusterSelector returns a boolean if a field has been set.
func (o *V1beta1GenericPlacementFields) HasClusterSelector() bool {
	if o != nil && !IsNil(o.ClusterSelector) {
		return true
	}

	return false
}

// SetClusterSelector gets a reference to the given V1LabelSelector and assigns it to the ClusterSelector field.
func (o *V1beta1GenericPlacementFields) SetClusterSelector(v V1LabelSelector) {
	o.ClusterSelector = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V1beta1GenericPlacementFields) GetClusters() []V1beta1GenericClusterReference {
	if o == nil || IsNil(o.Clusters) {
		var ret []V1beta1GenericClusterReference
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1GenericPlacementFields) GetClustersOk() ([]V1beta1GenericClusterReference, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V1beta1GenericPlacementFields) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []V1beta1GenericClusterReference and assigns it to the Clusters field.
func (o *V1beta1GenericPlacementFields) SetClusters(v []V1beta1GenericClusterReference) {
	o.Clusters = v
}

func (o V1beta1GenericPlacementFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1GenericPlacementFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterSelector) {
		toSerialize["clusterSelector"] = o.ClusterSelector
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

type NullableV1beta1GenericPlacementFields struct {
	value *V1beta1GenericPlacementFields
	isSet bool
}

func (v NullableV1beta1GenericPlacementFields) Get() *V1beta1GenericPlacementFields {
	return v.value
}

func (v *NullableV1beta1GenericPlacementFields) Set(val *V1beta1GenericPlacementFields) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1GenericPlacementFields) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1GenericPlacementFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1GenericPlacementFields(val *V1beta1GenericPlacementFields) *NullableV1beta1GenericPlacementFields {
	return &NullableV1beta1GenericPlacementFields{value: val, isSet: true}
}

func (v NullableV1beta1GenericPlacementFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1GenericPlacementFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

