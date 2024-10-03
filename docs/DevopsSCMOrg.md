# DevopsSCMOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Avatar** | Pointer to **string** | the url of organization avatar | [optional] 
**JenkinsOrganizationPipeline** | Pointer to **bool** | weather or not already have jenkins pipeline. | [optional] 
**Name** | Pointer to **string** | organization name | [optional] 

## Methods

### NewDevopsSCMOrg

`func NewDevopsSCMOrg() *DevopsSCMOrg`

NewDevopsSCMOrg instantiates a new DevopsSCMOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsSCMOrgWithDefaults

`func NewDevopsSCMOrgWithDefaults() *DevopsSCMOrg`

NewDevopsSCMOrgWithDefaults instantiates a new DevopsSCMOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsSCMOrg) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsSCMOrg) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsSCMOrg) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsSCMOrg) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsSCMOrg) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsSCMOrg) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsSCMOrg) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsSCMOrg) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAvatar

`func (o *DevopsSCMOrg) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *DevopsSCMOrg) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *DevopsSCMOrg) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *DevopsSCMOrg) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetJenkinsOrganizationPipeline

`func (o *DevopsSCMOrg) GetJenkinsOrganizationPipeline() bool`

GetJenkinsOrganizationPipeline returns the JenkinsOrganizationPipeline field if non-nil, zero value otherwise.

### GetJenkinsOrganizationPipelineOk

`func (o *DevopsSCMOrg) GetJenkinsOrganizationPipelineOk() (*bool, bool)`

GetJenkinsOrganizationPipelineOk returns a tuple with the JenkinsOrganizationPipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsOrganizationPipeline

`func (o *DevopsSCMOrg) SetJenkinsOrganizationPipeline(v bool)`

SetJenkinsOrganizationPipeline sets JenkinsOrganizationPipeline field to given value.

### HasJenkinsOrganizationPipeline

`func (o *DevopsSCMOrg) HasJenkinsOrganizationPipeline() bool`

HasJenkinsOrganizationPipeline returns a boolean if a field has been set.

### GetName

`func (o *DevopsSCMOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsSCMOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsSCMOrg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevopsSCMOrg) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


