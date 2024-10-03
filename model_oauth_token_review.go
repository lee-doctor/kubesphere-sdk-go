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

// checks if the OauthTokenReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OauthTokenReview{}

// OauthTokenReview struct for OauthTokenReview
type OauthTokenReview struct {
	// Kubernetes API version
	ApiVersion string `json:"apiVersion"`
	// kind of the API object
	Kind string `json:"kind"`
	Spec *OauthSpec `json:"spec,omitempty"`
	Status *OauthStatus `json:"status,omitempty"`
}

type _OauthTokenReview OauthTokenReview

// NewOauthTokenReview instantiates a new OauthTokenReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthTokenReview(apiVersion string, kind string) *OauthTokenReview {
	this := OauthTokenReview{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	return &this
}

// NewOauthTokenReviewWithDefaults instantiates a new OauthTokenReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthTokenReviewWithDefaults() *OauthTokenReview {
	this := OauthTokenReview{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *OauthTokenReview) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *OauthTokenReview) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *OauthTokenReview) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *OauthTokenReview) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *OauthTokenReview) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *OauthTokenReview) SetKind(v string) {
	o.Kind = v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *OauthTokenReview) GetSpec() OauthSpec {
	if o == nil || IsNil(o.Spec) {
		var ret OauthSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenReview) GetSpecOk() (*OauthSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *OauthTokenReview) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given OauthSpec and assigns it to the Spec field.
func (o *OauthTokenReview) SetSpec(v OauthSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OauthTokenReview) GetStatus() OauthStatus {
	if o == nil || IsNil(o.Status) {
		var ret OauthStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenReview) GetStatusOk() (*OauthStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OauthTokenReview) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OauthStatus and assigns it to the Status field.
func (o *OauthTokenReview) SetStatus(v OauthStatus) {
	o.Status = &v
}

func (o OauthTokenReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OauthTokenReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *OauthTokenReview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
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

	varOauthTokenReview := _OauthTokenReview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOauthTokenReview)

	if err != nil {
		return err
	}

	*o = OauthTokenReview(varOauthTokenReview)

	return err
}

type NullableOauthTokenReview struct {
	value *OauthTokenReview
	isSet bool
}

func (v NullableOauthTokenReview) Get() *OauthTokenReview {
	return v.value
}

func (v *NullableOauthTokenReview) Set(val *OauthTokenReview) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthTokenReview) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthTokenReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthTokenReview(val *OauthTokenReview) *NullableOauthTokenReview {
	return &NullableOauthTokenReview{value: val, isSet: true}
}

func (v NullableOauthTokenReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthTokenReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


