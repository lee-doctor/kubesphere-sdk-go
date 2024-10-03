# V1alpha3SvnSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | Pointer to **string** | credential id to access svn source | [optional] 
**Excludes** | Pointer to **string** | branches do not run pipeline | [optional] 
**Includes** | Pointer to **string** | branches to run pipeline | [optional] 
**Remote** | Pointer to **string** | remote address url | [optional] 
**ScmId** | Pointer to **string** | uid of scm | [optional] 

## Methods

### NewV1alpha3SvnSource

`func NewV1alpha3SvnSource() *V1alpha3SvnSource`

NewV1alpha3SvnSource instantiates a new V1alpha3SvnSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3SvnSourceWithDefaults

`func NewV1alpha3SvnSourceWithDefaults() *V1alpha3SvnSource`

NewV1alpha3SvnSourceWithDefaults instantiates a new V1alpha3SvnSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *V1alpha3SvnSource) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *V1alpha3SvnSource) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *V1alpha3SvnSource) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *V1alpha3SvnSource) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetExcludes

`func (o *V1alpha3SvnSource) GetExcludes() string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *V1alpha3SvnSource) GetExcludesOk() (*string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *V1alpha3SvnSource) SetExcludes(v string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *V1alpha3SvnSource) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### GetIncludes

`func (o *V1alpha3SvnSource) GetIncludes() string`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *V1alpha3SvnSource) GetIncludesOk() (*string, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *V1alpha3SvnSource) SetIncludes(v string)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *V1alpha3SvnSource) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetRemote

`func (o *V1alpha3SvnSource) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *V1alpha3SvnSource) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *V1alpha3SvnSource) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *V1alpha3SvnSource) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetScmId

`func (o *V1alpha3SvnSource) GetScmId() string`

GetScmId returns the ScmId field if non-nil, zero value otherwise.

### GetScmIdOk

`func (o *V1alpha3SvnSource) GetScmIdOk() (*string, bool)`

GetScmIdOk returns a tuple with the ScmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmId

`func (o *V1alpha3SvnSource) SetScmId(v string)`

SetScmId sets ScmId field to given value.

### HasScmId

`func (o *V1alpha3SvnSource) HasScmId() bool`

HasScmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


