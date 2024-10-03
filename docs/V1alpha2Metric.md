# V1alpha2Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | **float64** |  | 
**Min** | **float64** |  | 
**Samples** | Pointer to [**[]V1alpha2Sample**](V1alpha2Sample.md) |  | [optional] 

## Methods

### NewV1alpha2Metric

`func NewV1alpha2Metric(max float64, min float64, ) *V1alpha2Metric`

NewV1alpha2Metric instantiates a new V1alpha2Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2MetricWithDefaults

`func NewV1alpha2MetricWithDefaults() *V1alpha2Metric`

NewV1alpha2MetricWithDefaults instantiates a new V1alpha2Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *V1alpha2Metric) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V1alpha2Metric) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V1alpha2Metric) SetMax(v float64)`

SetMax sets Max field to given value.


### GetMin

`func (o *V1alpha2Metric) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V1alpha2Metric) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V1alpha2Metric) SetMin(v float64)`

SetMin sets Min field to given value.


### GetSamples

`func (o *V1alpha2Metric) GetSamples() []V1alpha2Sample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *V1alpha2Metric) GetSamplesOk() (*[]V1alpha2Sample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *V1alpha2Metric) SetSamples(v []V1alpha2Sample)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *V1alpha2Metric) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


