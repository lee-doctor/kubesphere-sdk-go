# V1ConfigFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** |  | 
**Author** | Pointer to **string** |  | [optional] 
**Config** | [**V1Config**](V1Config.md) |  | 
**Container** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**DockerVersion** | Pointer to **string** |  | [optional] 
**History** | Pointer to [**[]V1History**](V1History.md) |  | [optional] 
**Os** | **string** |  | 
**OsFeatures** | Pointer to **[]string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Rootfs** | [**V1RootFS**](V1RootFS.md) |  | 
**Variant** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ConfigFile

`func NewV1ConfigFile(architecture string, config V1Config, os string, rootfs V1RootFS, ) *V1ConfigFile`

NewV1ConfigFile instantiates a new V1ConfigFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigFileWithDefaults

`func NewV1ConfigFileWithDefaults() *V1ConfigFile`

NewV1ConfigFileWithDefaults instantiates a new V1ConfigFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *V1ConfigFile) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *V1ConfigFile) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *V1ConfigFile) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetAuthor

`func (o *V1ConfigFile) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *V1ConfigFile) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *V1ConfigFile) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *V1ConfigFile) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetConfig

`func (o *V1ConfigFile) GetConfig() V1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1ConfigFile) GetConfigOk() (*V1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1ConfigFile) SetConfig(v V1Config)`

SetConfig sets Config field to given value.


### GetContainer

`func (o *V1ConfigFile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V1ConfigFile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V1ConfigFile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V1ConfigFile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreated

`func (o *V1ConfigFile) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *V1ConfigFile) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *V1ConfigFile) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *V1ConfigFile) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDockerVersion

`func (o *V1ConfigFile) GetDockerVersion() string`

GetDockerVersion returns the DockerVersion field if non-nil, zero value otherwise.

### GetDockerVersionOk

`func (o *V1ConfigFile) GetDockerVersionOk() (*string, bool)`

GetDockerVersionOk returns a tuple with the DockerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerVersion

`func (o *V1ConfigFile) SetDockerVersion(v string)`

SetDockerVersion sets DockerVersion field to given value.

### HasDockerVersion

`func (o *V1ConfigFile) HasDockerVersion() bool`

HasDockerVersion returns a boolean if a field has been set.

### GetHistory

`func (o *V1ConfigFile) GetHistory() []V1History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *V1ConfigFile) GetHistoryOk() (*[]V1History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *V1ConfigFile) SetHistory(v []V1History)`

SetHistory sets History field to given value.

### HasHistory

`func (o *V1ConfigFile) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetOs

`func (o *V1ConfigFile) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *V1ConfigFile) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *V1ConfigFile) SetOs(v string)`

SetOs sets Os field to given value.


### GetOsFeatures

`func (o *V1ConfigFile) GetOsFeatures() []string`

GetOsFeatures returns the OsFeatures field if non-nil, zero value otherwise.

### GetOsFeaturesOk

`func (o *V1ConfigFile) GetOsFeaturesOk() (*[]string, bool)`

GetOsFeaturesOk returns a tuple with the OsFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFeatures

`func (o *V1ConfigFile) SetOsFeatures(v []string)`

SetOsFeatures sets OsFeatures field to given value.

### HasOsFeatures

`func (o *V1ConfigFile) HasOsFeatures() bool`

HasOsFeatures returns a boolean if a field has been set.

### GetOsVersion

`func (o *V1ConfigFile) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *V1ConfigFile) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *V1ConfigFile) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *V1ConfigFile) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetRootfs

`func (o *V1ConfigFile) GetRootfs() V1RootFS`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *V1ConfigFile) GetRootfsOk() (*V1RootFS, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *V1ConfigFile) SetRootfs(v V1RootFS)`

SetRootfs sets Rootfs field to given value.


### GetVariant

`func (o *V1ConfigFile) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *V1ConfigFile) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *V1ConfigFile) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *V1ConfigFile) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


