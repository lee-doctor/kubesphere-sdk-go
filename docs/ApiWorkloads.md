# ApiWorkloads

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]int32** | the number of unhealthy workloads | 
**Items** | Pointer to **map[string]interface{}** | unhealthy workloads | [optional] 
**Namespace** | **string** | the name of the namespace | 

## Methods

### NewApiWorkloads

`func NewApiWorkloads(data map[string]int32, namespace string, ) *ApiWorkloads`

NewApiWorkloads instantiates a new ApiWorkloads object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWorkloadsWithDefaults

`func NewApiWorkloadsWithDefaults() *ApiWorkloads`

NewApiWorkloadsWithDefaults instantiates a new ApiWorkloads object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiWorkloads) GetData() map[string]int32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiWorkloads) GetDataOk() (*map[string]int32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiWorkloads) SetData(v map[string]int32)`

SetData sets Data field to given value.


### GetItems

`func (o *ApiWorkloads) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApiWorkloads) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApiWorkloads) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *ApiWorkloads) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNamespace

`func (o *ApiWorkloads) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiWorkloads) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiWorkloads) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


