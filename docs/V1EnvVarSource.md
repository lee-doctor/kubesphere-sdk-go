# V1EnvVarSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapKeyRef** | Pointer to [**V1ConfigMapKeySelector**](V1ConfigMapKeySelector.md) |  | [optional] 
**FieldRef** | Pointer to [**V1ObjectFieldSelector**](V1ObjectFieldSelector.md) |  | [optional] 
**ResourceFieldRef** | Pointer to [**V1ResourceFieldSelector**](V1ResourceFieldSelector.md) |  | [optional] 
**SecretKeyRef** | Pointer to [**V1SecretKeySelector**](V1SecretKeySelector.md) |  | [optional] 

## Methods

### NewV1EnvVarSource

`func NewV1EnvVarSource() *V1EnvVarSource`

NewV1EnvVarSource instantiates a new V1EnvVarSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnvVarSourceWithDefaults

`func NewV1EnvVarSourceWithDefaults() *V1EnvVarSource`

NewV1EnvVarSourceWithDefaults instantiates a new V1EnvVarSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMapKeyRef

`func (o *V1EnvVarSource) GetConfigMapKeyRef() V1ConfigMapKeySelector`

GetConfigMapKeyRef returns the ConfigMapKeyRef field if non-nil, zero value otherwise.

### GetConfigMapKeyRefOk

`func (o *V1EnvVarSource) GetConfigMapKeyRefOk() (*V1ConfigMapKeySelector, bool)`

GetConfigMapKeyRefOk returns a tuple with the ConfigMapKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapKeyRef

`func (o *V1EnvVarSource) SetConfigMapKeyRef(v V1ConfigMapKeySelector)`

SetConfigMapKeyRef sets ConfigMapKeyRef field to given value.

### HasConfigMapKeyRef

`func (o *V1EnvVarSource) HasConfigMapKeyRef() bool`

HasConfigMapKeyRef returns a boolean if a field has been set.

### GetFieldRef

`func (o *V1EnvVarSource) GetFieldRef() V1ObjectFieldSelector`

GetFieldRef returns the FieldRef field if non-nil, zero value otherwise.

### GetFieldRefOk

`func (o *V1EnvVarSource) GetFieldRefOk() (*V1ObjectFieldSelector, bool)`

GetFieldRefOk returns a tuple with the FieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRef

`func (o *V1EnvVarSource) SetFieldRef(v V1ObjectFieldSelector)`

SetFieldRef sets FieldRef field to given value.

### HasFieldRef

`func (o *V1EnvVarSource) HasFieldRef() bool`

HasFieldRef returns a boolean if a field has been set.

### GetResourceFieldRef

`func (o *V1EnvVarSource) GetResourceFieldRef() V1ResourceFieldSelector`

GetResourceFieldRef returns the ResourceFieldRef field if non-nil, zero value otherwise.

### GetResourceFieldRefOk

`func (o *V1EnvVarSource) GetResourceFieldRefOk() (*V1ResourceFieldSelector, bool)`

GetResourceFieldRefOk returns a tuple with the ResourceFieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceFieldRef

`func (o *V1EnvVarSource) SetResourceFieldRef(v V1ResourceFieldSelector)`

SetResourceFieldRef sets ResourceFieldRef field to given value.

### HasResourceFieldRef

`func (o *V1EnvVarSource) HasResourceFieldRef() bool`

HasResourceFieldRef returns a boolean if a field has been set.

### GetSecretKeyRef

`func (o *V1EnvVarSource) GetSecretKeyRef() V1SecretKeySelector`

GetSecretKeyRef returns the SecretKeyRef field if non-nil, zero value otherwise.

### GetSecretKeyRefOk

`func (o *V1EnvVarSource) GetSecretKeyRefOk() (*V1SecretKeySelector, bool)`

GetSecretKeyRefOk returns a tuple with the SecretKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyRef

`func (o *V1EnvVarSource) SetSecretKeyRef(v V1SecretKeySelector)`

SetSecretKeyRef sets SecretKeyRef field to given value.

### HasSecretKeyRef

`func (o *V1EnvVarSource) HasSecretKeyRef() bool`

HasSecretKeyRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


