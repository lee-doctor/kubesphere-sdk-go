# ApiListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]map[string]interface{}** |  | 
**TotalItems** | **int32** |  | 

## Methods

### NewApiListResult

`func NewApiListResult(items []map[string]interface{}, totalItems int32, ) *ApiListResult`

NewApiListResult instantiates a new ApiListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiListResultWithDefaults

`func NewApiListResultWithDefaults() *ApiListResult`

NewApiListResultWithDefaults instantiates a new ApiListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ApiListResult) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApiListResult) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApiListResult) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.


### GetTotalItems

`func (o *ApiListResult) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiListResult) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiListResult) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


