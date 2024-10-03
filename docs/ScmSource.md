# ScmSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | The id of the source configuration management (SCM). | [optional] 

## Methods

### NewScmSource

`func NewScmSource() *ScmSource`

NewScmSource instantiates a new ScmSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScmSourceWithDefaults

`func NewScmSourceWithDefaults() *ScmSource`

NewScmSourceWithDefaults instantiates a new ScmSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *ScmSource) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ScmSource) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ScmSource) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ScmSource) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetApiUrl

`func (o *ScmSource) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ScmSource) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ScmSource) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ScmSource) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetId

`func (o *ScmSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScmSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScmSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScmSource) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


