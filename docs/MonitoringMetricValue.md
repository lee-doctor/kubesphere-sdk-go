# MonitoringMetricValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **map[string]string** | time series labels | [optional] 
**Value** | Pointer to **string** | time series, values of vector type | [optional] 
**Values** | Pointer to **[]map[string]interface{}** | time series, values of matrix type | [optional] 

## Methods

### NewMonitoringMetricValue

`func NewMonitoringMetricValue() *MonitoringMetricValue`

NewMonitoringMetricValue instantiates a new MonitoringMetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMetricValueWithDefaults

`func NewMonitoringMetricValueWithDefaults() *MonitoringMetricValue`

NewMonitoringMetricValueWithDefaults instantiates a new MonitoringMetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MonitoringMetricValue) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MonitoringMetricValue) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MonitoringMetricValue) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MonitoringMetricValue) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetValue

`func (o *MonitoringMetricValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MonitoringMetricValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MonitoringMetricValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MonitoringMetricValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *MonitoringMetricValue) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MonitoringMetricValue) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MonitoringMetricValue) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *MonitoringMetricValue) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


