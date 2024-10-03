# OauthTokenReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Kubernetes API version | 
**Kind** | **string** | kind of the API object | 
**Spec** | Pointer to [**OauthSpec**](OauthSpec.md) |  | [optional] 
**Status** | Pointer to [**OauthStatus**](OauthStatus.md) |  | [optional] 

## Methods

### NewOauthTokenReview

`func NewOauthTokenReview(apiVersion string, kind string, ) *OauthTokenReview`

NewOauthTokenReview instantiates a new OauthTokenReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthTokenReviewWithDefaults

`func NewOauthTokenReviewWithDefaults() *OauthTokenReview`

NewOauthTokenReviewWithDefaults instantiates a new OauthTokenReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OauthTokenReview) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OauthTokenReview) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OauthTokenReview) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *OauthTokenReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OauthTokenReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OauthTokenReview) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSpec

`func (o *OauthTokenReview) GetSpec() OauthSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OauthTokenReview) GetSpecOk() (*OauthSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OauthTokenReview) SetSpec(v OauthSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *OauthTokenReview) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *OauthTokenReview) GetStatus() OauthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OauthTokenReview) GetStatusOk() (*OauthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OauthTokenReview) SetStatus(v OauthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OauthTokenReview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


