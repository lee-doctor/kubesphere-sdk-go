# RegistriesManifestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | The digest of the content, as defined by the Registry V2 HTTP API Specificiation. Reference https://docs.docker.com/registry/spec/api/#digest-parameter | [optional] 
**MediaType** | Pointer to **string** | The MIME type of the image. | [optional] 
**Size** | Pointer to **int32** | The size in bytes of the image. | [optional] 

## Methods

### NewRegistriesManifestConfig

`func NewRegistriesManifestConfig() *RegistriesManifestConfig`

NewRegistriesManifestConfig instantiates a new RegistriesManifestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesManifestConfigWithDefaults

`func NewRegistriesManifestConfigWithDefaults() *RegistriesManifestConfig`

NewRegistriesManifestConfigWithDefaults instantiates a new RegistriesManifestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *RegistriesManifestConfig) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RegistriesManifestConfig) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RegistriesManifestConfig) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *RegistriesManifestConfig) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetMediaType

`func (o *RegistriesManifestConfig) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *RegistriesManifestConfig) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *RegistriesManifestConfig) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *RegistriesManifestConfig) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetSize

`func (o *RegistriesManifestConfig) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RegistriesManifestConfig) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RegistriesManifestConfig) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RegistriesManifestConfig) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


