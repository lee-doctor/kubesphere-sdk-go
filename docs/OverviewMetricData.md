# OverviewMetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**[]OverviewMetricValue**](OverviewMetricValue.md) |  | 
**ResultType** | **string** |  | 

## Methods

### NewOverviewMetricData

`func NewOverviewMetricData(result []OverviewMetricValue, resultType string, ) *OverviewMetricData`

NewOverviewMetricData instantiates a new OverviewMetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewMetricDataWithDefaults

`func NewOverviewMetricDataWithDefaults() *OverviewMetricData`

NewOverviewMetricDataWithDefaults instantiates a new OverviewMetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *OverviewMetricData) GetResult() []OverviewMetricValue`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OverviewMetricData) GetResultOk() (*[]OverviewMetricValue, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OverviewMetricData) SetResult(v []OverviewMetricValue)`

SetResult sets Result field to given value.


### GetResultType

`func (o *OverviewMetricData) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *OverviewMetricData) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *OverviewMetricData) SetResultType(v string)`

SetResultType sets ResultType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


