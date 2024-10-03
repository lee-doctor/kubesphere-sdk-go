# V1alpha3GitlabSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUri** | Pointer to **string** | The api url can specify the location of the gitlab apiserver.For private cloud configuration | [optional] 
**CredentialId** | Pointer to **string** | credential id to access gitlab source | [optional] 
**DiscoverBranches** | Pointer to **int32** | Discover branch configuration | [optional] 
**DiscoverPrFromForks** | Pointer to [**V1alpha3DiscoverPRFromForks**](V1alpha3DiscoverPRFromForks.md) |  | [optional] 
**DiscoverPrFromOrigin** | Pointer to **int32** | Discover origin PR configuration | [optional] 
**DiscoverTags** | Pointer to **bool** | Discover tags configuration | [optional] 
**GitCloneOption** | Pointer to [**V1alpha3GitCloneOption**](V1alpha3GitCloneOption.md) |  | [optional] 
**Owner** | Pointer to **string** | owner of gitlab repo | [optional] 
**RegexFilter** | Pointer to **string** | Regex used to match the name of the branch that needs to be run | [optional] 
**Repo** | Pointer to **string** | repo name of gitlab repo | [optional] 
**ScmId** | Pointer to **string** | uid of scm | [optional] 
**ServerName** | Pointer to **string** | the name of gitlab server which was configured in jenkins | [optional] 

## Methods

### NewV1alpha3GitlabSource

`func NewV1alpha3GitlabSource() *V1alpha3GitlabSource`

NewV1alpha3GitlabSource instantiates a new V1alpha3GitlabSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3GitlabSourceWithDefaults

`func NewV1alpha3GitlabSourceWithDefaults() *V1alpha3GitlabSource`

NewV1alpha3GitlabSourceWithDefaults instantiates a new V1alpha3GitlabSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUri

`func (o *V1alpha3GitlabSource) GetApiUri() string`

GetApiUri returns the ApiUri field if non-nil, zero value otherwise.

### GetApiUriOk

`func (o *V1alpha3GitlabSource) GetApiUriOk() (*string, bool)`

GetApiUriOk returns a tuple with the ApiUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUri

`func (o *V1alpha3GitlabSource) SetApiUri(v string)`

SetApiUri sets ApiUri field to given value.

### HasApiUri

`func (o *V1alpha3GitlabSource) HasApiUri() bool`

HasApiUri returns a boolean if a field has been set.

### GetCredentialId

`func (o *V1alpha3GitlabSource) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *V1alpha3GitlabSource) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *V1alpha3GitlabSource) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *V1alpha3GitlabSource) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetDiscoverBranches

`func (o *V1alpha3GitlabSource) GetDiscoverBranches() int32`

GetDiscoverBranches returns the DiscoverBranches field if non-nil, zero value otherwise.

### GetDiscoverBranchesOk

`func (o *V1alpha3GitlabSource) GetDiscoverBranchesOk() (*int32, bool)`

GetDiscoverBranchesOk returns a tuple with the DiscoverBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverBranches

`func (o *V1alpha3GitlabSource) SetDiscoverBranches(v int32)`

SetDiscoverBranches sets DiscoverBranches field to given value.

### HasDiscoverBranches

`func (o *V1alpha3GitlabSource) HasDiscoverBranches() bool`

HasDiscoverBranches returns a boolean if a field has been set.

### GetDiscoverPrFromForks

`func (o *V1alpha3GitlabSource) GetDiscoverPrFromForks() V1alpha3DiscoverPRFromForks`

GetDiscoverPrFromForks returns the DiscoverPrFromForks field if non-nil, zero value otherwise.

### GetDiscoverPrFromForksOk

`func (o *V1alpha3GitlabSource) GetDiscoverPrFromForksOk() (*V1alpha3DiscoverPRFromForks, bool)`

GetDiscoverPrFromForksOk returns a tuple with the DiscoverPrFromForks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPrFromForks

`func (o *V1alpha3GitlabSource) SetDiscoverPrFromForks(v V1alpha3DiscoverPRFromForks)`

SetDiscoverPrFromForks sets DiscoverPrFromForks field to given value.

### HasDiscoverPrFromForks

`func (o *V1alpha3GitlabSource) HasDiscoverPrFromForks() bool`

HasDiscoverPrFromForks returns a boolean if a field has been set.

### GetDiscoverPrFromOrigin

`func (o *V1alpha3GitlabSource) GetDiscoverPrFromOrigin() int32`

GetDiscoverPrFromOrigin returns the DiscoverPrFromOrigin field if non-nil, zero value otherwise.

### GetDiscoverPrFromOriginOk

`func (o *V1alpha3GitlabSource) GetDiscoverPrFromOriginOk() (*int32, bool)`

GetDiscoverPrFromOriginOk returns a tuple with the DiscoverPrFromOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverPrFromOrigin

`func (o *V1alpha3GitlabSource) SetDiscoverPrFromOrigin(v int32)`

SetDiscoverPrFromOrigin sets DiscoverPrFromOrigin field to given value.

### HasDiscoverPrFromOrigin

`func (o *V1alpha3GitlabSource) HasDiscoverPrFromOrigin() bool`

HasDiscoverPrFromOrigin returns a boolean if a field has been set.

### GetDiscoverTags

`func (o *V1alpha3GitlabSource) GetDiscoverTags() bool`

GetDiscoverTags returns the DiscoverTags field if non-nil, zero value otherwise.

### GetDiscoverTagsOk

`func (o *V1alpha3GitlabSource) GetDiscoverTagsOk() (*bool, bool)`

GetDiscoverTagsOk returns a tuple with the DiscoverTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverTags

`func (o *V1alpha3GitlabSource) SetDiscoverTags(v bool)`

SetDiscoverTags sets DiscoverTags field to given value.

### HasDiscoverTags

`func (o *V1alpha3GitlabSource) HasDiscoverTags() bool`

HasDiscoverTags returns a boolean if a field has been set.

### GetGitCloneOption

`func (o *V1alpha3GitlabSource) GetGitCloneOption() V1alpha3GitCloneOption`

GetGitCloneOption returns the GitCloneOption field if non-nil, zero value otherwise.

### GetGitCloneOptionOk

`func (o *V1alpha3GitlabSource) GetGitCloneOptionOk() (*V1alpha3GitCloneOption, bool)`

GetGitCloneOptionOk returns a tuple with the GitCloneOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCloneOption

`func (o *V1alpha3GitlabSource) SetGitCloneOption(v V1alpha3GitCloneOption)`

SetGitCloneOption sets GitCloneOption field to given value.

### HasGitCloneOption

`func (o *V1alpha3GitlabSource) HasGitCloneOption() bool`

HasGitCloneOption returns a boolean if a field has been set.

### GetOwner

`func (o *V1alpha3GitlabSource) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1alpha3GitlabSource) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1alpha3GitlabSource) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1alpha3GitlabSource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRegexFilter

`func (o *V1alpha3GitlabSource) GetRegexFilter() string`

GetRegexFilter returns the RegexFilter field if non-nil, zero value otherwise.

### GetRegexFilterOk

`func (o *V1alpha3GitlabSource) GetRegexFilterOk() (*string, bool)`

GetRegexFilterOk returns a tuple with the RegexFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexFilter

`func (o *V1alpha3GitlabSource) SetRegexFilter(v string)`

SetRegexFilter sets RegexFilter field to given value.

### HasRegexFilter

`func (o *V1alpha3GitlabSource) HasRegexFilter() bool`

HasRegexFilter returns a boolean if a field has been set.

### GetRepo

`func (o *V1alpha3GitlabSource) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *V1alpha3GitlabSource) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *V1alpha3GitlabSource) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *V1alpha3GitlabSource) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetScmId

`func (o *V1alpha3GitlabSource) GetScmId() string`

GetScmId returns the ScmId field if non-nil, zero value otherwise.

### GetScmIdOk

`func (o *V1alpha3GitlabSource) GetScmIdOk() (*string, bool)`

GetScmIdOk returns a tuple with the ScmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmId

`func (o *V1alpha3GitlabSource) SetScmId(v string)`

SetScmId sets ScmId field to given value.

### HasScmId

`func (o *V1alpha3GitlabSource) HasScmId() bool`

HasScmId returns a boolean if a field has been set.

### GetServerName

`func (o *V1alpha3GitlabSource) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *V1alpha3GitlabSource) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *V1alpha3GitlabSource) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *V1alpha3GitlabSource) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


