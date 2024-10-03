# LoggingStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | **int64** | total number of containers | 
**Logs** | **int64** | total number of logs | 

## Methods

### NewLoggingStatistics

`func NewLoggingStatistics(containers int64, logs int64, ) *LoggingStatistics`

NewLoggingStatistics instantiates a new LoggingStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingStatisticsWithDefaults

`func NewLoggingStatisticsWithDefaults() *LoggingStatistics`

NewLoggingStatisticsWithDefaults instantiates a new LoggingStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *LoggingStatistics) GetContainers() int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *LoggingStatistics) GetContainersOk() (*int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *LoggingStatistics) SetContainers(v int64)`

SetContainers sets Containers field to given value.


### GetLogs

`func (o *LoggingStatistics) GetLogs() int64`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LoggingStatistics) GetLogsOk() (*int64, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LoggingStatistics) SetLogs(v int64)`

SetLogs sets Logs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


