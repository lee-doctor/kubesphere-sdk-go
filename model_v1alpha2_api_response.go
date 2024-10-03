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

// checks if the V1alpha2APIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1alpha2APIResponse{}

// V1alpha2APIResponse struct for V1alpha2APIResponse
type V1alpha2APIResponse struct {
	Histogram *LoggingHistogram `json:"histogram,omitempty"`
	Query *LoggingLogs `json:"query,omitempty"`
	Statistics *LoggingStatistics `json:"statistics,omitempty"`
}

// NewV1alpha2APIResponse instantiates a new V1alpha2APIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1alpha2APIResponse() *V1alpha2APIResponse {
	this := V1alpha2APIResponse{}
	return &this
}

// NewV1alpha2APIResponseWithDefaults instantiates a new V1alpha2APIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1alpha2APIResponseWithDefaults() *V1alpha2APIResponse {
	this := V1alpha2APIResponse{}
	return &this
}

// GetHistogram returns the Histogram field value if set, zero value otherwise.
func (o *V1alpha2APIResponse) GetHistogram() LoggingHistogram {
	if o == nil || IsNil(o.Histogram) {
		var ret LoggingHistogram
		return ret
	}
	return *o.Histogram
}

// GetHistogramOk returns a tuple with the Histogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2APIResponse) GetHistogramOk() (*LoggingHistogram, bool) {
	if o == nil || IsNil(o.Histogram) {
		return nil, false
	}
	return o.Histogram, true
}

// HasHistogram returns a boolean if a field has been set.
func (o *V1alpha2APIResponse) HasHistogram() bool {
	if o != nil && !IsNil(o.Histogram) {
		return true
	}

	return false
}

// SetHistogram gets a reference to the given LoggingHistogram and assigns it to the Histogram field.
func (o *V1alpha2APIResponse) SetHistogram(v LoggingHistogram) {
	o.Histogram = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *V1alpha2APIResponse) GetQuery() LoggingLogs {
	if o == nil || IsNil(o.Query) {
		var ret LoggingLogs
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2APIResponse) GetQueryOk() (*LoggingLogs, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *V1alpha2APIResponse) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given LoggingLogs and assigns it to the Query field.
func (o *V1alpha2APIResponse) SetQuery(v LoggingLogs) {
	o.Query = &v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *V1alpha2APIResponse) GetStatistics() LoggingStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret LoggingStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1alpha2APIResponse) GetStatisticsOk() (*LoggingStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *V1alpha2APIResponse) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given LoggingStatistics and assigns it to the Statistics field.
func (o *V1alpha2APIResponse) SetStatistics(v LoggingStatistics) {
	o.Statistics = &v
}

func (o V1alpha2APIResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1alpha2APIResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Histogram) {
		toSerialize["histogram"] = o.Histogram
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	return toSerialize, nil
}

type NullableV1alpha2APIResponse struct {
	value *V1alpha2APIResponse
	isSet bool
}

func (v NullableV1alpha2APIResponse) Get() *V1alpha2APIResponse {
	return v.value
}

func (v *NullableV1alpha2APIResponse) Set(val *V1alpha2APIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1alpha2APIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1alpha2APIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1alpha2APIResponse(val *V1alpha2APIResponse) *NullableV1alpha2APIResponse {
	return &NullableV1alpha2APIResponse{value: val, isSet: true}
}

func (v NullableV1alpha2APIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1alpha2APIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


