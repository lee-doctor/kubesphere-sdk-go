# DevopsCheckScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **int32** | column e.g. 0 | [optional] 
**Line** | Pointer to **int32** | line e.g. 0 | [optional] 
**Message** | Pointer to **string** | message e.g. unexpected char: &#39;#&#39; | [optional] 
**Status** | Pointer to **string** | status e.g. fail | [optional] 

## Methods

### NewDevopsCheckScript

`func NewDevopsCheckScript() *DevopsCheckScript`

NewDevopsCheckScript instantiates a new DevopsCheckScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsCheckScriptWithDefaults

`func NewDevopsCheckScriptWithDefaults() *DevopsCheckScript`

NewDevopsCheckScriptWithDefaults instantiates a new DevopsCheckScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *DevopsCheckScript) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *DevopsCheckScript) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *DevopsCheckScript) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *DevopsCheckScript) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetLine

`func (o *DevopsCheckScript) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DevopsCheckScript) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DevopsCheckScript) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *DevopsCheckScript) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetMessage

`func (o *DevopsCheckScript) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DevopsCheckScript) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DevopsCheckScript) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DevopsCheckScript) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DevopsCheckScript) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevopsCheckScript) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevopsCheckScript) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevopsCheckScript) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


