# LoggingLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]LoggingRecord**](LoggingRecord.md) | actual array of results | [optional] 
**Total** | **int64** | total number of matched results | 

## Methods

### NewLoggingLogs

`func NewLoggingLogs(total int64, ) *LoggingLogs`

NewLoggingLogs instantiates a new LoggingLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingLogsWithDefaults

`func NewLoggingLogsWithDefaults() *LoggingLogs`

NewLoggingLogsWithDefaults instantiates a new LoggingLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *LoggingLogs) GetRecords() []LoggingRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *LoggingLogs) GetRecordsOk() (*[]LoggingRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *LoggingLogs) SetRecords(v []LoggingRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *LoggingLogs) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *LoggingLogs) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LoggingLogs) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LoggingLogs) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


