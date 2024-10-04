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

// checks if the V2RepositoryTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2RepositoryTags{}

// V2RepositoryTags struct for V2RepositoryTags
type V2RepositoryTags struct {
	Registry string `json:"registry"`
	Repository string `json:"repository"`
	Tags []string `json:"tags"`
	Total int32 `json:"total"`
}

type _V2RepositoryTags V2RepositoryTags

// NewV2RepositoryTags instantiates a new V2RepositoryTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2RepositoryTags(registry string, repository string, tags []string, total int32) *V2RepositoryTags {
	this := V2RepositoryTags{}
	this.Registry = registry
	this.Repository = repository
	this.Tags = tags
	this.Total = total
	return &this
}

// NewV2RepositoryTagsWithDefaults instantiates a new V2RepositoryTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2RepositoryTagsWithDefaults() *V2RepositoryTags {
	this := V2RepositoryTags{}
	return &this
}

// GetRegistry returns the Registry field value
func (o *V2RepositoryTags) GetRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *V2RepositoryTags) GetRegistryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *V2RepositoryTags) SetRegistry(v string) {
	o.Registry = v
}

// GetRepository returns the Repository field value
func (o *V2RepositoryTags) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *V2RepositoryTags) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *V2RepositoryTags) SetRepository(v string) {
	o.Repository = v
}

// GetTags returns the Tags field value
func (o *V2RepositoryTags) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *V2RepositoryTags) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *V2RepositoryTags) SetTags(v []string) {
	o.Tags = v
}

// GetTotal returns the Total field value
func (o *V2RepositoryTags) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *V2RepositoryTags) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *V2RepositoryTags) SetTotal(v int32) {
	o.Total = v
}

func (o V2RepositoryTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2RepositoryTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["registry"] = o.Registry
	toSerialize["repository"] = o.Repository
	toSerialize["tags"] = o.Tags
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *V2RepositoryTags) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"registry",
		"repository",
		"tags",
		"total",
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

	varV2RepositoryTags := _V2RepositoryTags{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2RepositoryTags)

	if err != nil {
		return err
	}

	*o = V2RepositoryTags(varV2RepositoryTags)

	return err
}

type NullableV2RepositoryTags struct {
	value *V2RepositoryTags
	isSet bool
}

func (v NullableV2RepositoryTags) Get() *V2RepositoryTags {
	return v.value
}

func (v *NullableV2RepositoryTags) Set(val *V2RepositoryTags) {
	v.value = val
	v.isSet = true
}

func (v NullableV2RepositoryTags) IsSet() bool {
	return v.isSet
}

func (v *NullableV2RepositoryTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2RepositoryTags(val *V2RepositoryTags) *NullableV2RepositoryTags {
	return &NullableV2RepositoryTags{value: val, isSet: true}
}

func (v NullableV2RepositoryTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2RepositoryTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

