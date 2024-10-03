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

// checks if the RegistriesImageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistriesImageDetails{}

// RegistriesImageDetails struct for RegistriesImageDetails
type RegistriesImageDetails struct {
	ImageBlob *RegistriesImageBlob `json:"imageBlob,omitempty"`
	ImageManifest *RegistriesImageManifest `json:"imageManifest,omitempty"`
	// image tag.
	ImageTag *string `json:"imageTag,omitempty"`
	// Status message.
	Message *string `json:"message,omitempty"`
	// registry domain.
	Registry *string `json:"registry,omitempty"`
	// Status is the status of the image search, such as \"succeeded\".
	Status *string `json:"status,omitempty"`
}

// NewRegistriesImageDetails instantiates a new RegistriesImageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistriesImageDetails() *RegistriesImageDetails {
	this := RegistriesImageDetails{}
	return &this
}

// NewRegistriesImageDetailsWithDefaults instantiates a new RegistriesImageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistriesImageDetailsWithDefaults() *RegistriesImageDetails {
	this := RegistriesImageDetails{}
	return &this
}

// GetImageBlob returns the ImageBlob field value if set, zero value otherwise.
func (o *RegistriesImageDetails) GetImageBlob() RegistriesImageBlob {
	if o == nil || IsNil(o.ImageBlob) {
		var ret RegistriesImageBlob
		return ret
	}
	return *o.ImageBlob
}

// GetImageBlobOk returns a tuple with the ImageBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageDetails) GetImageBlobOk() (*RegistriesImageBlob, bool) {
	if o == nil || IsNil(o.ImageBlob) {
		return nil, false
	}
	return o.ImageBlob, true
}

// HasImageBlob returns a boolean if a field has been set.
func (o *RegistriesImageDetails) HasImageBlob() bool {
	if o != nil && !IsNil(o.ImageBlob) {
		return true
	}

	return false
}

// SetImageBlob gets a reference to the given RegistriesImageBlob and assigns it to the ImageBlob field.
func (o *RegistriesImageDetails) SetImageBlob(v RegistriesImageBlob) {
	o.ImageBlob = &v
}

// GetImageManifest returns the ImageManifest field value if set, zero value otherwise.
func (o *RegistriesImageDetails) GetImageManifest() RegistriesImageManifest {
	if o == nil || IsNil(o.ImageManifest) {
		var ret RegistriesImageManifest
		return ret
	}
	return *o.ImageManifest
}

// GetImageManifestOk returns a tuple with the ImageManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageDetails) GetImageManifestOk() (*RegistriesImageManifest, bool) {
	if o == nil || IsNil(o.ImageManifest) {
		return nil, false
	}
	return o.ImageManifest, true
}

// HasImageManifest returns a boolean if a field has been set.
func (o *RegistriesImageDetails) HasImageManifest() bool {
	if o != nil && !IsNil(o.ImageManifest) {
		return true
	}

	return false
}

// SetImageManifest gets a reference to the given RegistriesImageManifest and assigns it to the ImageManifest field.
func (o *RegistriesImageDetails) SetImageManifest(v RegistriesImageManifest) {
	o.ImageManifest = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *RegistriesImageDetails) GetImageTag() string {
	if o == nil || IsNil(o.ImageTag) {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageDetails) GetImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.ImageTag) {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *RegistriesImageDetails) HasImageTag() bool {
	if o != nil && !IsNil(o.ImageTag) {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *RegistriesImageDetails) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RegistriesImageDetails) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageDetails) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RegistriesImageDetails) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RegistriesImageDetails) SetMessage(v string) {
	o.Message = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *RegistriesImageDetails) GetRegistry() string {
	if o == nil || IsNil(o.Registry) {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageDetails) GetRegistryOk() (*string, bool) {
	if o == nil || IsNil(o.Registry) {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *RegistriesImageDetails) HasRegistry() bool {
	if o != nil && !IsNil(o.Registry) {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *RegistriesImageDetails) SetRegistry(v string) {
	o.Registry = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegistriesImageDetails) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageDetails) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegistriesImageDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RegistriesImageDetails) SetStatus(v string) {
	o.Status = &v
}

func (o RegistriesImageDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistriesImageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageBlob) {
		toSerialize["imageBlob"] = o.ImageBlob
	}
	if !IsNil(o.ImageManifest) {
		toSerialize["imageManifest"] = o.ImageManifest
	}
	if !IsNil(o.ImageTag) {
		toSerialize["imageTag"] = o.ImageTag
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Registry) {
		toSerialize["registry"] = o.Registry
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableRegistriesImageDetails struct {
	value *RegistriesImageDetails
	isSet bool
}

func (v NullableRegistriesImageDetails) Get() *RegistriesImageDetails {
	return v.value
}

func (v *NullableRegistriesImageDetails) Set(val *RegistriesImageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistriesImageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistriesImageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistriesImageDetails(val *RegistriesImageDetails) *NullableRegistriesImageDetails {
	return &NullableRegistriesImageDetails{value: val, isSet: true}
}

func (v NullableRegistriesImageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistriesImageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

