# X509Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorityKeyId** | **string** |  | 
**BasicConstraintsValid** | **bool** |  | 
**CRLDistributionPoints** | **[]string** |  | 
**DNSNames** | **[]string** |  | 
**EmailAddresses** | **[]string** |  | 
**ExcludedDNSDomains** | **[]string** |  | 
**ExcludedEmailAddresses** | **[]string** |  | 
**ExcludedIPRanges** | [**[]NetIPNet**](NetIPNet.md) |  | 
**ExcludedURIDomains** | **[]string** |  | 
**ExtKeyUsage** | **[]int32** |  | 
**Extensions** | [**[]PkixExtension**](PkixExtension.md) |  | 
**ExtraExtensions** | [**[]PkixExtension**](PkixExtension.md) |  | 
**IPAddresses** | **[]string** |  | 
**IsCA** | **bool** |  | 
**Issuer** | [**PkixName**](PkixName.md) |  | 
**IssuingCertificateURL** | **[]string** |  | 
**KeyUsage** | **int32** |  | 
**MaxPathLen** | **int32** |  | 
**MaxPathLenZero** | **bool** |  | 
**NotAfter** | **time.Time** |  | 
**NotBefore** | **time.Time** |  | 
**OCSPServer** | **[]string** |  | 
**PermittedDNSDomains** | **[]string** |  | 
**PermittedDNSDomainsCritical** | **bool** |  | 
**PermittedEmailAddresses** | **[]string** |  | 
**PermittedIPRanges** | [**[]NetIPNet**](NetIPNet.md) |  | 
**PermittedURIDomains** | **[]string** |  | 
**Policies** | [**[]X509OID**](X509OID.md) |  | 
**PolicyIdentifiers** | **[][]int32** |  | 
**PublicKey** | **map[string]interface{}** |  | 
**PublicKeyAlgorithm** | **int32** |  | 
**Raw** | **string** |  | 
**RawIssuer** | **string** |  | 
**RawSubject** | **string** |  | 
**RawSubjectPublicKeyInfo** | **string** |  | 
**RawTBSCertificate** | **string** |  | 
**SerialNumber** | **string** |  | 
**Signature** | **string** |  | 
**SignatureAlgorithm** | **int32** |  | 
**Subject** | [**PkixName**](PkixName.md) |  | 
**SubjectKeyId** | **string** |  | 
**URIs** | [**[]UrlURL**](UrlURL.md) |  | 
**UnhandledCriticalExtensions** | **[][]int32** |  | 
**UnknownExtKeyUsage** | **[][]int32** |  | 
**Version** | **int32** |  | 

## Methods

### NewX509Certificate

`func NewX509Certificate(authorityKeyId string, basicConstraintsValid bool, cRLDistributionPoints []string, dNSNames []string, emailAddresses []string, excludedDNSDomains []string, excludedEmailAddresses []string, excludedIPRanges []NetIPNet, excludedURIDomains []string, extKeyUsage []int32, extensions []PkixExtension, extraExtensions []PkixExtension, iPAddresses []string, isCA bool, issuer PkixName, issuingCertificateURL []string, keyUsage int32, maxPathLen int32, maxPathLenZero bool, notAfter time.Time, notBefore time.Time, oCSPServer []string, permittedDNSDomains []string, permittedDNSDomainsCritical bool, permittedEmailAddresses []string, permittedIPRanges []NetIPNet, permittedURIDomains []string, policies []X509OID, policyIdentifiers [][]int32, publicKey map[string]interface{}, publicKeyAlgorithm int32, raw string, rawIssuer string, rawSubject string, rawSubjectPublicKeyInfo string, rawTBSCertificate string, serialNumber string, signature string, signatureAlgorithm int32, subject PkixName, subjectKeyId string, uRIs []UrlURL, unhandledCriticalExtensions [][]int32, unknownExtKeyUsage [][]int32, version int32, ) *X509Certificate`

NewX509Certificate instantiates a new X509Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateWithDefaults

`func NewX509CertificateWithDefaults() *X509Certificate`

NewX509CertificateWithDefaults instantiates a new X509Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorityKeyId

`func (o *X509Certificate) GetAuthorityKeyId() string`

GetAuthorityKeyId returns the AuthorityKeyId field if non-nil, zero value otherwise.

### GetAuthorityKeyIdOk

`func (o *X509Certificate) GetAuthorityKeyIdOk() (*string, bool)`

GetAuthorityKeyIdOk returns a tuple with the AuthorityKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyId

`func (o *X509Certificate) SetAuthorityKeyId(v string)`

SetAuthorityKeyId sets AuthorityKeyId field to given value.


### GetBasicConstraintsValid

`func (o *X509Certificate) GetBasicConstraintsValid() bool`

GetBasicConstraintsValid returns the BasicConstraintsValid field if non-nil, zero value otherwise.

### GetBasicConstraintsValidOk

`func (o *X509Certificate) GetBasicConstraintsValidOk() (*bool, bool)`

GetBasicConstraintsValidOk returns a tuple with the BasicConstraintsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraintsValid

`func (o *X509Certificate) SetBasicConstraintsValid(v bool)`

SetBasicConstraintsValid sets BasicConstraintsValid field to given value.


### GetCRLDistributionPoints

`func (o *X509Certificate) GetCRLDistributionPoints() []string`

GetCRLDistributionPoints returns the CRLDistributionPoints field if non-nil, zero value otherwise.

### GetCRLDistributionPointsOk

`func (o *X509Certificate) GetCRLDistributionPointsOk() (*[]string, bool)`

GetCRLDistributionPointsOk returns a tuple with the CRLDistributionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCRLDistributionPoints

`func (o *X509Certificate) SetCRLDistributionPoints(v []string)`

SetCRLDistributionPoints sets CRLDistributionPoints field to given value.


### GetDNSNames

`func (o *X509Certificate) GetDNSNames() []string`

GetDNSNames returns the DNSNames field if non-nil, zero value otherwise.

### GetDNSNamesOk

`func (o *X509Certificate) GetDNSNamesOk() (*[]string, bool)`

GetDNSNamesOk returns a tuple with the DNSNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSNames

`func (o *X509Certificate) SetDNSNames(v []string)`

SetDNSNames sets DNSNames field to given value.


### GetEmailAddresses

`func (o *X509Certificate) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *X509Certificate) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *X509Certificate) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.


### GetExcludedDNSDomains

`func (o *X509Certificate) GetExcludedDNSDomains() []string`

GetExcludedDNSDomains returns the ExcludedDNSDomains field if non-nil, zero value otherwise.

### GetExcludedDNSDomainsOk

`func (o *X509Certificate) GetExcludedDNSDomainsOk() (*[]string, bool)`

GetExcludedDNSDomainsOk returns a tuple with the ExcludedDNSDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDNSDomains

`func (o *X509Certificate) SetExcludedDNSDomains(v []string)`

SetExcludedDNSDomains sets ExcludedDNSDomains field to given value.


### GetExcludedEmailAddresses

`func (o *X509Certificate) GetExcludedEmailAddresses() []string`

GetExcludedEmailAddresses returns the ExcludedEmailAddresses field if non-nil, zero value otherwise.

### GetExcludedEmailAddressesOk

`func (o *X509Certificate) GetExcludedEmailAddressesOk() (*[]string, bool)`

GetExcludedEmailAddressesOk returns a tuple with the ExcludedEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEmailAddresses

`func (o *X509Certificate) SetExcludedEmailAddresses(v []string)`

SetExcludedEmailAddresses sets ExcludedEmailAddresses field to given value.


### GetExcludedIPRanges

`func (o *X509Certificate) GetExcludedIPRanges() []NetIPNet`

GetExcludedIPRanges returns the ExcludedIPRanges field if non-nil, zero value otherwise.

### GetExcludedIPRangesOk

`func (o *X509Certificate) GetExcludedIPRangesOk() (*[]NetIPNet, bool)`

GetExcludedIPRangesOk returns a tuple with the ExcludedIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedIPRanges

`func (o *X509Certificate) SetExcludedIPRanges(v []NetIPNet)`

SetExcludedIPRanges sets ExcludedIPRanges field to given value.


### GetExcludedURIDomains

`func (o *X509Certificate) GetExcludedURIDomains() []string`

GetExcludedURIDomains returns the ExcludedURIDomains field if non-nil, zero value otherwise.

### GetExcludedURIDomainsOk

`func (o *X509Certificate) GetExcludedURIDomainsOk() (*[]string, bool)`

GetExcludedURIDomainsOk returns a tuple with the ExcludedURIDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedURIDomains

`func (o *X509Certificate) SetExcludedURIDomains(v []string)`

SetExcludedURIDomains sets ExcludedURIDomains field to given value.


### GetExtKeyUsage

`func (o *X509Certificate) GetExtKeyUsage() []int32`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *X509Certificate) GetExtKeyUsageOk() (*[]int32, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *X509Certificate) SetExtKeyUsage(v []int32)`

SetExtKeyUsage sets ExtKeyUsage field to given value.


### GetExtensions

`func (o *X509Certificate) GetExtensions() []PkixExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *X509Certificate) GetExtensionsOk() (*[]PkixExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *X509Certificate) SetExtensions(v []PkixExtension)`

SetExtensions sets Extensions field to given value.


### GetExtraExtensions

`func (o *X509Certificate) GetExtraExtensions() []PkixExtension`

GetExtraExtensions returns the ExtraExtensions field if non-nil, zero value otherwise.

### GetExtraExtensionsOk

`func (o *X509Certificate) GetExtraExtensionsOk() (*[]PkixExtension, bool)`

GetExtraExtensionsOk returns a tuple with the ExtraExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraExtensions

`func (o *X509Certificate) SetExtraExtensions(v []PkixExtension)`

SetExtraExtensions sets ExtraExtensions field to given value.


### GetIPAddresses

`func (o *X509Certificate) GetIPAddresses() []string`

GetIPAddresses returns the IPAddresses field if non-nil, zero value otherwise.

### GetIPAddressesOk

`func (o *X509Certificate) GetIPAddressesOk() (*[]string, bool)`

GetIPAddressesOk returns a tuple with the IPAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddresses

`func (o *X509Certificate) SetIPAddresses(v []string)`

SetIPAddresses sets IPAddresses field to given value.


### GetIsCA

`func (o *X509Certificate) GetIsCA() bool`

GetIsCA returns the IsCA field if non-nil, zero value otherwise.

### GetIsCAOk

`func (o *X509Certificate) GetIsCAOk() (*bool, bool)`

GetIsCAOk returns a tuple with the IsCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCA

`func (o *X509Certificate) SetIsCA(v bool)`

SetIsCA sets IsCA field to given value.


### GetIssuer

`func (o *X509Certificate) GetIssuer() PkixName`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *X509Certificate) GetIssuerOk() (*PkixName, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *X509Certificate) SetIssuer(v PkixName)`

SetIssuer sets Issuer field to given value.


### GetIssuingCertificateURL

`func (o *X509Certificate) GetIssuingCertificateURL() []string`

GetIssuingCertificateURL returns the IssuingCertificateURL field if non-nil, zero value otherwise.

### GetIssuingCertificateURLOk

`func (o *X509Certificate) GetIssuingCertificateURLOk() (*[]string, bool)`

GetIssuingCertificateURLOk returns a tuple with the IssuingCertificateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCertificateURL

`func (o *X509Certificate) SetIssuingCertificateURL(v []string)`

SetIssuingCertificateURL sets IssuingCertificateURL field to given value.


### GetKeyUsage

`func (o *X509Certificate) GetKeyUsage() int32`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *X509Certificate) GetKeyUsageOk() (*int32, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *X509Certificate) SetKeyUsage(v int32)`

SetKeyUsage sets KeyUsage field to given value.


### GetMaxPathLen

`func (o *X509Certificate) GetMaxPathLen() int32`

GetMaxPathLen returns the MaxPathLen field if non-nil, zero value otherwise.

### GetMaxPathLenOk

`func (o *X509Certificate) GetMaxPathLenOk() (*int32, bool)`

GetMaxPathLenOk returns a tuple with the MaxPathLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLen

`func (o *X509Certificate) SetMaxPathLen(v int32)`

SetMaxPathLen sets MaxPathLen field to given value.


### GetMaxPathLenZero

`func (o *X509Certificate) GetMaxPathLenZero() bool`

GetMaxPathLenZero returns the MaxPathLenZero field if non-nil, zero value otherwise.

### GetMaxPathLenZeroOk

`func (o *X509Certificate) GetMaxPathLenZeroOk() (*bool, bool)`

GetMaxPathLenZeroOk returns a tuple with the MaxPathLenZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLenZero

`func (o *X509Certificate) SetMaxPathLenZero(v bool)`

SetMaxPathLenZero sets MaxPathLenZero field to given value.


### GetNotAfter

`func (o *X509Certificate) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *X509Certificate) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.


### GetNotBefore

`func (o *X509Certificate) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *X509Certificate) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.


### GetOCSPServer

`func (o *X509Certificate) GetOCSPServer() []string`

GetOCSPServer returns the OCSPServer field if non-nil, zero value otherwise.

### GetOCSPServerOk

`func (o *X509Certificate) GetOCSPServerOk() (*[]string, bool)`

GetOCSPServerOk returns a tuple with the OCSPServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCSPServer

`func (o *X509Certificate) SetOCSPServer(v []string)`

SetOCSPServer sets OCSPServer field to given value.


### GetPermittedDNSDomains

`func (o *X509Certificate) GetPermittedDNSDomains() []string`

GetPermittedDNSDomains returns the PermittedDNSDomains field if non-nil, zero value otherwise.

### GetPermittedDNSDomainsOk

`func (o *X509Certificate) GetPermittedDNSDomainsOk() (*[]string, bool)`

GetPermittedDNSDomainsOk returns a tuple with the PermittedDNSDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedDNSDomains

`func (o *X509Certificate) SetPermittedDNSDomains(v []string)`

SetPermittedDNSDomains sets PermittedDNSDomains field to given value.


### GetPermittedDNSDomainsCritical

`func (o *X509Certificate) GetPermittedDNSDomainsCritical() bool`

GetPermittedDNSDomainsCritical returns the PermittedDNSDomainsCritical field if non-nil, zero value otherwise.

### GetPermittedDNSDomainsCriticalOk

`func (o *X509Certificate) GetPermittedDNSDomainsCriticalOk() (*bool, bool)`

GetPermittedDNSDomainsCriticalOk returns a tuple with the PermittedDNSDomainsCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedDNSDomainsCritical

`func (o *X509Certificate) SetPermittedDNSDomainsCritical(v bool)`

SetPermittedDNSDomainsCritical sets PermittedDNSDomainsCritical field to given value.


### GetPermittedEmailAddresses

`func (o *X509Certificate) GetPermittedEmailAddresses() []string`

GetPermittedEmailAddresses returns the PermittedEmailAddresses field if non-nil, zero value otherwise.

### GetPermittedEmailAddressesOk

`func (o *X509Certificate) GetPermittedEmailAddressesOk() (*[]string, bool)`

GetPermittedEmailAddressesOk returns a tuple with the PermittedEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedEmailAddresses

`func (o *X509Certificate) SetPermittedEmailAddresses(v []string)`

SetPermittedEmailAddresses sets PermittedEmailAddresses field to given value.


### GetPermittedIPRanges

`func (o *X509Certificate) GetPermittedIPRanges() []NetIPNet`

GetPermittedIPRanges returns the PermittedIPRanges field if non-nil, zero value otherwise.

### GetPermittedIPRangesOk

`func (o *X509Certificate) GetPermittedIPRangesOk() (*[]NetIPNet, bool)`

GetPermittedIPRangesOk returns a tuple with the PermittedIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedIPRanges

`func (o *X509Certificate) SetPermittedIPRanges(v []NetIPNet)`

SetPermittedIPRanges sets PermittedIPRanges field to given value.


### GetPermittedURIDomains

`func (o *X509Certificate) GetPermittedURIDomains() []string`

GetPermittedURIDomains returns the PermittedURIDomains field if non-nil, zero value otherwise.

### GetPermittedURIDomainsOk

`func (o *X509Certificate) GetPermittedURIDomainsOk() (*[]string, bool)`

GetPermittedURIDomainsOk returns a tuple with the PermittedURIDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedURIDomains

`func (o *X509Certificate) SetPermittedURIDomains(v []string)`

SetPermittedURIDomains sets PermittedURIDomains field to given value.


### GetPolicies

`func (o *X509Certificate) GetPolicies() []X509OID`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *X509Certificate) GetPoliciesOk() (*[]X509OID, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *X509Certificate) SetPolicies(v []X509OID)`

SetPolicies sets Policies field to given value.


### GetPolicyIdentifiers

`func (o *X509Certificate) GetPolicyIdentifiers() [][]int32`

GetPolicyIdentifiers returns the PolicyIdentifiers field if non-nil, zero value otherwise.

### GetPolicyIdentifiersOk

`func (o *X509Certificate) GetPolicyIdentifiersOk() (*[][]int32, bool)`

GetPolicyIdentifiersOk returns a tuple with the PolicyIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentifiers

`func (o *X509Certificate) SetPolicyIdentifiers(v [][]int32)`

SetPolicyIdentifiers sets PolicyIdentifiers field to given value.


### GetPublicKey

`func (o *X509Certificate) GetPublicKey() map[string]interface{}`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *X509Certificate) GetPublicKeyOk() (*map[string]interface{}, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *X509Certificate) SetPublicKey(v map[string]interface{})`

SetPublicKey sets PublicKey field to given value.


### GetPublicKeyAlgorithm

`func (o *X509Certificate) GetPublicKeyAlgorithm() int32`

GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmOk

`func (o *X509Certificate) GetPublicKeyAlgorithmOk() (*int32, bool)`

GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithm

`func (o *X509Certificate) SetPublicKeyAlgorithm(v int32)`

SetPublicKeyAlgorithm sets PublicKeyAlgorithm field to given value.


### GetRaw

`func (o *X509Certificate) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *X509Certificate) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *X509Certificate) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetRawIssuer

`func (o *X509Certificate) GetRawIssuer() string`

GetRawIssuer returns the RawIssuer field if non-nil, zero value otherwise.

### GetRawIssuerOk

`func (o *X509Certificate) GetRawIssuerOk() (*string, bool)`

GetRawIssuerOk returns a tuple with the RawIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawIssuer

`func (o *X509Certificate) SetRawIssuer(v string)`

SetRawIssuer sets RawIssuer field to given value.


### GetRawSubject

`func (o *X509Certificate) GetRawSubject() string`

GetRawSubject returns the RawSubject field if non-nil, zero value otherwise.

### GetRawSubjectOk

`func (o *X509Certificate) GetRawSubjectOk() (*string, bool)`

GetRawSubjectOk returns a tuple with the RawSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSubject

`func (o *X509Certificate) SetRawSubject(v string)`

SetRawSubject sets RawSubject field to given value.


### GetRawSubjectPublicKeyInfo

`func (o *X509Certificate) GetRawSubjectPublicKeyInfo() string`

GetRawSubjectPublicKeyInfo returns the RawSubjectPublicKeyInfo field if non-nil, zero value otherwise.

### GetRawSubjectPublicKeyInfoOk

`func (o *X509Certificate) GetRawSubjectPublicKeyInfoOk() (*string, bool)`

GetRawSubjectPublicKeyInfoOk returns a tuple with the RawSubjectPublicKeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSubjectPublicKeyInfo

`func (o *X509Certificate) SetRawSubjectPublicKeyInfo(v string)`

SetRawSubjectPublicKeyInfo sets RawSubjectPublicKeyInfo field to given value.


### GetRawTBSCertificate

`func (o *X509Certificate) GetRawTBSCertificate() string`

GetRawTBSCertificate returns the RawTBSCertificate field if non-nil, zero value otherwise.

### GetRawTBSCertificateOk

`func (o *X509Certificate) GetRawTBSCertificateOk() (*string, bool)`

GetRawTBSCertificateOk returns a tuple with the RawTBSCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTBSCertificate

`func (o *X509Certificate) SetRawTBSCertificate(v string)`

SetRawTBSCertificate sets RawTBSCertificate field to given value.


### GetSerialNumber

`func (o *X509Certificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *X509Certificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *X509Certificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetSignature

`func (o *X509Certificate) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *X509Certificate) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *X509Certificate) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSignatureAlgorithm

`func (o *X509Certificate) GetSignatureAlgorithm() int32`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *X509Certificate) GetSignatureAlgorithmOk() (*int32, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *X509Certificate) SetSignatureAlgorithm(v int32)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetSubject

`func (o *X509Certificate) GetSubject() PkixName`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *X509Certificate) GetSubjectOk() (*PkixName, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *X509Certificate) SetSubject(v PkixName)`

SetSubject sets Subject field to given value.


### GetSubjectKeyId

`func (o *X509Certificate) GetSubjectKeyId() string`

GetSubjectKeyId returns the SubjectKeyId field if non-nil, zero value otherwise.

### GetSubjectKeyIdOk

`func (o *X509Certificate) GetSubjectKeyIdOk() (*string, bool)`

GetSubjectKeyIdOk returns a tuple with the SubjectKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeyId

`func (o *X509Certificate) SetSubjectKeyId(v string)`

SetSubjectKeyId sets SubjectKeyId field to given value.


### GetURIs

`func (o *X509Certificate) GetURIs() []UrlURL`

GetURIs returns the URIs field if non-nil, zero value otherwise.

### GetURIsOk

`func (o *X509Certificate) GetURIsOk() (*[]UrlURL, bool)`

GetURIsOk returns a tuple with the URIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURIs

`func (o *X509Certificate) SetURIs(v []UrlURL)`

SetURIs sets URIs field to given value.


### GetUnhandledCriticalExtensions

`func (o *X509Certificate) GetUnhandledCriticalExtensions() [][]int32`

GetUnhandledCriticalExtensions returns the UnhandledCriticalExtensions field if non-nil, zero value otherwise.

### GetUnhandledCriticalExtensionsOk

`func (o *X509Certificate) GetUnhandledCriticalExtensionsOk() (*[][]int32, bool)`

GetUnhandledCriticalExtensionsOk returns a tuple with the UnhandledCriticalExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhandledCriticalExtensions

`func (o *X509Certificate) SetUnhandledCriticalExtensions(v [][]int32)`

SetUnhandledCriticalExtensions sets UnhandledCriticalExtensions field to given value.


### GetUnknownExtKeyUsage

`func (o *X509Certificate) GetUnknownExtKeyUsage() [][]int32`

GetUnknownExtKeyUsage returns the UnknownExtKeyUsage field if non-nil, zero value otherwise.

### GetUnknownExtKeyUsageOk

`func (o *X509Certificate) GetUnknownExtKeyUsageOk() (*[][]int32, bool)`

GetUnknownExtKeyUsageOk returns a tuple with the UnknownExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownExtKeyUsage

`func (o *X509Certificate) SetUnknownExtKeyUsage(v [][]int32)`

SetUnknownExtKeyUsage sets UnknownExtKeyUsage field to given value.


### GetVersion

`func (o *X509Certificate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *X509Certificate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *X509Certificate) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


