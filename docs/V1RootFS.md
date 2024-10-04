# V1RootFS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffIds** | [**[]V1Hash**](V1Hash.md) |  | 
**Type** | **string** |  | 

## Methods

### NewV1RootFS

`func NewV1RootFS(diffIds []V1Hash, type_ string, ) *V1RootFS`

NewV1RootFS instantiates a new V1RootFS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RootFSWithDefaults

`func NewV1RootFSWithDefaults() *V1RootFS`

NewV1RootFSWithDefaults instantiates a new V1RootFS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffIds

`func (o *V1RootFS) GetDiffIds() []V1Hash`

GetDiffIds returns the DiffIds field if non-nil, zero value otherwise.

### GetDiffIdsOk

`func (o *V1RootFS) GetDiffIdsOk() (*[]V1Hash, bool)`

GetDiffIdsOk returns a tuple with the DiffIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffIds

`func (o *V1RootFS) SetDiffIds(v []V1Hash)`

SetDiffIds sets DiffIds field to given value.


### GetType

`func (o *V1RootFS) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1RootFS) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1RootFS) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


