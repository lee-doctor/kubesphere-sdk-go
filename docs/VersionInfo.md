# VersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildDate** | **string** |  | 
**Compiler** | **string** |  | 
**GitCommit** | **string** |  | 
**GitMajor** | **string** |  | 
**GitMinor** | **string** |  | 
**GitTreeState** | **string** |  | 
**GitVersion** | **string** |  | 
**GoVersion** | **string** |  | 
**Kubernetes** | Pointer to [**VersionInfo**](VersionInfo.md) |  | [optional] 
**Platform** | **string** |  | 

## Methods

### NewVersionInfo

`func NewVersionInfo(buildDate string, compiler string, gitCommit string, gitMajor string, gitMinor string, gitTreeState string, gitVersion string, goVersion string, platform string, ) *VersionInfo`

NewVersionInfo instantiates a new VersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoWithDefaults

`func NewVersionInfoWithDefaults() *VersionInfo`

NewVersionInfoWithDefaults instantiates a new VersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildDate

`func (o *VersionInfo) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *VersionInfo) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *VersionInfo) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.


### GetCompiler

`func (o *VersionInfo) GetCompiler() string`

GetCompiler returns the Compiler field if non-nil, zero value otherwise.

### GetCompilerOk

`func (o *VersionInfo) GetCompilerOk() (*string, bool)`

GetCompilerOk returns a tuple with the Compiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiler

`func (o *VersionInfo) SetCompiler(v string)`

SetCompiler sets Compiler field to given value.


### GetGitCommit

`func (o *VersionInfo) GetGitCommit() string`

GetGitCommit returns the GitCommit field if non-nil, zero value otherwise.

### GetGitCommitOk

`func (o *VersionInfo) GetGitCommitOk() (*string, bool)`

GetGitCommitOk returns a tuple with the GitCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommit

`func (o *VersionInfo) SetGitCommit(v string)`

SetGitCommit sets GitCommit field to given value.


### GetGitMajor

`func (o *VersionInfo) GetGitMajor() string`

GetGitMajor returns the GitMajor field if non-nil, zero value otherwise.

### GetGitMajorOk

`func (o *VersionInfo) GetGitMajorOk() (*string, bool)`

GetGitMajorOk returns a tuple with the GitMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitMajor

`func (o *VersionInfo) SetGitMajor(v string)`

SetGitMajor sets GitMajor field to given value.


### GetGitMinor

`func (o *VersionInfo) GetGitMinor() string`

GetGitMinor returns the GitMinor field if non-nil, zero value otherwise.

### GetGitMinorOk

`func (o *VersionInfo) GetGitMinorOk() (*string, bool)`

GetGitMinorOk returns a tuple with the GitMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitMinor

`func (o *VersionInfo) SetGitMinor(v string)`

SetGitMinor sets GitMinor field to given value.


### GetGitTreeState

`func (o *VersionInfo) GetGitTreeState() string`

GetGitTreeState returns the GitTreeState field if non-nil, zero value otherwise.

### GetGitTreeStateOk

`func (o *VersionInfo) GetGitTreeStateOk() (*string, bool)`

GetGitTreeStateOk returns a tuple with the GitTreeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTreeState

`func (o *VersionInfo) SetGitTreeState(v string)`

SetGitTreeState sets GitTreeState field to given value.


### GetGitVersion

`func (o *VersionInfo) GetGitVersion() string`

GetGitVersion returns the GitVersion field if non-nil, zero value otherwise.

### GetGitVersionOk

`func (o *VersionInfo) GetGitVersionOk() (*string, bool)`

GetGitVersionOk returns a tuple with the GitVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitVersion

`func (o *VersionInfo) SetGitVersion(v string)`

SetGitVersion sets GitVersion field to given value.


### GetGoVersion

`func (o *VersionInfo) GetGoVersion() string`

GetGoVersion returns the GoVersion field if non-nil, zero value otherwise.

### GetGoVersionOk

`func (o *VersionInfo) GetGoVersionOk() (*string, bool)`

GetGoVersionOk returns a tuple with the GoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoVersion

`func (o *VersionInfo) SetGoVersion(v string)`

SetGoVersion sets GoVersion field to given value.


### GetKubernetes

`func (o *VersionInfo) GetKubernetes() VersionInfo`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *VersionInfo) GetKubernetesOk() (*VersionInfo, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *VersionInfo) SetKubernetes(v VersionInfo)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *VersionInfo) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetPlatform

`func (o *VersionInfo) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VersionInfo) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VersionInfo) SetPlatform(v string)`

SetPlatform sets Platform field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


