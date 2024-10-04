# V1alpha2ResourceQuotaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespaces** | [**[]V1alpha2ResourceQuotaStatusByNamespace**](V1alpha2ResourceQuotaStatusByNamespace.md) |  | 
**Total** | [**V1ResourceQuotaStatus**](V1ResourceQuotaStatus.md) |  | 

## Methods

### NewV1alpha2ResourceQuotaStatus

`func NewV1alpha2ResourceQuotaStatus(namespaces []V1alpha2ResourceQuotaStatusByNamespace, total V1ResourceQuotaStatus, ) *V1alpha2ResourceQuotaStatus`

NewV1alpha2ResourceQuotaStatus instantiates a new V1alpha2ResourceQuotaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2ResourceQuotaStatusWithDefaults

`func NewV1alpha2ResourceQuotaStatusWithDefaults() *V1alpha2ResourceQuotaStatus`

NewV1alpha2ResourceQuotaStatusWithDefaults instantiates a new V1alpha2ResourceQuotaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaces

`func (o *V1alpha2ResourceQuotaStatus) GetNamespaces() []V1alpha2ResourceQuotaStatusByNamespace`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *V1alpha2ResourceQuotaStatus) GetNamespacesOk() (*[]V1alpha2ResourceQuotaStatusByNamespace, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *V1alpha2ResourceQuotaStatus) SetNamespaces(v []V1alpha2ResourceQuotaStatusByNamespace)`

SetNamespaces sets Namespaces field to given value.


### GetTotal

`func (o *V1alpha2ResourceQuotaStatus) GetTotal() V1ResourceQuotaStatus`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1alpha2ResourceQuotaStatus) GetTotalOk() (*V1ResourceQuotaStatus, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1alpha2ResourceQuotaStatus) SetTotal(v V1ResourceQuotaStatus)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


