# V1alpha3TimerTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | Pointer to **string** | jenkins cron script | [optional] 
**Interval** | Pointer to **string** | interval ms | [optional] 

## Methods

### NewV1alpha3TimerTrigger

`func NewV1alpha3TimerTrigger() *V1alpha3TimerTrigger`

NewV1alpha3TimerTrigger instantiates a new V1alpha3TimerTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3TimerTriggerWithDefaults

`func NewV1alpha3TimerTriggerWithDefaults() *V1alpha3TimerTrigger`

NewV1alpha3TimerTriggerWithDefaults instantiates a new V1alpha3TimerTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *V1alpha3TimerTrigger) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *V1alpha3TimerTrigger) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *V1alpha3TimerTrigger) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *V1alpha3TimerTrigger) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetInterval

`func (o *V1alpha3TimerTrigger) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V1alpha3TimerTrigger) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V1alpha3TimerTrigger) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *V1alpha3TimerTrigger) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


