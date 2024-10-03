# OpenpitrixCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ClusterType** | Pointer to **int64** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Endpoints** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 
**FrontgateId** | Pointer to **string** |  | [optional] 
**GlobalUuid** | Pointer to **string** |  | [optional] 
**MetadataRootAccess** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**RuntimeId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusTime** | Pointer to **string** |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 
**TransitionStatus** | Pointer to **string** |  | [optional] 
**UpgradeStatus** | Pointer to **string** |  | [optional] 
**UpgradeTime** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenpitrixCluster

`func NewOpenpitrixCluster() *OpenpitrixCluster`

NewOpenpitrixCluster instantiates a new OpenpitrixCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenpitrixClusterWithDefaults

`func NewOpenpitrixClusterWithDefaults() *OpenpitrixCluster`

NewOpenpitrixClusterWithDefaults instantiates a new OpenpitrixCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *OpenpitrixCluster) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *OpenpitrixCluster) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *OpenpitrixCluster) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *OpenpitrixCluster) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAppId

`func (o *OpenpitrixCluster) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *OpenpitrixCluster) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *OpenpitrixCluster) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *OpenpitrixCluster) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetClusterId

`func (o *OpenpitrixCluster) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *OpenpitrixCluster) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *OpenpitrixCluster) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *OpenpitrixCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterType

`func (o *OpenpitrixCluster) GetClusterType() int64`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *OpenpitrixCluster) GetClusterTypeOk() (*int64, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *OpenpitrixCluster) SetClusterType(v int64)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *OpenpitrixCluster) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetCreateTime

`func (o *OpenpitrixCluster) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OpenpitrixCluster) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OpenpitrixCluster) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OpenpitrixCluster) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDebug

`func (o *OpenpitrixCluster) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *OpenpitrixCluster) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *OpenpitrixCluster) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *OpenpitrixCluster) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDescription

`func (o *OpenpitrixCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenpitrixCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenpitrixCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenpitrixCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoints

`func (o *OpenpitrixCluster) GetEndpoints() string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *OpenpitrixCluster) GetEndpointsOk() (*string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *OpenpitrixCluster) SetEndpoints(v string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *OpenpitrixCluster) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEnv

`func (o *OpenpitrixCluster) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *OpenpitrixCluster) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *OpenpitrixCluster) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *OpenpitrixCluster) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetFrontgateId

`func (o *OpenpitrixCluster) GetFrontgateId() string`

GetFrontgateId returns the FrontgateId field if non-nil, zero value otherwise.

### GetFrontgateIdOk

`func (o *OpenpitrixCluster) GetFrontgateIdOk() (*string, bool)`

GetFrontgateIdOk returns a tuple with the FrontgateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontgateId

`func (o *OpenpitrixCluster) SetFrontgateId(v string)`

SetFrontgateId sets FrontgateId field to given value.

### HasFrontgateId

`func (o *OpenpitrixCluster) HasFrontgateId() bool`

HasFrontgateId returns a boolean if a field has been set.

### GetGlobalUuid

`func (o *OpenpitrixCluster) GetGlobalUuid() string`

GetGlobalUuid returns the GlobalUuid field if non-nil, zero value otherwise.

### GetGlobalUuidOk

`func (o *OpenpitrixCluster) GetGlobalUuidOk() (*string, bool)`

GetGlobalUuidOk returns a tuple with the GlobalUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalUuid

`func (o *OpenpitrixCluster) SetGlobalUuid(v string)`

SetGlobalUuid sets GlobalUuid field to given value.

### HasGlobalUuid

`func (o *OpenpitrixCluster) HasGlobalUuid() bool`

HasGlobalUuid returns a boolean if a field has been set.

### GetMetadataRootAccess

`func (o *OpenpitrixCluster) GetMetadataRootAccess() bool`

GetMetadataRootAccess returns the MetadataRootAccess field if non-nil, zero value otherwise.

### GetMetadataRootAccessOk

`func (o *OpenpitrixCluster) GetMetadataRootAccessOk() (*bool, bool)`

GetMetadataRootAccessOk returns a tuple with the MetadataRootAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRootAccess

`func (o *OpenpitrixCluster) SetMetadataRootAccess(v bool)`

SetMetadataRootAccess sets MetadataRootAccess field to given value.

### HasMetadataRootAccess

`func (o *OpenpitrixCluster) HasMetadataRootAccess() bool`

HasMetadataRootAccess returns a boolean if a field has been set.

### GetName

`func (o *OpenpitrixCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenpitrixCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenpitrixCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenpitrixCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *OpenpitrixCluster) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OpenpitrixCluster) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OpenpitrixCluster) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *OpenpitrixCluster) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRuntimeId

`func (o *OpenpitrixCluster) GetRuntimeId() string`

GetRuntimeId returns the RuntimeId field if non-nil, zero value otherwise.

### GetRuntimeIdOk

`func (o *OpenpitrixCluster) GetRuntimeIdOk() (*string, bool)`

GetRuntimeIdOk returns a tuple with the RuntimeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeId

`func (o *OpenpitrixCluster) SetRuntimeId(v string)`

SetRuntimeId sets RuntimeId field to given value.

### HasRuntimeId

`func (o *OpenpitrixCluster) HasRuntimeId() bool`

HasRuntimeId returns a boolean if a field has been set.

### GetStatus

`func (o *OpenpitrixCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenpitrixCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenpitrixCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenpitrixCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *OpenpitrixCluster) GetStatusTime() string`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *OpenpitrixCluster) GetStatusTimeOk() (*string, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *OpenpitrixCluster) SetStatusTime(v string)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *OpenpitrixCluster) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetSubnetId

`func (o *OpenpitrixCluster) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *OpenpitrixCluster) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *OpenpitrixCluster) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *OpenpitrixCluster) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTransitionStatus

`func (o *OpenpitrixCluster) GetTransitionStatus() string`

GetTransitionStatus returns the TransitionStatus field if non-nil, zero value otherwise.

### GetTransitionStatusOk

`func (o *OpenpitrixCluster) GetTransitionStatusOk() (*string, bool)`

GetTransitionStatusOk returns a tuple with the TransitionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionStatus

`func (o *OpenpitrixCluster) SetTransitionStatus(v string)`

SetTransitionStatus sets TransitionStatus field to given value.

### HasTransitionStatus

`func (o *OpenpitrixCluster) HasTransitionStatus() bool`

HasTransitionStatus returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *OpenpitrixCluster) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *OpenpitrixCluster) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *OpenpitrixCluster) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *OpenpitrixCluster) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### GetUpgradeTime

`func (o *OpenpitrixCluster) GetUpgradeTime() string`

GetUpgradeTime returns the UpgradeTime field if non-nil, zero value otherwise.

### GetUpgradeTimeOk

`func (o *OpenpitrixCluster) GetUpgradeTimeOk() (*string, bool)`

GetUpgradeTimeOk returns a tuple with the UpgradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTime

`func (o *OpenpitrixCluster) SetUpgradeTime(v string)`

SetUpgradeTime sets UpgradeTime field to given value.

### HasUpgradeTime

`func (o *OpenpitrixCluster) HasUpgradeTime() bool`

HasUpgradeTime returns a boolean if a field has been set.

### GetVersionId

`func (o *OpenpitrixCluster) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *OpenpitrixCluster) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *OpenpitrixCluster) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *OpenpitrixCluster) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVpcId

`func (o *OpenpitrixCluster) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *OpenpitrixCluster) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *OpenpitrixCluster) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *OpenpitrixCluster) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetZone

`func (o *OpenpitrixCluster) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *OpenpitrixCluster) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *OpenpitrixCluster) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *OpenpitrixCluster) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


