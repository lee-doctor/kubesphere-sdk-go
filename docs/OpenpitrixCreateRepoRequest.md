# OpenpitrixCreateRepoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDefaultStatus** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Providers** | **[]string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Workspace** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenpitrixCreateRepoRequest

`func NewOpenpitrixCreateRepoRequest(providers []string, ) *OpenpitrixCreateRepoRequest`

NewOpenpitrixCreateRepoRequest instantiates a new OpenpitrixCreateRepoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenpitrixCreateRepoRequestWithDefaults

`func NewOpenpitrixCreateRepoRequestWithDefaults() *OpenpitrixCreateRepoRequest`

NewOpenpitrixCreateRepoRequestWithDefaults instantiates a new OpenpitrixCreateRepoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDefaultStatus

`func (o *OpenpitrixCreateRepoRequest) GetAppDefaultStatus() string`

GetAppDefaultStatus returns the AppDefaultStatus field if non-nil, zero value otherwise.

### GetAppDefaultStatusOk

`func (o *OpenpitrixCreateRepoRequest) GetAppDefaultStatusOk() (*string, bool)`

GetAppDefaultStatusOk returns a tuple with the AppDefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDefaultStatus

`func (o *OpenpitrixCreateRepoRequest) SetAppDefaultStatus(v string)`

SetAppDefaultStatus sets AppDefaultStatus field to given value.

### HasAppDefaultStatus

`func (o *OpenpitrixCreateRepoRequest) HasAppDefaultStatus() bool`

HasAppDefaultStatus returns a boolean if a field has been set.

### GetCategoryId

`func (o *OpenpitrixCreateRepoRequest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *OpenpitrixCreateRepoRequest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *OpenpitrixCreateRepoRequest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *OpenpitrixCreateRepoRequest) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCredential

`func (o *OpenpitrixCreateRepoRequest) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OpenpitrixCreateRepoRequest) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OpenpitrixCreateRepoRequest) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OpenpitrixCreateRepoRequest) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDescription

`func (o *OpenpitrixCreateRepoRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenpitrixCreateRepoRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenpitrixCreateRepoRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenpitrixCreateRepoRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OpenpitrixCreateRepoRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenpitrixCreateRepoRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenpitrixCreateRepoRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenpitrixCreateRepoRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviders

`func (o *OpenpitrixCreateRepoRequest) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OpenpitrixCreateRepoRequest) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OpenpitrixCreateRepoRequest) SetProviders(v []string)`

SetProviders sets Providers field to given value.


### GetType

`func (o *OpenpitrixCreateRepoRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenpitrixCreateRepoRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenpitrixCreateRepoRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenpitrixCreateRepoRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *OpenpitrixCreateRepoRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OpenpitrixCreateRepoRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OpenpitrixCreateRepoRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OpenpitrixCreateRepoRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OpenpitrixCreateRepoRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OpenpitrixCreateRepoRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OpenpitrixCreateRepoRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OpenpitrixCreateRepoRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetWorkspace

`func (o *OpenpitrixCreateRepoRequest) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *OpenpitrixCreateRepoRequest) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *OpenpitrixCreateRepoRequest) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *OpenpitrixCreateRepoRequest) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


