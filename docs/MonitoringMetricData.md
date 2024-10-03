# MonitoringMetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]MonitoringMetricValue**](MonitoringMetricValue.md) | metric data including labels, time series and values | [optional] 
**ResultType** | Pointer to **string** | result type, one of matrix, vector | [optional] 

## Methods

### NewMonitoringMetricData

`func NewMonitoringMetricData() *MonitoringMetricData`

NewMonitoringMetricData instantiates a new MonitoringMetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMetricDataWithDefaults

`func NewMonitoringMetricDataWithDefaults() *MonitoringMetricData`

NewMonitoringMetricDataWithDefaults instantiates a new MonitoringMetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *MonitoringMetricData) GetResult() []MonitoringMetricValue`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MonitoringMetricData) GetResultOk() (*[]MonitoringMetricValue, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MonitoringMetricData) SetResult(v []MonitoringMetricValue)`

SetResult sets Result field to given value.

### HasResult

`func (o *MonitoringMetricData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultType

`func (o *MonitoringMetricData) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *MonitoringMetricData) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *MonitoringMetricData) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *MonitoringMetricData) HasResultType() bool`

HasResultType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


