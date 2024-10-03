# V1alpha1APIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Histogram** | Pointer to [**EventsHistogram**](EventsHistogram.md) |  | [optional] 
**Query** | Pointer to [**EventsEvents**](EventsEvents.md) |  | [optional] 
**Statistics** | Pointer to [**EventsStatistics**](EventsStatistics.md) |  | [optional] 

## Methods

### NewV1alpha1APIResponse

`func NewV1alpha1APIResponse() *V1alpha1APIResponse`

NewV1alpha1APIResponse instantiates a new V1alpha1APIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1APIResponseWithDefaults

`func NewV1alpha1APIResponseWithDefaults() *V1alpha1APIResponse`

NewV1alpha1APIResponseWithDefaults instantiates a new V1alpha1APIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistogram

`func (o *V1alpha1APIResponse) GetHistogram() EventsHistogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *V1alpha1APIResponse) GetHistogramOk() (*EventsHistogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *V1alpha1APIResponse) SetHistogram(v EventsHistogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *V1alpha1APIResponse) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetQuery

`func (o *V1alpha1APIResponse) GetQuery() EventsEvents`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1alpha1APIResponse) GetQueryOk() (*EventsEvents, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1alpha1APIResponse) SetQuery(v EventsEvents)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V1alpha1APIResponse) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetStatistics

`func (o *V1alpha1APIResponse) GetStatistics() EventsStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V1alpha1APIResponse) GetStatisticsOk() (*EventsStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V1alpha1APIResponse) SetStatistics(v EventsStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V1alpha1APIResponse) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


