# V1alpha3MultiBranchPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitbucketServerSource** | Pointer to [**V1alpha3BitbucketServerSource**](V1alpha3BitbucketServerSource.md) |  | [optional] 
**Description** | Pointer to **string** | description of pipeline | [optional] 
**Discarder** | Pointer to [**V1alpha3DiscarderProperty**](V1alpha3DiscarderProperty.md) |  | [optional] 
**GitSource** | Pointer to [**V1alpha3GitSource**](V1alpha3GitSource.md) |  | [optional] 
**GithubSource** | Pointer to [**V1alpha3GithubSource**](V1alpha3GithubSource.md) |  | [optional] 
**GitlabSource** | Pointer to [**V1alpha3GitlabSource**](V1alpha3GitlabSource.md) |  | [optional] 
**MultibranchJobTrigger** | Pointer to [**V1alpha3MultiBranchJobTrigger**](V1alpha3MultiBranchJobTrigger.md) |  | [optional] 
**Name** | **string** | name of pipeline | 
**ScriptPath** | **string** | script path in scm | 
**SingleSvnSource** | Pointer to [**V1alpha3SingleSvnSource**](V1alpha3SingleSvnSource.md) |  | [optional] 
**SourceType** | **string** | type of scm, such as github/git/svn | 
**SvnSource** | Pointer to [**V1alpha3SvnSource**](V1alpha3SvnSource.md) |  | [optional] 
**TimerTrigger** | Pointer to [**V1alpha3TimerTrigger**](V1alpha3TimerTrigger.md) |  | [optional] 

## Methods

### NewV1alpha3MultiBranchPipeline

`func NewV1alpha3MultiBranchPipeline(name string, scriptPath string, sourceType string, ) *V1alpha3MultiBranchPipeline`

NewV1alpha3MultiBranchPipeline instantiates a new V1alpha3MultiBranchPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3MultiBranchPipelineWithDefaults

`func NewV1alpha3MultiBranchPipelineWithDefaults() *V1alpha3MultiBranchPipeline`

NewV1alpha3MultiBranchPipelineWithDefaults instantiates a new V1alpha3MultiBranchPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitbucketServerSource

`func (o *V1alpha3MultiBranchPipeline) GetBitbucketServerSource() V1alpha3BitbucketServerSource`

GetBitbucketServerSource returns the BitbucketServerSource field if non-nil, zero value otherwise.

### GetBitbucketServerSourceOk

`func (o *V1alpha3MultiBranchPipeline) GetBitbucketServerSourceOk() (*V1alpha3BitbucketServerSource, bool)`

GetBitbucketServerSourceOk returns a tuple with the BitbucketServerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucketServerSource

`func (o *V1alpha3MultiBranchPipeline) SetBitbucketServerSource(v V1alpha3BitbucketServerSource)`

SetBitbucketServerSource sets BitbucketServerSource field to given value.

### HasBitbucketServerSource

`func (o *V1alpha3MultiBranchPipeline) HasBitbucketServerSource() bool`

HasBitbucketServerSource returns a boolean if a field has been set.

### GetDescription

`func (o *V1alpha3MultiBranchPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1alpha3MultiBranchPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1alpha3MultiBranchPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1alpha3MultiBranchPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscarder

`func (o *V1alpha3MultiBranchPipeline) GetDiscarder() V1alpha3DiscarderProperty`

GetDiscarder returns the Discarder field if non-nil, zero value otherwise.

### GetDiscarderOk

`func (o *V1alpha3MultiBranchPipeline) GetDiscarderOk() (*V1alpha3DiscarderProperty, bool)`

GetDiscarderOk returns a tuple with the Discarder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscarder

`func (o *V1alpha3MultiBranchPipeline) SetDiscarder(v V1alpha3DiscarderProperty)`

SetDiscarder sets Discarder field to given value.

### HasDiscarder

`func (o *V1alpha3MultiBranchPipeline) HasDiscarder() bool`

HasDiscarder returns a boolean if a field has been set.

### GetGitSource

`func (o *V1alpha3MultiBranchPipeline) GetGitSource() V1alpha3GitSource`

GetGitSource returns the GitSource field if non-nil, zero value otherwise.

### GetGitSourceOk

`func (o *V1alpha3MultiBranchPipeline) GetGitSourceOk() (*V1alpha3GitSource, bool)`

GetGitSourceOk returns a tuple with the GitSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSource

`func (o *V1alpha3MultiBranchPipeline) SetGitSource(v V1alpha3GitSource)`

SetGitSource sets GitSource field to given value.

### HasGitSource

`func (o *V1alpha3MultiBranchPipeline) HasGitSource() bool`

HasGitSource returns a boolean if a field has been set.

### GetGithubSource

`func (o *V1alpha3MultiBranchPipeline) GetGithubSource() V1alpha3GithubSource`

GetGithubSource returns the GithubSource field if non-nil, zero value otherwise.

### GetGithubSourceOk

`func (o *V1alpha3MultiBranchPipeline) GetGithubSourceOk() (*V1alpha3GithubSource, bool)`

GetGithubSourceOk returns a tuple with the GithubSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSource

`func (o *V1alpha3MultiBranchPipeline) SetGithubSource(v V1alpha3GithubSource)`

SetGithubSource sets GithubSource field to given value.

### HasGithubSource

`func (o *V1alpha3MultiBranchPipeline) HasGithubSource() bool`

HasGithubSource returns a boolean if a field has been set.

### GetGitlabSource

`func (o *V1alpha3MultiBranchPipeline) GetGitlabSource() V1alpha3GitlabSource`

GetGitlabSource returns the GitlabSource field if non-nil, zero value otherwise.

### GetGitlabSourceOk

`func (o *V1alpha3MultiBranchPipeline) GetGitlabSourceOk() (*V1alpha3GitlabSource, bool)`

GetGitlabSourceOk returns a tuple with the GitlabSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabSource

`func (o *V1alpha3MultiBranchPipeline) SetGitlabSource(v V1alpha3GitlabSource)`

SetGitlabSource sets GitlabSource field to given value.

### HasGitlabSource

`func (o *V1alpha3MultiBranchPipeline) HasGitlabSource() bool`

HasGitlabSource returns a boolean if a field has been set.

### GetMultibranchJobTrigger

`func (o *V1alpha3MultiBranchPipeline) GetMultibranchJobTrigger() V1alpha3MultiBranchJobTrigger`

GetMultibranchJobTrigger returns the MultibranchJobTrigger field if non-nil, zero value otherwise.

### GetMultibranchJobTriggerOk

`func (o *V1alpha3MultiBranchPipeline) GetMultibranchJobTriggerOk() (*V1alpha3MultiBranchJobTrigger, bool)`

GetMultibranchJobTriggerOk returns a tuple with the MultibranchJobTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultibranchJobTrigger

`func (o *V1alpha3MultiBranchPipeline) SetMultibranchJobTrigger(v V1alpha3MultiBranchJobTrigger)`

SetMultibranchJobTrigger sets MultibranchJobTrigger field to given value.

### HasMultibranchJobTrigger

`func (o *V1alpha3MultiBranchPipeline) HasMultibranchJobTrigger() bool`

HasMultibranchJobTrigger returns a boolean if a field has been set.

### GetName

`func (o *V1alpha3MultiBranchPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha3MultiBranchPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha3MultiBranchPipeline) SetName(v string)`

SetName sets Name field to given value.


### GetScriptPath

`func (o *V1alpha3MultiBranchPipeline) GetScriptPath() string`

GetScriptPath returns the ScriptPath field if non-nil, zero value otherwise.

### GetScriptPathOk

`func (o *V1alpha3MultiBranchPipeline) GetScriptPathOk() (*string, bool)`

GetScriptPathOk returns a tuple with the ScriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPath

`func (o *V1alpha3MultiBranchPipeline) SetScriptPath(v string)`

SetScriptPath sets ScriptPath field to given value.


### GetSingleSvnSource

`func (o *V1alpha3MultiBranchPipeline) GetSingleSvnSource() V1alpha3SingleSvnSource`

GetSingleSvnSource returns the SingleSvnSource field if non-nil, zero value otherwise.

### GetSingleSvnSourceOk

`func (o *V1alpha3MultiBranchPipeline) GetSingleSvnSourceOk() (*V1alpha3SingleSvnSource, bool)`

GetSingleSvnSourceOk returns a tuple with the SingleSvnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSvnSource

`func (o *V1alpha3MultiBranchPipeline) SetSingleSvnSource(v V1alpha3SingleSvnSource)`

SetSingleSvnSource sets SingleSvnSource field to given value.

### HasSingleSvnSource

`func (o *V1alpha3MultiBranchPipeline) HasSingleSvnSource() bool`

HasSingleSvnSource returns a boolean if a field has been set.

### GetSourceType

`func (o *V1alpha3MultiBranchPipeline) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *V1alpha3MultiBranchPipeline) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *V1alpha3MultiBranchPipeline) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSvnSource

`func (o *V1alpha3MultiBranchPipeline) GetSvnSource() V1alpha3SvnSource`

GetSvnSource returns the SvnSource field if non-nil, zero value otherwise.

### GetSvnSourceOk

`func (o *V1alpha3MultiBranchPipeline) GetSvnSourceOk() (*V1alpha3SvnSource, bool)`

GetSvnSourceOk returns a tuple with the SvnSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnSource

`func (o *V1alpha3MultiBranchPipeline) SetSvnSource(v V1alpha3SvnSource)`

SetSvnSource sets SvnSource field to given value.

### HasSvnSource

`func (o *V1alpha3MultiBranchPipeline) HasSvnSource() bool`

HasSvnSource returns a boolean if a field has been set.

### GetTimerTrigger

`func (o *V1alpha3MultiBranchPipeline) GetTimerTrigger() V1alpha3TimerTrigger`

GetTimerTrigger returns the TimerTrigger field if non-nil, zero value otherwise.

### GetTimerTriggerOk

`func (o *V1alpha3MultiBranchPipeline) GetTimerTriggerOk() (*V1alpha3TimerTrigger, bool)`

GetTimerTriggerOk returns a tuple with the TimerTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerTrigger

`func (o *V1alpha3MultiBranchPipeline) SetTimerTrigger(v V1alpha3TimerTrigger)`

SetTimerTrigger sets TimerTrigger field to given value.

### HasTimerTrigger

`func (o *V1alpha3MultiBranchPipeline) HasTimerTrigger() bool`

HasTimerTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


