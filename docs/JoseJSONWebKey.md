# JoseJSONWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**CertificateThumbprintSHA1** | **string** |  | 
**CertificateThumbprintSHA256** | **string** |  | 
**Certificates** | [**[]X509Certificate**](X509Certificate.md) |  | 
**CertificatesURL** | [**UrlURL**](UrlURL.md) |  | 
**Key** | **map[string]interface{}** |  | 
**KeyID** | **string** |  | 
**Use** | **string** |  | 

## Methods

### NewJoseJSONWebKey

`func NewJoseJSONWebKey(algorithm string, certificateThumbprintSHA1 string, certificateThumbprintSHA256 string, certificates []X509Certificate, certificatesURL UrlURL, key map[string]interface{}, keyID string, use string, ) *JoseJSONWebKey`

NewJoseJSONWebKey instantiates a new JoseJSONWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoseJSONWebKeyWithDefaults

`func NewJoseJSONWebKeyWithDefaults() *JoseJSONWebKey`

NewJoseJSONWebKeyWithDefaults instantiates a new JoseJSONWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *JoseJSONWebKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *JoseJSONWebKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *JoseJSONWebKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetCertificateThumbprintSHA1

`func (o *JoseJSONWebKey) GetCertificateThumbprintSHA1() string`

GetCertificateThumbprintSHA1 returns the CertificateThumbprintSHA1 field if non-nil, zero value otherwise.

### GetCertificateThumbprintSHA1Ok

`func (o *JoseJSONWebKey) GetCertificateThumbprintSHA1Ok() (*string, bool)`

GetCertificateThumbprintSHA1Ok returns a tuple with the CertificateThumbprintSHA1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprintSHA1

`func (o *JoseJSONWebKey) SetCertificateThumbprintSHA1(v string)`

SetCertificateThumbprintSHA1 sets CertificateThumbprintSHA1 field to given value.


### GetCertificateThumbprintSHA256

`func (o *JoseJSONWebKey) GetCertificateThumbprintSHA256() string`

GetCertificateThumbprintSHA256 returns the CertificateThumbprintSHA256 field if non-nil, zero value otherwise.

### GetCertificateThumbprintSHA256Ok

`func (o *JoseJSONWebKey) GetCertificateThumbprintSHA256Ok() (*string, bool)`

GetCertificateThumbprintSHA256Ok returns a tuple with the CertificateThumbprintSHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprintSHA256

`func (o *JoseJSONWebKey) SetCertificateThumbprintSHA256(v string)`

SetCertificateThumbprintSHA256 sets CertificateThumbprintSHA256 field to given value.


### GetCertificates

`func (o *JoseJSONWebKey) GetCertificates() []X509Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *JoseJSONWebKey) GetCertificatesOk() (*[]X509Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *JoseJSONWebKey) SetCertificates(v []X509Certificate)`

SetCertificates sets Certificates field to given value.


### GetCertificatesURL

`func (o *JoseJSONWebKey) GetCertificatesURL() UrlURL`

GetCertificatesURL returns the CertificatesURL field if non-nil, zero value otherwise.

### GetCertificatesURLOk

`func (o *JoseJSONWebKey) GetCertificatesURLOk() (*UrlURL, bool)`

GetCertificatesURLOk returns a tuple with the CertificatesURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatesURL

`func (o *JoseJSONWebKey) SetCertificatesURL(v UrlURL)`

SetCertificatesURL sets CertificatesURL field to given value.


### GetKey

`func (o *JoseJSONWebKey) GetKey() map[string]interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JoseJSONWebKey) GetKeyOk() (*map[string]interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JoseJSONWebKey) SetKey(v map[string]interface{})`

SetKey sets Key field to given value.


### GetKeyID

`func (o *JoseJSONWebKey) GetKeyID() string`

GetKeyID returns the KeyID field if non-nil, zero value otherwise.

### GetKeyIDOk

`func (o *JoseJSONWebKey) GetKeyIDOk() (*string, bool)`

GetKeyIDOk returns a tuple with the KeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyID

`func (o *JoseJSONWebKey) SetKeyID(v string)`

SetKeyID sets KeyID field to given value.


### GetUse

`func (o *JoseJSONWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *JoseJSONWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *JoseJSONWebKey) SetUse(v string)`

SetUse sets Use field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


