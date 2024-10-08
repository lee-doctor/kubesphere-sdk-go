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

// checks if the V1ClusterTrustBundleProjection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ClusterTrustBundleProjection{}

// V1ClusterTrustBundleProjection ClusterTrustBundleProjection describes how to select a set of ClusterTrustBundle objects and project their contents into the pod filesystem.
type V1ClusterTrustBundleProjection struct {
	LabelSelector *V1LabelSelector `json:"labelSelector,omitempty"`
	// Select a single ClusterTrustBundle by object name.  Mutually-exclusive with signerName and labelSelector.
	Name *string `json:"name,omitempty"`
	// If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available.  If using name, then the named ClusterTrustBundle is allowed not to exist.  If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.
	Optional *bool `json:"optional,omitempty"`
	// Relative path from the volume root to write the bundle.
	Path string `json:"path"`
	// Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name.  The contents of all selected ClusterTrustBundles will be unified and deduplicated.
	SignerName *string `json:"signerName,omitempty"`
}

type _V1ClusterTrustBundleProjection V1ClusterTrustBundleProjection

// NewV1ClusterTrustBundleProjection instantiates a new V1ClusterTrustBundleProjection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ClusterTrustBundleProjection(path string) *V1ClusterTrustBundleProjection {
	this := V1ClusterTrustBundleProjection{}
	this.Path = path
	return &this
}

// NewV1ClusterTrustBundleProjectionWithDefaults instantiates a new V1ClusterTrustBundleProjection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ClusterTrustBundleProjectionWithDefaults() *V1ClusterTrustBundleProjection {
	this := V1ClusterTrustBundleProjection{}
	return &this
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *V1ClusterTrustBundleProjection) GetLabelSelector() V1LabelSelector {
	if o == nil || IsNil(o.LabelSelector) {
		var ret V1LabelSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClusterTrustBundleProjection) GetLabelSelectorOk() (*V1LabelSelector, bool) {
	if o == nil || IsNil(o.LabelSelector) {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *V1ClusterTrustBundleProjection) HasLabelSelector() bool {
	if o != nil && !IsNil(o.LabelSelector) {
		return true
	}

	return false
}

// SetLabelSelector gets a reference to the given V1LabelSelector and assigns it to the LabelSelector field.
func (o *V1ClusterTrustBundleProjection) SetLabelSelector(v V1LabelSelector) {
	o.LabelSelector = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ClusterTrustBundleProjection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClusterTrustBundleProjection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ClusterTrustBundleProjection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ClusterTrustBundleProjection) SetName(v string) {
	o.Name = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *V1ClusterTrustBundleProjection) GetOptional() bool {
	if o == nil || IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClusterTrustBundleProjection) GetOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *V1ClusterTrustBundleProjection) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *V1ClusterTrustBundleProjection) SetOptional(v bool) {
	o.Optional = &v
}

// GetPath returns the Path field value
func (o *V1ClusterTrustBundleProjection) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *V1ClusterTrustBundleProjection) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *V1ClusterTrustBundleProjection) SetPath(v string) {
	o.Path = v
}

// GetSignerName returns the SignerName field value if set, zero value otherwise.
func (o *V1ClusterTrustBundleProjection) GetSignerName() string {
	if o == nil || IsNil(o.SignerName) {
		var ret string
		return ret
	}
	return *o.SignerName
}

// GetSignerNameOk returns a tuple with the SignerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClusterTrustBundleProjection) GetSignerNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignerName) {
		return nil, false
	}
	return o.SignerName, true
}

// HasSignerName returns a boolean if a field has been set.
func (o *V1ClusterTrustBundleProjection) HasSignerName() bool {
	if o != nil && !IsNil(o.SignerName) {
		return true
	}

	return false
}

// SetSignerName gets a reference to the given string and assigns it to the SignerName field.
func (o *V1ClusterTrustBundleProjection) SetSignerName(v string) {
	o.SignerName = &v
}

func (o V1ClusterTrustBundleProjection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ClusterTrustBundleProjection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelSelector) {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.SignerName) {
		toSerialize["signerName"] = o.SignerName
	}
	return toSerialize, nil
}

func (o *V1ClusterTrustBundleProjection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
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

	varV1ClusterTrustBundleProjection := _V1ClusterTrustBundleProjection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ClusterTrustBundleProjection)

	if err != nil {
		return err
	}

	*o = V1ClusterTrustBundleProjection(varV1ClusterTrustBundleProjection)

	return err
}

type NullableV1ClusterTrustBundleProjection struct {
	value *V1ClusterTrustBundleProjection
	isSet bool
}

func (v NullableV1ClusterTrustBundleProjection) Get() *V1ClusterTrustBundleProjection {
	return v.value
}

func (v *NullableV1ClusterTrustBundleProjection) Set(val *V1ClusterTrustBundleProjection) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ClusterTrustBundleProjection) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ClusterTrustBundleProjection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ClusterTrustBundleProjection(val *V1ClusterTrustBundleProjection) *NullableV1ClusterTrustBundleProjection {
	return &NullableV1ClusterTrustBundleProjection{value: val, isSet: true}
}

func (v NullableV1ClusterTrustBundleProjection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ClusterTrustBundleProjection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


