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

// checks if the MonitoringMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringMetrics{}

// MonitoringMetrics struct for MonitoringMetrics
type MonitoringMetrics struct {
	// current page returned
	Page *int32 `json:"page,omitempty"`
	// actual array of results
	Results []MonitoringMetric `json:"results"`
	// page size
	TotalItem *int32 `json:"total_item,omitempty"`
	// total number of pages
	TotalPage *int32 `json:"total_page,omitempty"`
}

type _MonitoringMetrics MonitoringMetrics

// NewMonitoringMetrics instantiates a new MonitoringMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringMetrics(results []MonitoringMetric) *MonitoringMetrics {
	this := MonitoringMetrics{}
	this.Results = results
	return &this
}

// NewMonitoringMetricsWithDefaults instantiates a new MonitoringMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringMetricsWithDefaults() *MonitoringMetrics {
	this := MonitoringMetrics{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *MonitoringMetrics) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMetrics) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *MonitoringMetrics) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *MonitoringMetrics) SetPage(v int32) {
	o.Page = &v
}

// GetResults returns the Results field value
func (o *MonitoringMetrics) GetResults() []MonitoringMetric {
	if o == nil {
		var ret []MonitoringMetric
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *MonitoringMetrics) GetResultsOk() ([]MonitoringMetric, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *MonitoringMetrics) SetResults(v []MonitoringMetric) {
	o.Results = v
}

// GetTotalItem returns the TotalItem field value if set, zero value otherwise.
func (o *MonitoringMetrics) GetTotalItem() int32 {
	if o == nil || IsNil(o.TotalItem) {
		var ret int32
		return ret
	}
	return *o.TotalItem
}

// GetTotalItemOk returns a tuple with the TotalItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMetrics) GetTotalItemOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItem) {
		return nil, false
	}
	return o.TotalItem, true
}

// HasTotalItem returns a boolean if a field has been set.
func (o *MonitoringMetrics) HasTotalItem() bool {
	if o != nil && !IsNil(o.TotalItem) {
		return true
	}

	return false
}

// SetTotalItem gets a reference to the given int32 and assigns it to the TotalItem field.
func (o *MonitoringMetrics) SetTotalItem(v int32) {
	o.TotalItem = &v
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise.
func (o *MonitoringMetrics) GetTotalPage() int32 {
	if o == nil || IsNil(o.TotalPage) {
		var ret int32
		return ret
	}
	return *o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMetrics) GetTotalPageOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPage) {
		return nil, false
	}
	return o.TotalPage, true
}

// HasTotalPage returns a boolean if a field has been set.
func (o *MonitoringMetrics) HasTotalPage() bool {
	if o != nil && !IsNil(o.TotalPage) {
		return true
	}

	return false
}

// SetTotalPage gets a reference to the given int32 and assigns it to the TotalPage field.
func (o *MonitoringMetrics) SetTotalPage(v int32) {
	o.TotalPage = &v
}

func (o MonitoringMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	toSerialize["results"] = o.Results
	if !IsNil(o.TotalItem) {
		toSerialize["total_item"] = o.TotalItem
	}
	if !IsNil(o.TotalPage) {
		toSerialize["total_page"] = o.TotalPage
	}
	return toSerialize, nil
}

func (o *MonitoringMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varMonitoringMetrics := _MonitoringMetrics{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMonitoringMetrics)

	if err != nil {
		return err
	}

	*o = MonitoringMetrics(varMonitoringMetrics)

	return err
}

type NullableMonitoringMetrics struct {
	value *MonitoringMetrics
	isSet bool
}

func (v NullableMonitoringMetrics) Get() *MonitoringMetrics {
	return v.value
}

func (v *NullableMonitoringMetrics) Set(val *MonitoringMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringMetrics(val *MonitoringMetrics) *NullableMonitoringMetrics {
	return &NullableMonitoringMetrics{value: val, isSet: true}
}

func (v NullableMonitoringMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

