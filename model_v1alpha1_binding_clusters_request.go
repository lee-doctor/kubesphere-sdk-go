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

// checks if the V1alpha1BindingClustersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1BindingClustersRequest{}

// V1alpha1BindingClustersRequest struct for V1alpha1BindingClustersRequest
type V1alpha1BindingClustersRequest struct {
	Clusters []string `json:"clusters"`
	Labels []string `json:"labels"`
}

type _V1alpha1BindingClustersRequest V1alpha1BindingClustersRequest

// NewV1alpha1BindingClustersRequest instantiates a new V1alpha1BindingClustersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1BindingClustersRequest(clusters []string, labels []string) *V1alpha1BindingClustersRequest {
	this := V1alpha1BindingClustersRequest{}
	this.Clusters = clusters
	this.Labels = labels
	return &this
}

// NewV1alpha1BindingClustersRequestWithDefaults instantiates a new V1alpha1BindingClustersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1BindingClustersRequestWithDefaults() *V1alpha1BindingClustersRequest {
	this := V1alpha1BindingClustersRequest{}
	return &this
}

// GetClusters returns the Clusters field value
func (o *V1alpha1BindingClustersRequest) GetClusters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *V1alpha1BindingClustersRequest) GetClustersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clusters, true
}

// SetClusters sets field value
func (o *V1alpha1BindingClustersRequest) SetClusters(v []string) {
	o.Clusters = v
}

// GetLabels returns the Labels field value
func (o *V1alpha1BindingClustersRequest) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *V1alpha1BindingClustersRequest) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *V1alpha1BindingClustersRequest) SetLabels(v []string) {
	o.Labels = v
}

func (o V1alpha1BindingClustersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1BindingClustersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusters"] = o.Clusters
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

func (o *V1alpha1BindingClustersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusters",
		"labels",
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

	varV1alpha1BindingClustersRequest := _V1alpha1BindingClustersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha1BindingClustersRequest)

	if err != nil {
		return err
	}

	*o = V1alpha1BindingClustersRequest(varV1alpha1BindingClustersRequest)

	return err
}

type NullableV1alpha1BindingClustersRequest struct {
	value *V1alpha1BindingClustersRequest
	isSet bool
}

func (v NullableV1alpha1BindingClustersRequest) Get() *V1alpha1BindingClustersRequest {
	return v.value
}

func (v *NullableV1alpha1BindingClustersRequest) Set(val *V1alpha1BindingClustersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1BindingClustersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1BindingClustersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1BindingClustersRequest(val *V1alpha1BindingClustersRequest) *NullableV1alpha1BindingClustersRequest {
	return &NullableV1alpha1BindingClustersRequest{value: val, isSet: true}
}

func (v NullableV1alpha1BindingClustersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1BindingClustersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


