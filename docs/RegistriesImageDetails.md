# RegistriesImageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBlob** | Pointer to [**RegistriesImageBlob**](RegistriesImageBlob.md) |  | [optional] 
**ImageManifest** | Pointer to [**RegistriesImageManifest**](RegistriesImageManifest.md) |  | [optional] 
**ImageTag** | Pointer to **string** | image tag. | [optional] 
**Message** | Pointer to **string** | Status message. | [optional] 
**Registry** | Pointer to **string** | registry domain. | [optional] 
**Status** | Pointer to **string** | Status is the status of the image search, such as \&quot;succeeded\&quot;. | [optional] 

## Methods

### NewRegistriesImageDetails

`func NewRegistriesImageDetails() *RegistriesImageDetails`

NewRegistriesImageDetails instantiates a new RegistriesImageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesImageDetailsWithDefaults

`func NewRegistriesImageDetailsWithDefaults() *RegistriesImageDetails`

NewRegistriesImageDetailsWithDefaults instantiates a new RegistriesImageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageBlob

`func (o *RegistriesImageDetails) GetImageBlob() RegistriesImageBlob`

GetImageBlob returns the ImageBlob field if non-nil, zero value otherwise.

### GetImageBlobOk

`func (o *RegistriesImageDetails) GetImageBlobOk() (*RegistriesImageBlob, bool)`

GetImageBlobOk returns a tuple with the ImageBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBlob

`func (o *RegistriesImageDetails) SetImageBlob(v RegistriesImageBlob)`

SetImageBlob sets ImageBlob field to given value.

### HasImageBlob

`func (o *RegistriesImageDetails) HasImageBlob() bool`

HasImageBlob returns a boolean if a field has been set.

### GetImageManifest

`func (o *RegistriesImageDetails) GetImageManifest() RegistriesImageManifest`

GetImageManifest returns the ImageManifest field if non-nil, zero value otherwise.

### GetImageManifestOk

`func (o *RegistriesImageDetails) GetImageManifestOk() (*RegistriesImageManifest, bool)`

GetImageManifestOk returns a tuple with the ImageManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageManifest

`func (o *RegistriesImageDetails) SetImageManifest(v RegistriesImageManifest)`

SetImageManifest sets ImageManifest field to given value.

### HasImageManifest

`func (o *RegistriesImageDetails) HasImageManifest() bool`

HasImageManifest returns a boolean if a field has been set.

### GetImageTag

`func (o *RegistriesImageDetails) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *RegistriesImageDetails) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *RegistriesImageDetails) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *RegistriesImageDetails) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetMessage

`func (o *RegistriesImageDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegistriesImageDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegistriesImageDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RegistriesImageDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRegistry

`func (o *RegistriesImageDetails) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *RegistriesImageDetails) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *RegistriesImageDetails) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *RegistriesImageDetails) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetStatus

`func (o *RegistriesImageDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistriesImageDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistriesImageDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegistriesImageDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


