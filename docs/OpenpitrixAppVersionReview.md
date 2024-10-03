# OpenpitrixAppVersionReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**AppName** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to [**map[string]OpenpitrixAppVersionReviewPhase**](OpenpitrixAppVersionReviewPhase.md) |  | [optional] 
**ReviewId** | Pointer to **string** |  | [optional] 
**Reviewer** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusTime** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VersionName** | Pointer to **string** |  | [optional] 
**VersionType** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenpitrixAppVersionReview

`func NewOpenpitrixAppVersionReview() *OpenpitrixAppVersionReview`

NewOpenpitrixAppVersionReview instantiates a new OpenpitrixAppVersionReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenpitrixAppVersionReviewWithDefaults

`func NewOpenpitrixAppVersionReviewWithDefaults() *OpenpitrixAppVersionReview`

NewOpenpitrixAppVersionReviewWithDefaults instantiates a new OpenpitrixAppVersionReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *OpenpitrixAppVersionReview) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *OpenpitrixAppVersionReview) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *OpenpitrixAppVersionReview) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *OpenpitrixAppVersionReview) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppName

`func (o *OpenpitrixAppVersionReview) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *OpenpitrixAppVersionReview) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *OpenpitrixAppVersionReview) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *OpenpitrixAppVersionReview) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetPhase

`func (o *OpenpitrixAppVersionReview) GetPhase() map[string]OpenpitrixAppVersionReviewPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *OpenpitrixAppVersionReview) GetPhaseOk() (*map[string]OpenpitrixAppVersionReviewPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *OpenpitrixAppVersionReview) SetPhase(v map[string]OpenpitrixAppVersionReviewPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *OpenpitrixAppVersionReview) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReviewId

`func (o *OpenpitrixAppVersionReview) GetReviewId() string`

GetReviewId returns the ReviewId field if non-nil, zero value otherwise.

### GetReviewIdOk

`func (o *OpenpitrixAppVersionReview) GetReviewIdOk() (*string, bool)`

GetReviewIdOk returns a tuple with the ReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewId

`func (o *OpenpitrixAppVersionReview) SetReviewId(v string)`

SetReviewId sets ReviewId field to given value.

### HasReviewId

`func (o *OpenpitrixAppVersionReview) HasReviewId() bool`

HasReviewId returns a boolean if a field has been set.

### GetReviewer

`func (o *OpenpitrixAppVersionReview) GetReviewer() string`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *OpenpitrixAppVersionReview) GetReviewerOk() (*string, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *OpenpitrixAppVersionReview) SetReviewer(v string)`

SetReviewer sets Reviewer field to given value.

### HasReviewer

`func (o *OpenpitrixAppVersionReview) HasReviewer() bool`

HasReviewer returns a boolean if a field has been set.

### GetStatus

`func (o *OpenpitrixAppVersionReview) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenpitrixAppVersionReview) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenpitrixAppVersionReview) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenpitrixAppVersionReview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *OpenpitrixAppVersionReview) GetStatusTime() string`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *OpenpitrixAppVersionReview) GetStatusTimeOk() (*string, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *OpenpitrixAppVersionReview) SetStatusTime(v string)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *OpenpitrixAppVersionReview) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetVersionId

`func (o *OpenpitrixAppVersionReview) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *OpenpitrixAppVersionReview) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *OpenpitrixAppVersionReview) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *OpenpitrixAppVersionReview) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *OpenpitrixAppVersionReview) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *OpenpitrixAppVersionReview) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *OpenpitrixAppVersionReview) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *OpenpitrixAppVersionReview) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetVersionType

`func (o *OpenpitrixAppVersionReview) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *OpenpitrixAppVersionReview) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *OpenpitrixAppVersionReview) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.

### HasVersionType

`func (o *OpenpitrixAppVersionReview) HasVersionType() bool`

HasVersionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


