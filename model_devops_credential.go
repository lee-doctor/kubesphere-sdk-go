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

// checks if the DevopsCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevopsCredential{}

// DevopsCredential struct for DevopsCredential
type DevopsCredential struct {
	// Credential's description'
	Description *string `json:"description,omitempty"`
	// Credential's display name
	DisplayName *string `json:"display_name,omitempty"`
	// Credential's domain,In ks we only use the default domain, default '_''
	Domain *string `json:"domain,omitempty"`
	Fingerprint *DevopsCredentialFingerprint `json:"fingerprint,omitempty"`
	// Id of Credential, e.g. dockerhub-id
	Id string `json:"id"`
	// Type of Credential, e.g. ssh/kubeconfig
	Type string `json:"type"`
}

type _DevopsCredential DevopsCredential

// NewDevopsCredential instantiates a new DevopsCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevopsCredential(id string, type_ string) *DevopsCredential {
	this := DevopsCredential{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewDevopsCredentialWithDefaults instantiates a new DevopsCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevopsCredentialWithDefaults() *DevopsCredential {
	this := DevopsCredential{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DevopsCredential) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCredential) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DevopsCredential) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DevopsCredential) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DevopsCredential) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCredential) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DevopsCredential) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DevopsCredential) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DevopsCredential) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCredential) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DevopsCredential) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DevopsCredential) SetDomain(v string) {
	o.Domain = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *DevopsCredential) GetFingerprint() DevopsCredentialFingerprint {
	if o == nil || IsNil(o.Fingerprint) {
		var ret DevopsCredentialFingerprint
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevopsCredential) GetFingerprintOk() (*DevopsCredentialFingerprint, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *DevopsCredential) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given DevopsCredentialFingerprint and assigns it to the Fingerprint field.
func (o *DevopsCredential) SetFingerprint(v DevopsCredentialFingerprint) {
	o.Fingerprint = &v
}

// GetId returns the Id field value
func (o *DevopsCredential) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DevopsCredential) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DevopsCredential) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *DevopsCredential) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DevopsCredential) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DevopsCredential) SetType(v string) {
	o.Type = v
}

func (o DevopsCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevopsCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *DevopsCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
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

	varDevopsCredential := _DevopsCredential{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDevopsCredential)

	if err != nil {
		return err
	}

	*o = DevopsCredential(varDevopsCredential)

	return err
}

type NullableDevopsCredential struct {
	value *DevopsCredential
	isSet bool
}

func (v NullableDevopsCredential) Get() *DevopsCredential {
	return v.value
}

func (v *NullableDevopsCredential) Set(val *DevopsCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableDevopsCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableDevopsCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevopsCredential(val *DevopsCredential) *NullableDevopsCredential {
	return &NullableDevopsCredential{value: val, isSet: true}
}

func (v NullableDevopsCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevopsCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


