# DevopsPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions. | [optional] 
**BranchNames** | Pointer to **[]string** | branch names | [optional] 
**Disabled** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayName** | Pointer to **string** | display name | [optional] 
**EstimatedDurationInMillis** | Pointer to **int32** | estimated duration time, unit is millis | [optional] 
**FullDisplayName** | Pointer to **string** | full display name | [optional] 
**FullName** | Pointer to **string** | full name | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**NumberOfFailingBranches** | Pointer to **int32** | number of failing branches | [optional] 
**NumberOfFailingPullRequests** | Pointer to **int32** | number of failing pull requests | [optional] 
**NumberOfFolders** | Pointer to **int32** | number of folders | [optional] 
**NumberOfPipelines** | Pointer to **int32** | number of pipelines | [optional] 
**NumberOfSuccessfulBranches** | Pointer to **int32** | number of successful pull requests | [optional] 
**NumberOfSuccessfulPullRequests** | Pointer to **int32** | number of successful pull requests | [optional] 
**Organization** | Pointer to **string** | the name of organization | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**PipelineFolderNames** | Pointer to **[]map[string]interface{}** | pipeline folder names | [optional] 
**ScmSource** | Pointer to [**ScmSource**](ScmSource.md) |  | [optional] 
**TotalNumberOfBranches** | Pointer to **int32** | total number of branches | [optional] 
**TotalNumberOfPullRequests** | Pointer to **int32** | total number of pull requests | [optional] 
**WeatherScore** | **int32** | the score to description the result of pipeline activity | 

## Methods

### NewDevopsPipeline

`func NewDevopsPipeline(weatherScore int32, ) *DevopsPipeline`

NewDevopsPipeline instantiates a new DevopsPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsPipelineWithDefaults

`func NewDevopsPipelineWithDefaults() *DevopsPipeline`

NewDevopsPipelineWithDefaults instantiates a new DevopsPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsPipeline) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsPipeline) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsPipeline) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsPipeline) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsPipeline) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsPipeline) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsPipeline) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsPipeline) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsPipeline) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsPipeline) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsPipeline) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsPipeline) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetBranchNames

`func (o *DevopsPipeline) GetBranchNames() []string`

GetBranchNames returns the BranchNames field if non-nil, zero value otherwise.

### GetBranchNamesOk

`func (o *DevopsPipeline) GetBranchNamesOk() (*[]string, bool)`

GetBranchNamesOk returns a tuple with the BranchNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchNames

`func (o *DevopsPipeline) SetBranchNames(v []string)`

SetBranchNames sets BranchNames field to given value.

### HasBranchNames

`func (o *DevopsPipeline) HasBranchNames() bool`

HasBranchNames returns a boolean if a field has been set.

### GetDisabled

`func (o *DevopsPipeline) GetDisabled() map[string]interface{}`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *DevopsPipeline) GetDisabledOk() (*map[string]interface{}, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *DevopsPipeline) SetDisabled(v map[string]interface{})`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *DevopsPipeline) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsPipeline) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsPipeline) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsPipeline) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsPipeline) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEstimatedDurationInMillis

`func (o *DevopsPipeline) GetEstimatedDurationInMillis() int32`

GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field if non-nil, zero value otherwise.

### GetEstimatedDurationInMillisOk

`func (o *DevopsPipeline) GetEstimatedDurationInMillisOk() (*int32, bool)`

GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMillis

`func (o *DevopsPipeline) SetEstimatedDurationInMillis(v int32)`

SetEstimatedDurationInMillis sets EstimatedDurationInMillis field to given value.

### HasEstimatedDurationInMillis

`func (o *DevopsPipeline) HasEstimatedDurationInMillis() bool`

HasEstimatedDurationInMillis returns a boolean if a field has been set.

### GetFullDisplayName

`func (o *DevopsPipeline) GetFullDisplayName() string`

GetFullDisplayName returns the FullDisplayName field if non-nil, zero value otherwise.

### GetFullDisplayNameOk

`func (o *DevopsPipeline) GetFullDisplayNameOk() (*string, bool)`

GetFullDisplayNameOk returns a tuple with the FullDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDisplayName

`func (o *DevopsPipeline) SetFullDisplayName(v string)`

SetFullDisplayName sets FullDisplayName field to given value.

### HasFullDisplayName

`func (o *DevopsPipeline) HasFullDisplayName() bool`

HasFullDisplayName returns a boolean if a field has been set.

### GetFullName

`func (o *DevopsPipeline) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DevopsPipeline) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DevopsPipeline) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DevopsPipeline) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetName

`func (o *DevopsPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevopsPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfFailingBranches

`func (o *DevopsPipeline) GetNumberOfFailingBranches() int32`

GetNumberOfFailingBranches returns the NumberOfFailingBranches field if non-nil, zero value otherwise.

### GetNumberOfFailingBranchesOk

`func (o *DevopsPipeline) GetNumberOfFailingBranchesOk() (*int32, bool)`

GetNumberOfFailingBranchesOk returns a tuple with the NumberOfFailingBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailingBranches

`func (o *DevopsPipeline) SetNumberOfFailingBranches(v int32)`

SetNumberOfFailingBranches sets NumberOfFailingBranches field to given value.

### HasNumberOfFailingBranches

`func (o *DevopsPipeline) HasNumberOfFailingBranches() bool`

HasNumberOfFailingBranches returns a boolean if a field has been set.

### GetNumberOfFailingPullRequests

`func (o *DevopsPipeline) GetNumberOfFailingPullRequests() int32`

GetNumberOfFailingPullRequests returns the NumberOfFailingPullRequests field if non-nil, zero value otherwise.

### GetNumberOfFailingPullRequestsOk

`func (o *DevopsPipeline) GetNumberOfFailingPullRequestsOk() (*int32, bool)`

GetNumberOfFailingPullRequestsOk returns a tuple with the NumberOfFailingPullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailingPullRequests

`func (o *DevopsPipeline) SetNumberOfFailingPullRequests(v int32)`

SetNumberOfFailingPullRequests sets NumberOfFailingPullRequests field to given value.

### HasNumberOfFailingPullRequests

`func (o *DevopsPipeline) HasNumberOfFailingPullRequests() bool`

HasNumberOfFailingPullRequests returns a boolean if a field has been set.

### GetNumberOfFolders

`func (o *DevopsPipeline) GetNumberOfFolders() int32`

GetNumberOfFolders returns the NumberOfFolders field if non-nil, zero value otherwise.

### GetNumberOfFoldersOk

`func (o *DevopsPipeline) GetNumberOfFoldersOk() (*int32, bool)`

GetNumberOfFoldersOk returns a tuple with the NumberOfFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFolders

`func (o *DevopsPipeline) SetNumberOfFolders(v int32)`

SetNumberOfFolders sets NumberOfFolders field to given value.

### HasNumberOfFolders

`func (o *DevopsPipeline) HasNumberOfFolders() bool`

HasNumberOfFolders returns a boolean if a field has been set.

### GetNumberOfPipelines

`func (o *DevopsPipeline) GetNumberOfPipelines() int32`

GetNumberOfPipelines returns the NumberOfPipelines field if non-nil, zero value otherwise.

### GetNumberOfPipelinesOk

`func (o *DevopsPipeline) GetNumberOfPipelinesOk() (*int32, bool)`

GetNumberOfPipelinesOk returns a tuple with the NumberOfPipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPipelines

`func (o *DevopsPipeline) SetNumberOfPipelines(v int32)`

SetNumberOfPipelines sets NumberOfPipelines field to given value.

### HasNumberOfPipelines

`func (o *DevopsPipeline) HasNumberOfPipelines() bool`

HasNumberOfPipelines returns a boolean if a field has been set.

### GetNumberOfSuccessfulBranches

`func (o *DevopsPipeline) GetNumberOfSuccessfulBranches() int32`

GetNumberOfSuccessfulBranches returns the NumberOfSuccessfulBranches field if non-nil, zero value otherwise.

### GetNumberOfSuccessfulBranchesOk

`func (o *DevopsPipeline) GetNumberOfSuccessfulBranchesOk() (*int32, bool)`

GetNumberOfSuccessfulBranchesOk returns a tuple with the NumberOfSuccessfulBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSuccessfulBranches

`func (o *DevopsPipeline) SetNumberOfSuccessfulBranches(v int32)`

SetNumberOfSuccessfulBranches sets NumberOfSuccessfulBranches field to given value.

### HasNumberOfSuccessfulBranches

`func (o *DevopsPipeline) HasNumberOfSuccessfulBranches() bool`

HasNumberOfSuccessfulBranches returns a boolean if a field has been set.

### GetNumberOfSuccessfulPullRequests

`func (o *DevopsPipeline) GetNumberOfSuccessfulPullRequests() int32`

GetNumberOfSuccessfulPullRequests returns the NumberOfSuccessfulPullRequests field if non-nil, zero value otherwise.

### GetNumberOfSuccessfulPullRequestsOk

`func (o *DevopsPipeline) GetNumberOfSuccessfulPullRequestsOk() (*int32, bool)`

GetNumberOfSuccessfulPullRequestsOk returns a tuple with the NumberOfSuccessfulPullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSuccessfulPullRequests

`func (o *DevopsPipeline) SetNumberOfSuccessfulPullRequests(v int32)`

SetNumberOfSuccessfulPullRequests sets NumberOfSuccessfulPullRequests field to given value.

### HasNumberOfSuccessfulPullRequests

`func (o *DevopsPipeline) HasNumberOfSuccessfulPullRequests() bool`

HasNumberOfSuccessfulPullRequests returns a boolean if a field has been set.

### GetOrganization

`func (o *DevopsPipeline) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DevopsPipeline) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DevopsPipeline) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DevopsPipeline) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetParameters

`func (o *DevopsPipeline) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DevopsPipeline) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DevopsPipeline) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DevopsPipeline) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPermissions

`func (o *DevopsPipeline) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DevopsPipeline) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DevopsPipeline) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DevopsPipeline) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPipelineFolderNames

`func (o *DevopsPipeline) GetPipelineFolderNames() []map[string]interface{}`

GetPipelineFolderNames returns the PipelineFolderNames field if non-nil, zero value otherwise.

### GetPipelineFolderNamesOk

`func (o *DevopsPipeline) GetPipelineFolderNamesOk() (*[]map[string]interface{}, bool)`

GetPipelineFolderNamesOk returns a tuple with the PipelineFolderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineFolderNames

`func (o *DevopsPipeline) SetPipelineFolderNames(v []map[string]interface{})`

SetPipelineFolderNames sets PipelineFolderNames field to given value.

### HasPipelineFolderNames

`func (o *DevopsPipeline) HasPipelineFolderNames() bool`

HasPipelineFolderNames returns a boolean if a field has been set.

### GetScmSource

`func (o *DevopsPipeline) GetScmSource() ScmSource`

GetScmSource returns the ScmSource field if non-nil, zero value otherwise.

### GetScmSourceOk

`func (o *DevopsPipeline) GetScmSourceOk() (*ScmSource, bool)`

GetScmSourceOk returns a tuple with the ScmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmSource

`func (o *DevopsPipeline) SetScmSource(v ScmSource)`

SetScmSource sets ScmSource field to given value.

### HasScmSource

`func (o *DevopsPipeline) HasScmSource() bool`

HasScmSource returns a boolean if a field has been set.

### GetTotalNumberOfBranches

`func (o *DevopsPipeline) GetTotalNumberOfBranches() int32`

GetTotalNumberOfBranches returns the TotalNumberOfBranches field if non-nil, zero value otherwise.

### GetTotalNumberOfBranchesOk

`func (o *DevopsPipeline) GetTotalNumberOfBranchesOk() (*int32, bool)`

GetTotalNumberOfBranchesOk returns a tuple with the TotalNumberOfBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfBranches

`func (o *DevopsPipeline) SetTotalNumberOfBranches(v int32)`

SetTotalNumberOfBranches sets TotalNumberOfBranches field to given value.

### HasTotalNumberOfBranches

`func (o *DevopsPipeline) HasTotalNumberOfBranches() bool`

HasTotalNumberOfBranches returns a boolean if a field has been set.

### GetTotalNumberOfPullRequests

`func (o *DevopsPipeline) GetTotalNumberOfPullRequests() int32`

GetTotalNumberOfPullRequests returns the TotalNumberOfPullRequests field if non-nil, zero value otherwise.

### GetTotalNumberOfPullRequestsOk

`func (o *DevopsPipeline) GetTotalNumberOfPullRequestsOk() (*int32, bool)`

GetTotalNumberOfPullRequestsOk returns a tuple with the TotalNumberOfPullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfPullRequests

`func (o *DevopsPipeline) SetTotalNumberOfPullRequests(v int32)`

SetTotalNumberOfPullRequests sets TotalNumberOfPullRequests field to given value.

### HasTotalNumberOfPullRequests

`func (o *DevopsPipeline) HasTotalNumberOfPullRequests() bool`

HasTotalNumberOfPullRequests returns a boolean if a field has been set.

### GetWeatherScore

`func (o *DevopsPipeline) GetWeatherScore() int32`

GetWeatherScore returns the WeatherScore field if non-nil, zero value otherwise.

### GetWeatherScoreOk

`func (o *DevopsPipeline) GetWeatherScoreOk() (*int32, bool)`

GetWeatherScoreOk returns a tuple with the WeatherScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeatherScore

`func (o *DevopsPipeline) SetWeatherScore(v int32)`

SetWeatherScore sets WeatherScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


