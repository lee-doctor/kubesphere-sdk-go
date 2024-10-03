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

// checks if the OpenpitrixCreateRepoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixCreateRepoRequest{}

// OpenpitrixCreateRepoRequest struct for OpenpitrixCreateRepoRequest
type OpenpitrixCreateRepoRequest struct {
	AppDefaultStatus *string `json:"app_default_status,omitempty"`
	CategoryId *string `json:"category_id,omitempty"`
	Credential *string `json:"credential,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Providers []string `json:"providers"`
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Workspace *string `json:"workspace,omitempty"`
}

type _OpenpitrixCreateRepoRequest OpenpitrixCreateRepoRequest

// NewOpenpitrixCreateRepoRequest instantiates a new OpenpitrixCreateRepoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixCreateRepoRequest(providers []string) *OpenpitrixCreateRepoRequest {
	this := OpenpitrixCreateRepoRequest{}
	this.Providers = providers
	return &this
}

// NewOpenpitrixCreateRepoRequestWithDefaults instantiates a new OpenpitrixCreateRepoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixCreateRepoRequestWithDefaults() *OpenpitrixCreateRepoRequest {
	this := OpenpitrixCreateRepoRequest{}
	return &this
}

// GetAppDefaultStatus returns the AppDefaultStatus field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetAppDefaultStatus() string {
	if o == nil || IsNil(o.AppDefaultStatus) {
		var ret string
		return ret
	}
	return *o.AppDefaultStatus
}

// GetAppDefaultStatusOk returns a tuple with the AppDefaultStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetAppDefaultStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AppDefaultStatus) {
		return nil, false
	}
	return o.AppDefaultStatus, true
}

// HasAppDefaultStatus returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasAppDefaultStatus() bool {
	if o != nil && !IsNil(o.AppDefaultStatus) {
		return true
	}

	return false
}

// SetAppDefaultStatus gets a reference to the given string and assigns it to the AppDefaultStatus field.
func (o *OpenpitrixCreateRepoRequest) SetAppDefaultStatus(v string) {
	o.AppDefaultStatus = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *OpenpitrixCreateRepoRequest) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetCredential() string {
	if o == nil || IsNil(o.Credential) {
		var ret string
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetCredentialOk() (*string, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given string and assigns it to the Credential field.
func (o *OpenpitrixCreateRepoRequest) SetCredential(v string) {
	o.Credential = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenpitrixCreateRepoRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenpitrixCreateRepoRequest) SetName(v string) {
	o.Name = &v
}

// GetProviders returns the Providers field value
func (o *OpenpitrixCreateRepoRequest) GetProviders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetProvidersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *OpenpitrixCreateRepoRequest) SetProviders(v []string) {
	o.Providers = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OpenpitrixCreateRepoRequest) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OpenpitrixCreateRepoRequest) SetUrl(v string) {
	o.Url = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OpenpitrixCreateRepoRequest) SetVisibility(v string) {
	o.Visibility = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *OpenpitrixCreateRepoRequest) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace) {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixCreateRepoRequest) GetWorkspaceOk() (*string, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *OpenpitrixCreateRepoRequest) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *OpenpitrixCreateRepoRequest) SetWorkspace(v string) {
	o.Workspace = &v
}

func (o OpenpitrixCreateRepoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixCreateRepoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppDefaultStatus) {
		toSerialize["app_default_status"] = o.AppDefaultStatus
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["providers"] = o.Providers
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	return toSerialize, nil
}

func (o *OpenpitrixCreateRepoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"providers",
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

	varOpenpitrixCreateRepoRequest := _OpenpitrixCreateRepoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenpitrixCreateRepoRequest)

	if err != nil {
		return err
	}

	*o = OpenpitrixCreateRepoRequest(varOpenpitrixCreateRepoRequest)

	return err
}

type NullableOpenpitrixCreateRepoRequest struct {
	value *OpenpitrixCreateRepoRequest
	isSet bool
}

func (v NullableOpenpitrixCreateRepoRequest) Get() *OpenpitrixCreateRepoRequest {
	return v.value
}

func (v *NullableOpenpitrixCreateRepoRequest) Set(val *OpenpitrixCreateRepoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixCreateRepoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixCreateRepoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixCreateRepoRequest(val *OpenpitrixCreateRepoRequest) *NullableOpenpitrixCreateRepoRequest {
	return &NullableOpenpitrixCreateRepoRequest{value: val, isSet: true}
}

func (v NullableOpenpitrixCreateRepoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixCreateRepoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

