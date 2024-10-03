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
	"bytes"
	"fmt"
)

// checks if the V1alpha2MetadataRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2MetadataRow{}

// V1alpha2MetadataRow struct for V1alpha2MetadataRow
type V1alpha2MetadataRow struct {
	DataType *string `json:"dataType,omitempty"`
	Id string `json:"id"`
	Label string `json:"label"`
	Priority *float64 `json:"priority,omitempty"`
	Truncate *int32 `json:"truncate,omitempty"`
	Value string `json:"value"`
}

type _V1alpha2MetadataRow V1alpha2MetadataRow

// NewV1alpha2MetadataRow instantiates a new V1alpha2MetadataRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2MetadataRow(id string, label string, value string) *V1alpha2MetadataRow {
	this := V1alpha2MetadataRow{}
	this.Id = id
	this.Label = label
	this.Value = value
	return &this
}

// NewV1alpha2MetadataRowWithDefaults instantiates a new V1alpha2MetadataRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2MetadataRowWithDefaults() *V1alpha2MetadataRow {
	this := V1alpha2MetadataRow{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *V1alpha2MetadataRow) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2MetadataRow) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *V1alpha2MetadataRow) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *V1alpha2MetadataRow) SetDataType(v string) {
	o.DataType = &v
}

// GetId returns the Id field value
func (o *V1alpha2MetadataRow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V1alpha2MetadataRow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V1alpha2MetadataRow) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *V1alpha2MetadataRow) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *V1alpha2MetadataRow) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *V1alpha2MetadataRow) SetLabel(v string) {
	o.Label = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *V1alpha2MetadataRow) GetPriority() float64 {
	if o == nil || IsNil(o.Priority) {
		var ret float64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2MetadataRow) GetPriorityOk() (*float64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *V1alpha2MetadataRow) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given float64 and assigns it to the Priority field.
func (o *V1alpha2MetadataRow) SetPriority(v float64) {
	o.Priority = &v
}

// GetTruncate returns the Truncate field value if set, zero value otherwise.
func (o *V1alpha2MetadataRow) GetTruncate() int32 {
	if o == nil || IsNil(o.Truncate) {
		var ret int32
		return ret
	}
	return *o.Truncate
}

// GetTruncateOk returns a tuple with the Truncate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2MetadataRow) GetTruncateOk() (*int32, bool) {
	if o == nil || IsNil(o.Truncate) {
		return nil, false
	}
	return o.Truncate, true
}

// HasTruncate returns a boolean if a field has been set.
func (o *V1alpha2MetadataRow) HasTruncate() bool {
	if o != nil && !IsNil(o.Truncate) {
		return true
	}

	return false
}

// SetTruncate gets a reference to the given int32 and assigns it to the Truncate field.
func (o *V1alpha2MetadataRow) SetTruncate(v int32) {
	o.Truncate = &v
}

// GetValue returns the Value field value
func (o *V1alpha2MetadataRow) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *V1alpha2MetadataRow) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *V1alpha2MetadataRow) SetValue(v string) {
	o.Value = v
}

func (o V1alpha2MetadataRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2MetadataRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Truncate) {
		toSerialize["truncate"] = o.Truncate
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *V1alpha2MetadataRow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"label",
		"value",
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

	varV1alpha2MetadataRow := _V1alpha2MetadataRow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2MetadataRow)

	if err != nil {
		return err
	}

	*o = V1alpha2MetadataRow(varV1alpha2MetadataRow)

	return err
}

type NullableV1alpha2MetadataRow struct {
	value *V1alpha2MetadataRow
	isSet bool
}

func (v NullableV1alpha2MetadataRow) Get() *V1alpha2MetadataRow {
	return v.value
}

func (v *NullableV1alpha2MetadataRow) Set(val *V1alpha2MetadataRow) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2MetadataRow) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2MetadataRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2MetadataRow(val *V1alpha2MetadataRow) *NullableV1alpha2MetadataRow {
	return &NullableV1alpha2MetadataRow{value: val, isSet: true}
}

func (v NullableV1alpha2MetadataRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2MetadataRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

