# OpenpitrixApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstraction** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**AppVersionTypes** | Pointer to **string** |  | [optional] 
**CategorySet** | [**[]OpenpitrixResourceCategory**](OpenpitrixResourceCategory.md) |  | 
**ChartName** | Pointer to **string** |  | [optional] 
**ClusterTotal** | Pointer to **int32** |  | [optional] 
**CompanyJoinTime** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanyProfile** | Pointer to **string** |  | [optional] 
**CompanyWebsite** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Isv** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **string** |  | [optional] 
**LatestAppVersion** | Pointer to [**OpenpitrixAppVersion**](OpenpitrixAppVersion.md) |  | [optional] 
**Maintainers** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Readme** | Pointer to **string** |  | [optional] 
**RepoId** | Pointer to **string** |  | [optional] 
**Screenshots** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusTime** | Pointer to **string** |  | [optional] 
**Tos** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenpitrixApp

`func NewOpenpitrixApp(categorySet []OpenpitrixResourceCategory, ) *OpenpitrixApp`

NewOpenpitrixApp instantiates a new OpenpitrixApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenpitrixAppWithDefaults

`func NewOpenpitrixAppWithDefaults() *OpenpitrixApp`

NewOpenpitrixAppWithDefaults instantiates a new OpenpitrixApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstraction

`func (o *OpenpitrixApp) GetAbstraction() string`

GetAbstraction returns the Abstraction field if non-nil, zero value otherwise.

### GetAbstractionOk

`func (o *OpenpitrixApp) GetAbstractionOk() (*string, bool)`

GetAbstractionOk returns a tuple with the Abstraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstraction

`func (o *OpenpitrixApp) SetAbstraction(v string)`

SetAbstraction sets Abstraction field to given value.

### HasAbstraction

`func (o *OpenpitrixApp) HasAbstraction() bool`

HasAbstraction returns a boolean if a field has been set.

### GetActive

`func (o *OpenpitrixApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OpenpitrixApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OpenpitrixApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OpenpitrixApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAppId

`func (o *OpenpitrixApp) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *OpenpitrixApp) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *OpenpitrixApp) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *OpenpitrixApp) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppVersionTypes

`func (o *OpenpitrixApp) GetAppVersionTypes() string`

GetAppVersionTypes returns the AppVersionTypes field if non-nil, zero value otherwise.

### GetAppVersionTypesOk

`func (o *OpenpitrixApp) GetAppVersionTypesOk() (*string, bool)`

GetAppVersionTypesOk returns a tuple with the AppVersionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersionTypes

`func (o *OpenpitrixApp) SetAppVersionTypes(v string)`

SetAppVersionTypes sets AppVersionTypes field to given value.

### HasAppVersionTypes

`func (o *OpenpitrixApp) HasAppVersionTypes() bool`

HasAppVersionTypes returns a boolean if a field has been set.

### GetCategorySet

`func (o *OpenpitrixApp) GetCategorySet() []OpenpitrixResourceCategory`

GetCategorySet returns the CategorySet field if non-nil, zero value otherwise.

### GetCategorySetOk

`func (o *OpenpitrixApp) GetCategorySetOk() (*[]OpenpitrixResourceCategory, bool)`

GetCategorySetOk returns a tuple with the CategorySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySet

`func (o *OpenpitrixApp) SetCategorySet(v []OpenpitrixResourceCategory)`

SetCategorySet sets CategorySet field to given value.


### GetChartName

`func (o *OpenpitrixApp) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *OpenpitrixApp) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *OpenpitrixApp) SetChartName(v string)`

SetChartName sets ChartName field to given value.

### HasChartName

`func (o *OpenpitrixApp) HasChartName() bool`

HasChartName returns a boolean if a field has been set.

### GetClusterTotal

`func (o *OpenpitrixApp) GetClusterTotal() int32`

GetClusterTotal returns the ClusterTotal field if non-nil, zero value otherwise.

### GetClusterTotalOk

`func (o *OpenpitrixApp) GetClusterTotalOk() (*int32, bool)`

GetClusterTotalOk returns a tuple with the ClusterTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTotal

`func (o *OpenpitrixApp) SetClusterTotal(v int32)`

SetClusterTotal sets ClusterTotal field to given value.

### HasClusterTotal

`func (o *OpenpitrixApp) HasClusterTotal() bool`

HasClusterTotal returns a boolean if a field has been set.

### GetCompanyJoinTime

`func (o *OpenpitrixApp) GetCompanyJoinTime() string`

GetCompanyJoinTime returns the CompanyJoinTime field if non-nil, zero value otherwise.

### GetCompanyJoinTimeOk

`func (o *OpenpitrixApp) GetCompanyJoinTimeOk() (*string, bool)`

GetCompanyJoinTimeOk returns a tuple with the CompanyJoinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyJoinTime

`func (o *OpenpitrixApp) SetCompanyJoinTime(v string)`

SetCompanyJoinTime sets CompanyJoinTime field to given value.

### HasCompanyJoinTime

`func (o *OpenpitrixApp) HasCompanyJoinTime() bool`

HasCompanyJoinTime returns a boolean if a field has been set.

### GetCompanyName

`func (o *OpenpitrixApp) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OpenpitrixApp) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OpenpitrixApp) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OpenpitrixApp) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyProfile

`func (o *OpenpitrixApp) GetCompanyProfile() string`

GetCompanyProfile returns the CompanyProfile field if non-nil, zero value otherwise.

### GetCompanyProfileOk

`func (o *OpenpitrixApp) GetCompanyProfileOk() (*string, bool)`

GetCompanyProfileOk returns a tuple with the CompanyProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyProfile

`func (o *OpenpitrixApp) SetCompanyProfile(v string)`

SetCompanyProfile sets CompanyProfile field to given value.

### HasCompanyProfile

`func (o *OpenpitrixApp) HasCompanyProfile() bool`

HasCompanyProfile returns a boolean if a field has been set.

### GetCompanyWebsite

`func (o *OpenpitrixApp) GetCompanyWebsite() string`

GetCompanyWebsite returns the CompanyWebsite field if non-nil, zero value otherwise.

### GetCompanyWebsiteOk

`func (o *OpenpitrixApp) GetCompanyWebsiteOk() (*string, bool)`

GetCompanyWebsiteOk returns a tuple with the CompanyWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyWebsite

`func (o *OpenpitrixApp) SetCompanyWebsite(v string)`

SetCompanyWebsite sets CompanyWebsite field to given value.

### HasCompanyWebsite

`func (o *OpenpitrixApp) HasCompanyWebsite() bool`

HasCompanyWebsite returns a boolean if a field has been set.

### GetCreateTime

`func (o *OpenpitrixApp) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OpenpitrixApp) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OpenpitrixApp) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OpenpitrixApp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDescription

`func (o *OpenpitrixApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenpitrixApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenpitrixApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenpitrixApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHome

`func (o *OpenpitrixApp) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *OpenpitrixApp) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *OpenpitrixApp) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *OpenpitrixApp) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIcon

`func (o *OpenpitrixApp) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *OpenpitrixApp) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *OpenpitrixApp) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *OpenpitrixApp) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsv

`func (o *OpenpitrixApp) GetIsv() string`

GetIsv returns the Isv field if non-nil, zero value otherwise.

### GetIsvOk

`func (o *OpenpitrixApp) GetIsvOk() (*string, bool)`

GetIsvOk returns a tuple with the Isv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsv

`func (o *OpenpitrixApp) SetIsv(v string)`

SetIsv sets Isv field to given value.

### HasIsv

`func (o *OpenpitrixApp) HasIsv() bool`

HasIsv returns a boolean if a field has been set.

### GetKeywords

`func (o *OpenpitrixApp) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *OpenpitrixApp) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *OpenpitrixApp) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *OpenpitrixApp) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLatestAppVersion

`func (o *OpenpitrixApp) GetLatestAppVersion() OpenpitrixAppVersion`

GetLatestAppVersion returns the LatestAppVersion field if non-nil, zero value otherwise.

### GetLatestAppVersionOk

`func (o *OpenpitrixApp) GetLatestAppVersionOk() (*OpenpitrixAppVersion, bool)`

GetLatestAppVersionOk returns a tuple with the LatestAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAppVersion

`func (o *OpenpitrixApp) SetLatestAppVersion(v OpenpitrixAppVersion)`

SetLatestAppVersion sets LatestAppVersion field to given value.

### HasLatestAppVersion

`func (o *OpenpitrixApp) HasLatestAppVersion() bool`

HasLatestAppVersion returns a boolean if a field has been set.

### GetMaintainers

`func (o *OpenpitrixApp) GetMaintainers() string`

GetMaintainers returns the Maintainers field if non-nil, zero value otherwise.

### GetMaintainersOk

`func (o *OpenpitrixApp) GetMaintainersOk() (*string, bool)`

GetMaintainersOk returns a tuple with the Maintainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainers

`func (o *OpenpitrixApp) SetMaintainers(v string)`

SetMaintainers sets Maintainers field to given value.

### HasMaintainers

`func (o *OpenpitrixApp) HasMaintainers() bool`

HasMaintainers returns a boolean if a field has been set.

### GetName

`func (o *OpenpitrixApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenpitrixApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenpitrixApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenpitrixApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *OpenpitrixApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OpenpitrixApp) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OpenpitrixApp) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *OpenpitrixApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetReadme

`func (o *OpenpitrixApp) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *OpenpitrixApp) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *OpenpitrixApp) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *OpenpitrixApp) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### GetRepoId

`func (o *OpenpitrixApp) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *OpenpitrixApp) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *OpenpitrixApp) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.

### HasRepoId

`func (o *OpenpitrixApp) HasRepoId() bool`

HasRepoId returns a boolean if a field has been set.

### GetScreenshots

`func (o *OpenpitrixApp) GetScreenshots() string`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *OpenpitrixApp) GetScreenshotsOk() (*string, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *OpenpitrixApp) SetScreenshots(v string)`

SetScreenshots sets Screenshots field to given value.

### HasScreenshots

`func (o *OpenpitrixApp) HasScreenshots() bool`

HasScreenshots returns a boolean if a field has been set.

### GetSources

`func (o *OpenpitrixApp) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *OpenpitrixApp) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *OpenpitrixApp) SetSources(v string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *OpenpitrixApp) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStatus

`func (o *OpenpitrixApp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenpitrixApp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenpitrixApp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenpitrixApp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *OpenpitrixApp) GetStatusTime() string`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *OpenpitrixApp) GetStatusTimeOk() (*string, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *OpenpitrixApp) SetStatusTime(v string)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *OpenpitrixApp) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetTos

`func (o *OpenpitrixApp) GetTos() string`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *OpenpitrixApp) GetTosOk() (*string, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *OpenpitrixApp) SetTos(v string)`

SetTos sets Tos field to given value.

### HasTos

`func (o *OpenpitrixApp) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetUpdateTime

`func (o *OpenpitrixApp) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *OpenpitrixApp) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *OpenpitrixApp) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *OpenpitrixApp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


