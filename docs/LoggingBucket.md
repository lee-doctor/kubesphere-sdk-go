# LoggingBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | total number of logs at intervals | 
**Time** | **int64** | timestamp | 

## Methods

### NewLoggingBucket

`func NewLoggingBucket(count int64, time int64, ) *LoggingBucket`

NewLoggingBucket instantiates a new LoggingBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingBucketWithDefaults

`func NewLoggingBucketWithDefaults() *LoggingBucket`

NewLoggingBucketWithDefaults instantiates a new LoggingBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *LoggingBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LoggingBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LoggingBucket) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTime

`func (o *LoggingBucket) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LoggingBucket) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LoggingBucket) SetTime(v int64)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


