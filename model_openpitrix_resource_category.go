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

// checks if the OpenpitrixResourceCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixResourceCategory{}

// OpenpitrixResourceCategory struct for OpenpitrixResourceCategory
type OpenpitrixResourceCategory struct {
	CategoryId *string `json:"category_id,omitempty"`
	CreateTime *string `json:"create_time,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusTime *string `json:"status_time,omitempty"`
}

// NewOpenpitrixResourceCategory instantiates a new OpenpitrixResourceCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixResourceCategory() *OpenpitrixResourceCategory {
	this := OpenpitrixResourceCategory{}
	return &this
}

// NewOpenpitrixResourceCategoryWithDefaults instantiates a new OpenpitrixResourceCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixResourceCategoryWithDefaults() *OpenpitrixResourceCategory {
	this := OpenpitrixResourceCategory{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *OpenpitrixResourceCategory) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixResourceCategory) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *OpenpitrixResourceCategory) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *OpenpitrixResourceCategory) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *OpenpitrixResourceCategory) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixResourceCategory) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *OpenpitrixResourceCategory) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *OpenpitrixResourceCategory) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *OpenpitrixResourceCategory) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixResourceCategory) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *OpenpitrixResourceCategory) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *OpenpitrixResourceCategory) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenpitrixResourceCategory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixResourceCategory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenpitrixResourceCategory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenpitrixResourceCategory) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OpenpitrixResourceCategory) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixResourceCategory) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OpenpitrixResourceCategory) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OpenpitrixResourceCategory) SetStatus(v string) {
	o.Status = &v
}

// GetStatusTime returns the StatusTime field value if set, zero value otherwise.
func (o *OpenpitrixResourceCategory) GetStatusTime() string {
	if o == nil || IsNil(o.StatusTime) {
		var ret string
		return ret
	}
	return *o.StatusTime
}

// GetStatusTimeOk returns a tuple with the StatusTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixResourceCategory) GetStatusTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StatusTime) {
		return nil, false
	}
	return o.StatusTime, true
}

// HasStatusTime returns a boolean if a field has been set.
func (o *OpenpitrixResourceCategory) HasStatusTime() bool {
	if o != nil && !IsNil(o.StatusTime) {
		return true
	}

	return false
}

// SetStatusTime gets a reference to the given string and assigns it to the StatusTime field.
func (o *OpenpitrixResourceCategory) SetStatusTime(v string) {
	o.StatusTime = &v
}

func (o OpenpitrixResourceCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixResourceCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusTime) {
		toSerialize["status_time"] = o.StatusTime
	}
	return toSerialize, nil
}

type NullableOpenpitrixResourceCategory struct {
	value *OpenpitrixResourceCategory
	isSet bool
}

func (v NullableOpenpitrixResourceCategory) Get() *OpenpitrixResourceCategory {
	return v.value
}

func (v *NullableOpenpitrixResourceCategory) Set(val *OpenpitrixResourceCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixResourceCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixResourceCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixResourceCategory(val *OpenpitrixResourceCategory) *NullableOpenpitrixResourceCategory {
	return &NullableOpenpitrixResourceCategory{value: val, isSet: true}
}

func (v NullableOpenpitrixResourceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixResourceCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


