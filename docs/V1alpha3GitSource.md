# V1alpha3GitSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | Pointer to **string** | credential id to access git source | [optional] 
**DiscoverBranches** | Pointer to **bool** | Whether to discover a branch | [optional] 
**DiscoverTags** | Pointer to **bool** | Discover tags configuration | [optional] 
**GitCloneOption** | Pointer to [**V1alpha3GitCloneOption**](V1alpha3GitCloneOption.md) |  | [optional] 
**RegexFilter** | Pointer to **string** | Regex used to match the name of the branch that needs to be run | [optional] 
**ScmId** | Pointer to **string** | uid of scm | [optional] 
**Url** | Pointer to **string** | url of git source | [optional] 

## Methods

### NewV1alpha3GitSource

`func NewV1alpha3GitSource() *V1alpha3GitSource`

NewV1alpha3GitSource instantiates a new V1alpha3GitSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3GitSourceWithDefaults

`func NewV1alpha3GitSourceWithDefaults() *V1alpha3GitSource`

NewV1alpha3GitSourceWithDefaults instantiates a new V1alpha3GitSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *V1alpha3GitSource) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *V1alpha3GitSource) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *V1alpha3GitSource) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *V1alpha3GitSource) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetDiscoverBranches

`func (o *V1alpha3GitSource) GetDiscoverBranches() bool`

GetDiscoverBranches returns the DiscoverBranches field if non-nil, zero value otherwise.

### GetDiscoverBranchesOk

`func (o *V1alpha3GitSource) GetDiscoverBranchesOk() (*bool, bool)`

GetDiscoverBranchesOk returns a tuple with the DiscoverBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverBranches

`func (o *V1alpha3GitSource) SetDiscoverBranches(v bool)`

SetDiscoverBranches sets DiscoverBranches field to given value.

### HasDiscoverBranches

`func (o *V1alpha3GitSource) HasDiscoverBranches() bool`

HasDiscoverBranches returns a boolean if a field has been set.

### GetDiscoverTags

`func (o *V1alpha3GitSource) GetDiscoverTags() bool`

GetDiscoverTags returns the DiscoverTags field if non-nil, zero value otherwise.

### GetDiscoverTagsOk

`func (o *V1alpha3GitSource) GetDiscoverTagsOk() (*bool, bool)`

GetDiscoverTagsOk returns a tuple with the DiscoverTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTags

`func (o *V1alpha3GitSource) SetDiscoverTags(v bool)`

SetDiscoverTags sets DiscoverTags field to given value.

### HasDiscoverTags

`func (o *V1alpha3GitSource) HasDiscoverTags() bool`

HasDiscoverTags returns a boolean if a field has been set.

### GetGitCloneOption

`func (o *V1alpha3GitSource) GetGitCloneOption() V1alpha3GitCloneOption`

GetGitCloneOption returns the GitCloneOption field if non-nil, zero value otherwise.

### GetGitCloneOptionOk

`func (o *V1alpha3GitSource) GetGitCloneOptionOk() (*V1alpha3GitCloneOption, bool)`

GetGitCloneOptionOk returns a tuple with the GitCloneOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCloneOption

`func (o *V1alpha3GitSource) SetGitCloneOption(v V1alpha3GitCloneOption)`

SetGitCloneOption sets GitCloneOption field to given value.

### HasGitCloneOption

`func (o *V1alpha3GitSource) HasGitCloneOption() bool`

HasGitCloneOption returns a boolean if a field has been set.

### GetRegexFilter

`func (o *V1alpha3GitSource) GetRegexFilter() string`

GetRegexFilter returns the RegexFilter field if non-nil, zero value otherwise.

### GetRegexFilterOk

`func (o *V1alpha3GitSource) GetRegexFilterOk() (*string, bool)`

GetRegexFilterOk returns a tuple with the RegexFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexFilter

`func (o *V1alpha3GitSource) SetRegexFilter(v string)`

SetRegexFilter sets RegexFilter field to given value.

### HasRegexFilter

`func (o *V1alpha3GitSource) HasRegexFilter() bool`

HasRegexFilter returns a boolean if a field has been set.

### GetScmId

`func (o *V1alpha3GitSource) GetScmId() string`

GetScmId returns the ScmId field if non-nil, zero value otherwise.

### GetScmIdOk

`func (o *V1alpha3GitSource) GetScmIdOk() (*string, bool)`

GetScmIdOk returns a tuple with the ScmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmId

`func (o *V1alpha3GitSource) SetScmId(v string)`

SetScmId sets ScmId field to given value.

### HasScmId

`func (o *V1alpha3GitSource) HasScmId() bool`

HasScmId returns a boolean if a field has been set.

### GetUrl

`func (o *V1alpha3GitSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1alpha3GitSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1alpha3GitSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1alpha3GitSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


