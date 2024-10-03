# LoggingHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Histograms** | [**[]LoggingBucket**](LoggingBucket.md) | actual array of histogram results | 
**Total** | **int64** | total number of logs | 

## Methods

### NewLoggingHistogram

`func NewLoggingHistogram(histograms []LoggingBucket, total int64, ) *LoggingHistogram`

NewLoggingHistogram instantiates a new LoggingHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingHistogramWithDefaults

`func NewLoggingHistogramWithDefaults() *LoggingHistogram`

NewLoggingHistogramWithDefaults instantiates a new LoggingHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistograms

`func (o *LoggingHistogram) GetHistograms() []LoggingBucket`

GetHistograms returns the Histograms field if non-nil, zero value otherwise.

### GetHistogramsOk

`func (o *LoggingHistogram) GetHistogramsOk() (*[]LoggingBucket, bool)`

GetHistogramsOk returns a tuple with the Histograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistograms

`func (o *LoggingHistogram) SetHistograms(v []LoggingBucket)`

SetHistograms sets Histograms field to given value.


### GetTotal

`func (o *LoggingHistogram) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LoggingHistogram) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LoggingHistogram) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


