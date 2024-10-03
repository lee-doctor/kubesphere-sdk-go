# V1alpha2GroupBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**GroupRef** | Pointer to [**V1alpha2GroupRef**](V1alpha2GroupRef.md) |  | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1alpha2GroupBinding

`func NewV1alpha2GroupBinding() *V1alpha2GroupBinding`

NewV1alpha2GroupBinding instantiates a new V1alpha2GroupBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2GroupBindingWithDefaults

`func NewV1alpha2GroupBindingWithDefaults() *V1alpha2GroupBinding`

NewV1alpha2GroupBindingWithDefaults instantiates a new V1alpha2GroupBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1alpha2GroupBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1alpha2GroupBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1alpha2GroupBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1alpha2GroupBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetGroupRef

`func (o *V1alpha2GroupBinding) GetGroupRef() V1alpha2GroupRef`

GetGroupRef returns the GroupRef field if non-nil, zero value otherwise.

### GetGroupRefOk

`func (o *V1alpha2GroupBinding) GetGroupRefOk() (*V1alpha2GroupRef, bool)`

GetGroupRefOk returns a tuple with the GroupRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRef

`func (o *V1alpha2GroupBinding) SetGroupRef(v V1alpha2GroupRef)`

SetGroupRef sets GroupRef field to given value.

### HasGroupRef

`func (o *V1alpha2GroupBinding) HasGroupRef() bool`

HasGroupRef returns a boolean if a field has been set.

### GetKind

`func (o *V1alpha2GroupBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1alpha2GroupBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1alpha2GroupBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1alpha2GroupBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1alpha2GroupBinding) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha2GroupBinding) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha2GroupBinding) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha2GroupBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsers

`func (o *V1alpha2GroupBinding) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V1alpha2GroupBinding) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V1alpha2GroupBinding) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V1alpha2GroupBinding) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


