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

// checks if the V1beta1IngressTLS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1IngressTLS{}

// V1beta1IngressTLS IngressTLS describes the transport layer security associated with an Ingress.
type V1beta1IngressTLS struct {
	// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Hosts []string `json:"hosts,omitempty"`
	// SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	SecretName *string `json:"secretName,omitempty"`
}

// NewV1beta1IngressTLS instantiates a new V1beta1IngressTLS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1IngressTLS() *V1beta1IngressTLS {
	this := V1beta1IngressTLS{}
	return &this
}

// NewV1beta1IngressTLSWithDefaults instantiates a new V1beta1IngressTLS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1IngressTLSWithDefaults() *V1beta1IngressTLS {
	this := V1beta1IngressTLS{}
	return &this
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *V1beta1IngressTLS) GetHosts() []string {
	if o == nil || IsNil(o.Hosts) {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1IngressTLS) GetHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *V1beta1IngressTLS) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *V1beta1IngressTLS) SetHosts(v []string) {
	o.Hosts = v
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *V1beta1IngressTLS) GetSecretName() string {
	if o == nil || IsNil(o.SecretName) {
		var ret string
		return ret
	}
	return *o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1IngressTLS) GetSecretNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecretName) {
		return nil, false
	}
	return o.SecretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *V1beta1IngressTLS) HasSecretName() bool {
	if o != nil && !IsNil(o.SecretName) {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the SecretName field.
func (o *V1beta1IngressTLS) SetSecretName(v string) {
	o.SecretName = &v
}

func (o V1beta1IngressTLS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1IngressTLS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.SecretName) {
		toSerialize["secretName"] = o.SecretName
	}
	return toSerialize, nil
}

type NullableV1beta1IngressTLS struct {
	value *V1beta1IngressTLS
	isSet bool
}

func (v NullableV1beta1IngressTLS) Get() *V1beta1IngressTLS {
	return v.value
}

func (v *NullableV1beta1IngressTLS) Set(val *V1beta1IngressTLS) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1IngressTLS) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1IngressTLS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1IngressTLS(val *V1beta1IngressTLS) *NullableV1beta1IngressTLS {
	return &NullableV1beta1IngressTLS{value: val, isSet: true}
}

func (v NullableV1beta1IngressTLS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1IngressTLS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

