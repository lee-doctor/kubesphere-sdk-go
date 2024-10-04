# V1alpha2ResourceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | [**V1alpha2ResourceQuotaSpec**](V1alpha2ResourceQuotaSpec.md) |  | 
**Status** | Pointer to [**V1alpha2ResourceQuotaStatus**](V1alpha2ResourceQuotaStatus.md) |  | [optional] 

## Methods

### NewV1alpha2ResourceQuota

`func NewV1alpha2ResourceQuota(spec V1alpha2ResourceQuotaSpec, ) *V1alpha2ResourceQuota`

NewV1alpha2ResourceQuota instantiates a new V1alpha2ResourceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2ResourceQuotaWithDefaults

`func NewV1alpha2ResourceQuotaWithDefaults() *V1alpha2ResourceQuota`

NewV1alpha2ResourceQuotaWithDefaults instantiates a new V1alpha2ResourceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1alpha2ResourceQuota) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1alpha2ResourceQuota) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1alpha2ResourceQuota) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1alpha2ResourceQuota) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1alpha2ResourceQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1alpha2ResourceQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1alpha2ResourceQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1alpha2ResourceQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1alpha2ResourceQuota) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha2ResourceQuota) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha2ResourceQuota) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha2ResourceQuota) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1alpha2ResourceQuota) GetSpec() V1alpha2ResourceQuotaSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1alpha2ResourceQuota) GetSpecOk() (*V1alpha2ResourceQuotaSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1alpha2ResourceQuota) SetSpec(v V1alpha2ResourceQuotaSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *V1alpha2ResourceQuota) GetStatus() V1alpha2ResourceQuotaStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1alpha2ResourceQuota) GetStatusOk() (*V1alpha2ResourceQuotaStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1alpha2ResourceQuota) SetStatus(v V1alpha2ResourceQuotaStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1alpha2ResourceQuota) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


