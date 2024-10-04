# OverviewMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**OverviewMetricData**](OverviewMetricData.md) |  | 
**MetricName** | **string** |  | 

## Methods

### NewOverviewMetric

`func NewOverviewMetric(data OverviewMetricData, metricName string, ) *OverviewMetric`

NewOverviewMetric instantiates a new OverviewMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewMetricWithDefaults

`func NewOverviewMetricWithDefaults() *OverviewMetric`

NewOverviewMetricWithDefaults instantiates a new OverviewMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OverviewMetric) GetData() OverviewMetricData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OverviewMetric) GetDataOk() (*OverviewMetricData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OverviewMetric) SetData(v OverviewMetricData)`

SetData sets Data field to given value.


### GetMetricName

`func (o *OverviewMetric) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *OverviewMetric) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *OverviewMetric) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


