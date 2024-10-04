# V1HealthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **int64** |  | [optional] 
**Retries** | Pointer to **int32** |  | [optional] 
**StartPeriod** | Pointer to **int64** |  | [optional] 
**Test** | Pointer to **[]string** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

## Methods

### NewV1HealthConfig

`func NewV1HealthConfig() *V1HealthConfig`

NewV1HealthConfig instantiates a new V1HealthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HealthConfigWithDefaults

`func NewV1HealthConfigWithDefaults() *V1HealthConfig`

NewV1HealthConfigWithDefaults instantiates a new V1HealthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *V1HealthConfig) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V1HealthConfig) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V1HealthConfig) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *V1HealthConfig) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRetries

`func (o *V1HealthConfig) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *V1HealthConfig) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *V1HealthConfig) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *V1HealthConfig) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetStartPeriod

`func (o *V1HealthConfig) GetStartPeriod() int64`

GetStartPeriod returns the StartPeriod field if non-nil, zero value otherwise.

### GetStartPeriodOk

`func (o *V1HealthConfig) GetStartPeriodOk() (*int64, bool)`

GetStartPeriodOk returns a tuple with the StartPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPeriod

`func (o *V1HealthConfig) SetStartPeriod(v int64)`

SetStartPeriod sets StartPeriod field to given value.

### HasStartPeriod

`func (o *V1HealthConfig) HasStartPeriod() bool`

HasStartPeriod returns a boolean if a field has been set.

### GetTest

`func (o *V1HealthConfig) GetTest() []string`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *V1HealthConfig) GetTestOk() (*[]string, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *V1HealthConfig) SetTest(v []string)`

SetTest sets Test field to given value.

### HasTest

`func (o *V1HealthConfig) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetTimeout

`func (o *V1HealthConfig) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *V1HealthConfig) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *V1HealthConfig) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *V1HealthConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


