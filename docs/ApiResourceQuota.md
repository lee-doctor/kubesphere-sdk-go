# ApiResourceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**V1ResourceQuotaStatus**](V1ResourceQuotaStatus.md) |  | 
**Namespace** | **string** | namespace | 

## Methods

### NewApiResourceQuota

`func NewApiResourceQuota(data V1ResourceQuotaStatus, namespace string, ) *ApiResourceQuota`

NewApiResourceQuota instantiates a new ApiResourceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResourceQuotaWithDefaults

`func NewApiResourceQuotaWithDefaults() *ApiResourceQuota`

NewApiResourceQuotaWithDefaults instantiates a new ApiResourceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiResourceQuota) GetData() V1ResourceQuotaStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiResourceQuota) GetDataOk() (*V1ResourceQuotaStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiResourceQuota) SetData(v V1ResourceQuotaStatus)`

SetData sets Data field to given value.


### GetNamespace

`func (o *ApiResourceQuota) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiResourceQuota) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiResourceQuota) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


