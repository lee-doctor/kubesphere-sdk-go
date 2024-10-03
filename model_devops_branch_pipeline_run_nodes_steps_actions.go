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

// checks if the DevopsBranchPipelineRunNodesStepsActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsBranchPipelineRunNodesStepsActions{}

// DevopsBranchPipelineRunNodesStepsActions struct for DevopsBranchPipelineRunNodesStepsActions
type DevopsBranchPipelineRunNodesStepsActions struct {
	Class *string `json:"_class,omitempty"`
	Links *Links `json:"_links,omitempty"`
	UrlName *string `json:"urlName,omitempty"`
}

// NewDevopsBranchPipelineRunNodesStepsActions instantiates a new DevopsBranchPipelineRunNodesStepsActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsBranchPipelineRunNodesStepsActions() *DevopsBranchPipelineRunNodesStepsActions {
	this := DevopsBranchPipelineRunNodesStepsActions{}
	return &this
}

// NewDevopsBranchPipelineRunNodesStepsActionsWithDefaults instantiates a new DevopsBranchPipelineRunNodesStepsActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsBranchPipelineRunNodesStepsActionsWithDefaults() *DevopsBranchPipelineRunNodesStepsActions {
	this := DevopsBranchPipelineRunNodesStepsActions{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *DevopsBranchPipelineRunNodesStepsActions) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsBranchPipelineRunNodesStepsActions) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *DevopsBranchPipelineRunNodesStepsActions) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *DevopsBranchPipelineRunNodesStepsActions) SetClass(v string) {
	o.Class = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DevopsBranchPipelineRunNodesStepsActions) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsBranchPipelineRunNodesStepsActions) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DevopsBranchPipelineRunNodesStepsActions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *DevopsBranchPipelineRunNodesStepsActions) SetLinks(v Links) {
	o.Links = &v
}

// GetUrlName returns the UrlName field value if set, zero value otherwise.
func (o *DevopsBranchPipelineRunNodesStepsActions) GetUrlName() string {
	if o == nil || IsNil(o.UrlName) {
		var ret string
		return ret
	}
	return *o.UrlName
}

// GetUrlNameOk returns a tuple with the UrlName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsBranchPipelineRunNodesStepsActions) GetUrlNameOk() (*string, bool) {
	if o == nil || IsNil(o.UrlName) {
		return nil, false
	}
	return o.UrlName, true
}

// HasUrlName returns a boolean if a field has been set.
func (o *DevopsBranchPipelineRunNodesStepsActions) HasUrlName() bool {
	if o != nil && !IsNil(o.UrlName) {
		return true
	}

	return false
}

// SetUrlName gets a reference to the given string and assigns it to the UrlName field.
func (o *DevopsBranchPipelineRunNodesStepsActions) SetUrlName(v string) {
	o.UrlName = &v
}

func (o DevopsBranchPipelineRunNodesStepsActions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsBranchPipelineRunNodesStepsActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["_class"] = o.Class
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.UrlName) {
		toSerialize["urlName"] = o.UrlName
	}
	return toSerialize, nil
}

type NullableDevopsBranchPipelineRunNodesStepsActions struct {
	value *DevopsBranchPipelineRunNodesStepsActions
	isSet bool
}

func (v NullableDevopsBranchPipelineRunNodesStepsActions) Get() *DevopsBranchPipelineRunNodesStepsActions {
	return v.value
}

func (v *NullableDevopsBranchPipelineRunNodesStepsActions) Set(val *DevopsBranchPipelineRunNodesStepsActions) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsBranchPipelineRunNodesStepsActions) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsBranchPipelineRunNodesStepsActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsBranchPipelineRunNodesStepsActions(val *DevopsBranchPipelineRunNodesStepsActions) *NullableDevopsBranchPipelineRunNodesStepsActions {
	return &NullableDevopsBranchPipelineRunNodesStepsActions{value: val, isSet: true}
}

func (v NullableDevopsBranchPipelineRunNodesStepsActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsBranchPipelineRunNodesStepsActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

