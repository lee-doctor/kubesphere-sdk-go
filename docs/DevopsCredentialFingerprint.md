# DevopsCredentialFingerprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | Credential&#39;s display name and description | [optional] 
**Hash** | Pointer to **string** | Credential&#39;s hash | [optional] 
**Usage** | Pointer to [**[]DevopsCredentialFingerprintUsage**](DevopsCredentialFingerprintUsage.md) | all usage of Credential | [optional] 

## Methods

### NewDevopsCredentialFingerprint

`func NewDevopsCredentialFingerprint() *DevopsCredentialFingerprint`

NewDevopsCredentialFingerprint instantiates a new DevopsCredentialFingerprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsCredentialFingerprintWithDefaults

`func NewDevopsCredentialFingerprintWithDefaults() *DevopsCredentialFingerprint`

NewDevopsCredentialFingerprintWithDefaults instantiates a new DevopsCredentialFingerprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *DevopsCredentialFingerprint) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DevopsCredentialFingerprint) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DevopsCredentialFingerprint) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DevopsCredentialFingerprint) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetHash

`func (o *DevopsCredentialFingerprint) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DevopsCredentialFingerprint) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DevopsCredentialFingerprint) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *DevopsCredentialFingerprint) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetUsage

`func (o *DevopsCredentialFingerprint) GetUsage() []DevopsCredentialFingerprintUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DevopsCredentialFingerprint) GetUsageOk() (*[]DevopsCredentialFingerprintUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DevopsCredentialFingerprint) SetUsage(v []DevopsCredentialFingerprintUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DevopsCredentialFingerprint) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


