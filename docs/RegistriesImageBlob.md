# RegistriesImageBlob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | The architecture field specifies the CPU architecture, for example amd64 or ppc64le. | [optional] 
**Config** | Pointer to [**RegistriesConfig**](RegistriesConfig.md) |  | [optional] 
**Container** | Pointer to **string** | Container id. | [optional] 
**ContainerConfig** | Pointer to [**RegistriesContainerConfig**](RegistriesContainerConfig.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Create time. | [optional] 
**DockerVersion** | Pointer to **string** | docker version. | [optional] 
**History** | Pointer to [**[]RegistriesHistory**](RegistriesHistory.md) | The data of history update. | [optional] 
**Os** | Pointer to **string** | Operating system. | [optional] 
**RootfsOmitempty** | [**RegistriesRootfs**](RegistriesRootfs.md) |  | 

## Methods

### NewRegistriesImageBlob

`func NewRegistriesImageBlob(rootfsOmitempty RegistriesRootfs, ) *RegistriesImageBlob`

NewRegistriesImageBlob instantiates a new RegistriesImageBlob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesImageBlobWithDefaults

`func NewRegistriesImageBlobWithDefaults() *RegistriesImageBlob`

NewRegistriesImageBlobWithDefaults instantiates a new RegistriesImageBlob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *RegistriesImageBlob) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *RegistriesImageBlob) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *RegistriesImageBlob) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *RegistriesImageBlob) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetConfig

`func (o *RegistriesImageBlob) GetConfig() RegistriesConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RegistriesImageBlob) GetConfigOk() (*RegistriesConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RegistriesImageBlob) SetConfig(v RegistriesConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RegistriesImageBlob) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainer

`func (o *RegistriesImageBlob) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *RegistriesImageBlob) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *RegistriesImageBlob) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *RegistriesImageBlob) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerConfig

`func (o *RegistriesImageBlob) GetContainerConfig() RegistriesContainerConfig`

GetContainerConfig returns the ContainerConfig field if non-nil, zero value otherwise.

### GetContainerConfigOk

`func (o *RegistriesImageBlob) GetContainerConfigOk() (*RegistriesContainerConfig, bool)`

GetContainerConfigOk returns a tuple with the ContainerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerConfig

`func (o *RegistriesImageBlob) SetContainerConfig(v RegistriesContainerConfig)`

SetContainerConfig sets ContainerConfig field to given value.

### HasContainerConfig

`func (o *RegistriesImageBlob) HasContainerConfig() bool`

HasContainerConfig returns a boolean if a field has been set.

### GetCreated

`func (o *RegistriesImageBlob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RegistriesImageBlob) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RegistriesImageBlob) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RegistriesImageBlob) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDockerVersion

`func (o *RegistriesImageBlob) GetDockerVersion() string`

GetDockerVersion returns the DockerVersion field if non-nil, zero value otherwise.

### GetDockerVersionOk

`func (o *RegistriesImageBlob) GetDockerVersionOk() (*string, bool)`

GetDockerVersionOk returns a tuple with the DockerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerVersion

`func (o *RegistriesImageBlob) SetDockerVersion(v string)`

SetDockerVersion sets DockerVersion field to given value.

### HasDockerVersion

`func (o *RegistriesImageBlob) HasDockerVersion() bool`

HasDockerVersion returns a boolean if a field has been set.

### GetHistory

`func (o *RegistriesImageBlob) GetHistory() []RegistriesHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *RegistriesImageBlob) GetHistoryOk() (*[]RegistriesHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *RegistriesImageBlob) SetHistory(v []RegistriesHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *RegistriesImageBlob) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetOs

`func (o *RegistriesImageBlob) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *RegistriesImageBlob) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *RegistriesImageBlob) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *RegistriesImageBlob) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetRootfsOmitempty

`func (o *RegistriesImageBlob) GetRootfsOmitempty() RegistriesRootfs`

GetRootfsOmitempty returns the RootfsOmitempty field if non-nil, zero value otherwise.

### GetRootfsOmitemptyOk

`func (o *RegistriesImageBlob) GetRootfsOmitemptyOk() (*RegistriesRootfs, bool)`

GetRootfsOmitemptyOk returns a tuple with the RootfsOmitempty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfsOmitempty

`func (o *RegistriesImageBlob) SetRootfsOmitempty(v RegistriesRootfs)`

SetRootfsOmitempty sets RootfsOmitempty field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


