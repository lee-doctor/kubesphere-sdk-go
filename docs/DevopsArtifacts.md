# DevopsArtifacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Downloadable** | Pointer to **bool** | downloadable or not | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Path** | Pointer to **string** | path | [optional] 
**Size** | Pointer to **int32** | size | [optional] 
**Url** | Pointer to **string** | The url for Download artifacts | [optional] 

## Methods

### NewDevopsArtifacts

`func NewDevopsArtifacts() *DevopsArtifacts`

NewDevopsArtifacts instantiates a new DevopsArtifacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsArtifactsWithDefaults

`func NewDevopsArtifactsWithDefaults() *DevopsArtifacts`

NewDevopsArtifactsWithDefaults instantiates a new DevopsArtifacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsArtifacts) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsArtifacts) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsArtifacts) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsArtifacts) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsArtifacts) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsArtifacts) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsArtifacts) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsArtifacts) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDownloadable

`func (o *DevopsArtifacts) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *DevopsArtifacts) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *DevopsArtifacts) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *DevopsArtifacts) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetId

`func (o *DevopsArtifacts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsArtifacts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsArtifacts) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsArtifacts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DevopsArtifacts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsArtifacts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsArtifacts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevopsArtifacts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *DevopsArtifacts) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DevopsArtifacts) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DevopsArtifacts) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DevopsArtifacts) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *DevopsArtifacts) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DevopsArtifacts) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DevopsArtifacts) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DevopsArtifacts) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *DevopsArtifacts) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevopsArtifacts) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevopsArtifacts) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevopsArtifacts) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


