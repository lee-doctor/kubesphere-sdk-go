# V1alpha2BadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **map[string]interface{}** |  | 
**Status** | **int32** |  | 

## Methods

### NewV1alpha2BadRequestError

`func NewV1alpha2BadRequestError(reason map[string]interface{}, status int32, ) *V1alpha2BadRequestError`

NewV1alpha2BadRequestError instantiates a new V1alpha2BadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2BadRequestErrorWithDefaults

`func NewV1alpha2BadRequestErrorWithDefaults() *V1alpha2BadRequestError`

NewV1alpha2BadRequestErrorWithDefaults instantiates a new V1alpha2BadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *V1alpha2BadRequestError) GetReason() map[string]interface{}`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1alpha2BadRequestError) GetReasonOk() (*map[string]interface{}, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1alpha2BadRequestError) SetReason(v map[string]interface{})`

SetReason sets Reason field to given value.


### GetStatus

`func (o *V1alpha2BadRequestError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha2BadRequestError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha2BadRequestError) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


