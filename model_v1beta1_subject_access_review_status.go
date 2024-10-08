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

// checks if the V1beta1SubjectAccessReviewStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1SubjectAccessReviewStatus{}

// V1beta1SubjectAccessReviewStatus struct for V1beta1SubjectAccessReviewStatus
type V1beta1SubjectAccessReviewStatus struct {
	Allowed bool `json:"allowed"`
	Denied *bool `json:"denied,omitempty"`
	EvaluationError *string `json:"evaluationError,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

type _V1beta1SubjectAccessReviewStatus V1beta1SubjectAccessReviewStatus

// NewV1beta1SubjectAccessReviewStatus instantiates a new V1beta1SubjectAccessReviewStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1SubjectAccessReviewStatus(allowed bool) *V1beta1SubjectAccessReviewStatus {
	this := V1beta1SubjectAccessReviewStatus{}
	this.Allowed = allowed
	return &this
}

// NewV1beta1SubjectAccessReviewStatusWithDefaults instantiates a new V1beta1SubjectAccessReviewStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1SubjectAccessReviewStatusWithDefaults() *V1beta1SubjectAccessReviewStatus {
	this := V1beta1SubjectAccessReviewStatus{}
	return &this
}

// GetAllowed returns the Allowed field value
func (o *V1beta1SubjectAccessReviewStatus) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *V1beta1SubjectAccessReviewStatus) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *V1beta1SubjectAccessReviewStatus) SetAllowed(v bool) {
	o.Allowed = v
}

// GetDenied returns the Denied field value if set, zero value otherwise.
func (o *V1beta1SubjectAccessReviewStatus) GetDenied() bool {
	if o == nil || IsNil(o.Denied) {
		var ret bool
		return ret
	}
	return *o.Denied
}

// GetDeniedOk returns a tuple with the Denied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1SubjectAccessReviewStatus) GetDeniedOk() (*bool, bool) {
	if o == nil || IsNil(o.Denied) {
		return nil, false
	}
	return o.Denied, true
}

// HasDenied returns a boolean if a field has been set.
func (o *V1beta1SubjectAccessReviewStatus) HasDenied() bool {
	if o != nil && !IsNil(o.Denied) {
		return true
	}

	return false
}

// SetDenied gets a reference to the given bool and assigns it to the Denied field.
func (o *V1beta1SubjectAccessReviewStatus) SetDenied(v bool) {
	o.Denied = &v
}

// GetEvaluationError returns the EvaluationError field value if set, zero value otherwise.
func (o *V1beta1SubjectAccessReviewStatus) GetEvaluationError() string {
	if o == nil || IsNil(o.EvaluationError) {
		var ret string
		return ret
	}
	return *o.EvaluationError
}

// GetEvaluationErrorOk returns a tuple with the EvaluationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1SubjectAccessReviewStatus) GetEvaluationErrorOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluationError) {
		return nil, false
	}
	return o.EvaluationError, true
}

// HasEvaluationError returns a boolean if a field has been set.
func (o *V1beta1SubjectAccessReviewStatus) HasEvaluationError() bool {
	if o != nil && !IsNil(o.EvaluationError) {
		return true
	}

	return false
}

// SetEvaluationError gets a reference to the given string and assigns it to the EvaluationError field.
func (o *V1beta1SubjectAccessReviewStatus) SetEvaluationError(v string) {
	o.EvaluationError = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1beta1SubjectAccessReviewStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1SubjectAccessReviewStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1beta1SubjectAccessReviewStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1beta1SubjectAccessReviewStatus) SetReason(v string) {
	o.Reason = &v
}

func (o V1beta1SubjectAccessReviewStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1SubjectAccessReviewStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed"] = o.Allowed
	if !IsNil(o.Denied) {
		toSerialize["denied"] = o.Denied
	}
	if !IsNil(o.EvaluationError) {
		toSerialize["evaluationError"] = o.EvaluationError
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *V1beta1SubjectAccessReviewStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowed",
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

	varV1beta1SubjectAccessReviewStatus := _V1beta1SubjectAccessReviewStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1beta1SubjectAccessReviewStatus)

	if err != nil {
		return err
	}

	*o = V1beta1SubjectAccessReviewStatus(varV1beta1SubjectAccessReviewStatus)

	return err
}

type NullableV1beta1SubjectAccessReviewStatus struct {
	value *V1beta1SubjectAccessReviewStatus
	isSet bool
}

func (v NullableV1beta1SubjectAccessReviewStatus) Get() *V1beta1SubjectAccessReviewStatus {
	return v.value
}

func (v *NullableV1beta1SubjectAccessReviewStatus) Set(val *V1beta1SubjectAccessReviewStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1SubjectAccessReviewStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1SubjectAccessReviewStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1SubjectAccessReviewStatus(val *V1beta1SubjectAccessReviewStatus) *NullableV1beta1SubjectAccessReviewStatus {
	return &NullableV1beta1SubjectAccessReviewStatus{value: val, isSet: true}
}

func (v NullableV1beta1SubjectAccessReviewStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1SubjectAccessReviewStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


