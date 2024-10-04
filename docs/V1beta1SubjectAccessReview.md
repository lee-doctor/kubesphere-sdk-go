# V1beta1SubjectAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Spec** | [**V1beta1SubjectAccessReviewSpec**](V1beta1SubjectAccessReviewSpec.md) |  | 
**Status** | Pointer to [**V1beta1SubjectAccessReviewStatus**](V1beta1SubjectAccessReviewStatus.md) |  | [optional] 

## Methods

### NewV1beta1SubjectAccessReview

`func NewV1beta1SubjectAccessReview(spec V1beta1SubjectAccessReviewSpec, ) *V1beta1SubjectAccessReview`

NewV1beta1SubjectAccessReview instantiates a new V1beta1SubjectAccessReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1SubjectAccessReviewWithDefaults

`func NewV1beta1SubjectAccessReviewWithDefaults() *V1beta1SubjectAccessReview`

NewV1beta1SubjectAccessReviewWithDefaults instantiates a new V1beta1SubjectAccessReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1beta1SubjectAccessReview) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1beta1SubjectAccessReview) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1beta1SubjectAccessReview) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1beta1SubjectAccessReview) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1beta1SubjectAccessReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1beta1SubjectAccessReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1beta1SubjectAccessReview) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1beta1SubjectAccessReview) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSpec

`func (o *V1beta1SubjectAccessReview) GetSpec() V1beta1SubjectAccessReviewSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1beta1SubjectAccessReview) GetSpecOk() (*V1beta1SubjectAccessReviewSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1beta1SubjectAccessReview) SetSpec(v V1beta1SubjectAccessReviewSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *V1beta1SubjectAccessReview) GetStatus() V1beta1SubjectAccessReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1beta1SubjectAccessReview) GetStatusOk() (*V1beta1SubjectAccessReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1beta1SubjectAccessReview) SetStatus(v V1beta1SubjectAccessReviewStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1beta1SubjectAccessReview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


