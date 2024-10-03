# DevopsSCMServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**ApiUrl** | Pointer to **string** | url of scm server | [optional] 
**Id** | Pointer to **string** | server id of scm server | [optional] 
**Name** | Pointer to **string** | name of scm server | [optional] 

## Methods

### NewDevopsSCMServer

`func NewDevopsSCMServer() *DevopsSCMServer`

NewDevopsSCMServer instantiates a new DevopsSCMServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsSCMServerWithDefaults

`func NewDevopsSCMServerWithDefaults() *DevopsSCMServer`

NewDevopsSCMServerWithDefaults instantiates a new DevopsSCMServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsSCMServer) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsSCMServer) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsSCMServer) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsSCMServer) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsSCMServer) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsSCMServer) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsSCMServer) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsSCMServer) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApiUrl

`func (o *DevopsSCMServer) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *DevopsSCMServer) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *DevopsSCMServer) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *DevopsSCMServer) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetId

`func (o *DevopsSCMServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsSCMServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsSCMServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsSCMServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DevopsSCMServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsSCMServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsSCMServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevopsSCMServer) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


