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

// checks if the V1VolumeProjection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1VolumeProjection{}

// V1VolumeProjection Projection that may be projected along with other supported volume types
type V1VolumeProjection struct {
	ClusterTrustBundle *V1ClusterTrustBundleProjection `json:"clusterTrustBundle,omitempty"`
	ConfigMap *V1ConfigMapProjection `json:"configMap,omitempty"`
	DownwardAPI *V1DownwardAPIProjection `json:"downwardAPI,omitempty"`
	Secret *V1SecretProjection `json:"secret,omitempty"`
	ServiceAccountToken *V1ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty"`
}

// NewV1VolumeProjection instantiates a new V1VolumeProjection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1VolumeProjection() *V1VolumeProjection {
	this := V1VolumeProjection{}
	return &this
}

// NewV1VolumeProjectionWithDefaults instantiates a new V1VolumeProjection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1VolumeProjectionWithDefaults() *V1VolumeProjection {
	this := V1VolumeProjection{}
	return &this
}

// GetClusterTrustBundle returns the ClusterTrustBundle field value if set, zero value otherwise.
func (o *V1VolumeProjection) GetClusterTrustBundle() V1ClusterTrustBundleProjection {
	if o == nil || IsNil(o.ClusterTrustBundle) {
		var ret V1ClusterTrustBundleProjection
		return ret
	}
	return *o.ClusterTrustBundle
}

// GetClusterTrustBundleOk returns a tuple with the ClusterTrustBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VolumeProjection) GetClusterTrustBundleOk() (*V1ClusterTrustBundleProjection, bool) {
	if o == nil || IsNil(o.ClusterTrustBundle) {
		return nil, false
	}
	return o.ClusterTrustBundle, true
}

// HasClusterTrustBundle returns a boolean if a field has been set.
func (o *V1VolumeProjection) HasClusterTrustBundle() bool {
	if o != nil && !IsNil(o.ClusterTrustBundle) {
		return true
	}

	return false
}

// SetClusterTrustBundle gets a reference to the given V1ClusterTrustBundleProjection and assigns it to the ClusterTrustBundle field.
func (o *V1VolumeProjection) SetClusterTrustBundle(v V1ClusterTrustBundleProjection) {
	o.ClusterTrustBundle = &v
}

// GetConfigMap returns the ConfigMap field value if set, zero value otherwise.
func (o *V1VolumeProjection) GetConfigMap() V1ConfigMapProjection {
	if o == nil || IsNil(o.ConfigMap) {
		var ret V1ConfigMapProjection
		return ret
	}
	return *o.ConfigMap
}

// GetConfigMapOk returns a tuple with the ConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VolumeProjection) GetConfigMapOk() (*V1ConfigMapProjection, bool) {
	if o == nil || IsNil(o.ConfigMap) {
		return nil, false
	}
	return o.ConfigMap, true
}

// HasConfigMap returns a boolean if a field has been set.
func (o *V1VolumeProjection) HasConfigMap() bool {
	if o != nil && !IsNil(o.ConfigMap) {
		return true
	}

	return false
}

// SetConfigMap gets a reference to the given V1ConfigMapProjection and assigns it to the ConfigMap field.
func (o *V1VolumeProjection) SetConfigMap(v V1ConfigMapProjection) {
	o.ConfigMap = &v
}

// GetDownwardAPI returns the DownwardAPI field value if set, zero value otherwise.
func (o *V1VolumeProjection) GetDownwardAPI() V1DownwardAPIProjection {
	if o == nil || IsNil(o.DownwardAPI) {
		var ret V1DownwardAPIProjection
		return ret
	}
	return *o.DownwardAPI
}

// GetDownwardAPIOk returns a tuple with the DownwardAPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VolumeProjection) GetDownwardAPIOk() (*V1DownwardAPIProjection, bool) {
	if o == nil || IsNil(o.DownwardAPI) {
		return nil, false
	}
	return o.DownwardAPI, true
}

// HasDownwardAPI returns a boolean if a field has been set.
func (o *V1VolumeProjection) HasDownwardAPI() bool {
	if o != nil && !IsNil(o.DownwardAPI) {
		return true
	}

	return false
}

// SetDownwardAPI gets a reference to the given V1DownwardAPIProjection and assigns it to the DownwardAPI field.
func (o *V1VolumeProjection) SetDownwardAPI(v V1DownwardAPIProjection) {
	o.DownwardAPI = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *V1VolumeProjection) GetSecret() V1SecretProjection {
	if o == nil || IsNil(o.Secret) {
		var ret V1SecretProjection
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VolumeProjection) GetSecretOk() (*V1SecretProjection, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *V1VolumeProjection) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given V1SecretProjection and assigns it to the Secret field.
func (o *V1VolumeProjection) SetSecret(v V1SecretProjection) {
	o.Secret = &v
}

// GetServiceAccountToken returns the ServiceAccountToken field value if set, zero value otherwise.
func (o *V1VolumeProjection) GetServiceAccountToken() V1ServiceAccountTokenProjection {
	if o == nil || IsNil(o.ServiceAccountToken) {
		var ret V1ServiceAccountTokenProjection
		return ret
	}
	return *o.ServiceAccountToken
}

// GetServiceAccountTokenOk returns a tuple with the ServiceAccountToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VolumeProjection) GetServiceAccountTokenOk() (*V1ServiceAccountTokenProjection, bool) {
	if o == nil || IsNil(o.ServiceAccountToken) {
		return nil, false
	}
	return o.ServiceAccountToken, true
}

// HasServiceAccountToken returns a boolean if a field has been set.
func (o *V1VolumeProjection) HasServiceAccountToken() bool {
	if o != nil && !IsNil(o.ServiceAccountToken) {
		return true
	}

	return false
}

// SetServiceAccountToken gets a reference to the given V1ServiceAccountTokenProjection and assigns it to the ServiceAccountToken field.
func (o *V1VolumeProjection) SetServiceAccountToken(v V1ServiceAccountTokenProjection) {
	o.ServiceAccountToken = &v
}

func (o V1VolumeProjection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1VolumeProjection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterTrustBundle) {
		toSerialize["clusterTrustBundle"] = o.ClusterTrustBundle
	}
	if !IsNil(o.ConfigMap) {
		toSerialize["configMap"] = o.ConfigMap
	}
	if !IsNil(o.DownwardAPI) {
		toSerialize["downwardAPI"] = o.DownwardAPI
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.ServiceAccountToken) {
		toSerialize["serviceAccountToken"] = o.ServiceAccountToken
	}
	return toSerialize, nil
}

type NullableV1VolumeProjection struct {
	value *V1VolumeProjection
	isSet bool
}

func (v NullableV1VolumeProjection) Get() *V1VolumeProjection {
	return v.value
}

func (v *NullableV1VolumeProjection) Set(val *V1VolumeProjection) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VolumeProjection) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VolumeProjection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VolumeProjection(val *V1VolumeProjection) *NullableV1VolumeProjection {
	return &NullableV1VolumeProjection{value: val, isSet: true}
}

func (v NullableV1VolumeProjection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VolumeProjection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


