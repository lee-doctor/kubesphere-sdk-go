# V1alpha3BitbucketServerSource

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

### NewV1alpha3BitbucketServerSource

`func NewV1alpha3BitbucketServerSource() *V1alpha3BitbucketServerSource`

NewV1alpha3BitbucketServerSource instantiates a new V1alpha3BitbucketServerSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3BitbucketServerSourceWithDefaults

`func NewV1alpha3BitbucketServerSourceWithDefaults() *V1alpha3BitbucketServerSource`

NewV1alpha3BitbucketServerSourceWithDefaults instantiates a new V1alpha3BitbucketServerSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUri

`func (o *V1alpha3BitbucketServerSource) GetApiUri() string`

GetApiUri returns the ApiUri field if non-nil, zero value otherwise.

### GetApiUriOk

`func (o *V1alpha3BitbucketServerSource) GetApiUriOk() (*string, bool)`

GetApiUriOk returns a tuple with the ApiUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUri

`func (o *V1alpha3BitbucketServerSource) SetApiUri(v string)`

SetApiUri sets ApiUri field to given value.

### HasApiUri

`func (o *V1alpha3BitbucketServerSource) HasApiUri() bool`

HasApiUri returns a boolean if a field has been set.

### GetCredentialId

`func (o *V1alpha3BitbucketServerSource) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *V1alpha3BitbucketServerSource) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *V1alpha3BitbucketServerSource) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *V1alpha3BitbucketServerSource) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetDiscoverBranches

`func (o *V1alpha3BitbucketServerSource) GetDiscoverBranches() int32`

GetDiscoverBranches returns the DiscoverBranches field if non-nil, zero value otherwise.

### GetDiscoverBranchesOk

`func (o *V1alpha3BitbucketServerSource) GetDiscoverBranchesOk() (*int32, bool)`

GetDiscoverBranchesOk returns a tuple with the DiscoverBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverBranches

`func (o *V1alpha3BitbucketServerSource) SetDiscoverBranches(v int32)`

SetDiscoverBranches sets DiscoverBranches field to given value.

### HasDiscoverBranches

`func (o *V1alpha3BitbucketServerSource) HasDiscoverBranches() bool`

HasDiscoverBranches returns a boolean if a field has been set.

### GetDiscoverPrFromForks

`func (o *V1alpha3BitbucketServerSource) GetDiscoverPrFromForks() V1alpha3DiscoverPRFromForks`

GetDiscoverPrFromForks returns the DiscoverPrFromForks field if non-nil, zero value otherwise.

### GetDiscoverPrFromForksOk

`func (o *V1alpha3BitbucketServerSource) GetDiscoverPrFromForksOk() (*V1alpha3DiscoverPRFromForks, bool)`

GetDiscoverPrFromForksOk returns a tuple with the DiscoverPrFromForks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPrFromForks

`func (o *V1alpha3BitbucketServerSource) SetDiscoverPrFromForks(v V1alpha3DiscoverPRFromForks)`

SetDiscoverPrFromForks sets DiscoverPrFromForks field to given value.

### HasDiscoverPrFromForks

`func (o *V1alpha3BitbucketServerSource) HasDiscoverPrFromForks() bool`

HasDiscoverPrFromForks returns a boolean if a field has been set.

### GetDiscoverPrFromOrigin

`func (o *V1alpha3BitbucketServerSource) GetDiscoverPrFromOrigin() int32`

GetDiscoverPrFromOrigin returns the DiscoverPrFromOrigin field if non-nil, zero value otherwise.

### GetDiscoverPrFromOriginOk

`func (o *V1alpha3BitbucketServerSource) GetDiscoverPrFromOriginOk() (*int32, bool)`

GetDiscoverPrFromOriginOk returns a tuple with the DiscoverPrFromOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPrFromOrigin

`func (o *V1alpha3BitbucketServerSource) SetDiscoverPrFromOrigin(v int32)`

SetDiscoverPrFromOrigin sets DiscoverPrFromOrigin field to given value.

### HasDiscoverPrFromOrigin

`func (o *V1alpha3BitbucketServerSource) HasDiscoverPrFromOrigin() bool`

HasDiscoverPrFromOrigin returns a boolean if a field has been set.

### GetDiscoverTags

`func (o *V1alpha3BitbucketServerSource) GetDiscoverTags() bool`

GetDiscoverTags returns the DiscoverTags field if non-nil, zero value otherwise.

### GetDiscoverTagsOk

`func (o *V1alpha3BitbucketServerSource) GetDiscoverTagsOk() (*bool, bool)`

GetDiscoverTagsOk returns a tuple with the DiscoverTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTags

`func (o *V1alpha3BitbucketServerSource) SetDiscoverTags(v bool)`

SetDiscoverTags sets DiscoverTags field to given value.

### HasDiscoverTags

`func (o *V1alpha3BitbucketServerSource) HasDiscoverTags() bool`

HasDiscoverTags returns a boolean if a field has been set.

### GetGitCloneOption

`func (o *V1alpha3BitbucketServerSource) GetGitCloneOption() V1alpha3GitCloneOption`

GetGitCloneOption returns the GitCloneOption field if non-nil, zero value otherwise.

### GetGitCloneOptionOk

`func (o *V1alpha3BitbucketServerSource) GetGitCloneOptionOk() (*V1alpha3GitCloneOption, bool)`

GetGitCloneOptionOk returns a tuple with the GitCloneOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCloneOption

`func (o *V1alpha3BitbucketServerSource) SetGitCloneOption(v V1alpha3GitCloneOption)`

SetGitCloneOption sets GitCloneOption field to given value.

### HasGitCloneOption

`func (o *V1alpha3BitbucketServerSource) HasGitCloneOption() bool`

HasGitCloneOption returns a boolean if a field has been set.

### GetOwner

`func (o *V1alpha3BitbucketServerSource) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1alpha3BitbucketServerSource) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1alpha3BitbucketServerSource) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1alpha3BitbucketServerSource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRegexFilter

`func (o *V1alpha3BitbucketServerSource) GetRegexFilter() string`

GetRegexFilter returns the RegexFilter field if non-nil, zero value otherwise.

### GetRegexFilterOk

`func (o *V1alpha3BitbucketServerSource) GetRegexFilterOk() (*string, bool)`

GetRegexFilterOk returns a tuple with the RegexFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexFilter

`func (o *V1alpha3BitbucketServerSource) SetRegexFilter(v string)`

SetRegexFilter sets RegexFilter field to given value.

### HasRegexFilter

`func (o *V1alpha3BitbucketServerSource) HasRegexFilter() bool`

HasRegexFilter returns a boolean if a field has been set.

### GetRepo

`func (o *V1alpha3BitbucketServerSource) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *V1alpha3BitbucketServerSource) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *V1alpha3BitbucketServerSource) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *V1alpha3BitbucketServerSource) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetScmId

`func (o *V1alpha3BitbucketServerSource) GetScmId() string`

GetScmId returns the ScmId field if non-nil, zero value otherwise.

### GetScmIdOk

`func (o *V1alpha3BitbucketServerSource) GetScmIdOk() (*string, bool)`

GetScmIdOk returns a tuple with the ScmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmId

`func (o *V1alpha3BitbucketServerSource) SetScmId(v string)`

SetScmId sets ScmId field to given value.

### HasScmId

`func (o *V1alpha3BitbucketServerSource) HasScmId() bool`

HasScmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


