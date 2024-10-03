# OpenpitrixRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDefaultStatus** | Pointer to **string** |  | [optional] 
**CategorySet** | [**[]OpenpitrixResourceCategory**](OpenpitrixResourceCategory.md) |  | 
**Controller** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | [**[]OpenpitrixRepoLabel**](OpenpitrixRepoLabel.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Providers** | **[]string** |  | 
**RepoId** | Pointer to **string** |  | [optional] 
**Selectors** | [**[]OpenpitrixRepoSelector**](OpenpitrixRepoSelector.md) |  | 
**Status** | Pointer to **string** |  | [optional] 
**StatusTime** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenpitrixRepo

`func NewOpenpitrixRepo(categorySet []OpenpitrixResourceCategory, labels []OpenpitrixRepoLabel, providers []string, selectors []OpenpitrixRepoSelector, ) *OpenpitrixRepo`

NewOpenpitrixRepo instantiates a new OpenpitrixRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenpitrixRepoWithDefaults

`func NewOpenpitrixRepoWithDefaults() *OpenpitrixRepo`

NewOpenpitrixRepoWithDefaults instantiates a new OpenpitrixRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDefaultStatus

`func (o *OpenpitrixRepo) GetAppDefaultStatus() string`

GetAppDefaultStatus returns the AppDefaultStatus field if non-nil, zero value otherwise.

### GetAppDefaultStatusOk

`func (o *OpenpitrixRepo) GetAppDefaultStatusOk() (*string, bool)`

GetAppDefaultStatusOk returns a tuple with the AppDefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDefaultStatus

`func (o *OpenpitrixRepo) SetAppDefaultStatus(v string)`

SetAppDefaultStatus sets AppDefaultStatus field to given value.

### HasAppDefaultStatus

`func (o *OpenpitrixRepo) HasAppDefaultStatus() bool`

HasAppDefaultStatus returns a boolean if a field has been set.

### GetCategorySet

`func (o *OpenpitrixRepo) GetCategorySet() []OpenpitrixResourceCategory`

GetCategorySet returns the CategorySet field if non-nil, zero value otherwise.

### GetCategorySetOk

`func (o *OpenpitrixRepo) GetCategorySetOk() (*[]OpenpitrixResourceCategory, bool)`

GetCategorySetOk returns a tuple with the CategorySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySet

`func (o *OpenpitrixRepo) SetCategorySet(v []OpenpitrixResourceCategory)`

SetCategorySet sets CategorySet field to given value.


### GetController

`func (o *OpenpitrixRepo) GetController() int32`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *OpenpitrixRepo) GetControllerOk() (*int32, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *OpenpitrixRepo) SetController(v int32)`

SetController sets Controller field to given value.

### HasController

`func (o *OpenpitrixRepo) HasController() bool`

HasController returns a boolean if a field has been set.

### GetCreateTime

`func (o *OpenpitrixRepo) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OpenpitrixRepo) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OpenpitrixRepo) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OpenpitrixRepo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCredential

`func (o *OpenpitrixRepo) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OpenpitrixRepo) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OpenpitrixRepo) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OpenpitrixRepo) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDescription

`func (o *OpenpitrixRepo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenpitrixRepo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenpitrixRepo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenpitrixRepo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *OpenpitrixRepo) GetLabels() []OpenpitrixRepoLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OpenpitrixRepo) GetLabelsOk() (*[]OpenpitrixRepoLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OpenpitrixRepo) SetLabels(v []OpenpitrixRepoLabel)`

SetLabels sets Labels field to given value.


### GetName

`func (o *OpenpitrixRepo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenpitrixRepo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenpitrixRepo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenpitrixRepo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *OpenpitrixRepo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OpenpitrixRepo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OpenpitrixRepo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *OpenpitrixRepo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProviders

`func (o *OpenpitrixRepo) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OpenpitrixRepo) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OpenpitrixRepo) SetProviders(v []string)`

SetProviders sets Providers field to given value.


### GetRepoId

`func (o *OpenpitrixRepo) GetRepoId() string`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *OpenpitrixRepo) GetRepoIdOk() (*string, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *OpenpitrixRepo) SetRepoId(v string)`

SetRepoId sets RepoId field to given value.

### HasRepoId

`func (o *OpenpitrixRepo) HasRepoId() bool`

HasRepoId returns a boolean if a field has been set.

### GetSelectors

`func (o *OpenpitrixRepo) GetSelectors() []OpenpitrixRepoSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *OpenpitrixRepo) GetSelectorsOk() (*[]OpenpitrixRepoSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *OpenpitrixRepo) SetSelectors(v []OpenpitrixRepoSelector)`

SetSelectors sets Selectors field to given value.


### GetStatus

`func (o *OpenpitrixRepo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenpitrixRepo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenpitrixRepo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenpitrixRepo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *OpenpitrixRepo) GetStatusTime() string`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *OpenpitrixRepo) GetStatusTimeOk() (*string, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *OpenpitrixRepo) SetStatusTime(v string)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *OpenpitrixRepo) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetType

`func (o *OpenpitrixRepo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenpitrixRepo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenpitrixRepo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenpitrixRepo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *OpenpitrixRepo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OpenpitrixRepo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OpenpitrixRepo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OpenpitrixRepo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OpenpitrixRepo) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OpenpitrixRepo) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OpenpitrixRepo) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OpenpitrixRepo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


