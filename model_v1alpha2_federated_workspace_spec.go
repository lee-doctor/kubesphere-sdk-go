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

// checks if the V1alpha2FederatedWorkspaceSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2FederatedWorkspaceSpec{}

// V1alpha2FederatedWorkspaceSpec struct for V1alpha2FederatedWorkspaceSpec
type V1alpha2FederatedWorkspaceSpec struct {
	Overrides []V1alpha2GenericOverrideItem `json:"overrides,omitempty"`
	Placement V1alpha2GenericPlacementFields `json:"placement"`
	Template V1alpha2WorkspaceTemplateSpec `json:"template"`
}

type _V1alpha2FederatedWorkspaceSpec V1alpha2FederatedWorkspaceSpec

// NewV1alpha2FederatedWorkspaceSpec instantiates a new V1alpha2FederatedWorkspaceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2FederatedWorkspaceSpec(placement V1alpha2GenericPlacementFields, template V1alpha2WorkspaceTemplateSpec) *V1alpha2FederatedWorkspaceSpec {
	this := V1alpha2FederatedWorkspaceSpec{}
	this.Placement = placement
	this.Template = template
	return &this
}

// NewV1alpha2FederatedWorkspaceSpecWithDefaults instantiates a new V1alpha2FederatedWorkspaceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2FederatedWorkspaceSpecWithDefaults() *V1alpha2FederatedWorkspaceSpec {
	this := V1alpha2FederatedWorkspaceSpec{}
	return &this
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *V1alpha2FederatedWorkspaceSpec) GetOverrides() []V1alpha2GenericOverrideItem {
	if o == nil || IsNil(o.Overrides) {
		var ret []V1alpha2GenericOverrideItem
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2FederatedWorkspaceSpec) GetOverridesOk() ([]V1alpha2GenericOverrideItem, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *V1alpha2FederatedWorkspaceSpec) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []V1alpha2GenericOverrideItem and assigns it to the Overrides field.
func (o *V1alpha2FederatedWorkspaceSpec) SetOverrides(v []V1alpha2GenericOverrideItem) {
	o.Overrides = v
}

// GetPlacement returns the Placement field value
func (o *V1alpha2FederatedWorkspaceSpec) GetPlacement() V1alpha2GenericPlacementFields {
	if o == nil {
		var ret V1alpha2GenericPlacementFields
		return ret
	}

	return o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value
// and a boolean to check if the value has been set.
func (o *V1alpha2FederatedWorkspaceSpec) GetPlacementOk() (*V1alpha2GenericPlacementFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Placement, true
}

// SetPlacement sets field value
func (o *V1alpha2FederatedWorkspaceSpec) SetPlacement(v V1alpha2GenericPlacementFields) {
	o.Placement = v
}

// GetTemplate returns the Template field value
func (o *V1alpha2FederatedWorkspaceSpec) GetTemplate() V1alpha2WorkspaceTemplateSpec {
	if o == nil {
		var ret V1alpha2WorkspaceTemplateSpec
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *V1alpha2FederatedWorkspaceSpec) GetTemplateOk() (*V1alpha2WorkspaceTemplateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *V1alpha2FederatedWorkspaceSpec) SetTemplate(v V1alpha2WorkspaceTemplateSpec) {
	o.Template = v
}

func (o V1alpha2FederatedWorkspaceSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2FederatedWorkspaceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	toSerialize["placement"] = o.Placement
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *V1alpha2FederatedWorkspaceSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"placement",
		"template",
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

	varV1alpha2FederatedWorkspaceSpec := _V1alpha2FederatedWorkspaceSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2FederatedWorkspaceSpec)

	if err != nil {
		return err
	}

	*o = V1alpha2FederatedWorkspaceSpec(varV1alpha2FederatedWorkspaceSpec)

	return err
}

type NullableV1alpha2FederatedWorkspaceSpec struct {
	value *V1alpha2FederatedWorkspaceSpec
	isSet bool
}

func (v NullableV1alpha2FederatedWorkspaceSpec) Get() *V1alpha2FederatedWorkspaceSpec {
	return v.value
}

func (v *NullableV1alpha2FederatedWorkspaceSpec) Set(val *V1alpha2FederatedWorkspaceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2FederatedWorkspaceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2FederatedWorkspaceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2FederatedWorkspaceSpec(val *V1alpha2FederatedWorkspaceSpec) *NullableV1alpha2FederatedWorkspaceSpec {
	return &NullableV1alpha2FederatedWorkspaceSpec{value: val, isSet: true}
}

func (v NullableV1alpha2FederatedWorkspaceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2FederatedWorkspaceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


