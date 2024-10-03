# GitAuthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteUrl** | **string** | git server url | 
**SecretRef** | Pointer to [**V1SecretReference**](V1SecretReference.md) |  | [optional] 

## Methods

### NewGitAuthInfo

`func NewGitAuthInfo(remoteUrl string, ) *GitAuthInfo`

NewGitAuthInfo instantiates a new GitAuthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitAuthInfoWithDefaults

`func NewGitAuthInfoWithDefaults() *GitAuthInfo`

NewGitAuthInfoWithDefaults instantiates a new GitAuthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteUrl

`func (o *GitAuthInfo) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *GitAuthInfo) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *GitAuthInfo) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetSecretRef

`func (o *GitAuthInfo) GetSecretRef() V1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *GitAuthInfo) GetSecretRefOk() (*V1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *GitAuthInfo) SetSecretRef(v V1SecretReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *GitAuthInfo) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


