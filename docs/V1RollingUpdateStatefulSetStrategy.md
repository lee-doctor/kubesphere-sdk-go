# V1RollingUpdateStatefulSetStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUnavailable** | Pointer to **string** | The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable. | [optional] 
**Partition** | Pointer to **int32** | Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0. | [optional] 

## Methods

### NewV1RollingUpdateStatefulSetStrategy

`func NewV1RollingUpdateStatefulSetStrategy() *V1RollingUpdateStatefulSetStrategy`

NewV1RollingUpdateStatefulSetStrategy instantiates a new V1RollingUpdateStatefulSetStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RollingUpdateStatefulSetStrategyWithDefaults

`func NewV1RollingUpdateStatefulSetStrategyWithDefaults() *V1RollingUpdateStatefulSetStrategy`

NewV1RollingUpdateStatefulSetStrategyWithDefaults instantiates a new V1RollingUpdateStatefulSetStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxUnavailable

`func (o *V1RollingUpdateStatefulSetStrategy) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *V1RollingUpdateStatefulSetStrategy) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *V1RollingUpdateStatefulSetStrategy) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *V1RollingUpdateStatefulSetStrategy) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetPartition

`func (o *V1RollingUpdateStatefulSetStrategy) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V1RollingUpdateStatefulSetStrategy) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V1RollingUpdateStatefulSetStrategy) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V1RollingUpdateStatefulSetStrategy) HasPartition() bool`

HasPartition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


