# V1alpha2Table

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]V1alpha2Column**](V1alpha2Column.md) |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**Rows** | [**[]V1alpha2Row**](V1alpha2Row.md) |  | 
**TruncationCount** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewV1alpha2Table

`func NewV1alpha2Table(columns []V1alpha2Column, id string, label string, rows []V1alpha2Row, type_ string, ) *V1alpha2Table`

NewV1alpha2Table instantiates a new V1alpha2Table object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2TableWithDefaults

`func NewV1alpha2TableWithDefaults() *V1alpha2Table`

NewV1alpha2TableWithDefaults instantiates a new V1alpha2Table object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *V1alpha2Table) GetColumns() []V1alpha2Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *V1alpha2Table) GetColumnsOk() (*[]V1alpha2Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *V1alpha2Table) SetColumns(v []V1alpha2Column)`

SetColumns sets Columns field to given value.


### GetId

`func (o *V1alpha2Table) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1alpha2Table) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1alpha2Table) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *V1alpha2Table) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1alpha2Table) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1alpha2Table) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRows

`func (o *V1alpha2Table) GetRows() []V1alpha2Row`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *V1alpha2Table) GetRowsOk() (*[]V1alpha2Row, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *V1alpha2Table) SetRows(v []V1alpha2Row)`

SetRows sets Rows field to given value.


### GetTruncationCount

`func (o *V1alpha2Table) GetTruncationCount() int32`

GetTruncationCount returns the TruncationCount field if non-nil, zero value otherwise.

### GetTruncationCountOk

`func (o *V1alpha2Table) GetTruncationCountOk() (*int32, bool)`

GetTruncationCountOk returns a tuple with the TruncationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationCount

`func (o *V1alpha2Table) SetTruncationCount(v int32)`

SetTruncationCount sets TruncationCount field to given value.

### HasTruncationCount

`func (o *V1alpha2Table) HasTruncationCount() bool`

HasTruncationCount returns a boolean if a field has been set.

### GetType

`func (o *V1alpha2Table) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha2Table) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha2Table) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


