# DevopsOrgRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Repositories** | Pointer to [**Repositories**](Repositories.md) |  | [optional] 

## Methods

### NewDevopsOrgRepo

`func NewDevopsOrgRepo() *DevopsOrgRepo`

NewDevopsOrgRepo instantiates a new DevopsOrgRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsOrgRepoWithDefaults

`func NewDevopsOrgRepoWithDefaults() *DevopsOrgRepo`

NewDevopsOrgRepoWithDefaults instantiates a new DevopsOrgRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsOrgRepo) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsOrgRepo) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsOrgRepo) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsOrgRepo) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsOrgRepo) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsOrgRepo) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsOrgRepo) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsOrgRepo) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRepositories

`func (o *DevopsOrgRepo) GetRepositories() Repositories`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *DevopsOrgRepo) GetRepositoriesOk() (*Repositories, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *DevopsOrgRepo) SetRepositories(v Repositories)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *DevopsOrgRepo) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


