# V1alpha2MetricRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**Group** | **string** |  | 
**ID** | **string** |  | 
**Label** | **string** |  | 
**Metric** | [**V1alpha2Metric**](V1alpha2Metric.md) |  | 
**Priority** | **float64** |  | 
**URL** | **string** |  | 
**Value** | **float64** |  | 
**ValueEmpty** | **bool** |  | 

## Methods

### NewV1alpha2MetricRow

`func NewV1alpha2MetricRow(format string, group string, iD string, label string, metric V1alpha2Metric, priority float64, uRL string, value float64, valueEmpty bool, ) *V1alpha2MetricRow`

NewV1alpha2MetricRow instantiates a new V1alpha2MetricRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2MetricRowWithDefaults

`func NewV1alpha2MetricRowWithDefaults() *V1alpha2MetricRow`

NewV1alpha2MetricRowWithDefaults instantiates a new V1alpha2MetricRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *V1alpha2MetricRow) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V1alpha2MetricRow) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V1alpha2MetricRow) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetGroup

`func (o *V1alpha2MetricRow) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V1alpha2MetricRow) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V1alpha2MetricRow) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetID

`func (o *V1alpha2MetricRow) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *V1alpha2MetricRow) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *V1alpha2MetricRow) SetID(v string)`

SetID sets ID field to given value.


### GetLabel

`func (o *V1alpha2MetricRow) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1alpha2MetricRow) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1alpha2MetricRow) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetric

`func (o *V1alpha2MetricRow) GetMetric() V1alpha2Metric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *V1alpha2MetricRow) GetMetricOk() (*V1alpha2Metric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *V1alpha2MetricRow) SetMetric(v V1alpha2Metric)`

SetMetric sets Metric field to given value.


### GetPriority

`func (o *V1alpha2MetricRow) GetPriority() float64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1alpha2MetricRow) GetPriorityOk() (*float64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1alpha2MetricRow) SetPriority(v float64)`

SetPriority sets Priority field to given value.


### GetURL

`func (o *V1alpha2MetricRow) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *V1alpha2MetricRow) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *V1alpha2MetricRow) SetURL(v string)`

SetURL sets URL field to given value.


### GetValue

`func (o *V1alpha2MetricRow) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1alpha2MetricRow) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1alpha2MetricRow) SetValue(v float64)`

SetValue sets Value field to given value.


### GetValueEmpty

`func (o *V1alpha2MetricRow) GetValueEmpty() bool`

GetValueEmpty returns the ValueEmpty field if non-nil, zero value otherwise.

### GetValueEmptyOk

`func (o *V1alpha2MetricRow) GetValueEmptyOk() (*bool, bool)`

GetValueEmptyOk returns a tuple with the ValueEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueEmpty

`func (o *V1alpha2MetricRow) SetValueEmpty(v bool)`

SetValueEmpty sets ValueEmpty field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


