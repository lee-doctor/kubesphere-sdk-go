# Repositories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Items** | Pointer to [**[]RepositoriesItems**](RepositoriesItems.md) |  | [optional] 
**LastPage** | Pointer to **map[string]interface{}** |  | [optional] 
**NextPage** | Pointer to **map[string]interface{}** |  | [optional] 
**PageSize** | Pointer to **int32** | page size | [optional] 

## Methods

### NewRepositories

`func NewRepositories() *Repositories`

NewRepositories instantiates a new Repositories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoriesWithDefaults

`func NewRepositoriesWithDefaults() *Repositories`

NewRepositoriesWithDefaults instantiates a new Repositories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *Repositories) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Repositories) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Repositories) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Repositories) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *Repositories) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Repositories) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Repositories) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Repositories) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *Repositories) GetItems() []RepositoriesItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Repositories) GetItemsOk() (*[]RepositoriesItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Repositories) SetItems(v []RepositoriesItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *Repositories) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLastPage

`func (o *Repositories) GetLastPage() map[string]interface{}`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *Repositories) GetLastPageOk() (*map[string]interface{}, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *Repositories) SetLastPage(v map[string]interface{})`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *Repositories) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetNextPage

`func (o *Repositories) GetNextPage() map[string]interface{}`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *Repositories) GetNextPageOk() (*map[string]interface{}, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *Repositories) SetNextPage(v map[string]interface{})`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *Repositories) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPageSize

`func (o *Repositories) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Repositories) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Repositories) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Repositories) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


