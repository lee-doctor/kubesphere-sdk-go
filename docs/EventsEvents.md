# EventsEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]V1Event**](V1Event.md) | actual array of results | 
**Total** | **int64** | total number of matched results | 

## Methods

### NewEventsEvents

`func NewEventsEvents(records []V1Event, total int64, ) *EventsEvents`

NewEventsEvents instantiates a new EventsEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsEventsWithDefaults

`func NewEventsEventsWithDefaults() *EventsEvents`

NewEventsEventsWithDefaults instantiates a new EventsEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *EventsEvents) GetRecords() []V1Event`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *EventsEvents) GetRecordsOk() (*[]V1Event, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *EventsEvents) SetRecords(v []V1Event)`

SetRecords sets Records field to given value.


### GetTotal

`func (o *EventsEvents) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EventsEvents) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EventsEvents) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


