# V1alpha2Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | **string** |  | 
**LabelMinor** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**[]V1alpha2MetadataRow**](V1alpha2MetadataRow.md) |  | [optional] 
**NodeId** | **string** |  | 

## Methods

### NewV1alpha2Connection

`func NewV1alpha2Connection(id string, label string, nodeId string, ) *V1alpha2Connection`

NewV1alpha2Connection instantiates a new V1alpha2Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2ConnectionWithDefaults

`func NewV1alpha2ConnectionWithDefaults() *V1alpha2Connection`

NewV1alpha2ConnectionWithDefaults instantiates a new V1alpha2Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1alpha2Connection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1alpha2Connection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1alpha2Connection) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *V1alpha2Connection) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1alpha2Connection) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1alpha2Connection) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelMinor

`func (o *V1alpha2Connection) GetLabelMinor() string`

GetLabelMinor returns the LabelMinor field if non-nil, zero value otherwise.

### GetLabelMinorOk

`func (o *V1alpha2Connection) GetLabelMinorOk() (*string, bool)`

GetLabelMinorOk returns a tuple with the LabelMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMinor

`func (o *V1alpha2Connection) SetLabelMinor(v string)`

SetLabelMinor sets LabelMinor field to given value.

### HasLabelMinor

`func (o *V1alpha2Connection) HasLabelMinor() bool`

HasLabelMinor returns a boolean if a field has been set.

### GetMetadata

`func (o *V1alpha2Connection) GetMetadata() []V1alpha2MetadataRow`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha2Connection) GetMetadataOk() (*[]V1alpha2MetadataRow, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha2Connection) SetMetadata(v []V1alpha2MetadataRow)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha2Connection) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNodeId

`func (o *V1alpha2Connection) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *V1alpha2Connection) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *V1alpha2Connection) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


