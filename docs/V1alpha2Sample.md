# V1alpha2Sample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**Value** | **float64** |  | 

## Methods

### NewV1alpha2Sample

`func NewV1alpha2Sample(date time.Time, value float64, ) *V1alpha2Sample`

NewV1alpha2Sample instantiates a new V1alpha2Sample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2SampleWithDefaults

`func NewV1alpha2SampleWithDefaults() *V1alpha2Sample`

NewV1alpha2SampleWithDefaults instantiates a new V1alpha2Sample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *V1alpha2Sample) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *V1alpha2Sample) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *V1alpha2Sample) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetValue

`func (o *V1alpha2Sample) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1alpha2Sample) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1alpha2Sample) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


