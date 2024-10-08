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
)

// checks if the V1alpha1Workspace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha1Workspace{}

// V1alpha1Workspace struct for V1alpha1Workspace
type V1alpha1Workspace struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec *V1alpha1WorkspaceSpec `json:"spec,omitempty"`
	Status map[string]interface{} `json:"status,omitempty"`
}

// NewV1alpha1Workspace instantiates a new V1alpha1Workspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha1Workspace() *V1alpha1Workspace {
	this := V1alpha1Workspace{}
	return &this
}

// NewV1alpha1WorkspaceWithDefaults instantiates a new V1alpha1Workspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha1WorkspaceWithDefaults() *V1alpha1Workspace {
	this := V1alpha1Workspace{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1alpha1Workspace) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Workspace) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1alpha1Workspace) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1alpha1Workspace) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1alpha1Workspace) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Workspace) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1alpha1Workspace) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1alpha1Workspace) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1alpha1Workspace) GetMetadata() V1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Workspace) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1alpha1Workspace) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1alpha1Workspace) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *V1alpha1Workspace) GetSpec() V1alpha1WorkspaceSpec {
	if o == nil || IsNil(o.Spec) {
		var ret V1alpha1WorkspaceSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Workspace) GetSpecOk() (*V1alpha1WorkspaceSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *V1alpha1Workspace) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given V1alpha1WorkspaceSpec and assigns it to the Spec field.
func (o *V1alpha1Workspace) SetSpec(v V1alpha1WorkspaceSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1alpha1Workspace) GetStatus() map[string]interface{} {
	if o == nil || IsNil(o.Status) {
		var ret map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha1Workspace) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1alpha1Workspace) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *V1alpha1Workspace) SetStatus(v map[string]interface{}) {
	o.Status = v
}

func (o V1alpha1Workspace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha1Workspace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV1alpha1Workspace struct {
	value *V1alpha1Workspace
	isSet bool
}

func (v NullableV1alpha1Workspace) Get() *V1alpha1Workspace {
	return v.value
}

func (v *NullableV1alpha1Workspace) Set(val *V1alpha1Workspace) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha1Workspace) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha1Workspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha1Workspace(val *V1alpha1Workspace) *NullableV1alpha1Workspace {
	return &NullableV1alpha1Workspace{value: val, isSet: true}
}

func (v NullableV1alpha1Workspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha1Workspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


