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

// checks if the V1alpha2GenericPlacementFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2GenericPlacementFields{}

// V1alpha2GenericPlacementFields struct for V1alpha2GenericPlacementFields
type V1alpha2GenericPlacementFields struct {
	ClusterSelector *V1LabelSelector `json:"clusterSelector,omitempty"`
	Clusters []V1alpha2GenericClusterReference `json:"clusters,omitempty"`
}

// NewV1alpha2GenericPlacementFields instantiates a new V1alpha2GenericPlacementFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2GenericPlacementFields() *V1alpha2GenericPlacementFields {
	this := V1alpha2GenericPlacementFields{}
	return &this
}

// NewV1alpha2GenericPlacementFieldsWithDefaults instantiates a new V1alpha2GenericPlacementFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2GenericPlacementFieldsWithDefaults() *V1alpha2GenericPlacementFields {
	this := V1alpha2GenericPlacementFields{}
	return &this
}

// GetClusterSelector returns the ClusterSelector field value if set, zero value otherwise.
func (o *V1alpha2GenericPlacementFields) GetClusterSelector() V1LabelSelector {
	if o == nil || IsNil(o.ClusterSelector) {
		var ret V1LabelSelector
		return ret
	}
	return *o.ClusterSelector
}

// GetClusterSelectorOk returns a tuple with the ClusterSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2GenericPlacementFields) GetClusterSelectorOk() (*V1LabelSelector, bool) {
	if o == nil || IsNil(o.ClusterSelector) {
		return nil, false
	}
	return o.ClusterSelector, true
}

// HasClusterSelector returns a boolean if a field has been set.
func (o *V1alpha2GenericPlacementFields) HasClusterSelector() bool {
	if o != nil && !IsNil(o.ClusterSelector) {
		return true
	}

	return false
}

// SetClusterSelector gets a reference to the given V1LabelSelector and assigns it to the ClusterSelector field.
func (o *V1alpha2GenericPlacementFields) SetClusterSelector(v V1LabelSelector) {
	o.ClusterSelector = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V1alpha2GenericPlacementFields) GetClusters() []V1alpha2GenericClusterReference {
	if o == nil || IsNil(o.Clusters) {
		var ret []V1alpha2GenericClusterReference
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2GenericPlacementFields) GetClustersOk() ([]V1alpha2GenericClusterReference, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V1alpha2GenericPlacementFields) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []V1alpha2GenericClusterReference and assigns it to the Clusters field.
func (o *V1alpha2GenericPlacementFields) SetClusters(v []V1alpha2GenericClusterReference) {
	o.Clusters = v
}

func (o V1alpha2GenericPlacementFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2GenericPlacementFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterSelector) {
		toSerialize["clusterSelector"] = o.ClusterSelector
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

type NullableV1alpha2GenericPlacementFields struct {
	value *V1alpha2GenericPlacementFields
	isSet bool
}

func (v NullableV1alpha2GenericPlacementFields) Get() *V1alpha2GenericPlacementFields {
	return v.value
}

func (v *NullableV1alpha2GenericPlacementFields) Set(val *V1alpha2GenericPlacementFields) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2GenericPlacementFields) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2GenericPlacementFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2GenericPlacementFields(val *V1alpha2GenericPlacementFields) *NullableV1alpha2GenericPlacementFields {
	return &NullableV1alpha2GenericPlacementFields{value: val, isSet: true}
}

func (v NullableV1alpha2GenericPlacementFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2GenericPlacementFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


