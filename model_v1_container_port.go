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

// checks if the V1ContainerPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ContainerPort{}

// V1ContainerPort ContainerPort represents a network port in a single container.
type V1ContainerPort struct {
	// Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort"`
	// What host IP to bind the external port to.
	HostIP *string `json:"hostIP,omitempty"`
	// Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
	HostPort *int32 `json:"hostPort,omitempty"`
	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
	Name *string `json:"name,omitempty"`
	// Protocol for port. Must be UDP, TCP, or SCTP. Defaults to \"TCP\".
	Protocol *string `json:"protocol,omitempty"`
}

type _V1ContainerPort V1ContainerPort

// NewV1ContainerPort instantiates a new V1ContainerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ContainerPort(containerPort int32) *V1ContainerPort {
	this := V1ContainerPort{}
	this.ContainerPort = containerPort
	return &this
}

// NewV1ContainerPortWithDefaults instantiates a new V1ContainerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ContainerPortWithDefaults() *V1ContainerPort {
	this := V1ContainerPort{}
	return &this
}

// GetContainerPort returns the ContainerPort field value
func (o *V1ContainerPort) GetContainerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContainerPort
}

// GetContainerPortOk returns a tuple with the ContainerPort field value
// and a boolean to check if the value has been set.
func (o *V1ContainerPort) GetContainerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerPort, true
}

// SetContainerPort sets field value
func (o *V1ContainerPort) SetContainerPort(v int32) {
	o.ContainerPort = v
}

// GetHostIP returns the HostIP field value if set, zero value otherwise.
func (o *V1ContainerPort) GetHostIP() string {
	if o == nil || IsNil(o.HostIP) {
		var ret string
		return ret
	}
	return *o.HostIP
}

// GetHostIPOk returns a tuple with the HostIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ContainerPort) GetHostIPOk() (*string, bool) {
	if o == nil || IsNil(o.HostIP) {
		return nil, false
	}
	return o.HostIP, true
}

// HasHostIP returns a boolean if a field has been set.
func (o *V1ContainerPort) HasHostIP() bool {
	if o != nil && !IsNil(o.HostIP) {
		return true
	}

	return false
}

// SetHostIP gets a reference to the given string and assigns it to the HostIP field.
func (o *V1ContainerPort) SetHostIP(v string) {
	o.HostIP = &v
}

// GetHostPort returns the HostPort field value if set, zero value otherwise.
func (o *V1ContainerPort) GetHostPort() int32 {
	if o == nil || IsNil(o.HostPort) {
		var ret int32
		return ret
	}
	return *o.HostPort
}

// GetHostPortOk returns a tuple with the HostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ContainerPort) GetHostPortOk() (*int32, bool) {
	if o == nil || IsNil(o.HostPort) {
		return nil, false
	}
	return o.HostPort, true
}

// HasHostPort returns a boolean if a field has been set.
func (o *V1ContainerPort) HasHostPort() bool {
	if o != nil && !IsNil(o.HostPort) {
		return true
	}

	return false
}

// SetHostPort gets a reference to the given int32 and assigns it to the HostPort field.
func (o *V1ContainerPort) SetHostPort(v int32) {
	o.HostPort = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ContainerPort) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ContainerPort) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ContainerPort) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ContainerPort) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V1ContainerPort) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ContainerPort) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V1ContainerPort) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *V1ContainerPort) SetProtocol(v string) {
	o.Protocol = &v
}

func (o V1ContainerPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ContainerPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerPort"] = o.ContainerPort
	if !IsNil(o.HostIP) {
		toSerialize["hostIP"] = o.HostIP
	}
	if !IsNil(o.HostPort) {
		toSerialize["hostPort"] = o.HostPort
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	return toSerialize, nil
}

func (o *V1ContainerPort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"containerPort",
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

	varV1ContainerPort := _V1ContainerPort{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ContainerPort)

	if err != nil {
		return err
	}

	*o = V1ContainerPort(varV1ContainerPort)

	return err
}

type NullableV1ContainerPort struct {
	value *V1ContainerPort
	isSet bool
}

func (v NullableV1ContainerPort) Get() *V1ContainerPort {
	return v.value
}

func (v *NullableV1ContainerPort) Set(val *V1ContainerPort) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ContainerPort) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ContainerPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ContainerPort(val *V1ContainerPort) *NullableV1ContainerPort {
	return &NullableV1ContainerPort{value: val, isSet: true}
}

func (v NullableV1ContainerPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ContainerPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


