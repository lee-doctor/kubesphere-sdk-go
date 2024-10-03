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

// checks if the V1RoleBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RoleBinding{}

// V1RoleBinding RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.
type V1RoleBinding struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	RoleRef V1RoleRef `json:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []V1Subject `json:"subjects,omitempty"`
}

type _V1RoleBinding V1RoleBinding

// NewV1RoleBinding instantiates a new V1RoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RoleBinding(roleRef V1RoleRef) *V1RoleBinding {
	this := V1RoleBinding{}
	this.RoleRef = roleRef
	return &this
}

// NewV1RoleBindingWithDefaults instantiates a new V1RoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RoleBindingWithDefaults() *V1RoleBinding {
	this := V1RoleBinding{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1RoleBinding) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RoleBinding) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1RoleBinding) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1RoleBinding) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1RoleBinding) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RoleBinding) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1RoleBinding) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1RoleBinding) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1RoleBinding) GetMetadata() V1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RoleBinding) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1RoleBinding) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1RoleBinding) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetRoleRef returns the RoleRef field value
func (o *V1RoleBinding) GetRoleRef() V1RoleRef {
	if o == nil {
		var ret V1RoleRef
		return ret
	}

	return o.RoleRef
}

// GetRoleRefOk returns a tuple with the RoleRef field value
// and a boolean to check if the value has been set.
func (o *V1RoleBinding) GetRoleRefOk() (*V1RoleRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleRef, true
}

// SetRoleRef sets field value
func (o *V1RoleBinding) SetRoleRef(v V1RoleRef) {
	o.RoleRef = v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *V1RoleBinding) GetSubjects() []V1Subject {
	if o == nil || IsNil(o.Subjects) {
		var ret []V1Subject
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RoleBinding) GetSubjectsOk() ([]V1Subject, bool) {
	if o == nil || IsNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *V1RoleBinding) HasSubjects() bool {
	if o != nil && !IsNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []V1Subject and assigns it to the Subjects field.
func (o *V1RoleBinding) SetSubjects(v []V1Subject) {
	o.Subjects = v
}

func (o V1RoleBinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RoleBinding) ToMap() (map[string]interface{}, error) {
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
	toSerialize["roleRef"] = o.RoleRef
	if !IsNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	return toSerialize, nil
}

func (o *V1RoleBinding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roleRef",
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

	varV1RoleBinding := _V1RoleBinding{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1RoleBinding)

	if err != nil {
		return err
	}

	*o = V1RoleBinding(varV1RoleBinding)

	return err
}

type NullableV1RoleBinding struct {
	value *V1RoleBinding
	isSet bool
}

func (v NullableV1RoleBinding) Get() *V1RoleBinding {
	return v.value
}

func (v *NullableV1RoleBinding) Set(val *V1RoleBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RoleBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RoleBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RoleBinding(val *V1RoleBinding) *NullableV1RoleBinding {
	return &NullableV1RoleBinding{value: val, isSet: true}
}

func (v NullableV1RoleBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RoleBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


