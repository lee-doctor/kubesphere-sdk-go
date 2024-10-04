# V1beta1SubjectAccessReviewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** |  | 
**Denied** | Pointer to **bool** |  | [optional] 
**EvaluationError** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewV1beta1SubjectAccessReviewStatus

`func NewV1beta1SubjectAccessReviewStatus(allowed bool, ) *V1beta1SubjectAccessReviewStatus`

NewV1beta1SubjectAccessReviewStatus instantiates a new V1beta1SubjectAccessReviewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1SubjectAccessReviewStatusWithDefaults

`func NewV1beta1SubjectAccessReviewStatusWithDefaults() *V1beta1SubjectAccessReviewStatus`

NewV1beta1SubjectAccessReviewStatusWithDefaults instantiates a new V1beta1SubjectAccessReviewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *V1beta1SubjectAccessReviewStatus) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *V1beta1SubjectAccessReviewStatus) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *V1beta1SubjectAccessReviewStatus) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetDenied

`func (o *V1beta1SubjectAccessReviewStatus) GetDenied() bool`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *V1beta1SubjectAccessReviewStatus) GetDeniedOk() (*bool, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *V1beta1SubjectAccessReviewStatus) SetDenied(v bool)`

SetDenied sets Denied field to given value.

### HasDenied

`func (o *V1beta1SubjectAccessReviewStatus) HasDenied() bool`

HasDenied returns a boolean if a field has been set.

### GetEvaluationError

`func (o *V1beta1SubjectAccessReviewStatus) GetEvaluationError() string`

GetEvaluationError returns the EvaluationError field if non-nil, zero value otherwise.

### GetEvaluationErrorOk

`func (o *V1beta1SubjectAccessReviewStatus) GetEvaluationErrorOk() (*string, bool)`

GetEvaluationErrorOk returns a tuple with the EvaluationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationError

`func (o *V1beta1SubjectAccessReviewStatus) SetEvaluationError(v string)`

SetEvaluationError sets EvaluationError field to given value.

### HasEvaluationError

`func (o *V1beta1SubjectAccessReviewStatus) HasEvaluationError() bool`

HasEvaluationError returns a boolean if a field has been set.

### GetReason

`func (o *V1beta1SubjectAccessReviewStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1beta1SubjectAccessReviewStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1beta1SubjectAccessReviewStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1beta1SubjectAccessReviewStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


