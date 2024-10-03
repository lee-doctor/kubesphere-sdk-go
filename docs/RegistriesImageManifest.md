# RegistriesImageManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**RegistriesManifestConfig**](RegistriesManifestConfig.md) |  | [optional] 
**Layers** | Pointer to [**[]RegistriesLayers**](RegistriesLayers.md) | Fields of an item in the layers list. | [optional] 
**MediaType** | Pointer to **string** | The MIME type of the manifest. | [optional] 
**SchemaVersion** | Pointer to **int32** | This field specifies the image manifest schema version as an integer. | [optional] 

## Methods

### NewRegistriesImageManifest

`func NewRegistriesImageManifest() *RegistriesImageManifest`

NewRegistriesImageManifest instantiates a new RegistriesImageManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesImageManifestWithDefaults

`func NewRegistriesImageManifestWithDefaults() *RegistriesImageManifest`

NewRegistriesImageManifestWithDefaults instantiates a new RegistriesImageManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *RegistriesImageManifest) GetConfig() RegistriesManifestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RegistriesImageManifest) GetConfigOk() (*RegistriesManifestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RegistriesImageManifest) SetConfig(v RegistriesManifestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RegistriesImageManifest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLayers

`func (o *RegistriesImageManifest) GetLayers() []RegistriesLayers`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *RegistriesImageManifest) GetLayersOk() (*[]RegistriesLayers, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *RegistriesImageManifest) SetLayers(v []RegistriesLayers)`

SetLayers sets Layers field to given value.

### HasLayers

`func (o *RegistriesImageManifest) HasLayers() bool`

HasLayers returns a boolean if a field has been set.

### GetMediaType

`func (o *RegistriesImageManifest) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *RegistriesImageManifest) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *RegistriesImageManifest) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *RegistriesImageManifest) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *RegistriesImageManifest) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RegistriesImageManifest) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RegistriesImageManifest) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RegistriesImageManifest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


