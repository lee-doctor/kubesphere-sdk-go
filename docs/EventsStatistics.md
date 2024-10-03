# EventsStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **int64** | total number of events | 
**Resources** | **int64** | total number of resources | 

## Methods

### NewEventsStatistics

`func NewEventsStatistics(events int64, resources int64, ) *EventsStatistics`

NewEventsStatistics instantiates a new EventsStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsStatisticsWithDefaults

`func NewEventsStatisticsWithDefaults() *EventsStatistics`

NewEventsStatisticsWithDefaults instantiates a new EventsStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsStatistics) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsStatistics) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsStatistics) SetEvents(v int64)`

SetEvents sets Events field to given value.


### GetResources

`func (o *EventsStatistics) GetResources() int64`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EventsStatistics) GetResourcesOk() (*int64, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EventsStatistics) SetResources(v int64)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


