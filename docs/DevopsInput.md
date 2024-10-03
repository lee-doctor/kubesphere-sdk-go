# DevopsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**DevopsInputLinks**](DevopsInputLinks.md) |  | [optional] 
**Id** | Pointer to **string** | the id of check action | [optional] 
**Message** | Pointer to **string** | the message of check action | [optional] 
**Ok** | Pointer to **string** | check status. e.g. \&quot;Proceed\&quot; | [optional] 
**Parameters** | Pointer to **[]map[string]interface{}** | the parameters of check action | [optional] 
**Submitter** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDevopsInput

`func NewDevopsInput() *DevopsInput`

NewDevopsInput instantiates a new DevopsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsInputWithDefaults

`func NewDevopsInputWithDefaults() *DevopsInput`

NewDevopsInputWithDefaults instantiates a new DevopsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsInput) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsInput) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsInput) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsInput) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsInput) GetLinks() DevopsInputLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsInput) GetLinksOk() (*DevopsInputLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsInput) SetLinks(v DevopsInputLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsInput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *DevopsInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevopsInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *DevopsInput) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DevopsInput) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DevopsInput) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DevopsInput) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOk

`func (o *DevopsInput) GetOk() string`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *DevopsInput) GetOkOk() (*string, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *DevopsInput) SetOk(v string)`

SetOk sets Ok field to given value.

### HasOk

`func (o *DevopsInput) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetParameters

`func (o *DevopsInput) GetParameters() []map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DevopsInput) GetParametersOk() (*[]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DevopsInput) SetParameters(v []map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DevopsInput) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSubmitter

`func (o *DevopsInput) GetSubmitter() map[string]interface{}`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *DevopsInput) GetSubmitterOk() (*map[string]interface{}, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *DevopsInput) SetSubmitter(v map[string]interface{})`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *DevopsInput) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


