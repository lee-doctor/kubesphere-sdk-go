# DevopsCheckPlayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abort** | Pointer to **bool** | abort or not | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Parameters** | Pointer to [**[]DevopsCheckPlayloadParameters**](DevopsCheckPlayloadParameters.md) |  | [optional] 

## Methods

### NewDevopsCheckPlayload

`func NewDevopsCheckPlayload() *DevopsCheckPlayload`

NewDevopsCheckPlayload instantiates a new DevopsCheckPlayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsCheckPlayloadWithDefaults

`func NewDevopsCheckPlayloadWithDefaults() *DevopsCheckPlayload`

NewDevopsCheckPlayloadWithDefaults instantiates a new DevopsCheckPlayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbort

`func (o *DevopsCheckPlayload) GetAbort() bool`

GetAbort returns the Abort field if non-nil, zero value otherwise.

### GetAbortOk

`func (o *DevopsCheckPlayload) GetAbortOk() (*bool, bool)`

GetAbortOk returns a tuple with the Abort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbort

`func (o *DevopsCheckPlayload) SetAbort(v bool)`

SetAbort sets Abort field to given value.

### HasAbort

`func (o *DevopsCheckPlayload) HasAbort() bool`

HasAbort returns a boolean if a field has been set.

### GetId

`func (o *DevopsCheckPlayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsCheckPlayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsCheckPlayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsCheckPlayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameters

`func (o *DevopsCheckPlayload) GetParameters() []DevopsCheckPlayloadParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DevopsCheckPlayload) GetParametersOk() (*[]DevopsCheckPlayloadParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DevopsCheckPlayload) SetParameters(v []DevopsCheckPlayloadParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DevopsCheckPlayload) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


