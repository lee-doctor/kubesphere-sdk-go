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
	"time"
	"bytes"
	"fmt"
)

// checks if the V1alpha2ComponentStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2ComponentStatus{}

// V1alpha2ComponentStatus struct for V1alpha2ComponentStatus
type V1alpha2ComponentStatus struct {
	// the number of healthy backend components
	HealthyBackends int32 `json:"healthyBackends"`
	Label map[string]interface{} `json:"label"`
	// component name
	Name string `json:"name"`
	// the name of the namespace
	Namespace string `json:"namespace"`
	// self link
	SelfLink string `json:"selfLink"`
	// started time
	StartedAt time.Time `json:"startedAt"`
	// the total replicas of each backend system component
	TotalBackends int32 `json:"totalBackends"`
}

type _V1alpha2ComponentStatus V1alpha2ComponentStatus

// NewV1alpha2ComponentStatus instantiates a new V1alpha2ComponentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2ComponentStatus(healthyBackends int32, label map[string]interface{}, name string, namespace string, selfLink string, startedAt time.Time, totalBackends int32) *V1alpha2ComponentStatus {
	this := V1alpha2ComponentStatus{}
	this.HealthyBackends = healthyBackends
	this.Label = label
	this.Name = name
	this.Namespace = namespace
	this.SelfLink = selfLink
	this.StartedAt = startedAt
	this.TotalBackends = totalBackends
	return &this
}

// NewV1alpha2ComponentStatusWithDefaults instantiates a new V1alpha2ComponentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2ComponentStatusWithDefaults() *V1alpha2ComponentStatus {
	this := V1alpha2ComponentStatus{}
	return &this
}

// GetHealthyBackends returns the HealthyBackends field value
func (o *V1alpha2ComponentStatus) GetHealthyBackends() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HealthyBackends
}

// GetHealthyBackendsOk returns a tuple with the HealthyBackends field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ComponentStatus) GetHealthyBackendsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HealthyBackends, true
}

// SetHealthyBackends sets field value
func (o *V1alpha2ComponentStatus) SetHealthyBackends(v int32) {
	o.HealthyBackends = v
}

// GetLabel returns the Label field value
func (o *V1alpha2ComponentStatus) GetLabel() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ComponentStatus) GetLabelOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Label, true
}

// SetLabel sets field value
func (o *V1alpha2ComponentStatus) SetLabel(v map[string]interface{}) {
	o.Label = v
}

// GetName returns the Name field value
func (o *V1alpha2ComponentStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ComponentStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1alpha2ComponentStatus) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *V1alpha2ComponentStatus) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ComponentStatus) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *V1alpha2ComponentStatus) SetNamespace(v string) {
	o.Namespace = v
}

// GetSelfLink returns the SelfLink field value
func (o *V1alpha2ComponentStatus) GetSelfLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SelfLink
}

// GetSelfLinkOk returns a tuple with the SelfLink field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ComponentStatus) GetSelfLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfLink, true
}

// SetSelfLink sets field value
func (o *V1alpha2ComponentStatus) SetSelfLink(v string) {
	o.SelfLink = v
}

// GetStartedAt returns the StartedAt field value
func (o *V1alpha2ComponentStatus) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ComponentStatus) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *V1alpha2ComponentStatus) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetTotalBackends returns the TotalBackends field value
func (o *V1alpha2ComponentStatus) GetTotalBackends() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalBackends
}

// GetTotalBackendsOk returns a tuple with the TotalBackends field value
// and a boolean to check if the value has been set.
func (o *V1alpha2ComponentStatus) GetTotalBackendsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBackends, true
}

// SetTotalBackends sets field value
func (o *V1alpha2ComponentStatus) SetTotalBackends(v int32) {
	o.TotalBackends = v
}

func (o V1alpha2ComponentStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2ComponentStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["healthyBackends"] = o.HealthyBackends
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	toSerialize["selfLink"] = o.SelfLink
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["totalBackends"] = o.TotalBackends
	return toSerialize, nil
}

func (o *V1alpha2ComponentStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"healthyBackends",
		"label",
		"name",
		"namespace",
		"selfLink",
		"startedAt",
		"totalBackends",
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

	varV1alpha2ComponentStatus := _V1alpha2ComponentStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1alpha2ComponentStatus)

	if err != nil {
		return err
	}

	*o = V1alpha2ComponentStatus(varV1alpha2ComponentStatus)

	return err
}

type NullableV1alpha2ComponentStatus struct {
	value *V1alpha2ComponentStatus
	isSet bool
}

func (v NullableV1alpha2ComponentStatus) Get() *V1alpha2ComponentStatus {
	return v.value
}

func (v *NullableV1alpha2ComponentStatus) Set(val *V1alpha2ComponentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2ComponentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2ComponentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2ComponentStatus(val *V1alpha2ComponentStatus) *NullableV1alpha2ComponentStatus {
	return &NullableV1alpha2ComponentStatus{value: val, isSet: true}
}

func (v NullableV1alpha2ComponentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2ComponentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


