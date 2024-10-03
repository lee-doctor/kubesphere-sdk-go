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

// checks if the DevopsCheckCronRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsCheckCronRes{}

// DevopsCheckCronRes struct for DevopsCheckCronRes
type DevopsCheckCronRes struct {
	// last run time.
	LastTime *string `json:"lastTime,omitempty"`
	// message
	Message *string `json:"message,omitempty"`
	// next run time.
	NextTime *string `json:"nextTime,omitempty"`
	// result e.g. ok, error
	Result *string `json:"result,omitempty"`
}

// NewDevopsCheckCronRes instantiates a new DevopsCheckCronRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsCheckCronRes() *DevopsCheckCronRes {
	this := DevopsCheckCronRes{}
	return &this
}

// NewDevopsCheckCronResWithDefaults instantiates a new DevopsCheckCronRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsCheckCronResWithDefaults() *DevopsCheckCronRes {
	this := DevopsCheckCronRes{}
	return &this
}

// GetLastTime returns the LastTime field value if set, zero value otherwise.
func (o *DevopsCheckCronRes) GetLastTime() string {
	if o == nil || IsNil(o.LastTime) {
		var ret string
		return ret
	}
	return *o.LastTime
}

// GetLastTimeOk returns a tuple with the LastTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCheckCronRes) GetLastTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastTime) {
		return nil, false
	}
	return o.LastTime, true
}

// HasLastTime returns a boolean if a field has been set.
func (o *DevopsCheckCronRes) HasLastTime() bool {
	if o != nil && !IsNil(o.LastTime) {
		return true
	}

	return false
}

// SetLastTime gets a reference to the given string and assigns it to the LastTime field.
func (o *DevopsCheckCronRes) SetLastTime(v string) {
	o.LastTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DevopsCheckCronRes) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCheckCronRes) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DevopsCheckCronRes) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DevopsCheckCronRes) SetMessage(v string) {
	o.Message = &v
}

// GetNextTime returns the NextTime field value if set, zero value otherwise.
func (o *DevopsCheckCronRes) GetNextTime() string {
	if o == nil || IsNil(o.NextTime) {
		var ret string
		return ret
	}
	return *o.NextTime
}

// GetNextTimeOk returns a tuple with the NextTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCheckCronRes) GetNextTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextTime) {
		return nil, false
	}
	return o.NextTime, true
}

// HasNextTime returns a boolean if a field has been set.
func (o *DevopsCheckCronRes) HasNextTime() bool {
	if o != nil && !IsNil(o.NextTime) {
		return true
	}

	return false
}

// SetNextTime gets a reference to the given string and assigns it to the NextTime field.
func (o *DevopsCheckCronRes) SetNextTime(v string) {
	o.NextTime = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DevopsCheckCronRes) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCheckCronRes) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DevopsCheckCronRes) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DevopsCheckCronRes) SetResult(v string) {
	o.Result = &v
}

func (o DevopsCheckCronRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsCheckCronRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastTime) {
		toSerialize["lastTime"] = o.LastTime
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.NextTime) {
		toSerialize["nextTime"] = o.NextTime
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableDevopsCheckCronRes struct {
	value *DevopsCheckCronRes
	isSet bool
}

func (v NullableDevopsCheckCronRes) Get() *DevopsCheckCronRes {
	return v.value
}

func (v *NullableDevopsCheckCronRes) Set(val *DevopsCheckCronRes) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsCheckCronRes) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsCheckCronRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsCheckCronRes(val *DevopsCheckCronRes) *NullableDevopsCheckCronRes {
	return &NullableDevopsCheckCronRes{value: val, isSet: true}
}

func (v NullableDevopsCheckCronRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsCheckCronRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


