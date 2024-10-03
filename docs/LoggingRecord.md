# LoggingRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | container name | [optional] 
**Log** | Pointer to **string** | log message | [optional] 
**Namespace** | Pointer to **string** | namespace | [optional] 
**Pod** | Pointer to **string** | pod name | [optional] 
**Time** | Pointer to **string** | log timestamp | [optional] 

## Methods

### NewLoggingRecord

`func NewLoggingRecord() *LoggingRecord`

NewLoggingRecord instantiates a new LoggingRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingRecordWithDefaults

`func NewLoggingRecordWithDefaults() *LoggingRecord`

NewLoggingRecordWithDefaults instantiates a new LoggingRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *LoggingRecord) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *LoggingRecord) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *LoggingRecord) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *LoggingRecord) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetLog

`func (o *LoggingRecord) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *LoggingRecord) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *LoggingRecord) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *LoggingRecord) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetNamespace

`func (o *LoggingRecord) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *LoggingRecord) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *LoggingRecord) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *LoggingRecord) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPod

`func (o *LoggingRecord) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *LoggingRecord) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *LoggingRecord) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *LoggingRecord) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetTime

`func (o *LoggingRecord) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LoggingRecord) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LoggingRecord) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *LoggingRecord) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


