# V1alpha3GithubSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUri** | Pointer to **string** | The api url can specify the location of the github apiserver.For private cloud configuration | [optional] 
**CredentialId** | Pointer to **string** | credential id to access github source | [optional] 
**DiscoverBranches** | Pointer to **int32** | Discover branch configuration | [optional] 
**DiscoverPrFromForks** | Pointer to [**V1alpha3DiscoverPRFromForks**](V1alpha3DiscoverPRFromForks.md) |  | [optional] 
**DiscoverPrFromOrigin** | Pointer to **int32** | Discover origin PR configuration | [optional] 
**DiscoverTags** | Pointer to **bool** | Discover tag configuration | [optional] 
**GitCloneOption** | Pointer to [**V1alpha3GitCloneOption**](V1alpha3GitCloneOption.md) |  | [optional] 
**Owner** | Pointer to **string** | owner of github repo | [optional] 
**RegexFilter** | Pointer to **string** | Regex used to match the name of the branch that needs to be run | [optional] 
**Repo** | Pointer to **string** | repo name of github repo | [optional] 
**ScmId** | Pointer to **string** | uid of scm | [optional] 

## Methods

### NewV1alpha3GithubSource

`func NewV1alpha3GithubSource() *V1alpha3GithubSource`

NewV1alpha3GithubSource instantiates a new V1alpha3GithubSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3GithubSourceWithDefaults

`func NewV1alpha3GithubSourceWithDefaults() *V1alpha3GithubSource`

NewV1alpha3GithubSourceWithDefaults instantiates a new V1alpha3GithubSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUri

`func (o *V1alpha3GithubSource) GetApiUri() string`

GetApiUri returns the ApiUri field if non-nil, zero value otherwise.

### GetApiUriOk

`func (o *V1alpha3GithubSource) GetApiUriOk() (*string, bool)`

GetApiUriOk returns a tuple with the ApiUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUri

`func (o *V1alpha3GithubSource) SetApiUri(v string)`

SetApiUri sets ApiUri field to given value.

### HasApiUri

`func (o *V1alpha3GithubSource) HasApiUri() bool`

HasApiUri returns a boolean if a field has been set.

### GetCredentialId

`func (o *V1alpha3GithubSource) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *V1alpha3GithubSource) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *V1alpha3GithubSource) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *V1alpha3GithubSource) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetDiscoverBranches

`func (o *V1alpha3GithubSource) GetDiscoverBranches() int32`

GetDiscoverBranches returns the DiscoverBranches field if non-nil, zero value otherwise.

### GetDiscoverBranchesOk

`func (o *V1alpha3GithubSource) GetDiscoverBranchesOk() (*int32, bool)`

GetDiscoverBranchesOk returns a tuple with the DiscoverBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverBranches

`func (o *V1alpha3GithubSource) SetDiscoverBranches(v int32)`

SetDiscoverBranches sets DiscoverBranches field to given value.

### HasDiscoverBranches

`func (o *V1alpha3GithubSource) HasDiscoverBranches() bool`

HasDiscoverBranches returns a boolean if a field has been set.

### GetDiscoverPrFromForks

`func (o *V1alpha3GithubSource) GetDiscoverPrFromForks() V1alpha3DiscoverPRFromForks`

GetDiscoverPrFromForks returns the DiscoverPrFromForks field if non-nil, zero value otherwise.

### GetDiscoverPrFromForksOk

`func (o *V1alpha3GithubSource) GetDiscoverPrFromForksOk() (*V1alpha3DiscoverPRFromForks, bool)`

GetDiscoverPrFromForksOk returns a tuple with the DiscoverPrFromForks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPrFromForks

`func (o *V1alpha3GithubSource) SetDiscoverPrFromForks(v V1alpha3DiscoverPRFromForks)`

SetDiscoverPrFromForks sets DiscoverPrFromForks field to given value.

### HasDiscoverPrFromForks

`func (o *V1alpha3GithubSource) HasDiscoverPrFromForks() bool`

HasDiscoverPrFromForks returns a boolean if a field has been set.

### GetDiscoverPrFromOrigin

`func (o *V1alpha3GithubSource) GetDiscoverPrFromOrigin() int32`

GetDiscoverPrFromOrigin returns the DiscoverPrFromOrigin field if non-nil, zero value otherwise.

### GetDiscoverPrFromOriginOk

`func (o *V1alpha3GithubSource) GetDiscoverPrFromOriginOk() (*int32, bool)`

GetDiscoverPrFromOriginOk returns a tuple with the DiscoverPrFromOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPrFromOrigin

`func (o *V1alpha3GithubSource) SetDiscoverPrFromOrigin(v int32)`

SetDiscoverPrFromOrigin sets DiscoverPrFromOrigin field to given value.

### HasDiscoverPrFromOrigin

`func (o *V1alpha3GithubSource) HasDiscoverPrFromOrigin() bool`

HasDiscoverPrFromOrigin returns a boolean if a field has been set.

### GetDiscoverTags

`func (o *V1alpha3GithubSource) GetDiscoverTags() bool`

GetDiscoverTags returns the DiscoverTags field if non-nil, zero value otherwise.

### GetDiscoverTagsOk

`func (o *V1alpha3GithubSource) GetDiscoverTagsOk() (*bool, bool)`

GetDiscoverTagsOk returns a tuple with the DiscoverTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTags

`func (o *V1alpha3GithubSource) SetDiscoverTags(v bool)`

SetDiscoverTags sets DiscoverTags field to given value.

### HasDiscoverTags

`func (o *V1alpha3GithubSource) HasDiscoverTags() bool`

HasDiscoverTags returns a boolean if a field has been set.

### GetGitCloneOption

`func (o *V1alpha3GithubSource) GetGitCloneOption() V1alpha3GitCloneOption`

GetGitCloneOption returns the GitCloneOption field if non-nil, zero value otherwise.

### GetGitCloneOptionOk

`func (o *V1alpha3GithubSource) GetGitCloneOptionOk() (*V1alpha3GitCloneOption, bool)`

GetGitCloneOptionOk returns a tuple with the GitCloneOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCloneOption

`func (o *V1alpha3GithubSource) SetGitCloneOption(v V1alpha3GitCloneOption)`

SetGitCloneOption sets GitCloneOption field to given value.

### HasGitCloneOption

`func (o *V1alpha3GithubSource) HasGitCloneOption() bool`

HasGitCloneOption returns a boolean if a field has been set.

### GetOwner

`func (o *V1alpha3GithubSource) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1alpha3GithubSource) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1alpha3GithubSource) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1alpha3GithubSource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRegexFilter

`func (o *V1alpha3GithubSource) GetRegexFilter() string`

GetRegexFilter returns the RegexFilter field if non-nil, zero value otherwise.

### GetRegexFilterOk

`func (o *V1alpha3GithubSource) GetRegexFilterOk() (*string, bool)`

GetRegexFilterOk returns a tuple with the RegexFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexFilter

`func (o *V1alpha3GithubSource) SetRegexFilter(v string)`

SetRegexFilter sets RegexFilter field to given value.

### HasRegexFilter

`func (o *V1alpha3GithubSource) HasRegexFilter() bool`

HasRegexFilter returns a boolean if a field has been set.

### GetRepo

`func (o *V1alpha3GithubSource) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *V1alpha3GithubSource) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *V1alpha3GithubSource) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *V1alpha3GithubSource) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetScmId

`func (o *V1alpha3GithubSource) GetScmId() string`

GetScmId returns the ScmId field if non-nil, zero value otherwise.

### GetScmIdOk

`func (o *V1alpha3GithubSource) GetScmIdOk() (*string, bool)`

GetScmIdOk returns a tuple with the ScmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmId

`func (o *V1alpha3GithubSource) SetScmId(v string)`

SetScmId sets ScmId field to given value.

### HasScmId

`func (o *V1alpha3GithubSource) HasScmId() bool`

HasScmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


