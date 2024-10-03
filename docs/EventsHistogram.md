# EventsHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]EventsBucket**](EventsBucket.md) | actual array of histogram results | 
**Total** | **int64** | total number of events | 

## Methods

### NewEventsHistogram

`func NewEventsHistogram(buckets []EventsBucket, total int64, ) *EventsHistogram`

NewEventsHistogram instantiates a new EventsHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsHistogramWithDefaults

`func NewEventsHistogramWithDefaults() *EventsHistogram`

NewEventsHistogramWithDefaults instantiates a new EventsHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *EventsHistogram) GetBuckets() []EventsBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *EventsHistogram) GetBucketsOk() (*[]EventsBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *EventsHistogram) SetBuckets(v []EventsBucket)`

SetBuckets sets Buckets field to given value.


### GetTotal

`func (o *EventsHistogram) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EventsHistogram) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EventsHistogram) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


