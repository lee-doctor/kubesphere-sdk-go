# V1alpha2APIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Histogram** | Pointer to [**LoggingHistogram**](LoggingHistogram.md) |  | [optional] 
**Query** | Pointer to [**LoggingLogs**](LoggingLogs.md) |  | [optional] 
**Statistics** | Pointer to [**LoggingStatistics**](LoggingStatistics.md) |  | [optional] 

## Methods

### NewV1alpha2APIResponse

`func NewV1alpha2APIResponse() *V1alpha2APIResponse`

NewV1alpha2APIResponse instantiates a new V1alpha2APIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2APIResponseWithDefaults

`func NewV1alpha2APIResponseWithDefaults() *V1alpha2APIResponse`

NewV1alpha2APIResponseWithDefaults instantiates a new V1alpha2APIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistogram

`func (o *V1alpha2APIResponse) GetHistogram() LoggingHistogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *V1alpha2APIResponse) GetHistogramOk() (*LoggingHistogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *V1alpha2APIResponse) SetHistogram(v LoggingHistogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *V1alpha2APIResponse) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetQuery

`func (o *V1alpha2APIResponse) GetQuery() LoggingLogs`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1alpha2APIResponse) GetQueryOk() (*LoggingLogs, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1alpha2APIResponse) SetQuery(v LoggingLogs)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V1alpha2APIResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetStatistics

`func (o *V1alpha2APIResponse) GetStatistics() LoggingStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V1alpha2APIResponse) GetStatisticsOk() (*LoggingStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V1alpha2APIResponse) SetStatistics(v LoggingStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V1alpha2APIResponse) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


