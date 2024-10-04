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

// checks if the V1ResourceQuotaSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ResourceQuotaSpec{}

// V1ResourceQuotaSpec ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
type V1ResourceQuotaSpec struct {
	// hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Hard *map[string]ResourceQuantity `json:"hard,omitempty"`
	ScopeSelector *V1ScopeSelector `json:"scopeSelector,omitempty"`
	// A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
	Scopes []string `json:"scopes,omitempty"`
}

// NewV1ResourceQuotaSpec instantiates a new V1ResourceQuotaSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ResourceQuotaSpec() *V1ResourceQuotaSpec {
	this := V1ResourceQuotaSpec{}
	return &this
}

// NewV1ResourceQuotaSpecWithDefaults instantiates a new V1ResourceQuotaSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ResourceQuotaSpecWithDefaults() *V1ResourceQuotaSpec {
	this := V1ResourceQuotaSpec{}
	return &this
}

// GetHard returns the Hard field value if set, zero value otherwise.
func (o *V1ResourceQuotaSpec) GetHard() map[string]ResourceQuantity {
	if o == nil || IsNil(o.Hard) {
		var ret map[string]ResourceQuantity
		return ret
	}
	return *o.Hard
}

// GetHardOk returns a tuple with the Hard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ResourceQuotaSpec) GetHardOk() (*map[string]ResourceQuantity, bool) {
	if o == nil || IsNil(o.Hard) {
		return nil, false
	}
	return o.Hard, true
}

// HasHard returns a boolean if a field has been set.
func (o *V1ResourceQuotaSpec) HasHard() bool {
	if o != nil && !IsNil(o.Hard) {
		return true
	}

	return false
}

// SetHard gets a reference to the given map[string]ResourceQuantity and assigns it to the Hard field.
func (o *V1ResourceQuotaSpec) SetHard(v map[string]ResourceQuantity) {
	o.Hard = &v
}

// GetScopeSelector returns the ScopeSelector field value if set, zero value otherwise.
func (o *V1ResourceQuotaSpec) GetScopeSelector() V1ScopeSelector {
	if o == nil || IsNil(o.ScopeSelector) {
		var ret V1ScopeSelector
		return ret
	}
	return *o.ScopeSelector
}

// GetScopeSelectorOk returns a tuple with the ScopeSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ResourceQuotaSpec) GetScopeSelectorOk() (*V1ScopeSelector, bool) {
	if o == nil || IsNil(o.ScopeSelector) {
		return nil, false
	}
	return o.ScopeSelector, true
}

// HasScopeSelector returns a boolean if a field has been set.
func (o *V1ResourceQuotaSpec) HasScopeSelector() bool {
	if o != nil && !IsNil(o.ScopeSelector) {
		return true
	}

	return false
}

// SetScopeSelector gets a reference to the given V1ScopeSelector and assigns it to the ScopeSelector field.
func (o *V1ResourceQuotaSpec) SetScopeSelector(v V1ScopeSelector) {
	o.ScopeSelector = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *V1ResourceQuotaSpec) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ResourceQuotaSpec) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *V1ResourceQuotaSpec) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *V1ResourceQuotaSpec) SetScopes(v []string) {
	o.Scopes = v
}

func (o V1ResourceQuotaSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ResourceQuotaSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hard) {
		toSerialize["hard"] = o.Hard
	}
	if !IsNil(o.ScopeSelector) {
		toSerialize["scopeSelector"] = o.ScopeSelector
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableV1ResourceQuotaSpec struct {
	value *V1ResourceQuotaSpec
	isSet bool
}

func (v NullableV1ResourceQuotaSpec) Get() *V1ResourceQuotaSpec {
	return v.value
}

func (v *NullableV1ResourceQuotaSpec) Set(val *V1ResourceQuotaSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ResourceQuotaSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ResourceQuotaSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ResourceQuotaSpec(val *V1ResourceQuotaSpec) *NullableV1ResourceQuotaSpec {
	return &NullableV1ResourceQuotaSpec{value: val, isSet: true}
}

func (v NullableV1ResourceQuotaSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ResourceQuotaSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

