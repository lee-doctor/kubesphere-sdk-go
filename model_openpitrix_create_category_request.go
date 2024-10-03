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

// checks if the OpenpitrixCreateCategoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixCreateCategoryRequest{}

// OpenpitrixCreateCategoryRequest struct for OpenpitrixCreateCategoryRequest
type OpenpitrixCreateCategoryRequest struct {
	Description *string `json:"description,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewOpenpitrixCreateCategoryRequest instantiates a new OpenpitrixCreateCategoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixCreateCategoryRequest() *OpenpitrixCreateCategoryRequest {
	this := OpenpitrixCreateCategoryRequest{}
	return &this
}

// NewOpenpitrixCreateCategoryRequestWithDefaults instantiates a new OpenpitrixCreateCategoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixCreateCategoryRequestWithDefaults() *OpenpitrixCreateCategoryRequest {
	this := OpenpitrixCreateCategoryRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenpitrixCreateCategoryRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateCategoryRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenpitrixCreateCategoryRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenpitrixCreateCategoryRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *OpenpitrixCreateCategoryRequest) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateCategoryRequest) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *OpenpitrixCreateCategoryRequest) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *OpenpitrixCreateCategoryRequest) SetIcon(v string) {
	o.Icon = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *OpenpitrixCreateCategoryRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateCategoryRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *OpenpitrixCreateCategoryRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *OpenpitrixCreateCategoryRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenpitrixCreateCategoryRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateCategoryRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenpitrixCreateCategoryRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenpitrixCreateCategoryRequest) SetName(v string) {
	o.Name = &v
}

func (o OpenpitrixCreateCategoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixCreateCategoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableOpenpitrixCreateCategoryRequest struct {
	value *OpenpitrixCreateCategoryRequest
	isSet bool
}

func (v NullableOpenpitrixCreateCategoryRequest) Get() *OpenpitrixCreateCategoryRequest {
	return v.value
}

func (v *NullableOpenpitrixCreateCategoryRequest) Set(val *OpenpitrixCreateCategoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixCreateCategoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixCreateCategoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixCreateCategoryRequest(val *OpenpitrixCreateCategoryRequest) *NullableOpenpitrixCreateCategoryRequest {
	return &NullableOpenpitrixCreateCategoryRequest{value: val, isSet: true}
}

func (v NullableOpenpitrixCreateCategoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixCreateCategoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


