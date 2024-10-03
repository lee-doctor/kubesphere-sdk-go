# V1beta1IngressBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**V1TypedLocalObjectReference**](V1TypedLocalObjectReference.md) |  | [optional] 
**ServiceName** | Pointer to **string** | Specifies the name of the referenced service. | [optional] 
**ServicePort** | Pointer to **string** | Specifies the port of the referenced service. | [optional] 

## Methods

### NewV1beta1IngressBackend

`func NewV1beta1IngressBackend() *V1beta1IngressBackend`

NewV1beta1IngressBackend instantiates a new V1beta1IngressBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1IngressBackendWithDefaults

`func NewV1beta1IngressBackendWithDefaults() *V1beta1IngressBackend`

NewV1beta1IngressBackendWithDefaults instantiates a new V1beta1IngressBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *V1beta1IngressBackend) GetResource() V1TypedLocalObjectReference`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *V1beta1IngressBackend) GetResourceOk() (*V1TypedLocalObjectReference, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *V1beta1IngressBackend) SetResource(v V1TypedLocalObjectReference)`

SetResource sets Resource field to given value.

### HasResource

`func (o *V1beta1IngressBackend) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetServiceName

`func (o *V1beta1IngressBackend) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1beta1IngressBackend) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1beta1IngressBackend) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1beta1IngressBackend) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServicePort

`func (o *V1beta1IngressBackend) GetServicePort() string`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *V1beta1IngressBackend) GetServicePortOk() (*string, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *V1beta1IngressBackend) SetServicePort(v string)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *V1beta1IngressBackend) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


