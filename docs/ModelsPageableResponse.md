# ModelsPageableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]map[string]interface{}** | paging data | 
**TotalCount** | **int32** | total count | 

## Methods

### NewModelsPageableResponse

`func NewModelsPageableResponse(items []map[string]interface{}, totalCount int32, ) *ModelsPageableResponse`

NewModelsPageableResponse instantiates a new ModelsPageableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsPageableResponseWithDefaults

`func NewModelsPageableResponseWithDefaults() *ModelsPageableResponse`

NewModelsPageableResponseWithDefaults instantiates a new ModelsPageableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ModelsPageableResponse) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelsPageableResponse) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelsPageableResponse) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ModelsPageableResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ModelsPageableResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ModelsPageableResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


