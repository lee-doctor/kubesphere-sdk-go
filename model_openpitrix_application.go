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

// checks if the OpenpitrixApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixApplication{}

// OpenpitrixApplication struct for OpenpitrixApplication
type OpenpitrixApplication struct {
	App *OpenpitrixApp `json:"app,omitempty"`
	Cluster *OpenpitrixCluster `json:"cluster,omitempty"`
	// application ingresses
	Ingresses []V1beta1Ingress `json:"ingresses,omitempty"`
	// application name
	Name string `json:"name"`
	// application services
	Services []V1Service `json:"services,omitempty"`
	Version *OpenpitrixAppVersion `json:"version,omitempty"`
	Workloads *OpenpitrixWorkLoads `json:"workloads,omitempty"`
}

type _OpenpitrixApplication OpenpitrixApplication

// NewOpenpitrixApplication instantiates a new OpenpitrixApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixApplication(name string) *OpenpitrixApplication {
	this := OpenpitrixApplication{}
	this.Name = name
	return &this
}

// NewOpenpitrixApplicationWithDefaults instantiates a new OpenpitrixApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixApplicationWithDefaults() *OpenpitrixApplication {
	this := OpenpitrixApplication{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *OpenpitrixApplication) GetApp() OpenpitrixApp {
	if o == nil || IsNil(o.App) {
		var ret OpenpitrixApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixApplication) GetAppOk() (*OpenpitrixApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *OpenpitrixApplication) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given OpenpitrixApp and assigns it to the App field.
func (o *OpenpitrixApplication) SetApp(v OpenpitrixApp) {
	o.App = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *OpenpitrixApplication) GetCluster() OpenpitrixCluster {
	if o == nil || IsNil(o.Cluster) {
		var ret OpenpitrixCluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixApplication) GetClusterOk() (*OpenpitrixCluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *OpenpitrixApplication) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given OpenpitrixCluster and assigns it to the Cluster field.
func (o *OpenpitrixApplication) SetCluster(v OpenpitrixCluster) {
	o.Cluster = &v
}

// GetIngresses returns the Ingresses field value if set, zero value otherwise.
func (o *OpenpitrixApplication) GetIngresses() []V1beta1Ingress {
	if o == nil || IsNil(o.Ingresses) {
		var ret []V1beta1Ingress
		return ret
	}
	return o.Ingresses
}

// GetIngressesOk returns a tuple with the Ingresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixApplication) GetIngressesOk() ([]V1beta1Ingress, bool) {
	if o == nil || IsNil(o.Ingresses) {
		return nil, false
	}
	return o.Ingresses, true
}

// HasIngresses returns a boolean if a field has been set.
func (o *OpenpitrixApplication) HasIngresses() bool {
	if o != nil && !IsNil(o.Ingresses) {
		return true
	}

	return false
}

// SetIngresses gets a reference to the given []V1beta1Ingress and assigns it to the Ingresses field.
func (o *OpenpitrixApplication) SetIngresses(v []V1beta1Ingress) {
	o.Ingresses = v
}

// GetName returns the Name field value
func (o *OpenpitrixApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenpitrixApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenpitrixApplication) SetName(v string) {
	o.Name = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *OpenpitrixApplication) GetServices() []V1Service {
	if o == nil || IsNil(o.Services) {
		var ret []V1Service
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixApplication) GetServicesOk() ([]V1Service, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *OpenpitrixApplication) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []V1Service and assigns it to the Services field.
func (o *OpenpitrixApplication) SetServices(v []V1Service) {
	o.Services = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OpenpitrixApplication) GetVersion() OpenpitrixAppVersion {
	if o == nil || IsNil(o.Version) {
		var ret OpenpitrixAppVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixApplication) GetVersionOk() (*OpenpitrixAppVersion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OpenpitrixApplication) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given OpenpitrixAppVersion and assigns it to the Version field.
func (o *OpenpitrixApplication) SetVersion(v OpenpitrixAppVersion) {
	o.Version = &v
}

// GetWorkloads returns the Workloads field value if set, zero value otherwise.
func (o *OpenpitrixApplication) GetWorkloads() OpenpitrixWorkLoads {
	if o == nil || IsNil(o.Workloads) {
		var ret OpenpitrixWorkLoads
		return ret
	}
	return *o.Workloads
}

// GetWorkloadsOk returns a tuple with the Workloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixApplication) GetWorkloadsOk() (*OpenpitrixWorkLoads, bool) {
	if o == nil || IsNil(o.Workloads) {
		return nil, false
	}
	return o.Workloads, true
}

// HasWorkloads returns a boolean if a field has been set.
func (o *OpenpitrixApplication) HasWorkloads() bool {
	if o != nil && !IsNil(o.Workloads) {
		return true
	}

	return false
}

// SetWorkloads gets a reference to the given OpenpitrixWorkLoads and assigns it to the Workloads field.
func (o *OpenpitrixApplication) SetWorkloads(v OpenpitrixWorkLoads) {
	o.Workloads = &v
}

func (o OpenpitrixApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Ingresses) {
		toSerialize["ingresses"] = o.Ingresses
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Workloads) {
		toSerialize["workloads"] = o.Workloads
	}
	return toSerialize, nil
}

func (o *OpenpitrixApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varOpenpitrixApplication := _OpenpitrixApplication{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenpitrixApplication)

	if err != nil {
		return err
	}

	*o = OpenpitrixApplication(varOpenpitrixApplication)

	return err
}

type NullableOpenpitrixApplication struct {
	value *OpenpitrixApplication
	isSet bool
}

func (v NullableOpenpitrixApplication) Get() *OpenpitrixApplication {
	return v.value
}

func (v *NullableOpenpitrixApplication) Set(val *OpenpitrixApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixApplication(val *OpenpitrixApplication) *NullableOpenpitrixApplication {
	return &NullableOpenpitrixApplication{value: val, isSet: true}
}

func (v NullableOpenpitrixApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

