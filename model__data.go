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

// checks if the Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Data{}

// Data struct for Data
type Data struct {
	Errors []DataErrors `json:"errors,omitempty"`
	// jenkinsfile
	Jenkinsfile *string `json:"jenkinsfile,omitempty"`
	// result e.g. success
	Result *string `json:"result,omitempty"`
}

// NewData instantiates a new Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewData() *Data {
	this := Data{}
	return &this
}

// NewDataWithDefaults instantiates a new Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataWithDefaults() *Data {
	this := Data{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Data) GetErrors() []DataErrors {
	if o == nil || IsNil(o.Errors) {
		var ret []DataErrors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetErrorsOk() ([]DataErrors, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Data) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []DataErrors and assigns it to the Errors field.
func (o *Data) SetErrors(v []DataErrors) {
	o.Errors = v
}

// GetJenkinsfile returns the Jenkinsfile field value if set, zero value otherwise.
func (o *Data) GetJenkinsfile() string {
	if o == nil || IsNil(o.Jenkinsfile) {
		var ret string
		return ret
	}
	return *o.Jenkinsfile
}

// GetJenkinsfileOk returns a tuple with the Jenkinsfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetJenkinsfileOk() (*string, bool) {
	if o == nil || IsNil(o.Jenkinsfile) {
		return nil, false
	}
	return o.Jenkinsfile, true
}

// HasJenkinsfile returns a boolean if a field has been set.
func (o *Data) HasJenkinsfile() bool {
	if o != nil && !IsNil(o.Jenkinsfile) {
		return true
	}

	return false
}

// SetJenkinsfile gets a reference to the given string and assigns it to the Jenkinsfile field.
func (o *Data) SetJenkinsfile(v string) {
	o.Jenkinsfile = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Data) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Data) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Data) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *Data) SetResult(v string) {
	o.Result = &v
}

func (o Data) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Jenkinsfile) {
		toSerialize["jenkinsfile"] = o.Jenkinsfile
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableData struct {
	value *Data
	isSet bool
}

func (v NullableData) Get() *Data {
	return v.value
}

func (v *NullableData) Set(val *Data) {
	v.value = val
	v.isSet = true
}

func (v NullableData) IsSet() bool {
	return v.isSet
}

func (v *NullableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableData(val *Data) *NullableData {
	return &NullableData{value: val, isSet: true}
}

func (v NullableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

