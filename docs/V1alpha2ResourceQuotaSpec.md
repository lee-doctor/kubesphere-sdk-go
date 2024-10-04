# V1alpha2ResourceQuotaSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | [**V1ResourceQuotaSpec**](V1ResourceQuotaSpec.md) |  | 
**Selector** | **map[string]string** |  | 

## Methods

### NewV1alpha2ResourceQuotaSpec

`func NewV1alpha2ResourceQuotaSpec(quota V1ResourceQuotaSpec, selector map[string]string, ) *V1alpha2ResourceQuotaSpec`

NewV1alpha2ResourceQuotaSpec instantiates a new V1alpha2ResourceQuotaSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2ResourceQuotaSpecWithDefaults

`func NewV1alpha2ResourceQuotaSpecWithDefaults() *V1alpha2ResourceQuotaSpec`

NewV1alpha2ResourceQuotaSpecWithDefaults instantiates a new V1alpha2ResourceQuotaSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *V1alpha2ResourceQuotaSpec) GetQuota() V1ResourceQuotaSpec`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *V1alpha2ResourceQuotaSpec) GetQuotaOk() (*V1ResourceQuotaSpec, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *V1alpha2ResourceQuotaSpec) SetQuota(v V1ResourceQuotaSpec)`

SetQuota sets Quota field to given value.


### GetSelector

`func (o *V1alpha2ResourceQuotaSpec) GetSelector() map[string]string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1alpha2ResourceQuotaSpec) GetSelectorOk() (*map[string]string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1alpha2ResourceQuotaSpec) SetSelector(v map[string]string)`

SetSelector sets Selector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


