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

// checks if the ApiWorkloads type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiWorkloads{}

// ApiWorkloads struct for ApiWorkloads
type ApiWorkloads struct {
	// the number of unhealthy workloads
	Data map[string]int32 `json:"data"`
	// unhealthy workloads
	Items map[string]interface{} `json:"items,omitempty"`
	// the name of the namespace
	Namespace string `json:"namespace"`
}

type _ApiWorkloads ApiWorkloads

// NewApiWorkloads instantiates a new ApiWorkloads object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWorkloads(data map[string]int32, namespace string) *ApiWorkloads {
	this := ApiWorkloads{}
	this.Data = data
	this.Namespace = namespace
	return &this
}

// NewApiWorkloadsWithDefaults instantiates a new ApiWorkloads object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWorkloadsWithDefaults() *ApiWorkloads {
	this := ApiWorkloads{}
	return &this
}

// GetData returns the Data field value
func (o *ApiWorkloads) GetData() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ApiWorkloads) GetDataOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ApiWorkloads) SetData(v map[string]int32) {
	o.Data = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ApiWorkloads) GetItems() map[string]interface{} {
	if o == nil || IsNil(o.Items) {
		var ret map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWorkloads) GetItemsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return map[string]interface{}{}, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ApiWorkloads) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given map[string]interface{} and assigns it to the Items field.
func (o *ApiWorkloads) SetItems(v map[string]interface{}) {
	o.Items = v
}

// GetNamespace returns the Namespace field value
func (o *ApiWorkloads) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ApiWorkloads) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ApiWorkloads) SetNamespace(v string) {
	o.Namespace = v
}

func (o ApiWorkloads) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiWorkloads) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

func (o *ApiWorkloads) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"namespace",
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

	varApiWorkloads := _ApiWorkloads{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiWorkloads)

	if err != nil {
		return err
	}

	*o = ApiWorkloads(varApiWorkloads)

	return err
}

type NullableApiWorkloads struct {
	value *ApiWorkloads
	isSet bool
}

func (v NullableApiWorkloads) Get() *ApiWorkloads {
	return v.value
}

func (v *NullableApiWorkloads) Set(val *ApiWorkloads) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWorkloads) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWorkloads) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWorkloads(val *ApiWorkloads) *NullableApiWorkloads {
	return &NullableApiWorkloads{value: val, isSet: true}
}

func (v NullableApiWorkloads) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWorkloads) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


