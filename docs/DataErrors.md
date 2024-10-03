# DataErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | error message | [optional] 
**Location** | Pointer to **[]string** | err location | [optional] 

## Methods

### NewDataErrors

`func NewDataErrors() *DataErrors`

NewDataErrors instantiates a new DataErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataErrorsWithDefaults

`func NewDataErrorsWithDefaults() *DataErrors`

NewDataErrorsWithDefaults instantiates a new DataErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DataErrors) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DataErrors) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DataErrors) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DataErrors) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLocation

`func (o *DataErrors) GetLocation() []string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DataErrors) GetLocationOk() (*[]string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DataErrors) SetLocation(v []string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DataErrors) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


