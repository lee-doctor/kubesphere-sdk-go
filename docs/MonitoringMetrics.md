# MonitoringMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | current page returned | [optional] 
**Results** | [**[]MonitoringMetric**](MonitoringMetric.md) | actual array of results | 
**TotalItem** | Pointer to **int32** | page size | [optional] 
**TotalPage** | Pointer to **int32** | total number of pages | [optional] 

## Methods

### NewMonitoringMetrics

`func NewMonitoringMetrics(results []MonitoringMetric, ) *MonitoringMetrics`

NewMonitoringMetrics instantiates a new MonitoringMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMetricsWithDefaults

`func NewMonitoringMetricsWithDefaults() *MonitoringMetrics`

NewMonitoringMetricsWithDefaults instantiates a new MonitoringMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *MonitoringMetrics) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *MonitoringMetrics) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *MonitoringMetrics) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *MonitoringMetrics) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *MonitoringMetrics) GetResults() []MonitoringMetric`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MonitoringMetrics) GetResultsOk() (*[]MonitoringMetric, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MonitoringMetrics) SetResults(v []MonitoringMetric)`

SetResults sets Results field to given value.


### GetTotalItem

`func (o *MonitoringMetrics) GetTotalItem() int32`

GetTotalItem returns the TotalItem field if non-nil, zero value otherwise.

### GetTotalItemOk

`func (o *MonitoringMetrics) GetTotalItemOk() (*int32, bool)`

GetTotalItemOk returns a tuple with the TotalItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItem

`func (o *MonitoringMetrics) SetTotalItem(v int32)`

SetTotalItem sets TotalItem field to given value.

### HasTotalItem

`func (o *MonitoringMetrics) HasTotalItem() bool`

HasTotalItem returns a boolean if a field has been set.

### GetTotalPage

`func (o *MonitoringMetrics) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *MonitoringMetrics) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *MonitoringMetrics) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *MonitoringMetrics) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


