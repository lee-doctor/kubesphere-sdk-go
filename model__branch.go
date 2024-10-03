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

// checks if the Branch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Branch{}

// Branch struct for Branch
type Branch struct {
	// primary or not
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// issues
	Issues []map[string]interface{} `json:"issues,omitempty"`
	// url
	Url *string `json:"url,omitempty"`
}

// NewBranch instantiates a new Branch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranch() *Branch {
	this := Branch{}
	return &this
}

// NewBranchWithDefaults instantiates a new Branch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchWithDefaults() *Branch {
	this := Branch{}
	return &this
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *Branch) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *Branch) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *Branch) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *Branch) GetIssues() []map[string]interface{} {
	if o == nil || IsNil(o.Issues) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetIssuesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *Branch) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []map[string]interface{} and assigns it to the Issues field.
func (o *Branch) SetIssues(v []map[string]interface{}) {
	o.Issues = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Branch) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Branch) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Branch) SetUrl(v string) {
	o.Url = &v
}

func (o Branch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Branch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableBranch struct {
	value *Branch
	isSet bool
}

func (v NullableBranch) Get() *Branch {
	return v.value
}

func (v *NullableBranch) Set(val *Branch) {
	v.value = val
	v.isSet = true
}

func (v NullableBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranch(val *Branch) *NullableBranch {
	return &NullableBranch{value: val, isSet: true}
}

func (v NullableBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


