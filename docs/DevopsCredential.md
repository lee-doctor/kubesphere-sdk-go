# DevopsCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Credential&#39;s description&#39; | [optional] 
**DisplayName** | Pointer to **string** | Credential&#39;s display name | [optional] 
**Domain** | Pointer to **string** | Credential&#39;s domain,In ks we only use the default domain, default &#39;_&#39;&#39; | [optional] 
**Fingerprint** | Pointer to [**DevopsCredentialFingerprint**](DevopsCredentialFingerprint.md) |  | [optional] 
**Id** | **string** | Id of Credential, e.g. dockerhub-id | 
**Type** | **string** | Type of Credential, e.g. ssh/kubeconfig | 

## Methods

### NewDevopsCredential

`func NewDevopsCredential(id string, type_ string, ) *DevopsCredential`

NewDevopsCredential instantiates a new DevopsCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsCredentialWithDefaults

`func NewDevopsCredentialWithDefaults() *DevopsCredential`

NewDevopsCredentialWithDefaults instantiates a new DevopsCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DevopsCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DevopsCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DevopsCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DevopsCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsCredential) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsCredential) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsCredential) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsCredential) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomain

`func (o *DevopsCredential) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DevopsCredential) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DevopsCredential) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DevopsCredential) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFingerprint

`func (o *DevopsCredential) GetFingerprint() DevopsCredentialFingerprint`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *DevopsCredential) GetFingerprintOk() (*DevopsCredentialFingerprint, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *DevopsCredential) SetFingerprint(v DevopsCredentialFingerprint)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *DevopsCredential) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetId

`func (o *DevopsCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevopsCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevopsCredential) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DevopsCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevopsCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevopsCredential) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


