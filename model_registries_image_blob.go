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

// checks if the RegistriesImageBlob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistriesImageBlob{}

// RegistriesImageBlob struct for RegistriesImageBlob
type RegistriesImageBlob struct {
	// The architecture field specifies the CPU architecture, for example amd64 or ppc64le.
	Architecture *string `json:"architecture,omitempty"`
	Config *RegistriesConfig `json:"config,omitempty"`
	// Container id.
	Container *string `json:"container,omitempty"`
	ContainerConfig *RegistriesContainerConfig `json:"container_config,omitempty"`
	// Create time.
	Created *time.Time `json:"created,omitempty"`
	// docker version.
	DockerVersion *string `json:"docker_version,omitempty"`
	// The data of history update.
	History []RegistriesHistory `json:"history,omitempty"`
	// Operating system.
	Os *string `json:"os,omitempty"`
	RootfsOmitempty RegistriesRootfs `json:"rootfs omitempty"`
}

type _RegistriesImageBlob RegistriesImageBlob

// NewRegistriesImageBlob instantiates a new RegistriesImageBlob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistriesImageBlob(rootfsOmitempty RegistriesRootfs) *RegistriesImageBlob {
	this := RegistriesImageBlob{}
	this.RootfsOmitempty = rootfsOmitempty
	return &this
}

// NewRegistriesImageBlobWithDefaults instantiates a new RegistriesImageBlob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistriesImageBlobWithDefaults() *RegistriesImageBlob {
	this := RegistriesImageBlob{}
	return &this
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *RegistriesImageBlob) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetConfig() RegistriesConfig {
	if o == nil || IsNil(o.Config) {
		var ret RegistriesConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetConfigOk() (*RegistriesConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given RegistriesConfig and assigns it to the Config field.
func (o *RegistriesImageBlob) SetConfig(v RegistriesConfig) {
	o.Config = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *RegistriesImageBlob) SetContainer(v string) {
	o.Container = &v
}

// GetContainerConfig returns the ContainerConfig field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetContainerConfig() RegistriesContainerConfig {
	if o == nil || IsNil(o.ContainerConfig) {
		var ret RegistriesContainerConfig
		return ret
	}
	return *o.ContainerConfig
}

// GetContainerConfigOk returns a tuple with the ContainerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetContainerConfigOk() (*RegistriesContainerConfig, bool) {
	if o == nil || IsNil(o.ContainerConfig) {
		return nil, false
	}
	return o.ContainerConfig, true
}

// HasContainerConfig returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasContainerConfig() bool {
	if o != nil && !IsNil(o.ContainerConfig) {
		return true
	}

	return false
}

// SetContainerConfig gets a reference to the given RegistriesContainerConfig and assigns it to the ContainerConfig field.
func (o *RegistriesImageBlob) SetContainerConfig(v RegistriesContainerConfig) {
	o.ContainerConfig = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RegistriesImageBlob) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDockerVersion returns the DockerVersion field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetDockerVersion() string {
	if o == nil || IsNil(o.DockerVersion) {
		var ret string
		return ret
	}
	return *o.DockerVersion
}

// GetDockerVersionOk returns a tuple with the DockerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetDockerVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DockerVersion) {
		return nil, false
	}
	return o.DockerVersion, true
}

// HasDockerVersion returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasDockerVersion() bool {
	if o != nil && !IsNil(o.DockerVersion) {
		return true
	}

	return false
}

// SetDockerVersion gets a reference to the given string and assigns it to the DockerVersion field.
func (o *RegistriesImageBlob) SetDockerVersion(v string) {
	o.DockerVersion = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetHistory() []RegistriesHistory {
	if o == nil || IsNil(o.History) {
		var ret []RegistriesHistory
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetHistoryOk() ([]RegistriesHistory, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []RegistriesHistory and assigns it to the History field.
func (o *RegistriesImageBlob) SetHistory(v []RegistriesHistory) {
	o.History = v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *RegistriesImageBlob) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *RegistriesImageBlob) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *RegistriesImageBlob) SetOs(v string) {
	o.Os = &v
}

// GetRootfsOmitempty returns the RootfsOmitempty field value
func (o *RegistriesImageBlob) GetRootfsOmitempty() RegistriesRootfs {
	if o == nil {
		var ret RegistriesRootfs
		return ret
	}

	return o.RootfsOmitempty
}

// GetRootfsOmitemptyOk returns a tuple with the RootfsOmitempty field value
// and a boolean to check if the value has been set.
func (o *RegistriesImageBlob) GetRootfsOmitemptyOk() (*RegistriesRootfs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootfsOmitempty, true
}

// SetRootfsOmitempty sets field value
func (o *RegistriesImageBlob) SetRootfsOmitempty(v RegistriesRootfs) {
	o.RootfsOmitempty = v
}

func (o RegistriesImageBlob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistriesImageBlob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Architecture) {
		toSerialize["architecture"] = o.Architecture
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if !IsNil(o.ContainerConfig) {
		toSerialize["container_config"] = o.ContainerConfig
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DockerVersion) {
		toSerialize["docker_version"] = o.DockerVersion
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	toSerialize["rootfs omitempty"] = o.RootfsOmitempty
	return toSerialize, nil
}

func (o *RegistriesImageBlob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rootfs omitempty",
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

	varRegistriesImageBlob := _RegistriesImageBlob{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistriesImageBlob)

	if err != nil {
		return err
	}

	*o = RegistriesImageBlob(varRegistriesImageBlob)

	return err
}

type NullableRegistriesImageBlob struct {
	value *RegistriesImageBlob
	isSet bool
}

func (v NullableRegistriesImageBlob) Get() *RegistriesImageBlob {
	return v.value
}

func (v *NullableRegistriesImageBlob) Set(val *RegistriesImageBlob) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistriesImageBlob) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistriesImageBlob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistriesImageBlob(val *RegistriesImageBlob) *NullableRegistriesImageBlob {
	return &NullableRegistriesImageBlob{value: val, isSet: true}
}

func (v NullableRegistriesImageBlob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistriesImageBlob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


