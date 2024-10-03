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

// checks if the V1alpha2Control type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2Control{}

// V1alpha2Control struct for V1alpha2Control
type V1alpha2Control struct {
	Confirmation *string `json:"confirmation,omitempty"`
	Human string `json:"human"`
	Icon string `json:"icon"`
	Id string `json:"id"`
	Rank int32 `json:"rank"`
}

type _V1alpha2Control V1alpha2Control

// NewV1alpha2Control instantiates a new V1alpha2Control object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2Control(human string, icon string, id string, rank int32) *V1alpha2Control {
	this := V1alpha2Control{}
	this.Human = human
	this.Icon = icon
	this.Id = id
	this.Rank = rank
	return &this
}

// NewV1alpha2ControlWithDefaults instantiates a new V1alpha2Control object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2ControlWithDefaults() *V1alpha2Control {
	this := V1alpha2Control{}
	return &this
}

// GetConfirmation returns the Confirmation field value if set, zero value otherwise.
func (o *V1alpha2Control) GetConfirmation() string {
	if o == nil || IsNil(o.Confirmation) {
		var ret string
		return ret
	}
	return *o.Confirmation
}

// GetConfirmationOk returns a tuple with the Confirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2Control) GetConfirmationOk() (*string, bool) {
	if o == nil || IsNil(o.Confirmation) {
		return nil, false
	}
	return o.Confirmation, true
}

// HasConfirmation returns a boolean if a field has been set.
func (o *V1alpha2Control) HasConfirmation() bool {
	if o != nil && !IsNil(o.Confirmation) {
		return true
	}

	return false
}

// SetConfirmation gets a reference to the given string and assigns it to the Confirmation field.
func (o *V1alpha2Control) SetConfirmation(v string) {
	o.Confirmation = &v
}

// GetHuman returns the Human field value
func (o *V1alpha2Control) GetHuman() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Human
}

// GetHumanOk returns a tuple with the Human field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Control) GetHumanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Human, true
}

// SetHuman sets field value
func (o *V1alpha2Control) SetHuman(v string) {
	o.Human = v
}

// GetIcon returns the Icon field value
func (o *V1alpha2Control) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Control) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *V1alpha2Control) SetIcon(v string) {
	o.Icon = v
}

// GetId returns the Id field value
func (o *V1alpha2Control) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Control) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V1alpha2Control) SetId(v string) {
	o.Id = v
}

// GetRank returns the Rank field value
func (o *V1alpha2Control) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *V1alpha2Control) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *V1alpha2Control) SetRank(v int32) {
	o.Rank = v
}

func (o V1alpha2Control) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2Control) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Confirmation) {
		toSerialize["confirmation"] = o.Confirmation
	}
	toSerialize["human"] = o.Human
	toSerialize["icon"] = o.Icon
	toSerialize["id"] = o.Id
	toSerialize["rank"] = o.Rank
	return toSerialize, nil
}

func (o *V1alpha2Control) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"human",
		"icon",
		"id",
		"rank",
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

	varV1alpha2Control := _V1alpha2Control{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2Control)

	if err != nil {
		return err
	}

	*o = V1alpha2Control(varV1alpha2Control)

	return err
}

type NullableV1alpha2Control struct {
	value *V1alpha2Control
	isSet bool
}

func (v NullableV1alpha2Control) Get() *V1alpha2Control {
	return v.value
}

func (v *NullableV1alpha2Control) Set(val *V1alpha2Control) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2Control) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2Control) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2Control(val *V1alpha2Control) *NullableV1alpha2Control {
	return &NullableV1alpha2Control{value: val, isSet: true}
}

func (v NullableV1alpha2Control) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2Control) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


