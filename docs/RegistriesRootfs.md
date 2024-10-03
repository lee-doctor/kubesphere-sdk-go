# RegistriesRootfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffIds** | Pointer to **[]string** | Contain ids of layer list | [optional] 
**Type** | Pointer to **string** | Root filesystem type, always \&quot;layers\&quot;  | [optional] 

## Methods

### NewRegistriesRootfs

`func NewRegistriesRootfs() *RegistriesRootfs`

NewRegistriesRootfs instantiates a new RegistriesRootfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesRootfsWithDefaults

`func NewRegistriesRootfsWithDefaults() *RegistriesRootfs`

NewRegistriesRootfsWithDefaults instantiates a new RegistriesRootfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffIds

`func (o *RegistriesRootfs) GetDiffIds() []string`

GetDiffIds returns the DiffIds field if non-nil, zero value otherwise.

### GetDiffIdsOk

`func (o *RegistriesRootfs) GetDiffIdsOk() (*[]string, bool)`

GetDiffIdsOk returns a tuple with the DiffIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffIds

`func (o *RegistriesRootfs) SetDiffIds(v []string)`

SetDiffIds sets DiffIds field to given value.

### HasDiffIds

`func (o *RegistriesRootfs) HasDiffIds() bool`

HasDiffIds returns a boolean if a field has been set.

### GetType

`func (o *RegistriesRootfs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistriesRootfs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistriesRootfs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistriesRootfs) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


