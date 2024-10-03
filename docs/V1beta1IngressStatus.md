# V1beta1IngressStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | Pointer to [**V1LoadBalancerStatus**](V1LoadBalancerStatus.md) |  | [optional] 

## Methods

### NewV1beta1IngressStatus

`func NewV1beta1IngressStatus() *V1beta1IngressStatus`

NewV1beta1IngressStatus instantiates a new V1beta1IngressStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1IngressStatusWithDefaults

`func NewV1beta1IngressStatusWithDefaults() *V1beta1IngressStatus`

NewV1beta1IngressStatusWithDefaults instantiates a new V1beta1IngressStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *V1beta1IngressStatus) GetLoadBalancer() V1LoadBalancerStatus`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *V1beta1IngressStatus) GetLoadBalancerOk() (*V1LoadBalancerStatus, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *V1beta1IngressStatus) SetLoadBalancer(v V1LoadBalancerStatus)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *V1beta1IngressStatus) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


