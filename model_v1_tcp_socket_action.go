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

// checks if the V1TCPSocketAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1TCPSocketAction{}

// V1TCPSocketAction TCPSocketAction describes an action based on opening a socket
type V1TCPSocketAction struct {
	// Optional: Host name to connect to, defaults to the pod IP.
	Host *string `json:"host,omitempty"`
	// Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	Port string `json:"port"`
}

type _V1TCPSocketAction V1TCPSocketAction

// NewV1TCPSocketAction instantiates a new V1TCPSocketAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TCPSocketAction(port string) *V1TCPSocketAction {
	this := V1TCPSocketAction{}
	this.Port = port
	return &this
}

// NewV1TCPSocketActionWithDefaults instantiates a new V1TCPSocketAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TCPSocketActionWithDefaults() *V1TCPSocketAction {
	this := V1TCPSocketAction{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *V1TCPSocketAction) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TCPSocketAction) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *V1TCPSocketAction) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *V1TCPSocketAction) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value
func (o *V1TCPSocketAction) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *V1TCPSocketAction) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *V1TCPSocketAction) SetPort(v string) {
	o.Port = v
}

func (o V1TCPSocketAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1TCPSocketAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

func (o *V1TCPSocketAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
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

	varV1TCPSocketAction := _V1TCPSocketAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1TCPSocketAction)

	if err != nil {
		return err
	}

	*o = V1TCPSocketAction(varV1TCPSocketAction)

	return err
}

type NullableV1TCPSocketAction struct {
	value *V1TCPSocketAction
	isSet bool
}

func (v NullableV1TCPSocketAction) Get() *V1TCPSocketAction {
	return v.value
}

func (v *NullableV1TCPSocketAction) Set(val *V1TCPSocketAction) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TCPSocketAction) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TCPSocketAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TCPSocketAction(val *V1TCPSocketAction) *NullableV1TCPSocketAction {
	return &NullableV1TCPSocketAction{value: val, isSet: true}
}

func (v NullableV1TCPSocketAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TCPSocketAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


