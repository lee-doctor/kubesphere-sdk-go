# RegistriesHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Created time. | [optional] 
**CreatedBy** | Pointer to **string** | Created command. | [optional] 
**EmptyLayer** | Pointer to **bool** | Layer empty or not. | [optional] 

## Methods

### NewRegistriesHistory

`func NewRegistriesHistory() *RegistriesHistory`

NewRegistriesHistory instantiates a new RegistriesHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesHistoryWithDefaults

`func NewRegistriesHistoryWithDefaults() *RegistriesHistory`

NewRegistriesHistoryWithDefaults instantiates a new RegistriesHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *RegistriesHistory) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RegistriesHistory) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RegistriesHistory) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RegistriesHistory) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RegistriesHistory) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RegistriesHistory) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RegistriesHistory) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RegistriesHistory) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEmptyLayer

`func (o *RegistriesHistory) GetEmptyLayer() bool`

GetEmptyLayer returns the EmptyLayer field if non-nil, zero value otherwise.

### GetEmptyLayerOk

`func (o *RegistriesHistory) GetEmptyLayerOk() (*bool, bool)`

GetEmptyLayerOk returns a tuple with the EmptyLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyLayer

`func (o *RegistriesHistory) SetEmptyLayer(v bool)`

SetEmptyLayer sets EmptyLayer field to given value.

### HasEmptyLayer

`func (o *RegistriesHistory) HasEmptyLayer() bool`

HasEmptyLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


