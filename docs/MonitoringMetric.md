# MonitoringMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MonitoringMetricData**](MonitoringMetricData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**MetricName** | Pointer to **string** | metric name, eg. scheduler_up_sum | [optional] 

## Methods

### NewMonitoringMetric

`func NewMonitoringMetric() *MonitoringMetric`

NewMonitoringMetric instantiates a new MonitoringMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMetricWithDefaults

`func NewMonitoringMetricWithDefaults() *MonitoringMetric`

NewMonitoringMetricWithDefaults instantiates a new MonitoringMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MonitoringMetric) GetData() MonitoringMetricData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MonitoringMetric) GetDataOk() (*MonitoringMetricData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MonitoringMetric) SetData(v MonitoringMetricData)`

SetData sets Data field to given value.

### HasData

`func (o *MonitoringMetric) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *MonitoringMetric) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MonitoringMetric) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MonitoringMetric) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MonitoringMetric) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMetricName

`func (o *MonitoringMetric) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MonitoringMetric) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MonitoringMetric) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *MonitoringMetric) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


