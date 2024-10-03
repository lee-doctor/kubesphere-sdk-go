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

// checks if the OpenpitrixAppVersionReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenpitrixAppVersionReview{}

// OpenpitrixAppVersionReview struct for OpenpitrixAppVersionReview
type OpenpitrixAppVersionReview struct {
	AppId *string `json:"app_id,omitempty"`
	AppName *string `json:"app_name,omitempty"`
	Phase *map[string]OpenpitrixAppVersionReviewPhase `json:"phase,omitempty"`
	ReviewId *string `json:"review_id,omitempty"`
	Reviewer *string `json:"reviewer,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusTime *string `json:"status_time,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
	VersionName *string `json:"version_name,omitempty"`
	VersionType *string `json:"version_type,omitempty"`
}

// NewOpenpitrixAppVersionReview instantiates a new OpenpitrixAppVersionReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenpitrixAppVersionReview() *OpenpitrixAppVersionReview {
	this := OpenpitrixAppVersionReview{}
	return &this
}

// NewOpenpitrixAppVersionReviewWithDefaults instantiates a new OpenpitrixAppVersionReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenpitrixAppVersionReviewWithDefaults() *OpenpitrixAppVersionReview {
	this := OpenpitrixAppVersionReview{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *OpenpitrixAppVersionReview) SetAppId(v string) {
	o.AppId = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *OpenpitrixAppVersionReview) SetAppName(v string) {
	o.AppName = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetPhase() map[string]OpenpitrixAppVersionReviewPhase {
	if o == nil || IsNil(o.Phase) {
		var ret map[string]OpenpitrixAppVersionReviewPhase
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetPhaseOk() (*map[string]OpenpitrixAppVersionReviewPhase, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given map[string]OpenpitrixAppVersionReviewPhase and assigns it to the Phase field.
func (o *OpenpitrixAppVersionReview) SetPhase(v map[string]OpenpitrixAppVersionReviewPhase) {
	o.Phase = &v
}

// GetReviewId returns the ReviewId field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetReviewId() string {
	if o == nil || IsNil(o.ReviewId) {
		var ret string
		return ret
	}
	return *o.ReviewId
}

// GetReviewIdOk returns a tuple with the ReviewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetReviewIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewId) {
		return nil, false
	}
	return o.ReviewId, true
}

// HasReviewId returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasReviewId() bool {
	if o != nil && !IsNil(o.ReviewId) {
		return true
	}

	return false
}

// SetReviewId gets a reference to the given string and assigns it to the ReviewId field.
func (o *OpenpitrixAppVersionReview) SetReviewId(v string) {
	o.ReviewId = &v
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetReviewer() string {
	if o == nil || IsNil(o.Reviewer) {
		var ret string
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetReviewerOk() (*string, bool) {
	if o == nil || IsNil(o.Reviewer) {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasReviewer() bool {
	if o != nil && !IsNil(o.Reviewer) {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given string and assigns it to the Reviewer field.
func (o *OpenpitrixAppVersionReview) SetReviewer(v string) {
	o.Reviewer = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OpenpitrixAppVersionReview) SetStatus(v string) {
	o.Status = &v
}

// GetStatusTime returns the StatusTime field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetStatusTime() string {
	if o == nil || IsNil(o.StatusTime) {
		var ret string
		return ret
	}
	return *o.StatusTime
}

// GetStatusTimeOk returns a tuple with the StatusTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetStatusTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StatusTime) {
		return nil, false
	}
	return o.StatusTime, true
}

// HasStatusTime returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasStatusTime() bool {
	if o != nil && !IsNil(o.StatusTime) {
		return true
	}

	return false
}

// SetStatusTime gets a reference to the given string and assigns it to the StatusTime field.
func (o *OpenpitrixAppVersionReview) SetStatusTime(v string) {
	o.StatusTime = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *OpenpitrixAppVersionReview) SetVersionId(v string) {
	o.VersionId = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetVersionName() string {
	if o == nil || IsNil(o.VersionName) {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetVersionNameOk() (*string, bool) {
	if o == nil || IsNil(o.VersionName) {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasVersionName() bool {
	if o != nil && !IsNil(o.VersionName) {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *OpenpitrixAppVersionReview) SetVersionName(v string) {
	o.VersionName = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *OpenpitrixAppVersionReview) GetVersionType() string {
	if o == nil || IsNil(o.VersionType) {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenpitrixAppVersionReview) GetVersionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VersionType) {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *OpenpitrixAppVersionReview) HasVersionType() bool {
	if o != nil && !IsNil(o.VersionType) {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *OpenpitrixAppVersionReview) SetVersionType(v string) {
	o.VersionType = &v
}

func (o OpenpitrixAppVersionReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenpitrixAppVersionReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !IsNil(o.ReviewId) {
		toSerialize["review_id"] = o.ReviewId
	}
	if !IsNil(o.Reviewer) {
		toSerialize["reviewer"] = o.Reviewer
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusTime) {
		toSerialize["status_time"] = o.StatusTime
	}
	if !IsNil(o.VersionId) {
		toSerialize["version_id"] = o.VersionId
	}
	if !IsNil(o.VersionName) {
		toSerialize["version_name"] = o.VersionName
	}
	if !IsNil(o.VersionType) {
		toSerialize["version_type"] = o.VersionType
	}
	return toSerialize, nil
}

type NullableOpenpitrixAppVersionReview struct {
	value *OpenpitrixAppVersionReview
	isSet bool
}

func (v NullableOpenpitrixAppVersionReview) Get() *OpenpitrixAppVersionReview {
	return v.value
}

func (v *NullableOpenpitrixAppVersionReview) Set(val *OpenpitrixAppVersionReview) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenpitrixAppVersionReview) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenpitrixAppVersionReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenpitrixAppVersionReview(val *OpenpitrixAppVersionReview) *NullableOpenpitrixAppVersionReview {
	return &NullableOpenpitrixAppVersionReview{value: val, isSet: true}
}

func (v NullableOpenpitrixAppVersionReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenpitrixAppVersionReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


