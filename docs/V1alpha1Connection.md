# V1alpha1Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalKubernetesAPIEndpoint** | Pointer to **string** |  | [optional] 
**Kubeconfig** | Pointer to **string** |  | [optional] 
**KubernetesAPIEndpoint** | Pointer to **string** |  | [optional] 
**KubernetesAPIServerPort** | Pointer to **int32** |  | [optional] 
**KubesphereAPIEndpoint** | Pointer to **string** |  | [optional] 
**KubesphereAPIServerPort** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1Connection

`func NewV1alpha1Connection() *V1alpha1Connection`

NewV1alpha1Connection instantiates a new V1alpha1Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ConnectionWithDefaults

`func NewV1alpha1ConnectionWithDefaults() *V1alpha1Connection`

NewV1alpha1ConnectionWithDefaults instantiates a new V1alpha1Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalKubernetesAPIEndpoint

`func (o *V1alpha1Connection) GetExternalKubernetesAPIEndpoint() string`

GetExternalKubernetesAPIEndpoint returns the ExternalKubernetesAPIEndpoint field if non-nil, zero value otherwise.

### GetExternalKubernetesAPIEndpointOk

`func (o *V1alpha1Connection) GetExternalKubernetesAPIEndpointOk() (*string, bool)`

GetExternalKubernetesAPIEndpointOk returns a tuple with the ExternalKubernetesAPIEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKubernetesAPIEndpoint

`func (o *V1alpha1Connection) SetExternalKubernetesAPIEndpoint(v string)`

SetExternalKubernetesAPIEndpoint sets ExternalKubernetesAPIEndpoint field to given value.

### HasExternalKubernetesAPIEndpoint

`func (o *V1alpha1Connection) HasExternalKubernetesAPIEndpoint() bool`

HasExternalKubernetesAPIEndpoint returns a boolean if a field has been set.

### GetKubeconfig

`func (o *V1alpha1Connection) GetKubeconfig() string`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *V1alpha1Connection) GetKubeconfigOk() (*string, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *V1alpha1Connection) SetKubeconfig(v string)`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *V1alpha1Connection) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetKubernetesAPIEndpoint

`func (o *V1alpha1Connection) GetKubernetesAPIEndpoint() string`

GetKubernetesAPIEndpoint returns the KubernetesAPIEndpoint field if non-nil, zero value otherwise.

### GetKubernetesAPIEndpointOk

`func (o *V1alpha1Connection) GetKubernetesAPIEndpointOk() (*string, bool)`

GetKubernetesAPIEndpointOk returns a tuple with the KubernetesAPIEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesAPIEndpoint

`func (o *V1alpha1Connection) SetKubernetesAPIEndpoint(v string)`

SetKubernetesAPIEndpoint sets KubernetesAPIEndpoint field to given value.

### HasKubernetesAPIEndpoint

`func (o *V1alpha1Connection) HasKubernetesAPIEndpoint() bool`

HasKubernetesAPIEndpoint returns a boolean if a field has been set.

### GetKubernetesAPIServerPort

`func (o *V1alpha1Connection) GetKubernetesAPIServerPort() int32`

GetKubernetesAPIServerPort returns the KubernetesAPIServerPort field if non-nil, zero value otherwise.

### GetKubernetesAPIServerPortOk

`func (o *V1alpha1Connection) GetKubernetesAPIServerPortOk() (*int32, bool)`

GetKubernetesAPIServerPortOk returns a tuple with the KubernetesAPIServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesAPIServerPort

`func (o *V1alpha1Connection) SetKubernetesAPIServerPort(v int32)`

SetKubernetesAPIServerPort sets KubernetesAPIServerPort field to given value.

### HasKubernetesAPIServerPort

`func (o *V1alpha1Connection) HasKubernetesAPIServerPort() bool`

HasKubernetesAPIServerPort returns a boolean if a field has been set.

### GetKubesphereAPIEndpoint

`func (o *V1alpha1Connection) GetKubesphereAPIEndpoint() string`

GetKubesphereAPIEndpoint returns the KubesphereAPIEndpoint field if non-nil, zero value otherwise.

### GetKubesphereAPIEndpointOk

`func (o *V1alpha1Connection) GetKubesphereAPIEndpointOk() (*string, bool)`

GetKubesphereAPIEndpointOk returns a tuple with the KubesphereAPIEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesphereAPIEndpoint

`func (o *V1alpha1Connection) SetKubesphereAPIEndpoint(v string)`

SetKubesphereAPIEndpoint sets KubesphereAPIEndpoint field to given value.

### HasKubesphereAPIEndpoint

`func (o *V1alpha1Connection) HasKubesphereAPIEndpoint() bool`

HasKubesphereAPIEndpoint returns a boolean if a field has been set.

### GetKubesphereAPIServerPort

`func (o *V1alpha1Connection) GetKubesphereAPIServerPort() int32`

GetKubesphereAPIServerPort returns the KubesphereAPIServerPort field if non-nil, zero value otherwise.

### GetKubesphereAPIServerPortOk

`func (o *V1alpha1Connection) GetKubesphereAPIServerPortOk() (*int32, bool)`

GetKubesphereAPIServerPortOk returns a tuple with the KubesphereAPIServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesphereAPIServerPort

`func (o *V1alpha1Connection) SetKubesphereAPIServerPort(v int32)`

SetKubesphereAPIServerPort sets KubesphereAPIServerPort field to given value.

### HasKubesphereAPIServerPort

`func (o *V1alpha1Connection) HasKubesphereAPIServerPort() bool`

HasKubesphereAPIServerPort returns a boolean if a field has been set.

### GetToken

`func (o *V1alpha1Connection) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1alpha1Connection) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1alpha1Connection) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1alpha1Connection) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *V1alpha1Connection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1Connection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1Connection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1alpha1Connection) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


