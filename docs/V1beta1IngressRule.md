# V1beta1IngressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \&quot;host\&quot; part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to    the IP in the Spec of the parent Ingress. 2. The &#x60;:&#x60; delimiter is not respected because ports are not allowed.    Currently the port of an Ingress is implicitly :80 for http and    :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.  Host can be \&quot;precise\&quot; which is a domain name without the terminating dot of a network host (e.g. \&quot;foo.bar.com\&quot;) or \&quot;wildcard\&quot;, which is a domain name prefixed with a single wildcard label (e.g. \&quot;*.foo.com\&quot;). The wildcard character &#39;*&#39; must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host &#x3D;&#x3D; \&quot;*\&quot;). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule. | [optional] 
**Http** | Pointer to [**V1beta1HTTPIngressRuleValue**](V1beta1HTTPIngressRuleValue.md) |  | [optional] 

## Methods

### NewV1beta1IngressRule

`func NewV1beta1IngressRule() *V1beta1IngressRule`

NewV1beta1IngressRule instantiates a new V1beta1IngressRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1IngressRuleWithDefaults

`func NewV1beta1IngressRuleWithDefaults() *V1beta1IngressRule`

NewV1beta1IngressRuleWithDefaults instantiates a new V1beta1IngressRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *V1beta1IngressRule) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1beta1IngressRule) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1beta1IngressRule) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1beta1IngressRule) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHttp

`func (o *V1beta1IngressRule) GetHttp() V1beta1HTTPIngressRuleValue`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *V1beta1IngressRule) GetHttpOk() (*V1beta1HTTPIngressRuleValue, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *V1beta1IngressRule) SetHttp(v V1beta1HTTPIngressRuleValue)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *V1beta1IngressRule) HasHttp() bool`

HasHttp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


