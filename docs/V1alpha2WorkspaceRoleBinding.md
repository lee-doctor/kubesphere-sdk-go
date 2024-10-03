# V1alpha2WorkspaceRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**RoleRef** | [**V1RoleRef**](V1RoleRef.md) |  | 
**Subjects** | Pointer to [**[]V1Subject**](V1Subject.md) |  | [optional] 

## Methods

### NewV1alpha2WorkspaceRoleBinding

`func NewV1alpha2WorkspaceRoleBinding(roleRef V1RoleRef, ) *V1alpha2WorkspaceRoleBinding`

NewV1alpha2WorkspaceRoleBinding instantiates a new V1alpha2WorkspaceRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2WorkspaceRoleBindingWithDefaults

`func NewV1alpha2WorkspaceRoleBindingWithDefaults() *V1alpha2WorkspaceRoleBinding`

NewV1alpha2WorkspaceRoleBindingWithDefaults instantiates a new V1alpha2WorkspaceRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1alpha2WorkspaceRoleBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1alpha2WorkspaceRoleBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1alpha2WorkspaceRoleBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1alpha2WorkspaceRoleBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1alpha2WorkspaceRoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1alpha2WorkspaceRoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1alpha2WorkspaceRoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1alpha2WorkspaceRoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1alpha2WorkspaceRoleBinding) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha2WorkspaceRoleBinding) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha2WorkspaceRoleBinding) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha2WorkspaceRoleBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRoleRef

`func (o *V1alpha2WorkspaceRoleBinding) GetRoleRef() V1RoleRef`

GetRoleRef returns the RoleRef field if non-nil, zero value otherwise.

### GetRoleRefOk

`func (o *V1alpha2WorkspaceRoleBinding) GetRoleRefOk() (*V1RoleRef, bool)`

GetRoleRefOk returns a tuple with the RoleRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRef

`func (o *V1alpha2WorkspaceRoleBinding) SetRoleRef(v V1RoleRef)`

SetRoleRef sets RoleRef field to given value.


### GetSubjects

`func (o *V1alpha2WorkspaceRoleBinding) GetSubjects() []V1Subject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *V1alpha2WorkspaceRoleBinding) GetSubjectsOk() (*[]V1Subject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *V1alpha2WorkspaceRoleBinding) SetSubjects(v []V1Subject)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *V1alpha2WorkspaceRoleBinding) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


