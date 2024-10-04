# V1alpha2ResourceQuotaStatusByNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hard** | Pointer to [**map[string]ResourceQuantity**](ResourceQuantity.md) | Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/ | [optional] 
**Namespace** | **string** |  | 
**Used** | Pointer to [**map[string]ResourceQuantity**](ResourceQuantity.md) | Used is the current observed total usage of the resource in the namespace. | [optional] 

## Methods

### NewV1alpha2ResourceQuotaStatusByNamespace

`func NewV1alpha2ResourceQuotaStatusByNamespace(namespace string, ) *V1alpha2ResourceQuotaStatusByNamespace`

NewV1alpha2ResourceQuotaStatusByNamespace instantiates a new V1alpha2ResourceQuotaStatusByNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2ResourceQuotaStatusByNamespaceWithDefaults

`func NewV1alpha2ResourceQuotaStatusByNamespaceWithDefaults() *V1alpha2ResourceQuotaStatusByNamespace`

NewV1alpha2ResourceQuotaStatusByNamespaceWithDefaults instantiates a new V1alpha2ResourceQuotaStatusByNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHard

`func (o *V1alpha2ResourceQuotaStatusByNamespace) GetHard() map[string]ResourceQuantity`

GetHard returns the Hard field if non-nil, zero value otherwise.

### GetHardOk

`func (o *V1alpha2ResourceQuotaStatusByNamespace) GetHardOk() (*map[string]ResourceQuantity, bool)`

GetHardOk returns a tuple with the Hard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHard

`func (o *V1alpha2ResourceQuotaStatusByNamespace) SetHard(v map[string]ResourceQuantity)`

SetHard sets Hard field to given value.

### HasHard

`func (o *V1alpha2ResourceQuotaStatusByNamespace) HasHard() bool`

HasHard returns a boolean if a field has been set.

### GetNamespace

`func (o *V1alpha2ResourceQuotaStatusByNamespace) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1alpha2ResourceQuotaStatusByNamespace) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1alpha2ResourceQuotaStatusByNamespace) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetUsed

`func (o *V1alpha2ResourceQuotaStatusByNamespace) GetUsed() map[string]ResourceQuantity`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *V1alpha2ResourceQuotaStatusByNamespace) GetUsedOk() (*map[string]ResourceQuantity, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *V1alpha2ResourceQuotaStatusByNamespace) SetUsed(v map[string]ResourceQuantity)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *V1alpha2ResourceQuotaStatusByNamespace) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


