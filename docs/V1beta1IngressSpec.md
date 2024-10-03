# V1beta1IngressSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | Pointer to [**V1beta1IngressBackend**](V1beta1IngressBackend.md) |  | [optional] 
**IngressClassName** | Pointer to **string** | IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated &#x60;kubernetes.io/ingress.class&#x60; annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation. | [optional] 
**Rules** | Pointer to [**[]V1beta1IngressRule**](V1beta1IngressRule.md) | A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend. | [optional] 
**Tls** | Pointer to [**[]V1beta1IngressTLS**](V1beta1IngressTLS.md) | TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. | [optional] 

## Methods

### NewV1beta1IngressSpec

`func NewV1beta1IngressSpec() *V1beta1IngressSpec`

NewV1beta1IngressSpec instantiates a new V1beta1IngressSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1IngressSpecWithDefaults

`func NewV1beta1IngressSpecWithDefaults() *V1beta1IngressSpec`

NewV1beta1IngressSpecWithDefaults instantiates a new V1beta1IngressSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *V1beta1IngressSpec) GetBackend() V1beta1IngressBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *V1beta1IngressSpec) GetBackendOk() (*V1beta1IngressBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *V1beta1IngressSpec) SetBackend(v V1beta1IngressBackend)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *V1beta1IngressSpec) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetIngressClassName

`func (o *V1beta1IngressSpec) GetIngressClassName() string`

GetIngressClassName returns the IngressClassName field if non-nil, zero value otherwise.

### GetIngressClassNameOk

`func (o *V1beta1IngressSpec) GetIngressClassNameOk() (*string, bool)`

GetIngressClassNameOk returns a tuple with the IngressClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressClassName

`func (o *V1beta1IngressSpec) SetIngressClassName(v string)`

SetIngressClassName sets IngressClassName field to given value.

### HasIngressClassName

`func (o *V1beta1IngressSpec) HasIngressClassName() bool`

HasIngressClassName returns a boolean if a field has been set.

### GetRules

`func (o *V1beta1IngressSpec) GetRules() []V1beta1IngressRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *V1beta1IngressSpec) GetRulesOk() (*[]V1beta1IngressRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *V1beta1IngressSpec) SetRules(v []V1beta1IngressRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *V1beta1IngressSpec) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTls

`func (o *V1beta1IngressSpec) GetTls() []V1beta1IngressTLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *V1beta1IngressSpec) GetTlsOk() (*[]V1beta1IngressTLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *V1beta1IngressSpec) SetTls(v []V1beta1IngressTLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *V1beta1IngressSpec) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


